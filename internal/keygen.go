package internal

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
)

func Keygen() {
	AppData := os.Getenv("LOCALAPPDATA")
	if AppData == "" {
		fmt.Println("Error Getting Appdata dir.")
		return
	}
	keyDir := filepath.Join(AppData, "armoracrypt", "Keys")

	//check if Master key already exist
	_, err := os.Stat(keyDir)
	if err == nil {
		fmt.Println("Master Key [DETECTED]")
		return
	} else {
		fmt.Println("Master Key [MISSING]")
	}

	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		fmt.Println("Error generating private keys!")
		return
	}
	fmt.Println("Generating Master Key...")
	//create root Masterkey file
	if err := os.MkdirAll(keyDir, 0600); err != nil {
		fmt.Printf("error creating key directory: %v", err)
		return
	}

	if err := os.WriteFile(keyDir+"/masterkey.bin", key, 0600); err != nil {
		fmt.Printf("error writing key file: %v", err)
		return
	}
	fmt.Println("Master key generated at:", keyDir)
}
