package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	// "fmt"
	"io"
	"net/http"
	"os"
)

func scanHandler(w http.ResponseWriter, r *http.Request) {
	infoLog("Incoming GET @ /scan")
	result := make(chan (DbResult), 1)

	scanEvent := DbEvent{
		eventType:  Scan,
		data:       nil,
		resultChan: result,
	}

	dispatch(scanEvent)

	w.Write([]byte("{\"status\": \"scanning\"}"))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	infoLog("Incoming POST @ /edit")

	var body []byte
	var err error

	if body, err = io.ReadAll(r.Body); err != nil {
		errLog(err)
	}

	var file File

	if err := json.Unmarshal(body, &file); err != nil {
		fmt.Println(string(body))
		errLog(err)
		return
	}

	path := musicDir + file.Name

	editMetadata(path, file.Metadata)

	resChan := make(chan DbResult)

	editEv := DbEvent{
		data:       file,
		eventType:  Edit,
		resultChan: resChan,
	}

	dispatch(editEv)

	res := <-resChan

	if res.err != nil {
		w.Write([]byte("{\"success\":false}"))
	}

	w.Write([]byte("{\"success\":true}"))
}

type CoverUpload struct {
	Data     string `json:"data"`
	MimeType string `json:"type"`
}

func setCoverHandler(w http.ResponseWriter, r *http.Request) {
	infoLog("Incoming POST @ /cover")

	uid := r.PathValue("uid")

	resChan := make(chan DbResult)

	dbE := DbEvent{
		eventType:  Query,
		data:       uid,
		resultChan: resChan,
	}

	dispatch(dbE)

	res := <-resChan

	file := res.data.(File)

	filePath := musicDir + file.Name

	var body []byte
	var err error

	if body, err = io.ReadAll(r.Body); err != nil {
		errLog(err)
	}

	var cover CoverUpload

	if err := json.Unmarshal(body, &cover); err != nil {
		errLog(err)
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(cover.Data)

	if err != nil {
		errLog(err)
	}

	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))
	hashed := hash.Sum(nil)

	dir := hex.EncodeToString(hashed)
	// fmt.Println(cover)
	coverFile := os.TempDir() + "/temp-" + dir + "." + strings.Split(cover.MimeType, "/")[1]

	infoLog(coverFile)

	if err := os.WriteFile(coverFile, imageData, os.FileMode(int(0777))); err != nil {
		errLog(err, "Couldn't create temp file")
		w.WriteHeader(500)
	}

	if err := setCover(filePath, coverFile); err != nil {
		errLog(err, "Couldn't set cover")
		w.WriteHeader(500)
	}

	os.Remove(coverFile)
	w.Write([]byte("{\"success\":true}"))
}

func coverHandler(w http.ResponseWriter, r *http.Request) {
	infoLog("Incoming GET @ /cover")

	uid := r.PathValue("uid")

	resChan := make(chan DbResult)

	dbE := DbEvent{
		eventType:  Query,
		data:       uid,
		resultChan: resChan,
	}

	dispatch(dbE)

	res := <-resChan

	file := res.data.(File)

	cover, err := getCover(file.Name)

	if err != nil {
		file, err := os.ReadFile("./album.svg")
		if err != nil {
			errLog(err, "Couldn't fetch cover")
		}

		w.Header().Add("Content-Type", "image/svg+xml")
		w.Write(file)
	}

	w.Write(cover)
}

func songsHandler(w http.ResponseWriter, r *http.Request) {
	infoLog("Incoming GET @ /songs")

	songsChan := make(chan (DbResult), 1)

	getEvent := DbEvent{
		eventType:  All,
		data:       nil,
		resultChan: songsChan,
	}

	dispatch(getEvent)

	result := <-songsChan

	songs := result.data.([]File)

	if songsJson, err := json.Marshal(songs); err != nil {
		errLog(err)
	} else {
		if _, err := w.Write(songsJson); err != nil {
			errLog(err)
		}
	}
}
