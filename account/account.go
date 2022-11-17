package account

import (
	"context"
	"fmt"
	"log"
	"math/big"
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

func (account *AccountSingleton) GetBalance(address string) string {
	//client, err := ethclient.Dial("https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

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

// func main2() {
// 	client, err := ethclient.Dial("https://cloudflare-eth.com")
// 	//client, err := ethclient.Dial("https://127.0.0.1:8545")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	account := common.HexToAddress("0xeE0D1B5aeAd1Ac95233e43D87413426BeF113C79")
// 	balance, err := client.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balance) // 25893180161173005034

// 	blockNumber := big.NewInt(15929513)
// 	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balanceAt) // 25729324269165216042

// 	fbalance := new(big.Float)
// 	fbalance.SetString(balanceAt.String())
// 	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(10)))
// 	fmt.Println(ethValue) // 25.729324269165216041

// 	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
// 	fmt.Println(pendingBalance) // 25729324269165216042

// }
