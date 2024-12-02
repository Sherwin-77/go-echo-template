package constants

type FilterType string

const (
	FilterResponseSingleOption   string = "singleOption"
	FilterResponseMultipleOption string = "multipleOption"
	FilterResponsePartialText    string = "partialText"

	FilterTypeExact   FilterType = "exact"
	FilterTypePartial FilterType = "partial"
	FilterTypeCustom  FilterType = "custom"
)
