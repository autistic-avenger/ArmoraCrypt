package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
)

func Encrypt(fp string)([]byte, error) {
	//get key
	AppData := os.Getenv("LOCALAPPDATA")
	if AppData == "" {
		fmt.Println("Error Getting Appdata dir.")
		return nil,fmt.Errorf("Error getting AppData")
	}
	keyDir := filepath.Join(AppData, "armoracrypt", "Keys")
	key, err := os.ReadFile(keyDir + "/masterkey.bin")
	if err != nil {
		return nil,err
	}
	//------------------------

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil,err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}

	gcm, err := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil,err
	}

	cypher := gcm.Seal(nonce,nonce,data,nil)
	return cypher,nil

}
