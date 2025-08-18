# Pokedex CLI

A command-line interface Pokédex application built in Go that allows you to explore Pokémon locations, catch Pokémon, and manage your collection using the PokéAPI.

## Features

- **Location Exploration**: Browse through Pokémon locations and discover which Pokémon can be found in each area
- **Pokémon Catching**: Attempt to catch Pokémon with a probability-based system
- **Collection Management**: Inspect caught Pokémon and view your complete Pokédex
- **Caching System**: Built-in HTTP response caching for improved performance
- **Interactive CLI**: Simple command-based interface with help system

## Installation

1. Clone the repository:
```bash
git clone https://github.com/max-durnea/pokedexcli.git
cd pokedexcli
```

2. Initialize Go modules:
```bash
go mod init github.com/max-durnea/pokedexcli
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

## Project Structure

```
pokedexcli/
├── main.go                    # Main application entry point and command handlers
├── internal/
│   ├── pokeapi/
│   │   ├── locations.go       # Location-related API functions
│   │   ├── pokemon.go         # Individual Pokémon data fetching
│   │   └── pokemons.go        # Pokémon list fetching for locations
│   └── pokecache/
│       ├── caching.go         # HTTP response caching implementation
│       └── caching_test.go    # Cache functionality tests
└── README.md
```

## Commands

Once the application is running, you can use the following commands:

- **`help`** - Display all available commands and their descriptions
- **`exit`** - Exit the Pokédex application
- **`map`** - List the next 20 location areas
- **`mapb`** - List the previous 20 location areas
- **`explore <location>`** - Explore a specific location and see available Pokémon
- **`catch <pokemon>`** - Attempt to catch a specified Pokémon
- **`inspect <pokemon>`** - View details of a caught Pokémon
- **`pokedex`** - List all your caught Pokémon

## Usage Examples

```bash
Pokedex > help
Welcome to the Pokedex!
Usage:

exit:  Exit the Pokedex
help:  Displays a help message
map:  Lists the next 20 locations everytime it's called
...

Pokedex > map
# Lists location areas

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:

 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
 -hp: 35
 -attack: 55
 -defense: 40
 -special-attack: 50
 -special-defense: 50
 -speed: 90
Types:
 - electric

Pokedex > pokedex
Your Pokedex:
 - pikachu
```

## Technical Details

### Caching System
The application implements a thread-safe cache with TTL (Time To Live) functionality:
- **Cache Duration**: 5 seconds for API responses
- **Automatic Cleanup**: Background goroutine removes expired entries
- **Concurrency Safe**: Uses mutexes to handle concurrent access

### Pokémon Catching Mechanics
The catch system uses a probability-based approach:
- Success rate is inversely related to the Pokémon's base experience
- Higher experience Pokémon are harder to catch
- Minimum catch rate of 5% for very powerful Pokémon

### API Integration
The application integrates with the [PokéAPI](https://pokeapi.co/) to fetch:
- Location area data with pagination
- Pokémon encounter information for specific locations
- Detailed Pokémon statistics and information

## Data Structures

### Core Types
- **`Config`**: Manages pagination state for location browsing
- **`command`**: Defines CLI command structure with callbacks
- **`PokemonInfo`**: Contains detailed Pokémon data (stats, types, etc.)
- **`LocationPage`**: Handles paginated location data
- **`PokemonList`**: Manages Pokémon encounters in locations

## Dependencies

- **Standard Library**: `fmt`, `strings`, `bufio`, `os`, `encoding/json`, `net/http`, `io`, `time`, `sync`, `math/rand`
- **External API**: [PokéAPI](https://pokeapi.co/) for Pokémon data

## Testing

Run the cache tests:
```bash
go test ./internal/pokecache
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source. Please check the repository for license details.

## Acknowledgments

- Data provided by [PokéAPI](https://pokeapi.co/)
- Inspired by the classic Pokémon games