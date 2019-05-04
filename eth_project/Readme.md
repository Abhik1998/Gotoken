
**Description**

 ***Tokeninteraction.sol***
 This smart contract is developed to take a note of funds transferred between the application and the Ethereum
 
 ```function balanceOf( address who ) constant returns (uint value) ```
 
 this function return the balance value in the contract
 Since we didnot specifically mention any value in the initial contructor so I just added a config/contract.js which   initializes the number of tokens initially present with the Ethereum as 1000.
 The #app# folder consists of the golang app which ,as of now, can create accounts and keep a track of those in the folder accounty as I have generated a sample account
 
 [Address link](https://github.com/Abhik1998/Gotoken/blob/master/eth_project/app/accounty/UTC--2019-05-03T15-52-14.230412907Z--46fffc00a18e3ffe9482af080091e04f44ed8a21)
 This consists of a single line data as:
```
{"address":"46fffc00a18e3ffe9482af080091e04f44ed8a21","crypto":{"cipher":"aes-128-ctr","ciphertext":"a80e844f688614efda6ef08570973e1d29a9a62a29794958d2fa0f56d505dff3","cipherparams":{"iv":"f5490a2439e1be633409f9c9077b3661"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"024021aa74d6eb1f612df85949cf78c4c76f5f82e6bc88cb56f81e94d153759a"},"mac":"76ce61872f8fb29a216ddceb13dc631108d05a1d83f367f8127a38eb3e7e1c40"},"id":"10b78cf2-bbf9-428d-9117-4f6bf284c13d","version":3}
```
Next steps:
To use the tokens and maintain track of all the tokens used by the app.

Points covered:
1) The smart contract keeps a track of the issued tokens and can return balance after a transaction ,i.e, of the
   number of ethers in the Ethereum account.(Tokeninteraction.sol)
2) The token exchange takes place in O(1) steps as after reading the code one can see that each method consists of
   2 steps of operation. No looping statements have been used.
3) Since the smart contract is using trnsactions taken on behalf of ethereum and returns balance of tokens present
   with Ethereum so security of Ethereum is assured.
4) A centralised appliction built with Golang which keeps a track of the account to be used for tokens transfer with 
   Ethereum.

   
