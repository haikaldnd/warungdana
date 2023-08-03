package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"warungdana/models"
	"warungdana/presenter"
	"warungdana/usecase"
	"warungdana/utils/constant"
)

type H map[string]interface{}

func Number2(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := usecase.Number2(db)

		switch res.Status.Message {
		case constant.FailedGetData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number3(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := usecase.Number3(db)

		switch res.Status.Message {
		case constant.FailedGetData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number4(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := usecase.Number4(db)

		switch res.Status.Message {
		case constant.FailedGetData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number5(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := usecase.Number5(db)

		switch res.Status.Message {
		case constant.FailedGetData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number6(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			request models.File
		)
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		res := usecase.Number6(db, &request)

		switch res.Status.Message {
		case constant.FailedWriteFile:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 500
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number7(db *sql.DB) echo.HandlerFunc {
	var res presenter.DefaultResponse
	return func(c echo.Context) error {
		request := models.FileOpen{FileName: c.QueryParam("file_name")}
		if request.FileName == "" {
			res.Status.Message = constant.RequestInvalid
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 500
			return c.JSON(res.HTTPCode, res)
		}
		res := usecase.Number7(db, &request)

		switch res.Status.Message {
		case constant.FailedOpenFile:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 500
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number8(db *sql.DB) echo.HandlerFunc {
	var res presenter.DefaultResponse
	return func(c echo.Context) error {
		request := models.InputCity{City: c.QueryParam("city")}
		if request.City == "" {
			res.Status.Message = constant.RequestInvalid
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 500
			return c.JSON(res.HTTPCode, res)
		}
		res := usecase.Number8(db, &request)

		switch res.Status.Message {
		case constant.DataNotFound:
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number9(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := usecase.Number9(db)

		switch res.Status.Message {
		case constant.DataNotFound:
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number9c(db *sql.DB) echo.HandlerFunc {
	var res presenter.DefaultResponse
	return func(c echo.Context) error {
		request := models.InputMulti{Data: c.QueryParam("input")}
		if request.Data == "" {
			res.Status.Message = constant.RequestInvalid
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 500
			return c.JSON(res.HTTPCode, res)
		}
		res := usecase.Number9c(db, &request)

		switch res.Status.Message {
		case constant.DataNotFound:
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 400
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number9d(db *sql.DB) echo.HandlerFunc {
	var res presenter.DefaultResponse
	return func(c echo.Context) error {
		request := models.InputMulti{Data: c.QueryParam("input")}
		if request.Data == "" {
			res.Status.Message = constant.RequestInvalid
			res.HTTPCode = http.StatusBadRequest
			res.Status.Code = 500
			return c.JSON(res.HTTPCode, res)
		}
		res := usecase.Number9d(db, &request)

		switch res.Status.Message {
		case constant.FailedConvertData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 500
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}

func Number10a(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		res := usecase.Number10a(db)

		switch res.Status.Message {
		case constant.FailedGetData:
			res.HTTPCode = http.StatusInternalServerError
			res.Status.Code = 500
		default:
			res.HTTPCode = http.StatusOK
			res.Status.Code = 200
		}
		return c.JSON(res.HTTPCode, res)
	}
}
