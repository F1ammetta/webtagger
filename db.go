package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ostafen/clover"
)

func insertSong(db *clover.DB, song File) error {
	doc := clover.NewDocumentOf(song)

	if songInDb(db, song.Uid) {
		return errors.New("Song already in DB")
	}

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

	var metadata map[string]any
	metadata = doc.Get("Metadata").(map[string]any)

	song.Uid = doc.Get("Uid").(string)
	song.Name = doc.Get("Name").(string)
	song.Size = float32(doc.Get("Size").(float64))
	song.Deleted = doc.Get("Deleted").(bool)

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
		doc.Set("Deleted", false)
		db.Save("songs", doc)
		return true
	}

	return false
}

type EventType int

const (
	Insert EventType = iota
	Query
	Edit
	Delete
	All
	Scan
	Clean
)

func (e EventType) String() string {
	return [...]string{"Insert", "Query", "Edit", "Delete", "All", "Scan", "Clean"}[e]
}

type DbEvent struct {
	eventType  EventType
	data       any
	resultChan chan (DbResult)
	// File for insertions, Edits, and Deletes, string for querys
}

type DbResult struct {
	data any
	err  error
}

var eventChan chan (DbEvent)

type DbStatus int

const (
	IO DbStatus = iota
	SCANNING
	IDLE
)

func (e DbStatus) String() string {
	return [...]string{"IO", "SCANNING", "IDLE"}[e]
}

var dbStatus DbStatus

func setStatus(eventType EventType) {
	switch eventType {
	case Query, All, Insert, Delete, Edit:
		dbStatus = IO
	case Scan:
		dbStatus = SCANNING
	}
}

func dispatch(event DbEvent) {
	eventChan <- event
}

func handleEvent(db *clover.DB, event DbEvent) {
	setStatus(event.eventType)
	result := event.resultChan
	defer close(result)
	switch event.eventType {
	case Insert:
		file := event.data.(File)

		if file.Uid == "" {
			result <- DbResult{
				data: nil,
				err:  errors.New("Invalid file"),
			}
			return
		}

		if err := insertSong(db, file); err != nil {
			result <- DbResult{
				data: nil,
				err:  err,
			}
		} else {
			result <- DbResult{
				data: nil,
				err:  nil,
			}
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
		file := event.data.(File)
		uid := file.Uid

		if err := editSong(db, uid, file); err != nil {
			result <- DbResult{
				data: nil,
				err:  err,
			}
		} else {
			result <- DbResult{
				data: nil,
				err:  nil,
			}
		}
		return
	case Delete:

	case Scan:
		flagFiles(db)
		cleanChan := make(chan struct{})
		go scanner(cleanChan)
		go clean(cleanChan)
	case Clean:
		cleaner(db)
	}
}

func cleaner(db *clover.DB) {
	query := db.Query("songs")

	err := query.Where(clover.Field("Deleted").Eq(true)).Delete()

	if err != nil {
		errLog(err)
	}

	infoLog("Cleaned db successfully")
}

func clean(cleanChan chan struct{}) {
	<-cleanChan

	resChan := make(chan DbResult)

	ev := DbEvent{
		data:       nil,
		eventType:  Clean,
		resultChan: resChan,
	}

	dispatch(ev)
	<-resChan
}

func flagFiles(db *clover.DB) {
	query := db.Query("songs")

	if res, err := query.FindAll(); err != nil {
		fmt.Println(err)
	} else {
		for _, doc := range res {
			doc.Set("Deleted", true)
			db.Save("songs", doc)
		}
	}
}

func eventLoop(ctx context.Context) {
	dbStatus = IDLE
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
			infoLog("DB Event", event.eventType.String())
			handleEvent(db, event)
			dbStatus = IDLE
		}
	}

}

func editSong(db *clover.DB, uid string, file File) error {
	query := db.Query("songs")

	doc, err := query.Where(clover.Field("Uid").Eq(uid)).FindFirst()

	if err != nil {
		fmt.Println(err)
	}

	if doc == nil {
		return errors.New("Couldn't find song")
	}

	meta := file.Metadata

	doc.Set("Metadata.Title", meta.Title)
	doc.Set("Metadata.Artist", meta.Artist)
	doc.Set("Metadata.Album", meta.Album)
	doc.Set("Metadata.Year", meta.Year)
	doc.Set("Metadata.Genre", meta.Genre)

	return db.Save("songs", doc)
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

	var song File
	song = songUnmarshal(doc)

	if err != nil {
		return File{}, errors.New("Couldn't find song")
	}

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
