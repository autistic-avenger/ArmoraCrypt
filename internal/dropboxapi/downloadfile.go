package dropboxapi

import (
	"armoracrypt/internal"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func DownloadFile(token string, dropBoxPath string) ([]byte,error) {
	req, _ := http.NewRequest(
		"POST",
		"https://content.dropboxapi.com/2/files/download",
		nil,
	)

	client := &http.Client{}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Dropbox-API-Arg", fmt.Sprintf("{\"path\":\"%s\"}", dropBoxPath))

	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	fmt.Println("Status:", resp.Status)

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
		token,_:=internal.CheckToken()
		DownloadFile(token,dropBoxPath)
	}else if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("Failed.")
	}

	data,_ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data,nil
}
