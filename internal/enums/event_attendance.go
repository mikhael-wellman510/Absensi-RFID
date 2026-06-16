package enums

import "log"

type EventAttendance string

const (
	CheckIn  = EventAttendance("CHECK_IN")
	CheckOut = EventAttendance("CHECK_OUT")
)

func (ea EventAttendance) IsValid() bool {

	switch ea {
	case CheckIn, CheckOut:
		return true
	default:
		log.Println("invalid event_attendance")
		return false
	}
}
