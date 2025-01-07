package repositories

import (
	"database/sql"
	"errors"

	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
)

type UserRepository interface {
	Create(user models.User) error
	GetUserBayEmail(email string) (*models.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user models.User) error {
	query := `INSERT INTO users (email, password) VALUES($1, $2)`
	_, err := r.db.Exec(query, user.Email, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetUserBayEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func InitializeDatabase(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
	_, err := db.Exec(query)
	return err
}
