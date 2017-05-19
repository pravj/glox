package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pravj/glox/cmd"
	"github.com/pravj/glox/scanner"
	"github.com/pravj/glox/errors"
	"github.com/spf13/cobra"
)

type glox struct {
	// contains unexported functions

	// hasError, flag for error handling
	hasError bool
}

func createInterpreter() *glox {
	return &glox{hasError:false}
}

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

func (i *glox) runSource(source string) {
	scanner := scanner.New(source)
	scanner.ScanTokens()
	// nonzero exit code in case of error in scanning
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	interpreter := createInterpreter()

	numArgs := len(args)
	if numArgs == 0 {
		// Read-Eval-Print-Loop
		interpreter.runLoop()
	} else if numArgs == 1 {
		// read the source file
		content, err := ioutil.ReadFile(args[0])
		errors.CheckError(err)

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
