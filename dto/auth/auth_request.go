package authdto

type Request_Register struct {
	Name     string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}

type Request_Login struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
