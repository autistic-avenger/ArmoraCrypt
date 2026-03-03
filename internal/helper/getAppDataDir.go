package helper

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataDir() string{
	var tokenDir, AppData string
	operatingSys := runtime.GOOS
	if operatingSys == "linux" {
		AppData = os.Getenv("HOME")
		tokenDir = filepath.Join(AppData, ".config", "armoracrypt")
	} else {
		AppData = os.Getenv("LOCALAPPDATA")
		tokenDir = filepath.Join(AppData, "armoracrypt")
	}
	return tokenDir
}