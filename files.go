package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

var exts = []string{
	"m4a", "mp3", "flac", "ogg", "opus",
}

var id = ""

var fileChan chan (File)

func fswatch() {
	w, err := fsnotify.NewWatcher()

	if err != nil {
		errLog(err, "Couldn't start fs watcher")
	}

	go watchLoop(w)

	if err := w.Add(musicDir); err != nil {
		errLog(err, "Couldn't add music path to fs watcher")
	}

}

func watchLoop(w *fsnotify.Watcher) {
	defer w.Close()
	for {
		// TODO: Handle unwanted events created from editing files
		// NOTE: Maybe use a channel from the event loop to notify file editings
		select {
		case err, ok := <-w.Errors:
			if !ok {
				break
			}
			errLog(err, "Watcher")
		case e, ok := <-w.Events:
			if !ok {
				break
			}
			switch e.Op {
			case fsnotify.Create:
				// TODO: Handle File Creation
				<-time.After(time.Millisecond * 100)
				path := e.Name

				ext := strings.Split(e.Name, ".")
				if !slices.Contains(exts, ext[len(ext)-1]) {
					continue
				}

				fd, err := os.Open(path)

				if err != nil {
					errLog(err)
				}

				// var file File

				info, err := fd.Stat()

				if err != nil {
					errLog(err)
				}

				var file File
				file.Name = info.Name()
				file.Size = float32(info.Size()) / 1000_000

				hash := sha256.New()
				hash.Write([]byte(file.Name))
				hashed := hash.Sum(nil)

				uid := hex.EncodeToString(hashed)

				file.Uid = uid
				file.Metadata, err = gatherMetadata(path)
				file.Deleted = false

				noChan := make(chan DbResult, 1)

				event := DbEvent{
					eventType:  Insert,
					data:       file,
					resultChan: noChan,
				}

				go dispatch(event)
			case fsnotify.Remove:
				// TODO: Handle File Removal
				dirs := strings.Split(e.Name, "/")
				name := dirs[len(dirs)-1]
				hash := sha256.New()
				hash.Write([]byte(name))
				hashed := hash.Sum(nil)

				uid := hex.EncodeToString(hashed)

				ext := strings.Split(e.Name, ".")
				if !slices.Contains(exts, ext[len(ext)-1]) {
					continue
				}

				noChan := make(chan DbResult, 1)

				event := DbEvent{
					eventType:  Delete,
					data:       uid,
					resultChan: noChan,
				}

				go dispatch(event)
			case fsnotify.Rename:
				// TODO: Handle File Renaming, or don't

			}

			infoLog(fmt.Sprintf("%s", e))
		}
	}
}

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

	var file File
	file.Name = info.Name()
	file.Size = float32(info.Size()) / 1000_000

	hash := sha256.New()
	hash.Write([]byte(file.Name))
	hashed := hash.Sum(nil)

	uid := hex.EncodeToString(hashed)

	file.Uid = uid
	file.Metadata, err = gatherMetadata(path)
	file.Deleted = false

	noChan := make(chan DbResult, 1)

	event := DbEvent{
		eventType:  Insert,
		data:       file,
		resultChan: noChan,
	}

	go dispatch(event)

	return nil
}

func scanner(cleanChan chan (struct{})) {
	err := filepath.WalkDir(musicDir, walkHandler)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	cleanChan <- struct{}{}
}
