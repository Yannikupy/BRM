package httpserver

type refreshRequest struct {
	Refresh string `json:"refresh"`
}

type authorizeRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
