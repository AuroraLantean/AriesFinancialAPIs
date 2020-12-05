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
	logI.Println("---------------== httpApyU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := apysUpdateColly(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err@ writeRowX")
	}

	logI.Println("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}

	// parse our response slice into JSON format
	//b, err := json.Marshal(ss)
	// if err != nil {
	// 	logI.Println("err:", err)
	// }
	//w.Write(b)
}

// httpAriesU ...
func httpAriesU(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpAriesU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := ariesU(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err@ writeRowX")
	}

	logI.Println("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpGetApyAries ...
func httpGetApyAries(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpGetApyAries")
	RewardsPool := r.URL.Query().Get("rewardspool")
	reqBody := ReqBody{RewardsPool: RewardsPool}
	logI.Println("over to lambda function")
	logI.Println("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := getApyAries(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err@ writeRowX")
	}

	logI.Println("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

/*
curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpApyReset'
| jq
*/
// httpApyReset ...
func httpApyReset(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpApyReset")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)
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
	logI.Println("------------== write to db")
	inputLambda := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := apysUpdate(inputLambda)
	logI.Println("result:", outputLambdaPt)
	outputLambdaPt.Data = apys
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err@ writeRowX")
	}

	logI.Println("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
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
	logI.Println("---------------== httpCreateUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := CreateUser(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpReadUser ...
func httpReadUser(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpReadUser")
	userID := r.URL.Query().Get("userID")
	ethAddress := r.URL.Query().Get("ethAddress")
	reqBody := ReqBody{EthereumAddr: ethAddress, UserID: userID}
	logI.Println("over to lambda function")
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := ReadUser(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardC ...
func httpRewardC(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpRewardC")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardC(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardR ...
func httpRewardR(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpRewardR")
	userID := r.URL.Query().Get("userID")
	vaultID := r.URL.Query().Get("vaultID")
	rewardID := r.URL.Query().Get("rewardID")

	reqBody := ReqBody{UserID: userID, VaultID: vaultID, RewardID: rewardID}
	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardR(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRewardD ...
func httpRewardD(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpRewardD")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardD(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpVaultEthAddrR ...
func httpVaultEthAddrR(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpVaultEthAddrR")
	vaultID := r.URL.Query().Get("vaultID")
	reqBody := ReqBody{VaultID: vaultID}
	logI.Println("over to lambda function")

	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := VaultEthAddrR(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
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
	logI.Println("---------------== httpDeleteUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := DeleteUser(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
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
	logI.Println("---------------== httpUpdateUser")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)

	logI.Println("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UpdateUser(inputLambda)
	logI.Println("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpCreateUsers ...
// func httpCreateUsers(w http.ResponseWriter, r *http.Request) {
// 	logI.Println("---------------== httpCreateUsers")
// 	var reqBody ReqBody
// 	err := json.NewDecoder(r.Body).Decode(&reqBody)
// 	if err != nil {
// 		logI.Println("json decode err:", err)
// 		return
// 	}
// 	logI.Println("reqBody:", reqBody)

// 	logI.Println("over to lambda function")
// 	inputLambda := InputLambda{Body: reqBody}
// 	outputLambdaPt, err := UsersC(inputLambda)
// 	logI.Println("result:", outputLambdaPt)
// 	if err != nil || outputLambdaPt.Mesg != "ok" {
// 		logI.Println("\n====>>>> err from this lambda function")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(*outputLambdaPt)
// 	if err != nil {
// 		logI.Println("Error @ json.NewEncoder:", err)
// 	}
// }

// // httpReadUsers ...
// func httpReadUsers(w http.ResponseWriter, r *http.Request) {
// 	logI.Println("---------------== httpReadUsers")
// 	offset := r.URL.Query().Get("offset")
// 	amount := r.URL.Query().Get("amount")

// 	reqBody := ReqBody{Offset: offset, Amount: amount}
// 	logI.Println("over to lambda function")
// 	inputLambda := InputLambda{Body: reqBody}
// 	outputLambdaPt, err := UsersR(inputLambda)
// 	logI.Println("result:", outputLambdaPt)
// 	if err != nil || outputLambdaPt.Mesg != "ok" {
// 		logI.Println("\n====>>>> err from this lambda function")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(*outputLambdaPt)
// 	if err != nil {
// 		logI.Println("Error @ json.NewEncoder:", err)
// 	}
// }

/*weth, yearnfinance, curvefi3pool, curvefiy, curvefibusd, curvefisbtc, daistablecoin, trueusd, usdc, geminidollar, tetherusd

doReadRow("SELECT * FROM account WHERE id = ?", 16, "accountID", accountID)

curl -v 'localhost:3000/httpApyR?sourceURL=https://stats.finance/yearn&perfPeriod=week' |jq
*/

// httpApyR ...
func httpApyR(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpApyR")
	// query := r.URL.Query()
	// sourceURL := query.Get("sourceURL")
	// perfPeriod := query.Get("perfPeriod")
	// logI.Println("sourceURL:", sourceURL, ", perfPeriod:", perfPeriod)
	sourceURL := "https://stats.finance/yearn"
	perfPeriod := "week"
	w.Header().Set("Content-Type", "application/json")

	var dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		dataName = "yearnFinance"
	default:
		logI.Println("sourceURL invalid:")
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

	logI.Println("------------==")
	vaultApySlicePtr, err := readLambda(inputLambda)
	logI.Println("readLambda result:", vaultApySlicePtr)
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
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpApyC ...
func httpApyC(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpApyC")
	// query := r.URL.Query()
	// rowName := query.Get("rowname")
	// logI.Println("rowName: " + rowName)

	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		logE.Println("json decode err:", err)
		return
	}
	logI.Println("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	//authHeader, updatedAt

	logI.Println("------------== write to db")
	outputLambdaPt, err := addRowDB(inputLambda)
	logI.Println("addRowDB outputLambdaPt:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		logI.Println("\n====>>>> err@ writeRowX")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		logI.Println("Error @ json.NewEncoder:", err)
	}
}

// httpRowD ...
func httpRowD(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== httpRowD")
}

func getAPIx(w http.ResponseWriter, r *http.Request) {
	logI.Println("---------------== getAPIx")
	// offset := r.URL.Query().Get("offset")
	// amount := r.URL.Query().Get("amount")

	// reqBody := ReqBody{Offset: offset, Amount: amount}
	// logI.Println("over to lambda function")
	// inputLambda := InputLambda{Body: reqBody}
	// outputLambdaPt, err := UsersR(inputLambda)
	// logI.Println("result:", outputLambdaPt)
	// if err != nil || outputLambdaPt.Mesg != "ok" {
	// 	logI.Println("\n====>>>> err from this lambda function")
	//}

	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(*anChainOutPt)
	// if err != nil {
	// 	logI.Println("Error @ json.NewEncoder:", err)
	// }
}

func ping(w http.ResponseWriter, r *http.Request) {
	logI.Println("Ping")
	testDB()
	w.Write([]byte("ping"))
}

func root(w http.ResponseWriter, r *http.Request) {
	logI.Println("root")
	w.Write([]byte("root"))
}

/*
 */
