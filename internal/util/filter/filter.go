package filter

import (
	"time"

	"github.com/google/uuid"
)

// StringFilter
//

type StringFilter struct {
	Eq    *string `json:"eq"`
	Regex *string `json:"regex"`
}

// UuidFilter
//

type UuidFilter struct {
	Eq *uuid.UUID  `json:"eq"`
	In []uuid.UUID `json:"in"`
	Ne *uuid.UUID  `json:"ne"`
}

// IntFilter
//

type IntFilter struct {
	Eq  *int `json:"eq"`
	Gt  *int `json:"gt"`
	Gte *int `json:"gte"`
	Lt  *int `json:"lt"`
	Lte *int `json:"lte"`
}

// Float64Filter
//

type Float64Filter struct {
	Eq  *float64 `json:"eq"`
	Gt  *float64 `json:"gt"`
	Gte *float64 `json:"gte"`
	Lt  *float64 `json:"lt"`
	Lte *float64 `json:"lte"`
}

// DateFilter
//

type DateFilter struct {
	Eq  *time.Time `json:"eq"`
	Gt  *time.Time `json:"gt"`
	Gte *time.Time `json:"gte"`
	Lt  *time.Time `json:"lt"`
	Lte *time.Time `json:"lte"`
}

// SortFilter
//

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type SortFilter struct {
	SortBy    string     `json:"sortBy"`
	SortOrder *SortOrder `json:"sortOrder"`
}

// PaginationFilter
//

type PaginationFilter struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}
