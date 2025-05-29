package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"subsequences/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input_file>\n", os.Args[0])
		os.Exit(1)
	}
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "No data in file")
		os.Exit(1)
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid n: %v\n", err)
		os.Exit(1)
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Insufficient data in file")
			os.Exit(1)
		}
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number at position %d: %v\n", i+1, err)
			os.Exit(1)
		}
		a[i] = x
	}

	result := internal.FindMinFullAlphabet(a)
	if result == -1 {
		fmt.Println("NONE")
	} else {
		fmt.Println(result)
	}
}
