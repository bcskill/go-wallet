package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEosGenerate(t *testing.T) {
	wallet := EosWallet{}
	account := wallet.Generate(false)
	//{"private_key":"5KDfdGiWXXTxb2QHVAZp4vMXTsbiNP39HrBtiz8rY9Ag3LnFoUM",
	//"public_key":"EOS7KrBYdccTUYxExpJiXDp2YXMjMxJHepKade93gxhuHw5PTmx9k",
	//"address":"",
	//"mnemonic":"soccer abuse buyer upset calm pass extra camp visa man economy elephant",
	//"keystore":""}
	fmt.Println(account)
}

func TestEosGenerateByPrivateKey(t *testing.T) {
	wallet := EosWallet{}
	privateKey := "5KDfdGiWXXTxb2QHVAZp4vMXTsbiNP39HrBtiz8rY9Ag3LnFoUM"
	account := wallet.GenerateByPrivateKey(privateKey, false)
	fmt.Println(account)
	publicKey := "EOS7KrBYdccTUYxExpJiXDp2YXMjMxJHepKade93gxhuHw5PTmx9k"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
}

func TestEosGenerateByMnemonic(t *testing.T) {
	wallet := EosWallet{}
	mnemonic := "soccer abuse buyer upset calm pass extra camp visa man economy elephant"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/0'/0'/0/0", false)
	fmt.Println(account)
	publicKey := "EOS7KrBYdccTUYxExpJiXDp2YXMjMxJHepKade93gxhuHw5PTmx9k"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
}