package repository

import (
	"context"

	"github.com/fumiyanakamura/go-sample-api/core/domain/model"
)

type UserRepository interface {
	Create(context.Context, model.User) (*model.User, error)
}
