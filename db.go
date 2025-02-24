package main

import (
	"fmt"
	"github.com/ostafen/clover"
)

func insertSong(db *clover.DB, song File) error {
	doc := clover.NewDocumentOf(song)

	err := db.Insert("songs", doc)
	if err != nil {
		return err
	}

	return nil
}

func findSong(db *clover.DB, uid string) File {
	query := db.Query("songs")

	doc, err := query.Where(clover.Field("Uid").Eq(uid)).FindFirst()

	if err != nil {
		fmt.Println(err)
	}

	var song File

	err = doc.Unmarshal(song)

	var metadata map[string]interface{}
	metadata = doc.Get("Metadata").(map[string]interface{})

	song.Uid = uid
	song.Name = doc.Get("Name").(string)
	song.Size = float32(doc.Get("Size").(float64))

	song.Metadata.Album = metadata["Album"].(string)
	song.Metadata.Title = metadata["Title"].(string)
	song.Metadata.Artist = metadata["Artist"].(string)
	song.Metadata.Genre = metadata["Genre"].(string)
	song.Metadata.Year = int(metadata["Year"].(int64))

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
