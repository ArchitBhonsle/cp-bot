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
func GetCodeforcesProblems(contestID string) types.Contest {
	contestURL := fmt.Sprintf("https://codeforces.com/contest/%v", contestID)

	var contest types.Contest
	collector := colly.NewCollector()

	collector.OnHTML(cfProblemsSelector, func(e *colly.HTMLElement) {
		contest = append(contest, &CodeforcesProblem{
			contest: contestID,
			problem: strings.Trim(e.Text, " \n"),
			url:     fmt.Sprintf("https://codeforces.com%v", e.Attr("href")),
		})
	})

	collector.Visit(contestURL)

	return contest
}

const cfTestcasesSelector = "div.sample-tests div.sample-test"

// Scrape will get the problem's inputs and corresponding outputs
func (p *CodeforcesProblem) Fetch(send chan *types.FetchedProblem) {
	var testcases []types.Testcase
	collector := colly.NewCollector()

	collector.OnHTML(cfTestcasesSelector, func(e *colly.HTMLElement) {
		testcases = append(testcases, types.Testcase{
			Input:  strings.Trim(e.ChildText(".input > pre"), " \n"),
			Output: strings.Trim(e.ChildText(".output > pre"), " \n"),
		})
	})

	collector.Visit(p.url)

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
