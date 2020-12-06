package main

/*@file main.go
@brief to determine which function to execute
*/

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
)

// LocStr ...
var LocStr = "Asia/Taipei"

// MaxDBConnLifetime ...
var MaxDBConnLifetime = 5

// IsProduction ... to set productin mode
var IsProduction = true // then set c value in the switch

// IsToScrape ...
var IsToScrape = true

// IsToRunFunc1 ...
var IsToRunFunc1 = 0

// MaxRiskScore ...
var MaxRiskScore = 90

//BlacCategories ...
var BlacCategories = []string{"hacker", "scam", "ransomware", "criminal", "money laundering", "terrorist financing"}

// EthereumNetwork ...
var EthereumNetwork = ""

//Cfg ...
var Cfg Config

var logW *log.Logger
var logI *log.Logger
var logE *log.Logger

// log1 ... to print logs
//var logE2 = logE.Println

// LogMode ...
var LogMode = 1

func log1(v ...interface{}) {
	switch LogMode {
	case 0:
		fmt.Println(v...)
	case 1:
		log.Println(v...)
	case 2:
		logI.Println(v...)
	case 3:
		logW.Println(v...)
	case 4:
		logE.Println(v...)
	default:
		fmt.Println(v...)
	}
}

// RwTokenPriceFake ...
var RwTokenPriceFake = float64(846.06242)

// RwTokenTotalLiquidityFake ...
var RwTokenTotalLiquidityFake = float64(4552)

func init() {
	//-------------------== Initial Conditions
	// append to the file or make a new file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	//log2 := log.Lshortfile

	logI = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logW = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logE = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// logI.Println("Something noteworthy happened")
	// logW.Println("There is something you should know about")
	// logE.Println("Something went wrong")
	//

	//-------------------==
	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		logE.Println("isEnvOk:", isEnvOk)
		return
	}

	//IsProduction = os.Getenv("ISPRODUCTION") == "1"
	//IsToScrape = os.Getenv("ISTOSCRAPE") == "1"
	IsToScrape = 1 == 1
	LogMode = 1 //dev: 0, log1: 1
	IsToRunFunc1 = 0
	IsProduction = 1 == 1

	err = cleanenv.ReadConfig("config.yml", &Cfg)
	if err != nil {
		logE.Println("reading config file failed")
		return
	}
	//logE.Println("Cfg:", Cfg)
	EthereumNetwork = Cfg.EthereumNetwork.Name
}

