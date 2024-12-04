package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return nil
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
