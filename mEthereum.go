package main

import (
	"log"
	"math/big"
	"os"

	"github.com/joho/godotenv"

	//"strconv"
	"context"
	"crypto/ecdsa"

	// "crypto/rand"
	// "encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
	//"github.com/miguelmota/go-solidity-sha3"
)

var rewardCtrtAddr string = "0xBF76248d5e3bfd1d4dDE4369Fe6163289A0267F6"

func preEthereum() {
	print("------------------== preEthereum")

	//------------------==get env and config files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	address1 := os.Getenv("ADDRESS1")
	address1pk := os.Getenv("ADDRESS1PK")//no 0x
	address2 := os.Getenv("ADDRESS2")
	print("address1:", address1, "\naddress1pk:", address1pk, "\naddress2:", address2)

	const configFileName = "config.txt"
	var myenv map[string]string
	if myenv, err = godotenv.Read(configFileName); err != nil {
		print("could not load config file from", configFileName, err)
	}
	addrLPToken := myenv["ADDRLPTOKEN"]
	addrRewardToken := myenv["ADDRREWARDTOKEN"]
	addrRewardCtrt := myenv["ADDRREWARDCTRT"]
	print("\naddrLPToken", addrLPToken, "\naddrRewardToken:", addrRewardToken, "\naddrRewardCtrt:", addrRewardCtrt)

	//os.Exit(0)
	//------------------==getting address from private key
	ctx := context.Background()
	addr1 := common.HexToAddress(address1)
	addr2 := common.HexToAddress(address2)
	addrByteLPToken := common.HexToAddress(addrLPToken)
	//addrByteRewardToken := common.HexToAddress(addrRewardToken)
	// addrByteRewardCtrt := common.HexToAddress(addrRewardCtrt)

	/*
		    key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		    if err != nil {
		        panic(err)
		    }
					privateKey, err := crypto.GenerateKey()
		    if err != nil {
		        log.Fatal(err)
		    }
	*/
	// type PrivateKey struct {
	//     PublicKey
	//     D   *big.Int
	// }
	prkECDSA1ptr, err := crypto.HexToECDSA(address1pk)
	//func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error)
	if err != nil {
		log.Fatal(err)
	}
	//prkBytes := crypto.FromECDSA(prkECDSA1ptr)
	//print(`prkBytes:`,prkBytes)
	print("\nprkECDSA1ptr Type: %T, Value: %v\n", prkECDSA1ptr, prkECDSA1ptr)

	pukCryptoPtr1 := prkECDSA1ptr.Public()
	// type "crypto".PublicKey
	print("\npukCryptoPtr1 Type: %T, Value: %v\n", pukCryptoPtr1, pukCryptoPtr1) //Type: *ecdsa.PublicKey
	pukECDSAptr1, ok := pukCryptoPtr1.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: pukCryptoPtr1 is not of type *ecdsa.PublicKey")
	}
	print("pukECDSAptr1 Type: %T, Value: %v\n", pukECDSAptr1, pukECDSAptr1)
	//(type *ecdsa.PublicKey)

	//publicKeyBytes := crypto.FromECDSAPub(pukECDSAptr1)
	// used for crypto.VerifySignature
	//func FromECDSAPub(pub *ecdsa.PublicKey) []byte

	fromAddress := crypto.PubkeyToAddress(*pukECDSAptr1)
	//func PubkeyToAddress(p ecdsa.PublicKey) common.Address
	print("\naddress1 from prvKey:", fromAddress.Hex())

	//----------------------== connecting to Ethereum
	// type Client struct {
	// 	// contains filtered or unexported fields
	// }
	// var conn Client

	EthNodeURL := "https://rinkeby.infura.io/v3/34d79804349241d8a6bbfb1351e33a62"
	//"http://127.0.0.1:8545"
	conn, err := ethclient.Dial(EthNodeURL)
	// For an IPC based RPC connection to a remote node: /mnt/sda5/ethereum/geth.ipc

	if err != nil {
		log.Fatalf("failed to connect to the Ethereum network: %v", err)
	}
	print("connection to Ethereum successful")

	balanceAddr1, _ := conn.BalanceAt(ctx, addr1, nil)
	print("Ether Balance addr1:", balanceAddr1)
	balanceAddr2, _ := conn.BalanceAt(ctx, addr2, nil)
	print("Ether Balance addr2:", balanceAddr2)

	nonceM, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	print("nonceM:", nonceM)

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	} // estimate gas price

	amountToSend := big.NewInt(1000)
	//amount.SetString("1000000000000000000000", 10)
	//0.01 wei = 10000000000000000
	//nonce = big.NewInt(int64(nonce))

	auth := bind.NewKeyedTransactor(prkECDSA1ptr)
	auth.Nonce = big.NewInt(int64(nonceM))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//https://github.com/what-the-func/golang-ethereum-transfer-tokens/blob/master/main.go

	instance, err := NewRewards(addrByteLPToken, conn)
	if err != nil {
		log.Fatal(err)
	}

	blockTimestamp, periodFinish, rewardRate, DURATION, err := instance.GetData1(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve GetData1: %v", err)
	}
	print(blockTimestamp, periodFinish, rewardRate, DURATION)

	balanceTokenUser1, err := instance.BalanceOf(&bind.CallOpts{}, addr1)
	if err != nil {
		log.Fatal(err)
	}
	print("Token balance of addr1:", balanceTokenUser1)

	// balanceTokenkUser2, err := instance.BalanceOf(&bind.CallOpts{}, addr2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// print("Token balance of addr2:", balanceTokenkUser2)

	//os.Exit(0)
	isToSend := false
	if isToSend {
		tx, err := instance.Stake(auth, amountToSend)
		if err != nil {
			log.Fatal(err)
		}
		print("transaction hash: %s", tx.Hash().Hex())
	} else {
		print("\nno transaction was made")
		print("\namountToSend: %s", amountToSend)
	}
}

/*


 */
