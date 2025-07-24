package repository

import (
	"database/sql"
	"user-service/model"
)

type UserRepository interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUserByID(id int64) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id)
	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return &user, err
}

func (r *userRepo) CreateUser(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (r *userRepo) UpdateUser(user *model.User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	return err
}

func (r *userRepo) DeleteUser(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
