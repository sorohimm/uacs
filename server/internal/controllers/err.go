package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func errStatusCode(err error) int {
	switch err {
	case mongo.ErrNoDocuments:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
