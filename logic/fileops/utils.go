package fileops

import (
	"io/ioutil"
	"path"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
	"github.com/spf13/cobra"
)

// ProblemPath gets the path corresponding to the given problem
func problemPath(p *types.FetchedProblem) string {
	problemInfo := p.ProblemInfo

	var website string
	switch problemInfo.Website {
	case types.Atcoder:
		website = "atcoder"
	case types.Codeforces:
		website = "codeforces"
	}

	problemPath := path.Join(website, problemInfo.Contest, problemInfo.Problem)

	return problemPath
}

// createFile creates a file at the given filePath and fill it with the
// specified content
func createFile(filePath string, content string) {
	contentAsBytes := []byte(filePath)
	err := ioutil.WriteFile(filePath, contentAsBytes, 0644)
	cobra.CheckErr(err)
}
