package apperr

import "fmt"

type AppErr struct {
	Code    *int
	Message string
	Reasons []Reason
}

type Reason struct {
	Description string
	Field       string
}

func New(message string) *AppErr {
	return &AppErr{
		Message: message,
		Reasons: []Reason{},
	}
}

func (err *AppErr) WithCode(code int) *AppErr {
	err.Code = &code

	return err
}

func (err *AppErr) WithReason(description, field string) *AppErr {
	err.Reasons = append(err.Reasons, Reason{Description: description, Field: field})

	return err
}

func (err *AppErr) Error() string {
	codeStr := ""
	if err.Code != nil {
		codeStr = fmt.Sprintf("code: %d", *err.Code)
	}

	if len(err.Reasons) == 0 {
		if codeStr != "" {

			return fmt.Sprintf("%s, message: %s", codeStr, err.Message)
		}

		return fmt.Sprintf("message: %s", err.Message)
	}

	reasons := ""
	for _, reason := range err.Reasons {
		reasons += fmt.Sprintf(", reason: {field: %s, description: %s}", reason.Field, reason.Description)
	}

	if codeStr != "" {

		return fmt.Sprintf("%s, message: %s%s", codeStr, err.Message, reasons)
	}

	return fmt.Sprintf("message: %s%s", err.Message, reasons)
}

func (err *AppErr) Is(target error) bool {
	if target, ok := target.(*AppErr); ok {
		if err.Code != nil {
			return err.Code == target.Code && err.Message == target.Message
		}
		return err.Message == target.Message
	}
	return false
}
