package fileops

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
	"github.com/spf13/viper"
)

// CreateFiles given a FetchedProblem create all it's corresponding files
func CreateFiles(p *types.FetchedProblem) {
	problemPath := path.Join(viper.GetString("directory"), problemPath(p))
	os.MkdirAll(problemPath, os.ModePerm)

	for index, testcase := range p.Testcases {
		createTestcaseFiles(problemPath, index, &testcase)
	}

	templateBytes, templateErr := ioutil.ReadFile(viper.GetString("template"))
	template := ""
	if templateErr == nil {
		template = string(templateBytes)
	}

	createFile(path.Join(problemPath, "sol.cpp"), template)
}

// createTestCaseFile will create the input and output files for the given
// Testcase
func createTestcaseFiles(problemPath string, index int, t *types.Testcase) {
	inputPath := path.Join(problemPath, fmt.Sprintf("inp%v.txt", index))
	outputPath := path.Join(problemPath, fmt.Sprintf("exp%v.txt", index))

	createFile(inputPath, t.Input)
	createFile(outputPath, t.Output)
}
