package importer_test

import (
	"math/rand"
	"testing"

	"github.com/stalko/viodata/importer"
	"github.com/stretchr/testify/assert"
)

const (
	IPAddress    = "200.106.141.15"
	countryCode  = "SI"
	country      = "Nepal"
	city         = "DuBuquemouth"
	latitude     = "-84.87503094689836"
	longitude    = "7.206435933364332"
	mysteryValue = "7823011346"
)

func TestRecordToModel(t *testing.T) {
	const (
		latitudeFloat     float64 = -84.87503094689836
		longitudeFloat    float64 = 7.206435933364332
		mysteryValueInt64 int64   = 7823011346
	)

	record := []string{
		IPAddress,
		countryCode,
		country,
		city,
		latitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.NoError(t, err)
	if assert.NotNil(t, model) {
		assert.Equal(t, model.IPAddress, IPAddress)

		if assert.NotNil(t, model.City) {
			assert.Equal(t, *model.City, city)
		}

		if assert.NotNil(t, model.CountryCode) {
			assert.Equal(t, *model.CountryCode, countryCode)
		}

		if assert.NotNil(t, model.Country) {
			assert.Equal(t, *model.Country, country)
		}

		if assert.NotNil(t, model.Latitude) {
			assert.Equal(t, *model.Latitude, latitudeFloat)
		}

		if assert.NotNil(t, model.Longitude) {
			assert.Equal(t, *model.Longitude, longitudeFloat)
		}

		if assert.NotNil(t, model.MysteryValue) {
			assert.Equal(t, *model.MysteryValue, mysteryValueInt64)
		}
	}

}

func TestRecordToModelEmptyRecord(t *testing.T) {
	record := []string{}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrRecordIsEmpty)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidIPAddress(t *testing.T) {
	const invalidIPAddress = "invalid_IP_address"

	record := []string{
		invalidIPAddress,
		countryCode,
		country,
		city,
		latitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidIPAddress)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidCountryCodeLength(t *testing.T) {
	const invalidCountryCode = "NL+"

	record := []string{
		IPAddress,
		invalidCountryCode,
		country,
		city,
		latitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidCountryCodeLength)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidCountryLength(t *testing.T) {
	invalidCountry := randomString(101)

	record := []string{
		IPAddress,
		countryCode,
		invalidCountry,
		city,
		latitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidCountryLength)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidCityLength(t *testing.T) {
	invalidCity := randomString(201)

	record := []string{
		IPAddress,
		countryCode,
		country,
		invalidCity,
		latitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidCityLength)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidLatitude(t *testing.T) {
	invalidLatitude := "-90.1"

	record := []string{
		IPAddress,
		countryCode,
		country,
		city,
		invalidLatitude,
		longitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidLatitude)
	assert.Nil(t, model)
}

func TestRecordToModelInvalidLongitude(t *testing.T) {
	invalidLongitude := "180.1"

	record := []string{
		IPAddress,
		countryCode,
		country,
		city,
		latitude,
		invalidLongitude,
		mysteryValue,
	}

	model, err := importer.RecordToModel(record)
	assert.ErrorIs(t, err, importer.ErrInvalidLongitude)
	assert.Nil(t, model)
}

func randomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
