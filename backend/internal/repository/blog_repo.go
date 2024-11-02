package repository

import (
	"database/sql"
	"fmt"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/internal/types"
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
				b.created_at 
			FROM 
				blogs b
			JOIN 
				users u ON b.userId = u.id 
			ORDER BY 
				b.created_at DESC 
			LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []types.BlogsResponse
	for rows.Next() {
		var blog types.BlogsResponse

		if err := rows.Scan(&blog.ID, &blog.Name, &blog.Title, &blog.Description, &blog.Content, &blog.Thumbnail, &blog.CreatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}
