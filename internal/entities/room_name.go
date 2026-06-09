package entities

type (
	RoomName struct {
		Base
		RoomName string `json:"roomName" gorm:"column:room_name;not null"`
	}
)
