package mongo

import (
	"context"
	"log"
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
	clientOptions := options.Client().ApplyURI(connection)
	clientOptions.SetConnectTimeout(15 * time.Second)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return &Storage{c: client}, nil
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

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
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
