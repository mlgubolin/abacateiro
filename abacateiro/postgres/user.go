package postgres

import (
	"abacateiro"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(user abacateiro.User) (abacateiro.User, error) {
	return abacateiro.User{}, nil
}
func (s *UserService) GetUser(id int) (abacateiro.User, error) {
	return abacateiro.User{}, nil
}
func (s *UserService) GetUsers() ([]abacateiro.User, error) {
	return []abacateiro.User{}, nil
}
func (s *UserService) UpdateUser(user abacateiro.User) (abacateiro.User, error) {
	return abacateiro.User{}, nil
}
func (s *UserService) DeleteUser(id int) error {
	return nil
}
