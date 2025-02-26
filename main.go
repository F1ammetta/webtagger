package main

import (
	// "errors"
	"context"
	"fmt"
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

	// meta, err := gatherMetadata("1-01 PSYCHO.flac")
	//
	// fmt.Println(meta)

	initDb()

	// err := editMetadata(
	// 	"1-01 PSYCHO.flac",
	// 	"soyeon.jpg",
	// 	Meta{
	// 		Title:  "sexo",
	// 		Artist: "tilin",
	// 		Album:  "prrr",
	// 		Year:   2032,
	// 		Genre:  "Huh",
	// 	},
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// }

	eventChan = make(chan (DbEvent))

	go eventLoop(ctx)

	<-time.After(time.Second)

	status := make(chan string)
	go scanner(status)

	// go stats(status)

	<-status

	songsChan := make(chan (DbResult), 1)
	getEvent := DbEvent{
		eventType:  All,
		data:       nil,
		resultChan: songsChan,
	}

	dispatch(getEvent)

	result := <-songsChan

	songs := result.data.([]File)

	fmt.Println(songs)

	// song := songInDb(db, "c1e9f5258e12594a710166176e8bdc2d6f415455937ec3aa9362a21f284a2b08")
	// fmt.Println(song)
}
