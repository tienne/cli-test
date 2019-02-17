package database

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/asdine/storm"
	"github.com/mozzet/cli/config"
	bolt "go.etcd.io/bbolt"
)

var (
	dbFile = "mozzet.db"
)

type DB struct {
	home string
	storm.DB
}

func Execute(fn func(*DB) error) error {
	db, err := current()
	if err != nil {
		return err
	}
	defer db.Close()

	return fn(db)
}

func current() (*DB, error) {
	home := config.ConfigPath()

	path := filepath.Join(home, dbFile)
	db, err := open(path)
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, fmt.Errorf("db is nil while no error in opening at '%s'", path)
	}

	return db, nil
}

func open(path string) (*DB, error) {
	stormDB, err := storm.Open(path, storm.BoltOptions(0600, &bolt.Options{Timeout: 2 * time.Second}))

	if err != nil {
		return nil, fmt.Errorf("opening db at %s: %s (any awless existing process running?)", path, err)
	}

	return &DB{*stormDB}, nil
}
