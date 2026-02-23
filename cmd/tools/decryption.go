package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var decrypt = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts the Files",
	Long:  "AES256-GCM based encyption that decrypts the Files",
	Run: func(cmd *cobra.Command, args []string) {
		fp, err := cmd.Flags().GetString("fp")
		if err != nil {
			fmt.Printf("Error getting fp flag for decryption!")
			return 
		}
		AbsPath, err := filepath.Abs(fp)
		if filepath.Ext(AbsPath) != ".crypt"{
			fmt.Println("Not An Encrypted File can't Decrypt!")
			return
		}
		if err != nil {
			fmt.Println("Error Getting Abs Path!")
			return
		}
		fileInfo, err := os.Stat(AbsPath)
		if err != nil {
			fmt.Println("Not a valid file path for Decryption!")
			return
		}
		data, err := internal.Decrypt(AbsPath)
		if err!=nil{
			return
		}

		fileName := fileInfo.Name()
		if filepath.Ext(fileName) == ".crypt" {
			fileName = fileName[:len(fileName)-len(".crypt")]
		}

		dirPath := filepath.Dir(AbsPath)
		os.WriteFile(filepath.Join(dirPath,fileName),data,0600)
		fmt.Println("Decrypted Successfyully.")
		fmt.Println("Location:",filepath.Dir(AbsPath))
		
		//check if its a Folder 
		isZip := (filepath.Ext(filepath.Join(dirPath,fileName))==".zip")
		if isZip{ //is a zipped Folder

			fmt.Println("Zip [Detected]")
			fmt.Printf("Unzipping.")
			time.Sleep(300*time.Millisecond)
			fmt.Printf("\rUnzipping..")
			time.Sleep(300*time.Millisecond)
			fmt.Printf("\rUnzipping...\n")
			time.Sleep(300*time.Millisecond)

			err = internal.Unzip(filepath.Join(dirPath,fileName),dirPath)
			if err!=nil{
				fmt.Println("Error unzipping!")
				return
			}
			err = os.Remove(filepath.Join(dirPath,fileName)) 
			if err!=nil{
				fmt.Println("Error deleting the zip file")
				return
			}
	

		}
		
	},
}

func init() {
	decrypt.Flags().String("fp", "", "filePath for decryption")
	cmd.RootCmd.AddCommand(decrypt)
}
