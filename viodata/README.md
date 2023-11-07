## VIO - DATA

Library for importing, validating and accessing saved data. PostgreSQL Database used as a storage provider. But can be replaced by implementing `./storage/storage.go` interface. `importer` contains interface and CSV implementation for importing file data and saving it in a storage. Some data can be skipped from saving in the database due to incorrect value( [Read more about validation](#validation) )

# Database

The main database is PostgreSQL. 

All changes to the database might be described in the migration query. It will protect the database and keep history of all changes.  ( [You can find more about how to use it in Dependencies](#dependencies) ). Each migration must include `up` and `down` action in case if something goes wrong. Be mindful, all migrations are stored in the database changelog table. Migration will be automatically executed after running `docker compose` or it can be executed manually (at your own risk) by executing following command:
```
make migrate-up
```

To make communication with the database there is a library: SQLC ( [You can find more about it in dependencies list](#dependencies) ). This library allows to generate GoLang code from SQL query defined in the folder `./db/queries/`. This library requires to have access to the schema definition as well. That's why one of the dependencies is `migrations` directory. Don't forget to re-generate code after editing queries by executing following command:
```
make gen-sql
```

The file's contents are fake. That's one more reason why the Database is not normalized. But it can be improved in the future by splitting `city` to independent table and placing `country_code` together with `countries` entity.

## GIS - under development(NOT READY)

Current version of the database supports GIS extension. This feature can improve storage and search over geolocation data. The data type is 
http://postgis.net/workshops/postgis-intro/geography.html and can use the most common SRID for geographic coordinates - 4326, which corresponds with “longitude/latitude on the WGS84 spheroid”. You can see the definition here: https://epsg.io/4326.

Right now it's not fully implemented due to conflict with `sqlcopy` operation required for fast inserting list of entities, more: https://www.postgresql.org/docs/current/sql-copy.html

# Validation

The validation happens on converting a CSV record to Model in `./model`. Here is the list of rules:

- IPAddress - string, max length: 45, validated by `net.ParseIP()`. IPv4 = varchar(15) = xxx.xxx.xxx.xxx, IPv6 = varchar(45) = 0000:0000:0000:0000:0000:0000:0000:0000
- CountryCode - *string, max length: 2. Should be 2 symbols by standard ALPHA-2 ISO 3166
- Country - *string, max length: 100. The longest country name has 56 symbols `The United Kingdom of Great Britain and Northern Ireland`. Just in case of appearance of new countries max length is 100.
- City - *string, max length: 200. The longest city name is `Krung Thep Mahanakhon Amon Rattanakosin Mahinthara Yuthaya Mahadilok Phop Noppharat Ratchathani Burirom Udomratchaniwet Mahasathan Amon Piman Awatan Sathit Sakkathattiya Witsanukam Prasit` with 187 symbols. There is room for additional 13 symbols, just in case.
- Latitude - *float64, max/min 90.0000000 to -90.0000000
- Longitude - *float64, max/min 180.0000000 to -180.0000000


# How to test

The main packages like: `db`, `importer`, `viodata`, `model` and `typeconverter` are covered with tests. You can execute tests by running command:
```
make test
```

For ease of testing there is a `mockgen` library ( [You can find more about it in dependencies list](#dependencies) ). Don't forget to re-generate mocks by executing following command:
```
make mockgen
```

# Dependencies
- [SQLC](https://sqlc.dev) is used to generate Go code and structs based on SQL queries. To add/edit SQL queries, simply modify [one of queries file](db/queries/ip_locations.sql) and run `make gen`. This will generate SQL boilerplate code for you.
- [Migrate](https://github.com/golang-migrate/migrate) generates database migration files. Use `bash gen_migration.sh [migration_name]` to generate new migrations. Example: `bash gen_migration.sh city_table`
- [Go Mock] (github.com/golang/mock/mockgen) generates mocks for interfaces. Make sure you've installed go package `go install github.com/golang/mock/mockgen@v1.6.0` and exported PATH: `export PATH=$PATH:$(go env GOPATH)/bin`
