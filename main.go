package main

import (
	fs "ascii-art-justify/ascii-art/res"
	"fmt"
)

func main() {
	size := fs.GetTerminalColumns()
	fmt.Println(size)
}
