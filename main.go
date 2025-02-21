package main

import (
	"fmt"
	// "github.com/bogem/id3v2/v2"
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"path/filepath"
)

const musicDir = "/home/fiammetta/Music/"

const (
	scanning = "scanning"
	ready    = "ready"
)

type Meta struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Year     int    `json:"year"`
	Duration int    `json:"duration"`
	Cover    string `json:"cover"` // base64 encoded image
}

type File struct {
	Uid      string  `json:"uid"`
	Name     string  `json:"name"`
	Size     float32 `json:"size"`
	Metadata Meta    `json:"metadata"`
}

var files []File
var filenames []string

func walkHandler(path string, d fs.DirEntry, err error) error {
	info, err := d.Info()

	if err != nil {
		fmt.Printf("err: %s", err)
		return nil
	}

	name := info.Name()

	size := float32(info.Size()) / 1000_000

	hash := sha256.New()
	hash.Write([]byte(name))
	hashed := hash.Sum(nil)
	uid := hex.EncodeToString(hashed)

	files = append(files, File{
		Name: name, Size: size, Uid: uid,
	})

	return nil
}

func scanner(status chan (string)) error {

	// status <- scanning

	err := filepath.WalkDir(musicDir, walkHandler)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	fmt.Println((files))

	// status <- ready
	status <- ready
	return nil
}

func stats(status chan (string)) {}

func main() {
	status := make(chan string)

	go scanner(status)

	go stats(status)

	<-status
}
