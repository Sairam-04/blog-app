package repository

import (
	"database/sql"
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/google/uuid"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) CreateBlog(blog *domain.Blog) (bool, error) {
	query := `INSERT INTO blogs (id, userId, title, description, content, thumbnail, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := r.db.Exec(query, blog.ID, blog.UserID, blog.Title, blog.Description, blog.Content, blog.Thumbnail, blog.CreatedAt, blog.UpdatedAt)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BlogRepository) GetBlogs(limit, offset int) ([]types.BlogsResponse, error) {
	query := `SELECT 
				b.id, 
				u.name, 
				b.title, 
				b.description, 
				b.content, 
				b.thumbnail, 
				b.created_at, 
				b.updated_at
			FROM 
				blogs b
			JOIN 
				users u ON b.userId = u.id 
			ORDER BY 
				b.updated_at DESC 
			LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []types.BlogsResponse
	for rows.Next() {
		var blog types.BlogsResponse

		if err := rows.Scan(&blog.ID, &blog.Name, &blog.Title, &blog.Description, &blog.Content, &blog.Thumbnail, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) GetUserBlogs(userId uuid.UUID) ([]types.BlogResponse, error) {
	query := `SELECT id, title, description, content, thumbnail, created_at, updated_at
				FROM blogs b WHERE b.userId = $1 ORDER BY b.updated_at DESC;`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []types.BlogResponse
	for rows.Next() {
		var blog types.BlogResponse

		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Content, &blog.Thumbnail, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) UpdateBlogByID(userId, blogId uuid.UUID, blog *types.UpdateBlogReq) (bool, error) {
	query := `UPDATE blogs SET
			  title = $1, description = $2, 
			  content = $3, thumbnail = $4, updated_at = $5
			  WHERE 
			  userId = $6 and id = $7;`
	_, err := r.db.Exec(query, blog.Title, blog.Description, blog.Content, blog.Thumbnail, time.Now(), userId, blogId)
	if err != nil {
		return false, err
	}
	return true, nil
}
