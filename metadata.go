package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func setCover(filePath string, coverPath string) error {
	cover := fmt.Sprintf("cover=%s", coverPath)

	cmd := exec.Command("tageditor",
		"-s",
		cover,
		"--max-padding", "100000",
		"-f", filePath,
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return err
	}

	infoLog(string(out))

	return nil
}

func editMetadata(filePath string, m Meta) error {

	title := fmt.Sprintf("title=%s", m.Title)
	artist := fmt.Sprintf("artist=%s", m.Artist)
	album := fmt.Sprintf("album=%s", m.Album)
	genre := fmt.Sprintf("genre=%s", m.Genre)
	year := fmt.Sprintf("year=%d", m.Year)

	cmd := exec.Command("tageditor",
		"-s",
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

func getCover(filePath string) ([]byte, error) {
	filePath = musicDir + filePath
	infoLog(filePath)
	cmd := exec.Command(
		"tageditor",
		"-e",
		"cover",
		"-f",
		filePath,
	)

	out, err := cmd.Output()

	if err != nil {
		infoLog(string(out))
		return nil, err
	}

	return out, nil
}

func gatherMetadata(filePath string) (Meta, error) {
	metadata := Meta{}

	cmd := exec.Command(
		"ffprobe",
		filePath,
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return Meta{}, err
	}

	input := string(out)

	var dateRegex *regexp.Regexp

	// Regular expressions for each field
	titleRegex := regexp.MustCompile(`(?i)title\s*:\s*(.+)`)
	albumRegex := regexp.MustCompile(`(?i)album\s*:\s*(.+)`)
	artistRegex := regexp.MustCompile(`(?i)artist\s*:\s*(.+)`)
	if strings.HasSuffix(filePath, "flac") {
		dateRegex = regexp.MustCompile(`(?i)date\s*:\s*(.+)`)
	} else {
		dateRegex = regexp.MustCompile(`(?i)year\s*:\s*(.+)`)
	}
	genreRegex := regexp.MustCompile(`(?i)genre\s*:\s*(.+)`)

	// Extract fields
	metadata.Title = extractField(titleRegex, input)
	metadata.Album = extractField(albumRegex, input)
	metadata.Artist = extractField(artistRegex, input)
	metadata.Year, _ = strconv.Atoi(extractField(dateRegex, input))
	metadata.Genre = extractField(genreRegex, input)

	return metadata, nil
}

func extractField(regex *regexp.Regexp, input string) string {
	match := regex.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1] // Return the captured group
	}
	return "" // Return empty string if no match
}
