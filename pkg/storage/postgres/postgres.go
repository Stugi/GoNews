package postgres

import (
	"context"
	"fmt"
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
		fmt.Println(err)
		return nil, err
	}
	s := Storage{
		db: db,
	}
	fmt.Println("Connected to Postgres!")
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
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at, published_at)
		VALUES ($1, $2, $3, $4, $5)`,
		post.AuthorID, post.Title, post.Content, post.CreatedAt, post.PublishedAt)

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (s *Storage) UpdatePost(post storage.Post) error {
	query := `
		UPDATE posts
		SET author_id = $1, title = $2, content = $3, created_at = $4, published_at = $5
		WHERE id = $6`
	// TODO: How update only field changed
	_, err := s.db.Exec(context.Background(), query,
		post.AuthorID, post.Title, post.Content, post.CreatedAt, post.PublishedAt, post.ID)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
func (s *Storage) DeletePost(ID int) error {
	err := s.db.QueryRow(context.Background(), `
		DELETE FROM posts
		WHERE id = $1`,
		ID).Scan(&ID)

	return err
}

type Author struct {
	ID   int
	Name string
}

func (s *Storage) AddAuthor(author Author) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO authors (name)
		VALUES ($1)`,
		author.Name)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (s *Storage) Authors() ([]Author, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id, name FROM authors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authors []Author
	for rows.Next() {
		var a Author
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	fmt.Println(authors)
	return authors, nil
}
