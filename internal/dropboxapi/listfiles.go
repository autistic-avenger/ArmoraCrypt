package dropboxapi

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func UploadFile(token string, dropboxPath string, localFilePath string) error {
	fileData, err := os.ReadFile(localFilePath)

	reqs, err := http.NewRequest(
		"POST", 
		"https://content.dropboxapi.com/2/files/upload", 
		bytes.NewReader(fileData),
	)
	if err!=nil{
		return err
	}

	reqs.Header.Set("Authorization","Bearer "+token)
	reqs.Header.Set("Content-Type","application/octet-stream")
	reqs.Header.Set("Dropbox-API-Arg", fmt.Sprintf("{\"autorename\":false,\"mode\":\"add\",\"mute\":false,\"path\":\"%s\",\"strict_conflict\":false}",dropboxPath))

	client := &http.Client{}
	resp,err:= client.Do(reqs)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	return nil
}
