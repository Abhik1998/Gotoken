package accounty

import (
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"
)

func GenerateAccount(keystoreDir string, pwd string) (common.Address, error) {
    ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
    acc, err := ks.NewAccount(pwd)
    if err != nil {
        return common.Address{}, err
    }

    return acc.Address, nil
}

