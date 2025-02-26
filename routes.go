package main

import (
	"encoding/json"
	"net/http"
)

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
