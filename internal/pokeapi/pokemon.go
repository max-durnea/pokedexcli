package pokeapi
import (
	"encoding/json"
	"net/http"
	"io"
)
type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func FetchPokemonInfo(url string) (PokemonInfo, error){
	bytes, ok := cache.Get(url)
	if ok {
		pokemonInfo:=PokemonInfo{}
		err := json.Unmarshal(bytes, &pokemonInfo)
		if(err!=nil){
			return PokemonInfo{}, err
		}
		cache.Add(url,bytes)
		return pokemonInfo, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(body,&pokemonInfo)
	if err != nil{
		return PokemonInfo{},err
	}
	cache.Add(url, body)
	return pokemonInfo,nil
}