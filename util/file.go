package util

import (
	"bufio"
	"log"
	"os"
)

func OpenLineByLine(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	return text
}
