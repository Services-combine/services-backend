package combine

type User struct {
	Id       int    `json:"-" db:"_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
