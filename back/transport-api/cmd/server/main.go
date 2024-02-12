package main

import (
	"transport-api/internal/ports/httpserver"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API
//	@host			localhost:8080
//	@BasePath		/api/v1

func main() {
	srv := httpserver.New("localhost:8080", nil)
	_ = srv.ListenAndServe()
}
