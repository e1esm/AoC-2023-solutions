package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("full.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	rows := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return
		}
		rows = append(rows, line)
	}
	fmt.Println()
	fmt.Println(getDigitsSum(rows))
}
