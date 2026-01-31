package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

func writeNFO(name string, season int, episode int, id int) error {
	type metadata struct {
		ID      int    `xml:"tvshow>uniqueid"`
		Type    string `xml:"tvshow>type,attr"`
		Default bool   `xml:"tvshow>default,attr"`
		Title   string `xml:"tvshow>title"`
		Season  int    `xml:"tvshow>season"`
		Episode int    `xml:"tvshow>episode"`
	}
	seasonNum := "0" + strconv.Itoa(season)
	if season > 9 {
		seasonNum = "" + strconv.Itoa(season)
	}
	episodeNum := "0" + strconv.Itoa(episode)
	if episode > 9 {
		seasonNum = "" + strconv.Itoa(episode)
	}
	fileName := name + " - s" + seasonNum + "e" + episodeNum + ".nfo"

	showMetadata := metadata{
		ID:      id,
		Type:    "tmdb",
		Default: true,
		Title:   name,
		Season:  season,
		Episode: episode,
	}
	xmlBytes, err := xml.MarshalIndent(showMetadata, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling XML")
		return err
	}

	xmlOutput := []byte(xml.Header + string(xmlBytes) + "\n")
	file, err := os.Create(fileName)
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
