package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

var exts = []string{
	"m4a", "mp3", "flac", "ogg", "opus",
}

var id = ""

var fileChan chan (File)

func walkHandler(path string, d fs.DirEntry, err error) error {
	info, err := d.Info()

	temp := strings.Split(path, ".")

	ext := temp[len(temp)-1]

	if !slices.Contains(exts, ext) {
		return nil
	}

	if err != nil {
		fmt.Printf("err: %s", err)
		return nil
	}

	if d.IsDir() {
		return nil
	}

	name := info.Name()

	size := float32(info.Size()) / 1000_000

	hash := sha256.New()
	hash.Write([]byte(name))
	hashed := hash.Sum(nil)

	uid := hex.EncodeToString(hashed)

	id = uid

	metadata, err := gatherMetadata(path)

	file := File{
		Name: name, Size: size, Uid: uid, Metadata: metadata,
	}

	noChan := make(chan DbResult, 1)

	event := DbEvent{
		eventType:  Insert,
		data:       file,
		resultChan: noChan,
	}

	go dispatch(event)

	return nil
}

func scanner() {
	err := filepath.WalkDir(musicDir, walkHandler)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

}
