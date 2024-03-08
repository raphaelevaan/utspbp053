package models

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	GameID   int    `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type Participant struct {
	ID        int `json:"id"`
	RoomID    int `json:"id_room"`
	AccountID int `json:"id_account"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type Game struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"max_player"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
