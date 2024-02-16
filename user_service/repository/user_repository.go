package repository

import (
	"database/sql"
	"errors"
	"github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *UserRepository) CreateUser(username, passwordHash string, role string) (int32, error) {
	var id int32
	err := r.db.QueryRow("INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id",
		username, passwordHash, role).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (r *UserRepository) GetUserByID(userID int) (*user_service.UserResponse, error) {
	var user user_service.UserResponse
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers() (user_service.UsersResponse, error) {
	var users user_service.UsersResponse
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return user_service.UsersResponse{}, err
	}
	if err := rows.Scan(&users); err != nil {
		return user_service.UsersResponse{}, err
	}
	return users, nil
}

type GetUserByPropertyResponse struct {
	Id           int32  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
	Role         string `db:"role"`
}

func (r *UserRepository) GetUserByProperty(username string) (*GetUserByPropertyResponse, error) {
	var user GetUserByPropertyResponse
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id int, newUsername *string, role *string) error {
	// Construct the SQL query
	query := "UPDATE users SET"
	params := []interface{}{}

	// Check if newUsername is provided
	if newUsername != nil {
		query += " username = ?,"
		params = append(params, *newUsername)
	}

	// Check if role is provided
	if role != nil {
		query += " role = ?,"
		params = append(params, *role)
	}

	// Trim the trailing comma from the query
	query = strings.TrimSuffix(query, ",")

	// Add WHERE clause for the user id
	query += " WHERE id = ?"
	params = append(params, id)

	// Execute the SQL query
	_, err := r.db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	// Construct the SQL query
	query := "DELETE FROM users WHERE id = ?"

	// Execute the SQL query
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
