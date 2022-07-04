package controllers

import (
	"net/http"
	"uacs/internal/services"
)

func errStatusCode(err error) int {
	switch err {
	case services.ErrIdNotFount:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
