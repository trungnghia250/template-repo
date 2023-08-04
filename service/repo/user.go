package repo

import (
	"FirstAssignment/service/model"
	"context"
	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	Delete(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

func (u *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Model(&model.User{}).Create(user).Error
}

func (u *userRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	res := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (u *userRepo) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	result := u.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	result := u.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, user *model.User) error {
	result := u.db.WithContext(ctx).Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
