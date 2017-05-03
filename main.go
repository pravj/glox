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

func runLoop() {
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

func runSource(source string) {
	scanner := scanner.New(source)
	fmt.Println(scanner)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	numArgs := len(args)
	if numArgs == 0 {
		// Read-Eval-Print-Loop
		runLoop()
	} else if numArgs == 1 {
		// read the source file
		content, err := ioutil.ReadFile(args[0])
		utils.CheckError(err)

		runSource(string(content))
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
