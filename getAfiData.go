package main

import (
	"strings"
)

/*@file getAfiData.go
@brief getAfiData API
see getAfiData function description below

expect frontend to give rewardsPoolCtrt address. then use that to find the pool info, then call that Rewards contract to get rewardRate

@author
@date   2020-12-07
*/
func getAfiData(inputLambda InputLambda) (*OutputLambda, error) {
	log1("-----------== getAfiData()")
	dump("inputLambda.Body:", inputLambda.Body)
	//reqBody := inputLambda.Body
	//addrRewardsPool := reqBody.RewardsPool

	var addrRewardsPool string
	var isRewardPoolFound bool
	var mesg string
	var pool RewardsPool
	for _, item := range Cfg.RewardsPools {
		if strings.ToLower(item.ID) == "002" {
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

	//ampRWint := 100000
	//----------== RW Token ... Get rwToken Price
	rwTokenData, err1, err2 := getTokenData(rwTokenPriceSource, loadingTime)
	if err1 != nil || err2 != nil {
		logE.Println("err@ doChromedpAndRegexp getting rwTokenPrice. err1:", err1, ", err2:", err2)
		log1("use fake data: rwTokenPrice =", RwTokenPriceFake, ", rwTotalLiquidity = ", RwTokenTotalLiquidityFake)
		rwTokenData = PairData{RwTokenPriceFake, RwTokenTotalLiquidityFake, 0, 0, 0}
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "err@ chromedpScraper rwTokenPrice",
		// }, nil
	}
	rwTokenPrice := rwTokenData.Price // same for all pools
	rwTotalLiquidity := rwTokenData.TotalLiquidity
	log1("scraped rwTokenPrice:", rwTokenPrice, ", rwTotalLiquidity:", rwTotalLiquidity)

	outputLambdaPt := &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: PairData{
			Price: rwTokenPrice,
			TotalLiquidity: rwTotalLiquidity,
		},
	}
	log1("\ngetAfiData is successful")
	return outputLambdaPt, nil
}

/*
curl -XPUT -d '{"token0":"afi","token1":"usdc","sourceName":"uniswap"}' 'localhost:3000/vaults/afi' | jq

*/
