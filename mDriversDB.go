package main

/*@file mDriversDB.go
@brief various driver functions
see function descriptions below

@author
@date
*/
import (
	"database/sql"
	"errors"
	"os"
	"time"

	//"github.com/graphql-go/graphql"
	_ "github.com/go-sql-driver/mysql"
)

// connection to DB
// import	(_ "github.com/go-sql-driver/mysql")
func dbConn() (*sql.DB,
	error) {
	print("-----------== dbConn()")
	dbID := toInt(os.Getenv("DATABASE_ID"))
	if dbID < 0 {
		print("DATABASE_ID in .env is invalid. dbID =", dbID)
		return nil, errors.New("DATABASE_ID in .env is invalid")
	}

	var dbDriver, dbUser, dbPass, dbURL, dbPort, dbSchema string
	if dbID == 0 {
		dbDriver = "mysql"
		dbUser = os.Getenv("DBUSER0")
		dbPass = os.Getenv("DBPASS0")
		dbURL = os.Getenv("DBURL0")
		dbPort = os.Getenv("DBPORT0")
		dbSchema = os.Getenv("DBSCHEMA0")
	} else if dbID == 1 {
		dbDriver = "mysql"
		dbUser = os.Getenv("DBUSER1")
		dbPass = os.Getenv("DBPASS1")
		dbPort = os.Getenv("DBPORT1")
		dbSchema = os.Getenv("DBSCHEMA1")
		if IsProduction {
			dbURL = os.Getenv("DBURL1prodc")
		} else {
			dbURL = os.Getenv("DBURL1local")
		}
	} else {
		print("dbID is not valid")
		return nil, errors.New("dbID is not valid")
	}

	if dbDriver == "" || dbUser == "" || dbPass == "" || dbURL == "" || dbPort == "" || dbSchema == "" {
		print("dbConn() inputs are not valid:")
		print("dbDriver:", dbDriver, ", dbUser:", dbUser, ", dbPass:", dbPass, ", dbURL:", dbURL, ", dbPort:", dbPort,
			", dbSchema:", dbSchema)
		return nil, errors.New("dbConn() inputs are not valid")
	}

	db, err := sql.Open(dbDriver,
		dbUser+":"+dbPass+"@tcp("+dbURL+":"+dbPort+")/"+dbSchema)
	if err != nil {
		print("err@ sql.Open():", err)
		//panic(err.Error())
		return db, err
	}
	print("sql.Open() successful")

	err = db.Ping()
	if err != nil {
		print("err@ db.Ping()")
		return db, err
	}
	print("db.Ping() successful. err:", err)
	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no maximum limit (which is also the default setting).
	//db.SetMaxOpenConns(1)

	/* Set the maximum number of concurrently idle connections
	to 5. Setting this. takes up memory which can otherwise be
	used for both your application and the database.
	*/
	//db.SetMaxIdleConns(1)

	/*Set the maximum lifetime of a connection to 1 hour.
	 */
	db.SetConnMaxLifetime(time.Duration(MaxDBConnLifetime) * time.Second)
	print("db.SetConnMaxLifetime() successful")
	return db, err
}

func readV3(db *sql.DB, stmtIn string, args ...interface{}) (NullString, NullString, NullString, string, error) {
	print("-----------== readV3()")
	nulstr := NullString{}
	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("db.Prepare() failed. err:", err)
		return nulstr, nulstr, nulstr, "db.Prepare() failed", err
	}
	print("check1")

	var v1, v2, v3 NullString
	err = stmt.QueryRow(args...).Scan(&v1, &v2, &v3)
	print("check2")

	var isOk string
	if !v1.Valid {
		print("v1 is nil:", v1)
		isOk = "v1 is nil"
	} else if v1.String == "" {
		print("v1 is empty:", v1)
		isOk = "v1 is empty"
	} else {
		isOk = "v1ok"
	}
	print("v1:", v1, ", v2:", v2, ", v3:", v3)

	if err != nil {
		isOk = "err@ stmt.QueryRow()"
		print(isOk+", err:", err)
		return v1, v2, v3, isOk, err
	}
	err = stmt.Close()
	if err != nil {
		isOk = "err@ stmt.Close()"
		print(isOk+", err:", err)
		return v1, v2, v3, isOk, err
	}
	return v1, v2, v3, isOk, err
}

