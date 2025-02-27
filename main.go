package main

import (
	// "errors"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	// "strconv"
	// "github.com/ostafen/clover"
)

const musicDir = "/home/fiammetta/Music/"

// const musicDir = "./"

const (
	scanning = "scanning"
	ready    = "ready"
)

type Meta struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

type File struct {
	Uid      string  `json:"uid"`
	Name     string  `json:"name"`
	Size     float32 `json:"size"`
	Metadata Meta    `json:"metadata"`
}

var filenames []string

func main() {
	ctx, _ := context.WithCancel(context.Background())

	initDb()

	eventChan = make(chan DbEvent)

	go eventLoop(ctx)

	<-time.After(time.Second)

	http.HandleFunc("/songs", songsHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/cover/{uid}", coverHandler)

	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	// song := songInDb(db, "c1e9f5258e12594a710166176e8bdc2d6f415455937ec3aa9362a21f284a2b08")
}

func infoLog(s ...string) {

	log.Output(1, fmt.Sprintf("INFO: %s", s))
}

func errLog(e error) {
	log.Output(1, fmt.Sprintf("ERROR: %s", e.Error()))
}
