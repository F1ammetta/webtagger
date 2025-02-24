package main

import (
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

func findSong(db *clover.DB, uid string) File {
	query := db.Query("songs")

	doc, err := query.Where(clover.Field("Uid").Eq(uid)).FindFirst()

	if err != nil {
		fmt.Println(err)
	}

	song := songUnmarshal(doc)

	return song
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
