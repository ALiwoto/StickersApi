package database

func (r *SearchPackRequest) BuildWhereQuery() (string, []any) {
	return "first_name LIKE ?", []any{"%" + r.PackTitle + "%"}
}
