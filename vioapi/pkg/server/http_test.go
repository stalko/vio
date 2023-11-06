package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stalko/vioapi/pkg/server"
	"github.com/stalko/viodata"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetIPInformation(t *testing.T) {
	const (
		IPAddress = "200.106.141.15"
	)
	var countryCode = "SI"
	var country = "Nepal"
	var city = "DuBuquemouth"
	var latitudeFloat float64 = -84.87503094689836
	var longitudeFloat float64 = 7.206435933364332
	var mysteryValueInt64 int64 = 7823011346

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	vdMock := viodata.NewMockVioData(ctrl)

	srv := server.NewHTTPServer("8080", zap.NewExample(), vdMock)

	router, ok := srv.(*server.HTTPServer)
	assert.Equal(t, ok, true)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/ip_location/%s", IPAddress), nil)
	assert.NoError(t, err)

	vdMock.EXPECT().GetIPLocationByIP(req.Context(), IPAddress).Return(&viodata.IPLocation{
		IPAddress:    IPAddress,
		CountryCode:  &countryCode,
		Country:      &country,
		City:         &city,
		Latitude:     &latitudeFloat,
		Longitude:    &longitudeFloat,
		MysteryValue: &mysteryValueInt64,
	}, nil)

	router.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := server.IPLocation{}

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, resp.IPAddress, IPAddress)

	if assert.NotNil(t, resp.CountryCode) {
		assert.Equal(t, *resp.CountryCode, countryCode)
	}

	if assert.NotNil(t, resp.Country) {
		assert.Equal(t, *resp.Country, country)
	}

	if assert.NotNil(t, resp.City) {
		assert.Equal(t, *resp.City, city)
	}

	if assert.NotNil(t, resp.Latitude) {
		assert.Equal(t, *resp.Latitude, latitudeFloat)
	}

	if assert.NotNil(t, resp.Longitude) {
		assert.Equal(t, *resp.Longitude, longitudeFloat)
	}

	if assert.NotNil(t, resp.MysteryValue) {
		assert.Equal(t, *resp.MysteryValue, mysteryValueInt64)
	}
}

func TestGetIPInformationNotFound(t *testing.T) {
	const (
		IPAddressNotFound = "200.106.141.15"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	vdMock := viodata.NewMockVioData(ctrl)

	srv := server.NewHTTPServer("8080", zap.NewExample(), vdMock)

	router, ok := srv.(*server.HTTPServer)
	assert.Equal(t, ok, true)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/ip_location/%s", IPAddressNotFound), nil)
	assert.NoError(t, err)

	vdMock.EXPECT().GetIPLocationByIP(req.Context(), IPAddressNotFound).Return(nil, viodata.ErrIPLocationNotFound)

	router.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	resp := server.ErrorMessage{}

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, resp.Status, http.StatusNotFound)
}

func TestGetIPInformationIncorrectIP(t *testing.T) {
	const (
		IPAddressNotFound = "incorrect_ip"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	vdMock := viodata.NewMockVioData(ctrl)

	srv := server.NewHTTPServer("8080", zap.NewExample(), vdMock)

	router, ok := srv.(*server.HTTPServer)
	assert.Equal(t, ok, true)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/ip_location/%s", IPAddressNotFound), nil)
	assert.NoError(t, err)

	router.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	resp := server.ErrorMessage{}

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, resp.Status, http.StatusBadRequest)
}
