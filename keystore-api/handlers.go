package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func putValueHandler(ctx echo.Context) error {
	key := ctx.Param("key")

	reqBody := struct {
		Value string `json:"value"`
	}{}

	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := putValue(key, reqBody.Value); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]string{
			key: reqBody.Value,
		},
	)
}

func getValueHandler(ctx echo.Context) error {
	key := ctx.Param("key")

	value, err := getValue(key)
	if err != nil {
		switch {
		case errors.Is(err, ErrNoSuchKey):
			return echo.NewHTTPError(http.StatusNotFound, "key not found")

		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]string{
			key: value,
		},
	)
}

func deleteValueHandler(ctx echo.Context) error {
	key := ctx.Param("key")

	if err := deleteValue(key); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "value deleted")
}
