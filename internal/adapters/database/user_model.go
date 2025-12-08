package database

import (
	"context"
	"errors"

	service "github.com/yeencloud/lib-base"
	"github.com/yeencloud/svc-identity/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct{}

type UserModel struct {
	gorm.Model

	ID string `gorm:"primary_key;unique;not null;default:null;<-:create"`

	Username string
}

func (r *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	val := tx.Statement.Context.Value("userID")
	id, ok := val.(string)
	if !ok {
		return errors.New("userID not found in context or not a string")
	}

	r.ID = id
	return nil
}

func (r UserRepo) AddUser(ctx context.Context, user domain.User) error {
	return service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&user).Error
	})
}

func userToDomain(model UserModel) domain.User { //nolint:unused
	return domain.User{
		ID:       model.ID,
		Username: model.Username,
	}
}

func userToModel(user domain.User) UserModel { //nolint:unused
	return UserModel{
		ID:       user.ID,
		Username: user.Username,
	}
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}
