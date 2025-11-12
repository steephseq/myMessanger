package storage

import (
	"auth-service/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

func (p *PostgresRepo) GetUserID(ul models.UserLogin) (string, error) {
	query := `SELECT 
		id 
		FROM auth_users 
		WHERE email=$1`
	var id string
	if err := p.db.Get(&id, query, ul.Email); err != nil {
		return "", err
	}
	return id, nil
}

func (p *PostgresRepo) GetUserPassword(ul models.UserLogin) (string, error) {
	query := `SELECT 
		password 
		FROM auth_users 
		WHERE email=$1`
	var password string
	if err := p.db.Get(&password, query, ul.Email); err != nil {
		return "", err
	}
	return password, nil
}

func (p *PostgresRepo) CreateUser(ur models.UserRegister) error {
	query := `INSERT
		INTO auth_users
		VALUES($1,$2,$3)`
	_, err := p.db.Exec(query, ur.Email, ur.Password, ur.Username)
	return err
}

func (p *PostgresRepo) CheckPasswordHash(ul models.UserLogin) error {
	storedPassword, err := p.GetUserPassword(ul)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(ul.Password))
}

func (p *PostgresRepo) ExistsUser(ur models.UserRegister) (bool, error) {
	query := `SELECT EXISTS(
		SELECT 1
		FROM auth_users 
		WHERE email=$1 OR username=$2)`
	var exists bool
	if err := p.db.Get(&exists, query, ur.Email, ur.Username); err != nil {
		return false, err
	}
	return exists, nil
}
