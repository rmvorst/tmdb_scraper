package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rmvorst/tmdb_scraper/internal/api"
	"github.com/rmvorst/tmdb_scraper/internal/env"
	"github.com/rmvorst/tmdb_scraper/internal/writers"
)

const baseURL = "https://api.themoviedb.org/3/tv/"

func logFatal(err error) {
	log.Fatalf("Error: %s\n", err)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Example usage: go run . <show_id> <num_episodes_season_1> <num_episodes_season_2> ...")
	}

	err := env.SetupEnv()
	if err != nil {
		logFatal(err)
	}

	envPath, err := env.GetEnvPath()
	if err != nil {
		logFatal(err)
	}
	envFPath := filepath.Join(envPath, ".env")
	godotenv.Load(envFPath)

	cfg := env.EnvConfig{
		ApiKey:  os.Getenv("API_KEY"),
		RootNFO: os.Getenv("NFO_ROOT"),
	}

	err = writers.ResetNFO(cfg.RootNFO)
	if err != nil {
		logFatal(err)
	}

	// Define flags (name, default, description)
	seasonFlag := flag.Int("season", 0, "Season Number")
	episodeFlag := flag.Int("episode", 0, "Episode Number")
	specialsFlag := flag.Bool("specials", false, "Specials Boolean")
	flag.Parse()
	args := flag.Args()

	// Get the tmdb-id and episode number overrides from args
	tmdbID := args[0]
	episodeNumOverride := args[1:]

	seasonListURL := baseURL + tmdbID
	seasons, err := getJSON[api.SeasonList](seasonListURL, cfg)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	seasonEpisodeNums, err := strconv.Atoi(episodeNumOverride[0])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Set loop initial conditions
	seasonNum := 1
	episodeNum := 1
	numSeasons := len(episodeNumOverride)
	overrideIDX := 0

	if *specialsFlag {
		seasonNum -= 1
	}

	// Start the loop
seasonLoop:
	for _, season := range seasons.Seasons {
		episodeListURL := seasonListURL + "/season/" + strconv.Itoa(season.SeasonNumber)
		episodes, err := getJSON[api.EpisodeList](episodeListURL, cfg)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		for _, episode := range episodes.Episodes {
			if checkFlag(seasonNum, seasonFlag) {
				if checkFlag(episodeNum, episodeFlag) {
					err := writers.WriteNFO(seasons.Name, episode.Name, episode.Summary, cfg.RootNFO, seasonNum, episodeNum, episode.ID)
					if err != nil {
						log.Fatalf("Error: %v\n", err)
					}
				}
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

func checkFlag(num int, flag *int) bool {
	switch *flag {
	case 0, num:
		return true
	default:
		return false
	}
}
