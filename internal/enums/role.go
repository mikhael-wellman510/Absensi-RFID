package enums

import "log"

type Role string

const (
	SuperAdmin  = Role("SUPER_ADMIN")
	SchoolAdmin = Role("SCHOOL_ADMIN")
	Teacher     = Role("TEACHER")
	Parent      = Role("PARENT")
)

func (r Role) IsValid() bool {
	switch r {
	case SuperAdmin, SchoolAdmin, Teacher, Parent:
		return true
	default:
		log.Println("invalid role")
		return false
	}
}
