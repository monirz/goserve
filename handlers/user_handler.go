package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/monirz/goserve"
	"github.com/monirz/goserve/response"
	"go.uber.org/zap"
)

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &goserve.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		s.logger.Error("invalid json input",
			zap.String("error", err.Error()),
		)
		response.ResponseError(w, http.StatusBadRequest, "invalid json input", err.Error())
		return
	}

	// validate input
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		s.logger.Error("validation error",
			zap.String("error", err.Error()),
		)

		response.ResponseError(w, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	user.UUID = uuid.NewString()

	id, err := s.UserService.CreateUser(user)
	if err != nil {
		s.logger.Error("error creating user",
			zap.String("error", err.Error()),
		)

		response.ResponseError(w, http.StatusInternalServerError, "error creating user", err.Error())
		return
	}

	s.logger.Info("user created successfully",
		zap.Int("id", int(id)),
		zap.String("uuid", user.UUID),
	)

	response.JSONResponse(w, "user created successfully", http.StatusCreated, 201, &goserve.User{
		UUID:  user.UUID,
		Email: user.Email,
	})

}
