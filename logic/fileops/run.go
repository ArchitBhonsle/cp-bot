package fileops

import (
	"fmt"
	"os/exec"
	"path"
)

func Run(directory string) error {
	solutionPath := path.Join(directory, "sol.cpp")
	executablePath := path.Join(directory, "sol")

	compileCmd := exec.Command(
    		"g++",
		"-std=c++17",
		"-Wshadow",
		"-Wall",
		fmt.Sprint("-o", executablePath),
		solutionPath,
		"-g",
		"-fsanitize=address",
		"-fsanitize=undefined",
		"-D_GLIBCXX_DEBUG",
	)

	output, err := compileCmd.CombinedOutput()
	if err != nil {
		println(string(output))
		println(err.Error())
		return nil
	}

	println("Solution successfully compiled")

	return nil
}
