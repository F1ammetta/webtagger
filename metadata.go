package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func editMetadata(filePath string, coverPath string, m Meta) error {

	cover := fmt.Sprintf("cover=%s", coverPath)
	title := fmt.Sprintf("title=%s", m.Title)
	artist := fmt.Sprintf("artist=%s", m.Artist)
	album := fmt.Sprintf("album=%s", m.Album)
	genre := fmt.Sprintf("genre=%s", m.Genre)
	year := fmt.Sprintf("year=%d", m.Year)

	cmd := exec.Command("tageditor",
		"-s",
		cover,
		title,
		artist,
		album,
		genre,
		year,
		"--max-padding", "100000",
		"-f", filePath,
	)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return err
	}

	output := string(out)

	if !strings.Contains(output, "applied") {
		return errors.New(fmt.Sprintf("An unexpected error occurred: %s", output))
	}

	return nil
}

const (
	DATE   = 0
	GENRE  = 1
	ALBUM  = 2
	ARTIST = 3
	TITLE  = 4
)

func gatherMetadata(filePath string) (Meta, error) {
	metadata := Meta{}

	cmd := exec.Command(
		"tageditor",
		"-g", "-n",
		"title", "artist",
		"album", "genre",
		"year", "-f",
		filePath,
	)

	out, err := cmd.Output()

	if err != nil {
		return Meta{}, err
	}

	output := string(out)

	lines := strings.Split(output, "\n")

	lines = lines[2 : len(lines)-1]

	for i, line := range lines {
		line := strings.TrimSpace(line)
		vals := strings.Fields(line)

		fmt.Println(vals)

		val := strings.Join(vals[1:], " ")

		switch i {
		case DATE:
			val = strings.Join(vals[2:], " ")
			metadata.Year, _ = strconv.Atoi(val)
		case GENRE:
			metadata.Genre = val
		case ALBUM:
			metadata.Album = val
		case ARTIST:
			metadata.Artist = val
		case TITLE:
			metadata.Title = val
		}

		if i == TITLE {
			break
		}
	}

	return metadata, nil
}
