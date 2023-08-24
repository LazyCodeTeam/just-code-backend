package util

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func FromPgString(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

func FromPgInt(value pgtype.Int4) *int {
	if !value.Valid {
		return nil
	}
	val := int(value.Int32)

	return &val
}

func ToPgString(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{String: *value, Valid: true}
}

func ToPgInt4(value *int) pgtype.Int4 {
	if value == nil {
		return pgtype.Int4{Valid: false}
	}

	return pgtype.Int4{Int32: int32(*value), Valid: true}
}

func DecodeUUID(uuid pgtype.UUID) string {
	src := uuid.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}

func EncodeUUID(uuid string) pgtype.UUID {
	uuid = uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]
	var out [16]byte
	copy(out[:], uuid)

	return pgtype.UUID{Bytes: out, Valid: true}
}
