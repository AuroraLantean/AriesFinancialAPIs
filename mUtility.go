package main

/*@file mUtility.go
@brief various utility functions and helper functions
see SendVerifCode function description below

@author
@date
*/
import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

// Log1 ... to print logs
//var Log1 = log.Println

// print1 ... to print logs
var print = fmt.Println

// logFatal ... to print logs
var logFatal = log.Fatal

// dump ... to print structs
var dump = spew.Dump

// loadEnv ...
func loadEnv() string {
	err := godotenv.Load()
	if err != nil {
		print("cannot load env file. err:", err)
		return "cannot load env file"
	}
	return "ok"
}

// GetLocalTime ...
/*@brief to generate current time
@param out: current time
@param  in: none
*/
func GetLocalTime(locStr string) (time.Time, string, error) {
	loc, err := time.LoadLocation(locStr)
	timeNow := time.Now().In(loc)
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	print("timeNow:", timeNow, ", timeNowStr:", timeNowStr, ", err:", err)
	return timeNow, timeNowStr, err
}

// ParseTime ...
/*@brief to parse our time format
@param out: time type variable
@param  in: our time format. e.g. 2020-08-14 07:17:31
*/
func ParseTime(timeStr string) (time.Time, error) {
	const longForm = "2006-01-2 15:04:05"
	//print("ParseTime() input:", timeStr)
	t, err := time.Parse(longForm, timeStr)
	print("parsed time:", t, ", err:", err)
	return t, err
}

// lambdaFunc ...
func lambdaFunc(input InputLambda) (*OutputLambda, error) {
	print("---------------== lambdaFunc")
	//dump("input:", input)
	//rawBody := input.Body
	timeout := 5
	delay := 7

	ch1 := make(chan *OutputLambda)
	go SubLambda(ch1, input, delay)
	//print("Goroutines#:", runtime.NumGoroutine()) // => 2

	var outputPtr *OutputLambda
	var outMesg string
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		outMesg = "timeout has been reached"
		print(outMesg)
		outputPtr = &OutputLambda{
			Code: "110099",
			Mesg: outMesg,
			Data: nil,
		}
	case outputPtr = <-ch1:
		print("channel value has been returned")
	}
	return outputPtr, nil
}

// SubLambda ...
func SubLambda(ch1 chan *OutputLambda, input InputLambda, delay int) { //wg *sync.WaitGroup
	print("----------== SubLambda")
	dump(input)
	time.Sleep(time.Duration(delay) * time.Second)

	ch1 <- &OutputLambda{
		Code: "0",
		Mesg: "ok",
		Data: RespUser{
			//ID: user[0],
		},
	}
	//wg.Done() //or (*wg).Done()
}

// ExecuteRoutine ...
func ExecuteRoutine(routineInputs RoutineInputs) (*RoutineOut, error) {
	print("---------------== ExecuteRoutine")
	dump("routineInputs:", routineInputs)
	routineName := routineInputs.RoutineName
	routineAddr := routineInputs.Address
	method := routineInputs.Method
	timeout := routineInputs.Timeout

	ch1 := make(chan *RoutineOut)
	switch {
	case routineName == "MakeGetRequest":
		go MakeGetRequest(ch1, routineAddr)
	case routineName == "MakeHTTPRequest":
		go MakeHTTPRequest(ch1, routineAddr, method)
	default:
		print("routineName has no match!")
		return &RoutineOut{"110030", "function input not valid",
			"NA"}, nil
	}
	//print("Goroutines#:", runtime.NumGoroutine()) // => 2

	var RoutineOutPtr *RoutineOut
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		print("routine takes too long. timeout has been reached")
		RoutineOutPtr = &RoutineOut{"110028",
			"routine takes too long: " + toStr(timeout) + " seconds", "NA"}
	case RoutineOutPtr = <-ch1:
		print("Success. CallGoroutine() channel value has been returned")
	}
	return RoutineOutPtr, nil
}

