package main

/*@file main.go
@brief to determine which function to execute
*/

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// LocStr ...
var LocStr = "Asia/Taipei"

// MaxDBConnLifetime ...
var MaxDBConnLifetime = 5

// IsProduction ... to set productin mode
var IsProduction = true // then set c value in the switch

// IsToScrape ...
var IsToScrape = true

// MaxRiskScore ...
var MaxRiskScore = 90

//BlacCategories ...
var BlacCategories = []string{"hacker", "scam", "ransomware", "criminal", "money laundering", "terrorist financing"}

func main() {
	//-------------------== Initial Conditions
	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		print(isEnvOk)
		return
	}

	port := ":" + os.Getenv("SERVER_PORT")
	if port == ":" {
		print("SERVER_PORT in .env is empty")
		return
	}

	IsProduction = os.Getenv("ISPRODUCTION") == "1"
	IsToScrape = os.Getenv("ISTOSCRAPE") == "1"
	print("port"+port, ", IsProduction:", IsProduction, ", IsToScrape:", IsToScrape)

	//-------------------== Routers
	router := mux.NewRouter()
	router.HandleFunc("/member", httpCreateUser).Methods("POST")
	router.HandleFunc("/member", httpReadUser).Methods("GET")
	router.HandleFunc("/member", httpUpdateUser).Methods("PUT")
	router.HandleFunc("/member", httpDeleteUser).Methods("DELETE")

	// router.HandleFunc("/securities", httpCreateUsers).Methods("POST")
	// router.HandleFunc("/securities", httpReadUsers).Methods("GET")

	router.HandleFunc("/reward", httpRewardC).Methods("POST")
	router.HandleFunc("/reward", httpRewardR).Methods("GET")
	router.HandleFunc("/reward", httpRewardD).Methods("DELETE")

	//router.HandleFunc("/vaultethaddr", httpVaultEthAddrR).Methods("GET")

	// router.HandleFunc("/aries10", httpApyC).Methods("POST")
	// router.HandleFunc("/aries10", httpApyR).Methods("GET")
	// router.HandleFunc("/aries10reset", httpApyReset).Methods("PUT")
	// router.HandleFunc("/aries10", httpApyU).Methods("PUT")

	router.HandleFunc("/ariesapy", httpAriesR).Methods("GET")
	router.HandleFunc("/ariesapy", httpAriesU).Methods("PUT")

	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/", root).Methods("GET")
	/*
		//protected routes
		router.HandleFunc("/protected",
			tokenVerifyMiddleWare(ctr.ProtectedEndpoint())).Methods("GET") // validate token > run middleware func
	*/

	if 99 == 0 {
		dataName := "AFI"
		subgraphName := "uniswapV2"
		graphqlOut, mesg := DoCallGraphQLAPI(dataName, subgraphName)
		if mesg != "ok" {
			print(mesg)
			return
		}
		pairData := graphqlOut.Data.UniSwapPairData
		token1Price := pairData.Token1Price
		uniSwapToken1 := pairData.UniSwapToken1
		totalLiquitidy1 := uniSwapToken1.TotalLiquidity
		print("token1Price:", token1Price, ", totalLiquitidy1:", totalLiquitidy1)
		bitSize := 64
		if f, err := strconv.ParseFloat(token1Price, bitSize); err == nil {
			print("float type:", f) // bitSize is 32 for float32 convertible, 64 for float64
		} else {
			print(err)
		}

	}
	if 99 == 1 {
		// proto := "btc"
		// bcAddr := "12t9YDPgwueZ9NyMgw519p7AA8isjr6SMw"

		bcAddr := "0x304a554a310c7e546dfe434669c62820b7d83490" //dao hacker with no score
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			return
		}
	}
	if 0 == 2 {
		reqBody := ReqBody{
			SourceName: "uniswap",
			Token0: "afi",
			Token1: "usdc",
		}
		inputLambda := InputLambda{
			Body: reqBody,
		}
		outputLambdaPt, err := ariesU(inputLambda)
		print("result:", outputLambdaPt)
		if err != nil || outputLambdaPt.Mesg != "ok" {
			print("\n====>>>> err@ writeRowX")
		}
		print("IsToScrape:", IsToScrape)
	}

	if 99 == 5 {
		sourceURL := "https://info.uniswap.org/pair/0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891"
		loadingTime := 7
		regexpStr := `[-+]?[0-9]*\.?[0-9]+`
		//ss, err := chromedpScraper(sourceURL, loadingTime)
		ss, err := chromedpScraperFake(sourceURL, loadingTime)
		if err != nil {
			log.Fatal(err)
		}
		print("scraper output:", ss)
		print("chromedpScraper is successful")
		for idx, v := range ss {
			print("idx", idx, ":", v)
			out, err := regexp2FindInBtw(v, regexpStr)
			if err != nil {
				print("err:", err)
			}
			print("out:", out)
		}
		//print("outerHTML1:", strings.TrimSpace(outerHTML1))
		//<div class="sc-bdVaJa KpMoH css-9on69b">$10.39</div>
	}
	if 1 == 2 {
		preEthereum();
	}

	print("listening on", port)
	log.Fatal(http.ListenAndServe(port, router))
}

/*
	//https://golang.org/pkg/net/http/#pkg-constants
	123 abc make http status code!!!
	respondWithError(w, http.StatusInternalServerError, errM)
	return

	w.Header().Set("Content-Type","application/json")
	respondWithJSON(w, "Update completed Successfully")



https://stats.finance/robots.txt ... ok

//--------------------== APY
ApyC
curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/aries10' | jq

ApyR
curl 'localhost:3000/aries10' | jq

ApyU
curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/aries10' | jq


Write
curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpWriteRow'

Fetch+Update
curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/update'

//--------------------==
curl 'localhost:3000/ping'
curl 'localhost:3000'
curl 'localhost:3000/anChain'

//---------- VaultEthAddr
VaultEthAddr R
curl 'localhost:3000/vaultethaddr?userID=1&vaultID=1' | jq

-H "Content-type: application/json"

UsersC
curl -XPOST -d ' {"ethereumAddrs":["0xD118CDb869B4DA6cE2bb5c47306789eA0f5A0024","0x8Db1535f716e9cA763bFaad5896c237c2c83449c","0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E"]}' 'localhost:3000/securities' | jq

-H "Content-type: application/json"


//--------------------== Deployed Domain
https://api.aries.financial/aries10 ... done by 0x48

https://api.aries.financial/
https://api.aries.financial/ping

//--------------------== future use
Read:
curl -v 'https://api.aries.financial/aries1?sourceURL=https://stats.finance/yearn&perfPeriod=week'

https://api.aries.financial/aries2?sourceURL=https://stats.finance/yearn&perfPeriod=week
http://localhost:3000/aries10?sourceURL=https://stats.finance/yearn&perfPeriod=week

curl -v 'localhost:3000/aries3?sourceURL=https://stats.finance/yearn&perfPeriod=week'


* MIT License
* ===========
*
* Copyright (c) 2020 Synthetix
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*/
