package account

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet interface {
	//create public key
	CreatePublicKey() string
	//create private key
	CreatePrivateKey() string
	//create address
	CreateAddress() string
}

var walletInstance *WalletSingleton

type WalletSingleton struct {
	once sync.Once
}

func (w *WalletSingleton) GetSingleton() *WalletSingleton {
	w.once.Do(func() {
		walletInstance = &WalletSingleton{}
	})
	return walletInstance
}

//create public key
func (w *WalletSingleton) CreatePublicKey() string {
	//generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return privateKey.PublicKey.X.String() + privateKey.PublicKey.Y.String()
}

//create private key
func (w *WalletSingleton) CreatePrivateKey() string {
	//generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	return privateKey.D.String()
}

//create address
func (w *WalletSingleton) CreateAddress() string {
	//generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).String()
}
