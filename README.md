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

## Test Locally on User Table APIs
if using the 2016 DAO hacker address with bad score: "0x304a554a310c7e546dfe434669c62820b7d83490", any API below should return error

------------== UserC ... Add User<br>
curl -XPOST -d '{"ethereumAddr":"your_eth_address"}' 'localhost:3000/user' | jq

curl -XPOST -d '{"ethereumAddr":"0x304a554a310c7e546dfe434669c62820b7d83490"}' 'localhost:3000/user' | jq

------------== UserR ... Read User<br>
Use either userID or ethAdress as param<br>
curl 'localhost:3000/user?userID=4' | jq

curl 'localhost:3000/user?ethAddress=your_eth_address' | jq

------------== UserU ... Update User<br>
Use either userID or ethAdress as param<br>
curl -XPUT -d '{"ethereumAddr":"your_eth_address","reward":"3.24"}' 'localhost:3000/user' | jq

curl -XPUT -d '{"ethereumAddr":"0x304a554a310c7e546dfe434669c62820b7d83490","reward":"3.24"}' 'localhost:3000/user' | jq

------------== UserD ... Delete User<br>
Use either userID or ethAdress as param<br>
curl -XDELETE -d '{"ethereumAddr":"your_eth_address"}' 'localhost:3000/user' | jq

curl -XDELETE -d '{"userID":"5"}' 'localhost:3000/user' | jq

------------== UsersC ... Add Multiple Users<br>
curl -XPOST -d ' {"ethereumAddrs":["addr1","addr2","addr3"]}' 'localhost:3000/users' | jq

## Test Locally on Reward Table APIs
------------== Reward C ... Add Reward<br>
curl -XPOST -d '{"userID":"6","vaultID":"1","reward":"3.33"}' 'localhost:3000/reward' | jq

------------== Reward R ... Read Reward<br>
Use either rewardID or (userID and vaultID) as param<br>
curl 'localhost:3000/reward?userID=6&vaultID=1' | jq

curl 'localhost:3000/reward?rewardID=11' | jq

------------== Reward D ... Delete Reward<br>
Use either rewardID or (userID and vaultID) as param<br>
curl -XDELETE -d '{"rewardID":"6"}' 'localhost:3000/reward' | jq

curl -XDELETE -d '{"userID":"4","vaultID":"5"}' 'localhost:3000/reward' | jq


## Test Locally on VaultEthAddr Table APIs
------------== VaultEthAddr R ... Read VaultEthAddr<br>
curl 'localhost:3000/vaultethaddr?userID=1&vaultID=1' | jq

## Test Locally on Adding An User Array
curl -XPOST -d ' {"ethereumAddrs":["ethAddr1","ethAddr2","ethAddr3"]}' 'localhost:3000/users' | jq


## Test Locally on data fetching APIs from source to DB
Read APYs from database: <br>
$ curl -v 'localhost:3000/vaults/apy'

Fetch data from source then update APYs in the database: <br>
$ curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/update'

Write fixed data into database: <br>
$ curl -XPOST -d '{"sourceURL":"https://stats.finance/yearn","perfPeriod":"week"}' 'localhost:3000/httpWriteRow'

Test server and database connection: <br>
$ curl 'localhost:3000/ping'

Test server: <br>
$ curl 'localhost:3000'


## Domain of Deployment
Read APYs from database: <br> [https://api.aries.financial/vaults/apy](https://api.aries.financial/vaults/apy)

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

<br>
## To Do
Make a config file to set smart contract addresses, descriptions, symbols, etc...