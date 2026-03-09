package api

type SeasonList struct {
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

type EpisodeList struct {
	Episodes []struct {
		AirDate    string `json:"air_date"`
		EpisodeNum int    `json:"episode_number"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Summary    string `json:"overview"`
	}
}
