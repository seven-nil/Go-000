// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"net/http"
)

// Injectors from wire.go:

func InitHttpHandler(initParam interface{}, ctx context.Context) http.Handler {
	return handler
}