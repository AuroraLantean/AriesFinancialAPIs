package main

import "database/sql"

/*@file RewardD.go
@brief RewardD API ... Reward mapping table
see RewardD function description below

@author
@date   2020-11-11
*/

// RewardD ...
func RewardD(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== RewardD")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	userID := reqBody.UserID
	vaultID := reqBody.VaultID
	rewardID := reqBody.RewardID

	ok1, err1 := checkInput(userID, 1, "userID")
	ok2, err2 := checkInput(vaultID, 1, "vaultID")
	ok3, err3 := checkInput(rewardID, 1, "rewardID")
	if ok3 {
		if toInt(rewardID) <= 0 {
			print("rewardID is not valid")
			return &OutputLambda{
				Code: "110000",
				Mesg: "API input not valid:",
				Data: nil,
			}, nil
		}

	} else if ok1 && ok2 {
		if toInt(userID) <= 0 || toInt(vaultID) <= 0 {
			print("userID and vaultID are not valid")
			return &OutputLambda{
				Code: "110000",
				Mesg: "API input not valid:",
				Data: nil,
			}, nil
		}

	} else {
		print("input values are not valid")
		print("ok1:", ok1, ", err1:", err1)
		print("ok2:", ok2, ", err2:", err2)
		print("ok3:", ok3, ", err3:", err3)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: rewardID or [userID, vaultID] is empty or invalid",
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

	var rewardRow []string
	var returnSize int
	var ethAddress, stmtIn, isOk string
	if ok3 {
		stmtIn1 := "SELECT id, userID, vaultID, reward, updatedAt FROM AriesFinancial.Reward WHERE id = ?"
		//SELECT * FROM table limit 5, 10 ... row 6~15
		returnSize = 5

		rewardRow, isOk, err = readRow(db, stmtIn1, returnSize, rewardID)
		userID := rewardRow[1]
		// check error conditions: row not found, error existing, or delete value is valid...
		switch {
		case err == sql.ErrNoRows:
			print("row not found")
			return &OutputLambda{
				Code: "110002",
				Mesg: "row not found",
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
			print("readRow is ok")
		}
		//------------== risk check
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

	} else if ok1 && ok2 {
		print("---------------== ok1 && ok2")
		stmtIn = "SELECT ethAddress, riskCheckedAt FROM AriesFinancial.User WHERE id = ?"
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

		stmtIn = "SELECT id, userID, vaultID, reward, updatedAt FROM AriesFinancial.Reward WHERE userID = ? AND vaultID = ?"
		//SELECT * FROM table limit 5, 10 ... row 6~15
		returnSize = 5

		rewardRow, isOk, err = readRow(db, stmtIn, returnSize, userID, vaultID)
		// check error conditions: row not found, error existing, or delete value is valid...
		switch {
		case err == sql.ErrNoRows:
			print("row not found")
			return &OutputLambda{
				Code: "110002",
				Mesg: "row not found",
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
			print("readRow is ok")
		}

	} else {
		print("ok3 or ok1&&ok2 is false")
	}

	//-----------------=== check risk
	bcAddr := ethAddress
	msgOut := DoCallAnChainAPI("eth", bcAddr)
	if msgOut != "ok" {
		return &OutputLambda{
			Code: "110009",
			Mesg: msgOut,
		}, nil
	}

	//-----------------===
	if ok3 {
		stmtIn = "DELETE FROM AriesFinancial.Reward WHERE id = ?"
		RewardRowID, err := writeRowV(db, stmtIn, rewardID)
		if err != nil {
			print("err@ writeRow() inserting new row", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ writeRowV",
				Data: nil,
			}, nil
		}
		print("affected Reward row number:", RewardRowID)
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
	} else if ok1 && ok2 {
		stmtIn = "DELETE FROM AriesFinancial.Reward WHERE userID = ? AND vaultID = ?"
		RewardRowID, err := writeRowV(db, stmtIn, userID, vaultID)
		if err != nil {
			print("err@ writeRow() inserting new row", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ writeRowV",
				Data: nil,
			}, nil
		}
		print("affected Reward row number:", RewardRowID)
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
	} else {
		print("ok3 or ok1&&ok2 is false")
	}

	return &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: nil,
	}, nil
}

/*
RewardD
curl -XPOST -d '{"userID":"1","vaultID":"1"}' 'localhost:3000/rewarddelete' | jq

*/
