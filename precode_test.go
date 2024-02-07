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

func doRequest(target string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", target, nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

// Тест 1: Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	requestedCount := 100
	target := fmt.Sprintf("/cafe?count=%d&city=moscow", requestedCount)
	responseRecorder := doRequest(target)
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())

}

// Тест 2: Город, который передаётся в параметре city, не поддерживается
func TestMainHandlerCityNotOk(t *testing.T) {
	wrongCity := "wrongCity"
	target := fmt.Sprintf("/cafe?count=2&city=%s", wrongCity)
	responseRecorder := doRequest(target)

	actualStatusCode := responseRecorder.Code
	assert.Equal(t, 400, actualStatusCode) // 400 - bad request

	expectedBody := "wrong city value"
	assert.Equal(t, expectedBody, responseRecorder.Body.String())
}

// Тест 3: Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerRequestIsOk(t *testing.T) {
	target := "/cafe?count=2&city=moscow"
	responseRecorder := doRequest(target)
	actualStatusCode := responseRecorder.Code
	require.Equal(t, 200, actualStatusCode) // 200 - ok
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 2)
	assert.Equal(t, "Мир кофе,Сладкоежка", responseRecorder.Body.String())
}

// Задачка на потом: ДОП тест 4 - Count не был передан (16-17 строка) - копирую тестовые функции и меняю таргеты.
// Задачка на потом: ДОП тест 5 - в count передано не число (23-24 строка) - копирую тестовые функции и меняю таргеты.
