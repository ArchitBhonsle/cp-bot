package types

// ProblemInfo is the problem's metadata
type ProblemInfo struct {
	Website Website
	Contest string
	Problem string
	URL     string
}

// Problem stores the data about a problem
type Problem interface {
	Fetch(chan *FetchedProblem)
	GetInfo() *ProblemInfo
}

// FetchedProblem has the path to save the problem to and it's testcases
type FetchedProblem struct {
	ProblemInfo *ProblemInfo
	Testcases   []Testcase
}

// Testcase represents a single test case made of input and output
type Testcase struct {
	Input  string
	Output string
}

// Contest is a slice of Problems
type Contest []Problem

// Website stores what website the problem is from
type Website int

const (
	// Invalid if the website isn't supported
	Invalid Website = iota
	// Atcoder if url is of the form https://atcoder.jp/contests/<contest-id>
	Atcoder
	// Codeforces if url is of the form https://codeforces.com/contest/<contest-id>
	Codeforces
)
