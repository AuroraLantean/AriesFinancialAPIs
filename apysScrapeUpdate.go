package main

/*@file apysScrapeUpdate.go
@brief apysScrapeUpdate API
see apysScrapeUpdate function description below

@author
@date   2020-11-11
*/
func apysScrapeUpdate(inputLambda InputLambda) (*OutputLambda, error) {
	print("-----------== apysScrapeUpdate()")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body
	sourceURL := reqBody.SourceURL
	// sourceURL := "https://stats.finance/yearn"
	// htmlPattern := reqBody.HTPMPattern
	// regexpStr := reqBody.RegexpStr

	var htmlPattern, regexpStr, dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		htmlPattern = ".MuiTable-root tbody tr"
		regexpStr = `/\s([^}]*)\%`
		dataName = "yearnFinance"
	default:
		print("sourceURL invalid")
		return &OutputLambda{
			Code: "000104",
			Mesg: "sourceURL invalid",
		}, nil
	}
	// "([a-z]+)"

	print("visiting", sourceURL)

	var ss []string
	if IsToScrape {
		ss1, err := collyScraper(sourceURL, htmlPattern)
		if err != nil {
			print("err@ collyScraper:", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ collyScraper",
			}, nil
		}
		ss = ss1
	} else {
		ss2, err := collyScraperFakeYFI1()
		if err != nil {
			print("err@ collyScraper:", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ collyScraper",
			}, nil
		}
		ss = ss2
	}

	//dump(ss)
	print(ss)
	print("-----------------== Vaults")
	apys := APYs{}
	for idx, target := range ss {
		if idx < 5 || idx > 15 {
			continue
		}
		print("------------==\n", target)
		apyN, err := regexp2FindInBtw(target, regexpStr)
		if err != nil {
			print("err:", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ regexp2FindInBtw",
			}, nil
		}
		print("apyN:", apyN)
		switch {
		case idx == 5:
			apys.WETH = apyN
		case idx == 6:
			apys.YFI = apyN
		case idx == 7:
			apys.CRV3 = apyN
		case idx == 8:
			apys.CRVY = apyN
		case idx == 9:
			apys.CRVBUSD = apyN
		case idx == 10:
			apys.CRVSBTC = apyN
		case idx == 11:
			apys.DAI = apyN
		case idx == 12:
			apys.TrueUSD = apyN
		case idx == 13:
			apys.USDC = apyN
		case idx == 14:
			apys.Gemini = apyN
		case idx == 15:
			apys.TetherUSD = apyN
		default:
			print("idx of APY not needed")
		}
	}
	//after looping through all table entries
	//print("-----------------== Delegated Vaults")
	print("------------==")
	dump("apys:", apys)

	// yearnFinance_week
	print("------------== write to db")
	inputLambda2 := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := apysUpdate(inputLambda2)

	print("\napysScrapeUpdate is successful")
	return outputLambdaPt, err
}
