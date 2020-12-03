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
	print("-----------== getApyAries()")
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
				print(mesg)
				// return &OutputLambda{
				// 	Code: "000104",
				// 	Mesg: mesg,
				// }, nil
			} else {
				print("found pool is on the same EthereumNetwork")
			}
			pool = item
			isRewardPoolFound = true
			break
		}
	}
	if !isRewardPoolFound {
		print("input token0, token1, sourceName are invalid")
		return &OutputLambda{
			Code: "000104",
			Mesg: "input token0, token1, sourceName are invalid",
		}, nil
	}
	print("\n-----------== Rewards Pool has been found in config file. \npool:", pool)
	addrRewardsPool = pool.RewardsCtrt
	lpTokenPriceSource := pool.LpTokenPriceSource
	totalLiquiditySource := pool.TotalLiquiditySource
	rwTokenPriceSource := pool.RwTokenPriceSorce
	print("addrRewardsPool:", addrRewardsPool)
	print("lpTokenPriceSource:", lpTokenPriceSource)
	print("totalLiquiditySource:", totalLiquiditySource)
	print("rwTokenPriceSource:", rwTokenPriceSource)

	loadingTime := pool.LoadingTime
	print("loadingTime:", loadingTime)

	//---------------== Get RewardRate and totalSupply
	outputLambdaPt := &OutputLambda{}
	rewardRate, totalSupply, err := getRewardsCtrtValues(addrRewardsPool, pool.Network)
	if err != nil {
		print("getRewardsCtrtValues failed", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ getRewardsCtrtValues",
		}, nil
	}
	print("rewardRate:", rewardRate, ", totalSupply:", totalSupply)
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
		print("err@ chromedpScraper:", err)
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ chromedpScraper",
		}, nil
	}
	lpTokenPrice := lpTokenData.Price
	lpTotalLiquidity := lpTokenData.TotalLiquidity
	print("lpTokenPrice:", lpTokenPrice, ", lpTotalLiquidity:", lpTotalLiquidity)
	ampLP64 := int64(ampLPint)
	lpPriceBI := float64ToBigInt(lpTokenPrice, ampLP64)
	print("lpPriceBI:", lpPriceBI)
	lpTotalLiquidityBI := float64ToBigInt(lpTotalLiquidity, ampLP64)
	print("lpTotalLiquidityBI:", lpTotalLiquidityBI)

	//-----------------== RW Token ... Get Token Price
	rwTokenData, err := getTokenData(rwTokenPriceSource, loadingTime)
	if err != nil {
		print("err@ chromedpScraper:", err)
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ chromedpScraper",
		}, nil
	}
	rwTokenPrice := rwTokenData.Price
	//lpTotalLiquidity := rwTokenData.TotalLiquidity
	print("rwTokenPrice:", rwTokenPrice)
	ampRW64 := int64(ampRWint)
	rwPriceBI := float64ToBigInt(rwTokenPrice, ampRW64)
	print("rwPriceBI:", rwPriceBI)

	print("-----------==\n")
	print("LIVE")
	print("Network:", pool.Network)
	print("Pool name:", pool.Name)
	print("address of RewardsPool:", addrRewardsPool)
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
		print("use TVL = totalStakedAmount * lpTokenPrice")
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)
		rwPriceBI = lpPriceBI

	case "011": // UniLP_USDC_AFI Pool
		print("use TVL = totalLiquidity")
		TVL = new(big.Int).Mul(base, lpTotalLiquidityBI)

	case "021", "031": // afUSDC, afUSDT pool
		print("use TVL = totalStakedAmount * lpTokenPrice")
		TVL = new(big.Int).Mul(totalSupply, lpPriceBI)
		rwPriceBI = lpPriceBI

	default:
		print("err@ pool.ID not found")
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ pool.ID not found",
		}, nil
	} // ampLP64 injected into TVL

	secondsPerWk := big.NewInt(604800)
	weeksPerYear := big.NewInt(52)
	//	t := big.Int
	print("rewardRate:", rewardRate)
	weeklyReward := new(big.Int).Mul(rewardRate, secondsPerWk)
	print("weeklyReward:", weeklyReward)

	yearlyReward := new(big.Int).Mul(weeklyReward, weeksPerYear)
	print("yearlyReward:", yearlyReward)

	yearlyPrice := new(big.Int).Mul(yearlyReward, rwPriceBI)
	print("yearlyPrice:", yearlyPrice)

	apt2int := 1000000
	amp2int64 := int64(apt2int)
	apt2f64 := float64(apt2int)
	amp2bi := big.NewInt(amp2int64)
	yearlyPriceAmp := new(big.Int).Mul(yearlyPrice, amp2bi)
	print("amp2bi:", amp2bi)
	print("yearlyPriceAmp:", yearlyPriceAmp)

	if len(TVL.Bits()) == 0 {
		return &OutputLambda{
			Code: "0",
			Mesg: "totalSupply is zero",
			Data: "NA",
		}, nil
	}
	APY1 := new(big.Int).Div(yearlyPriceAmp, TVL)
	print("APY1:", APY1)

	APY2bf := new(big.Float).SetInt(APY1)
	print("APY2bf:", APY2bf)

	ampBF := big.NewFloat(apt2f64)
	APY3bf := new(big.Float).Quo(APY2bf, ampBF)
	print("APY3bf:", APY3bf)

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
	print("\ngetApyAries is successful")
	return outputLambdaPt, err
}

/*
curl -XPUT -d '{"token0":"afi","token1":"usdc","sourceName":"uniswap"}' 'localhost:3000/vaults/afi' | jq

*/
