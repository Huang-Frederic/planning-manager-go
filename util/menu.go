package util

import "fmt"

func menu() bool {
	fmt.Println(" ")
	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Quitter")
	fmt.Println(" ")
	fmt.Printf("Choisissez une option : ")

	return userChoice("mainMenu", askUserChoice())
}

func askUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func subMenu() bool {
	userInput := 0
	for userInput < 1 || userInput > 2 {
		fmt.Println("1. Retourner au menu principal")
		fmt.Println("2. Quitter")
		fmt.Println(" ")
		fmt.Printf("Choisissez une option : ")
		userInput = askUserChoice()
		if userInput < 1 || userInput > 2 {
			fmt.Println("Erreur, votre numéro est invalide, veuillez resaisir")
		}
	}
	return userChoice("subMenu", userInput)
}

func userChoice(menu string, choice int) bool {
	switch menu {
	case "mainMenu":
		switch choice {
		case 1:
			listRooms()
			fmt.Println("List")
			return subMenu()
		case 2:
			CreateReservationFromUserInput()
			fmt.Println("Create")
			return subMenu()
		case 3:
			fmt.Println("Cancel")
			return subMenu()
		case 4:
			fmt.Println("View")
			return subMenu()
		case 5:
			return false
		default:
			fmt.Println("Erreur dans la saisie, veuillez re-taper votre chiffre")
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
			fmt.Println("ERROR, subMenu DEFAULT")
			return true
		}
	default:
		// Should never happens
		fmt.Println("ERROR, userChoice DEFAULT")
		return true
	}
}