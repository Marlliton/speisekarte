package apperr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErr_New(t *testing.T) {
	t.Run("should create a new AppErr with code and message", func(t *testing.T) {
		code := 404
		message := "not found"
		err := New(message).WithCode(code)

		assert.NotNil(t, err)
		assert.Equal(t, code, *err.Code)
		assert.Equal(t, message, err.Message)
		assert.Empty(t, err.Reasons)
	})
}

func TestAppErr_WithReason(t *testing.T) {
	t.Run("should add a reason to the AppErr", func(t *testing.T) {
		err := New("bad request").
			WithReason("invalid input", "name")

		assert.NotNil(t, err)
		assert.Equal(t, 1, len(err.Reasons))
		assert.Equal(t, "invalid input", err.Reasons[0].Description)
		assert.Equal(t, "name", err.Reasons[0].Field)
	})

	t.Run("should add multiple reasons to the AppErr", func(t *testing.T) {
		err := New("bad request").
			WithReason("invalid input", "name").
			WithReason("invalid format", "email")

		assert.NotNil(t, err)
		assert.Equal(t, 2, len(err.Reasons))
		assert.Equal(t, "invalid input", err.Reasons[0].Description)
		assert.Equal(t, "name", err.Reasons[0].Field)
		assert.Equal(t, "invalid format", err.Reasons[1].Description)
		assert.Equal(t, "email", err.Reasons[1].Field)
	})
}

func TestAppErr_Error(t *testing.T) {
	t.Run("should format error without reasons", func(t *testing.T) {
		err := New("not found").WithCode(404)
		expected := "code: 404, message: not found"

		assert.Equal(t, expected, err.Error())
	})

	t.Run("should format error with reasons", func(t *testing.T) {
		err := New("bad request").
			WithCode(400).
			WithReason("invalid input", "name").
			WithReason("invalid format", "email")

		expected := "code: 400, message: bad request, reason: {field: name, description: invalid input}, reason: {field: email, description: invalid format}"

		assert.Equal(t, expected, err.Error())
	})
}

func TestAppErr_Is(t *testing.T) {
	t.Run("should return true when errors are equal", func(t *testing.T) {
		err1 := New("not found")
		err2 := New("not found")

		assert.True(t, err1.Is(err2))
	})

	t.Run("should return false when errors have different codes", func(t *testing.T) {
		err1 := New("not found").WithCode(404)
		err2 := New("not found").WithCode(400)

		assert.False(t, err1.Is(err2))
	})

	t.Run("should return false when errors have different messages", func(t *testing.T) {
		err1 := New("not found")
		err2 := New("bad request")

		assert.False(t, err1.Is(err2))
	})

	t.Run("should return false when target is not an AppErr", func(t *testing.T) {
		err1 := New("not found")
		err2 := errors.New("not found")

		assert.False(t, err1.Is(err2))
	})
}
