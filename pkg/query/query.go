package query

import (
	"github.com/sherwin-77/go-echo-template/pkg/constants"
	"gorm.io/gorm"
	"net/url"
	"strconv"
	"strings"
)

type FilterParam struct {
	DisplayName  string
	FieldName    string
	InternalName string
	FilterType   constants.FilterType
	Callback     func(db *gorm.DB, value string) *gorm.DB
}

type SortParam struct {
	DisplayName string
	Field       string
	Direction   constants.SortDirection
}

type Builder struct {
	model          interface{}
	AllowedFilters []FilterParam
	AllowedSorts   []SortParam
	DefaultSort    SortParam
}

// ExtractFilters parses query params and matches them to allowed filters
func ExtractFilters(queryParams url.Values, allowedFilters []FilterParam) map[string]string {
	filters := make(map[string]string)
	for _, filter := range allowedFilters {
		if value := queryParams.Get(filter.InternalName); value != "" {
			filters[filter.InternalName] = value
		}
	}
	return filters
}

// ExtractSorting parses query params for sorting (field and direction)
func ExtractSorting(queryParams url.Values, allowedSorts []SortParam, defaultSort SortParam) (string, constants.SortDirection) {
	sortField := queryParams.Get("sort")
	sortDirection := constants.SortDirectionAscending
	if strings.HasPrefix(sortField, "-") {
		sortDirection = constants.SortDirectionDescending
		sortField = strings.TrimPrefix(sortField, "-")
	}

	isValidSort := false
	for _, sort := range allowedSorts {
		if sort.Field == sortField {
			isValidSort = true
			break
		}
	}

	if !isValidSort {
		sortField = defaultSort.Field
		sortDirection = defaultSort.Direction
	}

	return sortField, sortDirection
}

// ExtractPagination parses query params for pagination (limit and page)
func ExtractPagination(queryParams url.Values) (int, int) {
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil || limit < 1 {
		limit = constants.DefaultPerPage
	}
	limit = min(limit, constants.MaxPerPage)

	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	return limit, page
}

func (b *Builder) ApplyBuilder(db *gorm.DB, queryParams url.Values) *gorm.DB {
	filters := ExtractFilters(queryParams, b.AllowedFilters)

	for _, filter := range b.AllowedFilters {
		if value, ok := filters[filter.InternalName]; ok {
			switch filter.FilterType {
			case constants.FilterTypeExact:
				db = db.Where(filter.FieldName+" = ?", value)
			case constants.FilterTypePartial:
				db = db.Where(filter.FieldName+" ILIKE ?", "%"+value+"%")
			case constants.FilterTypeCustom:
				if filter.Callback != nil {
					db = filter.Callback(db, value)
				}
			}
		}
	}

	sortField, sortDirection := ExtractSorting(queryParams, b.AllowedSorts, b.DefaultSort)
	db = db.Order(sortField + " " + string(sortDirection))

	limit, page := ExtractPagination(queryParams)
	db = db.Offset((page - 1) * limit).Limit(limit)

	return db
}
