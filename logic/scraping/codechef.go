package scraping

import "fmt"

// GetCodechefProblems given a contest_id will scrape it's problems
func GetCodechefProblems(contest string) {
	contestURL := fmt.Sprintf("https://www.codechef.com/%v", contest)
	println(contestURL)
}
