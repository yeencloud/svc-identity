package database

import (
	"context"

	"github.com/samber/lo"
	service "github.com/yeencloud/lib-base"
	"github.com/yeencloud/svc-identity/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct{}

type UserModel struct {
	gorm.Model

	ID       string `gorm:"primary_key;unique;not null;default:null;<-:create"`
	Email    string `gorm:"index;unique;not null;default:null"`
	Username string `gorm:"index;unique;not null;default:null"`
	Password string
}

func (r UserRepo) AddUser(ctx context.Context, user domain.User, password string) error {
	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(lo.ToPtr(userToModel(user))).Error
	})

	if err != nil {
		return err
	}

	return service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&UserModel{}).Where("id = ?", user.ID).Update("password", password).Error
	})
}

func (r UserRepo) GetAuthByUsername(ctx context.Context, username string) (*domain.AuthInformation, error) {
	var m UserModel

	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).First(&m, "username = ?", username).First(&m).Error
	})
	if err != nil {
		return nil, err
	}

	return lo.ToPtr(authToDomain(m)), nil
}

func (r UserRepo) GetAuthByEmail(ctx context.Context, email string) (*domain.AuthInformation, error) {
	var m UserModel

	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).First(&m, "email = ?", email).First(&m).Error
	})
	if err != nil {
		return nil, err
	}

	return lo.ToPtr(authToDomain(m)), nil
}

func (r UserRepo) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var m UserModel
	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Where("id = ?", id).First(&m).Error
	})

	if err != nil {
		return nil, err
	}
	return lo.ToPtr(userToDomain(m)), nil
}

func userToDomain(model UserModel) domain.User { //nolint:unused
	return domain.User{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		Username:  model.Username,
		Email:     model.Email,
	}
}

func authToDomain(model UserModel) domain.AuthInformation {
	return domain.AuthInformation{
		ID:             model.ID,
		HashedPassword: model.Password,
	}
}

func userToModel(user domain.User) UserModel {
	return UserModel{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}
