package companies

type companyData struct {
	Title     string `json:"title"`
	Industry  string `json:"industry"`
	OwnerId   int    `json:"owner_id"`
	CreatedAt uint   `json:"created_at"`
}

type companyResponse struct {
	Data *companyData `json:"data"`
	Err  *string      `json:"error"`
}

type industryResponse struct {
	Industries map[int]string `json:"data"`
	Err        *string        `json:"error"`
}

type mainPageStatsData struct {
	ActiveLeadsAmount     int     `json:"active_leads_amount"`
	ActiveLeadsPrice      int     `json:"active_leads_price"`
	ClosedLeadsAmount     int     `json:"total_leads_amount"`
	ClosedLeadsPrice      int     `json:"closed_leads_price"`
	ActiveAdsAount        int     `json:"active_ads_aount"`
	CompanyAbsoluteRating float64 `json:"company_absolute_rating"`
	CompanyRelativeRating float64 `json:"company_relative_rating"`
}

type mainPageData struct {
	Title string            `json:"title"`
	Stats mainPageStatsData `json:"stats"`
}

type mainPageResponse struct {
	Data *mainPageData `json:"data"`
	Err  *string       `json:"error"`
}
