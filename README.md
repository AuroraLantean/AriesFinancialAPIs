# Aries Financial API Server

## HTTP request handlers
### func updateHttp(w http.ResponseWriter, r *http.Request)
- to extract request arguments, pass them down to updateLambda function below

### func httpWriteRow(w http.ResponseWriter, r *http.Request)
- to write fixed data into database

### func readHttp(w http.ResponseWriter, r *http.Request)
- to extract request arguments, pass them down to readLambda function below

### func ping(w http.ResponseWriter, r *http.Request) 
- to test server and db connection

### func root(w http.ResponseWriter, r *http.Request) 
- to test server

## Sub-Functions called from HTTP request handlers
### func readLambda(inputLambda InputLambda) (*([]VaultAPY), error) 
- to take input arguments, make fixed query about certain APYs on database 

### func updateLambda(inputLambda InputLambda) (*OutputLambda, error)
- to take input arguments, call scraper function on certain data source, call regular express to extract data, write data into database, send response

## Database functions
### func readRow(db *sql.DB, stmtIn string, returnSize int, args ...interface{}) ([]string, string, error)
- read from database

### func writeRowV(db *sql.DB, stmtIn string, args ...interface{}) (string, error)
- write data into database

<br>

## Installation
Install Go 1.15.3

Turn on Go module: <br>
$ export GO111MODULE=on

## Download this repo then setup dependencies
$ git clone https://github.com/aries-financial-defi/AriesFinancialAPIs.git  <br>
$ go mod tidy

## Run the API server
$ go run *.go

## Only build(produce an executable file)
For Linux 64 bit target environment(for example: AWS EC2): 
Run goBuild.sh:  <br>
$ ./goBuild.sh, and it will generate main.zip, which includes the binary and .env file for deployment

For producing an executable binary file and run it locally: <br>
$ go build -O main.go <br>
$ ./main

## Test Locally
Read APYs from database: <br>
$ curl -v 'localhost:3000/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week'

Fetch data from source then update APYs in the database: <br>
$ curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/update'

Write fixed data into database: <br>
$ curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpWriteRow'

Test server and database connection: <br>
$ curl 'localhost:3000/ping'

Test server: <br>
$ curl 'localhost:3000'


## Domain of Deployment
Read APYs from database: <br> [https://api.aries.financial/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week](https://api.aries.financial/vaults/apy?sourceURL=https://stats.finance/yearn&perfPeriod=week)

Fetch data from source then update APYs in the database:<br>
POST https://api.aries.financial/update <br>
Body: {"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}

<br>
Write fixed data into database:<br>
POST https://api.aries.financial/httpWriteRow <br>
Body: {"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}

<br>
<br>
Test server and database connection [https://api.aries.financial/ping](https://api.aries.financial/ping)
<br>
Test server [https://api.aries.financial](https://api.aries.financial)

