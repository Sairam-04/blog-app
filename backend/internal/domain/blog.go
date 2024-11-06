package domain

import (
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/google/uuid"
)

type Blog struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Thumbnail   string    `json:"Thumbnail" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type BlogRepository interface {
	CreateBlog(blog *Blog) (bool, error)
	GetBlogs(limit, offset int) ([]types.BlogsResponse, error)
	GetUserBlogs(userId uuid.UUID) ([]types.BlogResponse, error)
	UpdateBlogByID(userId, blogId uuid.UUID, blog *types.UpdateBlogReq) (bool, error)
	GetBlogByID(blogId uuid.UUID) (*types.BlogsResponse, error)
	Search(keyword string, limit, offset int) ([]types.BlogsResponse, error)
	DeleteBlog(userId, blogId uuid.UUID) error
}
