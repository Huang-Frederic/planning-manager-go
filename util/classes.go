package util

// Strcuture for rooms
type Room struct {
	ID       int
	Name     string
	Capacity int
}

// Structure for reservations
type Reservation struct {
	ID        int
	RoomID    int
	StartTime string
	EndTime   string
}