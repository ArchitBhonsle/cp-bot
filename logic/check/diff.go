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
	if terminalWidth > 80 {
		terminalWidth = 80
	}

	diffOutput, err := exec.Command(
		"diff",
		outFile,
		expFile,
		fmt.Sprintf("--width=%v", terminalWidth-2),
		"--side-by-side",
		"--ignore-space-change",
		"--ignore-trailing-space",
	).CombinedOutput()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
		default:
			return err
		}
	}

	if err == nil {
		diffOutput = nil
	}
	formatOutput(inputNumber, string(diffOutput), terminalWidth)
	fmt.Print(drawing.Reset)

	return nil
}

func formatOutput(inputNumber int, output string, terminalWidth int) {
	spaces := strings.Repeat(" ", terminalWidth-12)

	if output == "" {
		print(drawing.Green)
	} else {
		print(drawing.Red)
	}

	println(drawing.HorizontalTop(terminalWidth))

	fmt.Printf("%v out%v%vexp%v %v\n",
		drawing.Vertical,
		inputNumber,
		spaces,
		inputNumber,
		drawing.Vertical,
	)

	println(drawing.HorizontalBottom(terminalWidth))

	print(drawing.Reset)

	println(output)
}
