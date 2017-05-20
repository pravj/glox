// Package cmd implements the interface to handle command-line arguments.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "glox",
	Short: "Interpreter for the Lox language in Go",
	Long:  `Interpreter for the Lox language in Go`,
}

// Execute handles the execution of the base command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
