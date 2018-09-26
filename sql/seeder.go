package sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

// In case of MySQL database, the connection must enable option multiStatements=true

// SeedByFile reads SQL seed file and executes it against the given DB
func SeedByFile(db *sql.DB, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read the seed file %s, %s", filePath, err.Error())
	}

	return Seed(db, string(data))
}

// Seed uses given SQL commands for seeding against the database
func Seed(db *sql.DB, sqlCmds string) error {
	_, err := db.Exec(sqlCmds)
	if err != nil {
		return fmt.Errorf("could not seed data, %s", err.Error())
	}
	return nil
}

// CleanByFile executes cleaning up SQL commands in the given SQL file against the db
func CleanByFile(db *sql.DB, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could ")
	}
	return Clean(db, string(data))
}

// Clean uses given SQL commands for cleaning up the database
func Clean(db *sql.DB, sqlCmds string) error {
	_, err := db.Exec(sqlCmds)
	if err != nil {
		return fmt.Errorf("could not clean up the database, %s", err.Error())
	}
	return nil
}
