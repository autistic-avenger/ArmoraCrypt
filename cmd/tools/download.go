package tools

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Download = &cobra.Command{
	Use: "download",
	Short: "Download and decrypt the given files.",
	Long:"Downloads and decrypts the files locally using AES-256 functions",
	Run: func(cmd *cobra.Command, args []string) {
		var downloadType int8
		fmt.Println("------DOWNLOAD-----")
		fmt.Println("1.Folders")
		fmt.Println("2.Files")
		fmt.Scan(&downloadType)

		if downloadType == 1{


		}else if downloadType == 2{

		}else{
			fmt.Println("Not Allowed")
		}
	},
}