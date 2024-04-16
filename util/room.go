package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func AddRoom(name string, capacity int) error {
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

func GetRooms() ([]Room, error) {
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

func listRooms() {
    rooms, err := GetRooms() // Récupérer les salles de la base de données
    if err != nil {
        fmt.Println("Erreur lors de la récupération des salles : ", err)
    }

    fmt.Println("Salles disponibles :")
    for _, room := range rooms {
        fmt.Printf("- %s (Capacité : %d)\n", room.Name, room.Capacity)
    }
}