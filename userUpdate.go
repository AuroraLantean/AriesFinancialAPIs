package main

/*@file UpdateUser.go
@brief UpdateUser API
see UpdateUser function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

/*UpdateUser ...
User Update Scenario - UpdateUser: 
an existing user wants to update his info in our User DB table.
the API will check for this user's address via security function ->
if this address is okay, then update such user's info in our DB
else if this address is a hacker or something like that, response is to reject this request.
*/
func UpdateUser(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== UpdateUser")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	userID := reqBody.UserID
	ethAddress := reqBody.EthereumAddr
	reward := reqBody.Reward
	downlineIDs := reqBody.DownlineIDs
	uplineID := reqBody.UplineID

	ok1, err1 := checkInput(userID, 1, "User ID")
	ok2, err2 := checkInput(ethAddress, 42, "Ethereum Address")
	if !(ok1 != ok2) {
		print("input values are not valid")
		print("ok1:", ok1, ", err1:", err1)
		print("ok2:", ok2, ", err2:", err2)
		return &OutputLambda{
			Code: "110000",
			Mesg: "Only One of UserID and Ethereum Address should be valid",
			Data: nil,
		}, nil
		/*idx, err := getErr([]error{err1, err2})
		print("err@getErr:", err)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: index " + toStr(idx) + " of [Ethereum Address or UserID] is empty or invalid",
			Data: nil,
		}, nil
		*/
	}
	if ok1 && toInt(userID) <= 0 {
		return &OutputLambda{
			Code: "110000",
			Mesg: "userID invalid",
			Data: nil,
		}, nil
	}

	var ok3 bool
	var err3 error
	if toFloat(reward) >= 0 {
		ok3 = true
	}
	//ok3, err3 := checkInput(reward, 1, "reward")
	ok4, err4 := checkInput(uplineID, 1, "upID")
	ok5, err5 := checkInput(downlineIDs, 1, "downlineIDs")
	if !(onlyOneIsTrue([]bool{ok3, ok4, ok5})) {
		print("only one input should be valid")
		print("ok3:", ok3, ", err3:", err3)
		print("ok4:", ok4, ", err4:", err4)
		print("ok5:", ok5, ", err4:", err5)
		return &OutputLambda{
			Code: "110000",
			Mesg: "One of reward, uplineID and downlineIDs should be valid, the other two should be invalid",
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
	print("-----------== ")
	defer func() {
		errDbClose := db.Close()
		if errDbClose != nil {
			print("err@ db.Close:", errDbClose)
			return
		}
	}()

	//-----------------=== check bcAddr risk
	bcAddr := ethAddress
	msgOut := DoCallAnChainAPI("eth", bcAddr)
	if msgOut != "ok" {
		return &OutputLambda{
			Code: "110009",
			Mesg: msgOut,
		}, nil
	}

	_, timeNowStr, err := GetLocalTime(LocStr)
	if err != nil {
		print("err@ GetLocalTime", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ GetLocalTime",
		}, nil
	}

	//-----------------==
	var affectedRowID string
	if ok1 {

		if ok3 {
			//"UPDATE account SET pass = ?, zzz = ? WHERE id = ?"
			stmtIn := "UPDATE AriesFinancial.User SET reward = ?, riskCheckedAt = ? WHERE id = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.Reward, timeNowStr, userID)
		} else if ok4 {
			stmtIn := "UPDATE AriesFinancial.User SET uplineID = ?, riskCheckedAt = ? WHERE id = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.UplineID, timeNowStr, userID)
		} else if ok5 {
			stmtIn := "UPDATE AriesFinancial.User SET downlineIDs = ?, riskCheckedAt = ? WHERE id = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.DownlineIDs, timeNowStr, userID)
		} else {
			print("reward, uplineID, and downlineIDs are not valid")
		}

	} else if ok2 {
		if ok3 {
			stmtIn := "UPDATE AriesFinancial.User SET reward = ?, riskCheckedAt = ? WHERE ethAddress = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.Reward, timeNowStr, ethAddress)
		} else if ok4 {
			stmtIn := "UPDATE AriesFinancial.User SET uplineID = ?, riskCheckedAt = ? WHERE ethAddress = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.UplineID, timeNowStr, ethAddress)
		} else if ok5 {
			stmtIn := "UPDATE AriesFinancial.User SET downlineIDs = ?, riskCheckedAt = ? WHERE ethAddress = ?"
			affectedRowID, err = writeRowV(db, stmtIn, reqBody.DownlineIDs, timeNowStr, ethAddress)
		} else {
			print("reward, uplineID, and downlineIDs are not valid")
		}
	} else {
		print("both userID and ethAddress are not valid")
	}
	print("check2 affectedRowID:", affectedRowID, ", err:", err)

	// check error conditions: row not found, error existing, or delete value is valid...
	switch {
	case err == sql.ErrNoRows:
		print("row not found")
		return &OutputLambda{
			Code: "110001",
			Mesg: "row not found",
			Data: nil,
		}, nil

	case err != nil:
		print("err@ writeRow:", err)
		return &OutputLambda{
			Code: "110001",
			Mesg: "err@ writeRow",
			Data: nil,
		}, nil

	default:
		print("------== user row has been updated")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: nil,
	}, nil
}

/*
UpdateUser
curl -XPOST -d ' {"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f","reward":"2.9","downlineIDs":""}' 'localhost:3000/membership' | jq

curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f","reward":"","downlineIDs":"3,4"}' 'localhost:3000/membership' | jq

curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f","reward":"","uplineID":"7","downlineIDs":""}' 'localhost:3000/membership' | jq

*/
