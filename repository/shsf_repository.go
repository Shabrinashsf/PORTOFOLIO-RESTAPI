package repository

import (
	"context"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type (
	SHSFRepository interface {
		GetSubEventByID(ctx context.Context, tx *gorm.DB, subeventID string) (entity.Subevent, error)
		Register(ctx context.Context, tx *gorm.DB, SHSF entity.SHSF) (entity.SHSF, error)
		CreatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payment) (entity.Payment, error)
		CreateEventPayment(ctx context.Context, tx *gorm.DB, eventPayment entity.EventPayment) (entity.EventPayment, error)
		GetSHSFByUserID(ctx context.Context, tx *gorm.DB, userID any) ([]entity.SHSF, error)
		GetSHSFByUserIDAndSubeventID(ctx context.Context, tx *gorm.DB, userID string, subeventID string) (entity.SHSF, error)
		Update(ctx context.Context, tx *gorm.DB, shsf entity.SHSF) (entity.SHSF, error)
	}

	shsfRepository struct {
		db *gorm.DB
	}
)

func NewSHSFService(db *gorm.DB) SHSFRepository {
	return &shsfRepository{
		db: db,
	}
}

func (r *shsfRepository) GetSubEventByID(ctx context.Context, tx *gorm.DB, subeventID string) (entity.Subevent, error) {
	if tx == nil {
		tx = r.db
	}

	var subevent entity.Subevent
	if err := tx.WithContext(ctx).Where("id = ?", subeventID).Take(&subevent).Error; err != nil {
		return entity.Subevent{}, err
	}
	return subevent, nil
}

func (r *shsfRepository) Register(ctx context.Context, tx *gorm.DB, SHSF entity.SHSF) (entity.SHSF, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&SHSF).Error; err != nil {
		return entity.SHSF{}, dto.ErrGeneral
	}

	return SHSF, nil
}

func (r *shsfRepository) CreatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payment) (entity.Payment, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&payment).Error; err != nil {
		return entity.Payment{}, err
	}
	return payment, nil
}

func (r *shsfRepository) CreateEventPayment(ctx context.Context, tx *gorm.DB, eventPayment entity.EventPayment) (entity.EventPayment, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&eventPayment).Error; err != nil {
		return entity.EventPayment{}, err
	}
	return eventPayment, nil
}

func (r *shsfRepository) GetSHSFByUserID(ctx context.Context, tx *gorm.DB, userID any) ([]entity.SHSF, error) {
	var data []entity.SHSF

	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Find(&data).Error; err != nil {
		return []entity.SHSF{}, err
	}

	return data, nil
}

func (r *shsfRepository) GetSHSFByUserIDAndSubeventID(ctx context.Context, tx *gorm.DB, userID string, subeventID string) (entity.SHSF, error) {
	if tx == nil {
		tx = r.db
	}

	var shsf entity.SHSF
	if err := tx.WithContext(ctx).Where("user_id = ? AND subevent_id = ?", userID, subeventID).Find(&shsf).Error; err != nil {
		return entity.SHSF{}, err
	}

	return shsf, nil
}

func (r *shsfRepository) Update(ctx context.Context, tx *gorm.DB, shsf entity.SHSF) (entity.SHSF, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(shsf).Error; err != nil {
		return entity.SHSF{}, err
	}
	return shsf, nil
}
