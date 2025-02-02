package apperr

import "fmt"

type AppErr struct {
	Code    int
	Message string
	Reasons []Reason
}

type Reason struct {
	Description string
	Field       string
}

func New(code int, message string) *AppErr {
	return &AppErr{
		Code:    code,
		Message: message,
		Reasons: []Reason{},
	}
}

func (err *AppErr) WithReason(description, field string) *AppErr {
	err.Reasons = append(err.Reasons, Reason{Description: description, Field: field})

	return err
}

func (err *AppErr) Error() string {
	if len(err.Reasons) == 0 {
		return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
	}

	reasons := ""
	for _, reason := range err.Reasons {
		reasons += fmt.Sprintf(", reason: {field: %s, description: %s}", reason.Field, reason.Description)
	}

	return fmt.Sprintf("code: %d, message: %s%s", err.Code, err.Message, reasons)
}

func (err *AppErr) Is(target error) bool {
	if target, ok := target.(*AppErr); ok {
		return err.Code == target.Code && err.Message == target.Message
	}
	return false
}
