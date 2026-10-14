# Pokedex

A simple command-line Pokedex built in Go. It uses the [PokéAPI](https://pokeapi.co/) to browse locations and explore the Pokemon found in each area.

## Requirements

- Go 1.26.2 or later
- Internet access for PokéAPI requests

## Run

```bash
go run .
```

## Commands

Inside the REPL, use:

| Command | Description |
| --- | --- |
| `help` | Display available commands |
| `map` | Show the next page of locations |
| `mapb` | Show the previous page of locations |
| `explore <location-area>` | List Pokemon in a location area |
| `exit` | Close the Pokedex |

## Test

```bash
go test ./...
```
