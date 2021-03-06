package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

// AtcoderProblem represents a problem on atcoder.jp
type AtcoderProblem struct {
	contest string
	problem string
	url     string
}

const acProblemsSelector = "tbody > tr > td:first-child > a"

// GetAtcoderProblems given a contest_id will scrape it's problems
func GetAtcoderProblems(contestID string) (types.Contest, error) {
	tasksListURL := fmt.Sprintf("https://atcoder.jp/contests/%v/tasks", contestID)

	var contest types.Contest
	var err error
	collector := colly.NewCollector()

	collector.OnError(func(_r *colly.Response, collyError error) {
		err = collyError
	})

	collector.OnHTML(acProblemsSelector, func(e *colly.HTMLElement) {
		contest = append(contest, &AtcoderProblem{
			contest: contestID,
			problem: strings.Trim(e.Text, " \n"),
			url:     fmt.Sprintf("https://atcoder.jp%v", e.Attr("href")),
		})
	})

	collector.Visit(tasksListURL)

	return contest, err
}

const acTestcasesSelector = "span.lang-en h3 + pre"

// Scrape will get the problem's inputs and corresponding outputs
func (p *AtcoderProblem) Fetch(send chan *types.FetchedProblem) {
	var inputsAndOutputs []string
	collector := colly.NewCollector()

	collector.OnHTML(acTestcasesSelector, func(e *colly.HTMLElement) {
		inputsAndOutputs = append(inputsAndOutputs, strings.Trim(e.Text, " \n"))
	})

	collector.Visit(p.url)

	var testcases []types.Testcase
	for i := 0; i < len(inputsAndOutputs); i += 2 {
		testcases = append(testcases, types.Testcase{
			Input:  inputsAndOutputs[i],
			Output: inputsAndOutputs[i+1],
		})
	}

	send <- &types.FetchedProblem{
		ProblemInfo: p.GetInfo(),
		Testcases:   testcases,
	}
}

// GetInfo returns the corresponding problem's metadata
func (p *AtcoderProblem) GetInfo() *types.ProblemInfo {
	return &types.ProblemInfo{
		Website: types.Atcoder,
		Contest: p.contest,
		Problem: p.problem,
		URL:     p.url,
	}
}
