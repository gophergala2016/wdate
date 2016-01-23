package main

import (
	"encoding/gob"
	"os"
	"path/filepath"
)

var locs = make(map[string]int)
var dbFile = filepath.Join(os.Getenv("HOME"), ".wdate")

func loadDB() error {
	if !isFileExist(dbFile) {
		return nil
	}

	f, err := os.Open(dbFile)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(f)
	err = dec.Decode(&locs)
	if err != nil {
		return err
	}
	return nil
}

func addLocation(name string, offset int) error {
	locs[name] = offset
	f, err := os.Create(dbFile)
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(locs)
	if err != nil {
		return err
	}

	return nil
}

func removeLocation(name string) error {
	delete(locs, "name")
	f, err := os.Create(dbFile)
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(locs)
	if err != nil {
		return err
	}

	return nil
}

func isFileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}