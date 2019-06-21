package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/veer66/mapkha"
)

func main() {
	dict, err := mapkha.LoadDefaultDict()
	if err != nil {
		log.Fatal(err)

	}
	wordcut := mapkha.NewWordcut(dict)

	// Thank @iporsut for figuring out Scanner problem for me.
	max := 1024 * 1024 * 10

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, max), max)

	for scanner.Scan() {
		fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
