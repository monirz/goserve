package postgres

import (
	"database/sql"

	"github.com/monirz/goserve"
)

var _ goserve.UserService = (*UserService)(nil)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(*goserve.User) (int, error) {
	return 0, nil
}

func (u *UserService) FindUserByID(uuid string) (*goserve.User, error) {
	return nil, nil
}
func (u *UserService) GetUsers() (*[]goserve.User, error) {

	return nil, nil
}
