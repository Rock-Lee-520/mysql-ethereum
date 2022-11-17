package account

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

// type Wallet interface {
// 	//create public key
// 	CreatePublicKey() string
// 	//create private key
// 	CreatePrivateKey() string
// 	//create address
// 	CreateAddress() string
// }

// func main() {
// 	privateKey, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKeyBytes := crypto.FromECDSA(privateKey)
// 	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
// 	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

// 	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
// 	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

// 	hash := sha3.NewLegacyKeccak256()
// 	hash.Write(publicKeyBytes[1:])
// 	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e

// }

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
