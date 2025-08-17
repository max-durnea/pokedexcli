package pokeapi
import (
	"encoding/json"
	"net/http"
	"io"
)

type PokemonList struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func FetchPokemonList(url string) (*PokemonList, error){
	bytes, ok := cache.Get(url)
	if ok {
		pokemons:=PokemonList{}
		err := json.Unmarshal(bytes, &pokemons)
		if(err!=nil){
			return nil, err
		}
		cache.Add(url,bytes)
		return &pokemons, nil
	}
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	pokemons := PokemonList{}
	err = json.Unmarshal(body,&pokemons)
	if err != nil{
		return nil, err
	}
	cache.Add(url, body)
	return &pokemons, nil
}


