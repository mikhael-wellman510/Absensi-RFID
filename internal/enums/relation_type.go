package enums

import "log"

type RelationType string

const (
	Father   = RelationType("FATHER")
	Mother   = RelationType("MOTHER")
	Guardian = RelationType("GUARDIAN")
)

func (rt RelationType) IsValid() bool {
	switch rt {
	case Father, Mother, Guardian:
		return true
	default:
		log.Println("Invalid RelationType : ", rt)
		return false
	}
}
