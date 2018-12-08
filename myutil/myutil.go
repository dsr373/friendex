package myutil

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// CheckErr is a simple function that checks whether an error has occurred, and if so
// quits the program and logs the error message together with the context-provided "message"
func CheckErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s : %s", message, err)
	}
}

// ConfigDir returns the configuration folder, i.e. the current user's $HOME/.config/friendex
func ConfigDir() string {
	var home string
	var relativeConfDir = filepath.FromSlash(".config/friendex")

	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}
	return filepath.Join(home, relativeConfDir)
}
