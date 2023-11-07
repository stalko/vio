package model

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

const (
	IPAddress = iota
	CountryCode
	Country
	City
	Latitude
	Longitude
	MysteryValue
)

// CSVModel represents a single row in the CSV file
type CSVModel struct {
	IPAddress    string   // max length: 45
	CountryCode  *string  // max length: 2; ALPHA-2 ISO 3166
	Country      *string  // max length: 100
	City         *string  // max length: 200
	Latitude     *float64 // max/min 90.0000000 to -90.0000000
	Longitude    *float64 // max/min 180.0000000 to -180.0000000
	MysteryValue *int64
}

var ErrRecordIsEmpty = errors.New("record is empty")
var ErrInvalidIPAddress = errors.New("invalid IP address")
var ErrInvalidCountryCodeLength = errors.New("invalid country code length")
var ErrInvalidCountryLength = errors.New("invalid country length")
var ErrInvalidCityLength = errors.New("invalid city length")
var ErrInvalidLatitude = errors.New("invalid latitude, should be in range -90:90")
var ErrInvalidLongitude = errors.New("invalid longitude, should be in range -180:180")

// RecordToModel - validating and converting record string array to CSVModel
func RecordToModel(record []string) (*CSVModel, error) {
	var model CSVModel

	if len(record) == 0 {
		return nil, ErrRecordIsEmpty
	}

	for i, r := range record {
		switch i {
		case IPAddress:
			if net.ParseIP(r) == nil {
				return nil, ErrInvalidIPAddress
			}
			model.IPAddress = r
		case CountryCode:
			countryCode := strings.TrimSpace(r)
			if countryCode != "" {
				if len(countryCode) != 2 {
					return nil, ErrInvalidCountryCodeLength
				}

				model.CountryCode = &countryCode
			}
		case Country:
			country := strings.TrimSpace(r)
			if country != "" {
				if len(r) > 100 {
					return nil, ErrInvalidCountryLength
				}

				model.Country = &country
			}
		case City:
			city := strings.TrimSpace(r)
			if city != "" {
				if len(r) > 200 {
					return nil, ErrInvalidCityLength
				}

				model.City = &city
			}
		case Latitude:
			lat := strings.TrimSpace(r)
			if lat == "" {
				continue
			}

			s, err := strconv.ParseFloat(lat, 64)
			if err != nil {
				return nil, err
			}

			if -90 > s || s > 90 {
				return nil, ErrInvalidLatitude
			}

			model.Latitude = &s
		case Longitude:
			lon := strings.TrimSpace(r)
			if lon == "" {
				continue
			}

			s, err := strconv.ParseFloat(lon, 64)
			if err != nil {
				return nil, err
			}

			if -180 > s || s > 180 {
				return nil, ErrInvalidLongitude
			}

			model.Longitude = &s
		case MysteryValue:
			mv := strings.TrimSpace(r)
			if mv == "" {
				continue
			}

			s, err := strconv.ParseInt(mv, 0, 64)
			if err != nil {
				return nil, err
			}

			model.MysteryValue = &s
		}
	}
	return &model, nil
}
