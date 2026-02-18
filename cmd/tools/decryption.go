package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"fmt"
	"os"
	"path/filepath"
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
		}
		AbsPath, err := filepath.Abs(fp)
		if err != nil {
			fmt.Println("Error Getting Abs Path!")
			return
		}
		fileInfo, err := os.Stat(fp)
		if err != nil {
			fmt.Println("Not a valid file path for Decryption!")
			return
		}
		data, err := internal.Decrypt(AbsPath)

		fileName := fileInfo.Name()
		if filepath.Ext(fileName) == ".crypt" {
			fileName = fileName[:len(fileName)-len(".crypt")]
		}

		dirPath := filepath.Dir(AbsPath)
		os.WriteFile(filepath.Join(dirPath,fileName),data,0600)
		fmt.Println("Decrypted File Successfyully.")
		fmt.Println("Location:",filepath.Dir(AbsPath))
	},
}

func init() {
	decrypt.Flags().String("fp", "", "filePath for decryption")
	cmd.RootCmd.AddCommand(decrypt)
}
