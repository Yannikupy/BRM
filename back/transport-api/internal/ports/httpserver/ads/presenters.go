package ads

type addAdRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Price    int    `json:"price"`
	Industry string `json:"industry"`
}

type updateAdRequest struct {
	Text  string `json:"text"`
	Price int    `json:"price"`
}
