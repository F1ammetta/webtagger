package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const musicDir = "/srv/music/"

const (
	scanning = "scanning"
	ready    = "ready"
)

type Meta struct {
	Title  string `json:"title" clover:"Title"`
	Artist string `json:"artist" clover:"Artist"`
	Album  string `json:"album" clover:"Album"`
	Genre  string `json:"genre" clover:"Genre"`
	Year   int    `json:"year" clover:"Year"`
}

type File struct {
	Uid      string  `json:"uid" clover:"Uid"`
	Name     string  `json:"name" clover:"Name"`
	Size     float32 `json:"size" clover:"Size"`
	Metadata Meta    `json:"metadata" clover:"Metadata"`
	Deleted  bool    `json:"-" clover:"Deleted"`
}

var filenames []string

func main() {
	ctx, _ := context.WithCancel(context.Background())

	initDb()

	eventChan = make(chan DbEvent)

	go eventLoop(ctx)

	<-time.After(time.Second)

	// result := make(chan (DbResult), 1)
	//
	// scanEvent := DbEvent{
	// 	eventType:  Scan,
	// 	data:       nil,
	// 	resultChan: result,
	// }
	//
	// dispatch(scanEvent)
	//
	// <-result

	fswatch()

	http.HandleFunc("/songs", songsHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/cover/get/{uid}", coverHandler)
	http.HandleFunc("/cover/set/{uid}", setCoverHandler)
	http.HandleFunc("/edit/{uid}", editHandler)
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":6008", nil)

	// song := songInDb(db, "c1e9f5258e12594a710166176e8bdc2d6f415455937ec3aa9362a21f284a2b08")
}

func infoLog(s ...string) {

	log.Output(1, fmt.Sprintf("INFO: %s", s))
}

func errLog(e error, s ...string) {
	log.Output(1, fmt.Sprintf("ERROR: %s %s", s, e.Error()))
}
