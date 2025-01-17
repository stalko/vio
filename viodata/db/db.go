package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lucsky/cuid"
	"github.com/stalko/viodata/db/gen"
	tc "github.com/stalko/viodata/db/typeconverter"
	"github.com/stalko/viodata/storage"
	"go.uber.org/zap"
)

type dbImpl struct {
	logger  *zap.Logger
	querier gen.Querier
}

func NewDB(ctx context.Context, dsn string, logger *zap.Logger) (storage.Storage, error) {
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	querier := gen.New(dbpool)

	return &dbImpl{
		logger:  logger,
		querier: querier,
	}, nil
}

func NewDBFromQuerier(ctx context.Context, querier gen.Querier, logger *zap.Logger) storage.Storage {
	return &dbImpl{
		logger:  logger,
		querier: querier,
	}
}

// GetIPLocationsByIPAddress with argument IP address, returns information about the IP address' location (e.g. country, city)
// return an error storage.ErrIPLocationNotFound in case when select return's 0 entities from the database
func (db *dbImpl) GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (*storage.IPLocation, error) {
	ipLocation, err := db.querier.GetIPLocationsByIPAddress(ctx, ipAddress)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Info("ip_location entity now found in the database")
			return nil, storage.ErrIPLocationNotFound
		}
		db.logger.Error("can't get ip_location", zap.String("ip_address", ipAddress), zap.Error(err))
		return nil, fmt.Errorf("error getting ip_location: %w", err)
	}

	res := &storage.IPLocation{
		ID:           ipLocation.ID,
		IPAddress:    ipLocation.IpAddress,
		City:         tc.NewStringPointer(ipLocation.City),
		CountryCode:  tc.NewStringPointer(ipLocation.CountryCode),
		MysteryValue: tc.NewInt64Pointer(ipLocation.MysteryValue),
	}

	if ipLocation.Latitude.Valid {
		res.Latitude = &ipLocation.Latitude.Float64
	}

	if ipLocation.Longitude.Valid {
		res.Longitude = &ipLocation.Longitude.Float64
	}

	if ipLocation.CountryID.Valid {
		country, err := db.querier.GetCountryByID(ctx, ipLocation.CountryID.String)
		if err != nil {
			db.logger.Error("can't get country", zap.String("country_id", ipLocation.CountryID.String))
			return nil, fmt.Errorf("error getting country: %w", err)
		}

		res.CountryName = &country.Name
	}

	return res, nil
}

// BulkInsertIPLocation inserting list of IPLocations as well as countries for them to the database. Each country will be inserted in for loop. IPLocations will be inserted in one go with `copyfrom` feature.
func (db *dbImpl) BulkInsertIPLocation(ctx context.Context, IPLocations []storage.InsertIPLocation) error {
	insertLocations := []gen.BulkInsertIPLocationsParams{}

	for _, ipLoc := range IPLocations {
		var countryID pgtype.Text

		if ipLoc.CountryName != nil {

			//FIXME: Investigate why first 4-8 countries can't be inserted without retry mechanism.
			//First look shows that database is lagging on inserting first N entities, after a second - come back to normal
			err := backoff.Retry(func() error {
				country, err := db.querier.InsertCountry(ctx, gen.InsertCountryParams{
					ID:   cuid.New(),
					Name: *ipLoc.CountryName,
				})
				if err != nil {
					db.logger.Error("can't insert country", zap.Error(err), zap.String("country_name", *ipLoc.CountryName))
					return fmt.Errorf("error inverting a country: %w", err)
				}

				countryID = pgtype.Text{
					String: country.ID,
					Valid:  true,
				}

				return nil
			}, backoff.WithMaxRetries(backoff.NewConstantBackOff(3*time.Second), 3))
			if err != nil {
				db.logger.Error("can't insert country after retry", zap.Error(err), zap.String("country_name", *ipLoc.CountryName))
				return fmt.Errorf("error inverting a country: %w", err)
			}
		}

		insertLocations = append(insertLocations,
			gen.BulkInsertIPLocationsParams{
				ID:           cuid.New(),
				CountryID:    countryID,
				IpAddress:    ipLoc.IPAddress,
				CountryCode:  tc.NewNullString(ipLoc.CountryCode),
				City:         tc.NewNullString(ipLoc.City),
				MysteryValue: tc.NewNullInt64(ipLoc.MysteryValue),
				Latitude:     tc.NewNullFloat64(ipLoc.Latitude),
				Longitude:    tc.NewNullFloat64(ipLoc.Longitude),
			},
		)
	}

	_, err := db.querier.BulkInsertIPLocations(ctx, insertLocations)
	if err != nil {
		db.logger.Error("can't insert ip_locations", zap.Error(err), zap.Int("count", len(insertLocations)))
		return fmt.Errorf("error inverting ip_locations: %w", err)
	}

	return nil
}
