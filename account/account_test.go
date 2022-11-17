package account

import (
	"testing"

	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func TestGetBalance(t *testing.T) {
	var Account AccountSingleton
	//指针来调用
	instance := Account.GetSingleton()

	log.Info(instance.GetBalance("0xeE0D1B5aeAd1Ac95233e43D87413426BeF113C79"))
}
