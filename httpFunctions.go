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
	print("---------------== httpApyU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := apysScrapeUpdate(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err@ writeRowX")
	}

	print("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}

	// parse our response slice into JSON format
	//b, err := json.Marshal(ss)
	// if err != nil {
	// 	print("err:", err)
	// }
	//w.Write(b)
}

/*
curl -XPUT -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpApyReset'
| jq
*/
// httpApyReset ...
func httpApyReset(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpApyReset")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)
	sourceURL := reqBody.SourceURL

	var dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		dataName = "yearnFinance"
	default:
		print("sourceURL invalid:")
		return
	}

	apys := APYs{WETH: "0.01", AFI: "1.11", YFI: "2.22", CRV3: "3.33", CRVY: "4.44", CRVBUSD: "5.55", CRVSBTC: "6.66", DAI: "7.77", TrueUSD: "8.88", USDC: "9.99", Gemini: "10.00", TetherUSD: "11.11"}
	print("------------== write to db")
	inputLambda := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := apysUpdate(inputLambda)
	print("result:", outputLambdaPt)
	outputLambdaPt.Data = apys
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err@ writeRowX")
	}

	print("IsToScrape:", IsToScrape)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}

}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/httpUserLogin' | jq
*/
// httpUserC ...
func httpUserC(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUserC")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UserC(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpUserR ...
func httpUserR(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUserR")
	userID := r.URL.Query().Get("userID")
	ethAddress := r.URL.Query().Get("ethAddress")
	reqBody := ReqBody{EthereumAddr: ethAddress, UserID: userID}
	print("over to lambda function")
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UserR(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpRewardC ...
func httpRewardC(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpRewardC")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardC(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpRewardR ...
func httpRewardR(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpRewardR")
	userID := r.URL.Query().Get("userID")
	vaultID := r.URL.Query().Get("vaultID")
	rewardID := r.URL.Query().Get("rewardID")

	reqBody := ReqBody{UserID: userID, VaultID: vaultID, RewardID: rewardID}
	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardR(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpRewardD ...
func httpRewardD(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpRewardD")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := RewardD(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpVaultEthAddrR ...
func httpVaultEthAddrR(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpVaultEthAddrR")
	vaultID := r.URL.Query().Get("vaultID")
	reqBody := ReqBody{VaultID: vaultID}
	print("over to lambda function")

	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := VaultEthAddrR(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/httpUserD' | jq
*/
// httpUserD ...
func httpUserD(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUserD")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UserD(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

/*
0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f
0xCb6d5D1a7BE3cDf0A24CD945ff97e98FDa5D87C1
0xB197Fe6a0031b476B7b045a628A9Ce2421fa1D2E

curl -XPOST -d '{"EthereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/httpuserupdate' | jq
*/
// httpUserU ...
func httpUserU(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUserU")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UserU(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpUsersC ...
func httpUsersC(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUsersC")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)

	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UsersC(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpUsersR ...
func httpUsersR(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpUsersR")
	offset := r.URL.Query().Get("offset")
	amount := r.URL.Query().Get("amount")

	reqBody := ReqBody{Offset: offset, Amount: amount}
	print("over to lambda function")
	inputLambda := InputLambda{Body: reqBody}
	outputLambdaPt, err := UsersR(inputLambda)
	print("result:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err from this lambda function")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

/*weth, yearnfinance, curvefi3pool, curvefiy, curvefibusd, curvefisbtc, daistablecoin, trueusd, usdc, geminidollar, tetherusd

doReadRow("SELECT * FROM account WHERE id = ?", 16, "accountID", accountID)

curl -v 'localhost:3000/httpApyR?sourceURL=https://stats.finance/yearn&perfPeriod=week' |jq
*/

// httpApyR ...
func httpApyR(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpApyR")
	// query := r.URL.Query()
	// sourceURL := query.Get("sourceURL")
	// perfPeriod := query.Get("perfPeriod")
	// print("sourceURL:", sourceURL, ", perfPeriod:", perfPeriod)
	sourceURL := "https://stats.finance/yearn"
	perfPeriod := "week"
	w.Header().Set("Content-Type", "application/json")

	var dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		dataName = "yearnFinance"
	default:
		print("sourceURL invalid:")
		outputLambdaPt := &OutputLambda{
			Code: "000320",
			Mesg: "sourceURL invalid",
		}
		err := json.NewEncoder(w).Encode(*outputLambdaPt)
		if err != nil {
			print("Error @ json.NewEncoder:", err)
		}
		return
	}

	reqBody := ReqBody{PerfPeriod: perfPeriod}
	inputLambda := InputLambda{Body: reqBody, DataName: dataName}
	//authHeader, updatedAt

	print("------------==")
	vaultApySlicePtr, err := readLambda(inputLambda)
	print("readLambda result:", vaultApySlicePtr)
	if err != nil {
		print("\n===> err@ readLambda, err:", err)
		outputLambdaPt := &OutputLambda{
			Code: "000321",
			Mesg: "err@ readLambda",
			Data: err,
		}
		err := json.NewEncoder(w).Encode(*outputLambdaPt)
		if err != nil {
			print("Error @ json.NewEncoder:", err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(*vaultApySlicePtr)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpApyC ...
func httpApyC(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpApyC")
	// query := r.URL.Query()
	// rowName := query.Get("rowname")
	// print("rowName: " + rowName)

	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)
	inputLambda := InputLambda{Body: reqBody}
	//authHeader, updatedAt

	print("------------== write to db")
	outputLambdaPt, err := addRowDB(inputLambda)
	print("addRowDB outputLambdaPt:", outputLambdaPt)
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err@ writeRowX")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpRowD ...
func httpRowD(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpRowD")
}

func getAPIx(w http.ResponseWriter, r *http.Request) {
	print("---------------== getAPIx")
	// offset := r.URL.Query().Get("offset")
	// amount := r.URL.Query().Get("amount")

	// reqBody := ReqBody{Offset: offset, Amount: amount}
	// print("over to lambda function")
	// inputLambda := InputLambda{Body: reqBody}
	// outputLambdaPt, err := UsersR(inputLambda)
	// print("result:", outputLambdaPt)
	// if err != nil || outputLambdaPt.Mesg != "ok" {
	// 	print("\n====>>>> err from this lambda function")
	//}

	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(*anChainOutPt)
	// if err != nil {
	// 	print("Error @ json.NewEncoder:", err)
	// }
}

func ping(w http.ResponseWriter, r *http.Request) {
	print("Ping")
	testDB()
	w.Write([]byte("ping"))
}

func root(w http.ResponseWriter, r *http.Request) {
	print("root")
	w.Write([]byte("root"))
}

/*
 */
