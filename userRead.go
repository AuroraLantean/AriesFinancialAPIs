package main

/*@file ReadUser.go
@brief ReadUser API
see ReadUser function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

/*ReadUser ...
User Read Scenario - ReadUser: 
an existing user wants to read his info in our User DB table.
the API will check for this user's address via security function ->
if this address is okay, then fetch such user's info from our DB
else if this address is a hacker or something like that, response is to reject this request. 
*/
func ReadUser(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== ReadUser")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	userID := reqBody.UserID
	ethAddress := reqBody.EthereumAddr
	ok1, err1 := checkInput(userID, 1, "User ID")
	ok2, err2 := checkInput(ethAddress, 42, "Ethereum Address")

	if !(ok1 != ok2) {
		print("only one input should be valid")
		print("ok1:", ok1, ", err1:", err1)
		print("ok2:", ok2, ", err2:", err2)
		return &OutputLambda{
			Code: "110000",
			Mesg: "Only One of UserID and Ethereum Address should be valid",
			Data: nil,
		}, nil
	}
	if ok1 && toInt(userID) <= 0 {
		return &OutputLambda{
			Code: "110000",
			Mesg: "userID invalid",
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

	//-----------------==
	_, timeNowStr, err := GetLocalTime(LocStr)
	if err != nil {
		print("err@ GetLocalTime", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ GetLocalTime",
		}, nil
	}
	//-----------------==
	var userRow []string
	var isOk string
	var stmtIn string
	if ok1 {
		stmtIn = "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs, riskCheckedAt FROM AriesFinancial.User WHERE id = ?"
		returnSize := 7
		userRow, isOk, err = readRow(db, stmtIn, returnSize, userID)
		ethAddress = userRow[1]
		//riskCheckedAt := userRow[1]
		// check error conditions: row not found, error existing, or delete value is valid...
		switch {
		case err == sql.ErrNoRows:
			print("row not found")
			return &OutputLambda{
				Code: "110001",
				Mesg: "row not found for userID = " + userID,
				Data: nil,
			}, nil

		case isOk != "ok":
			print("err@ isOk:", err)
			return &OutputLambda{
				Code: "110001",
				Mesg: "err@ isOk read EthAddr: " + isOk,
				Data: nil,
			}, nil

		case err != nil:
			print("err@ readRow:", err)
			return &OutputLambda{
				Code: "110001",
				Mesg: "err@ read EthAddr from userID",
				Data: nil,
			}, nil

		default:
			print("Ethereum address is found:", ethAddress)
		}

		bcAddr := ethAddress
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			return &OutputLambda{
				Code: "110009",
				Mesg: msgOut,
			}, nil
		}

	} else if ok2 {
		//-----------------== check bcAddr risk
		bcAddr := ethAddress
		msgOut := DoCallAnChainAPI("eth", bcAddr)
		if msgOut != "ok" {
			return &OutputLambda{
				Code: "110009",
				Mesg: msgOut,
			}, nil
		}

		stmtIn = "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs, riskCheckedAt FROM AriesFinancial.User WHERE ethAddress = ?"
		returnSize := 7
		userRow, isOk, err = readRow(db, stmtIn, returnSize, ethAddress)

		// check error conditions: row not found, error existing, or delete value is valid...
		switch {
		case err == sql.ErrNoRows:
			print("row not found")
			return &OutputLambda{
				Code: "110009",
				Mesg: "userRow not found for given ethAddress",
			}, nil

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
			print("Ethereum address is found with userRow:", userRow)
		}
	} else {
		print("ok1 and ok2 are both false")
		return &OutputLambda{
			Code: "110009",
			Mesg: "userID and ethAddress are both empty",
		}, nil
	}

	//write riskCheckedAt to this user
	stmtIn = "UPDATE AriesFinancial.User SET riskCheckedAt = ? WHERE ethAddress = ?"
	affectedRowID, err := writeRowV(db, stmtIn, timeNowStr, ethAddress)
	switch {
	case err != nil:
		print("err@ writeRowV:", err)
		return &OutputLambda{
			Code: "110001",
			Mesg: "err@ writeRowV",
			Data: nil,
		}, nil

	default:
		print("writing riskCheckedAt to existing user successful. affectedRowID:", affectedRowID)
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: RespUser{
			ID:            userRow[0],
			EthereumAddr:  userRow[1],
			Reward:        userRow[2],
			UpdatedAt:     userRow[3],
			UplineID:      userRow[4],
			DownlineIDs:   userRow[5],
			RiskCheckedAt: timeNowStr,
		},
	}, nil
}

/*
UserR
curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/UserR' | jq

*/
