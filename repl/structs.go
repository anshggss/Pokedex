package repl

import "github.com/anshggss/pokedex/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commandRegistry map[string]cliCommand
	mapURL          string
	mapNextURL      string
	mapPrevURL      string
	mapCache        pokecache.Cache
	exploreCache    pokecache.Cache
	cmdArgs         []string
	exploreURL      string
	catchPokemon    string
	catchURL        string
	Pokedex         map[string]CatchPokemonInfo
}

type mapResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type exploreResponse struct {
	Pokemons []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Info PokeInfo `json:"pokemon"`
}

type PokeInfo struct {
	Name string `json:"name"`
}

type CatchPokemonInfo struct {
	Name   string `json:"name"`
	BaseXP int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}
