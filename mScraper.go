package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/gocolly/colly/v2"
)

/*@file scraper.go
@brief scraper API
see scraper function description below

@author
@date   2020-11-04
*/

//curl localhost:8080 -v -XPOST -d '{"xyz1":"john"}' | jq

// scraper ...
/*@brief scraper function to ...
@param out: Success or failure
@param  in: account
*/
func collyScraper(URL string, pattern string) ([]string, error) {
	log1("---------------== collyScraper()")
	//Create a new collector which will be in charge of collect the data from HTML
	c := colly.NewCollector()

	//Slices to store the data
	var response []string

	c.OnRequest(func(r *colly.Request) {
		log1("Visiting", r.URL)
	})
	// c.OnResponseHeaders(func(r *colly.Response) {
	// 	log1("Visited", r.Request.URL)
	// })

	//<table class="MuiTable-root">
	//pattern := "a[href]" //anchor tag with href
	c.OnHTML(pattern, func(e *colly.HTMLElement) {
		data1 := e.Text
		//data1 := e.Request.AbsoluteURL(e.Attr("href"))
		if data1 != "" {
			response = append(response, data1)
		}
	})
	/*
		c.OnHTML("#currencies-all tbody tr", func(e *colly.HTMLElement) {
			writer.Write([]string{
				e.ChildText(".currency-name-container"),
				e.ChildText(".col-symbol"),
				e.ChildAttr("a.price", "data-usd"),
	*/

	c.OnError(func(r *colly.Response, err error) {
		log1("c.OnError(): Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//Command to visit the website
	c.Visit(URL)

	return response, nil
}

/*
func fetchCoinMarketCap() ([]byte, error) {
	// Write CSV header
	//writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	// Find the review items
	// doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	band := s.Find("a").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	// })

	c.OnHTML("#currencies-all tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".currency-name-container"),
			e.ChildText(".col-symbol"),
			e.ChildAttr("a.price", "data-usd"),
			e.ChildAttr("a.volume", "data-usd"),
			e.ChildAttr(".market-cap", "data-usd"),
			e.ChildText(".percent-1h"),
			e.ChildText(".percent-24h"),
			e.ChildText(".percent-7d"),
		})
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
	b, err := json.Marshal(response)
	return b, err
}*/

func chromedpScraper(targetURL string, loadingTime int, isToScrape bool) ([]string, error) {
	log1("---------------== chromedpScraper()")
	if isToScrape {
		ctx, cancel := chromedp.NewContext(
			context.Background(),
			chromedp.WithLogf(log.Printf),
		)
		defer cancel()

		ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		var text1, text2 string
		log1("run chromedp")
		err := chromedp.Run(ctx,
			chromedp.Navigate(targetURL),
			// wait for element is visible(loaded)
			chromedp.WaitVisible(`body > #root`),
			chromedp.ActionFunc(func(context.Context) error {
				log1("waiting x seconds...")
				return nil
			}),
			chromedp.Sleep(time.Duration(loadingTime)*time.Second),
			//chromedp.OuterHTML(`#center .sc-bdVaJa.KpMoH.css-9on69b`, &outerHTML1),
			chromedp.Text(`#center .sc-bdVaJa.KpMoH.css-9on69b`, &text1),
			chromedp.Text(`#center .sc-bdVaJa.KpMoH.css-flugrv`, &text2),
		)
		/*chrome browser: copy selector
		#center > div > div.sc-kkGfuU.cDHqtk > div.sc-iQNlJl.fDoXeo > div > div.sc-eTuwsz.bFSUTM > div:nth-child(1) > div > div.sc-ifAKCX.sc-bZQynM.sc-gzVnrw.hatkAI > div.sc-bdVaJa.KpMoH.css-9on69b

		#center > div > div.sc-kkGfuU.cDHqtk > div.sc-iQNlJl.fDoXeo > div > div.sc-ifAKCX.sc-bZQynM.sc-dnqmqq.feLJDb > div:nth-child(1) > div > div.sc-bdVaJa.KpMoH.css-flugrv
		*/
		log1("text1:", text1, ", text2:", text2)
		if len(text2) > 2 {
			text2 = strings.TrimSpace(text2)[2:]
		}
		log1("chromedpScraper is successful")
		return []string{strings.TrimSpace(text1), text2}, err
	}
	log1("isToScrape is false")
	return []string{"$1022.03", "AFI = 33.00001 USDC"}, nil
}

func doChromedpAndRegexp(tokenPriceSource string, loadingTime int) (PairData, error) {
	log1("-----------== doChromedpAndRegexp()")
	regexpStr := `[-+]?[0-9]*\.?[0-9]+`
	pairData := PairData{}
	var ss []string
	var err error
	if tokenPriceSource == "" || loadingTime <= 0 {
		logE.Println("input invalid")
		return pairData, nil
	}
	ss, err = chromedpScraper(tokenPriceSource, loadingTime, IsToScrape)
	if err != nil {
		return pairData, err
	}

	log1("scraper output:", ss)
	pairData, err = doregexp2FindInBtw(ss, regexpStr)
	if err != nil {
		return pairData, err
	}
	return pairData, nil
}

func getTokenData(tokenPriceSource string, loadingTime int) (PairData, error) {
	log1("-----------== getTokenPairData()")
	var tokenPrice float64 = 1.0
	var totalLiquidity float64 = 4020.01

	if tokenPriceSource == "" {
		log1("no tokenPriceSource... use default tokenPrice and totalLiquidity...")
		return PairData{
			Price:          tokenPrice,
			TotalLiquidity: totalLiquidity,
		}, nil
	}
	if loadingTime <= 0 {
		return PairData{}, errors.New("loading time invalid")
	}
	pairData, err := doChromedpAndRegexp(tokenPriceSource, loadingTime)
	return pairData, err
}

func collyScraperFakeYFI1() ([]string, error) {
	print("---------------== collyScraperFakeYFI1()")
	return []string{" total$501,676,203", " vaults$202,552,207", " delegated vaults$13,858,140", " y.curve.fi$131,877,122", " busd.curve.fi$153,388,734",
		" Wrapped Ether (WETH)MKRVaultDAIDelegate3.13% / 1.16%55,850 WETH ($20,921,868)55,850 WETH ($20,921,868)",
		" yearn.finance (YFI)YFIGovernance15.32% / 10.28%2,132 YFI ($20,735,064)2,110 YFI ($20,526,738)",
		" curve.fi/3pool (3Crv)Curve3CrvVoterProxy10.61% / 12.60%17,844,285 3Crv17,842,190 3Crv",
		" curve.fi/y (yCRV)CurveYVoterProxy8.30% / 9.64%69,454,827 yCRV ($73,531,547)68,025,364 yCRV ($72,018,181)",
		" curve.fi/busd (crvBUSD)CurveBUSDVoterProxy20.96% / 18.30%14,336,686 crvBUSD ($15,210,865)12,588,041 crvBUSD ($13,355,597)",
		" curve.fi/sbtc (crvSBTC)CurveBTCVoterProxy10.91% / 3.27%946 crvSBTC ($12,753,473)920 crvSBTC ($12,408,102)",
		" Dai Stablecoin (DAI)DAICurve5.76% / 8.78%38,696,632 DAI ($39,083,599)18,679,534 DAI ($18,866,329)",
		" TrueUSD (TUSD)TUSDCurve2.24% / 4.56%7,563,217 TUSD ($7,551,215)1,514,076 TUSD ($1,511,376)",
		" USD Coin (USDC)DForceUSDC6.82% / 6.37%11,159,705 USDC ($11,144,852)11,159,706 USDC ($11,159,706)",
		" Gemini dollar (GUSD)CurveGUSDProxy0.01% / 0.02%423,390 GUSD0 GUSD",
		" Tether USD (USDT)DForceUSDT5.29% / 5.10%1,620,058 USDT ($1,619,724)1,619,647 USDT ($1,619,313)",
		" Aave Interest bearing LINK (aLINK)VaultUSDC2,002,930 aLINK ($20,469,948)6,950,018 USDC ($6,950,018) ChainLink Token (LINK)VaultUSDC33,061 LINK ($338,210)6,950,018 USDC ($6,950,018) Dai Stablecoin (DAI)2.20% / 3.66% / 3.95%7,702,227.27 DAI ($7,779,250)4,397,790.69 DAI  in Aave USD Coin (USDC)2.94% / 5.98% / 2.92%50,794,500.44 USDC ($50,726,893)50,123,808.84 USDC  in Aave Tether USD (USDT)3.28% / 3.27% / 2.68%34,294,024.75 USDT ($34,286,960)34,176,729.34 USDT  in Aave TrueUSD (TUSD)2.04% / 1.84% / 0.83%38,887,508.61 TUSD ($38,825,794)38,684,573.84 TUSD  in Aave Synth sUSD (sUSD)13.18% / 8.17% / 4.74%237,334.53 sUSD ($231,561)237,334.53 sUSD  in Aave Wrapped BTC (WBTC)0.00% / 23.81% / 5.07%1.99 WBTC ($26,664)0.44 WBTC  in Compoundnext: Aave Dai Stablecoin (DAI)3.52% / 5.07% / 3.82%12,023,133.74 DAI ($12,143,365)11,883,106.42 DAI  in Aave USD Coin (USDC)2.93% / 6.64% / 2.83%50,838,581.18 USDC ($50,770,915)50,447,360.43 USDC  in Aave Tether USD (USDT)3.25% / 3.75% / 2.61%39,740,602.53 USDT ($39,732,416)38,631,509.14 USDT  in Aave Binance USD (BUSD)0.32% / 0.27% / 1.89%50,749,193.61 BUSD ($50,742,038)50,423,841.93 BUSD  in Aave"}, nil
}
