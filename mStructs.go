package main

import "database/sql"

// RewardsPool ...
type RewardsPool struct {
	ID                   string `yaml:"id"`
	Name                 string `yaml:"name"`
	Network              string `yaml:"network"`
	LpTokenCtrt          string `yaml:"lpToken"`
	RwTokenCtrt          string `yaml:"rwToken"`
	RewardsCtrt          string `yaml:"rewardsCtrt"`
	LpTokenPriceSource   string `yaml:"lpTokenPriceSource"`
	TotalLiquiditySource string `yaml:"totalLiquiditySource"`
	LoadingTime          int    `yaml:"loadingTime"`
	RwTokenPriceSorce    string `yaml:"rwTokenPriceSource"`
}

// Config ... configuration file
type Config struct {
	EthereumNetwork struct {
		Name string `yaml:"name"`
	} `yaml:"EthereumNetwork"`

	RewardsPools []struct {
		ID                   string `yaml:"id"`
		Name                 string `yaml:"name"`
		Network              string `yaml:"network"`
		LpTokenCtrt          string `yaml:"lpToken"`
		RwTokenCtrt          string `yaml:"rwToken"`
		RewardsCtrt          string `yaml:"rewardsCtrt"`
		LpTokenPriceSource   string `yaml:"lpTokenPriceSource"`
		TotalLiquiditySource string `yaml:"totalLiquiditySource"`
		LoadingTime          int    `yaml:"loadingTime"`
		RwTokenPriceSorce    string `yaml:"rwTokenPriceSource"`
	} `yaml:"RewardsPools"`
}

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
	UpdatedAt  string `json:"updatedAt"`
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

// ReqBody of all APIs ... ALL inputs are string!
type ReqBody struct {
	SourceURL     string   `json:"sourceURL"`
	PerfPeriod    string   `json:"perfPeriod"`
	UserID        string   `json:"userID"`
	VaultID       string   `json:"vaultID"`
	RewardID      string   `json:"rewardID"`
	EthereumAddr  string   `json:"ethereumAddr"`
	Reward        string   `json:"reward"`
	UpdatedAt     string   `json:"updatedAt"`
	UplineID      string   `json:"uplineID"`
	DownlineIDs   string   `json:"downlineIDs"`
	EthereumAddrs []string `json:"ethereumAddrs"`
	Offset        string   `json:"offset"`
	Amount        string   `json:"amount"`
	SourceName    string   `json:"sourceName"`
	Token0        string   `json:"token0"`
	Token1        string   `json:"token1"`
	RewardsPool   string   `json:"rewardsPool"`
}

// HTPMPattern string `json:"htmlPattern"`
// RegexpStr   string `json:"regexpStr"`

// TableNameX ...
type TableNameX struct {
	ID           int    `json:"id"`
	EthereumAddr string `json:"ethereumAddr"`
}

// RespUser ... account and login details
/*Token       ... JWT Token
CreatedAt   ... time this account was created
*/
type RespUser struct {
	ID            string `json:"id"`
	EthereumAddr  string `json:"ethAddress"`
	Reward        string `json:"reward"`
	UpdatedAt     string `json:"updatedAt"`
	UplineID      string `json:"uplineID"`
	DownlineIDs   string `json:"downlineIDs"`
	RiskCheckedAt string `json:"riskCheckedAt"`
} //	Wallet      []CoinOut `json:"wallet"`

// Reward ...
/*ID          ... primary key
Reward        ... float
*/
type Reward struct {
	ID        string `json:"id"`
	UserID    string `json:"userID"`
	VaultID   string `json:"vaultID"`
	Reward    string `json:"reward"`
	UpdatedAt string `json:"updatedAt"`
}

// RespRewardC ... account and login details
type RespRewardC struct {
	NewRowID string `json:"newRowID"`
} //	Wallet      []CoinOut `json:"wallet"`

// RespRewardR ...
/*ID          ... primary key
Reward        ... float
*/
type RespRewardR struct {
	ID        string `json:"id"`
	UserID    string `json:"userID"`
	VaultID   string `json:"vaultID"`
	Reward    string `json:"reward"`
	UpdatedAt string `json:"updatedAt"`
}

