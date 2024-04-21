package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Function to list rooms
func listRooms() {
    reader := bufio.NewReader(os.Stdin)

    UNBUFFER, _ := reader.ReadString('\n')
	UNBUFFER = strings.TrimSpace(UNBUFFER)

    fmt.Print("Entrez l'heure de début du créneau (format YYYY-MM-DD HH:MM:SS) : ")
    startTimeInput, _ := reader.ReadString('\n')
	startTimeInput = strings.TrimSpace(startTimeInput)
	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeInput)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Format de date invalide pour l'heure de début.")
		fmt.Println("----------------------------------------------------")
		return
	}

    fmt.Print("Entrez l'heure de fin du créneau (format YYYY-MM-DD HH:MM:SS) : ")
    endTimeInput, _ := reader.ReadString('\n')
	endTimeInput = strings.TrimSpace(endTimeInput)
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeInput)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Format de date invalide pour l'heure de fin.")
		fmt.Println("----------------------------------------------------")
		return 
	}

	if startTime.After(endTime) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("La date de début ne peut pas être après la date de fin.")
		fmt.Println("----------------------------------------------------")
		return 
	}

    // Récupérez les salles de la base de données
    rooms, err := getRooms()
    if err != nil {
        fmt.Println("----------------------------------------------------")
        fmt.Println("Erreur lors de la récupération des salles : ", err)
        fmt.Println("----------------------------------------------------")
        return
    }

    // Affichez les salles disponibles pour le créneau spécifié
    fmt.Println("")
    fmt.Println("Salles disponibles pour le créneau spécifié :")
    fmt.Println("----------------------------------------------------")
    for _, room := range rooms {
        if isRoomAvailable(room.ID, startTime, endTime) {
            fmt.Printf("%d. %s (Capacité : %d)\n", room.ID, room.Name, room.Capacity)
        }
    }
    fmt.Println("")
}
