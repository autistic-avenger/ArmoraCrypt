package tools

import (
	"armoracrypt/cmd"
	"fmt"

	// "os"
	// "path/filepath"
	"github.com/spf13/cobra"
)

var encrypt = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts the Files",
	Long:  "AES256-GCM based encyption that encrypts the Files",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := cmd.Flags().GetString("fp")
		if err != nil {
			fmt.Println("Error fetching fp flag!")
		}

		// fileInfo, err := os.Stat(fp)
		// if err != nil {
		// 	fmt.Println("Not a valid file path!")
		// 	return
		// }

	},
}

func init() {
	encrypt.Flags().String("fp", "NotAllowed", "filepath of file to encrypt")
	encrypt.MarkFlagRequired("fp")
	cmd.RootCmd.AddCommand(encrypt)
}
