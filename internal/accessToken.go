package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func CheckToken()(string, error) {
	var tokenDir, AppData string
	operatingSys := runtime.GOOS
	if operatingSys == "linux" {
		AppData = os.Getenv("HOME")
		tokenDir = filepath.Join(AppData, ".config", "armoracrypt")
	} else {
		AppData = os.Getenv("LOCALAPPDATA")
		tokenDir = filepath.Join(AppData, "armoracrypt")
	}

	var token string
	_, err := os.Stat(tokenDir+"\\token.bin")
	if err == nil {
		byteofToken,_ := os.ReadFile(tokenDir+"\\token.bin")
		token = string(byteofToken)
		fmt.Println("Auth-Key [DETECTED]")
		return token,nil
	} else {
		fmt.Println("Auth-Key [MISSING]")
	}
	fmt.Printf("Enter Your DropBox Auth Key:")
	fmt.Scan(&token)

	err = os.WriteFile(tokenDir+"\\token.bin",[]byte(token),0755)
	if err!=nil{
		return "",err
	}
	fmt.Println("TOKEN STORED!")
	return token,nil

}
