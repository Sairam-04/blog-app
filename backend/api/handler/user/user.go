package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/service"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/Sairam-04/blog-app/backend/utils"
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
	fmt.Println("Logging in a User..")
}
