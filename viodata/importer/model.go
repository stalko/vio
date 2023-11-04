package importer

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
	IPAddress    string   //45
	CountryCode  *string  //2
	Country      *string  //100
	City         *string  //200
	Latitude     *float64 //max/min 90.0000000 to -90.0000000
	Longitude    *float64 // max/min 180.0000000 to -180.0000000
	MysteryValue *int64
}

var ErrRecordIsEmpty = errors.New("record is empty")
var ErrRInvalidIPAddress = errors.New("invalid IP address")

func RecordToModel(record []string) (*CSVModel, error) {
	var model CSVModel

	if len(record) == 0 {
		return nil, ErrRecordIsEmpty
	}

	for i, r := range record {
		switch i {
		case IPAddress:
			if net.ParseIP(r) == nil {
				return nil, ErrRInvalidIPAddress
			}
			model.IPAddress = r
		case CountryCode:
			countryCode := r
			if strings.TrimSpace(countryCode) != "" {
				model.CountryCode = &countryCode
			}
		case Country:
			country := r
			if strings.TrimSpace(country) != "" {
				model.Country = &country
			}
		case City:
			city := r
			if strings.TrimSpace(city) != "" {
				model.City = &city
			}
		case Latitude:
			if strings.TrimSpace(r) == "" {
				continue
			}

			s, err := strconv.ParseFloat(r, 64)
			if err != nil {
				return nil, err
			}

			model.Latitude = &s
		case Longitude:
			if strings.TrimSpace(r) == "" {
				continue
			}

			s, err := strconv.ParseFloat(r, 64)
			if err != nil {
				return nil, err
			}

			model.Longitude = &s
		case MysteryValue:
			if strings.TrimSpace(r) == "" {
				continue
			}

			s, err := strconv.ParseInt(r, 0, 64)
			if err != nil {
				return nil, err
			}

			model.MysteryValue = &s
		}
	}
	return &model, nil
}
