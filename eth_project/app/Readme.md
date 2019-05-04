All about app

**APP SAMPLE**

Since, it may be difficult for the reviewers to check out through the app. I am attaching a sample image of use of the app.

![Example](https://github.com/Abhik1998/Gotoken/blob/master/eth_project/app/WhatsApp%20Image%202019-05-03%20at%2021.24.18.jpeg)

In this example, I have created an account in the application and it connects to Ethereum address ,hence token exchange can take place in between Ethereum and the application in Golang. 

App has been extended:
Exchange of tokens can take place with Ethereum and tokens can be locked in the app.

Code Architechture:

[Line 56](https://github.com/Abhik1998/Gotoken/blob/38947a6ee8e0ecfa474ad0a6fe9353575157dd91/eth_project/app/main.go#L56)
 Till this line we have the code to add an account connecting to Ethereum for transferring tokens from and to centralised
   application.
   I recommend first of all to install all of the go packages in the `GOPATH`
   ```
   go get -u all
   ```
   This is because we are using `crypto/sha3` for ethereum interaction.
   After the lines 56 we have:
 ```
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
```
All of these refer to address from which tokens are to be transferred to the receiver address.

 I have provided a gas limit at the last few lines of the `main()` function.
 
 The last method provides option to user to set his address and paraphrase for security.
 
