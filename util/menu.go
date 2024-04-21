package util

import "fmt"

// Main Menu
func menu() bool {
	fmt.Println(" ")
	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Exporter les réservations en JSON")
	fmt.Println("6. Quitter")
	fmt.Println(" ")
	fmt.Printf("Choisissez une option : ")

	return userChoice("mainMenu", askUserChoice())
}

// Ask user choice
func askUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

// Sub Menu
func subMenu() bool {
	userInput := 0
	for userInput < 1 || userInput > 2 {
		fmt.Println("1. Retourner au menu principal")
		fmt.Println("2. Quitter")
		fmt.Println(" ")
		fmt.Printf("Choisissez une option : ")
		userInput = askUserChoice()
		if userInput < 1 || userInput > 2 {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Erreur, votre numéro est invalide, veuillez resaisir")
			fmt.Println("----------------------------------------------------")
		}
	}
	return userChoice("subMenu", userInput)
}

// User choice
func userChoice(menu string, choice int) bool {
	switch menu {
	case "mainMenu":
		switch choice {
		case 1:
			listRooms()
			return subMenu()
		case 2:
			createReservationFromUserInput()
			return subMenu()
		case 3:
			cancelReservationFromUserInput()
			return subMenu()
		case 4:
			listReservations()
			return subMenu()
		case 5:
			exportReservationsToJSON()
			return subMenu()
		case 6:
			return false
		default:
			fmt.Println("----------------------------------------------------")
			fmt.Println("Erreur dans la saisie, veuillez re-taper votre chiffre")
			fmt.Println("----------------------------------------------------")
			return true
		}
	case "subMenu":
		switch choice {
		case 1:
			return true
		case 2:
			return false
		default:
			// Should never happens
			fmt.Println("----------------------------------------------------")
			fmt.Println("ERROR, subMenu DEFAULT")
			fmt.Println("----------------------------------------------------")
			return true
		}
	default:
		// Should never happens
		fmt.Println("----------------------------------------------------")
		fmt.Println("ERROR, userChoice DEFAULT")
		fmt.Println("----------------------------------------------------")
		return true
	}
}