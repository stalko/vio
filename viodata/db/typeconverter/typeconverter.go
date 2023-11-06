package typeconverter

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func NewNullString(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{
		String: *s,
		Valid:  true,
	}
}

func NewStringPointer(s pgtype.Text) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func NewNullInt64(v *int64) pgtype.Int8 {
	if v == nil {
		return pgtype.Int8{}
	}
	return pgtype.Int8{
		Int64: *v,
		Valid: true,
	}
}

func NewInt64Pointer(v pgtype.Int8) *int64 {
	if v.Valid {
		return &v.Int64
	}
	return nil
}

func NewNullFloat64(v *float64) pgtype.Float8 {
	if v == nil {
		return pgtype.Float8{}
	}
	return pgtype.Float8{
		Float64: *v,
		Valid:   true,
	}
}
