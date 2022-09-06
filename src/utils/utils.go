package utils

import "ws-3xt/src/models"

func GetRoomById(rooms []models.Room, roomID string) (models.Room, bool) {
	for _, r := range rooms {
		if r.ID == roomID {
			return r, true
		}

	}
	return models.Room{}, false
}
