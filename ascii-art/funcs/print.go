package funcs

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Print(words []string, font string) (string, error) { //print the art
	num := 0
	var result string

	preWord := ""
	err := Standard(font)
	if err != nil {
		return "", err
	}
	for _, word := range words {
		num = num + 1
		asciiValues := Math(word)
		if len(asciiValues) == 0 {
			if num < len(words) || preWord != "" {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		path := "ascii-art/fonts/" + font + ".txt"
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		lineNumber := 0
		for i := 0; i < 8; i++ {
			if (word) == "\r" {
				if i == 0 && word == "\r" {
					result += "\n"
				}
				continue
			}
			for _, line := range asciiValues {
				file.Seek(0, io.SeekStart)
				scanner = bufio.NewScanner(file)
				lineNumber = 0

				for lineNumber < line+i && scanner.Scan() {
					lineNumber++
				}
				result += fmt.Sprintf("%s", scanner.Text())
			}
			if result != "" {
				result += "\n"
			}

		}
		preWord = word
	}
	return result, err
}
