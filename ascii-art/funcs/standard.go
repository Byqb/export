package funcs

import (
	"fmt"
	"io"
	"os"
)

func Standard(font string) error {
	file, err := os.Open("ascii-art/fonts/" + font + ".txt") // open the file
	if err != nil {                                          // handle error
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	data := make([]byte, 1024) // create buffer
	for {
		n, err := file.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			return err
		}

		if n == 0 {
			break
		}
	}
	return nil
}
