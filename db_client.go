package main

import (
	"database/sql"
	"os"

	// _ "modernc.org/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	db *sql.DB
}

func NewDBClient(dbFile string) DBClient {
	client := DBClient{nil}
	client.open(dbFile)
	return client
}

func (id *DBClient) open(dbFile string) {
	os.Remove(dbFile)
	file, err := os.Create(dbFile)
	if err != nil {
		panic(err.Error())
	}
	file.Close()

	// Alt open options:
	// dbFile = "file:" + dbFile + "?mode=memory&cache=shared"
	// db, err := sql.Open("sqlite", dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err.Error())
	}
	id.db = db
}

func (id *DBClient) exec(sql string) (sql.Result, error) {
	id.db.Prepare(sql)
	result, err := id.db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
	return result, err
}

func (id *DBClient) query(sql string) []string {
	result, err := id.db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	var rows []string
	for result.Next() {
		var col string
		err := result.Scan(&col)
		if err != nil {
			panic(err.Error())
		}
		rows = append(rows, col)
	}
	return rows
}
