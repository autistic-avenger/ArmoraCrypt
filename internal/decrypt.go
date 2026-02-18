package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func Decrypt(fp string) ([]byte, error) {
	//get key
	AppData := os.Getenv("LOCALAPPDATA")
	if AppData == "" {
		fmt.Println("Error Getting Appdata dir.")
		return nil, fmt.Errorf("Error getting AppData")
	}
	keyDir := filepath.Join(AppData, "armoracrypt", "Keys")
	key, err := os.ReadFile(keyDir + "/masterkey.bin")
	if err != nil {
		return nil, err
	}
	//------------------------

	encryptedData, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println("Error Reading file for Decryption!")
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error genering Cipher")
		return nil, err
	}

	gcm, _ := cipher.NewGCM(block)

	nonceSize := gcm.NonceSize()
	nonce, encryptedData := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(encryptedData), nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
