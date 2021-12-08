package util

import "database/sql"

func ToNullString(eventStr *string) sql.NullString {
	var nullString sql.NullString
	if len(*eventStr) == 0 {
		nullString.Valid = false
	} else {
		nullString.String = *eventStr
		nullString.Valid = true
	}
	return nullString
}

func ToNullInt64(eventInt64 *int64) sql.NullInt64 {
	var nullInt64 sql.NullInt64
	nullInt64.Valid = (eventInt64 != nil)
	if !nullInt64.Valid {
		return nullInt64
	}

	nullInt64.Int64 = *eventInt64
	return nullInt64
}
