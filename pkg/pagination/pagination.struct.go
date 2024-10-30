package pagination

type Pagination struct {
	Search     string      `json:"search,omitempty" query:"search"`
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	Filter     int      `json:"filter,omitempty" query:"filter"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int32        `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

