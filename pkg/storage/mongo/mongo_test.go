package mongo

import (
	"stugi/gonews/pkg/storage"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestStorage_AddPost(t *testing.T) {
	type fields struct {
		c *mongo.Client
	}
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "AddPost",
			fields: fields{
				c: nil,
			},
			args: args{
				post: storage.Post{
					ID:          1,
					AuthorID:    1,
					AuthorName:  "test",
					Title:       "test",
					Content:     "test",
					CreatedAt:   1,
					PublishedAt: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				c: tt.fields.c,
			}
			if err := s.AddPost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
