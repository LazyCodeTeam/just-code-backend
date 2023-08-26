package util

import (
	"encoding/hex"
	"fmt"
	"log/slog"

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

func FromPgUUID(uuid pgtype.UUID) string {
	src := uuid.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}

func ToPgUUID(uuid string) pgtype.UUID {
	uuid = uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]

	out, err := hex.DecodeString(uuid)
	if err != nil {
		slog.Error("Error decoding UUID", "err", err)
		return pgtype.UUID{Valid: false}
	}
	var buf [16]byte
	copy(buf[:], out)

	return pgtype.UUID{Bytes: buf, Valid: true}
}
