package storage

import (
	"context"
	"errors"
)

type Storage interface {
	InsertIPLocation(ctx context.Context, IPLocation InsertIPLocation) error
	BulkInsertIPLocation(ctx context.Context, IPLocations []InsertIPLocation) error

	GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (*IPLocation, error)
	GetCountIPLocationsByIPAddress(ctx context.Context, ipAddress string) (int64, error)
}

var ErrIPLocationNotFound = errors.New("ip_location not found")

type InsertIPLocation struct {
	IPAddress    string
	CountryName  *string
	CountryCode  *string
	City         *string
	Latitude     *float64
	Longitude    *float64
	MysteryValue *int64
}
type IPLocation struct {
	ID           string
	IPAddress    string
	CountryName  *string
	CountryCode  *string
	City         *string
	Latitude     *float64
	Longitude    *float64
	MysteryValue *int64
}
