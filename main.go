package main

/*@file main.go
@brief to determine which function to execute
*/

import (
	"log"
	"net/http"
)

// LocStr ...
var LocStr = "Asia/Taipei"

// MaxDBConnLifetime ...
var MaxDBConnLifetime = 5

// DatabaseID ...
var DatabaseID = 0 // 0: AWS RDS, 1: AWS Proxy DB

// IsProduction ... to set productin mode
var IsProduction = 1 // then set c value in the switch

// IsToFetch ...
var IsToFetch = true

// PortStr ...
var PortStr = "3000"

func main() {
	port := ":" + PortStr

	http.HandleFunc("/update", updateHTTP)
	http.HandleFunc("/vaults/apy", readHTTP)
	http.HandleFunc("/httpAddRow", httpAddRow)
	http.HandleFunc("/httpWriteRow", httpWriteRow)
	http.HandleFunc("/httpDeleteRow", httpDeleteRow)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", root)
	print("listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

/*
https://stats.finance/robots.txt ... ok


*/
