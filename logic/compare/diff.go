package compare

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
)

func Diff(directory string) error {
	return nil
}

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

	println("Solution successfully compiled")

	return nil
}

func CountInputs(directory string) (int, error) {
	pattern := path.Join(directory, "inp*.txt")

	count, err := filepath.Glob(pattern)
	if err != nil {
		return -1, err
	}

	return len(count), nil
}

func Execute(directory string, inputNumber int) error {
	executable := path.Join(directory, "sol")
	executeCmd := exec.Command(executable)

	// Taking input
	inputBuffer, inputBufferError := ioutil.ReadFile(
		path.Join(directory, fmt.Sprintf("inp%v.txt", inputNumber)))
	if inputBufferError != nil {
		return inputBufferError
	}
	stdin, stdinErr := executeCmd.StdinPipe()
	if stdinErr != nil {
		return stdinErr
	}
	go func() {
		defer stdin.Close()
		stdin.Write(inputBuffer)
	}()

	// Recieving output
	output, runtimeErr := executeCmd.CombinedOutput()
	if runtimeErr != nil {
		return runtimeErr
	}
	outputBufferError := ioutil.WriteFile(
		path.Join(directory, fmt.Sprintf("out%v.txt", inputNumber)),
		output,
		0644,
	)
	if outputBufferError != nil {
		return outputBufferError
	}

	return nil
}

