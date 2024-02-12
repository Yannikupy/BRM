package httpserver

type tokensData struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type tokensResponse struct {
	Data *tokensData `json:"data"`
	Err  *string     `json:"error"`
}
