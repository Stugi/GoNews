package main

import (
	"fmt"
	"log"
	"net/http"
	"stugi/gonews/pkg/api"
	"stugi/gonews/pkg/storage"
	"stugi/gonews/pkg/storage/memdb"
	"stugi/gonews/pkg/storage/mongo"
	"stugi/gonews/pkg/storage/postgres"
)

const (
	port        = ":8080"
	postgresUrl = "postgres://postgres:postgres@server.domain/posts"
	mongoUrl    = "mongodb://localhost:27017/"
)

// Server структура сервера.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var srv server

	db1 := memdb.New()

	db2, err := postgres.New(postgresUrl)

	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	db3, err := mongo.New(mongoUrl)

	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	_, _, _ = db1, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db3

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	err = http.ListenAndServe(port, srv.api.Router())

	if err != nil {
		log.Fatal(err)
	}
}
