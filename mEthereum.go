package main

import (
	"errors"
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

var n1 = big.NewInt(-1)

func getRewardsCtrtValues(addrRewardsPool string, network string) (*big.Int, *big.Int, error) {
	log1("------------== getRewardsCtrtValues")
	if addrRewardsPool == "" {
		return n1, n1, errors.New("reward contract address is invalid")
	}

	log1("gett env file")
	err := godotenv.Load()
	if err != nil {
		log1("err@ loading .env file.", err)
		return n1, n1, err
	}
	address1 := os.Getenv("ADDRESS1")
	address1pk := os.Getenv("ADDRESS1PK") //no 0x
	address2 := os.Getenv("ADDRESS2")
	log1("address1:", address1, "\naddress2:", address2)
	//"\naddress1pk:", address1pk,

	// addrLPToken := myenv["ADDRLPTOKEN"]
	// addrRewardToken := myenv["ADDRREWARDTOKEN"]
	//myenv["ADDRREWARDCTRT"]
	log1("\nnetwork:", network,
		"\naddrRewardsPool:", addrRewardsPool)
	//"\naddrLPToken", addrLPToken, "\naddrRewardToken:", addrRewardToken,

	//os.Exit(0)
	//------------------==getting address from private key
	ctx := context.Background()
	addr1 := common.HexToAddress(address1)
	addr2 := common.HexToAddress(address2)
	//addrByteLPToken := common.HexToAddress(addrLPToken)
	//addrByteRewardToken := common.HexToAddress(addrRewardToken)
	addrByteRewardCtrt := common.HexToAddress(addrRewardsPool)

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
		log1("err@ HexToECDSA.", err)
		return n1, n1, err
	}
	//prkBytes := crypto.FromECDSA(prkECDSA1ptr)
	//log1(`prkBytes:`,prkBytes)

	//log1("\nprkECDSA1ptr Type: %T, Value: %v\n", prkECDSA1ptr, prkECDSA1ptr)

	pukCryptoPtr1 := prkECDSA1ptr.Public()
	// type "crypto".PublicKey

	//log1("\npukCryptoPtr1 Type: %T, Value: %v\n", pukCryptoPtr1, pukCryptoPtr1) //Type: *ecdsa.PublicKey
	pukECDSAptr1, ok := pukCryptoPtr1.(*ecdsa.PublicKey)
	if !ok {
		log1("err@ cannot assert type: pukCryptoPtr1 is not of type *ecdsa.PublicKey.")
		return n1, n1, nil
	}
	//log1("pukECDSAptr1 Type: %T, Value: %v\n", pukECDSAptr1, pukECDSAptr1)
	//(type *ecdsa.PublicKey)

	//publicKeyBytes := crypto.FromECDSAPub(pukECDSAptr1)
	// used for crypto.VerifySignature
	//func FromECDSAPub(pub *ecdsa.PublicKey) []byte

	fromAddress := crypto.PubkeyToAddress(*pukECDSAptr1)
	//func PubkeyToAddress(p ecdsa.PublicKey) common.Address
	log1("\naddress1 from prvKey:", fromAddress.Hex())

	//----------------------== connecting to Ethereum
	// type Client struct {
	// 	// contains filtered or unexported fields
	// }
	// var conn Client

	var EthNodeURL string
	switch network{
		case "mainnet":
			EthNodeURL = os.Getenv("ETHEREUMMAIN")
		case "rinkeby":
			EthNodeURL = os.Getenv("ETHEREUMRINKEBY")
	}

	conn, err := ethclient.Dial(EthNodeURL)
	// For an IPC based RPC connection to a remote node: /mnt/sda5/ethereum/geth.ipc
	if err != nil {
		log1("err@ ethclient.Dial().", err)
		return n1, n1, err
	}
	log1("connection to Ethereum successful")

	balanceAddr1, _ := conn.BalanceAt(ctx, addr1, nil)
	log1("Ether Balance addr1:", balanceAddr1)
	balanceAddr2, _ := conn.BalanceAt(ctx, addr2, nil)
	log1("Ether Balance addr2:", balanceAddr2)

	nonceM, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log1("err@ PendingNonceAt.", err)
		return n1, n1, err
	}
	log1("nonceM:", nonceM)

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log1("err@ SuggestGasPrice.", err)
		return n1, n1, err
	} // estimate gas price

	auth := bind.NewKeyedTransactor(prkECDSA1ptr)
	auth.Nonce = big.NewInt(int64(nonceM))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//https://github.com/what-the-func/golang-ethereum-transfer-tokens/blob/master/main.go

	instRewards, err := NewRewards(addrByteRewardCtrt, conn)
	if err != nil {
		log1("Failed to make new Rewards contract instance", err)
		return n1, n1, err
	}
	//var instERC20 = nil
	// instRewardToken

	//log1("------------------== getRewardsCtrtValues")
	amountToSend := big.NewInt(1000)
	//amount.SetString("1000000000000000000000", 10)
	//0.01 wei = 10000000000000000
	//nonce = big.NewInt(int64(nonce))

	stakedBalanceUser1, err := instRewards.BalanceOf(nil, addr1) // &bind.CallOpts{}
	if err != nil {
		log1("Failed to retrieve BalanceOf", err)
		return n1, n1, err
	}
	log1("staked balance of addr1:", stakedBalanceUser1)

	rewardRate, err := instRewards.RewardRate(nil)
	if err != nil {
		log1("Failed to retrieve RewardRate", err)
		return n1, n1, err
	}
	log1("rewardRate:", rewardRate)

	totalSupply, err := instRewards.TotalSupply(nil)
	if err != nil {
		log1("Failed to retrieve totalSupply")
		return n1, n1, err
	}
	log1("totalSupply:", totalSupply)

	// blockTimestamp, periodFinish, rewardRate, DURATION, err := instRewards.GetData1(nil)
	// if err != nil {
	// 	log1("Failed to retrieve GetData1: %v", err)
	// }
	// log1("blockTimestamp:", blockTimestamp, ", periodFinish:", periodFinish, ", rewardRate:", rewardRate, ", DURATION:", DURATION)

	// balanceTokenkUser2, err := instRewards.BalanceOf(&bind.CallOpts{}, addr2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log1("Token balance of addr2:", balanceTokenkUser2)

	//os.Exit(0)
	isToSend := false
	if isToSend {
		tx, err := instRewards.Stake(auth, amountToSend)
		if err != nil {
			return n1, n1, err
		}
		log1("transaction hash: %s", tx.Hash().Hex())
	} else {
		log1("\nno transaction was made")
		log1("\namountToSend: %s", amountToSend)
	}

	return rewardRate, totalSupply, nil

}

/*


 */
