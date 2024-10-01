package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ParseTocFile(path, pattern string) string {
	file, err := os.Open(filepath.Join(path, filepath.Base(path)+".toc"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(pattern)
	
	for scanner.Scan() {
		matches := regex.FindStringSubmatch(scanner.Text())
		if len(matches) > 1 {
			return matches[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return ""
}
