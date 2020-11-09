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

curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/update' | jq

*/

// updateHTTP ...
func updateHTTP(w http.ResponseWriter, r *http.Request) {
	print("---------------== updateHTTP")
	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		print("json decode err:", err)
		return
	}
	print("reqBody:", reqBody)
	sourceURL := reqBody.SourceURL
	// htmlPattern := reqBody.HTPMPattern
	// regexpStr := reqBody.RegexpStr

	// sourceURL := "https://stats.finance/yearn"

	var htmlPattern, regexpStr, dataName string
	switch {
	case sourceURL == "https://stats.finance/yearn":
		htmlPattern = ".MuiTable-root tbody tr"
		regexpStr = `/\s([^}]*)\%`
		dataName = "yearnFinance"
	default:
		print("sourceURL invalid:")
		return
	}
	// "([a-z]+)"

	//Verify the param "URL" exists
	// URL := r.URL.Query().Get("url")
	// if URL == "" {
	// 	print("missing URL argument")
	// 	return
	// }
	print("visiting", sourceURL)

	var ss []string
	if IsToFetch {
		ss1, err := collyScraper(sourceURL, htmlPattern)
		if err != nil {
			print("failed to serialize response:", err)
			return
		}
		ss = ss1
	} else {
		ss2, err := collyScraperFakeYFI1()
		if err != nil {
			print("failed to serialize response:", err)
			return
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
			return
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
	inputLambda := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := updateLambda(inputLambda)
	print("updateLambda result:", outputLambdaPt)
	outputLambdaPt.Data = apys
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err@ writeRowX")
	}
	//-----------------==
	// outputLambdaPt = &OutputLambda{
	// 	Code: "000000",
	// 	Mesg: "ok",
	// 	Data: apys,
	// }
	print("IsToFetch:", IsToFetch)
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
curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpWriteRow'
| jq
*/
// httpWriteRow ...
func httpWriteRow(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpWriteRow")
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

	apys := APYs{WETH: "0.34", AFI: "1.13", YFI: "2.22", CRV3: "3.3", CRVY: "4.4", CRVBUSD: "5.5", CRVSBTC: "6.6", DAI: "7.7", TrueUSD: "8.8", USDC: "9.9", Gemini: "10.0", TetherUSD: "11.11"}
	print("------------== write to db")
	inputLambda := InputLambda{Body: reqBody, DataName: dataName, APYboWeek: apys}
	outputLambdaPt, err := updateLambda(inputLambda)
	print("updateLambda result:", outputLambdaPt)
	outputLambdaPt.Data = apys
	if err != nil || outputLambdaPt.Mesg != "ok" {
		print("\n====>>>> err@ writeRowX")
	}

	print("IsToFetch:", IsToFetch)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}

}

/*weth, yearnfinance, curvefi3pool, curvefiy, curvefibusd, curvefisbtc, daistablecoin, trueusd, usdc, geminidollar, tetherusd

doReadRow("SELECT * FROM account WHERE id = ?", 16, "accountID", accountID)

curl -v 'localhost:3000/readHTTP?sourceURL=https://stats.finance/yearn&perfPeriod=week' |jq
*/

// readHTTP ...
func readHTTP(w http.ResponseWriter, r *http.Request) {
	print("---------------== readHTTP")
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
		print("\n===> err@ readLambda, err:",err)
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
	//-----------------==
	// outputLambdaPt = &OutputLambda{
	// 	Code: "000000",
	// 	Mesg: "ok",
	// 	Data: apys,
	// }
	err = json.NewEncoder(w).Encode(*vaultApySlicePtr)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpAddRow ...
func httpAddRow(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpAddRow")
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
	//-----------------==
	// outputLambdaPt = &OutputLambda{
	// 	Code: "000000",
	// 	Mesg: "ok",
	// 	Data: apys,
	// }
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*outputLambdaPt)
	if err != nil {
		print("Error @ json.NewEncoder:", err)
	}
}

// httpDeleteRow ...
func httpDeleteRow(w http.ResponseWriter, r *http.Request) {
	print("---------------== httpDeleteRow")
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
	acc, err := readAccount(db, "SELECT * FROM account WHERE account_name = ?", reqBody.AccountName, "")
	print("err@readAccount():", err)

	// check error conditions: row not found, error existing, or delete value is valid...
	switch {
	case err == sql.ErrNoRows:
		print("row not found")
		return &OutputLambda{
			Code: "110006",
			Mesg: "account name not found for " + reqBody.AccountName,
			Data: nil,
		}, nil

	case err != nil:
		print("err:", err)
		return &OutputLambda{
			Code: "110001",
			Mesg: "cannot read from database. Table does not exist or column name mismatch",
			Data: nil,
		}, nil

	default:
		print("account is found")
	}
	print("acc from readAccount():", acc.ID, acc.AccountName, acc.CreatedAt, acc.UpdatedAt, acc.DeletedAt)
	print("account is found and valid")
	print("password is correct! accountID:", accountID)

*/
