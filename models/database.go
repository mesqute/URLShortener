package models

import (
	"database/sql"
	"log"
)

type database interface {
	Get(link string) (URL string, err error)
	Insert(URL string) (link string, err error)
}

type Database database

var db Database

// SetDB устанавливает активную БД
func SetDB(name string) {
	switch name {
	case "in-memory":
		db = &inMemory{items: make(map[string]value)}
	case "postgresql":
		connStr := "user=postgres host=db dbname=postgres password=0000 sslmode=disable"
		base, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		err = base.Ping()
		if err != nil {
			log.Fatal(err)
		}
		db = &postgreSQL{db: base}
	default:
		db = nil
	}
}

// GetDB возвращает активную БД
func GetDB() Database {
	return db
}
