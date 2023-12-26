package util

import (
	"github.com/fatih/color"
	"os"
	"runtime"
)

// UserHomeDir https://github.com/spf13/viper/blob/e36638d8786b0b58231039fc6d7db32b904dd1ba/util.go#L140
func UserHomeDir() string {
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// The color package also disables color output if the NO_COLOR environment variable is set to a non-empty string.
// https://github.com/fatih/color?tab=readme-ov-file#disableenable-colorx

func LogErrorAndExit(err error) {
	c := color.New(color.FgRed).Add(color.Bold)
	_, _ = c.Printf("Error: %v\n", err)
	os.Exit(1)
}
