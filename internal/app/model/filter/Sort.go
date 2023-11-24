package filter

import (
	"net/http"
	"slices"
	"urbathon-backend-2023/pkg/errorHandler"
)

type Sort struct {
	Field *string `json:"field" example:"name"`
	Order *string `json:"order" example:"asc"`
}

type Sortable interface {
	GetSortableFields() []string
	GetSort() *Sort
}

func ValidateSort(s *Sort, sortableField []string) (*Sort, *errorHandler.HttpErr) {
	if s.Field != nil && !slices.Contains(sortableField, *s.Field) {
		return nil, errorHandler.New("cant sort with this field", http.StatusBadRequest)
	}
	if s.Order != nil && *s.Order != "asc" && *s.Order != "desc" {
		return nil, errorHandler.New("sort order must be asc or desc", http.StatusBadRequest)
	}
	return s, nil
}
