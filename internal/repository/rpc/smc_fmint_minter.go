// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rpc

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

// DefiFMintMinterABI is the input ABI used to generate the binding from.
const DefiFMintMinterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEBT_EXCEEDED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEPOSIT_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_ALLOWANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_COLLATERAL_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_MINTING_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NOT_AUTHORIZED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_DEPLETED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_EARLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_NONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARD_CLAIM_REJECTED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_ZERO_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIFantomMintAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"collateralCanDecrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralLowestDebtRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralRatioDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"debtCanIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFeeDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustRepay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardCanClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEligibilityRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardIsEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"debtValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"collateralValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebtPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"checkCollateralCanDecrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"checkDebtCanIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPriceDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"tokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DefiFMintMinter is an auto generated Go binding around an Ethereum contract.
type DefiFMintMinter struct {
	DefiFMintMinterCaller     // Read-only binding to the contract
	DefiFMintMinterTransactor // Write-only binding to the contract
	DefiFMintMinterFilterer   // Log filterer for contract events
}

// DefiFMintMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiFMintMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiFMintMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFMintMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiFMintMinterSession struct {
	Contract     *DefiFMintMinter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiFMintMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiFMintMinterCallerSession struct {
	Contract *DefiFMintMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DefiFMintMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiFMintMinterTransactorSession struct {
	Contract     *DefiFMintMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DefiFMintMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiFMintMinterRaw struct {
	Contract *DefiFMintMinter // Generic contract binding to access the raw methods on
}

// DefiFMintMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiFMintMinterCallerRaw struct {
	Contract *DefiFMintMinterCaller // Generic read-only contract binding to access the raw methods on
}

// DefiFMintMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiFMintMinterTransactorRaw struct {
	Contract *DefiFMintMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiFMintMinter creates a new instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinter(address common.Address, backend bind.ContractBackend) (*DefiFMintMinter, error) {
	contract, err := bindDefiFMintMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinter{DefiFMintMinterCaller: DefiFMintMinterCaller{contract: contract}, DefiFMintMinterTransactor: DefiFMintMinterTransactor{contract: contract}, DefiFMintMinterFilterer: DefiFMintMinterFilterer{contract: contract}}, nil
}

// NewDefiFMintMinterCaller creates a new read-only instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterCaller(address common.Address, caller bind.ContractCaller) (*DefiFMintMinterCaller, error) {
	contract, err := bindDefiFMintMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterCaller{contract: contract}, nil
}

// NewDefiFMintMinterTransactor creates a new write-only instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiFMintMinterTransactor, error) {
	contract, err := bindDefiFMintMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterTransactor{contract: contract}, nil
}

// NewDefiFMintMinterFilterer creates a new log filterer instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFMintMinterFilterer, error) {
	contract, err := bindDefiFMintMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterFilterer{contract: contract}, nil
}

