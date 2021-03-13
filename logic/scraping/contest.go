package scraping

import (
	"net/url"
	"regexp"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

func getWebsiteFromHostname(hostname string) types.Website {
	regex := regexp.MustCompile(`(?:\w*\.)?(\w*)(?:\.\w*)`)

	websiteName := regex.FindStringSubmatch(hostname)[1]
	switch websiteName {
	case "codeforces":
		return types.Codeforces
	case "atcoder":
		return types.Atcoder
	default:
		return types.Invalid
	}
}

// GetContest returns a Contest parsed from the URL
func GetContest(rawurl string) (types.Contest, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	website := getWebsiteFromHostname(parsedURL.Hostname())
	contestID := getContestIDFromPath(parsedURL.EscapedPath(), website)

	var contest types.Contest
	switch website {
	case types.Atcoder:
		contest = GetAtcoderProblems(contestID)
	case types.Codeforces:
		contest = GetCodeforcesProblems(contestID)
	}

	return contest, nil
}

func getContestIDFromPath(path string, website types.Website) string {
	switch website {
	case types.Atcoder:
		regex := regexp.MustCompile(`/contests/(\w+)`)
		return regex.FindStringSubmatch(path)[1]
	case types.Codeforces:
		regex := regexp.MustCompile(`/contest/(\w+)`)
		return regex.FindStringSubmatch(path)[1]
	default:
		return ""
	}
}
