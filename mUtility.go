package main

/*@file mUtility.go
@brief various utility functions and helper functions
see SendVerifCode function description below

@author
@date
*/
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/dlclark/regexp2"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

// Log1 ... to print logs
//var Log1 = log.Println

// print1 ... to print logs
var print = fmt.Println

// logFatal ... to print logs
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logErr(str string, err error) {
	if err != nil {
		log.Println(str, err)
	}
}

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

func respondWithError(w http.ResponseWriter, status int, err Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	logErr("Error @ RespondWithJSON: ", err)
}

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
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
	bodyStr := routineInputs.Body
	method := routineInputs.Method
	timeout := routineInputs.Timeout

	ch1 := make(chan *RoutineOut)
	switch {
	case routineName == "MakeHTTPGET":
		go MakeHTTPGET(ch1, routineAddr)
	case routineName == "MakeHTTPRequest":
		go MakeHTTPRequest(ch1, routineAddr, method)
	case routineName == "MakeHTTPPOST":
		go MakeHTTPPOST(ch1, routineAddr, bodyStr)
	case routineName == "MakeGraphqlRequest":
		go MakeGraphqlRequest(ch1, routineAddr, bodyStr)

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

// MakeGraphqlRequest ...
func MakeGraphqlRequest(ch1 chan *RoutineOut,
	requestURL string, bodyStr string) {
	print("----------== MakeGraphqlRequest")
	dump(requestURL, bodyStr)
	//requestBody := strings.NewReader(bodyStr)
	jsonData := map[string]string{
		"query": bodyStr}

	/*
		jsonData := map[string]string{
			"query": `{	pair(id: "0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891") {
				token0 {
					id
					symbol
					derivedETH
				}
				token1 {
					id
					symbol
					derivedETH
				}
				token0Price
				token1Price
			}
			}`,
		}*/
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(request)

	if err != nil {
		print("http.Post():", err)
		ch1 <- &RoutineOut{"110023", "er@ http.Post()", "NA"}
	}
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

	err = resp.Body.Close()
	if err != nil {
		print("response close resp.Body.Close():", err)
		ch1 <- &RoutineOut{"110032", "err@ resp.Body.Close()",
			respStr}
	}
	print("successful@ MakeGraphqlRequest")
	ch1 <- &RoutineOut{"0", "ok", respBody}
}

// MakeHTTPPOST ...
func MakeHTTPPOST(ch1 chan *RoutineOut,
	requestURL string, bodyStr string) {
	print("----------== MakeHTTPPOST")
	dump(requestURL, bodyStr)
	//requestBody := strings.NewReader(bodyStr)

	requestBody := strings.NewReader(`{	pair(id: "0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891") {
			token0 {
				id
				symbol
				derivedETH
			}
			token1 {
				id
				symbol
				derivedETH
			}
			token0Price
			token1Price
		}
		}`)
	print("requestBody:", requestBody)
	resp, err := http.Post(
		requestURL,
		"application/json; charset=UTF-8",
		requestBody,
	)

	if err != nil {
		print("http.Post():", err)
		ch1 <- &RoutineOut{"110023", "er@ http.Post()", "NA"}
	}
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
	print("successful@ MakeHTTPPOST")
	ch1 <- &RoutineOut{"0", "ok", respBody}
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

func float64ToBigInt(val float64, mag int64) *big.Int {
	vBF := new(big.Float)
	vBF.SetFloat64(val)
	// Set precision if required.
	// vBF.SetPrec(64)

	magBF := new(big.Float)
	magBF.SetInt(big.NewInt(mag))

	vBF.Mul(vBF, magBF)

	result := new(big.Int)
	vBF.Int(result) // store converted number in result
	return result
}

// MakeHTTPGET ...
func MakeHTTPGET(ch1 chan *RoutineOut,
	requestURL string) {
	print("----------== MakeHTTPGET")
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

func doregexp2FindInBtw(ss []string, regexpStr string) (PairData, error) {
	pairData := PairData{}
	for idx, v := range ss {
		print("idx", idx, ":", v)
		out, err := regexp2FindInBtw(v, regexpStr)
		if err != nil {
			print("err:", err)
			return PairData{}, nil
		}
		print("out:", out)
		switch {
		case idx == 0:
			pairData.TotalLiquidity = toFloat(out)
		case idx == 1:
			pairData.Price = toFloat(out)
		default:
			print("idx not needed")
		}
	}
	print("\npairData:")
	dump(pairData)
	return pairData, nil
}

func regexp2FindInBtw(inputStr string, pattern string) (string, error) {
	var strOut string
	re := regexp2.MustCompile(pattern, 0)
	isMatch, err := re.MatchString(inputStr)
	if re.MatchTimeout*time.Second > 3 {
		return strOut, errors.New("err@ re.MatchTimeout")
	}
	if err != nil {
		return strOut, errors.New("err@ re.MatchString")
	}
	fmt.Println("isMatch:", isMatch)
	if isMatch {
		if m, err := re.FindStringMatch(inputStr); m != nil {
			if err != nil {
				return strOut, errors.New("err@ re.FindStringMatch")
			}
			// the whole match is always group 0
			strOut = m.String()
			fmt.Printf("Group 0: %v\n", "=="+strOut+"==")
			return strOut, nil
			//return removeBothEnds(strOut), nil
			// you can get all the groups too
			// gps := m.Groups()

			// // a group can be captured multiple times, so each cap is separately addressable
			// fmt.Println("Group 1, first capture", gps[1].Captures[0].String())
			// fmt.Println("Group 1, second capture", gps[1].Captures[1].String())
		}
	}
	return strOut, nil
}

func removeBothEnds(strIn string) string {
	if len(strIn) < 4 {
		return ""
	}
	s1 := strIn[2:]
	last := len(s1) - 1
	return s1[:last]
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
