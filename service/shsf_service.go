package service

import (
	"context"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/repository"
	"github.com/google/uuid"
)

type (
	SHSFService interface {
		Register(ctx context.Context, req dto.SHSFCreateRequest, userID string) (dto.SHSFMe, error)
		GetMe(ctx context.Context, userID any) ([]dto.SHSFMe, error)
		Update(ctx context.Context, req dto.SHSFUpdateRequest, userID string, subeventID string) (dto.SHSFMe, error)
	}

	shsfService struct {
		shsfRepo repository.SHSFRepository
	}
)

func NewSHSFService(shsfRepo repository.SHSFRepository) SHSFService {
	return &shsfService{
		shsfRepo: shsfRepo,
	}
}

func (s *shsfService) Register(ctx context.Context, req dto.SHSFCreateRequest, userID string) (dto.SHSFMe, error) {
	mu.Lock()
	defer mu.Unlock()

	UserID, err := uuid.Parse(userID)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrParseUUID
	}

	Sub_eventID, err := uuid.Parse(req.SubeventID)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrParseSEventID
	}

	subevent, err := s.shsfRepo.GetSubEventByID(ctx, nil, req.SubeventID)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrGetSubEventByID
	}

	shsf := entity.SHSF{
		UserID:       UserID,
		SubeventID:   Sub_eventID,
		SubeventName: subevent.Name,
		Name:         req.Name,
		Email:        req.Email,
		NoTelp:       req.TelpNumber,
		AdminStatus:  entity.ADMIN_STATUS_PENDING,
	}

	shsfReg, err := s.shsfRepo.Register(ctx, nil, shsf)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrCreateSHSF
	}

	transID := uuid.New()

	trans := entity.Payment{
		ID:            transID,
		UserID:        UserID,
		PaymentStatus: "PENDING",
		InvoiceURL:    "sementara kosong",
		PaidAmount:    0,
	}

	_, err = s.shsfRepo.CreatePayment(ctx, nil, trans)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrCreateTransaction
	}

	eventTrans := entity.EventPayment{
		RegistID:  shsfReg.ID,
		PaymentID: transID,
	}

	_, err = s.shsfRepo.CreateEventPayment(ctx, nil, eventTrans)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrCreateEventTransaction
	}

	return dto.SHSFMe{
		ID:          shsf.ID.String(),
		Name:        shsf.Name,
		Email:       shsf.Email,
		TelpNumber:  shsf.NoTelp,
		Subevent:    shsf.SubeventName,
		SubeventID:  shsf.SubeventID.String(),
		AdminStatus: shsf.AdminStatus,
	}, nil
}

func (s *shsfService) GetMe(ctx context.Context, userID any) ([]dto.SHSFMe, error) {
	shsf, err := s.shsfRepo.GetSHSFByUserID(ctx, nil, userID)
	if err != nil {
		return []dto.SHSFMe{}, dto.ErrGetSHSFByUserID
	}

	var returnData []dto.SHSFMe
	for _, elemen := range shsf {
		returnData = append(returnData, dto.SHSFMe{
			ID:          elemen.ID.String(),
			Name:        elemen.Name,
			Email:       elemen.Email,
			TelpNumber:  elemen.NoTelp,
			Subevent:    elemen.SubeventName,
			SubeventID:  elemen.SubeventID.String(),
			AdminStatus: elemen.AdminStatus,
		})
	}

	return returnData, nil
}

func (s *shsfService) Update(ctx context.Context, req dto.SHSFUpdateRequest, userID string, subeventID string) (dto.SHSFMe, error) {
	shsf, err := s.shsfRepo.GetSHSFByUserIDAndSubeventID(ctx, nil, userID, subeventID)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrGetSHSFByUserIDAndSubeventID
	}
	if shsf.AdminStatus != entity.ADMIN_STATUS_REVISION {
		return dto.SHSFMe{}, dto.ErrSHSFNotInRevisionState
	}

	shsf.Name = req.Name
	shsf.Email = req.Email
	shsf.NoTelp = req.TelpNumber
	shsf.AdminStatus = entity.ADMIN_STATUS_REVISED

	shsfUpd, err := s.shsfRepo.Update(ctx, nil, shsf)
	if err != nil {
		return dto.SHSFMe{}, dto.ErrUpdateSHSF
	}

	return dto.SHSFMe{
		ID:          shsfUpd.ID.String(),
		Name:        shsfUpd.Name,
		Email:       shsfUpd.Email,
		TelpNumber:  shsfUpd.NoTelp,
		Subevent:    shsfUpd.SubeventName,
		SubeventID:  shsfUpd.SubeventID.String(),
		AdminStatus: shsfUpd.AdminStatus,
	}, nil
}
