// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: copyfrom.go

package gen

import (
	"context"
)

// iteratorForBulkInsertIPLocations implements pgx.CopyFromSource.
type iteratorForBulkInsertIPLocations struct {
	rows                 []BulkInsertIPLocationsParams
	skippedFirstNextCall bool
}

func (r *iteratorForBulkInsertIPLocations) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForBulkInsertIPLocations) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].IpAddress,
		r.rows[0].CountryID,
		r.rows[0].CountryCode,
		r.rows[0].City,
		r.rows[0].Latitude,
		r.rows[0].Longitude,
		r.rows[0].MysteryValue,
	}, nil
}

func (r iteratorForBulkInsertIPLocations) Err() error {
	return nil
}

func (q *Queries) BulkInsertIPLocations(ctx context.Context, arg []BulkInsertIPLocationsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"ip_locations"}, []string{"id", "ip_address", "country_id", "country_code", "city", "latitude", "longitude", "mystery_value"}, &iteratorForBulkInsertIPLocations{rows: arg})
}
