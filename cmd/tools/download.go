package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"armoracrypt/internal/dropboxapi"
	"armoracrypt/internal/helper"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var Download = &cobra.Command{
	Use: "download",
	Short: "Download and decrypt the given files.",
	Long:"Downloads and decrypts the files locally using AES-256 functions",
	Run: func(cmd *cobra.Command, args []string) {
		token ,_ := internal.CheckToken()
		var downloadType int8
		fmt.Println("------DOWNLOAD-----")
		fmt.Println("1.Folders")
		fmt.Printf("2.Files\n>> ")
		fmt.Scan(&downloadType)

		if downloadType == 1{
			AppData := helper.GetAppDataDir()
			folderTxtDir := filepath.Join(AppData,"METADATA","Folders.txt")
			data,err := os.ReadFile(folderTxtDir)
			if err!=nil{
				fmt.Println("ERR:Nothing to Download?")
				return 
			}
			folders := strings.Split(string(data),"\n")
			fmt.Println("\n-----WHICH FOLDER TO FETCH-----")
			for index,folder := range folders{
				if folder == ""{
					continue
				}
				fmt.Printf("%d:%s\n",index+1,folder)
			}	
			fmt.Printf(">> ")
			var realIndex int8 
			fmt.Scan(&realIndex)
			dropboxPath := "/ARMORA/FOLDERS/"+folders[realIndex-1]+".zip.crypt"
			fileData,err := dropboxapi.DownloadFile(token,dropboxPath)
			if err!=nil{
				fmt.Println("Error Downloading File:",err)
				return
			}
			downloadDir := filepath.Join(helper.GetDownloadDir(),folders[realIndex-1]+".zip.crypt")

			os.WriteFile(downloadDir,fileData,0755)
			_,err =exec.Command("armoracrypt","decrypt","--fp",downloadDir).Output()
			if err!=nil{
				fmt.Println("ERR:Decrypting Files")
				return 
			}
			fmt.Println("Downloading Complete...\n[DOWNLOAD FOLDER]")
			err = os.Remove(downloadDir)
			if err!=nil{
				fmt.Println("REMOVING FAILED.")
			}
		}else if downloadType == 2{
			AppData := helper.GetAppDataDir()
			fileTxtDir := filepath.Join(AppData,"METADATA","Files.txt")
			data,err := os.ReadFile(fileTxtDir)
			if err!=nil{
				fmt.Println("ERR:Nothing to Download?")
				return 
			}

			files := strings.Split(string(data),"\n")
			fmt.Println("\n-----WHICH FOLDER TO FETCH-----")
			for index,folder := range files{
				if folder == ""{
					continue
				}
				fmt.Printf("%d:%s\n",index+1,folder)
			}	
			fmt.Printf(">> ")
			var realIndex int8 
			fmt.Scan(&realIndex)
			dropboxPath := "/ARMORA/FILES/"+files[realIndex-1]+".crypt"
			fileData,err := dropboxapi.DownloadFile(token,dropboxPath)
			if err!=nil{
				fmt.Println("Error Downloading File:",err)
				return
			}
			downloadDir := filepath.Join(helper.GetDownloadDir(),files[realIndex-1]+".crypt")

			os.WriteFile(downloadDir,fileData,0755)
			_,err =exec.Command("armoracrypt","decrypt","--fp",downloadDir).Output()
			if err!=nil{
				fmt.Println("ERR:Decrypting Files")
				return 
			}
			fmt.Println("Downloading Complete...\n[DOWNLOAD FOLDER]")
			err = os.Remove(downloadDir)
			if err!=nil{
				fmt.Println("REMOVING FAILED.")
			}

		}else{
			fmt.Println("Not Allowed")
		}
	},
}

func init(){
	cmd.RootCmd.AddCommand(Download)
}