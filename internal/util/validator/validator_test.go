package validator_test

import (
	"testing"

	"github.com/kazmerdome/best-ever-golang-starter/internal/util/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	type test struct {
		Name string `json:"name" bson:"name" validate:"required"`
	}

	// When calling validator, it should should validate the data and pass
	//
	ts := new(test)
	ts.Name = "test"
	err := validator.Validate(ts)
	assert.NoError(t, err)

	// When calling validator with invalid data, it should should validate the data and fail
	//
	tf := new(test)
	err = validator.Validate(tf)
	assert.Equal(t, err.Error(), "Key: 'test.Name' Error:Field validation for 'Name' failed on the 'required' tag")
}
