package repository

import "github.com/jmoiron/sqlx"

type userRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepoDB(db *sqlx.DB) *userRepoImpl {
	return &userRepoImpl{
		db: db,
	}
}

func (u *userRepoImpl) GetAllUsers() ([]User, error) {
	users := []User{}
	query := "select id, username, password from users"
	err := u.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepoImpl) GetUserByUsername(username string) (*User, error) {
	user := User{}
	query := "select id, username, password from users where username=?"
	err := u.db.Select(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepoImpl) CreateUser(username string, password string) (*User, error) {
	query := "insert into users (username, password) values (?, ?)"
	result, err := u.db.Exec(query, username, password)
	if err != nil {
		return nil, err
	}
	created_id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	created_user := User{
		Id:       int(created_id),
		Username: username,
		Password: password,
	}
	return &created_user, nil
}
