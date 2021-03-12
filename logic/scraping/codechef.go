package scraping

import "fmt"

// CodechefProblem represents a problem on codechef.com
type CodechefProblem struct {
	contest string
	problem string
	url     string
}

// GetCodechefProblems given a contest_id will scrape it's problems
func GetCodechefProblems(contest string) {
	contestURL := fmt.Sprintf("https://www.codechef.com/%v", contest)
	println(contestURL)
}
