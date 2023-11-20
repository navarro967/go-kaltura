package kaltura

const (
	LT             = "lt"
	LTE            = "lte"
	GT             = "gt"
	GTE            = "gte"
	EQ             = "eq"
	LIKE           = "like"
	XLIKE          = "xlike"
	LIKEX          = "likex"
	IN             = "in"
	NOT_IN         = "notin"
	NOT            = "not"
	BIT_AND        = "bitand"
	BIT_OR         = "bitor"
	MULTI_LIKE_OR  = "mlikeor"
	MULTI_LIKE_AND = "mlikeand"
	MATCH_OR       = "matchor"
	MATCH_AND      = "matchand"
	NOT_CONTAINS   = "notcontains"
)

const (
	CREATED_AT_ASC  = "+createdAt"
	CREATED_AT_DESC = "-createdAt"
	UPDATED_AT_ASC  = "+updatedAt"
	UPDATED_AT_DESC = "-updatedAt"
	ID_ASC          = "+id"
	ID_DESC         = "-id"
)

type KalturaFilter struct {
	//AdvancedSearch map[string]string  `json:"advancedSearch,omitempty"`
	Pager  KalturaFilterPager `json:"pager,omitempty"`
	Filter KalturaBaseFilter  `json:"filter"`
}

type KalturaBaseFilter struct {
	OrderBy string `json:"orderBy,omitempty"`
}

type KalturaFilterPager struct {
	PageSize  int `json:"pageSize,omitempty"`
	PageIndex int `json:"pageIndex,omitempty"`
}
