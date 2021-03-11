package scraping

import "fmt"

// GetCodeforcesProblems given a contest_id will scrape it's problems
func GetCodeforcesProblems(contest string) {
	contestURL := fmt.Sprintf("https://codeforces.com/contests/%v", contest)
	println(contestURL)
}