func main() {
	port := ":" + os.Getenv("SERVER_PORT")
	if port == ":" {
		logE.Println("SERVER_PORT in .env is empty")
		return
	}
	log1("port"+port, ", IsProduction:", IsProduction, ", IsToScrape:", IsToScrape)
	switch LogMode {
	case 0:
		log1("log in consoles")
	case 1:
		log1("use log file")
	case 2:
		log1("use log file with Info")
	case 3:
		log1("use log file with Warning")
	case 4:
		log1("use log file with Error")
	default:
		log1("log in consoles")
	}
	//-------------------== Routers
	router := mux.NewRouter()
	router.HandleFunc("/member", httpCreateUser).Methods("POST")
	router.HandleFunc("/member", httpReadUser).Methods("GET")

	router.HandleFunc("/reward", httpRewardC).Methods("POST")
	router.HandleFunc("/reward", httpRewardR).Methods("GET")
	router.HandleFunc("/reward", httpRewardD).Methods("DELETE")

	router.HandleFunc("/ariesapy", httpGetApyAries).Methods("GET")
	router.HandleFunc("/ariesapy", httpAriesU).Methods("PUT")

	router.HandleFunc("/cors", httpCorsGet).Methods("GET")
	router.HandleFunc("/cors", httpCorsPost).Methods("POST")
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/", root).Methods("GET")
	/*
		//protected routes
		router.HandleFunc("/protected",
			tokenVerifyMiddleWare(ctr.ProtectedEndpoint())).Methods("GET") // validate token > run middleware func
	*/

	//-------------------== test functions
	if IsToRunFunc1 == 1 {
		choice := 41
		var addrRewardsPool string
		switch choice {
		case 1: // mainnet
			addrRewardsPool = "0x9cd43309c9e122a13b466391babc5dec8be1e01e"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0x9cd43309c9e122a13b466391babc5dec8be1e01e
			curl 'localhost:3000/ariesapy?rewardspool=0x9cd43309C9E122A13b466391bABC5deC8bE1E01E' | jq
			*/
		case 11:
			addrRewardsPool = "0xd40cade3f71c20ba6fe940e431c890dc100e97d6"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0xd40cade3f71c20ba6fe940e431c890dc100e97d6
			curl 'localhost:3000/ariesapy?rewardspool=0xd40cade3f71c20ba6fe940e431c890dc100e97d6' | jq
			*/
		case 21:
			addrRewardsPool = "0xAC7DE028cCe2a99e9399aB0bE198Bc950994f50C"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0xAC7DE028cCe2a99e9399aB0bE198Bc950994f50C
			curl 'localhost:3000/ariesapy?rewardspool=0xAC7DE028cCe2a99e9399aB0bE198Bc950994f50C' | jq
			*/
		case 31:
			addrRewardsPool = "0x825241bA78700c11a4615523dF4B70F78C7384aa"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0x825241bA78700c11a4615523dF4B70F78C7384aa
			curl 'localhost:3000/ariesapy?rewardspool=0x825241bA78700c11a4615523dF4B70F78C7384aa' | jq
			*/
		case 41:
			addrRewardsPool = "0x8667D16150AcAA1FF19AcC5E5c64Bf0Ba1d551b3"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0x8667D16150AcAA1FF19AcC5E5c64Bf0Ba1d551b3
			curl 'localhost:3000/ariesapy?rewardspool=0x8667D16150AcAA1FF19AcC5E5c64Bf0Ba1d551b3' | jq
			*/
		default:
			addrRewardsPool = ""
		}
		print("addrRewardsPool:", addrRewardsPool)
		reqBody := ReqBody{
			RewardsPool: addrRewardsPool,
		}
		inputLambda := InputLambda{
			Body: reqBody,
		}
		outputLambdaPt, err := getApyAries(inputLambda)
		print("result:", outputLambdaPt)
		if err != nil || outputLambdaPt.Mesg != "ok" {
			print("\n====>>>> err@ writeRowX")
		}
		print("addrRewardsPool:", addrRewardsPool)
		print("IsToScrape:", IsToScrape)
	} /*
		-------==
		curl 'localhost:3000/ariesapy?rewardspool=0x825241bA78700c11a4615523dF4B70F78C7384aa' | jq
	*/

	if IsToRunFunc1 == 2 {
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
	if IsToRunFunc1 == 3 {
		// proto := "btc"
		// bcAddr := "12t9YDPgwueZ9NyMgw519p7AA8isjr6SMw"

		bcAddr := "0x304a554a310c7e546dfe434669c62820b7d83490" //dao hacker with no score
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			return
		}
	}
	if IsToRunFunc1 == 4 {
		reqBody := ReqBody{
			SourceName: "uniswap",
			Token0:     "afi",
			Token1:     "usdc",
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

	if IsToRunFunc1 == 5 {
		sourceURL := "https://info.uniswap.org/pair/0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891"
		loadingTime := 7
		IsToScrape = false
		regexpStr := `[-+]?[0-9]*\.?[0-9]+`
		ss, err := chromedpScraper(sourceURL, loadingTime, IsToScrape)
		if err != nil {
			log.Fatal(err)
		}
		print("scraper output:", ss)
		print("chromedpScraper is successful")
		for idx, v := range ss {
			print("idx", idx, ":", v)
			v2 := strings.Replace(v, ",", "", -1)
			out, err := regexp2FindInBtw(v2, regexpStr)
			if err != nil {
				print("err:", err)
			}
			print("out:", out)
		}
		//print("outerHTML1:", strings.TrimSpace(outerHTML1))
		//<div class="sc-bdVaJa KpMoH css-9on69b">$10.39</div>
	}
	if IsToRunFunc1 == 6 {
		//setupEthereum()
		rewardsPool := "0xd40cade3f71c20ba6fe940e431c890dc100e97d6"
		rewardRate, totalSupply, err := getRewardsCtrtValues(rewardsPool, "mainnet")
		if err != nil {
			print("setupEthereum failed")
			os.Exit(1)
		}
		print("rewardRate:", rewardRate, ", totalSupply:", totalSupply)
	}
	if IsToRunFunc1 == 7 {
		tokenPrice := float64(33.00001)
		mag := int64(100000)
		tokenPriceBI := float64ToBigInt(tokenPrice, mag)
		print("tokenPriceBI:", tokenPriceBI)

		maginfierBF := big.NewFloat(1000000)
		var ourAPYbn, _ = new(big.Int).SetString("547040", 10)
		print("ourAPYbn:", ourAPYbn)
		ourAPYf := new(big.Float).SetInt(ourAPYbn)
		print("ourAPYf:", ourAPYf)
		ourAPYbf := new(big.Float).Quo(ourAPYf, maginfierBF)
		print("ourAPYbf:", ourAPYbf)
	}

	//print("\nport"+port, ", IsProduction:", IsProduction, ", IsToScrape:", IsToScrape)
	log1("listening on", port)
	if IsToRunFunc1 == 0 {
		log.Fatal(http.ListenAndServe(port, router))
	}
}

/*
	//https://golang.org/pkg/net/http/#pkg-constants
	123 abc make http status code!!!
	respondWithError(w, http.StatusInternalServerError, errM)
	return

	w.Header().Set("Content-Type","application/json")
	respondWithJSON(w, "Update completed Successfully")

		case 1: // rinkeby
			addrRewardsPool = "0xbf76248d5e3bfd1d4dde4369fe6163289a0267f6"
			/*https://af-api.aries.financial/ariesapy?rewardspool=0xbf76248d5e3bfd1d4dde4369fe6163289a0267f6
			curl 'localhost:3000/ariesapy?rewardspool=0xbf76248d5e3bfd1d4dde4369fe6163289a0267f6' | jq
			


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

//--------------------==
	//router.HandleFunc("/member", httpUpdateUser).Methods("PUT")
	//router.HandleFunc("/member", httpDeleteUser).Methods("DELETE")

	// router.HandleFunc("/securities", httpCreateUsers).Methods("POST")
	// router.HandleFunc("/securities", httpReadUsers).Methods("GET")

		//router.HandleFunc("/vaultethaddr", httpVaultEthAddrR).Methods("GET")

	// router.HandleFunc("/aries10", httpApyC).Methods("POST")
	// router.HandleFunc("/aries10", httpApyR).Methods("GET")
	// router.HandleFunc("/aries10reset", httpApyReset).Methods("PUT")
	// router.HandleFunc("/aries10", httpApyU).Methods("PUT")

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
