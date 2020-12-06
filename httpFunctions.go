package main

/*@file
@brief to determine which function to execute
*/

import (
	"encoding/json"
	"net/http"
)

/*
curl -s 'http://127.0.0.1:3000/ping'
curl -s 'http://127.0.0.1:3000/YFIStats?url=https://stats.finance/yearn'

curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/vaults/apy' | jq
*/

// httpApyU ...
func httpApyU(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpApyU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := apysUpdateColly(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	log1("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}

	// parse our response slice into JSON format
	//b, err := json.Marshal(ss)
	// if err != nil {
	// 	log1("err:", err)
	// }
	//w.Write(b)
}

// httpAriesU ...
func httpAriesU(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpAriesU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := ariesU(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	log1("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpGetApyAries ...
func httpGetApyAries(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpGetApyAries")
	RewardsPool := r.URL.Query().Get("rewardspool")
	reqBody := ReqBody{RewardsPool: RewardsPool}
	log1("over to lambda function")
	log1("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := getApyAries(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	log1("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

func httpCorsGet(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpCorsGet")
	outputLambdaPt := &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: "123.4567",
	}
	log1("result:", outputLambdaPt)
	w.Header().Set("Content-Type", "application/json")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  
	//json.NewEncoder(w).Encode("OKOK")
	err := json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

func httpCorsPost(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpCorsPost")
	outputLambdaPt := &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: "123.4567",
	}
	log1("result:", outputLambdaPt)
	w.Header().Set("Content-Type", "application/json")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}


/*
curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpApyReset'
| jq
*/
// httpApyReset ...
func httpApyReset(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpApyReset")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)
	sourceURL := reqBody.SourceURL

	var dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		dataName = "yearnFinance"
	default:
		logE.Println("sourceURL invalid:")
		return
	}

	apys := APYs{WETH: "0.01", AFI: "1.11", YFI: "2.22", CRV3: "3.33", CRVY: "4.44", CRVBUSD: "5.55", CRVSBTC: "6.66", DAI: "7.77", TrueUSD: "8.88", USDC: "9.99", Gemini: "10.00", TetherUSD: "11.11"}
	log1("------------== write to db")
	inputLambda := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := apysUpdate(inputLambda)
	log1("result:", outputLambdaPt)
	outputLambdaPt.Data = apys
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	log1("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}

}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/httpUserLogin' | jq
*/
// httpCreateUser ...
func httpCreateUser(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpCreateUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := CreateUser(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpReadUser ...
func httpReadUser(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpReadUser")
	userID := r.URL.Query().Get("userID")
	ethAddress := r.URL.Query().Get("ethAddress")
	reqBody := ReqBody{EthereumAddr: ethAddress, UserID: userID}
	log1("over to lambda function")
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := ReadUser(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardC ...
func httpRewardC(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpRewardC")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardC(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardR ...
func httpRewardR(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpRewardR")
	userID := r.URL.Query().Get("userID")
	vaultID := r.URL.Query().Get("vaultID")
	rewardID := r.URL.Query().Get("rewardID")

	reqBody := ReqBody{UserID: userID, VaultID: vaultID, RewardID: rewardID}
	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardR(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardD ...
func httpRewardD(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpRewardD")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardD(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpVaultEthAddrR ...
func httpVaultEthAddrR(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpVaultEthAddrR")
	vaultID := r.URL.Query().Get("vaultID")
	reqBody := ReqBody{VaultID: vaultID}
	log1("over to lambda function")

	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := VaultEthAddrR(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/httpUser' | jq
*/
// httpDeleteUser ...
func httpDeleteUser(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpDeleteUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := DeleteUser(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/member' | jq
*/
// httpUpdateUser ...
func httpUpdateUser(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpUpdateUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)

	log1("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UpdateUser(inputLambda)
	log1("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logE.Println("\nerr@ lambda function output")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpCreateUsers ...
// func httpCreateUsers(w http.ResponseWriter, r *http.Request) {
// 	log1("---------------== httpCreateUsers")
// 	var reqBody ReqBody
// 	err := json.NewDecoder(r.Body).Decode(&reqBody)
// 	if err != nil {
// 		log1("json decode err:", err)
// 		return
// 	}
// 	log1("reqBody:", reqBody)

// 	log1("over to lambda function")
// 	inputLambda := InputLambda{Body: reqBody}
// 	outputLambdaPt, err := UsersC(inputLambda)
// 	log1("result:", outputLambdaPt)
// 	if err != nil || outputLambdaPt.Mesg != "ok" {
// 		log1("\n====>>>> err from this lambda function")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(*outputLambdaPt)
// 	if err != nil {
// 		logE.Println("Error @ json.NewEncoder:", err)
// 	}
// }

// // httpReadUsers ...
// func httpReadUsers(w http.ResponseWriter, r *http.Request) {
// 	log1("---------------== httpReadUsers")
// 	offset := r.URL.Query().Get("offset")
// 	amount := r.URL.Query().Get("amount")

// 	reqBody := ReqBody{Offset: offset, Amount: amount}
// 	log1("over to lambda function")
// 	inputLambda := InputLambda{Body: reqBody}
// 	outputLambdaPt, err := UsersR(inputLambda)
// 	log1("result:", outputLambdaPt)
// 	if err != nil || outputLambdaPt.Mesg != "ok" {
// 		log1("\n====>>>> err from this lambda function")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(*outputLambdaPt)
// 	if err != nil {
// 		logE.Println("Error @ json.NewEncoder:", err)
// 	}
// }

/*weth, yearnfinance, curvefi3pool, curvefiy, curvefibusd, curvefisbtc, daistablecoin, trueusd, usdc, geminidollar, tetherusd

doReadRow("SELECT * FROM account WHERE id = ?", 16, "accountID", accountID)

curl -v 'localhost:3000/httpApyR?sourceURL=https://stats.finance/yearn&perfPeriod=week' |jq
*/

// httpApyR ...
func httpApyR(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpApyR")
	// query := r.URL.Query()
	// sourceURL := query.Get("sourceURL")
	// perfPeriod := query.Get("perfPeriod")
	// log1("sourceURL:", sourceURL, ", perfPeriod:", perfPeriod)
	sourceURL := "https://stats.finance/yearn"
	perfPeriod := "week"
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		dataName = "yearnFinance"
	default:
		log1("sourceURL invalid:")
		outputLambdaPt := &OutputLambda{
			Code: "000320",
			Mesg: "sourceURL invalid",
		}
		err := json.NewEncoder(w).Encode(*outputLambdaPt)
		if err != nil {
			logE.Println("Error @ json.NewEncoder:", err)
		}
		return
	}

	reqBody := ReqBody{PerfPeriod: perfPeriod}
	inputLambda := InputLambda{Body: reqBody, DataName: dataName}
	//authHeader, updatedAt

	log1("------------==")
	vaultApySlicePtr, err := readLambda(inputLambda)
	log1("readLambda result:", vaultApySlicePtr)
	if err != nil {
		logE.Println("\n===> err@ readLambda, err:", err)
		outputLambdaPt := &OutputLambda{
			Code: "000321",
			Mesg: "err@ readLambda",
			Data: err,
		}
		err := json.NewEncoder(w).Encode(*outputLambdaPt)
		if err != nil {
			logE.Println("Error @ json.NewEncoder:", err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(*vaultApySlicePtr)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpApyC ...
func httpApyC(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpApyC")
	// query := r.URL.Query()
	// rowName := query.Get("rowname")
	// log1("rowName: " + rowName)

	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	log1("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	//authHeader, updatedAt

	log1("------------== write to db")
	outputLambdaPt, err := addRowDB(inputLambda)
	log1("addRowDB outputLambdaPt:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		log1("\n====>>>> err@ writeRowX")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logE.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRowD ...
func httpRowD(w http.ResponseWriter, r *http.Request) {
	log1("---------------== httpRowD")
}

func getAPIx(w http.ResponseWriter, r *http.Request) {
	log1("---------------== getAPIx")
	// offset := r.URL.Query().Get("offset")
	// amount := r.URL.Query().Get("amount")

	// reqBody := ReqBody{Offset: offset, Amount: amount}
	// log1("over to lambda function")
	// inputLambda := InputLambda{Body: reqBody}
	// outputLambdaPt, err := UsersR(inputLambda)
	// log1("result:", outputLambdaPt)
	// if err != nil || outputLambdaPt.Mesg != "ok" {
	// 	log1("\n====>>>> err from this lambda function")
	//}

	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(*anChainOutPt)
	// if err != nil {
	// 	logE.Println("Error @ json.NewEncoder:", err)
	// }
}

func ping(w http.ResponseWriter, r *http.Request) {
	log1("Ping")
	testDB()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("ping"))
}

func root(w http.ResponseWriter, r *http.Request) {
	log1("root")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("root"))
}

/*
 */
