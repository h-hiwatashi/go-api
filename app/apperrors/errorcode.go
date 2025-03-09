package apperrors

type ErrCode string

const (
	Unknown              ErrCode = "U000"
	ErrCodeInternalError ErrCode = "INTERNAL_ERROR"
	ErrCodeBadRequest    ErrCode = "BAD_REQUEST"
)
