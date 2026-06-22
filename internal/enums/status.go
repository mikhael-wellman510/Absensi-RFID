package enums

import "log"

type Status string

const (
	Active      = Status("ACTIVE")
	Graduated   = Status("GRADUATED")
	Transferred = Status("TRANSFERRED")
	Inactive    = Status("INACTIVE")
)

func (s Status) IsValid() bool {

	switch s {
	case Active, Graduated, Transferred, Inactive:
		return true
	default:
		log.Println("invalid status")
		return false
	}

}
