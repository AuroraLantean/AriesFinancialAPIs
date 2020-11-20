package main

/*@file main.go
@brief to determine which function to execute
*/

import (
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/user", httpUserC).Methods("POST")
	router.HandleFunc("/user", httpUserR).Methods("GET")
	router.HandleFunc("/user", httpUserU).Methods("PUT")
	router.HandleFunc("/user", httpUserD).Methods("DELETE")

	router.HandleFunc("/users", httpUsersC).Methods("POST")
	router.HandleFunc("/users", httpUsersR).Methods("GET")

	router.HandleFunc("/reward", httpRewardC).Methods("POST")
	router.HandleFunc("/reward", httpRewardR).Methods("GET")
	router.HandleFunc("/reward", httpRewardD).Methods("DELETE")

	router.HandleFunc("/vaultethaddr", httpVaultEthAddrR).Methods("GET")

	router.HandleFunc("/vaults/apy", httpApyC).Methods("POST")
	router.HandleFunc("/vaults/apyreset", httpApyReset).Methods("PUT")
	router.HandleFunc("/vaults/apy", httpApyU).Methods("PUT")
	//router.HandleFunc("/httpApyD", httpRowD).Methods("DELETE")

	router.HandleFunc("/vaults/apy", httpApyR).Methods("GET")
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/", root).Methods("GET")
	/*
		//protected routes
		router.HandleFunc("/protected",
			tokenVerifyMiddleWare(ctr.ProtectedEndpoint())).Methods("GET") // validate token > run middleware func
	*/

	if 1 == 2 {
		// proto := "btc"
		// bcAddr := "12t9YDPgwueZ9NyMgw519p7AA8isjr6SMw"

		bcAddr := "0x304a554a310c7e546dfe434669c62820b7d83490" //dao hacker with no score
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			return
		}
	}

	print("listening on", port)
	log.Fatal(http.ListenAndServe(port, router))
}

/*
https://stats.finance/robots.txt ... ok

//--------------------== APY
ApyC
curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/vaults/apy' | jq

ApyR
curl 'localhost:3000/vaults/apy' | jq

ApyU
curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/vaults/apy' | jq


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
curl -XPOST -d ' {"ethereumAddrs":["0xD118CDb869B4DA6cE2bb5c47306789eA0f5A0024","0x8Db1535f716e9cA763bFaad5896c237c2c83449c","0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E"]}' 'localhost:3000/users' | jq

-H "Content-type: application/json" 


//--------------------== Deployed Domain
https://api.aries.financial/vaults/apy
https://api.aries.financial/
https://api.aries.financial/ping

//--------------------== future use
Read:
curl -v 'https://api.aries.financial/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week'

https://api.aries.financial/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week
http://localhost:3000/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week

curl -v 'localhost:3000/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week'

*/
