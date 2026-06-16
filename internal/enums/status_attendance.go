package enums

import "log"

type StatusAttendance string

const (
	Present    = StatusAttendance("PRESENT")
	Late       = StatusAttendance("LATE")
	Absent     = StatusAttendance("ABSENT")
	Sick       = StatusAttendance("SICK")
	Permission = StatusAttendance("PERMISSION")
)

func (sa StatusAttendance) IsValid() bool {

	switch sa {
	case Present, Late, Absent, Sick, Permission:
		return true
	default:
		log.Println("invalid status_attendance")
		return false
	}
}
