package pagination

import (
	"errors"
	"math"
)

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	Data        any   `json:"data"`
}

func PaginationData(currentPage int, limit int, totalItems int64, data any) (*Pagination, error) {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	if totalPages == 0 {
		return nil, nil
	}
	if currentPage > totalPages {
		return nil, errors.New("Page not found")
	}

	return &Pagination{
		CurrentPage: currentPage,
		PerPage:     limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		Data:        data,
	}, nil
}
