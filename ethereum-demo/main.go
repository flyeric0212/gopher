/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/3/14 上午11:44
 */
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
	"strings"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

// go run main.go token.go
func main() {
	fmt.Println("start")

	handle_start_time := time.Now()
	//NewAccount()

	//NewAccount2()

	Tranfer("0x4AB5463C0B0F7a2D36e56e312A160D255918fA39")
	fmt.Printf("costs: %f s \n", time.Now().Sub(handle_start_time).Seconds())

	//addr, _ := NewAccount()
	//Tranfer(addr)

	//Tranfer("0x537e0ad4f869fa026cd2eb7523fdf7e361539a85")

	//1200000000000000000000
	//TransferFrom("0x537e0ad4F869FA026Cd2EB7523FDF7E361539a85", "0xB0cAb88a2D0ddb9EFBf46fACaC2e4491677A3787")
}

func NewAccount() (addr string, key string) {
	var err error

	ks := keystore.NewKeyStore("/Users/liangbo/Library/Ethereum/rinkeby/keystore/", keystore.StandardScryptN, keystore.StandardScryptP)
	address, _ := ks.NewAccount("jx19910212")
	//key_json, err := ks.Export(address, "jx19910212", "jx19910212")
	//if err != nil {
	//	fmt.Printf("err: %v \n", err)
	//	return
	//}
	addr = address.Address.Hex()

	// 获取keystore
	var key_store string
	err, key_store = exec_shell("/Users/liangbo/Documents/go_workspace/src/gopher/ethereum_demo/keystore.sh " +
		"/Users/liangbo/Library/Ethereum/rinkeby/keystore/" + " " + strings.ToLower(addr[2:]))
	if err != nil {
		fmt.Printf("call exec_shell fail.[err:%v]", err)
		return
	}

	key = key_store
	fmt.Printf("address: %s, private_key: %s \n", addr, key_store)
	return
}

func httpPost(url string, args string) (err error, res string) {
	resp, err := http.Post(url, "application/json", strings.NewReader(args))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res = string(body)
	return
}

func exec_shell(s string) (err error, res string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	res = out.String()
	return
}

func NewAccount2() (addr string, key string) {

	password := "jx19910212"
	// 创建帐号
	http_req := `{"jsonrpc":"2.0","method":"personal_newAccount","params":["` + password + `"],"id":1}`
	var http_resp string
	err, http_resp := httpPost("http://127.0.0.1:8545", http_req)
	if err != nil {
		fmt.Printf("call httpPost fail.[err:%v]", err)
		return
	}
	type EthereumRPC struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int64  `json:"id"`
		Result  string `json:"result"`
	}

	// 获取帐号地址
	var ethereum_rpc EthereumRPC
	err = json.Unmarshal([]byte(http_resp), &ethereum_rpc)
	if nil != err {
		fmt.Printf("call json.Unmarshal fail.[err:%v]", err)
		return
	}

	// 获取keystore
	var key_store string
	err, key_store = exec_shell("/Users/liangbo/Documents/go_workspace/src/gopher/ethereum_demo/keystore.sh " +
		"/Users/liangbo/Library/Ethereum/rinkeby/keystore/" + " " + ethereum_rpc.Result[2:])
	if err != nil {
		fmt.Printf("call exec_shell fail.[err:%v]", err)
		return
	}

	addr = ethereum_rpc.Result
	key = key_store
	fmt.Printf("address: %s, private_key: %s \n", addr, key)
	return
}

