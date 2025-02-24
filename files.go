package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ostafen/clover"
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

	fileChan <- File{
		Name: name, Size: size, Uid: uid, Metadata: metadata,
	}

	return nil
}

func handleInserts() {
	db, err := clover.Open("./data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := 0

	for {
		file := <-fileChan

		if file.Uid == "" {
			break
		}

		insertSong(db, file)
		c += 1
		fmt.Println(c)
	}

}

func scanner(status chan (string)) {
	// status <- scanning
	fileChan = make(chan File)

	go handleInserts()

	err := filepath.WalkDir(musicDir, walkHandler)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	fileChan <- File{
		Uid: "",
	}

	fmt.Println(len(files))

	// status <- ready
	status <- ready
}
