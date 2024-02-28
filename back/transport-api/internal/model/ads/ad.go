package ads

type Ad struct {
	Id           uint64
	CompanyId    uint64
	Title        string
	Text         string
	Industry     uint64
	Price        uint
	CreationDate int64
	CreatedBy    uint64
	Responsible  uint64
	IsDeleted    bool
}

type UpdateAd struct {
	Title       string
	Text        string
	Industry    uint64
	Price       uint
	Responsible uint64
}

type ListParams struct {
	Search *AdSearcher
	Filter *AdFilter
	Sort   *AdSorter
	Limit  uint
	Offset uint
}

type AdSearcher struct {
	Pattern string
}

type AdFilter struct {
	ByCompany bool
	CompanyId uint64

	ByIndustry bool
	Industry   uint64
}

type AdSorter struct {
	ByPriceAsc  bool
	ByPriceDesc bool

	ByDateAsc  bool
	ByDateDesc bool
}
