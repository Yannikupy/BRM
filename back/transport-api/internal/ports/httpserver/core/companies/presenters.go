package companies

type updateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    uint64 `json:"industry"`
	OwnerId     uint64 `json:"owner_id"`
}
