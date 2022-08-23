package main

import (
	"fmt"
	read_line "git.supremind.info/gobase/io/read-line"
)

func main() {
	file := read_line.DoWriteFile("./test.txt")
	for i := 0; i < 5; i++ {
		_ = read_line.Do(fmt.Sprintf("%v", i), file)
	}
}
