package mongo

import (
	"context"
	"stugi/gonews/pkg/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	databaseName   = "gonews"
	collectionName = "posts"
)

type Storage struct {
	c *mongo.Client
}

func New(connection string) (*Storage, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(connection))

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	return &Storage{
		c: client,
	}, nil
}

func (s *Storage) AddPost(post storage.Post) error {
	collection := s.c.Database(databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.c.Database(databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []storage.Post
	for cursor.Next(context.Background()) {
		var post storage.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Storage) DeletePost(ID int) error {
	collection := s.c.Database(databaseName).Collection(collectionName)

	_, err := collection.DeleteOne(context.Background(), bson.M{"id": ID})
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdatePost(post storage.Post) error {
	collection := s.c.Database(databaseName).Collection(collectionName)

	_, err := collection.UpdateOne(context.Background(), bson.M{"id": post.ID}, bson.M{"$set": post})
	if err != nil {
		return err
	}
	return nil
}
