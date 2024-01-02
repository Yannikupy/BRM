package companies

type updateCompanyRequest struct {
	Title    string `json:"title"`
	Industry int    `json:"industry"`
	OwnerId  int    `json:"owner_id"`
}
