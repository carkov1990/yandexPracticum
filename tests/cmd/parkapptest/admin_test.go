package main

import (
	"bytes"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/Yandex-Practicum/go-autotests/internal/random"
)

// TestUserAuth checks registration and authentication
func (suite *ParkappSuite) TestUserAuth() {
	login := random.ASCIIString(4, 10)

	suite.Run("add_user", func() {
		m := []byte(`
			{
				"login": "` + login + `"
			}
		`)

		req := resty.New().
			SetBaseURL(suite.parkappServerAddress).
			R().
			SetHeader("Content-Type", "application/json").
			SetBody(m)

		resp, err := req.Post("/api/user/add")

		noRespErr := suite.Assert().NoErrorf(err, "Ошибка при попытке сделать запрос на добавление юзера")
		validStatus := suite.Assert().Equalf(http.StatusOK, resp.StatusCode(),
			"Несоответствие статус кода ответа ожидаемому в хендлере '%s %s'", req.Method, req.URL,
		)

		if !noRespErr || !validStatus {
			dump := dumpRequest(suite.T(), req.RawRequest, bytes.NewReader(m))
			suite.T().Logf("Оригинальный запрос:\n\n%s", dump)
		}
	})

	suite.Run("remove_user", func() {
		m := []byte(`
			{
				"login": "` + login + `"
			}
		`)

		req := resty.New().
			SetBaseURL(suite.parkappServerAddress).
			R().
			SetHeader("Content-Type", "application/json").
			SetBody(m)

		resp, err := req.Post("/api/user/remove")

		noRespErr := suite.Assert().NoErrorf(err, "Ошибка при попытке сделать запрос на добавление юзера")
		validStatus := suite.Assert().Equalf(http.StatusOK, resp.StatusCode(),
			"Несоответствие статус кода ответа ожидаемому в хендлере '%s %s'", req.Method, req.URL,
		)

		if !noRespErr || !validStatus {
			dump := dumpRequest(suite.T(), req.RawRequest, bytes.NewReader(m))
			suite.T().Logf("Оригинальный запрос:\n\n%s", dump)
		}
	})
}
