package categoriesdto


type Response_Categories struct {
	Id          int    `json:"id" gorm:"type: int;PRIMARY_KEY"`
	Name        string `json:"name"`
}
