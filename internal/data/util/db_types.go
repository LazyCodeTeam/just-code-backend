package util

import "github.com/jackc/pgx/v5/pgtype"

func UnwrapDbString(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

func ToDbString(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{String: *value, Valid: true}
}
