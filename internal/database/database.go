package database

import (
	"encoding/json"
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
	Books []data_types.Book 
}

func NewDB(path string) (*DB, error) {
	db := DB{
		path: path,
		mux: &sync.RWMutex{},
	}

	err := db.ensureDB()
	return &db, err
}

func (db *DB) ensureDB() error {
	if _, err := os.ReadFile(db.path); errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}
	return nil
}

func (db *DB) createDB() error {
	dbStruct := DBStructure{
		Books: make([]data_types.Book, 0),
	}	
	return db.writeDB(dbStruct)
}

func (db *DB) writeDB(payload interface{}) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	dbStruct, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, dbStruct, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) loadDB() (DBStructure, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	data, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, err 
	}

	dbStruct := DBStructure{}
	err = json.Unmarshal(data, &dbStruct)
	if err != nil {
		return DBStructure{}, err 
	}

	return dbStruct, nil
} 