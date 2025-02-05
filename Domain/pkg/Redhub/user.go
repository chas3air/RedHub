package redhub

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v6();primary_key"`
	Login    string    `json:"login" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
	Role     string    `json:"role" gorm:"not null"`
	Nick     string    `json:"nick" gorm:"not null"`
}

func NewUser(login, password, role, nick string) *User {
	return &User{
		ID:       uuid.New(),
		Login:    login,
		Password: password,
		Role:     role,
		Nick:     nick,
	}
}

func (u *User) GetID() uuid.UUID   { return u.ID }
func (u *User) SetID(id uuid.UUID) { u.ID = id }

func (u *User) GetLogin() string      { return u.Login }
func (u *User) SetLogin(login string) { u.Login = login }

func (u *User) GetPassword() string         { return u.Password }
func (u *User) SetPassword(password string) { u.Password = password }

func (u *User) GetRole() string     { return u.Role }
func (u *User) SetRole(role string) { u.Role = role }

func (u *User) GetNick() string     { return u.Nick }
func (u *User) SetNick(nick string) { u.Nick = nick }
