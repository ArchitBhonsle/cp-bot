package check

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/ArchitBhonsle/cp-bot/logic/drawing"
)

// Diff prints out the differences between each of the expected output and the
// output produced by the user's solution
func Diff(directory string, inputNumber int) error {
	outFile := path.Join(directory, fmt.Sprintf("out%v.txt", inputNumber))
	expFile := path.Join(directory, fmt.Sprintf("exp%v.txt", inputNumber))

	terminalWidth, err := drawing.TerminalWidth()
	if err != nil {
		return nil
	}
	diffOutput, err := exec.Command(
		"diff",
		outFile,
		expFile,
		"-b",
		"-y",
		fmt.Sprintf("--width=%v", terminalWidth-2),
	).CombinedOutput()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
		default:
			return err
		}
	}

	formatOutput(inputNumber, string(diffOutput), terminalWidth)

	return nil
}

func formatOutput(inputNumber int, output string, terminalWidth int) {
	spaces := strings.Repeat(" ", terminalWidth-12)

	println(drawing.HorizontalTop(terminalWidth))

	fmt.Printf("%v out%v%vexp%v %v\n",
		drawing.Vertical,
		inputNumber,
		spaces,
		inputNumber,
		drawing.Vertical,
	)

	println(drawing.HorizontalBottom(terminalWidth))

	println(output)
}
