package template

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	apperror2 "template/app/internal/apperror"
	"template/app/pkg/logging"
)

const (
	usersURL = "/api/v1/users"
)

type Handler struct {
	Logger      logging.Logger
	UserService Service
	Validator   *validator.Validate
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(usersURL, apperror2.Middleware(h.CreateUser)).
		Methods(http.MethodPost)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("CreateUser")

	w.Header().Set("Content-Type", "application/json")

	var userDto CreateUserDTO
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		return apperror2.BadRequestError("Error deserializing JSON")
	}

	err = h.Validator.Struct(userDto)

	if err != nil {
		return apperror2.BadRequestError("Error validating JSON")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = h.UserService.Create(r.Context(), *userDto.Hashed(hashedPassword))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)

	return nil
}
