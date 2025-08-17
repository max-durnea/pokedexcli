package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input:=cleanInput(input)
		fmt.Printf("Your command was: %v\n",cleaned_input[0])
	}

}

func cleanInput(text string) []string{
	lowerStr := strings.ToLower(text)
	parts := strings.Fields(lowerStr)
	return parts
}