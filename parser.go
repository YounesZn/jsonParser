package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input:")
	if scanner.Scan() {
		input := scanner.Text()
		tokens := Tokenizer(input)
		for _, token := range tokens {
			fmt.Printf("Type: %v, Value: %v\n", token.tokenType.String(), token.value)
		}
	} else {
		log.Fatal("Failed to read input")
	}
}
