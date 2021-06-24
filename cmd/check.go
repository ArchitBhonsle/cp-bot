/*
Copyright © 2021 Archit Bhonsle <abhonsle2000@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"sync"

	"github.com/ArchitBhonsle/cp-bot/logic/check"
	"github.com/spf13/cobra"
)

var directoryFlag string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the program against the input files",
	Long: `This command command is used to check a specific solution file with the 
corressponding of input files and diff the produced output against the expected
output. Example:

cp-bot check
cp-bot check --dir /path/to/problem`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Deciding the directory to run the command for
		problemDirectory, errGetwd := os.Getwd()
		if errGetwd != nil {
			return errGetwd
		}
		if directoryFlag != "" {
			problemDirectory = directoryFlag
		}

		// Compile the user's solution
		if compileErr := check.Compile(problemDirectory); compileErr != nil {
			return compileErr
		}

		// Count the number input files
		count, countErr := check.CountInputs(problemDirectory)
		if countErr != nil {
			return countErr
		}

		// Execute the solution against each of the input files
		var executeWaitGroup sync.WaitGroup
		for i := 0; i < count; i++ {
			executeWaitGroup.Add(1)

			go func(inputNumber int) {
				defer executeWaitGroup.Done()
				// TODO error checking here
				check.Execute(problemDirectory, inputNumber)
			}(i)
		}
		executeWaitGroup.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	rootCmd.PersistentFlags().StringVar(&directoryFlag, "dir", "", "directory to run command on")
}
