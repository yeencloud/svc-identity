package database

import (
	"context"

	service "github.com/yeencloud/lib-base"
	"github.com/yeencloud/svc-identity/internal/domain"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	ID string `gorm:"primary_key;unique;not null;default:null;<-:create"`

	Username string
}

func (r *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = tx.Statement.Context.Value("userID").(string)
	return nil
}

func (r *UserModel) AddUser(ctx context.Context, user domain.User) error {
	return service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&user).Error
	})
}

func userToDomain(model UserModel) domain.User {
	return domain.User{
		ID:       model.ID,
		Username: model.Username,
	}
}

func userToModel(user domain.User) UserModel {
	return UserModel{
		ID:       user.ID,
		Username: user.Username,
	}
}
