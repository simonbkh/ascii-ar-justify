package fs

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalColumns() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()

	columns := strings.Fields(string(output))[1]
	cols, _ := strconv.Atoi(columns)

	return cols
}
