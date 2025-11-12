package storage

import (
	"user-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(connString string) (*PostgresRepo, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{db: db}, nil
}

func (p *PostgresRepo) AddUserIntoDB(u models.User) error {
	query := `INSERT 
		INTO users 
		(id, email, password, name, username, last_seen, avatar_url, bio)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	_, err := p.db.Exec(query, u.ID, u.Email, u.Password, u.Name, u.Username, u.LastSeen, u.AvatarURL, u.Bio)
	return err
}
