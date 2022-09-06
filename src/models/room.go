package models

type Room struct {
	ID    string `json:"room_id"`
	User1 User   `json:"user1"`
	User2 User   `json:"user2"`
}
