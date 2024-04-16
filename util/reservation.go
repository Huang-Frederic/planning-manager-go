package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateReservation(roomID int, startTime time.Time, endTime time.Time) error {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO Reservation (roomID, startTime, endTime) VALUES (?, ?, ?)
	`, roomID, startTime, endTime)
	if err != nil {
		return err
	}

	return nil
}

func CreateReservationFromUserInput() error {
	fmt.Println("Création d'une nouvelle réservation")

    reader := bufio.NewReader(os.Stdin)

	fmt.Println("Entrez l'identifiant de la salle : ")
	roomIDInput, _ := reader.ReadString('\n')
	roomIDInput = strings.TrimSpace(roomIDInput)
    roomID, _ := strconv.Atoi(roomIDInput)

	fmt.Println("Entrez l'heure de début de la réservation (format YYYY-MM-DD HH:MM:SS) : ")
	startTimeInput, _ := reader.ReadString('\n')
	startTimeInput = strings.TrimSpace(startTimeInput)
	startTime, _ := time.Parse("2006-01-02 15:04:05", startTimeInput)

	fmt.Println("Entrez l'heure de fin de la réservation (format YYYY-MM-DD HH:MM:SS) : ")
	endTimeInput, _ := reader.ReadString('\n')
	endTimeInput = strings.TrimSpace(endTimeInput)
	endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeInput)

	err := CreateReservation(roomID, startTime, endTime)
	if err != nil {
		return err
	}

	fmt.Println("La réservation a été créée avec succès.")
	return nil
}