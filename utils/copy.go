package utils

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(value string) {
	err := clipboard.WriteAll(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, "⚠️ Error copying to clipboard : ", err)
	} else {
		fmt.Println("📋 Copied to clipboard successfully !")
	}
}
