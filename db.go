package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ostafen/clover"
)

func insertSong(db *clover.DB, song File) error {
	doc := clover.NewDocumentOf(song)

	fmt.Println("achoo")
	if songInDb(db, song.Uid) {
		return errors.New("Song already in DB")
	}
	fmt.Println("salud")

	err := db.Insert("songs", doc)
	if err != nil {
		return err
	}

	return nil
}

func getSongs(db *clover.DB) []File {
	var files []File

	query := db.Query("songs")

	if res, err := query.FindAll(); err != nil {
		fmt.Println(err)
	} else {

		for _, doc := range res {
			files = append(files, songUnmarshal(doc))
		}
	}

	return files
}

func songUnmarshal(doc *clover.Document) File {
	var song File

	var metadata map[string]interface{}
	metadata = doc.Get("Metadata").(map[string]interface{})

	song.Uid = doc.Get("Uid").(string)
	song.Name = doc.Get("Name").(string)
	song.Size = float32(doc.Get("Size").(float64))

	song.Metadata.Album = metadata["Album"].(string)
	song.Metadata.Title = metadata["Title"].(string)
	song.Metadata.Artist = metadata["Artist"].(string)
	song.Metadata.Genre = metadata["Genre"].(string)
	song.Metadata.Year = int(metadata["Year"].(int64))

	return song
}
func songInDb(db *clover.DB, uid string) bool {
	query := db.Query("songs")

	doc, err := query.Where(clover.Field("Uid").Eq(uid)).FindFirst()

	if err != nil {
		fmt.Println(err)
	}

	if doc == nil {
		return false
	}

	if doc.Get("Uid").(string) == uid {
		return true
	}

	return false
}

type EventType int

const (
	Insert EventType = iota + 1
	Query
	Edit
	Delete
	All
)

func (e EventType) String() string {
	return [...]string{"Unknown", "Insert", "Query", "Edit", "Delete", "All"}[e]
}

type DbEvent struct {
	eventType  EventType
	data       interface{}
	resultChan chan (DbResult)
	// File for insertions, Edits, and Deletes, string for querys
}

type DbResult struct {
	data interface{}
	err  error
}

var eventChan chan (DbEvent)

func dispatch(event DbEvent) {
	eventChan <- event
}

func handleEvent(db *clover.DB, event DbEvent) {
	result := event.resultChan
	defer close(result)
	switch event.eventType {
	case Insert:
		file := event.data.(File)

		fmt.Println("pene")

		if file.Uid == "" {
			result <- DbResult{
				data: nil,
				err:  errors.New("Invalid file"),
			}
			return
		}

		fmt.Println("pene2")
		if err := insertSong(db, file); err != nil {
			result <- DbResult{
				data: nil,
				err:  err,
			}
			fmt.Println("pene3")
		} else {
			result <- DbResult{
				data: nil,
				err:  nil,
			}
			fmt.Println("pene4")
		}

		return
	case Query:
		uid := event.data.(string)

		if file, err := findSong(db, uid); err != nil {
			result <- DbResult{
				data: file,
				err:  err,
			}
		} else {
			result <- DbResult{
				data: file,
				err:  nil,
			}
		}
		return
	case All:
		result <- DbResult{
			data: getSongs(db),
			err:  nil,
		}
		return
	case Edit:

	case Delete:

	}
}

func eventLoop(ctx context.Context) {
	db, err := clover.Open("./data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		select {
		case <-ctx.Done():
			return

		case event := <-eventChan:
			fmt.Println("Received event:", event.eventType.String())
			handleEvent(db, event)
		}
	}

}

func findSong(db *clover.DB, uid string) (File, error) {
	query := db.Query("songs")

	doc, err := query.Where(clover.Field("Uid").Eq(uid)).FindFirst()

	if err != nil {
		fmt.Println(err)
	}

	if doc == nil {
		return File{}, errors.New("Couldn't find song")
	}

	song := songUnmarshal(doc)

	if song.Uid == "" {
		return File{}, errors.New("Couldn't find song")
	}

	return song, nil
}

func initDb() {

	db, err := clover.Open("./data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exists, err := db.HasCollection("songs")

	if err != nil {
		fmt.Println(err)
	}

	// Ensure the collection exists
	if !exists {
		db.CreateCollection("songs")
	}

	fmt.Println("Database initialized!")
}
