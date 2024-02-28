package ads

type addAdRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Industry uint64 `json:"industry"`
	Price    uint   `json:"price"`
}

type updateAdRequest struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Industry    uint64 `json:"industry"`
	Price       uint   `json:"price"`
	Responsible uint64 `json:"responsible"`
}
