# tmdb_scraper

This tool is intended to download episode information from the movie db (tmdb) and write the info into an NFO file for use in Kodi or Jellyfin. Optional flags allow for a user to specify the season number or episode number (or both) for a more targetted execution.

## Motivation

tmdb_scraper was built to fix issues that arise from Jellyfin's automated metadata creation algorithms. Some seasons, especially with anime, are not entered into tmdb correctly. This creates some issues with a user's expected folder structure once adding the shows to Jellyfin. You either comply with tmdb and let yourself be confuesed by what season and episode you are really on, or you define the metadata yourself manually. To automate the manual fix as best as possible, this tool was written to create the NFO files which are used by Jellyfin to correct the metadata that was scraped incorrectly. 

## Installation
`go install github.com/rmvorst/tmbd_scraper@latest`

## Prerequisites
User must already have a tmdb api key. This should be stored in the .env file as API_KEY

## Usage

Optional flags are:
- '--season' - specify the season wanted to be grabbed and written to an NFO
- '--episode' - specify the episde wanted to be grabbed and written to an NFO

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

## Output

The output is an NFO file. For info, see the [Jellyfin Documentation](https://jellyfin.org/docs/general/server/metadata/nfo/)

## ENV

The .env file consists of two fields: `API_KEY` and `NFO_ROOT`.

The `API_KEY` should be the users tmdb API Key.

The `NFO_ROOT` should be the folder the user wants to save the created NFO files.

If a .env file with the correct fields does not exist, this tool will guide the creation of it automatically in the $HOME/.config/tmdb_scraper folder.
