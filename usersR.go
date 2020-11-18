package main

/*@file UsersR.go
@brief UsersR API
see UsersR function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

// UsersR ...
func UsersR(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== UsersR")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	offset := reqBody.Offset
	amount := reqBody.Amount
	ok1, err1 := checkInput(offset, 1, "Offset")
	ok2, err2 := checkInput(amount, 1, "Amount")

	if !(ok1 && ok2) {
		print("inputs invalid")
		print("ok1:", ok1, ", err1:", err1)
		print("ok2:", ok2, ", err2:", err2)
		return &OutputLambda{
			Code: "110000",
			Mesg: "Only One of UserID and Ethereum Address should be valid",
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

	//-----------------===
	stmtIn1 := "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs FROM AriesFinancial.User WHERE limit ?, ?"
	//SELECT * FROM table limit 5, 10 ... row 6~15
	returnSize := 6

	_, isOk, err := readRow(db, stmtIn1, returnSize, offset, amount)
	// check error conditions: row not found, error existing, or delete value is valid...
	switch {
	case err == sql.ErrNoRows:
		print("row not found")

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
		print("Ethereum address already exists for idx")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: nil,
	}, nil
}

/*
UsersR
curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/UsersR' | jq

*/
