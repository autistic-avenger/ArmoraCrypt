package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"

)

func Decrypt() error {
	//get key
	AppData := os.Getenv("LOCALAPPDATA")
	if AppData == "" {
		fmt.Println("Error Getting Appdata dir.")
		return fmt.Errorf("Error getting AppData")
	}
	keyDir := filepath.Join(AppData, "armoracrypt", "Keys")
	key, err := os.ReadFile(keyDir + "/masterkey.bin")
	if err != nil {
		return err
	}
	//------------------------

	fp,err := filepath.Abs("./!OPERATIONS")
	if err!=nil{
		fmt.Println("Error getting ABS fp for Decryption!")
		return err
	}
	encryptedData,err := os.ReadFile(filepath.Join(fp,"Encrypted.crypt"))
	if err!=nil{
		fmt.Println("Error Reading file for Decryption!")
		return err
	}

	block ,err := aes.NewCipher(key)
	if err!=nil{
		fmt.Println("Error genering Cipher")
		return err
	}

	gcm,_ := cipher.NewGCM(block)

	nonceSize := gcm.NonceSize()
	nonce, encryptedData := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(encryptedData), nil)
	os.WriteFile(fp+"/Decrypted.png",plaintext,0600)
	return nil
}