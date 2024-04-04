package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWheOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response := responseRecorder.Result()

	//проверка возврата код ответа - 200
	require.Equal(t, 200, response.StatusCode, "Unexpected status code")
	//тело ответа не пустое
	require.NotEmpty(t, response.Body, "Response body is empty")

	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response := responseRecorder.Result()

	body, _ := io.ReadAll(response.Body)
	list := strings.Split(string(body), ",")

	assert.Len(t, list, totalCount, "wrong count value")

	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=TestingCity", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "Unexepected status code")
	body, _ := io.ReadAll(response.Body)
	//проверяем что тело ответа содержитт сообщение об ошибке
	require.Contains(t, string(body), "wrong city value")

	// здесь нужно добавить необходимые проверки
}
