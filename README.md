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
$ ./run.sh

## Zip binary and config and .env file
$ ./goZip.sh

## Clean log file:
$ ./goClean.sh

## Only build(produce an executable file)
For Linux 64 bit target environment(for example: AWS EC2): 
Run goBuild.sh:  <br>
$ ./goBuild.sh, and it will generate main.zip, which includes the binary and .env file for deployment

For producing an executable binary file and run it locally: <br>
$ go build -O main.go <br>
$ ./main

## Test Locally on User Table APIs
if using the 2016 DAO hacker address with bad score: "0x304a554a310c7e546dfe434669c62820b7d83490", any API below should return error

------------== Create User<br>
request type: POST, endpoint: /member <br>
request body: {"ethereumAddr":"your_eth_address"}<br>

------------== Read User<br>
request type: GET, endpoint: /member<br>
request params: userID or ethAddress<br>

/member?ethAddress=your_eth_address<br>
/member?userID=4<br>

------------== Update User<br>
request type: PUT, endpoint: /member <br>
request body: Use either userID or ethAdress<br>

{"ethereumAddr":"your_eth_address","reward":"3.24"}<br>
{"ethereumAddr":"0x304a554a310c7e546dfe434669c62820b7d83490","reward":"3.24"}<br>

------------== Delete User<br>
request type: DELETE, endpoint: /member <br>
request body: Use either userID or ethAdress<br>

{"ethereumAddr":"your_eth_address"}<br>
{"userID":"5"}<br>

------------== Get Rewards Pool APY<br>
request type: GET, endpoint: /ariesapy<br>
request params: reward contract address<br>

/ariesapy?rewardspool={rewards contract address}

------------== Read Reward<br>
Use either rewardID or (userID and vaultID) as param<br>
curl 'localhost:3000/reward?userID=6&vaultID=1' | jq

curl 'localhost:3000/reward?rewardID=11' | jq

------------== Delete Reward<br>
Use either rewardID or (userID and vaultID) as param<br>
curl -XDELETE -d '{"rewardID":"6"}' 'localhost:3000/reward' | jq

curl -XDELETE -d '{"userID":"4","vaultID":"5"}' 'localhost:3000/reward' | jq





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