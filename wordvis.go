package wordvis

import (
	"fmt"
	"github.com/jpvio/word-vis/reading"
)

// Run wordvis
func Run() {
	fmt.Println("words")
	fmt.Println(reading.ReadFileIntoWordArray("thing"))
}
