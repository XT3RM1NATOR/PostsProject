package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateSession(userID int, refreshToken string, expiresAt time.Time) error {
	_, err := r.db.Exec("INSERT INTO sessions (user_id, refresh_token, expires_at) VALUES ($1, $2, $3)",
		userID, refreshToken, expiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetUserIDByRefreshToken(refreshToken string) (int, error) {
	var userID int
	err := r.db.Get(&userID, "SELECT user_id FROM sessions WHERE refresh_token = $1", refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("User not found")
		}
		return 0, err
	}
	return userID, nil
}

func (r *AuthRepository) DeleteSession(sessionID int) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE id = $1", sessionID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) UpdateSession(sessionID int, expiresAt time.Time) error {
	_, err := r.db.Exec("UPDATE sessions SET expires_at = $1 WHERE id = $2", expiresAt, sessionID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*User, error) {
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

func (r *AuthRepository) GetSessionByRefreshToken(refreshToken string) (int, error) {
	var sessionID int
	err := r.db.Get(sessionID, "SELECT * FROM sessions WHERE refresh_token = $1", refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, errors.New("session not found")
		}
		return -1, err
	}
	return sessionID, nil
}

func (r *AuthRepository) CreateUser(username, email, passwordHash string, role string) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)",
		username, email, passwordHash, role)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetUserByID(userID int) (*User, error) {
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
