package repository

import (
	"database/sql"
	"errors"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/util"
	"log"
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

type Session struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	refreshToken string    `db:"refresh_token"`
	expiresAt    time.Time `db:"expires_at"`
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

func (r *AuthRepository) GetSessionByRefreshToken(refreshToken string) (int, error) {
	var sessionID Session
	err := r.db.Get(sessionID, "SELECT * FROM sessions WHERE refresh_token = $1", refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, errors.New("session not found")
		}
		return -1, err
	}
	return sessionID.ID, nil
}

func (r *AuthRepository) GetUserIdBySessionId(id int) (int, string, error) {
	var session Session
	err := r.db.Get(session, "SELECT * FROM sessions WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, "", errors.New("session not found")
		}
		return -1, "", err
	}
	role, err := util.ParseToken(session.refreshToken)
	if err != nil {
		log.Print("Error parsing the token")
	}
	return session.UserID, role.Role, nil
}
