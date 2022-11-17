package account

import (
	"testing"

	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func TestCreatePublicKey(t *testing.T) {
	var Wallet WalletSingleton
	//指针来调用
	instance := Wallet.GetSingleton()

	log.Info(instance.CreatePrivateKey())
}

func TestCreatePrivateKey(t *testing.T) {
	var Wallet WalletSingleton
	//指针来调用
	instance := Wallet.GetSingleton()

	log.Info(instance.CreatePrivateKey())
}

func TestCreateAddress(t *testing.T) {
	var Wallet WalletSingleton
	//指针来调用
	instance := Wallet.GetSingleton()

	log.Info(instance.CreateAddress())
}
