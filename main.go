package main

import (
	// "errors"
	"fmt"
	"time"
	// "strconv"
	"github.com/ostafen/clover"
)

// const musicDir = "/home/fiammetta/Music/"
const musicDir = "./"

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

var files []File
var filenames []string

func main() {
	status := make(chan string)

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

	go scanner(status)

	// go stats(status)

	<-status
	<-time.After(3 * time.Second)

	db, err := clover.Open("./data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	song := findSong(db, id)

	fmt.Println(song)
}
