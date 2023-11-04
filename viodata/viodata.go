package viodata

import (
	"context"
	"errors"

	"github.com/stalko/viodata/storage"
	"go.uber.org/zap"
)

type VioData interface {
	GetIPLocationByIP(ctx context.Context, IP string) (*IPLocation, error)
}

var ErrIPLocationNotFound = errors.New("ip_location entity not found in the storage")

type IPLocation struct {
	IPAddress    string
	CountryCode  *string
	Country      *string
	City         *string
	Latitude     *float64
	Longitude    *float64
	MysteryValue *int64
}

type vioData struct {
	s      storage.Storage
	logger *zap.Logger
}

func NewVioData(s storage.Storage, logger *zap.Logger) VioData {
	return &vioData{
		s:      s,
		logger: logger,
	}
}

// GetIPLocationByIP implements VioData.
func (vd *vioData) GetIPLocationByIP(ctx context.Context, IP string) (*IPLocation, error) {
	ipLocation, err := vd.s.GetIPLocationsByIPAddress(ctx, IP)
	if err != nil {
		if errors.Is(err, storage.ErrIPLocationNotFound) {
			return nil, ErrIPLocationNotFound
		}
		vd.logger.Error("can't retrieve ip_location from storage", zap.Error(err), zap.String("ip", IP))
		return nil, err
	}

	return &IPLocation{
		IPAddress:    ipLocation.IPAddress,
		CountryCode:  ipLocation.CountryCode,
		Country:      ipLocation.CountryName,
		City:         ipLocation.City,
		Latitude:     ipLocation.Latitude,
		Longitude:    ipLocation.Longitude,
		MysteryValue: ipLocation.MysteryValue,
	}, nil
}
