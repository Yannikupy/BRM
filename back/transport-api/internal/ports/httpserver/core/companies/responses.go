package companies

import "transport-api/internal/model/core"

func errorResponse(err error) companyResponse {
	if err == nil {
		return companyResponse{}
	}
	errStr := err.Error()
	return companyResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func companyToCompanyData(company core.Company) companyData {
	return companyData{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}
}

func ownerToOwnerData(owner core.Employee) ownerData {
	return ownerData{
		Id:           owner.Id,
		CompanyId:    owner.CompanyId,
		FirstName:    owner.FirstName,
		SecondName:   owner.SecondName,
		Email:        owner.Email,
		JobTitle:     owner.JobTitle,
		Department:   owner.Department,
		CreationDate: owner.CreationDate,
		IsDeleted:    owner.IsDeleted,
	}
}

type companyAndOwnerResponse struct {
	Data *companyAndOwnerData `json:"data"`
	Err  *string              `json:"err"`
}

type companyAndOwnerData struct {
	Company companyData `json:"company"`
	Owner   ownerData   `json:"owner"`
}

type ownerData struct {
	Id           uint   `json:"id"`
	CompanyId    uint   `json:"company_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	CreationDate int64  `json:"creation_date"`
	IsDeleted    bool   `json:"is_deleted"`
}

type companyResponse struct {
	Data *companyData `json:"data"`
	Err  *string      `json:"error"`
}

type companyData struct {
	Id           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Industry     uint    `json:"industry"`
	OwnerId      uint    `json:"owner_id"`
	Rating       float64 `json:"rating"`
	CreationDate int64   `json:"creation_date"`
	IsDeleted    bool    `json:"is_deleted"`
}

type industryResponse struct {
	Industries []string `json:"data"`
	Err        *string  `json:"error"`
}

type mainPageResponse struct {
	Data *mainPageData `json:"data"`
	Err  *string       `json:"error"`
}

type mainPageData struct {
	Title string            `json:"title"`
	Stats mainPageStatsData `json:"stats"`
}

type mainPageStatsData struct {
	ActiveLeadsAmount     int     `json:"active_leads_amount"`
	ActiveLeadsPrice      int     `json:"active_leads_price"`
	ClosedLeadsAmount     int     `json:"total_leads_amount"`
	ClosedLeadsPrice      int     `json:"closed_leads_price"`
	ActiveAdsAmount       int     `json:"active_ads_aount"`
	CompanyAbsoluteRating float64 `json:"company_absolute_rating"`
	CompanyRelativeRating float64 `json:"company_relative_rating"`
}
