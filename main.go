package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const baseURL = "https://api.themoviedb.org/3/tv/"
const envPath = "/data/projects/tmdb_scraper/.env"

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

type envConfig struct {
	apiKey  string
	rootNFO string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Example usage: go run . <show_id> <num_episodes_season_1> <num_episodes_season_2> ...")
	}
	godotenv.Load(envPath)

	cfg := envConfig{
		apiKey:  os.Getenv("API_KEY"),
		rootNFO: os.Getenv("NFO_ROOT"),
	}
	err := os.RemoveAll(cfg.rootNFO)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	err = os.Mkdir(cfg.rootNFO, 0755)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	tmdbID := os.Args[1]
	episodeNumOverride := os.Args[2:]
	numSeasons := len(episodeNumOverride)
	overrideIDX := 0

	seasonListURL := baseURL + tmdbID
	seasons, err := getJSON[seasonList](seasonListURL, cfg)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	seasonEpisodeNums, err := strconv.Atoi(episodeNumOverride[overrideIDX])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	seasonNum := 1
	episodeNum := 1
seasonLoop:
	for _, season := range seasons.Seasons {
		episodeListURL := seasonListURL + "/season/" + strconv.Itoa(season.SeasonNumber)
		episodes, err := getJSON[episodeList](episodeListURL, cfg)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		for _, episode := range episodes.Episodes {
			err := writeNFO(seasons.Name, seasonNum, episodeNum, episode.Name, episode.ID)
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
			episodeNum += 1
			if episodeNum > seasonEpisodeNums {
				overrideIDX += 1
				if overrideIDX > numSeasons-1 {
					break seasonLoop
				}
				seasonEpisodeNums, err = strconv.Atoi(episodeNumOverride[overrideIDX])
				if err != nil {
					log.Fatalf("Error: %v\n", err)
				}
				seasonNum += 1
				episodeNum = 1
			}
		}
	}
	fmt.Println("Successfully wrote all NFO files.")
}
