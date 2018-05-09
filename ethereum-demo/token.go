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

// MyAdvancedTokenABI is the input ABI used to generate the binding from.
const MyAdvancedTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newSellPrice\",\"type\":\"uint256\"},{\"name\":\"newBuyPrice\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"centralMinter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"FrozenFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// MyAdvancedTokenBin is the compiled bytecode used for deploying new contracts.
const MyAdvancedTokenBin = `0x60606040526003805460ff19166012179055341561001c57600080fd5b604051610e6d380380610e6d833981016040528080519190602001805182019190602001805182019190602001805160008054600160a060020a03191633600160a060020a0316908117825560035460ff16600a0a8802600481905590825260056020526040909120559150849050838360018280516100a09291602001906100f6565b5060028180516100b49291602001906100f6565b50505050600160a060020a038116156100e35760008054600160a060020a031916600160a060020a0383161790555b5050600260075550506004600855610191565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013757805160ff1916838001178555610164565b82800160010185558215610164579182015b82811115610164578251825591602001919060010190610149565b50610170929150610174565b5090565b61018e91905b80821115610170576000815560010161017a565b90565b610ccd806101a06000396000f3006060604052600436106101275763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305fefda7811461012c57806306fdde0314610147578063095ea7b3146101d157806318160ddd1461020757806323b872dd1461022c578063313ce5671461025457806342966c681461027d5780634b7503341461029357806370a08231146102a657806379c65068146102c557806379cc6790146102e75780638620410b146103095780638da5cb5b1461031c57806395d89b411461034b578063a6f2ae3a1461035e578063a9059cbb14610366578063b414d4b614610388578063cae9ca51146103a7578063dd62ed3e1461040c578063e4849b3214610431578063e724529c14610447578063f2fde38b1461046b575b600080fd5b341561013757600080fd5b61014560043560243561048a565b005b341561015257600080fd5b61015a6104b0565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561019657808201518382015260200161017e565b50505050905090810190601f1680156101c35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101dc57600080fd5b6101f3600160a060020a036004351660243561054e565b604051901515815260200160405180910390f35b341561021257600080fd5b61021a61057e565b60405190815260200160405180910390f35b341561023757600080fd5b6101f3600160a060020a0360043581169060243516604435610584565b341561025f57600080fd5b6102676105fb565b60405160ff909116815260200160405180910390f35b341561028857600080fd5b6101f3600435610604565b341561029e57600080fd5b61021a61068f565b34156102b157600080fd5b61021a600160a060020a0360043516610695565b34156102d057600080fd5b610145600160a060020a03600435166024356106a7565b34156102f257600080fd5b6101f3600160a060020a036004351660243561076d565b341561031457600080fd5b61021a610849565b341561032757600080fd5b61032f61084f565b604051600160a060020a03909116815260200160405180910390f35b341561035657600080fd5b61015a61085e565b6101456108c9565b341561037157600080fd5b610145600160a060020a03600435166024356108e9565b341561039357600080fd5b6101f3600160a060020a03600435166108f8565b34156103b257600080fd5b6101f360048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061090d95505050505050565b341561041757600080fd5b61021a600160a060020a0360043581169060243516610a3b565b341561043c57600080fd5b610145600435610a58565b341561045257600080fd5b610145600160a060020a03600435166024351515610ab5565b341561047657600080fd5b610145600160a060020a0360043516610b41565b60005433600160a060020a039081169116146104a557600080fd5b600791909155600855565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105465780601f1061051b57610100808354040283529160200191610546565b820191906000526020600020905b81548152906001019060200180831161052957829003601f168201915b505050505081565b600160a060020a033381166000908152600660209081526040808320938616835292905220819055600192915050565b60045481565b600160a060020a038084166000908152600660209081526040808320339094168352929052908120548211156105b957600080fd5b600160a060020a03808516600090815260066020908152604080832033909416835292905220805483900390556105f1848484610b8b565b5060019392505050565b60035460ff1681565b600160a060020a0333166000908152600560205260408120548290101561062a57600080fd5b600160a060020a03331660008181526005602052604090819020805485900390556004805485900390557fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a2506001919050565b60075481565b60056020526000908152604090205481565b60005433600160a060020a039081169116146106c257600080fd5b600160a060020a03808316600090815260056020526040808220805485019055600480548501905530909216917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905190815260200160405180910390a381600160a060020a031630600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b600160a060020a0382166000908152600560205260408120548290101561079357600080fd5b600160a060020a03808416600090815260066020908152604080832033909416835292905220548211156107c657600080fd5b600160a060020a038084166000818152600560209081526040808320805488900390556006825280832033909516835293905282902080548590039055600480548590039055907fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a250600192915050565b60085481565b600054600160a060020a031681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105465780601f1061051b57610100808354040283529160200191610546565b6000600854348115156108d857fe5b0490506108e6303383610b8b565b50565b6108f4338383610b8b565b5050565b60096020526000908152604090205460ff1681565b60008361091a818561054e565b15610a335780600160a060020a0316638f4ffcb1338630876040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156109d05780820151838201526020016109b8565b50505050905090810190601f1680156109fd5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1515610a1e57600080fd5b5af11515610a2b57600080fd5b505050600191505b509392505050565b600660209081526000928352604080842090915290825290205481565b6007548102600160a060020a033016311015610a7357600080fd5b610a7e333083610b8b565b33600160a060020a03166108fc60075483029081150290604051600060405180830381858888f1935050505015156108e657600080fd5b60005433600160a060020a03908116911614610ad057600080fd5b600160a060020a03821660009081526009602052604090819020805460ff19168315151790557f48335238b4855f35377ed80f164e8c6f3c366e54ac00b96a6402d4a9814a03a5908390839051600160a060020a039092168252151560208201526040908101905180910390a15050565b60005433600160a060020a03908116911614610b5c57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a0382161515610ba057600080fd5b600160a060020a038316600090815260056020526040902054819011610bc557600080fd5b600160a060020a03821660009081526005602052604090205481810111610beb57600080fd5b600160a060020a03831660009081526009602052604090205460ff1615610c1157600080fd5b600160a060020a03821660009081526009602052604090205460ff1615610c3757600080fd5b600160a060020a038084166000818152600560205260408082208054869003905592851680825290839020805485019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905190815260200160405180910390a35050505600a165627a7a723058205ae474e2c4754cd951ec9b654d746ca3ade0cd1ea15bc19a7479954c3d9651470029`

