package repositories

import (
	"database/sql"
	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
)

type PasswordRepository interface {
	Create(PayloadPassword models.Password) error
	GetPasswordByUsername(username string) (models.Password, error)
}

type PostgresPasswordRepository struct {
	db *sql.DB
}

func (p *PostgresPasswordRepository) GetPasswordByUsername(username string) (models.Password, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostgresPasswordRepository(db *sql.DB) *PostgresPasswordRepository {
	return &PostgresPasswordRepository{db: db}
}

func (p *PostgresPasswordRepository) Create(payload models.Password) error {
	query := `INSERT INTO passwords (user_id, service_name, service_username, service_password) VALUES ($1, $2, $3, $4)`
	_, err := p.db.Exec(query, payload.UserId, payload.ServiceName, payload.ServiceUserName, payload.ServicePassword)

	if err != nil {
		return err
	}
	return nil
}

func InitializePasswordDB(db *sql.DB) error {

	query := `
		CREATE TABLE IF NOT EXISTS passwords (
        id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL,
    	service_name VARCHAR(255) NOT NULL,
    	service_username VARCHAR(255),
    	service_password VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);`
	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
