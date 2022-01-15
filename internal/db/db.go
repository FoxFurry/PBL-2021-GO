package db

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

func GetDB() *sql.DB {
	dbOnce.Do(func() {
		var err error

		dbFile := viper.GetString("db.path") + viper.GetString("db.name")

		checkDBFile(dbFile)
		dbInstance, err = sql.Open("sqlite3", dbFile)

		if err != nil {
			log.Fatalf("Could not open database: %v", err)
		}

		createUserTable(dbInstance)
		createRoomTable(dbInstance)
		createRoomParticipantTable(dbInstance)
		createMessageTable(dbInstance)
	})

	return dbInstance
}

func checkDBFile(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.Create(path)
	}
}

func createUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "user" (
		"id"	INTEGER,
		"full_name"	TEXT,
		"email"	TEXT,
		"password"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create user table")
	}
}

func createRoomParticipantTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "room_participant" (
		"id"	INTEGER,
		"user_id"	INTEGER,
		"room_id"	INTEGER,
		"role"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create user table")
	}
}

func createRoomTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "room" (
		"id"	INTEGER,
		"name"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create user table")
	}
}

func createMessageTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "message" (
		"id"	INTEGER,
		"sender_id"	INTEGER,
		"room_id" INTEGER,
		"data" TEXT,
		"time" INT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create message table")
	}
}
