package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	showLineNums := flag.Bool("n", false, "show line numbers")
	flag.Parse()

	fileNames := flag.Args()

	for i := 0; i < len(fileNames); i++ {
		processFile(fileNames[i], *showLineNums)
	}
}

func processFile(filename string, showLineNums bool) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 1
	for scanner.Scan() {
		if showLineNums {
			fmt.Printf("%5d  ", lineNo)
		}
		fmt.Println(scanner.Text())
		lineNo += 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
