package main

import "database/sql"

/*@file RewardC.go
@brief RewardC API ... Reward mapping table
see RewardC function description below

@author
@date   2020-11-11
*/

// RewardC ...
func RewardC(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== RewardC")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	userID := reqBody.UserID
	vaultID := reqBody.VaultID
	reward := reqBody.Reward

	ok1, err1 := checkInput(userID, 1, "userID")
	ok2, err2 := checkInput(vaultID, 1, "vaultID")
	ok3, err3 := checkInput(reward, 1, "reward")
	if !(ok1 && ok2 && ok3) {
		print("input values are not valid")
		print("ok1:", ok1, ", err1:", err1)
		print("ok2:", ok2, ", err2:", err2)
		print("ok3:", ok3, ", err3:", err3)
		idx, err := getErr([]error{err1, err2, err3})
		print("err@getErr:", err)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: index " + toStr(idx) + " of [userID, VaultID, Reward] is empty or invalid",
			Data: nil,
		}, nil
	}

	if !(toInt(userID) > 0 && toInt(vaultID) > 0 && toFloat(reward) > 0) {
		print("input value is not valid: userID and vaultID should be integer, but reward should be float64")
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

	//-----------------=== check bcAddr risk
	_, timeNowStr, err := GetLocalTime(LocStr)
	if err != nil {
		print("err@ GetLocalTime", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ GetLocalTime",
			Data: nil,
		}, nil
	}

	stmtIn := "SELECT ethAddress, riskCheckedAt FROM AriesFinancial.User WHERE id = ?"
	returnSize := 2
	result, isOk, err := readRow(db, stmtIn, returnSize, userID)
	ethAddress := result[0]
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

	bcAddr := ethAddress
	msgOut := DoCallAnChainAPI("eth", bcAddr)
	if msgOut != "ok" {
		return &OutputLambda{
			Code: "110009",
			Mesg: msgOut,
		}, nil
	}

	stmtIn = "UPDATE AriesFinancial.User SET riskCheckedAt = ? WHERE id = ?"
	affectedRowID, err := writeRowV(db, stmtIn, timeNowStr, userID)
	print("affectedRowID:", affectedRowID)
	if err != nil {
		print("err@ writeRow:", err)
		return &OutputLambda{
			Code: "110001",
			Mesg: "err@ writeRow riskCheckedAt",
			Data: nil,
		}, nil
	}
	print("------== user riskCheckedAt has been updated")

	//-----------------===
	stmtIn = "INSERT INTO AriesFinancial.Reward (userID, vaultID, reward, updatedAt) VALUES (?,?,?,?)"
	RewardRowID, err := writeRowV(db, stmtIn, userID, vaultID, reward, timeNowStr)
	if err != nil {
		print("err@ writeRow() inserting new row", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ writeRowV",
			Data: nil,
		}, nil
	}
	print("new Reward row id:", RewardRowID)
	print("db writing is successful")

	// check error conditions: row not found, error existing, or delete value is valid...
	switch {
	case err != nil:
		print("err@ readRow:", err)
		return &OutputLambda{
			Code: "110001",
			Mesg: "err@ readRow",
			Data: nil,
		}, nil

	default:
		print("inserting a new Reward row is successful")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: RespRewardC{NewRowID: RewardRowID},
	}, nil
}

/*
RewardC
curl -XPOST -d '{"userID":5,"vaultID":1,"reward":2.34}' 'localhost:3000/RewardC' | jq

*/
