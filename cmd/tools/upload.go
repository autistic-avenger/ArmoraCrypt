package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	dropboxapi "armoracrypt/internal/dropboxapi"
	"armoracrypt/internal/helper"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var Upload = &cobra.Command{
	Use:   "upload",
	Short: "safe upload to cloud",
	Long:  "Encrypts and Uploads file/folder to the Cloud Storage",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Keygen()
		fp, err := cmd.Flags().GetString("fp")
		if err != nil {
			fmt.Println("Error getting fp Flag")
		}
		Abs, _ := filepath.Abs(fp)
		info, err := os.Stat(Abs)
		if err != nil {
			fmt.Println("Incorrect PATH")
			return
		}
		if info.IsDir() {
			_, err := exec.Command("armoracrypt", "encrypt", "--d", Abs).Output()
			if err != nil {
				fmt.Println("Error Uploading File")
				return
			}
			token, err := internal.CheckToken()
			if err != nil {
				fmt.Println(err)
				return
			}
			fileInfo, _ := os.Stat(Abs)
			//aPi works:
			//token ,relativeDROPPATH,localPATH

			err = dropboxapi.UploadFile(token, "/ARMORA/FOLDERS/"+fileInfo.Name()+".zip.crypt", Abs+".zip.crypt")
			if err != nil {
				fmt.Println("Error Uploading...", err)
				return
			}

			appDataDir := helper.GetAppDataDir()

			os.MkdirAll(filepath.Join(appDataDir, "METADATA"), 0755)
			fileDir := filepath.Join(appDataDir, "METADATA", "Folders.txt")
			readfile,_:= os.ReadFile(fileDir)
			if !strings.Contains(string(readfile),fileInfo.Name()){

				metadata, _ := os.OpenFile(fileDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0755)
				metadata.Write([]byte(fileInfo.Name()+"\n"))
				defer metadata.Close()
			}

			fmt.Println("[UPLOADED SUCCESSFULLY]")
			defer os.Remove(Abs + ".zip.crypt")

		} else {
			_, err := exec.Command("armoracrypt", "encrypt", "--fp", Abs).Output()
			if err != nil {
				fmt.Println("Error Uploading File")
				return
			}

			token, err := internal.CheckToken()
			if err != nil {
				fmt.Println(err)
				return
			}
			fileInfo, _ := os.Stat(Abs)

			//aPi works:
			//token ,relativeDROPPATH,localPATH
			err = dropboxapi.UploadFile(token, "/ARMORA/FILES/"+fileInfo.Name()+".crypt", Abs+".crypt")
			if err != nil {
				fmt.Println("Error Uploading...")
				return
			}

			
			appDataDir := helper.GetAppDataDir()

			os.MkdirAll(filepath.Join(appDataDir, "METADATA"), 0755)
			fileDir := filepath.Join(appDataDir, "METADATA", "Files.txt")
			readfile,_:= os.ReadFile(fileDir)
			if !strings.Contains(string(readfile),fileInfo.Name()){
				metadata, _ := os.OpenFile(fileDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0755)
				metadata.Write([]byte(fileInfo.Name()+"\n"))
				defer metadata.Close()

			}

			fmt.Println("[UPLOADED SUCCESSFULLY]")
			defer os.Remove(Abs + ".crypt")
		}

	},
}

func init() {
	Upload.Flags().String("fp", "/NarendraMODI/ISRAEL", "Upload to cloud")
	cmd.RootCmd.AddCommand(Upload)
}
