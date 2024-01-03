package main

import "auth/internal/ports/httpserver"

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API авторизации
//	@host			localhost:8082
//	@BasePath		/auth/v1

func main() {
	srv := httpserver.New("localhost:8082", nil)
	_ = srv.ListenAndServe()
}
