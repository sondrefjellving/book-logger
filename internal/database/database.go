package database

import (
	"errors"
	"os"
	"sync"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)

type DB struct {
	path string
	mux *sync.RWMutex
}

type DBStructure struct {
	books []data_types.Book 
}

func NewDB(path string) (*DB, error) {
	db := DB{
		path: path,
		mux: &sync.RWMutex{},
	}

	return &db, db.setupDB()
}

func (db *DB) setupDB() error {
	return db.ensureDB()
}

func (db *DB) ensureDB() error {
	if _, err := os.ReadFile(db.path); errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}
	return nil
}

func (db *DB) createDB() error {
	dbStruct := DBStructure{
		books: make([]data_types.Book, 0),
	}	
	return db.writeDB(dbStruct)
}

func (db *DB) writeDB(payload interface{}) error {

}