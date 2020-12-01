package main

/*@file ariesU.go
@brief ariesU API
see ariesU function description below

@author
@date   2020-11-30
*/
func ariesU(inputLambda InputLambda) (*OutputLambda, error) {
	print("-----------== ariesU()")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body
	sourceName := reqBody.SourceName
	token0 := reqBody.Token0
	token1 := reqBody.Token1

	var loadingTime int
	var sourceURL string
	switch {
	// case token0 != "afi": print("here1")
	// case token1 != "usdc": print("here2")
	// case sourceName != "uniswap": print("here3")
	case token0 == "afi" && token1 == "usdc" && sourceName == "uniswap":
		print("matched1")
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

	print("\n------------== regexp2 parsing")
	pairData := PairData{}
	for idx, v := range ss {
		print("idx", idx, ":", v)
		out, err := regexp2FindInBtw(v, regexpStr)
		if err != nil {
			print("err:", err)
		}
		print("out:", out)
		switch {
		case idx == 0:
			pairData.Price = toFloat(out)
		case idx == 1:
			pairData.TotalLiquidity = toFloat(out)
		default:
			print("idx not needed")
		}
	}
	print("\npairData:")
	dump(pairData)

	print("------------== Get Rewards contract rewardRate")

	print("------------== Calculate APY")

	print("------------== write to db")

	//inputLambda2 := InputLambda{Body: reqBody}
	//outputLambdaPt, err := apysUpdate(inputLambda2)
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
	}
	print("\nariesU is successful")
	return outputLambdaPt, err
}
/*
curl -XPUT -d '{"token0":"afi","token1":"usdc","sourceName":"uniswap"}' 'localhost:3000/vaults/afi' | jq

*/