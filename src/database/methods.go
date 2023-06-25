package database

func (r *SearchPackRequest) BuildWhereQuery() (string, []any) {
	return "pack_title LIKE ?", []any{"%" + r.PackTitle + "%"}
}
