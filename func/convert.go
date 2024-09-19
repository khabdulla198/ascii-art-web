package web

import (
	"bufio"
	"log"
	"os"
)

func Convert(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		//log.Fatalf("Error: %v", err)
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f) //reads file yuh

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: failed to read file: %s", err)
	}
	return lines, nil
}
