package reading

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileIntoWordArray(path string) []string {
	var result []string
	// Open the file.
	f, _ := os.Open("resources/file.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)

		result = append(result, strings.ToLower(word))
	}
	return result
}
