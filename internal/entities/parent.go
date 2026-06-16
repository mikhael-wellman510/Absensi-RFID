package entities

/*
Entity Updated
*/
type (
	Parent struct {
		Base
		UserId     string `json:"userId" gorm:"column:user_id;not null"`
		Occupation string `json:"occupation" gorm:"column:occupation;not null"`
		Address    string `json:"address" gorm:"column:address;not null"`
		User       User   `gorm:"foreignKey:UserId"`
	}

	ParentRequest struct {
		Id         string `json:"id"`
		UserId     string `json:"userId"`
		Occupation string `json:"occupation"`
		Address    string `json:"address"`
	}
	ParentResponse struct {
		Id           string       `json:"id"`
		UserResponse UserResponse `json:"userResponse"`
		Occupation   string       `json:"occupation"`
		Address      string       `json:"address"`
	}
)
