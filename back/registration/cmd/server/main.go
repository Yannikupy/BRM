package main

import "registration/internal/ports/httpserver"

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API регистрации
//	@host			localhost:8081
//	@BasePath		/registration

func main() {
	srv := httpserver.New("localhost:8081", nil)
	_ = srv.ListenAndServe()
}
