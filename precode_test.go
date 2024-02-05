package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест 1: Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	target := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}

// Тест 2: Город, который передаётся в параметре city, не поддерживается
func TestMainHandlerCityNotOk(t *testing.T) {
	WrongCity := "wrongCity"
	req := httptest.NewRequest("GET", "/cafe?count=2&city="+WrongCity, nil)
	ResponseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(ResponseRecorder, req)

	actualStatusCode := ResponseRecorder.Code
	assert.Equal(t, 400, actualStatusCode) // 400 - bad request

	expectedBody := "wrong city value"
	assert.Equal(t, expectedBody, ResponseRecorder.Body.String())
}

// Тест 3: Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerRequestIsOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	actualStatusCode := responseRecorder.Code
	require.Equal(t, 200, actualStatusCode) // 200 - ok
	require.NotEmpty(t, responseRecorder.Body.String())
}
