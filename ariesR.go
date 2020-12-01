package main

/*@file tokenDataByChromedpR.go
@brief tokenDataByChromedpR API
see tokenDataByChromedpR function description below

@author
@date   2020-11-20
*/
func tokenDataByChromedpR(inputLambda InputLambda) (*OutputLambda, error) {
	print("-----------== tokenDataByChromedpR()")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body
	sourceName := reqBody.SourceName
	token0 := reqBody.Token0
	token1 := reqBody.Token1

	var loadingTime int
	var sourceURL string
	switch {
	case token0 == "afi" && token1 == "usdc" && sourceName == "uniswap":
		sourceURL = "https://info.uniswap.org/pair/0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891"
		print("here1")
		loadingTime = 8
	default:
		print("input token0, token1, sourceName are invalid")
		return &OutputLambda{
			Code: "000104",
			Mesg: "input token0, token1, sourceName are invalid",
		}, nil
	}
	print("visiting", sourceURL, ", loadingTime:", loadingTime)

	regexpStr := `[-+]?[0-9]*\.?[0-9]+`
	print("regexpStr:", regexpStr)

	var err error
	var ss []string
	if IsToScrape {
		ss, err = chromedpScraper(sourceURL, loadingTime)
	} else {
		ss, err = chromedpScraperFake(sourceURL, loadingTime)
	}
	if err != nil {
		print("err@ chromedpScraper:", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ chromedpScraper",
		}, nil
	}
	print("scraper output:", ss)
	print("chromedpScraper is successful")

	print("------------== calculate APY")
	tokenPrice := toFloat(ss[0])
	pairTotalLiquidity := toFloat(ss[1])
	totalValueLocked := pairTotalLiquidity * tokenPrice
	print("totalValueLocked:", totalValueLocked)
	/*
  totalStakedAmountInPool for each pool = total iquidity per pair
	TVL = totalStakedAmountInPool* AFIPrice
	    AFIWeeklyROI = (weekly_reward * 1.01 / TVL) * 100
	    apy = YFIWeeklyROI * 52
	totalStakedAmountInPool is the amount of LP token staked in that pool
	weekly_reward is the total weekly reward token number for that pool
	*/

	outputLambdaPt := &OutputLambda{}
	if err != nil {
		outputLambdaPt = &OutputLambda{
			Code: "000104",
			Mesg: "mesg here",
		}
	}
	outputLambdaPt = &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: 0,//APY
	}
	print("\ntokenDataByChromedpR is successful")
	return outputLambdaPt, err
}

/*
curl -XPUT -d '{"token0":"afi","token1":"usdc","sourceName":"uniswap"}' 'localhost:3000/vaults/afi' | jq

*/
