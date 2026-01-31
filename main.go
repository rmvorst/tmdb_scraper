package main

import (
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

	seasonListURL := baseURL + tmdbID[0]
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

		for _, episode := range episodes.Episodes {
			seasonNum := "0" + strconv.Itoa(season.SeasonNumber)
			if season.SeasonNumber > 9 {
				seasonNum = "" + strconv.Itoa(season.SeasonNumber)
			}
			episodeNum := "0" + strconv.Itoa(episode.EpisodeNum)
			if episode.EpisodeNum > 9 {
				seasonNum = "" + strconv.Itoa(episode.EpisodeNum)
			}
			fileName := seasons.Name + " - s" + seasonNum + "e" + episodeNum + ".nfo"
		}
	}

}
