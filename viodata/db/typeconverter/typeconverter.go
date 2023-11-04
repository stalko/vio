package typeconverter

import (
	"database/sql"

	"github.com/cridenour/go-postgis"
)

func NewNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NewStringPointer(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func NewNullInt64(v *int64) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *v,
		Valid: true,
	}
}

func NewInt64Pointer(v sql.NullInt64) *int64 {
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

func NewNullFloat64(v *float64) sql.NullFloat64 {
	if v == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
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
