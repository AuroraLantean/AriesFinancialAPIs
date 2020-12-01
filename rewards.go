// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RewardsABI is the input ABI used to generate the binding from.
const RewardsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getData1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getData2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnlyRewardDistribution\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"setLpToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Rewards *RewardsCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Rewards *RewardsSession) DURATION() (*big.Int, error) {
	return _Rewards.Contract.DURATION(&_Rewards.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Rewards *RewardsCallerSession) DURATION() (*big.Int, error) {
	return _Rewards.Contract.DURATION(&_Rewards.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Rewards *RewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Rewards *RewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Rewards.Contract.BalanceOf(&_Rewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Rewards *RewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Rewards.Contract.BalanceOf(&_Rewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Rewards *RewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Rewards *RewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _Rewards.Contract.Earned(&_Rewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Rewards *RewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Rewards.Contract.Earned(&_Rewards.CallOpts, account)
}

// Erc20Token is a free data retrieval call binding the contract method 0x8a13eea7.
//
// Solidity: function erc20Token() view returns(address)
func (_Rewards *RewardsCaller) Erc20Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "erc20Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Token is a free data retrieval call binding the contract method 0x8a13eea7.
//
// Solidity: function erc20Token() view returns(address)
func (_Rewards *RewardsSession) Erc20Token() (common.Address, error) {
	return _Rewards.Contract.Erc20Token(&_Rewards.CallOpts)
}

// Erc20Token is a free data retrieval call binding the contract method 0x8a13eea7.
//
// Solidity: function erc20Token() view returns(address)
func (_Rewards *RewardsCallerSession) Erc20Token() (common.Address, error) {
	return _Rewards.Contract.Erc20Token(&_Rewards.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_Rewards *RewardsCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_Rewards *RewardsSession) GetBlockTimestamp() (*big.Int, error) {
	return _Rewards.Contract.GetBlockTimestamp(&_Rewards.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_Rewards *RewardsCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _Rewards.Contract.GetBlockTimestamp(&_Rewards.CallOpts)
}

// GetData1 is a free data retrieval call binding the contract method 0x9944cc71.
//
// Solidity: function getData1() view returns(uint256, uint256, uint256, uint256)
func (_Rewards *RewardsCaller) GetData1(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getData1")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetData1 is a free data retrieval call binding the contract method 0x9944cc71.
//
// Solidity: function getData1() view returns(uint256, uint256, uint256, uint256)
func (_Rewards *RewardsSession) GetData1() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Rewards.Contract.GetData1(&_Rewards.CallOpts)
}

// GetData1 is a free data retrieval call binding the contract method 0x9944cc71.
//
// Solidity: function getData1() view returns(uint256, uint256, uint256, uint256)
func (_Rewards *RewardsCallerSession) GetData1() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Rewards.Contract.GetData1(&_Rewards.CallOpts)
}

// GetData2 is a free data retrieval call binding the contract method 0xa898fd70.
//
// Solidity: function getData2() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Rewards *RewardsCaller) GetData2(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getData2")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetData2 is a free data retrieval call binding the contract method 0xa898fd70.
//
// Solidity: function getData2() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Rewards *RewardsSession) GetData2() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Rewards.Contract.GetData2(&_Rewards.CallOpts)
}

// GetData2 is a free data retrieval call binding the contract method 0xa898fd70.
//
// Solidity: function getData2() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Rewards *RewardsCallerSession) GetData2() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Rewards.Contract.GetData2(&_Rewards.CallOpts)
}

// IsOnlyRewardDistribution is a free data retrieval call binding the contract method 0x3cbc5d28.
//
// Solidity: function isOnlyRewardDistribution() view returns(bool)
func (_Rewards *RewardsCaller) IsOnlyRewardDistribution(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "isOnlyRewardDistribution")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnlyRewardDistribution is a free data retrieval call binding the contract method 0x3cbc5d28.
//
// Solidity: function isOnlyRewardDistribution() view returns(bool)
func (_Rewards *RewardsSession) IsOnlyRewardDistribution() (bool, error) {
	return _Rewards.Contract.IsOnlyRewardDistribution(&_Rewards.CallOpts)
}

// IsOnlyRewardDistribution is a free data retrieval call binding the contract method 0x3cbc5d28.
//
// Solidity: function isOnlyRewardDistribution() view returns(bool)
func (_Rewards *RewardsCallerSession) IsOnlyRewardDistribution() (bool, error) {
	return _Rewards.Contract.IsOnlyRewardDistribution(&_Rewards.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Rewards *RewardsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Rewards *RewardsSession) IsOwner() (bool, error) {
	return _Rewards.Contract.IsOwner(&_Rewards.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Rewards *RewardsCallerSession) IsOwner() (bool, error) {
	return _Rewards.Contract.IsOwner(&_Rewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Rewards *RewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Rewards *RewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Rewards.Contract.LastTimeRewardApplicable(&_Rewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Rewards *RewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Rewards.Contract.LastTimeRewardApplicable(&_Rewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Rewards *RewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Rewards *RewardsSession) LastUpdateTime() (*big.Int, error) {
	return _Rewards.Contract.LastUpdateTime(&_Rewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Rewards *RewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Rewards.Contract.LastUpdateTime(&_Rewards.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Rewards *RewardsCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Rewards *RewardsSession) LpToken() (common.Address, error) {
	return _Rewards.Contract.LpToken(&_Rewards.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Rewards *RewardsCallerSession) LpToken() (common.Address, error) {
	return _Rewards.Contract.LpToken(&_Rewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCallerSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Rewards *RewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Rewards *RewardsSession) PeriodFinish() (*big.Int, error) {
	return _Rewards.Contract.PeriodFinish(&_Rewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Rewards *RewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _Rewards.Contract.PeriodFinish(&_Rewards.CallOpts)
}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_Rewards *RewardsCaller) RewardDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_Rewards *RewardsSession) RewardDistribution() (common.Address, error) {
	return _Rewards.Contract.RewardDistribution(&_Rewards.CallOpts)
}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_Rewards *RewardsCallerSession) RewardDistribution() (common.Address, error) {
	return _Rewards.Contract.RewardDistribution(&_Rewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Rewards *RewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Rewards *RewardsSession) RewardPerToken() (*big.Int, error) {
	return _Rewards.Contract.RewardPerToken(&_Rewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Rewards *RewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _Rewards.Contract.RewardPerToken(&_Rewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewards *RewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewards *RewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _Rewards.Contract.RewardPerTokenStored(&_Rewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewards *RewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Rewards.Contract.RewardPerTokenStored(&_Rewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Rewards *RewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Rewards *RewardsSession) RewardRate() (*big.Int, error) {
	return _Rewards.Contract.RewardRate(&_Rewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Rewards *RewardsCallerSession) RewardRate() (*big.Int, error) {
	return _Rewards.Contract.RewardRate(&_Rewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Rewards(&_Rewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Rewards(&_Rewards.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsSession) TotalSupply() (*big.Int, error) {
	return _Rewards.Contract.TotalSupply(&_Rewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _Rewards.Contract.TotalSupply(&_Rewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Rewards *RewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Rewards *RewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.UserRewardPerTokenPaid(&_Rewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Rewards *RewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.UserRewardPerTokenPaid(&_Rewards.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Rewards *RewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Rewards *RewardsSession) Exit() (*types.Transaction, error) {
	return _Rewards.Contract.Exit(&_Rewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Rewards *RewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _Rewards.Contract.Exit(&_Rewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Rewards *RewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Rewards *RewardsSession) GetReward() (*types.Transaction, error) {
	return _Rewards.Contract.GetReward(&_Rewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Rewards *RewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _Rewards.Contract.GetReward(&_Rewards.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Rewards *RewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Rewards *RewardsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.NotifyRewardAmount(&_Rewards.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Rewards *RewardsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.NotifyRewardAmount(&_Rewards.TransactOpts, reward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_Rewards *RewardsTransactor) SetLpToken(opts *bind.TransactOpts, _lpToken common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setLpToken", _lpToken)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_Rewards *RewardsSession) SetLpToken(_lpToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetLpToken(&_Rewards.TransactOpts, _lpToken)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_Rewards *RewardsTransactorSession) SetLpToken(_lpToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetLpToken(&_Rewards.TransactOpts, _lpToken)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Rewards *RewardsTransactor) SetRewardDistribution(opts *bind.TransactOpts, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setRewardDistribution", _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Rewards *RewardsSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetRewardDistribution(&_Rewards.TransactOpts, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Rewards *RewardsTransactorSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetRewardDistribution(&_Rewards.TransactOpts, _rewardDistribution)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Rewards *RewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Rewards *RewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Stake(&_Rewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Rewards *RewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Stake(&_Rewards.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Rewards *RewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Rewards *RewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Withdraw(&_Rewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Rewards *RewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Withdraw(&_Rewards.TransactOpts, amount)
}

// RewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewards contract.
type RewardsOwnershipTransferredIterator struct {
	Event *RewardsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsOwnershipTransferred represents a OwnershipTransferred event raised by the Rewards contract.
type RewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsOwnershipTransferredIterator{contract: _Rewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsOwnershipTransferred)
				if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsOwnershipTransferred, error) {
	event := new(RewardsOwnershipTransferred)
	if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Rewards contract.
type RewardsRewardAddedIterator struct {
	Event *RewardsRewardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsRewardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardAdded represents a RewardAdded event raised by the Rewards contract.
type RewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Rewards *RewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*RewardsRewardAddedIterator, error) {

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &RewardsRewardAddedIterator{contract: _Rewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Rewards *RewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *RewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardAdded)
				if err := _Rewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Rewards *RewardsFilterer) ParseRewardAdded(log types.Log) (*RewardsRewardAdded, error) {
	event := new(RewardsRewardAdded)
	if err := _Rewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Rewards contract.
type RewardsRewardPaidIterator struct {
	Event *RewardsRewardPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsRewardPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardPaid represents a RewardPaid event raised by the Rewards contract.
type RewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Rewards *RewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*RewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardPaidIterator{contract: _Rewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Rewards *RewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *RewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardPaid)
				if err := _Rewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Rewards *RewardsFilterer) ParseRewardPaid(log types.Log) (*RewardsRewardPaid, error) {
	event := new(RewardsRewardPaid)
	if err := _Rewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Rewards contract.
type RewardsStakedIterator struct {
	Event *RewardsStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsStaked represents a Staked event raised by the Rewards contract.
type RewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*RewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsStakedIterator{contract: _Rewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *RewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsStaked)
				if err := _Rewards.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) ParseStaked(log types.Log) (*RewardsStaked, error) {
	event := new(RewardsStaked)
	if err := _Rewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Rewards contract.
type RewardsWithdrawnIterator struct {
	Event *RewardsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsWithdrawn represents a Withdrawn event raised by the Rewards contract.
type RewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*RewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsWithdrawnIterator{contract: _Rewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsWithdrawn)
				if err := _Rewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) ParseWithdrawn(log types.Log) (*RewardsWithdrawn, error) {
	event := new(RewardsWithdrawn)
	if err := _Rewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
