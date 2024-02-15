package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(username, passwordHash string, role string) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		username, passwordHash, role)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID int) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers() ([]User, error) {
	var users []User
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	if err := rows.Scan(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
