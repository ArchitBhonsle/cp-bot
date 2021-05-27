package fileops

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CreateFiles(p *types.FetchedProblem) {
	problemPath := path.Join(viper.GetString("directory"), ProblemPath(p))
	os.MkdirAll(problemPath, os.ModePerm)

	for index, testcase := range p.Testcases {
		createTestcaseFiles(problemPath, index, &testcase)
	}
}

func createTestcaseFiles(problemPath string, index int, t *types.Testcase) {
	inputPath := path.Join(problemPath, fmt.Sprintf("i%v.txt", index))
	outputPath := path.Join(problemPath, fmt.Sprintf("e%v.txt", index))

	createFile(inputPath, t.Input)
	createFile(outputPath, t.Output)
}

func createFile(filePath string, content string) {
	contentAsBytes := []byte(filePath)
	err := ioutil.WriteFile(filePath, contentAsBytes, 0644)
	cobra.CheckErr(err)
}
