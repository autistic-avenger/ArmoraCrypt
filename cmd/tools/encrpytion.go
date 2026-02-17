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
		if err != nil {
			fmt.Println("Error fetching fp flag!")
		}
		_, err = os.Stat(fp)
		if err != nil {
			fmt.Println("Not a valid file path!")
			return
		}
		AbsPath, err := filepath.Abs(fp)
		if err != nil {
			fmt.Println("Error Getting Abs Path!")
			return
		}
		cypher ,err := internal.Encrypt(AbsPath)
		if err != nil {
			fmt.Println(AbsPath)
			fmt.Println("Error encrypting File!")
			return
		}
		absWriteFilePath,err := filepath.Abs("./!OPERATIONS")
		if err!=nil{
			fmt.Println("Error abs file")
		}
		fmt.Println(absWriteFilePath)
		joinedFP := filepath.Join(absWriteFilePath,"Encrypted")
		err = os.WriteFile(joinedFP+".crypt",cypher,0600)
	},
}

func init() {
	encrypt.Flags().String("fp", "NotAllowed", "filepath of file to encrypt")
	encrypt.MarkFlagRequired("fp")
	cmd.RootCmd.AddCommand(encrypt)
}
