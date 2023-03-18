package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/monirz/goserve"
)

type mockUserService struct {
	users map[int]goserve.User
}

func (m *mockUserService) CreateUser(user *goserve.User) (int, error) {

	if user.Email == "" {
		return 0, errors.New("email cannot be empty")
	}
	if user.Password == "" {
		return 0, errors.New("password cannot be empty")
	}

	id := 123
	user.UUID = "123e4567-e89b-12d3-a456-426655440000"
	user.Email = "test@example.com"

	m.users[id] = *user

	return id, nil
}

func (m *mockUserService) FindUserByID(uuid string) (*goserve.User, error) {
	for _, user := range m.users {
		if user.UUID == uuid {
			return &user, nil
		}
	}
	return nil, nil
}

func (m *mockUserService) GetUsers() (*[]goserve.User, error) {
	users := []goserve.User{}
	for _, v := range m.users {
		users = append(users, v)
	}
	return &users, nil
}

func TestCreateUserHandler(t *testing.T) {
	mockService := &mockUserService{users: make(map[int]goserve.User)}

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	s := NewServer(db)
	s.UserService = mockService

	t.Run("CreateUserHandler with valid request", func(t *testing.T) {
		user := &goserve.User{
			UUID:     "123e4567-e89b-12d3-a456-426655440000",
			Email:    "test@example.com",
			Password: "password",
		}

		jsonStr, _ := json.Marshal(user)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(s.CreateUserHandler)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d but got %d", http.StatusCreated, rr.Code)
		}

		var respData map[string]interface{}
		if err = json.Unmarshal(rr.Body.Bytes(), &respData); err != nil {
			t.Fatal(err)
		}

		respUser, ok := respData["data"].(map[string]interface{})
		if !ok {
			t.Error("could not parse response data")
		}

		if respUser["email"] != user.Email {
			t.Errorf("expected email %s but got %v", user.Email, respUser["email"])
		}

	})

	//TODO add more test case
}
