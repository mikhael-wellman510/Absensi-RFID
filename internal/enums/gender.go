package enums

import "log"

type Gender string

const (
	GenderMale   = Gender("Male")
	GenderFemale = Gender("Female")
)

func (gender Gender) IsValid() bool {
	switch gender {
	case GenderMale, GenderFemale:
		return true
	default:
		log.Println("Gender is invalid")
		return false
	}
}
