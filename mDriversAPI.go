package main

/*@file mDriversAPI.go
@brief various driver functions
see function descriptions below

@author
@date
*/
import (
	//	"errors"
	//	"time"

	"encoding/json"
	"net/url"
	"os"
)

// DoCallAnChainAPI ...
/*@brief to call DoCallAnChainAPI
@param out: result and error message
@param  in:
*/
func DoCallAnChainAPI(proto string, bcAddr string) string {
	print("-----------== DoCallAnChainAPI()")
	apiChoice := 2
	anChainOutPt, msg, err := CallAnChainAPI(apiChoice, proto, bcAddr)
	if msg != "ok" || err != nil {
		print("msg:", msg, ". err:", err)
		return "msg or err exists"
	}
	dump(anChainOutPt)
	anChainOut := *anChainOutPt
	msgOut := ""
	if anChainOut.ErrMsg != "" || anChainOut.Status != 200 {
		msgOut = "EthAddress risk checking failed: " + anChainOut.ErrMsg
		print(msgOut)
		return msgOut
	}
	riskScore := anChainOut.Data[bcAddr].Risk.Score
	print("riskScore:", riskScore)
	if riskScore > MaxRiskScore {
		msgOut = "EthAddress risk score is too high"
		print(msgOut)
		return msgOut
	}
	category := anChainOut.Data[bcAddr].Self.Category
	print("category:", category, "\nBlacCategories", BlacCategories)
	if strSliceHasAny(category, BlacCategories) {
		msgOut = "EthAddress risk category is on BlacCategories"
		print(msgOut)
		return msgOut
	}
	return "ok"
}

// CallAnChainAPI ...
/*@brief to call AnChain APIs
@param out: result and error message
@param  in:
*/
func CallAnChainAPI(choice int, proto string, bcAddr string) (*AnChainOut, string, error) {
	print("-----------== CallAnChainAPI()")
	anChainOut := AnChainOut{}
	routineName := "MakeHTTPRequest" //"MakeGetRequest"
	timeoutSecs := 3

	var endpoint, restType string
	switch {
	case choice == 1:
		endpoint = "address_info"
		restType = "GET"
	case choice == 2:
		endpoint = "address_risk_score"
		restType = "GET"
	case choice == 3:
		endpoint = "address_risk_info"
		restType = "GET"
	default:
		print("choice has no match")
	}

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		return &anChainOut, isEnvOk, nil
	}

	anChainKey := os.Getenv("ANCHAINKEY")
	if anChainKey == "" {
		return &anChainOut, "AnChainKey empty", nil
	}

	endpoint1 := url.QueryEscape(endpoint)
	proto1 := url.QueryEscape(proto)
	bcAddr1 := url.QueryEscape(bcAddr)
	routineAddr := "https://bei.anchainai.com/api/" + endpoint1 + "?proto=" + proto1 + "&address=" + bcAddr1 + "&apikey=" + anChainKey
	print("apiStr:", routineAddr)

	routineInputs := RoutineInputs{routineName,
		routineAddr, restType, timeoutSecs}
	routineOutPtr, err := ExecuteRoutine(routineInputs)
	//print("ExecuteRoutine result:", routineOutPtr)
	print("err:", err)

	byteSlice, ok := ((*routineOutPtr).RespRoutine).([]byte)
	if !ok {
		print("err@ RespRoutine not of []byte")
		return &anChainOut, "RespRoutine not of []byte", nil
	}

	err = json.Unmarshal(byteSlice, &anChainOut)
	if err != nil {
		print("err@ json.Unmarshal()", err)
	}
	return &anChainOut, "ok", nil

	// req, err := http.NewRequest("GET", apiStr, nil)
	// if err != nil {
	// 		log.Print(err)
	// 		return "", err
	// }

	// q := req.URL.Query()
	// //q.Add("api_key", "key_from_environment_or_flag")
	// //q.Add("another_thing", "foo & bar")
	// req.URL.RawQuery = q.Encode()

	// print(req.URL.String())

}

/*

 */
