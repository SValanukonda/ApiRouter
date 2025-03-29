package apperrors

type AppError struct {
	ErrorCode string
	Error     error
}

func NewAppError(errcode string, msg error) *AppError {
	return &AppError{
		ErrorCode: errcode,
		Error:     msg,
	}
}

func GetErrorMsgFromErrorCode(errcode string) string {
	return ErroCodes[errcode]
}
