package repository

import (
	"context"
	"errors"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		GetAllUser(ctx context.Context) ([]entity.User, error)
		VerifyEmail(code string) (entity.User, error)
		UpdateIsVerified(user entity.User) error
		GetUserByID(parsedID uuid.UUID) (entity.User, error)
		UpdateUser(tx *gorm.DB, user entity.User) (entity.User, error)
		DeleteUser(tx *gorm.DB, user entity.User) error
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

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	var user entity.User
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil // Email belum ada, tidak dianggap error
		}
		return entity.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) VerifyEmail(code string) (entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "verification_code  = ?", code)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) UpdateIsVerified(user entity.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) GetUserByID(parsedID uuid.UUID) (entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "id = ?", parsedID)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) UpdateUser(tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.Model(&user).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