// DeployMyAdvancedToken deploys a new Ethereum contract, binding an instance of MyAdvancedToken to it.
func DeployMyAdvancedToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string, centralMinter common.Address) (common.Address, *types.Transaction, *MyAdvancedToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MyAdvancedTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyAdvancedTokenBin), backend, initialSupply, tokenName, tokenSymbol, centralMinter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyAdvancedToken{MyAdvancedTokenCaller: MyAdvancedTokenCaller{contract: contract}, MyAdvancedTokenTransactor: MyAdvancedTokenTransactor{contract: contract}, MyAdvancedTokenFilterer: MyAdvancedTokenFilterer{contract: contract}}, nil
}

// MyAdvancedToken is an auto generated Go binding around an Ethereum contract.
type MyAdvancedToken struct {
	MyAdvancedTokenCaller     // Read-only binding to the contract
	MyAdvancedTokenTransactor // Write-only binding to the contract
	MyAdvancedTokenFilterer   // Log filterer for contract events
}

// MyAdvancedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyAdvancedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyAdvancedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyAdvancedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyAdvancedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyAdvancedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyAdvancedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyAdvancedTokenSession struct {
	Contract     *MyAdvancedToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyAdvancedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyAdvancedTokenCallerSession struct {
	Contract *MyAdvancedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MyAdvancedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyAdvancedTokenTransactorSession struct {
	Contract     *MyAdvancedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MyAdvancedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyAdvancedTokenRaw struct {
	Contract *MyAdvancedToken // Generic contract binding to access the raw methods on
}

// MyAdvancedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyAdvancedTokenCallerRaw struct {
	Contract *MyAdvancedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyAdvancedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyAdvancedTokenTransactorRaw struct {
	Contract *MyAdvancedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyAdvancedToken creates a new instance of MyAdvancedToken, bound to a specific deployed contract.
func NewMyAdvancedToken(address common.Address, backend bind.ContractBackend) (*MyAdvancedToken, error) {
	contract, err := bindMyAdvancedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedToken{MyAdvancedTokenCaller: MyAdvancedTokenCaller{contract: contract}, MyAdvancedTokenTransactor: MyAdvancedTokenTransactor{contract: contract}, MyAdvancedTokenFilterer: MyAdvancedTokenFilterer{contract: contract}}, nil
}

// NewMyAdvancedTokenCaller creates a new read-only instance of MyAdvancedToken, bound to a specific deployed contract.
func NewMyAdvancedTokenCaller(address common.Address, caller bind.ContractCaller) (*MyAdvancedTokenCaller, error) {
	contract, err := bindMyAdvancedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenCaller{contract: contract}, nil
}

// NewMyAdvancedTokenTransactor creates a new write-only instance of MyAdvancedToken, bound to a specific deployed contract.
func NewMyAdvancedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyAdvancedTokenTransactor, error) {
	contract, err := bindMyAdvancedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenTransactor{contract: contract}, nil
}

// NewMyAdvancedTokenFilterer creates a new log filterer instance of MyAdvancedToken, bound to a specific deployed contract.
func NewMyAdvancedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyAdvancedTokenFilterer, error) {
	contract, err := bindMyAdvancedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenFilterer{contract: contract}, nil
}

// bindMyAdvancedToken binds a generic wrapper to an already deployed contract.
func bindMyAdvancedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyAdvancedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyAdvancedToken *MyAdvancedTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyAdvancedToken.Contract.MyAdvancedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyAdvancedToken *MyAdvancedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.MyAdvancedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyAdvancedToken *MyAdvancedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.MyAdvancedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyAdvancedToken *MyAdvancedTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyAdvancedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyAdvancedToken *MyAdvancedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyAdvancedToken *MyAdvancedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyAdvancedToken.Contract.Allowance(&_MyAdvancedToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyAdvancedToken.Contract.Allowance(&_MyAdvancedToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyAdvancedToken.Contract.BalanceOf(&_MyAdvancedToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyAdvancedToken.Contract.BalanceOf(&_MyAdvancedToken.CallOpts, arg0)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCaller) BuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "buyPrice")
	return *ret0, err
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenSession) BuyPrice() (*big.Int, error) {
	return _MyAdvancedToken.Contract.BuyPrice(&_MyAdvancedToken.CallOpts)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) BuyPrice() (*big.Int, error) {
	return _MyAdvancedToken.Contract.BuyPrice(&_MyAdvancedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyAdvancedToken *MyAdvancedTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyAdvancedToken *MyAdvancedTokenSession) Decimals() (uint8, error) {
	return _MyAdvancedToken.Contract.Decimals(&_MyAdvancedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) Decimals() (uint8, error) {
	return _MyAdvancedToken.Contract.Decimals(&_MyAdvancedToken.CallOpts)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_MyAdvancedToken *MyAdvancedTokenCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "frozenAccount", arg0)
	return *ret0, err
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _MyAdvancedToken.Contract.FrozenAccount(&_MyAdvancedToken.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _MyAdvancedToken.Contract.FrozenAccount(&_MyAdvancedToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenSession) Name() (string, error) {
	return _MyAdvancedToken.Contract.Name(&_MyAdvancedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) Name() (string, error) {
	return _MyAdvancedToken.Contract.Name(&_MyAdvancedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MyAdvancedToken *MyAdvancedTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MyAdvancedToken *MyAdvancedTokenSession) Owner() (common.Address, error) {
	return _MyAdvancedToken.Contract.Owner(&_MyAdvancedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) Owner() (common.Address, error) {
	return _MyAdvancedToken.Contract.Owner(&_MyAdvancedToken.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCaller) SellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "sellPrice")
	return *ret0, err
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenSession) SellPrice() (*big.Int, error) {
	return _MyAdvancedToken.Contract.SellPrice(&_MyAdvancedToken.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) SellPrice() (*big.Int, error) {
	return _MyAdvancedToken.Contract.SellPrice(&_MyAdvancedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenSession) Symbol() (string, error) {
	return _MyAdvancedToken.Contract.Symbol(&_MyAdvancedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) Symbol() (string, error) {
	return _MyAdvancedToken.Contract.Symbol(&_MyAdvancedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyAdvancedToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenSession) TotalSupply() (*big.Int, error) {
	return _MyAdvancedToken.Contract.TotalSupply(&_MyAdvancedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyAdvancedToken *MyAdvancedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MyAdvancedToken.Contract.TotalSupply(&_MyAdvancedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Approve(&_MyAdvancedToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Approve(&_MyAdvancedToken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.ApproveAndCall(&_MyAdvancedToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.ApproveAndCall(&_MyAdvancedToken.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Burn(&_MyAdvancedToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Burn(&_MyAdvancedToken.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.BurnFrom(&_MyAdvancedToken.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.BurnFrom(&_MyAdvancedToken.TransactOpts, _from, _value)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) Buy() (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Buy(&_MyAdvancedToken.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) Buy() (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Buy(&_MyAdvancedToken.TransactOpts)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "freezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.FreezeAccount(&_MyAdvancedToken.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.FreezeAccount(&_MyAdvancedToken.TransactOpts, target, freeze)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.MintToken(&_MyAdvancedToken.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.MintToken(&_MyAdvancedToken.TransactOpts, target, mintedAmount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Sell(&_MyAdvancedToken.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Sell(&_MyAdvancedToken.TransactOpts, amount)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) SetPrices(opts *bind.TransactOpts, newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "setPrices", newSellPrice, newBuyPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) SetPrices(newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.SetPrices(&_MyAdvancedToken.TransactOpts, newSellPrice, newBuyPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) SetPrices(newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.SetPrices(&_MyAdvancedToken.TransactOpts, newSellPrice, newBuyPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Transfer(&_MyAdvancedToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.Transfer(&_MyAdvancedToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.TransferFrom(&_MyAdvancedToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.TransferFrom(&_MyAdvancedToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyAdvancedToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MyAdvancedToken *MyAdvancedTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.TransferOwnership(&_MyAdvancedToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MyAdvancedToken *MyAdvancedTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyAdvancedToken.Contract.TransferOwnership(&_MyAdvancedToken.TransactOpts, newOwner)
}

// MyAdvancedTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the MyAdvancedToken contract.
type MyAdvancedTokenBurnIterator struct {
	Event *MyAdvancedTokenBurn // Event containing the contract specifics and raw log

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
func (it *MyAdvancedTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyAdvancedTokenBurn)
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
		it.Event = new(MyAdvancedTokenBurn)
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
func (it *MyAdvancedTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyAdvancedTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyAdvancedTokenBurn represents a Burn event raised by the MyAdvancedToken contract.
type MyAdvancedTokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*MyAdvancedTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MyAdvancedToken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenBurnIterator{contract: _MyAdvancedToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *MyAdvancedTokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MyAdvancedToken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyAdvancedTokenBurn)
				if err := _MyAdvancedToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// MyAdvancedTokenFrozenFundsIterator is returned from FilterFrozenFunds and is used to iterate over the raw logs and unpacked data for FrozenFunds events raised by the MyAdvancedToken contract.
type MyAdvancedTokenFrozenFundsIterator struct {
	Event *MyAdvancedTokenFrozenFunds // Event containing the contract specifics and raw log

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
func (it *MyAdvancedTokenFrozenFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyAdvancedTokenFrozenFunds)
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
		it.Event = new(MyAdvancedTokenFrozenFunds)
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
func (it *MyAdvancedTokenFrozenFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyAdvancedTokenFrozenFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyAdvancedTokenFrozenFunds represents a FrozenFunds event raised by the MyAdvancedToken contract.
type MyAdvancedTokenFrozenFunds struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFrozenFunds is a free log retrieval operation binding the contract event 0x48335238b4855f35377ed80f164e8c6f3c366e54ac00b96a6402d4a9814a03a5.
//
// Solidity: event FrozenFunds(target address, frozen bool)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) FilterFrozenFunds(opts *bind.FilterOpts) (*MyAdvancedTokenFrozenFundsIterator, error) {

	logs, sub, err := _MyAdvancedToken.contract.FilterLogs(opts, "FrozenFunds")
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenFrozenFundsIterator{contract: _MyAdvancedToken.contract, event: "FrozenFunds", logs: logs, sub: sub}, nil
}

// WatchFrozenFunds is a free log subscription operation binding the contract event 0x48335238b4855f35377ed80f164e8c6f3c366e54ac00b96a6402d4a9814a03a5.
//
// Solidity: event FrozenFunds(target address, frozen bool)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) WatchFrozenFunds(opts *bind.WatchOpts, sink chan<- *MyAdvancedTokenFrozenFunds) (event.Subscription, error) {

	logs, sub, err := _MyAdvancedToken.contract.WatchLogs(opts, "FrozenFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyAdvancedTokenFrozenFunds)
				if err := _MyAdvancedToken.contract.UnpackLog(event, "FrozenFunds", log); err != nil {
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

// MyAdvancedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyAdvancedToken contract.
type MyAdvancedTokenTransferIterator struct {
	Event *MyAdvancedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyAdvancedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyAdvancedTokenTransfer)
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
		it.Event = new(MyAdvancedTokenTransfer)
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
func (it *MyAdvancedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyAdvancedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyAdvancedTokenTransfer represents a Transfer event raised by the MyAdvancedToken contract.
type MyAdvancedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyAdvancedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyAdvancedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyAdvancedTokenTransferIterator{contract: _MyAdvancedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyAdvancedToken *MyAdvancedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyAdvancedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyAdvancedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyAdvancedTokenTransfer)
				if err := _MyAdvancedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101668061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008c575b600080fd5b341561005b57600080fd5b6100636100ba565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561009757600080fd5b6100b873ffffffffffffffffffffffffffffffffffffffff600435166100d6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100fe57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555600a165627a7a7230582014bd8bf299fb1f821fa405199311fbc8d485a17282036b98e59bd49708a42a8a0029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x60606040526002805460ff19166012179055341561001c57600080fd5b604051610a1b380380610a1b833981016040528080519190602001805182019190602001805160025460ff16600a0a85026003819055600160a060020a03331660009081526004602052604081209190915592019190508280516100849291602001906100a1565b5060018180516100989291602001906100a1565b5050505061013c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100e257805160ff191683800117855561010f565b8280016001018555821561010f579182015b8281111561010f5782518255916020019190600101906100f4565b5061011b92915061011f565b5090565b61013991905b8082111561011b5760008155600101610125565b90565b6108d08061014b6000396000f3006060604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461017e57806323b872dd146101a3578063313ce567146101cb57806342966c68146101f457806370a082311461020a57806379cc67901461022957806395d89b411461024b578063a9059cbb1461025e578063cae9ca5114610282578063dd62ed3e146102e7575b600080fd5b34156100c957600080fd5b6100d161030c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561010d5780820151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015357600080fd5b61016a600160a060020a03600435166024356103aa565b604051901515815260200160405180910390f35b341561018957600080fd5b6101916103da565b60405190815260200160405180910390f35b34156101ae57600080fd5b61016a600160a060020a03600435811690602435166044356103e0565b34156101d657600080fd5b6101de610457565b60405160ff909116815260200160405180910390f35b34156101ff57600080fd5b61016a600435610460565b341561021557600080fd5b610191600160a060020a03600435166104eb565b341561023457600080fd5b61016a600160a060020a03600435166024356104fd565b341561025657600080fd5b6100d16105d9565b341561026957600080fd5b610280600160a060020a0360043516602435610644565b005b341561028d57600080fd5b61016a60048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061065395505050505050565b34156102f257600080fd5b610191600160a060020a0360043581169060243516610781565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a25780601f10610377576101008083540402835291602001916103a2565b820191906000526020600020905b81548152906001019060200180831161038557829003601f168201915b505050505081565b600160a060020a033381166000908152600560209081526040808320938616835292905220819055600192915050565b60035481565b600160a060020a0380841660009081526005602090815260408083203390941683529290529081205482111561041557600080fd5b600160a060020a038085166000908152600560209081526040808320339094168352929052208054839003905561044d84848461079e565b5060019392505050565b60025460ff1681565b600160a060020a0333166000908152600460205260408120548290101561048657600080fd5b600160a060020a03331660008181526004602052604090819020805485900390556003805485900390557fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a2506001919050565b60046020526000908152604090205481565b600160a060020a0382166000908152600460205260408120548290101561052357600080fd5b600160a060020a038084166000908152600560209081526040808320339094168352929052205482111561055657600080fd5b600160a060020a038084166000818152600460209081526040808320805488900390556005825280832033909516835293905282902080548590039055600380548590039055907fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a250600192915050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a25780601f10610377576101008083540402835291602001916103a2565b61064f33838361079e565b5050565b60008361066081856103aa565b156107795780600160a060020a0316638f4ffcb1338630876040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156107165780820151838201526020016106fe565b50505050905090810190601f1680156107435780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b151561076457600080fd5b5af1151561077157600080fd5b505050600191505b509392505050565b600560209081526000928352604080842090915290825290205481565b6000600160a060020a03831615156107b557600080fd5b600160a060020a038416600090815260046020526040902054829010156107db57600080fd5b600160a060020a0383166000908152600460205260409020548281011161080157600080fd5b50600160a060020a0380831660008181526004602052604080822080549488168084528284208054888103909155938590528154870190915591909301927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3600160a060020a0380841660009081526004602052604080822054928716825290205401811461089e57fe5b505050505600a165627a7a7230582080f7da6ad0c40ed42d78d573e0034169a5eb017be82ec9dfe59177c82e98ea590029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Token *TokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Token *TokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Token *TokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Token *TokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Token *TokenSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, _from, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Token contract.
type TokenBurnIterator struct {
	Event *TokenBurn // Event containing the contract specifics and raw log

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
func (it *TokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurn)
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
		it.Event = new(TokenBurn)
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
func (it *TokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurn represents a Burn event raised by the Token contract.
type TokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_Token *TokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*TokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenBurnIterator{contract: _Token.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_Token *TokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *TokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurn)
				if err := _Token.contract.UnpackLog(event, "Burn", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TokenRecipientABI is the input ABI used to generate the binding from.
const TokenRecipientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenRecipientBin is the compiled bytecode used for deploying new contracts.
const TokenRecipientBin = `0x`

// DeployTokenRecipient deploys a new Ethereum contract, binding an instance of TokenRecipient to it.
func DeployTokenRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRecipient, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// TokenRecipient is an auto generated Go binding around an Ethereum contract.
type TokenRecipient struct {
	TokenRecipientCaller     // Read-only binding to the contract
	TokenRecipientTransactor // Write-only binding to the contract
	TokenRecipientFilterer   // Log filterer for contract events
}

// TokenRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRecipientSession struct {
	Contract     *TokenRecipient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRecipientCallerSession struct {
	Contract *TokenRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRecipientTransactorSession struct {
	Contract     *TokenRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRecipientRaw struct {
	Contract *TokenRecipient // Generic contract binding to access the raw methods on
}

// TokenRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRecipientCallerRaw struct {
	Contract *TokenRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRecipientTransactorRaw struct {
	Contract *TokenRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRecipient creates a new instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipient(address common.Address, backend bind.ContractBackend) (*TokenRecipient, error) {
	contract, err := bindTokenRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// NewTokenRecipientCaller creates a new read-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientCaller(address common.Address, caller bind.ContractCaller) (*TokenRecipientCaller, error) {
	contract, err := bindTokenRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientCaller{contract: contract}, nil
}

// NewTokenRecipientTransactor creates a new write-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRecipientTransactor, error) {
	contract, err := bindTokenRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientTransactor{contract: contract}, nil
}

// NewTokenRecipientFilterer creates a new log filterer instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRecipientFilterer, error) {
	contract, err := bindTokenRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientFilterer{contract: contract}, nil
}

// bindTokenRecipient binds a generic wrapper to an already deployed contract.
func bindTokenRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.TokenRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactor) ReceiveApproval(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.contract.Transact(opts, "receiveApproval", _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactorSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}
