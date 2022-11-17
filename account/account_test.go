package account

import (
	"testing"

	Client "github.com/Rock-liyi/p2pdb-ethereum/client"

	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func TestGetBalance(t *testing.T) {
	var Account AccountSingleton
	//指针来调用
	accountInstance := Account.GetSingleton()

	var Client Client.ClientSingleton
	//指针来调用
	clientInstance := Client.GetSingleton()
	newClient := clientInstance.NewClient()
	log.Info(accountInstance.GetBalance("0xeE0D1B5aeAd1Ac95233e43D87413426BeF113C79", newClient))
}

func TestCheckAddress(t *testing.T) {
	var Account AccountSingleton
	//指针来调用
	accountInstance := Account.GetSingleton()

	var Client Client.ClientSingleton
	//指针来调用
	clientInstance := Client.GetSingleton()
	newClient := clientInstance.NewClient()
	log.Info(accountInstance.CheckAddress("0xeE0D1B5aeAd1Ac95233e43D87413426BeF113C79", newClient))
}
