package internal

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Keygen() {
	var keyDir, AppData string
	if runtime.GOOS == "linux" {
		AppData = os.Getenv("HOME")
		keyDir = filepath.Join(AppData, ".config", "armoracrypt", "Keys")
	} else {
		AppData = os.Getenv("LOCALAPPDATA")
		keyDir = filepath.Join(AppData, "armoracrypt", "Keys")
	}
	if AppData == "" {
		fmt.Println("Error Getting Appdata dir.")
		return
	}

	//check if Master key already exist
	_, err := os.Stat(keyDir)
	if err == nil {
		fmt.Println("Master Key [DETECTED]")
		fmt.Println("Location:", keyDir)
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
	if err := os.MkdirAll(keyDir, 0755); err != nil {
		fmt.Printf("error creating key directory: %v", err)
		return
	}

	if err := os.WriteFile(keyDir+"/masterkey.bin", key, 0755); err != nil {
		fmt.Printf("error writing key file: %v", err)
		return
	}
	fmt.Println("Master key generated at:", keyDir)
}
