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
	assert.Equal(t, 200, response.StatusCode, "Unexpected status code")
	//тело ответа не пустое
	assert.NotEmpty(t, response.Body, "Response body is empty")

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response := responseRecorder.Result()
	require.Equal(t, 200, response.StatusCode, "Unexpected status code")

	body, _ := io.ReadAll(response.Body)
	list := strings.Split(string(body), ",")

	assert.Len(t, list, totalCount, "wrong count value")

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=TestingCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "Unexepected status code")
	body, _ := io.ReadAll(response.Body)
	//проверяем что тело ответа содержитт сообщение об ошибке
	require.Contains(t, string(body), "wrong city value")

}
