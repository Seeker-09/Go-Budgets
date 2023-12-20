package main

import "database/sql"

func openDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "budgets.db")
	if err != nil {
		return db, err
	}
	return db, nil
}