package rmq

type jobRequest struct {
	JobType string `json:"job_type"`
	JobBody []byte `json:"job_body"`
}

type jobResponse struct {
	JobType   string `json:"job_type"`
	JobResult []byte `json:"job_result"`
	Error     string `json:"error"`
}

// Все возможные значения JobType
const (
	unknown               = ""
	createCompanyAndOwner = "create_company_and_owner"
	updateCompany         = "update_company"
	deleteCompany         = "delete_company"
	createEmployee        = "create_employee"
	updateEmployee        = "update_employee"
	deleteEmployee        = "delete_employee"
	createContact         = "create_contact"
	updateContact         = "update_contact"
	deleteContact         = "delete_contact"
)
