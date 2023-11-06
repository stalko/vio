package db_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stalko/viodata/db"
	"github.com/stalko/viodata/db/gen"
	"github.com/stalko/viodata/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInsertIPLocation(t *testing.T) {
	countryName := "Ukraine"
	var countInsertedRows int64 = 1
	ipLocation := storage.InsertIPLocation{
		IPAddress:   "",
		CountryName: &countryName,
	}
	country := gen.InsertCountryRow{
		ID:   "country_id",
		Name: countryName,
	}

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := gen.NewMockQuerier(ctrl)
	mockQuerier.EXPECT().InsertCountry(ctx, gomock.Any()).Return(country, nil)
	mockQuerier.EXPECT().BulkInsertIPLocations(ctx, gomock.Any()).Return(countInsertedRows, nil)

	database := db.NewDBFromQuerier(ctx, mockQuerier, zap.NewExample())

	err := database.BulkInsertIPLocation(ctx, []storage.InsertIPLocation{
		ipLocation,
	})
	assert.NoError(t, err)
}

func TestGetIPLocationsByIPAddress(t *testing.T) {
	country := gen.Country{
		ID:   "country_id",
		Name: "Ukraine",
	}

	IPLocation := gen.GetIPLocationsByIPAddressRow{
		ID: "ip_address_id",
		CountryID: pgtype.Text{
			String: country.ID,
			Valid:  true,
		},
		IpAddress: "ip_address",
		CountryCode: pgtype.Text{
			String: "country_code",
			Valid:  true,
		},
		City: pgtype.Text{
			String: "city",
			Valid:  true,
		},
		Latitude: pgtype.Float8{
			Float64: 1,
			Valid:   true,
		},
		Longitude: pgtype.Float8{
			Float64: 2,
			Valid:   true,
		},
		MysteryValue: pgtype.Int8{
			Int64: 3,
			Valid: true,
		},
	}

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := gen.NewMockQuerier(ctrl)
	mockQuerier.EXPECT().GetCountryByID(ctx, country.ID).Return(country, nil)
	mockQuerier.EXPECT().GetIPLocationsByIPAddress(ctx, IPLocation.IpAddress).Return(IPLocation, nil)

	database := db.NewDBFromQuerier(ctx, mockQuerier, zap.NewExample())

	res, err := database.GetIPLocationsByIPAddress(ctx, IPLocation.IpAddress)
	assert.NoError(t, err)

	if assert.NotNil(t, res) {
		assert.Equal(t, res.ID, IPLocation.ID)

		assert.Equal(t, res.IPAddress, IPLocation.IpAddress)

		if assert.NotNil(t, res.CountryName) {
			assert.Equal(t, *res.CountryName, country.Name)
		}

		if assert.NotNil(t, res.CountryCode) {
			assert.Equal(t, *res.CountryCode, IPLocation.CountryCode.String)
		}

		if assert.NotNil(t, res.Latitude) {
			assert.Equal(t, *res.Latitude, IPLocation.Latitude.Float64)
		}

		if assert.NotNil(t, res.Longitude) {
			assert.Equal(t, *res.Longitude, IPLocation.Longitude.Float64)
		}

		if assert.NotNil(t, res.MysteryValue) {
			assert.Equal(t, *res.MysteryValue, IPLocation.MysteryValue.Int64)
		}
	}
}
