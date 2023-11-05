package typeconverter

import (
	"database/sql"

	"github.com/cridenour/go-postgis"
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

func NewPoint(lat, lon *float64) *postgis.Point {
	if lat != nil && lon != nil {
		return &postgis.Point{
			X: *lat,
			Y: *lon,
		}
	}
	return nil
}

func NewLatLon(p *postgis.Point) (*float64, *float64) {
	if p != nil {
		return &p.X, &p.Y
	}
	return nil, nil
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

func NewFloat64Pointer(v sql.NullFloat64) *float64 {
	if v.Valid {
		return &v.Float64
	}
	return nil
}
