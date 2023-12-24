package main

import (
	"os"
	"runtime"
)

// UserHomeDir https://github.com/spf13/viper/blob/e36638d8786b0b58231039fc6d7db32b904dd1ba/util.go#L140
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
