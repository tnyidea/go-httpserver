package config

import (
	"context"
	"github.com/tnyidea/go-sample-userdata/models"
)

const UserServiceContextDatabase = "userserviceContextDatabase"

func NewContext() (context.Context, error) {
	ctx := context.Background()

	db, err := models.NewUserDatabase("../data/us-500.json")
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, UserServiceContextDatabase, db)

	return ctx, nil
}
