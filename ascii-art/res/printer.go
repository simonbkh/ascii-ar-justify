package fs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var IsOtput bool

func Printer(inputLine string, slice [][]string) string {
	var s string
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	raw, _ := cmd.Output()

	for i := 0; i < len(raw); i++ {
		if raw[i] == 32 && i != len(raw)-1 {
			raw = raw[i+1:len(raw)-1]
			break
		}
	}
	size, _ := strconv.ParseInt(string(raw),10,0)
	
		for j := 0; j < 8; j++ {
			for a := 0; a < int(size)-len(inputLine)*8; a++ {
				fmt.Print(" ")
			}
			for _, char := range inputLine {
				index := int(char) - 32
				if IsOtput {
					s += slice[index][j]
				} else {
					fmt.Print(slice[index][j])
				}

			}
			if IsOtput {
				s += "\n"
			} else {
				fmt.Println()
			}

		}
	
	return s
}
