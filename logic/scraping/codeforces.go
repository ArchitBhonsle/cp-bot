package scraping

import (
	"fmt"
	"strings"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
	"github.com/gocolly/colly"
)

// CodeforcesProblem represents a problem on codeforces.com
type CodeforcesProblem struct {
	contest string
	problem string
	url     string
}

const cfProblemsSelector = "table.problems > tbody > tr > td:nth-of-type(1) > a"

// GetCodeforcesProblems given a contest_id will scrape it's problems
func GetCodeforcesProblems(contestID string) (types.Contest, error) {
	contestURL := fmt.Sprintf("https://codeforces.com/contest/%v", contestID)

	var contest types.Contest
	var err error
	collector := colly.NewCollector()

	collector.OnError(func(_r *colly.Response, collyError error) {
		err = collyError
	})

	collector.OnHTML(cfProblemsSelector, func(e *colly.HTMLElement) {
		contest = append(contest, &CodeforcesProblem{
			contest: contestID,
			problem: strings.Trim(e.Text, " \n"),
			url:     fmt.Sprintf("https://codeforces.com%v", e.Attr("href")),
		})
	})

	collector.Visit(contestURL)

	return contest, err
}

const cfTestcasesSelector = "div.sample-tests div.sample-test"

// Scrape will get the problem's inputs and corresponding outputs
func (p *CodeforcesProblem) Fetch(send chan *types.FetchedProblem) {
	var testcases []types.Testcase
	var inputs []string
	var outputs []string
	collector := colly.NewCollector()

	collector.OnHTML(cfTestcasesSelector, func(e *colly.HTMLElement) {
		e.ForEach(".input > pre", func(idx int, ce *colly.HTMLElement) {
			inputs = append(inputs, strings.Trim(ce.Text, " \n"))
		})
		e.ForEach(".output > pre", func(idx int, ce *colly.HTMLElement) {
			outputs = append(outputs, strings.Trim(ce.Text, " \n"))
		})
	})

	collector.Visit(p.url)

	// TODO add len(inputs) == len(outputs) check
	for i := 0; i < len(inputs); i += 1 {
		testcases = append(testcases, types.Testcase{
			Input:  inputs[i],
			Output: outputs[i],
		})
	}

	send <- &types.FetchedProblem{
		ProblemInfo: p.GetInfo(),
		Testcases:   testcases,
	}
}

// GetInfo returns the corresponding problem's metadata
func (p *CodeforcesProblem) GetInfo() *types.ProblemInfo {
	return &types.ProblemInfo{
		Website: types.Codeforces,
		Contest: p.contest,
		Problem: p.problem,
		URL:     p.url,
	}
}
