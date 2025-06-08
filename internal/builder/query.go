package builder

import (
	"github.com/sherwin-77/go-echo-template/pkg/constants"
	"github.com/sherwin-77/go-echo-template/pkg/query"
)

func NewUserQueryBuilder() query.Builder {
	return query.NewBuilder(
		nil,
		[]query.FilterParam{
			{DisplayName: "Email", FieldName: "email", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{DisplayName: "Name", FieldName: "name", InternalName: "username", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
		},
		[]query.SortParam{
			{DisplayName: "Email", FieldName: "email"},
			{DisplayName: "Name", FieldName: "username"},
			{DisplayName: "Created At", FieldName: "created_at"},
		},
		query.SortParam{DisplayName: "Name", FieldName: "username", Direction: query.SortDirectionAscending},
	)
}
