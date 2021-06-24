package check

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
)

// CountInputs counts the number of input files in the current directory
func CountInputs(directory string) (int, error) {
	pattern := path.Join(directory, "inp*.txt")

	count, err := filepath.Glob(pattern)
	if err != nil {
		return -1, err
	}

	return len(count), nil
}

// Execute the solution in the given directory using the specified input file
// and pipes the output to the corressponding output file
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

	// Receiving output
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
