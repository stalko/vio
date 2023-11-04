package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
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

type DBConfig struct {
	ConnMaxIdleTime   time.Duration
	MaxOpenConns      int
	BackoffRetryCount uint64
	BackoffDuration   time.Duration
}

func NewDB(ctx context.Context, dsn string, logger *zap.Logger, config DBConfig) (storage.Storage, error) {
	var db *sql.DB
	var err error

	err = backoff.Retry(func() error {
		if db, err = sql.Open("pgx", dsn); err != nil {
			return err
		}
		if err := db.PingContext(ctx); err != nil {
			return err
		}

		return nil
	}, backoff.WithMaxRetries(backoff.NewConstantBackOff(config.BackoffDuration), config.BackoffRetryCount))

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetMaxOpenConns(config.MaxOpenConns)

	querier, err := gen.Prepare(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL queries: %w", err)
	}

	return &dbImpl{
		logger:  logger,
		querier: querier,
	}, nil
}

// GetCountIPLocationsByIPAddress
func (db *dbImpl) GetCountIPLocationsByIPAddress(ctx context.Context, ipAddress string) (int64, error) {
	count, err := db.querier.GetCountIPLocationsByIPAddress(ctx, ipAddress)
	if err != nil {
		db.logger.Error("can't count ip_locations", zap.String("ip_address", ipAddress), zap.Error(err))
		return 0, fmt.Errorf("error counting amount of ip_locations: %w", err)
	}
	return count, nil
}

// GetIPLocationsByIPAddress
func (db *dbImpl) GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (*storage.IPLocation, error) {
	ipLocation, err := db.querier.GetIPLocationsByIPAddress(ctx, ipAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

	if ipLocation.Latitude != 0 {
		res.Latitude = &ipLocation.Latitude
	}

	if ipLocation.Longitude != 0 {
		res.Longitude = &ipLocation.Longitude
	}

	if ipLocation.CountryID.Valid {
		country, err := db.querier.GetCountryByID(ctx, ipLocation.CountryID.Int32)
		if err != nil {
			db.logger.Error("can't get country", zap.Int32("country_id", ipLocation.CountryID.Int32))
			return nil, fmt.Errorf("error getting country: %w", err)
		}

		res.CountryName = &country.Name
	}

	return res, nil
}

// InsertIPLocation
func (db *dbImpl) InsertIPLocation(ctx context.Context, IPAddress string, countryName *string, countryCode *string, city *string, lat *float64, lon *float64, mysteryValue *int64) error {
	var countryID sql.NullInt32

	if countryName != nil {
		country, err := db.querier.InsertCountry(ctx, *countryName)
		if err != nil {
			db.logger.Error("can't insert country", zap.Error(err), zap.String("country_name", *countryName))
			return fmt.Errorf("error inverting a country: %w", err)
		}
		countryID = sql.NullInt32{
			Int32: country.ID,
			Valid: true,
		}
	}

	err := db.querier.InsertIPLocation(ctx, gen.InsertIPLocationParams{
		ID:           cuid.New(),
		CountryID:    countryID,
		IpAddress:    IPAddress,
		CountryCode:  tc.NewNullString(countryCode),
		City:         tc.NewNullString(city),
		MysteryValue: tc.NewNullInt64(mysteryValue),
		Latitude:     tc.NewNullFloat64(lat),
		Longitude:    tc.NewNullFloat64(lon),
	})
	if err != nil {
		db.logger.Error("can't insert ip_location", zap.Error(err), zap.String("ip_address", IPAddress))
		return fmt.Errorf("error inverting an ip_location: %w", err)
	}

	return nil
}