// readRow ...
func readRow(db *sql.DB, stmtIn string, returnSize int, args ...interface{}) ([]string, string, error) {
	print("-----------== readRow()")
	//nulstr := NullString{}
	nullStrSlice := make([]NullString, returnSize)
	print("returnSize:", returnSize)
	//print("nullStrSlice:", nullStrSlice)
	nss := make([]interface{}, returnSize)
	strSlice := make([]string, returnSize)
	for i := 0; i < returnSize; i++ {
		//print("i:", i)
		nss[i] = &NullString{}
		strSlice[i] = ""
		//&v will give the address of slice!
	}
	//print("nss:", nss)

	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("db.Prepare() failed. err:", err)
		return strSlice, "db.Prepare() failed", err
	}
	print("check1")
	err = stmt.QueryRow(args...).Scan(nss...)
	print("check2")
	print("s0:", (nss[0].(*NullString)))
	print("s1:", (nss[1].(*NullString)))
	var isOk string
	if err != nil {
		isOk = "err@ stmt.QueryRow()"
		print(isOk+", err:", err)
		return strSlice, isOk, err
	}

	for i, v := range nss {
		nullStrSlice[i] = *v.(*NullString)
	}
	//print("nullStrSlice at the end", nullStrSlice)

	v0 := nullStrSlice[0]
	if !v0.Valid {
		print("v0 is nil:", v0)
		isOk = "v0 is nil"
	} else if v0.String == "" {
		print("v0 is empty:", v0)
		isOk = "v0 is empty"
	} else {
		isOk = "ok"
	}

	for i, v := range nullStrSlice {
		if v.Valid {
			strSlice[i] = v.String
		}
	}

	err = stmt.Close()
	if err != nil {
		isOk = "err@ stmt.Close()"
		print(isOk+", err:", err)
		return strSlice, isOk, err
	}
	print("readRow() is successful")
	return strSlice, isOk, err
	//return nullStrSlice, isOk, err
}

func doReadRow(stmtIn string, returnSize int, inputType string,
	inputs ...interface{}) (string, string,
	error) {
	print("-----------== doReadRow()")

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		pExitErr("isEnvOk:"+isEnvOk, nil)
	}

	db, err := dbConn()
	if err != nil {
		pExitErr("err@dbConn()", err)
	}
	print("dbConn() connection is ok")

	var accountID string
	var strSlice []string
	var readOk string
	if inputType == "email" {
		print("input is email!")
		stmtInE := "SELECT * FROM email WHERE email_address = ?"
		outE, readOkE, errE := readRow(db, stmtInE, 6, "user1@domain.com")
		if errE != nil {
			readOk = "err@ readRow()"
			print(readOk+", errE:", errE)
			return "NA", readOk, err
		}
		accountID = outE[1]
		print("outE:", outE, ", readOkE:", readOkE, ", errE:",
			errE, ", accountID:", accountID)
		strSlice, readOk, err = readRow(db,
			"SELECT * FROM account WHERE id = ?", 16, accountID)

	} else {
		strSlice, readOk, err = readRow(db, stmtIn, returnSize, inputs)
	}

	if err != nil {
		readOk = "err@ readRow(accountID)"
		print(readOk+", err:", err)
		return "NA", readOk, err
	}
	err = db.Close()
	if err != nil {
		readOk = "err@ db.Close()"
		print(readOk+", err:", err)
		return "NA", readOk, err
	}

	//print("account id:", strSlice[0].String, strSlice[15].String, "readOk:", readOk, ", err:", err)
	//print("avatar :", strSlice[11].String, ", token:", strSlice[12].String)
	//createdAt, _, _, readOk, err := readV3(db, "SELECT created_at, updated_at, id FROM account WHERE account_id = ?", accountID)
	//print("createdAt:", createdAt, ", readOk:", readOk, ", err:", err)
	return strSlice[0], readOk, err
}

// readTableXrow ...
/*@brief to find a row with matched queryValue
@param nullStrSlice: row data
@param  in: execution sql statement, search condition
*/
func readTableXrow(db *sql.DB, stmtIn string, queryValue string) (TableNameX, error) {
	print("-----------== readTableXrow()")
	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("err@ db.Prepare(), err:", err)
		return TableNameX{}, err
	}
	print("check1, queryValue:", queryValue)

	row := stmt.QueryRow(queryValue)
	print("row:", row)

	item := TableNameX{}
	err = row.Scan(&item.ID, &item.EthereumAddr)
	dump(item)

	if err != nil {
		print("err@ row.Scan(), err:", err)
		return item, err
	}
	err = stmt.Close()
	if err != nil {
		print("err@ stmt.Close(), err:", err)
		return item, err
	}
	return item, err
}

