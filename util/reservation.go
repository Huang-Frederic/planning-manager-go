package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Function to create a reservation from user input
func createReservationFromUserInput() error {
	fmt.Println("")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Création d'une nouvelle réservation")
	fmt.Println("----------------------------------------------------")

	// Demandez à l'utilisateur de fournir les informations
	reader := bufio.NewReader(os.Stdin)

	UNBUFFER, _ := reader.ReadString('\n')
	UNBUFFER = strings.TrimSpace(UNBUFFER)

	fmt.Print("Entrez l'identifiant de la salle : ")
	roomIDInput, _ := reader.ReadString('\n')
	roomIDInput = strings.TrimSpace(roomIDInput)
	roomID, _ := strconv.Atoi(roomIDInput)

	fmt.Print("Entrez l'heure de début de la réservation (format YYYY-MM-DD HH:MM:SS) : ")
	startTimeInput, _ := reader.ReadString('\n')
	startTimeInput = strings.TrimSpace(startTimeInput)
	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeInput)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Format de date invalide pour l'heure de début.")
		fmt.Println("----------------------------------------------------")
		return err
	}

	fmt.Print("Entrez l'heure de fin de la réservation (format YYYY-MM-DD HH:MM:SS) : ")
	endTimeInput, _ := reader.ReadString('\n')
	endTimeInput = strings.TrimSpace(endTimeInput)
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeInput)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Format de date invalide pour l'heure de fin.")
		fmt.Println("----------------------------------------------------")
		return err
	}

	if startTime.After(endTime) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("La date de début ne peut pas être après la date de fin.")
		fmt.Println("----------------------------------------------------")
		return nil
	}


	if !isRoomAvailable(roomID, startTime, endTime) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("La salle n'est pas disponible pour cette période.")
		fmt.Println("----------------------------------------------------")
		return nil
	}

	reservation := Reservation{
		RoomID:    roomID,
		StartTime: startTime.Format("2006-01-02 15:04:05"), // Convert time.Time to string in the desired format
		EndTime:   endTime.Format("2006-01-02 15:04:05"),   // Convert time.Time to string in the desired format
	}

	err = createReservation(reservation)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("La réservation n'a pas pu être crée.")
		fmt.Println("----------------------------------------------------")
		fmt.Println("")
		return err
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println("La réservation a été créée avec succès.")
	fmt.Println("----------------------------------------------------")
	fmt.Println("")
	return nil
}

// Function to list reservations
func listReservations() error {
    var roomID int
    var err error

    // Demandez à l'utilisateur l'identifiant de la salle
    reader := bufio.NewReader(os.Stdin)
	UNBUFFER, _ := reader.ReadString('\n')
	UNBUFFER = strings.TrimSpace(UNBUFFER)

    fmt.Print("Entrez l'identifiant de la salle (mettez 0 pour afficher toutes les réservations) : ")
    roomIDInput, _ := reader.ReadString('\n')
    roomIDInput = strings.TrimSpace(roomIDInput)
    roomID, err = strconv.Atoi(roomIDInput)
    if err != nil {
        fmt.Println("----------------------------------------------------")
        fmt.Println("ID de salle invalide.")
        fmt.Println("----------------------------------------------------")
        return err
    }

    rows, err := getReservations(roomID)

    // Parcourir les lignes résultantes et afficher les réservations
    // Vérifier si des réservations ont été trouvées
    if roomID == 0 {
        fmt.Println("")
        fmt.Println("Réservations pour toutes les salles :")
        fmt.Println("----------------------------------------------------")
    } else {
        fmt.Println("")
        fmt.Printf("Réservations pour la salle %d :\n", roomID)
        fmt.Println("----------------------------------------------------")
    }

    for _, reservation := range rows {
        if roomID == 0 {
            fmt.Printf(" %d. Salle: %d - Début: %s - Fin: %s\n", reservation.ID, reservation.RoomID, reservation.StartTime, reservation.EndTime)
        } else {
            fmt.Printf(" %d. Début: %s - Fin: %s\n", reservation.ID, reservation.StartTime, reservation.EndTime)
        }
    }
    fmt.Println("")
    return nil
}

// Function to cancel a reservation from user input
func cancelReservationFromUserInput() error {
	fmt.Println("")
    fmt.Println("----------------------------------------------------")
    fmt.Println("Annulation d'une réservation")
    fmt.Println("----------------------------------------------------")

    reader := bufio.NewReader(os.Stdin)

	UNBUFFER, _ := reader.ReadString('\n')
	UNBUFFER = strings.TrimSpace(UNBUFFER)

    fmt.Print("Entrez l'ID de la réservation à annuler : ")
    reservationIDInput, _ := reader.ReadString('\n')
    reservationIDInput = strings.TrimSpace(reservationIDInput)
    reservationID, err := strconv.Atoi(reservationIDInput)
    if err != nil {
        fmt.Println("----------------------------------------------------")
        fmt.Println("ID de réservation invalide.")
        fmt.Println("----------------------------------------------------")
        return err
    }

    // Annulez la réservation
    err = cancelReservation(reservationID)
    if err != nil {
        fmt.Println("----------------------------------------------------")
        fmt.Println("La réservation n'a pas pu être annulée.")
        fmt.Println("----------------------------------------------------")
        fmt.Println("")
        return err
    }

    return nil
}

func exportReservationsToJSON() error {
    // Récupérer toutes les réservations
    allReservations, err := getReservations(0)
    if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Erreur de récupération des réservations depuis la base de données \n")
		fmt.Println("----------------------------------------------------")
        return err
    }

    // Convertir les réservations en JSON
    reservationsJSON, err := json.MarshalIndent(allReservations, "", "    ")
    if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Erreur de conversion vers le format JSON \n")
		fmt.Println("----------------------------------------------------")
        return err
    }

    // Générer un nom de fichier unique basé sur le temps actuel
    currentTime := time.Now()
    filename := fmt.Sprintf("export_datas/reservations_%s.json", currentTime.Format("2006-01-02_15-04-05"))

    err = os.WriteFile(filename, reservationsJSON, 0644)
    if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Les réservations ont été exportées avec succès dans le fichier %s\n", filename)
		fmt.Println("----------------------------------------------------")
        return err
    }
	fmt.Println("")
    fmt.Println("----------------------------------------------------")
    fmt.Printf("Les réservations ont été exportées avec succès dans le fichier %s\n", filename)
    fmt.Println("----------------------------------------------------")
    return nil
}
