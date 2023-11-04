package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/stalko/vioapi/docs"
	"github.com/stalko/viodata"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	srv     *http.Server
	logger  *zap.Logger
	vioData viodata.VioData
}

// @title VIO API
// @version 1.0
// @description This a service to access VIO data

// @contact.name Artur Odnostalko
// @contact.url https://linkedin.com/in/a.odnostalko
// @contact.email stalko23@gmail.com

// @host localhost:8080
// @BasePath /
func NewHTTPServer(port string, logger *zap.Logger, vioData viodata.VioData) Server {
	s := &HTTPServer{
		logger:  logger,
		vioData: vioData,
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //render swagger documentation from the folder: ./docs
	router.GET("/ip_location/:ip", s.GetIPInformation)

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	return s
}

type ErrorMessage struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

type IPLocation struct {
	IPAddress    string   `json:"ip_address"`
	CountryCode  *string  `json:"country_code,omitempty"`
	Country      *string  `json:"country,omitempty"`
	City         *string  `json:"city,omitempty"`
	Latitude     *float64 `json:"latitude,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty"`
	MysteryValue *int64   `json:"mystery_value,omitempty"`
}

// GetIPInformation godoc
// @Summary Get information about the IP address' location (e.g. country, city)
// @Description
// @Tags    ip_information
// @Accept  json
// @Produce json
// @Success 200 {object} IPLocation
// @Failure 400 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage "Location for given IP - not found"
// @Param ip path string true "IP address" example(220.235.222.173)
// @Security Token, OAuth2Password
// @Router /ip_location/{ip} [get]
func (s *HTTPServer) GetIPInformation(c *gin.Context) {
	ip := c.Param("ip")

	if net.ParseIP(ip) == nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{
			Error:  "Invalid IP address",
			Status: http.StatusBadRequest,
		})
		return
	}

	model, err := s.vioData.GetIPLocationByIP(c.Request.Context(), ip)
	if err != nil {
		if errors.Is(err, viodata.ErrIPLocationNotFound) {
			c.JSON(http.StatusNotFound,
				ErrorMessage{
					Error:  "IP location not found",
					Status: http.StatusNotFound,
				})
			return
		}

		c.JSON(http.StatusBadRequest,
			ErrorMessage{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})
		return
	}

	resp := IPLocation{
		IPAddress:    model.IPAddress,
		CountryCode:  model.CountryCode,
		Country:      model.Country,
		City:         model.City,
		Latitude:     model.Latitude,
		Longitude:    model.Longitude,
		MysteryValue: model.MysteryValue,
	}

	c.JSON(http.StatusOK, resp)
}

func (s *HTTPServer) Run() error {
	s.logger.Info(fmt.Sprintf("running HTTP server on %s", s.srv.Addr))
	return s.srv.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) {
	s.logger.Info(fmt.Sprintln("shutting down HTTP server"))
	s.srv.Shutdown(ctx)
}
