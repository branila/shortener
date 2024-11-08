package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const dbName = "shortener.json"

var (
	db    = make(map[string]string)
	mutex = &sync.RWMutex{}
)

func SyncDb() error {
	jsonData, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return fmt.Errorf("Error marshalling data: %e", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting working directory: %e", err)
	}

	dbPath := filepath.Join(wd, dbName)

	err = os.WriteFile(dbPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing data to file: %e", err)
	}

	return nil
}

func GetId(id string) (string, bool) {
	mutex.RLock()
	id, exists := db[id]
	mutex.RUnlock()

	return id, exists
}

func SetId(id, url string) {
	mutex.Lock()
	db[id] = url
	SyncDb()
	mutex.Unlock()
}
