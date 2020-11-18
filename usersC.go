package main

/*@file UsersC.go
@brief UsersC API
see UsersC function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

// UsersC ...
func UsersC(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== UsersC")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	ethAddrs := reqBody.EthereumAddrs
	ok1 := true
	var err1 error
	var idx1 int
	for idx, v := range ethAddrs {
		okIdx, errIdx := checkInput(v, 42, "EthAddr"+toStr(idx))
		if !okIdx {
			print("errIdx:", errIdx)
			ok1 = okIdx
			err1 = errIdx
			idx1 = idx
			break
		}
	}
	if !(ok1) {
		print("one input value is not valid")
		print("ok1:", ok1, ", err1:", err1)
		//idx, err := getErr([]error{err1})
		//print("err@getErr:", err)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: index " + toStr(idx1) + " of Ethereum Addresses is empty or invalid",
			Data: nil,
		}, nil
	}

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		return &OutputLambda{
			Code: "110017",
			Mesg: isEnvOk,
			Data: nil,
		}, nil
	}
	print("check1 loadEnv is ok")

	db, err := dbConn()
	if err != nil {
		print("err@dbConn():", err)
		return &OutputLambda{
			Code: "110009",
			Mesg: "cannot connect to database",
			Data: nil,
		}, nil
	}
	print("check2 dbConn() connection is ok")

	defer func() {
		errDbClose := db.Close()
		if errDbClose != nil {
			print("err@ db.Close:", errDbClose)
			return
		}
	}()

	//-----------------== Security
	/*
		routineOutPtr, err := SendSmsByEVERY8D(rawBody.CountryCode, phoneNum1, msg)
		Log1("SendSmsByEVERY8D() result:", routineOutPtr, err)
		outCode := (*routineOutPtr).Code
		if err != nil || outCode != "0" {
			Log1("err@ Sending SMS failed:", err)

			if avatar == "311" {
				return &Output{
					Status: "null",
					Code:   outCode,
					Mesg:   (*routineOutPtr).Mesg,
					Data:   RespFa1SendSMS2{
						VerifyCheckID: verifyCheckID,
						DueTime:       dueTime,
						ExternalResp: (*routineOutPtr).RespRoutine,
						VerifyCode: verifCode,
					},
				}, nil
			}
			return &Output{
				Status: "null",
				Code:   outCode,
				Mesg:   (*routineOutPtr).Mesg,
				Data:   nil,
			}, nil
		} //return &Output{}, nil
		Log1("check8. Sending SMS successful")
	*/

	//-----------------== Add into DB
	_, timeNowStr, err := GetLocalTime(LocStr)
	if err != nil {
		print("err@ GetLocalTime", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ GetLocalTime",
			Data: nil,
		}, nil
	}

	stmtIn1 := "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs FROM AriesFinancial.User WHERE ethAddress = ?"
	returnSize := 6

	print("about to enter into ethAddrs loop")
	for idx, ethAddress := range ethAddrs {
		print("idx", idx, ethAddress)
		//-----------------=== check bcAddr risk
		bcAddr := ethAddress
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			print("[WARNING] address found with high risk:", bcAddr)
			continue
		}

		//-----------------=== write into db
		_, isOk, err := readRow(db, stmtIn1, returnSize, ethAddress)
		idxStr := toStr(idx)
		// check error conditions: row not found, error existing, or delete value is valid...
		switch {
		case err == sql.ErrNoRows:
			print("row not found")

			stmtIn2 := "INSERT INTO AriesFinancial.User (ethAddress, updatedAt, riskCheckedAt) VALUES (?,?,?)"
			userID, err := writeRowV(db, stmtIn2, ethAddress, timeNowStr, timeNowStr)
			if err != nil {
				print("err@ writeRow() inserting new row", err)
				return &OutputLambda{
					Code: "000104",
					Mesg: "err@ writeRowV on idx " + idxStr,
					Data: nil,
				}, nil
			}
			print("new userID:", userID)
			print("db writing is successful")

		case isOk != "ok":
			print("err@ isOk:", err)
			return &OutputLambda{
				Code: "110001",
				Mesg: "err@ isOk: " + isOk,
				Data: nil,
			}, nil

		case err != nil:
			print("err@ readRow:", err)
			return &OutputLambda{
				Code: "110001",
				Mesg: "err@ readRow",
				Data: nil,
			}, nil

		default:
			print("Ethereum address already exists for idx", idxStr)
		}
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: nil,
	}, nil
}

/*
UsersC
curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/UsersC' | jq

*/
