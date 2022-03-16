package main

import (
	"fmt"
	"io"      // io.EOF
	"strings" // strings.NewReader
)

func main() {
	reader := strings.NewReader("Hello Reader")
	buffer := make([]byte, 8, 8)
	for {
		n, err := reader.Read(buffer)
		fmt.Printf("n = %v err = %v buffer = %v", n, err, buffer)
		fmt.Printf("buffer[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}
}
