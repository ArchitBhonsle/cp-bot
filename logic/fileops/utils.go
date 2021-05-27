package fileops

import (
	"path"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

// ProblemPath gets the path corresponding to the given problem
func ProblemPath(p *types.FetchedProblem) string {
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
