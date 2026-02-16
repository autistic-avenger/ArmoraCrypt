package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "armoracrypt",
	Short: "Encrypt & Decrypt Files From CLI!",
	Long:  "Uses AES-256 to do safe and secure encryption/decryption of the files directly from the CLI!",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
