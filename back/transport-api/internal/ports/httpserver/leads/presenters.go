package leads

type addLeadRequest struct {
	ProducerId    int    `json:"producer_id"`
	ConsumerId    int    `json:"consumer_id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Price         int    `json:"price"`
	ResponsibleId int    `json:"responsible_id"`
}

type updateLeadRequest struct {
	Title         string `json:"title"`
	Text          string `json:"text"`
	Price         int    `json:"price"`
	ResponsibleId int    `json:"responsible_id"`
	Status        int    `json:"status"`
}
