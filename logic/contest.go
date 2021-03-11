package logic

import (
	"net/url"
	"regexp"

	"github.com/ArchitBhonsle/cp-bot/logic/scraping"
)

// Website stores what website the problem is from
type Website int

const (
	// Invalid if the website isn't supported
	Invalid Website = iota
	// Atcoder if url is of the form https://atcoder.jp/contests/<contest-id>
	Atcoder
	// Codechef if url is of the form https://www.codechef.com/<contest-id>
	Codechef
	// Codeforces if url is of the form https://codeforces.com/contest/<contest-id>
	Codeforces
)

func getWebsiteFromHostname(hostname string) Website {
	regex := regexp.MustCompile(`(?:\w*\.)?(\w*)(?:\.\w*)`)

	websiteName := regex.FindStringSubmatch(hostname)[1]
	switch websiteName {
	case "codeforces":
		return Codeforces
	case "codechef":
		return Codechef
	case "atcoder":
		return Atcoder
	default:
		return Invalid
	}
}

// Contest is a slice of Problems
type Contest []Problem

// GetContest returns a Contest parsed from the URL
func GetContest(rawurl string) (Contest, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	website := getWebsiteFromHostname(parsedURL.Hostname())
	contestID := getContestIDFromPath(parsedURL.EscapedPath(), website)
	switch website {
	case Atcoder:
		scraping.GetAtcoderProblems(contestID)
	case Codechef:
		scraping.GetCodechefProblems(contestID)
	case Codeforces:
		scraping.GetCodeforcesProblems(contestID)
	}

	return nil, nil
}

func getContestIDFromPath(path string, website Website) string {
	switch website {
	case Atcoder:
		regex := regexp.MustCompile(`/contests/(\w+)`)
		return regex.FindStringSubmatch(path)[1]
	case Codechef:
		regex := regexp.MustCompile(`/(\w+)`)
		return regex.FindStringSubmatch(path)[1]
	case Codeforces:
		regex := regexp.MustCompile(`/contest/(\w+)`)
		return regex.FindStringSubmatch(path)[1]
	default:
		return ""
	}
}
