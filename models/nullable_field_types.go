package models

import "database/sql"

type NullInt64 = sql.NullInt64

type NullInt32 = sql.NullInt32

func NullInt64Value(f NullInt64) int64 {
	if f.Valid {
		return f.Int64
	} else {
		return 0
	}
}
