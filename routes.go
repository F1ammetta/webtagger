package main

import (
	"encoding/json"
	"net/http"
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
		errLog(err)
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
