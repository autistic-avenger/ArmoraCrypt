package internal

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(zipfilePath string,unzipOut string) error {
	unzipper,err := zip.OpenReader(zipfilePath)
	if err!=nil{
		return err
	}
	defer unzipper.Close()

	for _,zipPath := range unzipper.File{
		isDir := zipPath.FileInfo().IsDir()
		if isDir{
			//create Dir of all zipped things 
			err = os.MkdirAll(filepath.Clean(filepath.Join(unzipOut,filepath.Base(zipfilePath)[:len(filepath.Base(zipfilePath))-4],zipPath.Name)),0600)
			if err!=nil{
				return err
			}
		}
		if !isDir{

			outputFileData,err:= os.Create(filepath.Join(unzipOut,filepath.Base(zipfilePath)[:len(filepath.Base(zipfilePath))-4],zipPath.Name))
			if err!=nil{
				return err
			}
	
			zippedData,err := zipPath.Open()
			if err!=nil{
				return err
			}
			defer zippedData.Close()
			io.Copy(outputFileData,zippedData)
		}
	}

	return nil
}