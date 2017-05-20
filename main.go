// Package main implements the public facing interface to the glox interpreter.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pravj/glox/cmd"
	"github.com/pravj/glox/scanner"
	"github.com/pravj/glox/utils"
	"github.com/spf13/cobra"
)

// glox represents the interpreter.
type glox struct {
	// hasError, flag for error handling.
	hasError bool
}

// createInterpreter returns a new glox instance.
func createInterpreter() *glox {
	return &glox{hasError: false}
}

// runLoop implements the REPL interface for glox.
func (i *glox) runLoop() {
	promptPrefix := ">>> "

	replScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(promptPrefix)

		scanned := replScanner.Scan()
		if !scanned {
			return
		}

		line := replScanner.Text()
		fmt.Println(line)
	}
}

// runSource starts scanning of the source code.
func (i *glox) runSource(source string) {
	scanner := scanner.New(source)
	scanner.ScanTokens()
	// TODO: nonzero exit code in case of error in scanning
}

// RunRootCmd is the function to be executed as the base command.
func RunRootCmd(cmd *cobra.Command, args []string) {
	interpreter := createInterpreter()

	numArgs := len(args)
	if numArgs == 0 {
		// Read-Eval-Print-Loop
		interpreter.runLoop()
	} else if numArgs == 1 {
		// read the source file
		content, err := ioutil.ReadFile(args[0])
		utils.CheckError(err)

		interpreter.runSource(string(content))
	} else {
		fmt.Println("Usage: glox [script]")
	}
}

func init() {
	cmd.RootCmd.Run = RunRootCmd
}

func main() {
	cmd.Execute()
}
