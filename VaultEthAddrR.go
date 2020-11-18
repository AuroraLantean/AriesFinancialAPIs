package main

/*@file VaultEthAddrR.go
@brief VaultEthAddrR API
see VaultEthAddrR function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

// VaultEthAddrR ...
func VaultEthAddrR(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== VaultEthAddrR")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	vaultID := reqBody.VaultID

	ok2, err2 := checkInput(vaultID, 1, "vaultID")
	if !(ok2) {
		print("input values are not valid")
		print("ok2:", ok2, ", err2:", err2)
		idx, err := getErr([]error{err2})
		print("err@getErr:", err)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: index " + toStr(idx) + " of [userID, vaultID] is empty or invalid",
			Data: nil,
		}, nil
	}
	if !(toInt(vaultID) > 0) {
		print("input value is not valid")
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid:",
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
	stmtIn1 := "SELECT vaultID, ethAddr FROM AriesFinancial.VaultEthAddr WHERE vaultID = ?"
	//SELECT * FROM table limit 5, 10 ... row 6~15
	returnSize := 2

	rowOut, isOk, err := readRow(db, stmtIn1, returnSize, vaultID)
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
		print("readRow is ok")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: RespVaultEthAddrR{
			VaultID:      rowOut[0],
			EthereumAddr: rowOut[1],
		},
	}, nil
}

/*
VaultEthAddrR
curl 'localhost:3000/vaultethaddr?userID=1&vaultID=1' | jq

*/
