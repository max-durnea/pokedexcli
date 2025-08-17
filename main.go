package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
type command struct {
	name string
	description string
	callback func() error
}
var Commands map[string]command
func main() {
	Commands = map[string]command{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
    }
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input:=cleanInput(input)
		val,ok := Commands[cleaned_input[0]]
		if ok {
			val.callback()
		}else{
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string{
	lowerStr := strings.ToLower(text)
	parts := strings.Fields(lowerStr)
	return parts
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!");
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range Commands{
		fmt.Printf("%v: %v\n",value.name, value.description)
	}
	return nil
}