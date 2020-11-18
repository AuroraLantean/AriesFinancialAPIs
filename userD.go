package main

/*@file UserD.go
@brief UserD API
see UserD function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

// UserD ...
func UserD(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== UserD")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	ethAddress := reqBody.EthereumAddr
	userID := reqBody.UserID
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

	//-----------------=== check bcAddr risk
	if ok1 {
		stmtIn := "SELECT ethAddress, riskCheckedAt FROM AriesFinancial.User WHERE id = ?"
		returnSize := 2
		result, isOk, err := readRow(db, stmtIn, returnSize, userID)
		ethAddress = result[0]
		//riskCheckedAt := result[1]
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
	}

	bcAddr := ethAddress
	msgOut := DoCallAnChainAPI("eth", bcAddr)
	if msgOut != "ok" {
		return &OutputLambda{
			Code: "110009",
			Mesg: msgOut,
		}, nil
	}

	// stmtIn := "UPDATE AriesFinancial.User SET riskCheckedAt = ? WHERE id = ?"
	// affectedRowID, err := writeRowV(db, stmtIn, timeNowStr, userID)
	// print("affectedRowID:", affectedRowID)
	// if err != nil {
	// 	print("err@ writeRow:", err)
	// 	return &OutputLambda{
	// 		Code: "110001",
	// 		Mesg: "err@ writeRow riskCheckedAt",
	// 		Data: nil,
	// 	}, nil
	// }
	// print("------== user riskCheckedAt has been updated")

	//-----------------==
	var affectedRowID string
	if ok1 {
		stmtIn := "DELETE FROM AriesFinancial.User WHERE id = ?"
		affectedRowID, err = writeRowV(db, stmtIn, userID)

	} else if ok2 {
		stmtIn := "DELETE FROM AriesFinancial.User WHERE ethAddress = ?"
		affectedRowID, err = writeRowV(db, stmtIn, ethAddress)
	} else {
		print("both userID and Ethereum Address are not valid")
	}
	print("check2 deleted affectedRowID:", affectedRowID, ", err:", err)

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
		print("------== user row has been deleted")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: nil,
	}, nil
}

/*
UserD
curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/userdelete' | jq

*/
