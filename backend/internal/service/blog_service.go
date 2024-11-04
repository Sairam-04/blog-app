package service

import (
	"fmt"
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/google/uuid"
)

type BlogService struct {
	blogRepo domain.BlogRepository
}

func NewBlogService(blogRepo domain.BlogRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

func (s *BlogService) CreateNewBlog(userId uuid.UUID, blog *domain.Blog) error {
	blog.ID = uuid.New()
	blog.UserID = userId
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	ok, err := s.blogRepo.CreateBlog(blog)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to post a blog")
	}
	return nil
}

func (s *BlogService) GetAllBlogs(limit, offset int) ([]types.BlogsResponse, error) {
	blogs, err := s.blogRepo.GetBlogs(limit, offset)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (s *BlogService) UserBlogs(userId uuid.UUID) ([]types.BlogResponse, error) {
	blogs, err := s.blogRepo.GetUserBlogs(userId)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (s *BlogService) UdpateBlogService(userId uuid.UUID, blogId uuid.UUID, blog *types.UpdateBlogReq) error {
	ok, err := s.blogRepo.UpdateBlogByID(userId, blogId, blog)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("unable to update the blog")
	}
	return nil
}
