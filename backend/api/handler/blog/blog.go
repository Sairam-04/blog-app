package blog

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/service"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/Sairam-04/blog-app/backend/utils"
	"github.com/google/uuid"
)

type BlogHandler struct {
	blogService *service.BlogService
}

func NewBlogHandler(blogService *service.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

func (h *BlogHandler) PostBlog(w http.ResponseWriter, r *http.Request) {
	var blog domain.Blog
	userID, ok := r.Context().Value(types.UserIDKey{}).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "user is not authorized")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid payload request")
		return
	}
	parsedUserId, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	err = h.blogService.CreateNewBlog(parsedUserId, &blog)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, &types.GeneralResponse{
		Success: true,
		Message: "Blog Created Successfully",
		Error:   "",
	})
}

func (h *BlogHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || limit <= 0 {
		offset = 0
	}
	blogs, err := h.blogService.GetAllBlogs(limit, offset)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &types.GetAllBlogsResp{
		Success: true,
		Message: "Fetched All blogs",
		Error:   "",
		Blogs:   blogs,
	})
}
