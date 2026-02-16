package tools

import (
	"armoracrypt/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

var encrypt = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts the Files",
	Long:  "AES256-GCM based encyption that encrypts the Files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("&@&#@&$&*@&**%*@#")
	},
}

func init() {
	cmd.RootCmd.AddCommand(encrypt)
}
