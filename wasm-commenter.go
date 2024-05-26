package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: wasm-commenter <input.wat> <output.wat>")
		os.Exit(1)
	}

	if !strings.HasSuffix(os.Args[1], ".wat") {
		fmt.Println("Input file must have .wat extension")
		os.Exit(1)
	}

	fmt.Printf("Commenting %v...\n", os.Args[1])

	var (
		commentedLines int
		start          = time.Now()
	)

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	output := func() *os.File {
		file, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		return file
	}()
	defer output.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content := scanner.Text()

		for _, operation := range Operations {
			operationName := strings.ToLower(
				strings.Replace(operation.Name, "_", ".", -1),
			)
			if strings.Contains(content, operationName) {
				content = fmt.Sprintf("%v ;; %v", content, operation.Comment)
				commentedLines++
			}
		}

		io.WriteString(output, content+"\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Commented %v lines in %v\n", commentedLines, time.Since(start))
}
