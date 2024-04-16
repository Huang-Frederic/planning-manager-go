package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func dbConnect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
	  return nil, err
	}

	dbHost := os.Getenv("DBHost")
	dbPort := os.Getenv("DBPort")
	dbUser := os.Getenv("DBUser")
	dbPassword := os.Getenv("DBPassword")
	dbName := os.Getenv("DBName")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = dbCreate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}


func dbCreate(db *sql.DB) error {

	// Room
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Room (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			capacity INT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	var rowCount int
	err = db.QueryRow("SELECT COUNT(*) FROM Room").Scan(&rowCount)
	if err != nil {
		return err
	}
	if rowCount == 0 {
		_, err = db.Exec(`
			INSERT INTO Room (name, capacity) VALUES
			('Salle 1', 30),
			('Salle 2', 40),
			('Salle 3', 50)
		`)
		if err != nil {
			return err
		}
	}

	// Reservation
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Reservation (
			id INT AUTO_INCREMENT PRIMARY KEY,
			roomID INT NOT NULL,
			startTime DATETIME NOT NULL,
			endTime DATETIME NOT NULL,
			FOREIGN KEY (roomID) REFERENCES Room(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
