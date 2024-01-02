package leads

type leadData struct {
	LeadId        int    `json:"lead_id"`
	ProducerId    int    `json:"producer_id"`
	ConsumerId    int    `json:"consumer_id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Price         int    `json:"price"`
	Status        int    `json:"status"`
	CreatedAt     uint   `json:"created_at"`
	ResponsibleId int    `json:"responsible_id"`
}

type stageResponse struct {
	Stages map[int]string `json:"data"`
	Err    *string        `json:"error"`
}

type leadResponse struct {
	Data *leadData `json:"data"`
	Err  *string   `json:"error"`
}

type leadListResponse struct {
	Data []leadData `json:"data"`
	Err  *string    `json:"error"`
}
