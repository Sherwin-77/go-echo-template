package constants

type SortDirection string

// NOTE: This is default keyword for PSQL/MySQL database. Modify it if you use other database
const (
	SortDirectionAscending  SortDirection = "ASC"
	SortDirectionDescending SortDirection = "DESC"
)
