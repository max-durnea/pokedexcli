package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/max-durnea/pokedexcli/pokeapi"
)
type command struct {
	name string
	description string
	callback func(cfg *Config) error
}
type Config struct {
	Next string
	Prev string
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
		"map": {
			name: "map",
			description: "Lists the next 20 locations everytime it's called",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous 20 locations everytime it's called",
			callback: commandBmap,
		},
    }
	locationConfig := Config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Prev: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input:=cleanInput(input)
		val,ok := Commands[cleaned_input[0]]
		if ok {
			val.callback(&locationConfig)
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

func commandExit(cfg *Config) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!");
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range Commands{
		fmt.Printf("%v: %v\n",value.name, value.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	locationPage, err := pokeapi.FetchLocationPage(cfg.Next)
	if(err!=nil){
		return err
	}
	fmt.Println(locationPage.Previous)
	fmt.Println(locationPage.Next)
	cfg.Next = locationPage.Next
	cfg.Prev = locationPage.Previous
	for _, value := range locationPage.Results {
		fmt.Println(value.Name)
	}
	return nil
}
func commandBmap(cfg *Config) error {
	if(cfg.Next== "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20") {
		fmt.Println("You are on the first page!\n")
		return nil
	}
	locationPage, err := pokeapi.FetchLocationPage(cfg.Prev)
	if(err!=nil){
		return err
	}
	cfg.Next = locationPage.Next
	cfg.Prev = locationPage.Previous
	for _, value := range locationPage.Results {
		fmt.Println(value.Name)
	}
	return nil
}