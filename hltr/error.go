package hltr

import (
	"errors"
	"fmt"
)

type HLTRError interface {
	Error() string
	Fields() map[string]string
	SetMessage(message string)
	Set(field, err string)
	HasErrors() bool
	ToHttpError() map[string]any
}
type hltrError struct {
	message     string
	fieldErrors map[string]any
}

func NewError() *hltrError {
	return &hltrError{
		fieldErrors: make(map[string]any),
	}
}
func (fe *hltrError) Set(field, err string) {
	fe.fieldErrors[field] = err
}

func (fe *hltrError) Error() string {
	if len(fe.fieldErrors) == 0 {
		return ""
	}
	err := errors.New("validation error")

	return err.Error() + ": " + fmt.Sprintf("%v", fe.fieldErrors)
}

func (fe *hltrError) Fields() map[string]any {
	return fe.fieldErrors
}

func (fe *hltrError) HasErrors() bool {
	return len(fe.fieldErrors) > 0
}

func (fe *hltrError) SetMessage(message string) {
	fe.message = message
}

func (fe *hltrError) ToHttpError() map[string]any {
	return map[string]any{
		"message": fe.message,
		"fields":  fe.fieldErrors,
	}
}
