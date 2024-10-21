package postgres

import (
	"context"
	"stugi/gonews/pkg/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New(connection string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT 
		p.id, 
		p.author_id, 
		a.name as author_name, 
		p.title, 
		p.content, 
		p.created_at, 
		p.published_at 
	FROM posts p
	LEFT JOIN authors a ON p.author_id = a.id`)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.AuthorName, &p.Title, &p.Content, &p.CreatedAt, &p.PublishedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (s *Storage) AddPost(post storage.Post) error {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at, published_at)
		VALUES ($1, $2, $3, $4, $5)`,
		post.AuthorID, post.Title, post.Content, post.CreatedAt, post.PublishedAt).Scan(&post.ID)

	return err
}
func (s *Storage) UpdatePost(post storage.Post) error {
	err := s.db.QueryRow(context.Background(), `
		UPDATE posts
		SET author_id = $1, title = $2, content = $3, created_at = $4, published_at = $5
		WHERE id = $6`,
		post.AuthorID, post.Title, post.Content, post.CreatedAt, post.PublishedAt, post.ID).Scan(&post.ID)

	return err
}
func (s *Storage) DeletePost(ID int) error {
	err := s.db.QueryRow(context.Background(), `
		DELETE FROM posts
		WHERE id = $1`,
		ID).Scan(&ID)

	return err
}
