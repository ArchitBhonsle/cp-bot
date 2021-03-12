package scraping

import (
	"fmt"
)

// CodeforcesProblem represents a problem on codeforces.com
type CodeforcesProblem struct {
	contest string
	problem string
	url     string
}

// GetCodeforcesProblems given a contest_id will scrape it's problems
func GetCodeforcesProblems(contest string) {
	contestURL := fmt.Sprintf("https://codeforces.com/contests/%v", contest)
	println(contestURL)
}
