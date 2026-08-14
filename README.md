# tmdb_scraper

## Description

This tool is intended to download episode information from the movie db (tmdb) and write the info into an NFO file for use in Kodi or Jellyfin. Optional flags allow for a user to specify the season number or episode number (or both) for a more targetted execution.

## Motivation

tmdb_scraper was built to fix issues that arise from Jellyfin's automated metadata creation algorithms. Some seasons, especially with anime, are not entered into tmdb correctly. This creates some issues with a user's expected folder structure once adding the shows to Jellyfin. You either comply with tmdb and let yourself be confuesed by what season and episode you are really on, or you define the metadata yourself manually. To automate the manual fix as best as possible, this tool was written to create the NFO files which are used by Jellyfin to correct the metadata that was scraped incorrectly. 

## Quick Start

### Installation
`go install github.com/rmvorst/tmdb_scraper@latest`

### Prerequisites
User must already have a tmdb api key. This should be stored in the .env file as API_KEY.
The API Key can be accessed through the TMDB user account settings: https://www.themoviedb.org/settings/api

## Usage

Optional flags are:
- '--season' - Specify the season wanted to be grabbed and written to an NFO. Follow this flag with the season number as an int.
- '--episode' - Specify the episode wanted to be grabbed and written to an NFO Follow this flag with the episode number as an int.
- '--specials' - A boolean flag that specifies there are specials listed in the TMDB database. They will be treated as Season 0. No value needs to follow this flag.
- '--output' - A string flag that allows a user to specify the output folder. If not included, the NFOROOT in the user's config is used. The string representation of the folderpath to the desired output folder should follow this flag.
- '--debug' - Run debug lines that describe current environment, as well as the info of the show being fetched. No value needs to follow this flag.

## Examples:

### Read and write to an NFO the metadata of all episodes from the show with tmdb-id of 209867

`tmdb_scraper 209867 28 10`

209867 is the tmdb-id of the show being accessed

28 is the number of episodes in season 1

10 is the number of episodes in season 2

### Read and write to an NFO the metadata of all episodes from season 1 of the show with tmdb-id of 209867

`tmdb_scraper --season 1 209867 28 10`

209867 is the tmdb-id of the show being accessed

28 is the number of episodes in season 1

10 is the number of episodes in season 2

### Read and write to an NFO the metadata of episode 10 from season 1 of the show with tmdb-id of 209867

`tmdb_scraper --season 1 --episode 10 209867 28 10`

209867 is the tmdb-id of the show being accessed

28 is the number of episodes in season 1

10 is the number of episodes in season 2

### Read and write to an NFO the metadata of episode 10 from season 3 of the show with tmdb-id of 82684 to a defined location. There are specials in this series.

`tmdb_scraper --season 3 --episode 10 --specials --output "/path/to/output" 82684 16 24 24 24 24`

82684 is the tmdb-id of the show being accessed

16 is the number of specials

24 is the number of episodes in season 1

24 is the number of episodes in season 2

24 is the number of episodes in season 3

24 is the number of episodes in season 4

## Contributing

### Clone the repo

```bash
git clone https://github.com/rmvorst/tmdb_scraper
cd tmdb_scraper
```

### Build the compiled binary

```bash
go build
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## Output

The output is an NFO file. For info, see the [Jellyfin Documentation](https://jellyfin.org/docs/general/server/metadata/nfo/)

## ENV

The .env file consists of two fields: `API_KEY` and `NFO_ROOT`.

The `API_KEY` should be the users tmdb API Key.

The `NFO_ROOT` should be the folder the user wants to save the created NFO files.

If a .env file with the correct fields does not exist, this tool will guide the creation of it automatically in the $HOME/.config/tmdb_scraper folder.
