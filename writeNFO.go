package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writeNFO(seriesName, episodeName, summary string, season, episode, id int) error {
	type UniqueID struct {
		Type    string `xml:"type,attr"`
		Default bool   `xml:"default,attr"`
		Value   int    `xml:",chardata"`
	}

	type metadata struct {
		XMLName xml.Name `xml:"tvshow"`

		ID      UniqueID `xml:"uniqueid"`
		Title   string   `xml:"title"`
		Season  int      `xml:"season"`
		Episode int      `xml:"episode"`
		Plot    string   `xml:"plot"`
	}

	seasonNum := "0" + strconv.Itoa(season)
	if season > 9 {
		seasonNum = "" + strconv.Itoa(season)
	}
	episodeNum := "0" + strconv.Itoa(episode)
	if episode > 9 {
		episodeNum = "" + strconv.Itoa(episode)
	}
	fileName := "/data/tor/NFOs/" + seriesName + " - s" + seasonNum + "e" + episodeNum + ".nfo"
	fileName = strings.NewReplacer(":", "").Replace(fileName)

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
