package webscaper

import (
	"net/url"
	"regexp"
)

// Website stores what website the problem is from
type Website int

const (
	// Invalid if the website isn't supported
	Invalid Website = iota
	// Codeforces if url is of the form https://codeforces.com/contest/<contest-id>
	Codeforces
	// Codechef if url is of the form https://www.codechef.com/<contest-id>
	Codechef
	// Atcoder if url is of the form https://atcoder.jp/contests/<contest-id>
	Atcoder
)

// Problem stores the data about a problem
type Problem struct {
	website Website
	contest string
	problem string
}

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

// Problems returns a slice of Problems parsed from the URL
func Problems(rawurl string) ([]Problem, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	website := getWebsiteFromHostname(parsedURL.Hostname())
	println(website)

	return nil, nil
}