// VaultToEthAddr ...
/*
VaultID          ... primary key
EthereumAddr  ... string
*/
type VaultToEthAddr struct {
	VaultID      string `json:"vaultID"`
	EthereumAddr string `json:"ethereumAddr"`
}

// RespVaultEthAddrR ...
/*ID          ... primary key
Reward        ... float
*/
type RespVaultEthAddrR struct {
	VaultID      string `json:"vaultID"`
	EthereumAddr string `json:"ethereumAddr"`
}

// RoutineInputs ...
type RoutineInputs struct {
	RoutineName string `json:"routineName"`
	Address     string `json:"requestAddr"`
	Body        string `json:"body"`
	Method      string `json:"method"`
	Timeout     int    `json:"timeout"`
}

// RoutineOut ...
type RoutineOut struct {
	Code        string      `json:"Code"`
	Mesg        string      `json:"mesg"`
	RespRoutine interface{} `json:"respRoutine"`
}

// AnChainOut ...
type AnChainOut struct {
	Data   map[string]AnChainData `json:"data"`
	ErrMsg string                 `json:"err_msg"`
	Status int                    `json:"status"`
}

// GraphqlOut ...
type GraphqlOut struct {
	Data   UniSwapData  `json:"data"`
	Errors []GraphqlErr `json:"errors"`
}

// GraphqlErr ...
type GraphqlErr struct {
	Locations []LineColumn `json:"locations"`
	Message   string       `json:"message"`
}

// LineColumn ...
type LineColumn struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// UniSwapToken ...
type UniSwapToken struct {
	ID                 string `json:"id"`
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Decimals           string `json:"decimals"`
	TotalSupply        string `json:"totalSupply"`
	TradeVolume        string `json:"tradeVolume"`
	TradeVolumeUSD     string `json:"tradeVolumeUSD"`
	UntrackedVolumeUSD string `json:"untrackedVolumeUSD"`
	TxCount            string `json:"txCount"`
	TotalLiquidity     string `json:"totalLiquidity"`
	DerivedETH         string `json:"derivedETH"`
	MostLiquidPairs    string `json:"mostLiquidPairs"`
}

// UniSwapPair ...
type UniSwapPair struct {
	UniSwapToken0 UniSwapToken `json:"token0"`
	Token0Price   string       `json:"token0Price"`
	UniSwapToken1 UniSwapToken `json:"token1"`
	Token1Price   string       `json:"token1Price"`
}

// PairData ...
type PairData struct {
	Price            float64 `json:"price"`
	TotalLiquidity   float64 `json:"totalLiquidity"`
	TotalValueLocked float64 `json:"totalValueLocked"`
	WeeklyROI        float64 `json:"weeklyROI"`
	APY              float64 `json:"apy"`
}

// UniSwapData ...
type UniSwapData struct {
	UniSwapPairData UniSwapPair `json:"pair"`
}

// AnChainData ...
type AnChainData struct {
	IsAddrValid bool     `json:"is_address_valid"`
	Risk        Risk     `json:"risk"`
	Self        Self     `json:"self"`
	Activity    Activity `json:"activity"`
}

// Activity ...
type Activity struct {
	SuspicousActivity         []SuspiciousActivity `json:"suspicious_activity"`
	SuspiciousActivityDeclare string               `json:"suspicious_activity_declare"`
	VerdictTime               int                  `json:"verdict_time"`
}

// SuspiciousActivity ...
type SuspiciousActivity struct {
	AggrType    string   `json:"aggr_type"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Entity      string   `json:"entity"`
	TxnCnt      int      `json:"txn_cnt"`
	TxnDirect   int      `json:"txn_direct"`
	TxnHashes   []string `json:"txn_hashes"`
	TxnVol      float64  `json:"txn_vol"`
}

// Self ...
type Self struct {
	Category []string `json:"category"`
	Detail   []string `json:"detail"`
}

// Risk ...
type Risk struct {
	Level       int `json:"level"`
	Score       int `json:"score"`
	VerdictTime int `json:"verdict_time"`
}

// Error ...
type Error struct {
	Message string `json:"message"`
	MesgRaw string `json:"mesgraw"`
}
