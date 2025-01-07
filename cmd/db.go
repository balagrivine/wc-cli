package main

import (
	"fmt"
)

// This is the simplest key-value datastore implementation
// Store data, give data back

type Database struct {
	key uint
	val map[string]string
}

func (db *Database) store(key uint, val map[string]string) {
	db.key = key
	db.val = val
}

func (db *Database) retrieve(key uint) map[string]string {
	var result map[string]string
	if db.key == key {
		result = db.val
	}
	return result
}

// func (db *Database) update(

func main() {
	db := &Database{}

	personalInfo := map[string]string {"name": "Bala"}
	db.store(1, personalInfo)

	name := db.retrieve(1)
	fmt.Println(name)
}
