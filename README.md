# tmdb_scraper

## Installation
`go install github.com/rmvorst/tmbd_scraper@latest`

## Prerequisites
User must already have a tmdb api key. This should be stored in the .env file as API_KEY

## Usage

`tmdb_scraper 209867 28 10`

209867 is the tmdb-id of the show being accessed

28 is the number of episodes in season 1

10 is the number of episodes in season 2

## Description

This tool is intended to download episode information from the movie db (tmdb) and write the info into an NFO file for use in Kodi or Jellyfin. Currently, this is quite dumb and will pull info for all seasons and all episodes. Future updates will account for this to make it more seamless of an experience

## Output

The output is an NFO file. For info, see the [Jellyfin Documentation](https://jellyfin.org/docs/general/server/metadata/nfo/)

## ENV

The .env file consists of two fields: `API_KEY` and `NFO_ROOT`.

The `API_KEY` should be the users tmdb API Key.

The `NFO_ROOT` should be the folder the user wants to save the created NFO files.

If a .env file with the correct fields does not exist, this tool will guide the creation of it automatically in the $HOME/.config/tmdb_scraper folder.
