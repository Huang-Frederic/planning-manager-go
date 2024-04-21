package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// ---------------------------------------------------------
// DB FUNCTIONS
// ---------------------------------------------------------

// Function to connect to the db
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

	err = init_db(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Function to create the tables if they don't exist
func init_db(db *sql.DB) error {
	// Check if Room table exists
	var roomCount int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ?", os.Getenv("DBName"), "Room").Scan(&roomCount)
	if err != nil {
		return err
	}

	if roomCount == 0 {
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

	var reservationCount int
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ?", os.Getenv("DBName"), "Reservation").Scan(&reservationCount)
	if err != nil {
		return err
	}

	if reservationCount == 0 {
		_, err := db.Exec(`
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

		_, err = db.Exec(`
		INSERT INTO Reservation (roomID, startTime, endTime) VALUES
		(1, "2024-06-01 10:00:00", "2024-06-01 11:00:00"),
		(1, "2024-06-10 14:00:00", "2024-06-10 15:00:00"),
		(1, "2024-06-20 16:00:00", "2024-06-20 17:00:00"),
		(2, "2024-07-05 09:00:00", "2024-07-05 10:00:00"),
		(2, "2024-07-15 11:00:00", "2024-07-15 12:00:00"),
		(2, "2024-07-25 13:00:00", "2024-07-25 14:00:00"),
		(3, "2024-08-02 08:00:00", "2024-08-02 09:00:00"),
		(3, "2024-08-12 12:00:00", "2024-08-12 13:00:00"),
		(3, "2024-08-22 15:00:00", "2024-08-22 16:00:00")
		`)
		if err != nil {
			return err
		}
	}
	return nil
}

// ---------------------------------------------------------
// DB BOTH USED FUNCTIONS
// ---------------------------------------------------------


// Function to check if a room is available at a given time
func isRoomAvailable(roomID int, startTime, endTime time.Time) bool {
	db, err := dbConnect()
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Erreur de connexion à la base de données:", err)
		fmt.Println("----------------------------------------------------")
		return false
	}
	defer db.Close()

	query := `
		SELECT COUNT(*)
		FROM Reservation
		WHERE roomID = ? AND ((startTime <= ? AND endTime >= ?) OR (startTime <= ? AND endTime >= ?))
	`
	var count int
	err = db.QueryRow(query, roomID, startTime, endTime, startTime, endTime).Scan(&count)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Erreur lors de la requête SQL:", err)
		fmt.Println("----------------------------------------------------")
		return false
	}
	return count == 0
}

// ---------------------------------------------------------
// DB ROOM FUNCTIONS
// ---------------------------------------------------------

// Function to add a room to the db
func addRoom(name string, capacity int) error {
    db, err := dbConnect()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO Room (name, capacity) VALUES (?, ?)", name, capacity)
    if err != nil {
        return err
    }

    return nil
}

// Function to get all rooms from the db
func getRooms() ([]Room, error) {
    db, err := dbConnect()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, name, capacity FROM Room")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var rooms []Room

    for rows.Next() {
        var room Room
        err := rows.Scan(&room.ID, &room.Name, &room.Capacity)
        if err != nil {
            return nil, err
        }
        rooms = append(rooms, room)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return rooms, nil
}

// ---------------------------------------------------------
// DB RESERVATION FUNCTIONS
// ---------------------------------------------------------

// Function to create a reservation
func createReservation(reservation Reservation) error {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO Reservation (roomID, startTime, endTime) VALUES (?, ?, ?)
	`, reservation.RoomID, reservation.StartTime, reservation.EndTime)
	if err != nil {
		return err
	}

	return nil
}

// Function to cancel a reservation
func cancelReservation(reservationID int) error {
    db, err := dbConnect()
    if err != nil {
        return err
    }
    defer db.Close()

    // Vérifiez d'abord si la réservation existe
    var count int
    err = db.QueryRow(`SELECT COUNT(*) FROM Reservation WHERE id = ?`, reservationID).Scan(&count)
    if err != nil {
        return err
    }

    if count == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("L'ID selectionnée n'existe pas'")
		fmt.Println("----------------------------------------------------")
		fmt.Println("")
        return err
    }

    // Supprimez la réservation
    _, err = db.Exec(`DELETE FROM Reservation WHERE id = ?`, reservationID)
    if err != nil {
        return err
    }

    fmt.Println("----------------------------------------------------")
    fmt.Println("La réservation a été annulée avec succès.")
    fmt.Println("----------------------------------------------------")
    fmt.Println("")

    return nil
}

// Function to get reservations
func getReservations(roomID int) ([]Reservation, error) {
	var reservations []Reservation
	var rows *sql.Rows

	db, err := dbConnect()
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Connexion à la base de données échouée.")
		fmt.Println("----------------------------------------------------")
		fmt.Println("")
		return nil, err
	}
	defer db.Close()

	// Exécuter la requête SQL en fonction de la valeur de roomID
	if roomID == 0 {
		rows, err = db.Query(`SELECT * FROM Reservation`)
		if err != nil {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Les informations n'ont pas pu être récupérées.")
			fmt.Println("----------------------------------------------------")
			fmt.Println("")
			return nil, err
		}
	} else {
		rows, err = db.Query(`SELECT * FROM Reservation WHERE roomID = ?`, roomID)
		if err != nil {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Les informations n'ont pas pu être récupérées pour cette salle.")
			fmt.Println("----------------------------------------------------")
			fmt.Println("")
			return nil, err
		}
	}
	defer rows.Close()

	// Stocker les réservations dans un tableau
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.ID, &reservation.RoomID, &reservation.StartTime, &reservation.EndTime)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
