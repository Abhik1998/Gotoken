package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
    "app/accounty"
    "context"
    "log"
    "math/big"
    "crypto/ecdsa"
    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
    "github.com/ethereum/go-ethereum/ethclient"
)

const choiceCreateAccount = 1

func main() {
    fmt.Println("Welcome to Go app, use it!")

    fmt.Println("What action would you like to perform?")
    fmt.Printf("    Type: '%d' to Create new account.\n", choiceCreateAccount)

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        cmd, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println(err)
            os.Exit(0)
        }

        switch cmd {
		case choiceCreateAccount:
			keystoreDir, passphrase, err := scanCreateAccountArgs()
            if err != nil {
            fmt.Println(err)
            os.Exit(1)
            }
			acc, err := accounty.GenerateAccount(keystoreDir, passphrase)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		
			fmt.Printf("New account '%s' successfully generated and stored inside the '%s' directory.\n", acc.Hex(), keystoreDir)
		}
		

        fmt.Println("Done.")
        os.Exit(0)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("...")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("...")
	tokenAddress := common.HexToAddress("...")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("Token amount: %s", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gas limit: %d", gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

fmt.Printf("Tokens sent at TX: %s", signedTx.Hash().Hex())
}
func scanCreateAccountArgs() (keystoreDir string, passphrase string, err error) {
    keystoreDir = ""
    passphrase = ""

    fmt.Println("Paste the keystoreDir path where your new account should be saved:")
    _, err = fmt.Scanf("%s",&keystoreDir)
    if err != nil {
        return
    }

    fmt.Println("Type the passphrase you want to encrypt your account with:")
    _, err = fmt.Scanf("%s",&passphrase)
    if err != nil {
        return
    }

    return
}



