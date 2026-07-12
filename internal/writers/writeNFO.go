package writers

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ResetNFO(fpath string) error {
	err := os.MkdirAll(fpath, 0755)
	if err != nil {
		return fmt.Errorf("Error: Cannot make NFO Directory")
	}

	return nil
}

func WriteNFO(seriesName, episodeName, summary, fpath string, season, episode, id int) error {
	// Create and clean the NFO file name
	seasonNum := "0" + strconv.Itoa(season)
	if season > 9 {
		seasonNum = "" + strconv.Itoa(season)
	}
	episodeNum := "0" + strconv.Itoa(episode)
	if episode > 9 {
		episodeNum = "" + strconv.Itoa(episode)
	}
	fname := seriesName + " - s" + seasonNum + "e" + episodeNum + ".nfo"
	clean_fname := cleanFname(fname)
	fullPath := filepath.Join(fpath, clean_fname)

	// Set up data structure
	showMetadata := metadata{
		ID: UniqueID{
			Type:    "tmdb",
			Default: true,
			Value:   id,
		},
		Title:   episodeName,
		Season:  season,
		Episode: episode,
		Plot:    summary,
	}

	// Marshal data to xml format
	xmlBytes, err := xml.MarshalIndent(showMetadata, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling XML")
		return err
	}

	// Create and write xml formatted data to the NFO file
	xmlOutput := []byte(xml.Header + string(xmlBytes) + "\n")
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Error creating NFO")
		return err
	}
	defer file.Close()

	_, err = file.Write(xmlOutput)
	if err != nil {
		fmt.Println("Error writing NFO")
		return err
	}

	return nil
}

func cleanFname(fname string) string {
	fname = strings.NewReplacer(":", "").Replace(fname)

	return fname
}
