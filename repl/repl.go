package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/anshggss/pokedex/internal/pokecache"
)

func Repl() {

	// A new scanner to read the repl commands
	sc := bufio.NewScanner(os.Stdin)

	// Registry for commands
	commandReg := map[string]cliCommand{

		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": cliCommand{
			name:        "help",
			description: "About the REPL",
			callback:    commandHelp,
		},
		"map": cliCommand{
			name:        "map",
			description: "Lists 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Lists 20 locations on the map",
			callback:    commandMapB,
		},
		"explore": cliCommand{
			name:        "explore",
			description: "Lists all the pokemons in a location area",
			callback:    commandExplore,
		},
		"catch": cliCommand{
			name:        "catch",
			description: "Catches a given Pokemon",
			callback:    commandCatch,
		},
		"inspect": cliCommand{
			name:        "inspect",
			description: "Inspects a given Pokemon",
			callback:    commandInspect,
		},
	}

	// Config that stores the state used by all the commands
	dexConfig := &config{
		commandRegistry: commandReg,
		mapURL:          "https://pokeapi.co/api/v2/location-area/",
		mapNextURL:      "",
		mapPrevURL:      "",
		mapCache:        pokecache.NewCache(10 * time.Second),
		exploreCache:    pokecache.NewCache(10 * time.Second),
		cmdArgs:         nil,
		exploreURL:      "https://pokeapi.co/api/v2/location-area/",
		catchPokemon:    "",
		catchURL:        "https://pokeapi.co/api/v2/pokemon/",
		Pokedex:         map[string]CatchPokemonInfo{},
		inspectPokemon:  "",
	}

	for {

		fmt.Print("Pokedex > ")

		for sc.Scan() {

			input := sc.Text()
			if len(input) == 0 {
				fmt.Print("Pokedex > ")
				continue
			}

			cleanedInput := cleanInput(input)

			command, ok := commandReg[cleanedInput[0]]
			if !ok {
				fmt.Println("Unknown command")
				fmt.Print("Pokedex > ")
				continue
			}
			if command.name == "explore" {
				dexConfig.cmdArgs = cleanedInput[1:]
			}
			if command.name == "catch" {
				dexConfig.catchPokemon = cleanedInput[1]
			}
			if command.name == "inspect" {
				dexConfig.inspectPokemon = cleanedInput[1]
			}

			err := command.callback(dexConfig)
			if err != nil {
				fmt.Printf("Error with callback: %s\n", err.Error())
			}

			fmt.Print("Pokedex > ")
		}

		if err := sc.Err(); err != nil {
			fmt.Println("Something bad happened")
		}
	}
}
