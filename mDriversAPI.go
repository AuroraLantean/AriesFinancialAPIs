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
		routineAddr, "", restType, timeoutSecs}
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
// DoCallGraphQLAPI ...
/*@brief to call DoCallGraphQLAPI
@param out: result and error message
@param  in:
*/
func DoCallGraphQLAPI(dataName string, subgraphName string) (*GraphqlOut, string) {
	print("-----------== DoCallGraphQLAPI()")

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		return nil, "err@ loadEnv"
	}

	var queryStr, targetURL, uniswapV2TokenID string
	switch {
	case dataName == "AFI":
		uniswapV2TokenID = "0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891"
	default:
		print("dataName is invalid")
		return nil, "invalid dataName"
	}
	print("uniswapV2TokenID:", uniswapV2TokenID)
	// ` + uniswapV2TokenID + `

	switch {
	case subgraphName == "uniswapV2":
		targetURL = os.Getenv("UNISWAPV2")

		//queryStr = "xyz"
		queryStr = `{	pair(id: "0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891") {
			token0 {
				id
				name
				symbol
				decimals
				totalSupply
				tradeVolume
				tradeVolumeUSD
				untrackedVolumeUSD
				derivedETH
				tradeVolumeUSD
				totalLiquidity
			}
			token1 {
				id
				name
				symbol
				decimals
				totalSupply
				tradeVolume
				tradeVolumeUSD
				derivedETH
				tradeVolumeUSD
				untrackedVolumeUSD
				totalLiquidity
			}
			token0Price
			token1Price
		}
		}`

	default:
		print("subgraphName is invalid")
		return nil, "invalid subgraphName"
	}

	print("targetURL:", targetURL, "\nqueryStr:", queryStr)
	if targetURL == "" {
		return nil, "targetURL is empty"
	}
	if queryStr == "" {
		return nil, "queryStr is empty"
	}
	graphqlOutPt, msg, err := CallGraphQLAPI(queryStr, targetURL)
	if msg != "ok" || err != nil {
		print("msg:", msg, ". err:", err)
		return nil, "msg or err exists"
	}
	dump(graphqlOutPt)
	graphqlOut := *graphqlOutPt
	msgOut := ""
	if len(graphqlOut.Errors) > 0 && graphqlOut.Errors[0].Message != "" {
		msgOut = "Graphql query failed: " + graphqlOut.Errors[0].Message
		print(msgOut)
		return nil, msgOut
	}
	return graphqlOutPt, "ok"
}

// CallGraphQLAPI ...
/*@brief to call AnChain APIs
@param out: result and error message
@param  in:
*/
func CallGraphQLAPI(queryStr string, targetURL string) (*GraphqlOut, string, error) {
	print("-----------== CallGraphQLAPI()")
	GraphqlOut := GraphqlOut{}
	routineName := "MakeGraphqlRequest" //"MakeGetRequest"
	routineAddr := targetURL
	timeoutSecs := 3

	//endpoint1 := url.QueryEscape(endpoint)

	routineInputs := RoutineInputs{routineName,
		routineAddr, queryStr, "", timeoutSecs}
	routineOutPtr, err := ExecuteRoutine(routineInputs)
	//print("ExecuteRoutine result:", routineOutPtr)
	print("err:", err)

	byteSlice, ok := ((*routineOutPtr).RespRoutine).([]byte)
	if !ok {
		print("err@ RespRoutine not of []byte")
		return &GraphqlOut, "RespRoutine not of []byte", nil
	}

	err = json.Unmarshal(byteSlice, &GraphqlOut)
	if err != nil {
		print("err@ json.Unmarshal()", err)
	}
	return &GraphqlOut, "ok", nil
}
