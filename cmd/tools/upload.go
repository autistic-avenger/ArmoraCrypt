package tools

import (
	"armoracrypt/cmd"
	"armoracrypt/internal"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)
var Upload = &cobra.Command{
	Use: "upload",
	Short: "safe upload to cloud",
	Long: "Encrypts and Uploads file/folder to the Cloud Storage",
	Run:func(cmd *cobra.Command, args []string) {
		internal.Keygen()
		fp,err := cmd.Flags().GetString("fp")
		if err!=nil{
			fmt.Println("Error getting fp Flag")
		}
		Abs,_:= filepath.Abs(fp)
		info ,err := os.Stat(Abs)
		if err!=nil{
			fmt.Println("Incorrect PATH")
			return
		}
		if info.IsDir(){
			_,err := exec.Command("armoracrypt", "encrypt", "--d", Abs).Output()
			if err!=nil{
				fmt.Println("Error Uploading File")
				return
			}
			token,err := internal.CheckToken()
			if err!=nil{
				fmt.Println(err)
				return
			}
			
			fmt.Println("[UPLOADED SUCCESSFULLY]")
			defer os.Remove(Abs+".zip.crypt")

		}else{
			_,err := exec.Command("armoracrypt","encrypt","--fp",Abs).Output()
			if err!=nil{
				fmt.Println("Error Uploading File")
				return
			}
			
			fmt.Println("[UPLOADED SUCCESSFULLY]")
			defer os.Remove(Abs+".crypt")
		}




	},
} 

func init(){
	Upload.Flags().String("fp","","Upload to cloud")
	cmd.RootCmd.AddCommand(Upload)
}