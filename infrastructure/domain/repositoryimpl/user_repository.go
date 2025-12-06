package repositoryimpl

import (
	"context"

	"github.com/fumiyanakamura/go-sample-api/core/domain/model"
	"github.com/fumiyanakamura/go-sample-api/core/domain/repository"
	gormHandler "github.com/fumiyanakamura/go-sample-api/infrastructure/db/gorm"
	"gorm.io/gorm"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	handler *gorm.DB
}

func NewUserRepository() *UserRepository {
	h, _ := gormHandler.GetInstance()
	return &UserRepository{
		handler: h,
	}
}

// TODO: 実装とテスト追加
func (r *UserRepository) Create(context.Context, model.User) (*model.User, error) {
	return nil, nil
}