func Tranfer(to_address string) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	//首先导入上面生成的账户密钥（json）和密码
	// account 2
	keystore_json := "{\"address\":\"b0cab88a2d0ddb9efbf46facac2e4491677a3787\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"40cd36724af2877347ff88551e95c6b09c0f831ab5cfcb3612bb4546430c0fc5\",\"cipherparams\":{\"iv\":\"2dd7d8a43442d1ad7b6439daaf6290ee\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"f3785e4a1413638329d7229c149f142a3a32b6d6d67e68629142647dd389adfa\"},\"mac\":\"480a539a7c18019e0f1bc252d7a7759d243496bda581d26e0dbfb281be6d9126\"},\"id\":\"ed8638e2-44d4-48d2-8f43-52d4d0be2b22\",\"version\":3}"
	auth, err := bind.NewTransactor(strings.NewReader(keystore_json), "jx19910212")

	// account 1
	//keystore_json := "{\"address\":\"848e4c128827eb56ba3d42e29c53a3bfd7bd92cf\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"0374960044dd31765c2ec1e001db25f8565ce6680f57e9482874d9a515d9a564\",\"cipherparams\":{\"iv\":\"c59b0276f8923187e6519562ce26a86d\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"993d6cd53f5ed4786f6b3d14cda3c221af63f83c95f36254c03d17dcc5e87548\"},\"mac\":\"3b458d9b5ee00856af54cf957fcd8989c9cb625816a5d33330087a4e4e2a155a\"},\"id\":\"d34f4227-9dea-4f95-8b9d-13dba47d3812\",\"version\":3}"
	// auth, err := bind.NewTransactor(strings.NewReader(keystore_json), "jx19910212")

	// main account
	//keystore_json := "{\"address\":\"537e0ad4f869fa026cd2eb7523fdf7e361539a85\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"5907c1367c301bd439639ca377c6208ddc26074726ff2abdb27c0183843ddf1c\",\"cipherparams\":{\"iv\":\"2c9b884fc16f2269d02ed181b2620785\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"84f038ce7372b7f4136332016fd1102c5e7652be0b25965b2c32aa7d6f1efc73\"},\"mac\":\"45bbd9f3b37b76807ef19d4267d0f7044d47b29b127867a52aa9fd739ef015bf\"},\"id\":\"f38150e2-d0fd-4fb7-ad88-81e13e1357ce\",\"version\":3}"
	//auth, err := bind.NewTransactor(strings.NewReader(keystore_json), "Jx19910212@")
	if nil != err {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("from: %s \n", auth.From.Hex())

	// 查看合约地址
	// https://etherscan.io/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0
	// https://rinkeby.etherscan.io/token/0x2F814dBebf9d5ac77bCda9a192B387FC43325873 wtccoin
	// https://rinkeby.etherscan.io/token/0x5Eb4db894e254510f83Ae4c4311Dd431C92E94Ba gogocoin
	token, err := NewMyAdvancedToken(common.HexToAddress("0x5Eb4db894e254510f83Ae4c4311Dd431C92E94Ba"), client)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	balance, _ := token.BalanceOf(nil, common.HexToAddress(to_address))
	fmt.Println("balance:%v", balance.String())

	//每个代币都会有相应的位数，例如代币是18位，那么我们转账的时候，需要在金额后面加18个0
	decimal, err := token.Decimals(nil)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	amount := big.NewFloat(100.00)
	//这是处理位数的代码段
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	fmt.Printf("limit: %v, price: %v \n", auth.GasLimit, auth.GasPrice)

	auth.GasLimit = uint64(3000000)

	// 查看交易信息
	// https://rinkeby.etherscan.io/tx/0xe891847b31413c9a48b1de2ccd6397b4bd54e849a888e6f1e048fae701a8870c
	tx, err := token.Transfer(auth, common.HexToAddress(to_address), convertAmount)
	if nil != err {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("result: %v\n", tx)

	fmt.Printf("cost: %v, price: %v, gas: %v", tx.Cost(), tx.GasPrice(), tx.Gas())

}

func TransferFrom(from_address string, to_address string) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	//首先导入上面生成的账户密钥（json）和密码
	//keystore_json := "{\"address\":\"848e4c128827eb56ba3d42e29c53a3bfd7bd92cf\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"0374960044dd31765c2ec1e001db25f8565ce6680f57e9482874d9a515d9a564\",\"cipherparams\":{\"iv\":\"c59b0276f8923187e6519562ce26a86d\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"993d6cd53f5ed4786f6b3d14cda3c221af63f83c95f36254c03d17dcc5e87548\"},\"mac\":\"3b458d9b5ee00856af54cf957fcd8989c9cb625816a5d33330087a4e4e2a155a\"},\"id\":\"d34f4227-9dea-4f95-8b9d-13dba47d3812\",\"version\":3}"
	//auth, err := bind.NewTransactor(strings.NewReader(keystore_json), "jx19910212")

	keystore_json := "{\"address\":\"537e0ad4f869fa026cd2eb7523fdf7e361539a85\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"5907c1367c301bd439639ca377c6208ddc26074726ff2abdb27c0183843ddf1c\",\"cipherparams\":{\"iv\":\"2c9b884fc16f2269d02ed181b2620785\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"84f038ce7372b7f4136332016fd1102c5e7652be0b25965b2c32aa7d6f1efc73\"},\"mac\":\"45bbd9f3b37b76807ef19d4267d0f7044d47b29b127867a52aa9fd739ef015bf\"},\"id\":\"f38150e2-d0fd-4fb7-ad88-81e13e1357ce\",\"version\":3}"
	auth, err := bind.NewTransactor(strings.NewReader(keystore_json), "Jx19910212@")
	if nil != err {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("from: %s \n", auth.From.Hex())

	// 查看合约地址
	// https://etherscan.io/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0
	// https://rinkeby.etherscan.io/token/0x2F814dBebf9d5ac77bCda9a192B387FC43325873
	token, err := NewMyAdvancedToken(common.HexToAddress("0x2F814dBebf9d5ac77bCda9a192B387FC43325873"), client)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	//每个代币都会有相应的位数，例如eos是18位，那么我们转账的时候，需要在金额后面加18个0
	decimal, err := token.Decimals(nil)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	amount := big.NewFloat(100.00)
	//这是处理位数的代码段
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	// 查看交易信息
	// https://rinkeby.etherscan.io/tx/0xe891847b31413c9a48b1de2ccd6397b4bd54e849a888e6f1e048fae701a8870c
	txs, err := token.TransferFrom(auth, common.HexToAddress(from_address), common.HexToAddress(to_address), convertAmount)
	if nil != err {
		fmt.Printf("err: %v \n", err)
		return
	}

	fmt.Printf("result: %v\n", txs)
}
