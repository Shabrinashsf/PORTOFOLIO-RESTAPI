package repository

import (
	"context"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (models.User, bool, error)
		RegisterUser(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error)
		GetAllUser(ctx context.Context) ([]models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (models.User, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var user models.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return models.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
