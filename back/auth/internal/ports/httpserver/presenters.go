package httpserver

type refreshRequest struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type logoutRequest struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
