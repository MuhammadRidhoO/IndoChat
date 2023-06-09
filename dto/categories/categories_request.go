package categoriesdto

type Request_Categories struct {
	Id          int
	Name        string `json:"name" validate:"required"`
}
