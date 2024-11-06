package user

import (
	"encoding/json"
	"net/http"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/service"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/Sairam-04/blog-app/backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid payload request")
		return
	}

	if err := validator.New().Struct(user); err != nil {
		// typecast err to validator
		validateErrs := err.(validator.ValidationErrors)
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.ValidationError(validateErrs))
		return
	}

	token, err := h.userService.RegisterUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.RespondWithJSON(w, http.StatusCreated, &types.UserResponse{
		Success: true,
		Token:   token,
		Message: "User Registered Successfully",
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid payload request")
		return
	}
	token, err := h.userService.LoginUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, &types.UserResponse{
		Success: true,
		Token:   token,
		Message: "User LoggedIn Successfully",
	})
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(types.UserIDKey{}).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "user is not authorized")
		return
	}
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	users, err := h.userService.GetUser(parsedUserId)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "user not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusOK, users)

}
