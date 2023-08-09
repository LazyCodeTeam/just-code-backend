package util

import "github.com/jackc/pgx/v5/pgtype"

func UnwrapDbString(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
