package main

import (
	"math/big"
	"strings"
)

/*@file getApyAries.go
@brief getApyAries API
see getApyAries function description below

expect frontend to give rewardsPoolCtrt address. then use that to find the pool info, then call that Rewards contract to get rewardRate

@author
@date   2020-11-20
*/
func getApyAries(inputLambda InputLambda) (*OutputLambda, error) {
	log1("-----------== getApyAries()")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body
	addrRewardsPool := reqBody.RewardsPool
	// sourceName := reqBody.SourceName
	// token0 := reqBody.Token0
	// token1 := reqBody.Token1

	var isRewardPoolFound bool
	var mesg string
	var pool RewardsPool
	for _, item := range Cfg.RewardsPools {
		if strings.ToLower(item.RewardsCtrt) == strings.ToLower(addrRewardsPool) {
			if EthereumNetwork != item.Network {
				mesg = "reward contract network is different from current operating network"
				log1(mesg)
			} else {
				log1("found pool is on the same EthereumNetwork")
			}
			pool = item
			isRewardPoolFound = true
			break
		}
	}
	if !isRewardPoolFound {
		logE.Println("reward pool not found")
		return &OutputLambda{
			Code: "000104",
			Mesg: "reward pool not found",
		}, nil
	}
	log1("\n-----------== Rewards Pool has been found in config file. \npool:", pool)
	addrRewardsPool = pool.RewardsCtrt
	lpTokenPriceSource := pool.LpTokenPriceSource
	totalLiquiditySource := pool.TotalLiquiditySource
	rwTokenPriceSource := pool.RwTokenPriceSorce
	loadingTime := pool.LoadingTime
	decimalPlace := pool.DecimalPlace
	log1("addrRewardsPool:", addrRewardsPool)
	log1("lpTokenPriceSource:", lpTokenPriceSource)
	log1("totalLiquiditySource:", totalLiquiditySource)
	log1("rwTokenPriceSource:", rwTokenPriceSource)
	log1("decimalPlace:", decimalPlace)
	log1("loadingTime:", loadingTime)

	//---------------== Get RewardRate and totalSupply
	outputLambdaPt := &OutputLambda{}
	rewardRate, totalSupply, err := getRewardsCtrtValues(addrRewardsPool, pool.Network)
	if err != nil {
		logE.Println("getRewardsCtrtValues failed", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ getRewardsCtrtValues",
		}, nil
	}
	log1("rewardRate:", rewardRate, ", totalSupply:", totalSupply)
	if len(totalSupply.Bits()) == 0 {
		logE.Println("totalSupply is zero")
		return &OutputLambda{
			Code: "0",
			Mesg: "totalSupply is zero",
			Data: "NA",
		}, nil
	}

	ampLPint := 100000
	ampRWint := 100000
	//----------== RW Token ... Get rwToken Price
	rwTokenData, err := getTokenData(rwTokenPriceSource, loadingTime)
	if err != nil {
		logE.Println("err@ chromedpScraper rwTokenPrice:", err)
		log1("use fake data: tokenPrice =", RwTokenPriceFake, ", totalLiquidity = ", RwTokenTotalLiquidityFake)
		rwTokenData = PairData{RwTokenPriceFake, RwTokenTotalLiquidityFake, 0, 0, 0}
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "err@ chromedpScraper rwTokenPrice",
		// }, nil
	}
	rwTokenPrice := rwTokenData.Price // same for all pools
	rwTotalLiquidity := rwTokenData.TotalLiquidity
	ampRW64 := int64(ampRWint)
	log1("scraped rwTokenPrice:", rwTokenPrice, ", rwTotalLiquidity:", rwTotalLiquidity)
	ampLP64 := int64(ampLPint)
	rwTotalLiquidityBI := float64ToBigInt(rwTotalLiquidity, ampLP64)
	log1("rwTotalLiquidityBI:", rwTotalLiquidityBI)

	//----------== LP Token ...lpToken Price, Total Liquidity
	lpTokenData, err := getTokenData(lpTokenPriceSource, loadingTime)
	if err != nil {
		logE.Println("err@ chromedpScraper rwTokenPrice:", err)
		log1("use fake data: tokenPrice =", RwTokenPriceFake, ", totalLiquidity = ", RwTokenTotalLiquidityFake)
		lpTokenData = PairData{RwTokenPriceFake, RwTokenTotalLiquidityFake, 0, 0, 0}
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "err@ chromedpScraper lpTokenPrice",
		// }, nil
	}
	lpTokenPrice := lpTokenData.Price
	log1("scraped lpTokenPrice:", lpTokenPrice)


	log1("-----------==")
	log1("LIVE Network on", pool.Network)
	log1("Pool name:", pool.Name)
	log1("address of RewardsPool:", addrRewardsPool)
	var TVL, lpPriceBI, rwPriceBI *big.Int
	switch pool.ID {
	case "001", "002": //AFI Governance
		log1("use TVL = totalStakedAmount * lpTokenPrice")
		//lpTokenPrice = 33.0
		lpTokenPrice = rwTokenPrice
		log1("rwTokenPrice:", rwTokenPrice, ", lpTokenPrice:", lpTokenPrice)
		rwPriceBI = float64ToBigInt(rwTokenPrice, ampRW64)
		lpPriceBI = rwPriceBI
		log1("rwPriceBI:", rwPriceBI)
		log1("lpPriceBI:", lpPriceBI)
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)

	case "041": //afiDAI
		log1("use TVL = totalStakedAmount * lpTokenPrice")
		//rwTokenPrice = 33.0
		lpTokenPrice = 1.0
		log1("rwTokenPrice:", rwTokenPrice, ", lpTokenPrice:", lpTokenPrice)
		rwPriceBI = float64ToBigInt(rwTokenPrice, ampRW64)
		lpPriceBI = float64ToBigInt(lpTokenPrice, ampLP64)
		//log1("lpPriceBI:", lpPriceBI, ", rwPriceBI:", rwPriceBI)
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)

	case "011": // UniLP_USDC_AFI Pool
		log1("use TVL = uniswap totalLiquidity")
		base, isOk := new(big.Int).SetString("1000000000000000000", 10)
		if !isOk {
			logE.Println("making 1e18 bigInt failed")
			return &OutputLambda{
				Code: "000105",
				Mesg: "err@ initializing base 18 zeros as big int",
			}, nil
		}
		TVL = new(big.Int).Mul(base, rwTotalLiquidityBI)
		//rwTokenPrice = 33.0
		lpTokenPrice = 0
		log1("rwTokenPrice:", rwTokenPrice)
		rwPriceBI = float64ToBigInt(rwTokenPrice, ampRW64)

	case "021", "031": // afUSDC, afUSDT
		log1("pool ID =", pool.ID, "... uses 6 decimal places!")
		log1("use TVL = totalStakedAmount * lpTokenPrice * dpDif")
		//rwTokenPrice = 33.0
		lpTokenPrice = 1.0
		log1("rwTokenPrice:", rwTokenPrice, ", lpTokenPrice:", lpTokenPrice)
		rwPriceBI = float64ToBigInt(rwTokenPrice, ampRW64)
		lpPriceBI = float64ToBigInt(lpTokenPrice, ampLP64)
		//log1("lpPriceBI:", lpPriceBI, ", rwPriceBI:", rwPriceBI)
		TVL1 := new(big.Int).Mul(totalSupply, lpPriceBI)

		base, isOk := new(big.Int).SetString("1000000000000", 10)
		if !isOk {
			logE.Println("making 1e12 bigInt failed")
			return &OutputLambda{
				Code: "000105",
				Mesg: "err@ initializing base 6 zeros as big int",
			}, nil
		}
		TVL = new(big.Int).Mul(TVL1, base)

	default:
		logE.Println("err@ pool.ID not found")
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ pool.ID not found",
		}, nil
	} // ampLP64 injected into TVL

	if len(TVL.Bits()) == 0 {
		logE.Println("totalSupply is zero")
		return &OutputLambda{
			Code: "0",
			Mesg: "totalSupply is zero",
			Data: "NA",
		}, nil
	}

	secondsPerWk := big.NewInt(604800)
	weeksPerYear := big.NewInt(52)
	//	t := big.Int
	log1("rewardRate:", rewardRate)
	weeklyReward := new(big.Int).Mul(rewardRate, secondsPerWk)
	log1("weeklyReward:", weeklyReward)

	log1("rwTokenPrice:", rwTokenPrice)
	weeklyRewardPrice := new(big.Int).Mul(weeklyReward, rwPriceBI)
	log1("weeklyRewardPrice +5z:", weeklyRewardPrice)

	apt2int := 1000000
	amp2int64 := int64(apt2int)
	apt2f64 := float64(apt2int)
	amp2bi := big.NewInt(amp2int64)
	log1("amp2bi:", amp2bi)
	weeklyRewardPriceAmp := new(big.Int).Mul(weeklyRewardPrice, amp2bi)
	log1("weeklyRewardPriceAmp +11z:", weeklyRewardPriceAmp)

	weeklyROI := new(big.Int).Div(weeklyRewardPriceAmp, TVL)
	log1("weeklyROI:", weeklyROI)

	APY1 := new(big.Int).Mul(weeklyROI, weeksPerYear)

	// yearlyReward := new(big.Int).Mul(weeklyReward, weeksPerYear)
	// log1("yearlyReward:", yearlyReward)
	// yearlyPrice := new(big.Int).Mul(yearlyReward, rwPriceBI)
	//log1("yearlyPrice:", yearlyPrice)

	//yearlyPriceAmp := new(big.Int).Mul(yearlyPrice, amp2bi)
	//log1("yearlyPriceAmp:", yearlyPriceAmp)
	// APY1 := new(big.Int).Div(yearlyPriceAmp, TVL)

	log1("APY1:", APY1)

	APY2bf := new(big.Float).SetInt(APY1)
	log1("APY2bf:", APY2bf)

	ampBF := big.NewFloat(apt2f64)
	APY3bf := new(big.Float).Quo(APY2bf, ampBF)
	log1("APY3bf:", APY3bf)

	/*
		totalStakedAmountInPool for each pool = total iquidity per pair.totalStakedAmountInPool is the amount of LP token staked in that pool
		weeklyReward is the total weekly reward token number for that pool
	*/

	if err != nil {
		logE.Println("err exists")
		return &OutputLambda{
			Code: "000104",
			Mesg: "mesg here",
		}, nil
	}
	outputLambdaPt = &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: APY3bf,
	}
	log1("\ngetApyAries is successful")
	return outputLambdaPt, err
}

/*
curl -XPUT -d '{"token0":"afi","token1":"usdc","sourceName":"uniswap"}' 'localhost:3000/vaults/afi' | jq

*/
