package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/max-durnea/pokedexcli/internal/pokeapi"
	"math/rand"
)
type command struct {
	name string
	description string
	callback func(cfg *Config, args []string) error
}
type Config struct {
	Next string
	Prev string
}
var Commands map[string]command
var Pokemons map[string]pokeapi.PokemonInfo
func main() {
	Pokemons = map[string]pokeapi.PokemonInfo{}
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
		"explore": {
			name: "explore",
			description: "Lists pokemons from a location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Throw a pokeball at the specified pokemon",
			callback: commandCatch,
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
			val.callback(&locationConfig,cleaned_input[1:])
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

func commandExit(cfg *Config, args []string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!");
	return nil
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range Commands{
		fmt.Printf("%v: %v\n",value.name, value.description)
	}
	return nil
}

func commandMap(cfg *Config, args []string) error {
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
func commandBmap(cfg *Config, args []string) error {
	if(cfg.Next== "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20") {
		fmt.Println("You are on the first page!")
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

func commandExplore(cfg *Config, args []string) error {
	fmt.Printf("Exploring %v...\n",args[0])
	url:=fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v",args[0])
	pokemons,err:=pokeapi.FetchPokemonList(url)
	if err!=nil {
		return err
	}
	fmt.Println("Found Pokemon:\n")
	for _, value := range pokemons.PokemonEncounters{
		fmt.Printf(" - %v\n",value.Pokemon.Name)
	}
	fmt.Println()
	return nil
}

func commandCatch(cfg *Config, args []string) error {
	url:=fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v",args[0])
	pokemon,err:=pokeapi.FetchPokemonInfo(url)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n",pokemon.Name)
	//fmt.Println(pokemon)
	maxExp := 300
	chance := 1.0 -float64(pokemon.BaseExperience/maxExp)
	if(chance<0.05){
		chance=0.05
	}
	roll := rand.Float64()
	if roll < chance {
		fmt.Printf("%v was caught!\n",pokemon.Name)
		Pokemons[args[0]]=pokemon
	}else{
		fmt.Printf("%v escaped!\n",pokemon.Name)
	}
	return nil
}

func commandInspect(cfg *Config, args []string) error{
	
}