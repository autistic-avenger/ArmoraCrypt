package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func Zip(dirpath string) error {
	zipDirName := filepath.Base(dirpath)
	outPutDir := dirpath[:len(dirpath)-len(zipDirName)-1]
	fmt.Printf("Creating a zip.")
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("\rCreating a zip..")
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("\rCreating a zip...\n")
	time.Sleep(100 * time.Millisecond)

	zippedFileHandler, err := os.Create(filepath.Join(outPutDir, zipDirName+".zip"))
	if err != nil {
		return err
	}
	defer zippedFileHandler.Close()

	zipWriter := zip.NewWriter(zippedFileHandler)
	defer zipWriter.Close()

	filepath.Walk(dirpath, func(path string, info fs.FileInfo, err error) error {
		relpath, err := filepath.Rel(dirpath, path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			zipWriter.Create(relpath + "/")
			return nil
		}

		fileData, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileData.Close()
		zipFileData,err :=zipWriter.Create(relpath)
		if err !=nil{
			return nil
		}
		_,err = io.Copy(zipFileData,fileData)
		if err!=nil{
			return nil
		}
		fmt.Printf("\r\033[0KWriting %v",info.Name())
		
		return nil
	})

	return nil

}
