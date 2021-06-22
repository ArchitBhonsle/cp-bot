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

	"github.com/ArchitBhonsle/cp-bot/logic/compare"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Test the program against the input files",
	Long: `This command command is used to run a specific solution file
with the corressponging of input files and check the produced output against
the expected output. Example:

cp-bot run
cp-bot run /path/to/problem`,
	RunE: func(cmd *cobra.Command, args []string) error {
		problemDirectory, errGetwd := os.Getwd()
		if errGetwd != nil {
			return errGetwd
		}

		if len(args) == 1 {
			problemDirectory = args[0]
		}

		errRun := compare.Diff(problemDirectory)
		if errRun != nil {
			return errRun
		}

		if compileErr := compare.Compile(problemDirectory); compileErr != nil {
			return compileErr
		}

		count, countErr := compare.CountInputs(problemDirectory)
		if countErr != nil {
			return nil
		}
		for i := 0; i < count; i++ {
			compare.Execute(problemDirectory, i)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
