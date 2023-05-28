package controller

import (
	"testing"

	"example.com/test_axxonsoft/v2/dto"

	"github.com/stretchr/testify/assert"
)

func TestTaskValidator_Validate(t *testing.T) {
	validator := TaskValidator{}

	// Test case 1: Valid taskDto
	taskDto := &dto.TaskDTO{
		Method: "GET",
		Url:    "http://example.com",
	}
	errors := validator.validate(taskDto)
	assert.NotNil(t, errors)

	// Test case 2: Empty method field
	taskDto = &dto.TaskDTO{
		Method: "",
		Url:    "http://example.com",
	}
	errors = validator.validate(taskDto)
	assert.NotNil(t, errors)
	assert.Equal(t, true, errors.HasErrors)
	assert.Contains(t, errors.Errors, "Method field is required")

	// Test case 3: Invalid method value
	taskDto = &dto.TaskDTO{
		Method: "INVALID",
		Url:    "http://example.com",
	}
	errors = validator.validate(taskDto)
	assert.NotNil(t, errors)
	assert.Equal(t, true, errors.HasErrors)
	assert.Contains(t, errors.Errors, "method is not allowed")

	// Test case 4: Invalid URL
	taskDto = &dto.TaskDTO{
		Method: "GET",
		Url:    "example",
	}
	errors = validator.validate(taskDto)
	assert.NotNil(t, errors)
	assert.Equal(t, true, errors.HasErrors)
	assert.Contains(t, errors.Errors, "URL is not valid")
}
