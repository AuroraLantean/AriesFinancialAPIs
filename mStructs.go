package main

import "database/sql"

// APYs ... based on last week's performance
type APYs struct {
	ID         string `json:"id"`
	SrcAndPP   string `json:"srcAndPP"`
	WETH       string `json:"weth"`
	AFI        string `json:"afi"`
	YFI        string `json:"yfi"`
	CRV3       string `json:"crv3"`
	CRVY       string `json:"crvy"`
	CRVBUSD    string `json:"crvbusd"`
	CRVSBTC    string `json:"crvsbtc"`
	DAI        string `json:"dai"`
	TrueUSD    string `json:"trueusd"`
	USDC       string `json:"usdc"`
	Gemini     string `json:"gemini"`
	TetherUSD  string `json:"tetherUSD"`
	EthDai     string `json:"eth_dai"`
	EthUSDC    string `json:"eth_usdc"`
	EthUSDT    string `json:"eth_usdt"`
	EthWBTC    string `json:"eth_wbtc"`
	CrvRENWBTC string `json:"crv_renwbtc"`
	WBTC       string `json:"wbtc"`
	RENBTC     string `json:"renbtc"`
	WbtcTBTC   string `json:"wbtc_tbtc"`
	FARM       string `json:"farm"`
} /*curvefi3pool, curvefiy, curvefibusd, curvefisbtc
 */

// Wallet           []Coin `json:"wallet"`
// AccountID        int    `json:"account_id"`

// NullString ...
type NullString struct {
	sql.NullString
}

// OutputLambda ... API Response details
/*Code    ... response code
Mesg    ... response message
Data    ... response data
*/
type OutputLambda struct {
	Code string      `json:"code"`
	Mesg string      `json:"mesg"`
	Data interface{} `json:"data"`
}

// VaultAPY ...
type VaultAPY struct {
	ApyOneMonthSample  float64 `json:"apyOneMonthSample"`
	Symbol             string  `json:"symbol"`
	Timestamp          int64   `json:"timestamp"`
	ApyOneWeekSample   float64 `json:"apyOneWeekSample"`
	ApyInceptionSample float64 `json:"apyInceptionSample"`
	Address            string  `json:"address"`
	Name               string  `json:"name"`
	VaultSymbol        string  `json:"vaultSymbol"`
	Boost              BoostT  `json:"boost"`
	ApyOneDaySample    float64 `json:"apyOneDaySample"`
	ApyThreeDaySample  float64 `json:"apyThreeDaySample"`
	TokenAddress       string  `json:"tokenAddress"`
	Description        string  `json:"description"`
	ApyLoanscan        float64 `json:"apyLoanscan"`
	PoolApy            float64 `json:"poolApy"`
	VaultAddress       string  `json:"vaultAddress"`
}

/*
[
  {
    "apyOneMonthSample": 6.2998062451804016,
    "symbol": "yCRV",
    "timestamp": 1604626617893,
    "apyOneWeekSample": 10.233916938611037,
    "apyInceptionSample": 43.79820272711738,
    "address": "0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c",
    "name": "curve.fi/y LP",
    "vaultSymbol": "yUSD",
    "boost": {},
    "apyOneDaySample": 10.151891381361345,
    "apyThreeDaySample": 8.761802478742508,
    "tokenAddress": "0xdf5e0e81dff6faf3a7e52ba697820c5e32d806a8",
    "description": "yDAI/yUSDC/yUSDT/yTUSD",
    "apyLoanscan": 15.589686203503739,
    "poolApy": 4.936633183461183,
    "vaultAddress": "0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c"
  },
]
*/

// BoostT ...
type BoostT struct {
	GaugeBalance   float64 `json:"gaugeBalance"`
	WorkingBalance float64 `json:"workingBalance"`
	MaxBoost       float64 `json:"maxBoost"`
	GaugeTotal     float64 `json:"gaugeTotal"`
	VecrvTotal     float64 `json:"vecrvTotal"`
	MinVecrv       float64 `json:"minVecrv"`
	Boost          float64 `json:"boost"`
	VecrvBalance   float64 `json:"vecrvBalance"`
	WorkingTotal   float64 `json:"workingTotal"`
}

/*
  "boost": {
    "gaugeBalance": 70861853.99377115,
    "workingBalance": 29859189.876523312,
    "maxBoost": 1.5511842607592454,
    "gaugeTotal": 139632969.79810524,
    "vecrvTotal": 37634333.5469901,
    "minVecrv": 19098918.06222891,
    "boost": 1.053429602588015,
    "vecrvBalance": 681529.4812465438,
    "workingTotal": 71024125.27361171
  },
*/

// RespNameY ... account and login details
/*Token       ... JWT Token
CreatedAt   ... time this account was created
*/
type RespNameY struct {
	Token string `json:"token"`
	Level int    `json:"level"`
} //	Wallet      []CoinOut `json:"wallet"`

// InputLambda of all Lambda functions ...
type InputLambda struct {
	Body       ReqBody `json:"reqBody"`
	DataName   string  `json:"dataName"`
	APYboDay   APYs    `json:"apyboday"`
	APYboWeek  APYs    `json:"apyboweek"`
	APYboMonth APYs    `json:"apybomonth"`
} // rows of diff source, columns of diff farms
//srcAndPP: source plus performance period
//AuthHeader  string  `json:"Authorization"`
//URLvalues   url.Values  `json:"urlValues"`
//ContentType string  `json:"contentType"`

// ReqBody of all APIs ...
type ReqBody struct {
	SourceURL  string `json:"sourceURL"`
	PerfPeriod string `json:"perfPeriod"`
}

// HTPMPattern string `json:"htmlPattern"`
// RegexpStr   string `json:"regexpStr"`

// TableNameX ... account verification table details
/*ID          ... primary key
MailAddress ... Email address
*/
type TableNameX struct {
	ID          int    `json:"id"`
	MailAddress string `json:"mail_address"`
}

// RoutineInputs ...
type RoutineInputs struct {
	RoutineName string `json:"routineName"`
	Address     string `json:"requestAddr"`
	Method      string `json:"method"`
	Timeout     int    `json:"timeout"`
}

// RoutineOut ...
type RoutineOut struct {
	Code        string `json:"Code"`
	Mesg        string `json:"mesg"`
	RespRoutine string `json:"respRoutine"`
}