// readRewards ...
/*@brief to find verification rows
@param out: row data
@param  in: execution sql statement, search condition
*/
func readRewards(db *sql.DB, stmtIn string, args ...interface{}) ([]Reward, error) {
	print("-----------== readRewards()")
	items := make([]Reward, 0)
	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("err@ db.Prepare(), err:", err)
		return items, err
	}
	print("check1")

	rows, err := stmt.Query(args...)
	if err == sql.ErrNoRows {
		print("no row is found:")
		return items, nil
	} else if err != nil {
		print("err@ stmt.Query():", err)
		return items, err
	}
	print("rows:", rows)
	if rows == nil {
		return items, nil
	}

	for rows.Next() {
		item := Reward{}
		err = rows.Scan(&item.ID, &item.UserID, &item.VaultID, &item.Reward, &item.UpdatedAt)
		if err != nil {
			print("err@ rows.Scan():", err)
			return items, err
		}
		items = append(items, item)
	}

	err = stmt.Close()
	if err != nil {
		print("err@ stmt.Close(), err:", err)
		return items, err
	}
	return items, nil
} // &item.EthereumAddr,

/*@brief to write data into db
@param nullStrSlice: string, error
@param  in: db, execution sql statement, values to be written into
*/
func writeRowV(db *sql.DB, stmtIn string, args ...interface{}) (string, error) {
	print("-----------== writeRowV()")
	//print("dbInput:", dbInput)
	dump(args...)
	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("err@ db.Prepare(), err:", err)
		return "err@db.Prepare()", err
	}
	print("check0")

	if len(args) == 0 {
		return "err", errors.New("writeRowV() has no args")
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		print("err @stmt.Exec():", err)
		return "err", err
	}

	rowAff, err := result.RowsAffected()
	print("affected row:", rowAff)
	if err != nil {
		print("err @RowsAffected():", err)
		return "err", err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		print("err @lastInsertID():", err)
		return "err", err
	}

	err = stmt.Close()
	if err != nil {
		print("err@ stmt.Close(), err:", err)
		return toStr64(lastInsertID), err
	}
	return toStr64(lastInsertID), err
}

func testDB() {
	print("-----------== testDB()")
	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		print("err@ loadEnv")
	}
	print("check1 loadEnv is ok")

	db, err := dbConn()
	if err != nil {
		print("err@dbConn():", err)
	}
	print("check2 db connection is ok")

	defer func() {
		errDbClose := db.Close()
		if errDbClose != nil {
			print("err@ db.Close:", errDbClose)
			return
		}
	}()
}

/*type writeRowArg struct {
	APYboDay ApyBoDP `json:"apybodp"`
	APYboWeek ApyBoWP `json:"apybowp"`
	APYboMonth ApyBoMP `json:"apybomp"`
}*/
func addRowDB(inputLambda InputLambda) (*OutputLambda, error) {
	print("-----------== writeRowS()")
	//print("dbInput:", dbInput)
	dump(inputLambda)
	// writeRowArg := WriteRowArg{
	// 	SrcPlusPP: "YearnFinance_Week",
	// }
	reqBody := inputLambda.Body
	if inputLambda.DataName == "" || reqBody.PerfPeriod == "" {
		return &OutputLambda{
			Code: "000103",
			Mesg: "err@ dataName or perfPeriod is empty",
		}, nil
	}
	print("check3 rowName is ok")
	rowName := inputLambda.DataName + "_" + reqBody.PerfPeriod

	isEnvOk := loadEnv()
	if isEnvOk != "ok" {
		return &OutputLambda{
			Code: "000101",
			Mesg: "err@ loadEnv",
		}, nil
	}
	print("check1 loadEnv is ok")

	db, err := dbConn()
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

	stmtIn := "INSERT INTO apy (sourceName) VALUES (?)"
	//"INSERT INTO account (account_name, password, nickname, avatar) VALUES (?,?,?,?)"

	stmt, err := db.Prepare(stmtIn)
	if err != nil {
		print("err@ db.Prepare():", err)
		return &OutputLambda{
			Code: "000104",
			Mesg: "err@db.Prepare()",
		}, err
	}
	print("check4 prepare stmt is ok")

	defer func() {
		errStmtClose := stmt.Close()
		if errStmtClose != nil {
			print("err@ stmt.Close():", errStmtClose)
			return
		}
	}()

	//aw := args.APYboWeek
	result, err := stmt.Exec(rowName)
	if err != nil {
		print("err @stmt.Exec():", err)
		return &OutputLambda{
			Code: "000105",
			Mesg: "err@ stmt.Exec",
		}, err
	}
	print("check5 stmt.Exec is ok")

	rowAff, err := result.RowsAffected()
	print("affected row:", rowAff)
	if err != nil {
		print("err @RowsAffected:", err)
		return &OutputLambda{
			Code: "000106",
			Mesg: "err@ RowsAffected",
		}, err
	}
	print("check6 RowsAffected is ok")

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		print("err @lastInsertID():", err)
		return &OutputLambda{
			Code: "000107",
			Mesg: "err@ LastInsertId"}, err
	}

	print("\ndb writing is successful")
	return &OutputLambda{
		Code: "000000",
		Mesg: "ok",
		Data: toStr64(lastInsertID),
	}, nil
}
