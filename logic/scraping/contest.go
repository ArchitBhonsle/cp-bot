package scraping

import (
	"errors"
	"net/url"
	"regexp"

	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

func getWebsiteFromHostname(hostname string) (types.Website, error) {
	regex := regexp.MustCompile(`(?:\w*\.)?(\w*)(?:\.\w*)`)
	matches := regex.FindStringSubmatch(hostname)
	if len(matches) != 2 {
		return types.Invalid, errors.New("website couldn't be parsed")
	}
	websiteName := matches[1]

	var website types.Website
	switch websiteName {
	case "codeforces":
		website = types.Codeforces
	case "atcoder":
		website = types.Atcoder
	default:
		return types.Invalid, errors.New("website not supported")
	}

	return website, nil
}

// GetContest returns a Contest parsed from the URL
func GetContest(rawurl string) (types.Contest, error) {
	parsedURL, parseURLError := url.Parse(rawurl)
	if parseURLError != nil {
		return nil, parseURLError
	}

	website, websiteParsingError := getWebsiteFromHostname(parsedURL.Hostname())
	if websiteParsingError != nil {
		return nil, websiteParsingError
	}

	contestID, contestIDError := getContestIDFromPath(parsedURL.EscapedPath(), website)
	if contestIDError != nil {
		return nil, contestIDError
	}

	switch website {
	case types.Atcoder:
		return GetAtcoderProblems(contestID)
	case types.Codeforces:
		return GetCodeforcesProblems(contestID)
	}

	return nil, errors.New("something that wasn't supposed to go wrong went wrong")
}

func getContestIDFromPath(path string, website types.Website) (string, error) {
	var matches []string
	switch website {
	case types.Atcoder:
		regex := regexp.MustCompile(`/contests/(\w+)`)
		matches = regex.FindStringSubmatch(path)
	case types.Codeforces:
		regex := regexp.MustCompile(`/contest/(\w+)`)
		matches = regex.FindStringSubmatch(path)
	}

	if len(matches) != 2 {
		return "", errors.New("contest ID not found")
	}
	contestID := matches[1]

	return contestID, nil
}
