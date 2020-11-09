package main

/*
 */
func updateLambda(inputLambda InputLambda) (*OutputLambda, error) {
	print("-----------== updateLambda()")
	dump(inputLambda)
	perfPeriod := inputLambda.Body.PerfPeriod
	dataName := inputLambda.DataName
	if dataName == "" || perfPeriod == "" {
		return &OutputLambda{
			Code: "000100",
			Mesg: "err@ dataSource or perfPeriod is empty",
		}, nil
	}
	print("check0 DataName and PerfPeriod are ok")

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		return &OutputLambda{
			Code: "000101",
			Mesg: "err@ loadEnv",
		}, nil
	}
	print("check1 loadEnv is ok")

	db, err := dbConn(DatabaseID)
	if err != nil {
		print("err@dbConn():", err)
		return &OutputLambda{
			Code: "000102",
			Mesg: "err@ dbConn",
		}, nil
	}
	print("check2 db connection is ok")

	defer func() {
		errDbClose := db.Close()
		if errDbClose != nil {
			print("err@ db.Close:", errDbClose)
			return
		}
	}()

	//-----------------===
	srcAndPP := dataName + "_" + perfPeriod

	stmtIn := "UPDATE AriesFinancial.APY SET WETH = ?, AFI = ?, YFI = ?, CRV3 = ?, CRVY  = ?, CRVBUSD = ?, CRVSBTC = ?, DAI = ?, TrueUSD = ?, USDC = ?, Gemini = ?, TetherUSD = ? WHERE srcAndPP = ?"
	//"UPDATE account SET name = ?, time = ? WHERE id = ?"

	var ap APYs
	switch {
	case perfPeriod == "day":
		ap = inputLambda.APYboDay
	case perfPeriod == "week":
		ap = inputLambda.APYboWeek
	case perfPeriod == "month":
		ap = inputLambda.APYboMonth
	default:
		print("err@ perfPeriod invalid")
		return &OutputLambda{
			Code: "000103",
			Mesg: "err@ perfPeriod invalid",
		}, nil
	}

	updatedRow, err := writeRowV(db, stmtIn, ap.WETH, ap.AFI, ap.YFI, ap.CRV3, ap.CRVY, ap.CRVBUSD, ap.CRVSBTC, ap.DAI, ap.TrueUSD, ap.USDC, ap.Gemini, ap.TetherUSD, srcAndPP)
	//"INSERT INTO account (name, mobile, pic) VALUES (?,?,?)"
	if err != nil {
		print("err@ writeRow() insert new row", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@ writeRowV",
			Data: nil,
		}, nil
	}
	print("updatedRow:", updatedRow)

	print("\ndb writing is successful")
	return &OutputLambda{
		Code: "000000",
		Mesg: "ok",
		Data: updatedRow,
	}, nil
}
