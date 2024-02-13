package companies

type updateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    uint   `json:"industry"`
	OwnerId     uint   `json:"owner_id"`
}
