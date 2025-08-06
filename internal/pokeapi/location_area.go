package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) LARequest(currentID int) (string, error) {
	urlString := baseURL + "/location-area/" + strconv.Itoa(currentID) + "/"

	if val, ok := c.cache.Get(urlString); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			rS := fmt.Sprintf("Error reading data: %v", err)
			return rS, err
		}
		locationAreaName := string(locationArea.Name)
		return locationAreaName, nil
	}

	resp, err := http.Get(urlString)
	if err != nil {
		rS := fmt.Sprintf("Error contacting PokeAPI: %v", err)
		return rS, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		rS := fmt.Sprintf("Error reading response: %v", err)
		return rS, err
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		rS := fmt.Sprintf("Error reading data: %v", err)
		return rS, err
	}
	locationAreaName := string(locationArea.Name)
	c.cache.Add(urlString, data)
	return locationAreaName, nil
}
