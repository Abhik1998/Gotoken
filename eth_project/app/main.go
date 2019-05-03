package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"app/accounty"
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



