package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"

	"github.com/spf13/cobra"
)

var masterKeyInit = &cobra.Command{
	Use:   "init",
	Short: "Manages the application's master encryption key",
	Long:  `The masterkey command provides operations for generating,master encryption key`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Keygen()
	},
}
func init(){
	cmd.RootCmd.AddCommand(masterKeyInit)
}
