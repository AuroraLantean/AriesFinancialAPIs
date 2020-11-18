package main

/*@file UserC.go
@brief UserC API
see UserC function description below

@author
@date   2020-11-11
*/
import (
	"database/sql"
)

// UserC ...
func UserC(inputLambda InputLambda) (*OutputLambda, error) {
	print("---------------== UserC")
	dump("inputLambda.Body:", inputLambda.Body)
	reqBody := inputLambda.Body

	ethAddress := reqBody.EthereumAddr
	ok1, err1 := checkInput(ethAddress, 42, "Ethereum Address")

	if !(ok1) {
		print("input value is not valid")
		print("ok1:", ok1, ", err1:", err1)
		idx, err := getErr([]error{err1})
		print("err@getErr:", err)
		return &OutputLambda{
			Code: "110000",
			Mesg: "API input not valid: index " + toStr(idx) + " of [Ethereum Address] is empty or invalid",
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

	//-----------------===
	stmtIn := "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs, riskCheckedAt FROM AriesFinancial.User WHERE ethAddress = ?"
	returnSize := 7
	userRow, isOk, err := readRow(db, stmtIn, returnSize, ethAddress)

	// check error conditions: row not found, error existing, or delete value is valid...
	var userID string
	switch {
	case err == sql.ErrNoRows:
		print("row not found")

		stmtIn = "INSERT INTO AriesFinancial.User (ethAddress, updatedAt, riskCheckedAt) VALUES (?,?,?)"
		userID, err = writeRowV(db, stmtIn, ethAddress, timeNowStr, timeNowStr)
		if err != nil {
			print("err@ writeRow() inserting new row", err)
			return &OutputLambda{
				Code: "000104",
				Mesg: "err@ writeRowV",
			}, nil
		}
		print("userID:", userID)
		print("adding userRow is successful")
		return &OutputLambda{
			Code: "0",
			Mesg: "ok",
			Data: RespUser{
				ID:            userID,
				EthereumAddr:  ethAddress,
				UpdatedAt:     timeNowStr,
				RiskCheckedAt: timeNowStr,
			},
		}, nil
		/*
			stmtIn = "SELECT id, ethAddress, reward, updatedAt, uplineID, downlineIDs FROM AriesFinancial.User WHERE id = ?"
			returnSize := 6
			userRow, isOk, err = readRow(db, stmtIn, returnSize, userID)
			switch {
			case err == sql.ErrNoRows:
				print("row not found")
				return &OutputLambda{
					Code: "110001",
					Mesg: "err@ row not found",
				}, nil

			case isOk != "ok":
				print("err@ isOk:", err)
				return &OutputLambda{
					Code: "110001",
					Mesg: "err@ isOk: " + isOk,
				}, nil

			case err != nil:
				print("err@ readRow:", err)
				return &OutputLambda{
					Code: "110001",
					Mesg: "err@ readRow",
				}, nil

			default:
				print("newly added userRow:", userRow)
			}*/

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
UserC
curl -XPOST -d '{"ethereumAddr":"0x054f48Ae455dcf918F75bD28e8256Fd6fb02d27f"}' 'localhost:3000/UserC' | jq

*/
