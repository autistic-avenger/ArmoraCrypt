package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"github.com/spf13/cobra"
)

var decrypt = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts the Files",
	Long:  "AES256-GCM based encyption that decrypts the Files",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Decrypt()
	},
}

func init() {
	cmd.RootCmd.AddCommand(decrypt)
}
