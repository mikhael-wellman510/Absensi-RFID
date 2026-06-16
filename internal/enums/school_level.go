package enums

import "log"

type SchoolLevel string

const (
	SD  = SchoolLevel("SD")
	SMP = SchoolLevel("SMP")
	SMA = SchoolLevel("SMA")
)

func (sl SchoolLevel) IsValid() bool {
	switch sl {
	case SD, SMP, SMA:
		return true
	default:
		log.Println("invalid school level", sl)
		return false
	}
}
