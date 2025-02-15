package handler

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {

	// arrange
	expectedStatusCode := 200
	expectedBody := "Healthy status :)"
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// act
	GetHealth(w, req)
	resp := w.Result()
	rawBody, _ := io.ReadAll(resp.Body)
	actualBody := string(rawBody)

	// assert
	assert.Equal(t, expectedStatusCode, resp.StatusCode, "The status code should be equal")
	assert.Equal(t, expectedBody, actualBody)

}
