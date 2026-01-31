package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const baseURL = "https://api.themoviedb.org/3/tv/"

type seasonList struct {
	Name        string `json:"name"`
	NumEpisodes int    `json:"number_of_episodes"`
	NumSeasons  int    `json:"number_of_seasons"`
	ID          int    `json:"id"`
	Seasons     []struct {
		AirDate      string `json:"air_date"`
		NumEpisodes  int    `json:"episode_count"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
}

type episodeList struct {
	Episodes []struct {
		AirDate    string `json:"air_date"`
		EpisodeNum int    `json:"episode_number"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Summary    string `json:"overview"`
	}
}

func main() {
	tmdbID := os.Args

	seasonListURL := baseURL + tmdbID[1]
	seasons, err := getJSON[seasonList](seasonListURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, season := range seasons.Seasons {
		episodeListURL := seasonListURL + "/season/" + strconv.Itoa(season.SeasonNumber)
		episodes, err := getJSON[episodeList](episodeListURL)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		episodeNum := 1
		for _, episode := range episodes.Episodes {
			err := writeNFO(seasons.Name, season.SeasonNumber, episodeNum, episode.Name, episode.ID)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			episodeNum += 1
		}
	}
	fmt.Println("Successfully wrote all NFO files.")
}
