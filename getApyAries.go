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
				// return &OutputLambda{
				// 	Code: "000104",
				// 	Mesg: mesg,
				// }, nil
			} else {
				log1("found pool is on the same EthereumNetwork")
			}
			pool = item
			isRewardPoolFound = true
			break
		}
	}
	if !isRewardPoolFound {
		log1("input token0, token1, sourceName are invalid")
		return &OutputLambda{
			Code: "000104",
			Mesg: "input token0, token1, sourceName are invalid",
		}, nil
	}
	log1("\n-----------== Rewards Pool has been found in config file. \npool:", pool)
	addrRewardsPool = pool.RewardsCtrt
	lpTokenPriceSource := pool.LpTokenPriceSource
	totalLiquiditySource := pool.TotalLiquiditySource
	rwTokenPriceSource := pool.RwTokenPriceSorce
	log1("addrRewardsPool:", addrRewardsPool)
	log1("lpTokenPriceSource:", lpTokenPriceSource)
	log1("totalLiquiditySource:", totalLiquiditySource)
	log1("rwTokenPriceSource:", rwTokenPriceSource)

	loadingTime := pool.LoadingTime
	log1("loadingTime:", loadingTime)

	//---------------== Get RewardRate and totalSupply
	outputLambdaPt := &OutputLambda{}
	rewardRate, totalSupply, err := getRewardsCtrtValues(addrRewardsPool, pool.Network)
	if err != nil {
		log1("getRewardsCtrtValues failed", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ getRewardsCtrtValues",
		}, nil
	}
	log1("rewardRate:", rewardRate, ", totalSupply:", totalSupply)
	if len(totalSupply.Bits()) == 0 {
		return &OutputLambda{
			Code: "0",
			Mesg: "totalSupply is zero",
			Data: "NA",
		}, nil
	}
	//-----------------== LP Token ...Token Price, Total Liquidity
	ampLPint := 100000
	ampRWint := 100000
	lpTokenData, err := getTokenData(lpTokenPriceSource, loadingTime)
	if err != nil {
		log1("err@ chromedpScraper lptokenPrice:", err)
		lpTokenData = PairData{33.01, 873.1, 0, 0, 0}
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "err@ chromedpScraper lpTokenPrice",
		// }, nil
	}

	lpTokenPrice := lpTokenData.Price
	lpTotalLiquidity := lpTokenData.TotalLiquidity
	log1("lpTokenPrice:", lpTokenPrice, ", lpTotalLiquidity:", lpTotalLiquidity)
	ampLP64 := int64(ampLPint)
	lpPriceBI := float64ToBigInt(lpTokenPrice, ampLP64)
	log1("lpPriceBI:", lpPriceBI)
	lpTotalLiquidityBI := float64ToBigInt(lpTotalLiquidity, ampLP64)
	log1("lpTotalLiquidityBI:", lpTotalLiquidityBI)

	//-----------------== RW Token ... Get Token Price
	rwTokenData, err := getTokenData(rwTokenPriceSource, loadingTime)
	if err != nil {
		log1("err@ chromedpScraper rwTokenPrice:", err)
		rwTokenData = PairData{33.02, 873.2, 0, 0, 0}
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "err@ chromedpScraper rwTokenPrice",
		// }, nil
	}
	rwTokenPrice := rwTokenData.Price
	//lpTotalLiquidity := rwTokenData.TotalLiquidity
	log1("rwTokenPrice:", rwTokenPrice)
	ampRW64 := int64(ampRWint)
	rwPriceBI := float64ToBigInt(rwTokenPrice, ampRW64)
	log1("rwPriceBI:", rwPriceBI)

	log1("-----------==\n")
	log1("LIVE \nNetwork:", pool.Network)
	log1("Pool name:", pool.Name)
	log1("address of RewardsPool:", addrRewardsPool)
	var TVL *big.Int
	base, isOk := new(big.Int).SetString("1000000000000000000", 10)
	if !isOk {
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ initializing base 18 zeros as big int",
		}, nil
	}
	switch pool.ID {
	case "001", "002": //AFI Governance Pool
		log1("use TVL = totalStakedAmount * lpTokenPrice")
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)
		rwPriceBI = lpPriceBI

	case "011": // UniLP_USDC_AFI Pool
		log1("use TVL = totalLiquidity")
		TVL = new(big.Int).Mul(base, lpTotalLiquidityBI)

	case "021", "031": // afUSDC, afUSDT pool
		log1("use TVL = totalStakedAmount * lpTokenPrice")
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)
		rwPriceBI = lpPriceBI

	default:
		log1("err@ pool.ID not found")
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ pool.ID not found",
		}, nil
	} // ampLP64 injected into TVL

	secondsPerWk := big.NewInt(604800)
	weeksPerYear := big.NewInt(52)
	//	t := big.Int
	log1("rewardRate:", rewardRate)
	weeklyReward := new(big.Int).Mul(rewardRate, secondsPerWk)
	log1("weeklyReward:", weeklyReward)

	yearlyReward := new(big.Int).Mul(weeklyReward, weeksPerYear)
	log1("yearlyReward:", yearlyReward)

	yearlyPrice := new(big.Int).Mul(yearlyReward, rwPriceBI)
	log1("yearlyPrice:", yearlyPrice)

	apt2int := 1000000
	amp2int64 := int64(apt2int)
	apt2f64 := float64(apt2int)
	amp2bi := big.NewInt(amp2int64)
	yearlyPriceAmp := new(big.Int).Mul(yearlyPrice, amp2bi)
	log1("amp2bi:", amp2bi)
	log1("yearlyPriceAmp:", yearlyPriceAmp)

	if len(TVL.Bits()) == 0 {
		return &OutputLambda{
			Code: "0",
			Mesg: "totalSupply is zero",
			Data: "NA",
		}, nil
	}
	APY1 := new(big.Int).Div(yearlyPriceAmp, TVL)
	log1("APY1:", APY1)

	APY2bf := new(big.Float).SetInt(APY1)
	log1("APY2bf:", APY2bf)

	ampBF := big.NewFloat(apt2f64)
	APY3bf := new(big.Float).Quo(APY2bf, ampBF)
	log1("APY3bf:", APY3bf)

	/*
		  totalStakedAmountInPool for each pool = total iquidity per pair.	totalStakedAmountInPool is the amount of LP token staked in that pool
					weeklyReward is the total weekly reward token number for that pool
	*/

	if err != nil {
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
