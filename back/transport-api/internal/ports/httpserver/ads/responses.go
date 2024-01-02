package ads

type adData struct {
	AdId          int    `json:"ad_id"`
	CompanyId     int    `json:"company_id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Price         int    `json:"price"`
	Industry      string `json:"industry"`
	CreatedAt     uint   `json:"created_at"`
	CreatedId     int    `json:"created_id"`
	ResponsibleId int    `json:"responsible_id"`
}

type adResponse struct {
	Data *adData `json:"data"`
	Err  *string `json:"error"`
}

type adListResponse struct {
	Data []adData `json:"data"`
	Err  *string  `json:"error"`
}
