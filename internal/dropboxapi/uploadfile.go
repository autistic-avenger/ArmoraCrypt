package dropboxapi

import (
	"armoracrypt/internal"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func UploadFile(token string, dropboxPath string, localFilePath string) error {
	fileData, err := os.ReadFile(localFilePath)

	reqs, err := http.NewRequest(
		"POST",
		"https://content.dropboxapi.com/2/files/upload",
		bytes.NewReader(fileData),
	)
	if err != nil {
		return err
	}

	reqs.Header.Set("Authorization", "Bearer "+token)
	reqs.Header.Set("Content-Type", "application/octet-stream")
	reqs.Header.Set("Dropbox-API-Arg", fmt.Sprintf("{\"autorename\":false,\"mode\":\"overwrite\",\"mute\":false,\"path\":\"%s\",\"strict_conflict\":false}", dropboxPath))

	client := &http.Client{}
	resp, err := client.Do(reqs)
	if err != nil {
		return err
	}
	fmt.Println("Status Code:",resp.StatusCode)
	if resp.StatusCode == 401{
		fmt.Println("Your auth key expired..")
		//reget token 
		var tokenDir, AppData string
		operatingSys := runtime.GOOS
		if operatingSys == "linux" {
			AppData = os.Getenv("HOME")
			tokenDir = filepath.Join(AppData, ".config", "armoracrypt")
		} else {
			AppData = os.Getenv("LOCALAPPDATA")
			tokenDir = filepath.Join(AppData, "armoracrypt")
		}

		os.Remove(tokenDir+"/token.bin")
		token,_ := internal.CheckToken()
		UploadFile(token,dropboxPath,localFilePath)

	}else if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("Failed.")
	}
	defer resp.Body.Close()

	return nil
}
