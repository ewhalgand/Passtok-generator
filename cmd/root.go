/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ptk",
	Short: "Password or token generator",
	Long: `passtok is a simple and flexible CLI application to generate secure passwords or API tokens directly from your terminal.

	You can customize the type and length of the generated string using flags.

	Examples:
		ptk --version
		ptk gen
		ptk gen --type password --length 16
		ptk gen --type password --length 16 --copy
		ptk gen -t token -l 3

	Useful for developers, sysadmins, or anyone needing secure credentials on the fly.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("ptk version %v\n", version)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passtok.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(generateCmd)
	rootCmd.Flags().BoolP("version", "v", false, "Package version")
}
