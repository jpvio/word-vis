package reading

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileIntoWordArray(path string) ([]string, error) {
	var result []string
	// Open the file.
	f, err := os.Open("resources/file.txt")
	if err != nil {
		return result, err
	}

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		word := scanner.Text()
		result = append(result, strings.ToLower(word))
	}
	return result, nil
}