// bindDefiFMintMinter binds a generic wrapper to an already deployed contract.
func bindDefiFMintMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiFMintMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintMinter *DefiFMintMinterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiFMintMinter.Contract.DefiFMintMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintMinter *DefiFMintMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.DefiFMintMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintMinter *DefiFMintMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.DefiFMintMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintMinter *DefiFMintMinterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiFMintMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintMinter *DefiFMintMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintMinter *DefiFMintMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.contract.Transact(opts, method, params...)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRDEBTEXCEEDED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_DEBT_EXCEEDED")
	return *ret0, err
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEBTEXCEEDED(&_DefiFMintMinter.CallOpts)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEBTEXCEEDED(&_DefiFMintMinter.CallOpts)
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRDEPOSITPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_DEPOSIT_PROHIBITED")
	return *ret0, err
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEPOSITPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEPOSITPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWALLOWANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_ALLOWANCE")
	return *ret0, err
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWALLOWANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWALLOWANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_AMOUNT")
	return *ret0, err
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_BALANCE")
	return *ret0, err
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWBALANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWBALANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWBALANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWBALANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWCOLLATERALRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_COLLATERAL_RATIO")
	return *ret0, err
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWCOLLATERALRATIO(&_DefiFMintMinter.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWCOLLATERALRATIO(&_DefiFMintMinter.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRMINTINGPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_MINTING_PROHIBITED")
	return *ret0, err
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRMINTINGPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRMINTINGPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOTAUTHORIZED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NOT_AUTHORIZED")
	return *ret0, err
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOTAUTHORIZED(&_DefiFMintMinter.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOTAUTHORIZED(&_DefiFMintMinter.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOCOLLATERAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_COLLATERAL")
	return *ret0, err
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOCOLLATERAL(&_DefiFMintMinter.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOCOLLATERAL(&_DefiFMintMinter.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_ERROR")
	return *ret0, err
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOERROR() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOERROR(&_DefiFMintMinter.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOERROR() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOERROR(&_DefiFMintMinter.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_REWARD")
	return *ret0, err
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOREWARD() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOREWARD(&_DefiFMintMinter.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOREWARD() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOREWARD(&_DefiFMintMinter.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_VALUE")
	return *ret0, err
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOVALUE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOVALUE(&_DefiFMintMinter.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOVALUE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOVALUE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSDEPLETED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_DEPLETED")
	return *ret0, err
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSDEPLETED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSDEPLETED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSEARLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_EARLY")
	return *ret0, err
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSEARLY(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSEARLY(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSNONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_NONE")
	return *ret0, err
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSNONE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSNONE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSNONE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSNONE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDCLAIMREJECTED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARD_CLAIM_REJECTED")
	return *ret0, err
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDCLAIMREJECTED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDCLAIMREJECTED(&_DefiFMintMinter.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRZEROAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_ZERO_AMOUNT")
	return *ret0, err
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRZEROAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRZEROAMOUNT(&_DefiFMintMinter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "addressProvider")
	return *ret0, err
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) AddressProvider() (common.Address, error) {
	return _DefiFMintMinter.Contract.AddressProvider(&_DefiFMintMinter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) AddressProvider() (common.Address, error) {
	return _DefiFMintMinter.Contract.AddressProvider(&_DefiFMintMinter.CallOpts)
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CanDeposit(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "canDeposit", _token)
	return *ret0, err
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanDeposit(&_DefiFMintMinter.CallOpts, _token)
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanDeposit(&_DefiFMintMinter.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CanMint(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "canMint", _token)
	return *ret0, err
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanMint(&_DefiFMintMinter.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanMint(&_DefiFMintMinter.CallOpts, _token)
}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CheckCollateralCanDecrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "checkCollateralCanDecrease", _account, _token, _amount)
	return *ret0, err
}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CheckCollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckCollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CheckCollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckCollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CheckDebtCanIncrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "checkDebtCanIncrease", _account, _token, _amount)
	return *ret0, err
}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CheckDebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckDebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CheckDebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckDebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralCanDecrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralCanDecrease", _account, _token, _amount)
	return *ret0, err
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralLowestDebtRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralLowestDebtRatio4dec")
	return *ret0, err
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralRatioDecimalsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralRatioDecimalsCorrection")
	return *ret0, err
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralRatioDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralRatioDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralRatioDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralRatioDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralValueOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralValueOf", _account)
	return *ret0, err
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtCanIncrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtCanIncrease", _account, _token, _amount)
	return *ret0, err
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.DebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.DebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtValueOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtValueOf", _account)
	return *ret0, err
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "fMintFee")
	return *ret0, err
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintFee() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee(&_DefiFMintMinter.CallOpts)
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintFee() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee(&_DefiFMintMinter.CallOpts)
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFeeDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "fMintFeeDigitsCorrection")
	return *ret0, err
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintFeeDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFeeDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintFeeDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFeeDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FeePool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "feePool")
	return *ret0, err
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FeePool() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FeePool() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetCollateralPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getCollateralPool")
	return *ret0, err
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetCollateralPool(&_DefiFMintMinter.CallOpts)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetCollateralPool(&_DefiFMintMinter.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetDebtPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getDebtPool")
	return *ret0, err
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetDebtPool(&_DefiFMintMinter.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetDebtPool(&_DefiFMintMinter.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getFeeToken")
	return *ret0, err
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) GetFeeToken() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetFeeToken(&_DefiFMintMinter.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetFeeToken() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetFeeToken(&_DefiFMintMinter.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getPrice", _token)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0x173a58fe.
//
// Solidity: function getPriceDigitsCorrection(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetPriceDigitsCorrection(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getPriceDigitsCorrection", _token)
	return *ret0, err
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0x173a58fe.
//
// Solidity: function getPriceDigitsCorrection(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetPriceDigitsCorrection(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPriceDigitsCorrection(&_DefiFMintMinter.CallOpts, _token)
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0x173a58fe.
//
// Solidity: function getPriceDigitsCorrection(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetPriceDigitsCorrection(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPriceDigitsCorrection(&_DefiFMintMinter.CallOpts, _token)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardCanClaim(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardCanClaim", _account)
	return *ret0, err
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardCanClaim(&_DefiFMintMinter.CallOpts, _account)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardCanClaim(&_DefiFMintMinter.CallOpts, _account)
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEligibilityRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardEligibilityRatio4dec")
	return *ret0, err
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardIsEligible(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardIsEligible", _account)
	return *ret0, err
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardIsEligible(&_DefiFMintMinter.CallOpts, _account)
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardIsEligible(&_DefiFMintMinter.CallOpts, _account)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) TokenValue(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "tokenValue", _token, _amount)
	return *ret0, err
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.TokenValue(&_DefiFMintMinter.CallOpts, _token, _amount)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.TokenValue(&_DefiFMintMinter.CallOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Deposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Deposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Mint(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mint", _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Mint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Mint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustDeposit", _token, _amount)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustDeposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustDeposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustMint(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustMint", _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustMint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustMint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustRepay(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustRepay", _token, _amount)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustRepay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustRepay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustWithdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustWithdraw", _token, _amount)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Repay(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "repay", _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Repay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Repay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) RewardUpdate(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "rewardUpdate", _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardUpdate(&_DefiFMintMinter.TransactOpts, _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardUpdate(&_DefiFMintMinter.TransactOpts, _account)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Withdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Withdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// DefiFMintMinterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the DefiFMintMinter contract.
type DefiFMintMinterDepositedIterator struct {
	Event *DefiFMintMinterDeposited // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterDeposited)
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
		it.Event = new(DefiFMintMinterDeposited)
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
func (it *DefiFMintMinterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterDeposited represents a Deposited event raised by the DefiFMintMinter contract.
type DefiFMintMinterDeposited struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterDeposited(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Deposited", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterDepositedIterator{contract: _DefiFMintMinter.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterDeposited, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Deposited", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterDeposited)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseDeposited(log types.Log) (*DefiFMintMinterDeposited, error) {
	event := new(DefiFMintMinterDeposited)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the DefiFMintMinter contract.
type DefiFMintMinterMintedIterator struct {
	Event *DefiFMintMinterMinted // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterMinted)
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
		it.Event = new(DefiFMintMinterMinted)
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
func (it *DefiFMintMinterMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterMinted represents a Minted event raised by the DefiFMintMinter contract.
type DefiFMintMinterMinted struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterMinted(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterMintedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Minted", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterMintedIterator{contract: _DefiFMintMinter.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterMinted, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Minted", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterMinted)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseMinted(log types.Log) (*DefiFMintMinterMinted, error) {
	event := new(DefiFMintMinterMinted)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the DefiFMintMinter contract.
type DefiFMintMinterRepaidIterator struct {
	Event *DefiFMintMinterRepaid // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterRepaid)
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
		it.Event = new(DefiFMintMinterRepaid)
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
func (it *DefiFMintMinterRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterRepaid represents a Repaid event raised by the DefiFMintMinter contract.
type DefiFMintMinterRepaid struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterRepaid(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterRepaidIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Repaid", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterRepaidIterator{contract: _DefiFMintMinter.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterRepaid, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Repaid", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterRepaid)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Repaid", log); err != nil {
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

// ParseRepaid is a log parse operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseRepaid(log types.Log) (*DefiFMintMinterRepaid, error) {
	event := new(DefiFMintMinterRepaid)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the DefiFMintMinter contract.
type DefiFMintMinterWithdrawnIterator struct {
	Event *DefiFMintMinterWithdrawn // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterWithdrawn)
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
		it.Event = new(DefiFMintMinterWithdrawn)
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
func (it *DefiFMintMinterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterWithdrawn represents a Withdrawn event raised by the DefiFMintMinter contract.
type DefiFMintMinterWithdrawn struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Withdrawn", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterWithdrawnIterator{contract: _DefiFMintMinter.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterWithdrawn, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Withdrawn", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterWithdrawn)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseWithdrawn(log types.Log) (*DefiFMintMinterWithdrawn, error) {
	event := new(DefiFMintMinterWithdrawn)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}