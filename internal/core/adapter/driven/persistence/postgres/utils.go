package postgres

import "database/sql"

func stringToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
