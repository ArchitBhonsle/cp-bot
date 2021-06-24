package check

import (
	"fmt"
	"os/exec"
	"path"
)

// Compile compiles the solution in the specified directory
// TODO support more languages
func Compile(directory string) error {
	solution := path.Join(directory, "sol.cpp")
	executable := path.Join(directory, "sol")

	compileCmd := exec.Command(
		"g++",
		"-std=c++17",
		"-Wshadow",
		"-Wall",
		fmt.Sprint("-o", executable),
		solution,
		"-g",
		"-fsanitize=address",
		"-fsanitize=undefined",
		"-D_GLIBCXX_DEBUG",
	)

	if output, err := compileCmd.CombinedOutput(); err != nil {
		println(string(output))
		return err
	}

	return nil
}
