package storage

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(int) error   // удаление публикации по ID
}

type Post struct {
	ID          int
	AuthorID    int
	AuthorName  string
	Title       string
	Content     string
	CreatedAt   int64
	PublishedAt int64
}
