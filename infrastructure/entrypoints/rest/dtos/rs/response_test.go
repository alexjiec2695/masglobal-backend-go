package rs_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/infrastructure/entrypoints/rest/dtos/rs"
	"testing"
)

func TestNewResponse(t *testing.T) {
	data := struct {
		name  string
		email string
	}{"test", "test@test.com"}
	response := rs.NewResponse(data)

	assert.Equal(t, data, response["data"])
}

func TestNewErrorResponse_WithDetail(t *testing.T) {
	detail := struct {
		name  string
		email string
	}{"The name is required", "The email is required"}
	response := rs.NewErrorResponse(400, "Validation error", detail)
	err := response["error"].(map[string]interface{})

	assert.Equal(t, 400, err["code"])
	assert.Equal(t, "Validation error", err["description"])
	assert.Equal(t, detail, err["detail"])
}

func TestNewErrorResponse_WithoutDetail(t *testing.T) {
	response := rs.NewErrorResponse(401, "Unauthorized", nil)
	err := response["error"].(map[string]interface{})

	assert.Equal(t, 401, err["code"])
	assert.Equal(t, "Unauthorized", err["description"])
	assert.Nil(t, err["detail"])
}
