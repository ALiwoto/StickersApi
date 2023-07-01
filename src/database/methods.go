package database

import "strings"

func (r *SearchPackRequest) BuildWhereQuery() (string, []any) {
	return "LOWER(pack_title) LIKE ?", []any{"%" + strings.ToLower(r.PackTitle) + "%"}
}