// MakeHTTPRequest ...
func MakeHTTPRequest(ch1 chan *RoutineOut,
	requestURL string, method string) {
	print("----------== MakeHTTPRequest")
	dump(requestURL)
	client := &http.Client{}
	req, err := http.NewRequest(method, requestURL, nil)
	//resp, err := http.Get(requestURL)
	if err != nil {
		print("http.NewRequest():", err)
		ch1 <- &RoutineOut{"110023", "er@ http.NewRequest()", "NA"}
	}
	resp, err := client.Do(req)
	if resp == nil || resp.Body == nil {
		print("err@ resp or resp.Body is niil:", resp)
		ch1 <- &RoutineOut{"110035", "HTTP response is nil or its body is nil", "NA"}
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print("err@ reading response error: ioutil.ReadAll:", err)
		ch1 <- &RoutineOut{"110031", "reading response error", "NA"}
	}
	respStr := string(respBody)
	print("\nrespStr:", respStr)

	/*
		items := strings.Split(respStr, ",")
		if len(items) < 1 {
			print("err@ response length not valid")
			ch1 <- &RoutineOut{"110033", "response length not valid", respStr}
		}
		balance := toFloat(items[0])
		if balance < 0 {
			print("failed")
			ch1 <- &RoutineOut{"110034", "failed", respStr}
		}
	*/

	err = resp.Body.Close()
	if err != nil {
		print("response close resp.Body.Close():", err)
		ch1 <- &RoutineOut{"110032", "err@ resp.Body.Close()",
			respStr}
	}
	print("successful")
	ch1 <- &RoutineOut{"0", "ok", respBody}
}

// to convert int to string
func toStr(i int) string {
	return strconv.Itoa(i)
}

// to convert int to string
func toStr64(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// to convert string to int
func toInt(s string) int {
	if s == "" {
		print("err@ toInt: input string is empty:", s)
		return -111
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		print("err@ toInt: input string:", s, ", err:", err)
		return -111
	}
	return i
}

// to convert string to float
func toFloat(s string) float64 {
	if s == "" {
		print("input string is empty")
		return -111.00
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		print("err@converting string to a float. s:", s, ", err:", err)
		return -111.00
		//print(f) // bitSize is 32 for float32 convertible,
		// 64 for float64
	}
	return f
}

// MakeGetRequest ...
func MakeGetRequest(ch1 chan *RoutineOut,
	requestURL string) {
	print("----------== MakeGetRequest")
	dump(requestURL)
	resp, err := http.Get(requestURL)
	if err != nil {
		print("error@http.Get():", err)
		ch1 <- &RoutineOut{"110023", "error@http.Get()", "NA"}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print("reading response error@ioutil.ReadAll:", err)
		ch1 <- &RoutineOut{"110031", "reading response error", "NA"}
	}
	respStr := string(respBody)
	print("respStr:", respStr)

	items := strings.Split(respStr, ",")
	if len(items) < 1 {
		print("response length not valid")
		ch1 <- &RoutineOut{"110033", "response length not valid", respStr}
	}
	balance := toFloat(items[0])
	if balance < 0 {
		print("delivery failed")
		ch1 <- &RoutineOut{"110034", "delivery failed", respStr}
	}
	/*Response String:
	 */

	err = resp.Body.Close()
	if err != nil {
		print("response close error@ resp.Body.Close():", err)
		ch1 <- &RoutineOut{"110032", "response close error",
			respStr}
	}
	print("successful")
	ch1 <- &RoutineOut{"0", "OK", respStr}
}

// to check input for minimum length
func checkStrFixLength(s string, fixedLen int, inputName string) (bool, error) {
	if s == "" {
		print(inputName + " is empty")
		return false, errors.New(inputName + " is empty:")
	}
	strlen := utf8.RuneCountInString(s)
	if strlen != fixedLen {
		print(inputName, "of length", strlen, "should be of", toStr(fixedLen), "characters in length")
		return false, errors.New(inputName + " should be of " + toStr(fixedLen) + " characters in length")
	}
	print(inputName + " is valid via checkStrFixLength")
	return true, nil
}

// to check input for minimum length
func checkInput(s string, minLen int, inputName string) (bool, error) {
	if s == "" {
		print(inputName + " is empty")
		return false, errors.New(inputName + " is empty:")
	}
	if utf8.RuneCountInString(s) < minLen {
		print(inputName + " should be at least " + toStr(minLen) + " characters in length")
		return false, errors.New(inputName + " should be at least " + toStr(minLen) + " characters in length")
	}
	print(inputName + " is valid via checkInput")
	return true, nil
}

// to check for only one true value
func onlyOneIsTrue(bools []bool) bool {
	correct := false
	alreadyFound := false
	for _, v := range bools {
		if v {
			correct = true
			if alreadyFound {
				correct = false
				break
			} else {
				alreadyFound = true
			}
		}
	}
	return correct
}

func checkCharLength(s string) int {
	return utf8.RuneCountInString(s)
}

func strSliceHas(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func strSliceHasAny(s []string, e []string) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}
	return false
}

func getErr(errs []error) (int, error) {
	for idx, err := range errs {
		print(err)
		if err != nil {
			return idx, err
		}
	}
	return -1, nil
}

// to log fatal error and stop execution
func pExitErr(mesg string, err error) {
	if err != nil {
		print(mesg, err)
		os.Exit(1)
	}
}
