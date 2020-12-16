//+build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func InitHttpHandler(initParam interface{}, ctx context.Context) http.Handler {
	wire.Build(initParam.Impl, initStruct, handler.InitHandler)
	return &mux.Router{}
}
