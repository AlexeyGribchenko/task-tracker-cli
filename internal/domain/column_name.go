package domain

import "errors"

type ColumnName string

const (
	columnName        = "name"
	columnState       = "status"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDescription = "description"
)

var (
	ErrInvalidColumnName = errors.New("Invalid column name")
)

func ParseColumnName(name string) (ColumnName, error) {
	switch name {
	case columnName, columnState, columnCreatedAt, columnUpdatedAt, columnDescription:
		return ColumnName(name), nil
	}
	return ColumnName(""), ErrInvalidColumnName
}
