package domain

import (
	"errors"
	"strings"
)

type ColumnTitle string

const (
	ColumnId          = "id"
	ColumnName        = "name"
	ColumnStatus      = "status"
	ColumnCreatedAt   = "created_at"
	ColumnUpdatedAt   = "updated_at"
	ColumnDescription = "description"
)

var (
	ErrInvalidColumnName = errors.New("Invalid column name")
)

// ParseColumnName parses column name
// If name valid returns ColumnTitle
// If column name invalid, returns ErrInvalidColumnName
func ParseColumnName(name string) (ColumnTitle, error) {

	name = strings.ToLower(name)

	switch name {
	case ColumnId, ColumnName, ColumnStatus, ColumnCreatedAt, ColumnUpdatedAt, ColumnDescription:
		return ColumnTitle(name), nil
	}
	return ColumnTitle(""), ErrInvalidColumnName
}
