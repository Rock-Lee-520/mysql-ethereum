package account

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account interface {
}

var accountInstance *AccountSingleton

type AccountSingleton struct {
	once sync.Once
}

func (account *AccountSingleton) GetSingleton() *AccountSingleton {
	account.once.Do(func() {
		accountInstance = &AccountSingleton{}
	})
	return accountInstance
}

func (account *AccountSingleton) GetBalance(address string, client *ethclient.Client) string {
	//client, err := ethclient.Dial("https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	// client, err := ethclient.Dial("https://cloudflare-eth.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	accountAddress := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 1000000000000000000

	// wei转换为以太坊
	eth := new(big.Float)
	eth.SetString(balance.String())

	//	eth.Div(eth, big.NewFloat(math.Pow10(18)))

	return eth.String()
}

func (account *AccountSingleton) CheckAddress(address string, client *ethclient.Client) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString(address)) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString(address)) // is valid: false

	// client, err := ethclient.Dial("https://cloudflare-eth.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 0x Protocol Token (ZRX) smart contract address
	addr := common.HexToAddress(address)
	bytecode, err := client.CodeAt(context.Background(), addr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	addr = common.HexToAddress(address)
	bytecode, err = client.CodeAt(context.Background(), addr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
	return isContract
}
