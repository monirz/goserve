package goserve

type UserService interface {
	CreateUser(*User) (int, error)
	FindUserByID(uuid string) (*User, error)
	GetUsers() (*[]User, error)
}

type User struct {
	ID       int64  `json:"id,omitempty"`
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
