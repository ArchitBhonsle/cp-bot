package fileops

import (
	"fmt"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

// ProblemPath gets the path corresponding to the given problem
func ProblemPath(problemInfo *types.ProblemInfo) string {
	var website string
	switch problemInfo.Website {
	case types.Atcoder:
		website = "atcoder"
	case types.Codeforces:
		website = "coderforces"
	}

	path := fmt.Sprintf("%v/%v/%v", website, problemInfo.Contest, problemInfo.Problem)

	return path
}
