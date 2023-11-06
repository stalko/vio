package viodata_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stalko/viodata"
	"github.com/stalko/viodata/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetIPLocationByIP(t *testing.T) {
	const (
		IPAddress = "200.106.141.15"
	)
	var countryCode = "SI"
	var country = "Nepal"
	var city = "DuBuquemouth"
	var latitudeFloat float64 = -84.87503094689836
	var longitudeFloat float64 = 7.206435933364332
	var mysteryValueInt64 int64 = 7823011346

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().GetIPLocationsByIPAddress(ctx, IPAddress).Return(&storage.IPLocation{
		ID:           "id",
		IPAddress:    IPAddress,
		CountryName:  &country,
		CountryCode:  &countryCode,
		City:         &city,
		Latitude:     &latitudeFloat,
		Longitude:    &longitudeFloat,
		MysteryValue: &mysteryValueInt64,
	}, nil)

	vd := viodata.NewVioData(mockStorage, zap.NewExample())

	res, err := vd.GetIPLocationByIP(ctx, IPAddress)
	assert.NoError(t, err)

	if assert.NotNil(t, res) {
		assert.Equal(t, res.IPAddress, IPAddress)

		if assert.NotNil(t, res.City) {
			assert.Equal(t, *res.City, city)
		}

		if assert.NotNil(t, res.CountryCode) {
			assert.Equal(t, *res.CountryCode, countryCode)
		}

		if assert.NotNil(t, res.Country) {
			assert.Equal(t, *res.Country, country)
		}

		if assert.NotNil(t, res.Latitude) {
			assert.Equal(t, *res.Latitude, latitudeFloat)
		}

		if assert.NotNil(t, res.Longitude) {
			assert.Equal(t, *res.Longitude, longitudeFloat)
		}

		if assert.NotNil(t, res.MysteryValue) {
			assert.Equal(t, *res.MysteryValue, mysteryValueInt64)
		}
	}
}
