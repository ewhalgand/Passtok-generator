/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"passtok/utils"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var length int
var typ string
var copyFlag bool

func generatePassword(l int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for range make([]struct{}, l) {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

func generateToken(l int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for range make([]struct{}, l) {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a random password or token",
	Long: `Generate a secure password or token with customizable options.

	Examples:
		ptk gen                   # Generate a default 16-character password
		ptk gen --type (password or token) --length (number of character)     # Generate a number of characters based on the type
		pkt gen --type (password or token) --length (number of character) --copy # Generate a number of characters based on the type and copy it

	Supported types:
		--length Number of characters in the generated string (default 16)
		--type  Type of string to generate: "password" or "token" (default "password")
		--copy Copy the generated string to clipboard
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("🔐 Type demandé : %s (%d caractères)\n", typ, length)
		switch typ {
		case "password":
			result := generatePassword(length)
			fmt.Println("🔑 Your password : ", result)

			if copyFlag {
				utils.CopyToClipboard(result)
			}
		case "token":
			result := generateToken(length)
			fmt.Println("🔑 Your token : ", result)

			if copyFlag {
				utils.CopyToClipboard(result)
			}
		default:
			fmt.Fprintln(os.Stderr, "❌ Type not found. Use --type password or --type token")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().IntVarP(&length, "length", "l", 16, "Password or token length")
	generateCmd.Flags().StringVarP(&typ, "type", "t", "password", "Generated type : password or token")
	generateCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "Copy result to clipboard")
}
