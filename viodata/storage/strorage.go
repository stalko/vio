package storage

import (
	"context"
	"errors"
)

type Storage interface {
	InsertIPLocation(ctx context.Context, IPAddress string, countryName *string, countryCode *string, city *string, lat *float64, lon *float64, mysteryValue *int64) error

	GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (*IPLocation, error)
	GetCountIPLocationsByIPAddress(ctx context.Context, ipAddress string) (int64, error)
}

var ErrIPLocationNotFound = errors.New("ip_location not found")

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
