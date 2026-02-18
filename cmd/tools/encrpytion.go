package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var encrypt = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts the Files",
	Long:  "AES256-GCM based encyption that encrypts the Files",
	Run: func(cmd *cobra.Command, args []string) {
		fp, err := cmd.Flags().GetString("fp")

		//if its a file encryption
		if fp != "" {
			if err != nil {
				fmt.Println("Error fetching fp flag!")
			}
			fileInfo, err := os.Stat(fp)
			if err != nil {
				fmt.Println("Not a valid file path!")
				return
			}
			AbsPath, err := filepath.Abs(fp)
			if err != nil {
				fmt.Println("Error Getting Abs Path!")
				return
			}
			absWriteFilePath := filepath.Dir(AbsPath) 
			cypher, err := internal.Encrypt(AbsPath)
			if err != nil {
				fmt.Println("Error encrypting File!")
				return
			}
			
			fmt.Println(absWriteFilePath)
			joinedFP := filepath.Join(absWriteFilePath, fileInfo.Name())
			err = os.WriteFile(joinedFP+".crypt", cypher, 0600)

		} else {
			//if its a dir encryption
			dir, err := cmd.Flags().GetString("d")
			if dir == "" {
				return
			}
			if err != nil {
				fmt.Println("Error getting dir flag!")
			}
			_, err = os.ReadDir(dir)
			if err != nil {
				fmt.Println("Not a valid DIR path!")
				return
			}
			AbsDir, _ := filepath.Abs(dir)
			if err != nil {
				fmt.Println("Error Getting Dir Abs Path!")
				return
			}
			fmt.Println("DIR:", AbsDir)

		}

	},
}

func init() {
	encrypt.Flags().String("fp", "", "filepath of file to encrypt")
	encrypt.Flags().String("d", "", "folder path")
	cmd.RootCmd.AddCommand(encrypt)
}
