package repository

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type UsersRepo interface {
	GetAllUsers() ([]User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(username string, password string) (*User, error)
}
