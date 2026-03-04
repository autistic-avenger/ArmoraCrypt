package helper

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDownloadDir() string {
	oper := runtime.GOOS
	if oper == "linux" {
		return filepath.Join(os.Getenv("HOME"), "Downloads")
	} else {
		return filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	}
}
