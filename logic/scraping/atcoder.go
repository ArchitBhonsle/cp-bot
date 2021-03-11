package scraping

import "fmt"

// GetAtcoderProblems given a contest_id will scrape it's problems
func GetAtcoderProblems(contest string) {
	contestURL := fmt.Sprintf("https://atcoder.jp/contests/%v", contest)
	println(contestURL)
}
