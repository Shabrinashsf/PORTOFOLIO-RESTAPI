package dto

import "errors"

const (
	//failed
	MESSAGE_FAILED_REGISTER_SHSF = "failed create shsf"
	MESSAGE_FAILED_UPDATE_SHSF   = "failed update shsf"
	MESSAGE_FAILED_GET_SHSF      = "failed get user SHSFs subevent"

	//success
	MESSAGE_SUCCESS_REGISTER_SHSF = "success create shsf"
	MESSAGE_SUCCESS_UPDATE_SHSF   = "success update shsf"
	MESSAGE_SUCCESS_GET_SHSF      = "success get user SHSFs subevent"
)

var (
	ErrParseUUID                    = errors.New("failed parsing to uuid")
	ErrParseSEventID                = errors.New("failed parsing subevent_id")
	ErrGetSubEventByID              = errors.New("failed get subevent by id")
	ErrCreateSHSF                   = errors.New("failed create SHSF")
	ErrCreateTransaction            = errors.New("failed to create transaction")
	ErrCreateEventTransaction       = errors.New("failed to create event transaction")
	ErrGetSHSFByUserID              = errors.New("failed to get shsf by user id")
	ErrGetSHSFByUserIDAndSubeventID = errors.New("failed get shsf by user id and subevent id")
	ErrSHSFNotInRevisionState       = errors.New("this shsf is not in revision state")
	ErrUpdateSHSF                   = errors.New("failed update shsf")
)

type (
	SHSFCreateRequest struct {
		Name       string `json:"name" form:"name" binding:"required,max=100"`
		Email      string `json:"email" form:"email" binding:"required,max=100"`
		TelpNumber string `json:"phone_number" form:"phone_number" binding:"required,max=30"`
		SubeventID string `json:"sub_event_id" form:"sub_event_id" binding:"required,max=100"`
	}

	SHSFMe struct {
		ID          string `json:"id" form:"id" binding:"required,max=700"`
		Name        string `json:"name" form:"name" binding:"required,max=700"`
		Email       string `json:"email" form:"email" binding:"required,max=100"`
		TelpNumber  string `json:"phone_number" form:"phone_number" binding:"required,max=30"`
		Subevent    string `json:"sub_event" form:"sub_event" binding:"required,max=100"`
		SubeventID  string `json:"sub_event_id"`
		AdminStatus string `json:"admin_status"`
	}

	SHSFUpdateRequest struct {
		Name       string `json:"name" form:"name"`
		Email      string `json:"email" form:"email"`
		TelpNumber string `json:"phone_number" form:"phone_number"`
	}
)
