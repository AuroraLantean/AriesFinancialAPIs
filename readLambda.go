package main

import (
	"database/sql"
)

/*
curl -v http://localhost:3000/readHttp?dataName=yearnFinance&perfPeriod=week | jq
*/
func readLambda(inputLambda InputLambda) (*([]VaultAPY), error) {
	print("-----------== readLambda()")
	dump(inputLambda)
	perfPeriod := inputLambda.Body.PerfPeriod
	dataName := inputLambda.DataName
	sliceOut := []VaultAPY{}
	if dataName == "" || perfPeriod == "" {
		print("err@ perfPeriod or dataName not valid:", perfPeriod, dataName)
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000100",
		// 	Mesg: "err@ dataSource or perfPeriod is empty",
		// }, nil
	}
	print("check0 DataName and PerfPeriod are ok")

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		print("err@ loadEnv():", isEnvOk)
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000101",
		// 	Mesg: "err@ loadEnv",
		// }, nil
	}
	print("check1 loadEnv is ok")

	db, err := dbConn()
	if err != nil {
		print("err@dbConn():", err)
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000102",
		// 	Mesg: "err@ dbConn",
		// }, nil
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

	//SELECT * FROM account_email WHERE mail_address = ?
	stmtIn := "SELECT id, srcAndPP, WETH, AFI, CRV3, CRVY, CRVBUSD, CRVSBTC, DAI, TrueUSD, USDC, TetherUSD FROM AriesFinancial.APY WHERE srcAndPP = ?"
	//YFI, Gemini,
	symbols := []string{"WETH", "AFI", "CRV3", "CRVY", "CRVBUSD", "CRVSBTC", "DAI", "TrueUSD", "USDC", "TetherUSD"} //"YFI", "Gemini",

	returnSize := len(symbols) + 2 //12 // columns

	addrs := []string{"0x25B192d931dD8e473A2F2B53D8BB02b83aE6A4b0", "0xdE726E878373A321d788e361a368F26AB398A7D4", "0xf908a9B8Bc339221813Af9C7E380CE845964E266", "0x47561aADd55b829C9756CD8fE0016eCAD88dFbDC", "0x129C86C01abAE3d2C90B4507E62B33F0617ccB34", "0xFfD0662a840bdE1403CDcc090Fc7157b06c86219", "0xe0469D912c781e727a365fE89D8BcfF0de654BB7", "0x218911E240f4CCAEa0839e3f1f992E3aCb692Ad6", "0x0279eF39C3029af541cbabCF8e83Afa0c96E8782", "0xfdAF86cBa91672e81dA03D2f4Fe951505EE4F468"}

	names := []string{"WETH", "Aries.Financial", "Curve.fi/3pool LP", "Curve.fi/y LP", "Curve.fi/busd LP", "Curve.fi/sbtc LP", "DAI", "TrueUSD", "USD Coin", "TetherUSD"}
	// "YFI", "Gemini",
	tokenAddrs := []string{"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xdE726E878373A321d788e361a368F26AB398A7D4", "0x6c3f90f043a72fa612cbac8115ee7e52bde6e490", "0xdf5e0e81dff6faf3a7e52ba697820c5e32d806a8", "0x3B3Ac5386837Dc563660FB6a0937DFAa5924333B", "0x075b1bb99792c9E1041bA13afEf80C91a1e70fB3", "0x6b175474e89094c44da98b954eedeac495271d0f", "0x0000000000085d4780B73119b644AE5ecd22b376", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdAC17F958D2ee523a2206206994597C13D831ec7"}

	description := []string{"Wrappeth Ether", "Aries.Financial token", "Curve 3pool LP Token", "Curve Y pool LP Token", "Curve bUSD pool LP Token", "Curve renBTC pool LP Token", "DAI Stable coin", "TrueUSD", "USD Coin", "TetherUSD"}
	//"yearn.finance token", "Gemini Dollar",
	vaultAddrs := []string{"0x25B192d931dD8e473A2F2B53D8BB02b83aE6A4b0", "0xdE726E878373A321d788e361a368F26AB398A7D4", "0xf908a9B8Bc339221813Af9C7E380CE845964E266", "0x47561aADd55b829C9756CD8fE0016eCAD88dFbDC", "0x129C86C01abAE3d2C90B4507E62B33F0617ccB34", "0xFfD0662a840bdE1403CDcc090Fc7157b06c86219", "0xe0469D912c781e727a365fE89D8BcfF0de654BB7", "0x218911E240f4CCAEa0839e3f1f992E3aCb692Ad6", "0x0279eF39C3029af541cbabCF8e83Afa0c96E8782", "0xfdAF86cBa91672e81dA03D2f4Fe951505EE4F468"}

	if len(symbols) == len(addrs) && len(addrs) == len(names) && len(names) == len(tokenAddrs) && len(tokenAddrs) == len(description) && len(description) == len(vaultAddrs) {
		print("All vault slices have the same length")
	} else {
		print("err@ found one vault slice has different length")
		return &sliceOut, nil
	}

	//ETH_DAI, ETH_USDC, ETH_USDT, ETH_WBTC, CRV_RENWBTC, WBTC, RENBTC, WBTC_TBTC, FARM
	//Curvefi3pool, Curvefiy,
	srcAndPP := dataName + "_" + perfPeriod
	out, readOk, err := readRow(db, stmtIn, returnSize, srcAndPP)
	switch {
	case err == sql.ErrNoRows:
		print("row not found")
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000104",
		// 	Mesg: "row not found for " + srcAndPP,
		// }, nil

	case err != nil:
		print("err@ readRow():", err)
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000105",
		// 	Mesg: "cannot read from database. Table does not exist or column name mismatch",
		// }, nil

	case readOk != "ok":
		print("err@ row.Scan(), readOk:", readOk, " err:", err)
		return &sliceOut, nil
		// return &OutputLambda{
		// 	Code: "000106",
		// 	Mesg: "err@ stmt.Exec. " + readOk,
		// }, err

	default:
		print("row is found")
	}
	print("check4 readRow is ok")
	dump("readRow result:", out)

	//-----------------== make output format
	print("\ndb reading is successful")
	//sliceOut := []VaultAPY{}
	var idxD int
	for idx, value := range out {
		if idx > 1 {
			idxD = idx - 2
			sliceOut = append(sliceOut, VaultAPY{
				ApyOneWeekSample: toFloat(value),
				Symbol:           symbols[idxD],
				Address:          addrs[idxD],
				Name:             names[idxD],
				TokenAddress:     tokenAddrs[idxD],
				Description:      description[idxD],
				VaultAddress:     vaultAddrs[idxD],
			})
		}
	}
	return &sliceOut, nil

	// return &OutputLambda{
	// 	Code: "000000",
	// 	Mesg: "ok",
	// 	Data: out,
	// }, nil
}
