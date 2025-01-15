package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory_Validations(t *testing.T) {
	tests := []struct {
		name         string
		categoryName string
		expectedErrs bool
	}{
		{
			name:         "valid category name",
			categoryName: "Valid Category",
			expectedErrs: false,
		},
		{
			name:         "empty category name",
			categoryName: "",
			expectedErrs: true,
		},
		{
			name:         "short category name",
			categoryName: "ab",
			expectedErrs: true,
		},
		{
			name:         "long category name",
			categoryName: "This category name is definitely too long",
			expectedErrs: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c, errs := NewCategory(tc.categoryName)

			if !tc.expectedErrs {
				assert.Nil(t, errs)
				assert.NotNil(t, c)
			} else {
				assert.Nil(t, c)
				assert.NotNil(t, errs)
			}
		})
	}
}
