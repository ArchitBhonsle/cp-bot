package drawing

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		return -1, err
	}

	dimensions := strings.Split(strings.TrimSpace(string(out)), " ")
	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return -1, nil
	}

	return width, nil
}
