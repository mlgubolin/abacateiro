package abacateiro

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
	Document string `json:"user_document"`
}

// DTO de saída sem o campo Password
type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Document string `json:"user_document"`
}

type UserUpdate struct {
	Name        *string `json:"name"`
	DisplayName *string `json:"display_name"`
	Email       *string `json:"email"`
	Activated   *bool   `json:"activated"`
}

const (
	MinUsernameLen = 3
	MaxUsernameLen = 30
	MinPasswordLen = 6
	MaxPasswordLen = 50
)

var EmailRegex = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)

func (u *User) Verify() error {
	switch {
	case u.Name == "":
		return errors.New("nome é necessário")
	case utf8.RuneCountInString(u.Password) < MinPasswordLen:
		return errors.New("senha muito pequena")
	case utf8.RuneCountInString(u.Password) > MaxPasswordLen:
		return errors.New("senha muito grande")
	case u.Email == "" || !EmailRegex.MatchString(u.Email):
		return errors.New("email inválido")
	}
	return nil
}

// Função para converter um único User para UserResponse
func ToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Document: user.Document,
	}
}

// Função para converter um slice de User para um slice de UserResponse
func ToUserResponses(users []User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

type UserService interface {
	CreateUser(user User) (User, error)
	GetUser(id int) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) error
}
