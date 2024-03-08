package models

type Game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
}

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Participant struct {
	ID        int `json:"id"`
	RoomID    int `json:"room_id"`
	AccountID int `json:"account_id"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

// ParticipantsResponse struct represents the response format for participants
type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}
