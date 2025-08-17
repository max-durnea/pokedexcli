package pokeapi
import (
	"encoding/json"
	"net/http"
	"io"
	"time"
	"github.com/max-durnea/pokedexcli/internal/pokecache"
)
type LocationPage struct {
	Results []Location 
	Next string
	Previous string
}
type Location struct {
	Name string
}
var cache = pokecache.NewCache(5*time.Second)

func FetchLocationPage(url string) (*LocationPage, error) {
	bytes, ok := cache.Get(url)
	if ok {
		locationPage :=LocationPage{}
		err := json.Unmarshal(bytes, &locationPage)
		if(err!=nil){
			return nil, err
		}
		cache.Add(url,bytes)
		return &locationPage, nil
	}

	resp, err := http.Get(url)
	if(err!=nil){
		return nil, err
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	locationPage := LocationPage{}
	err = json.Unmarshal(body, &locationPage)
	if(err!=nil){
		return nil, err
	}
	cache.Add(url,body)
	return &locationPage, nil
}