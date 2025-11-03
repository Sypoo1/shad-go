package main

import (
	"bufio"
	"fmt"
	"os"
)

func helper(err error) {
	if err != nil {
		panic(err)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}

}

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("put paths")
	}

	linesMap := make(map[string]int)

	for _, filePath := range args[1:] {
		file, err := os.Open(filePath)
		helper(err)

		defer closeFile(file)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			linesMap[line] = linesMap[line] + 1
		}

		helper(scanner.Err())
	}

	for line, count := range linesMap {
		if count >= 2 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
