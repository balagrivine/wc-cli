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

func (db *Database) set(key uint, val map[string]string) {
	db.key = key
	db.val = val
}

func (db *Database) get(key uint) map[string]string {
	var result map[string]string
		if db.key == key {
			result = db.val
		}
	return result
}

func (db *Database) update(key uint, val map[string]string) {
	if db.key == key {
		db.val = val
	}
}

func (db *Database) remove(key uint) {
	if db.key == key {
		db.val = nil
	}
}

// func (db *Database) update(

func main() {
	db := Database{}

	personalInfo := map[string]string{"name": "Bala"}
	db.set(1, personalInfo)

	name := db.get(1)
	fmt.Println(name)

	db.update(1, map[string]string{"name": "Ochieng"})
	
	name = db.get(1)
	fmt.Println(name)

	db.remove(1)
	name = db.get(1)
	fmt.Println(name)
}
