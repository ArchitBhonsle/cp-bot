package scraping

import (
	"fmt"
	"strings"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
	"github.com/gocolly/colly"
)

// AtcoderProblem represents a problem on atcoder.jp
type AtcoderProblem struct {
	Contest string
	Problem string
	URL     string
}

const taskLinksSelector = "tbody > tr > td:first-child > a"

// GetAtcoderProblems given a contest_id will scrape it's problems
func GetAtcoderProblems(contestID string) types.Contest {
	contestURL := fmt.Sprintf("https://atcoder.jp/contests/%v", contestID)

	tasksListURL := fmt.Sprintf("%v/tasks", contestURL)

	var contest types.Contest

	collector := colly.NewCollector()

	collector.OnHTML(taskLinksSelector, func(e *colly.HTMLElement) {
		contest = append(contest, &AtcoderProblem{
			Contest: contestID,
			Problem: e.Text,
			URL:     fmt.Sprintf("https://atcoder.jp%v", e.Attr("href")),
		})
	})

	collector.Visit(tasksListURL)

	return contest
}

const testCasesSelector = "span.lang-en h3+pre"

// Scrape will get the problem's inputs and corresponding outputs
func (p *AtcoderProblem) Scrape(sync chan bool) []types.TestCase {
	collector := colly.NewCollector()

	var inputsAndOutputs []string
	collector.OnHTML(testCasesSelector, func(e *colly.HTMLElement) {
		inputsAndOutputs = append(inputsAndOutputs, strings.Trim(e.Text, " \n"))
	})

	collector.Visit(p.URL)

	var testCases []types.TestCase
	for i := 0; i < len(inputsAndOutputs); i += 2 {
		testCases = append(testCases, types.TestCase{
			Input:  inputsAndOutputs[i],
			Output: inputsAndOutputs[i+1],
		})
	}

	return testCases
}

// GetInfo returns the corresponding problem's metadata
func (p *AtcoderProblem) GetInfo() types.ProblemInfo {
	return types.ProblemInfo{
		Contest: p.Contest,
		Problem: p.Problem,
		URL:     p.URL,
	}
}
