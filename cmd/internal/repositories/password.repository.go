package repositories

import (
	"database/sql"

	"github.com/davi-sant/asafe-vault-go/cmd/internal/models"
)

type PasswordRepository interface {
	Create(PayloadPassword models.Password) error
	GetPasswordByServiceName(user_id int64, service_name string) ([]models.Password, error)
	GetAll(userID int64) ([]models.Password, error)
}

type PostgresPasswordRepository struct {
	db *sql.DB
}

func NewPostgresPasswordRepository(db *sql.DB) *PostgresPasswordRepository {
	return &PostgresPasswordRepository{db: db}
}

func (p *PostgresPasswordRepository) GetPasswordByServiceName(user_id int64, service_name string) ([]models.Password, error) {
	query := `SELECT id, user_id, service_name, service_username, service_password FROM passwords WHERE user_id = $1 AND service_name ILIKE $2`

	rows, err := p.db.Query(query, user_id, service_name)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var passwords []models.Password

	for rows.Next() {
		var password models.Password
		err := rows.Scan(&password.Id, &password.UserId, &password.ServiceName, &password.ServiceUserName, &password.ServicePassword)

		if err != nil {
			return nil, err
		}

		passwords = append(passwords, password)
	}
	return passwords, nil
}

func (p *PostgresPasswordRepository) GetAll(userID int64) ([]models.Password, error) {

	query := `SELECT id, user_id, service_name, service_username, service_password FROM passwords WHERE user_id = $1`
	rows, err := p.db.Query(query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []models.Password
	for rows.Next() {
		var password models.Password
		err := rows.Scan(&password.Id, &password.UserId, &password.ServiceName, &password.ServiceUserName, &password.ServicePassword)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return passwords, nil

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
