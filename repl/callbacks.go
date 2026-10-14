package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(c *config) error {

	var url string

	if c.mapNextURL == "" {
		url = c.mapURL
	} else {
		url = c.mapNextURL
	}

	val, ok := c.mapCache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		val, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		c.mapCache.Add(url, val)
	}

	var decoded mapResponse

	err := json.Unmarshal(val, &decoded)
	if err != nil {
		return err
	}

	c.mapNextURL = decoded.Next
	c.mapPrevURL = decoded.Previous
	for _, v := range decoded.Results {
		fmt.Println(v.Name)
	}

	return nil
}
func commandMapB(c *config) error {

	var url string
	if c.mapPrevURL == "" {
		url = c.mapURL
	} else {
		url = c.mapPrevURL
	}

	val, ok := c.mapCache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		val, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println("Didn't use the cache bruh")
		c.mapCache.Add(url, val)
	}

	var decoded mapResponse

	err := json.Unmarshal(val, &decoded)
	if err != nil {
		return err
	}

	c.mapNextURL = decoded.Next
	c.mapPrevURL = decoded.Previous
	for _, v := range decoded.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandExplore(c *config) error {

	if len(c.cmdArgs) == 0 {
		return fmt.Errorf("Dude provide some arguments")
	}

	cached, ok := c.exploreCache.Get(c.cmdArgs[0])
	if !ok {
		fullURL := c.exploreURL + c.cmdArgs[0] + "/"
		res, err := http.Get(fullURL)
		if err != nil {
			return err
		}

		cached, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		c.exploreCache.Add(c.cmdArgs[0], cached)

		res.Body.Close()
	}

	var decoded exploreResponse
	if err := json.Unmarshal(cached, &decoded); err != nil {
		return err
	}

	for _, pokemon := range decoded.Pokemons {
		fmt.Println(pokemon.Info.Name)
	}

	return nil
}

func commandCatch(c *config) error {

	if c.catchPokemon == "" {
		return fmt.Errorf("Enter a pokemon name")
	}

	fullURL := c.catchURL + c.catchPokemon
	res, err := http.Get(fullURL)

	if err != nil {
		return err
	}
	if int(res.StatusCode/100) != 2 {
		return fmt.Errorf("Invalid Pokemon name")
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var decoded CatchPokemonInfo
	if err := decoder.Decode(&decoded); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", decoded.Name)

	chance := int(decoded.BaseXP / 20)
	myNumber := rand.Intn(chance)
	finNumber := rand.Intn(chance)

	if myNumber != finNumber {
		fmt.Printf("%s escaped!\n", c.catchPokemon)
	} else {
		fmt.Printf("%s was caught!\n", c.catchPokemon)
		c.Pokedex[decoded.Name] = decoded
	}

	return nil
}
