package pokeapi
import (
	"encoding/json"
	"net/http"
	"io"
)
type LocationPage struct {
	Results []Location 
	Next string
	Previous string
}
type Location struct {
	Name string
}
func FetchLocationPage(url string) (*LocationPage, error) {
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
	return &locationPage, nil
}