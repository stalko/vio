## TODO:

- Create comments


// * Define a data format suitable for the data contained in the CSV file;
// * Sanitise the entries: the file comes from an unreliable source, this means that the entries can be duplicated, may miss some value, the value can not be in the correct format or completely bogus;
// * At the end of the import process, return some statistics about the time elapsed, as well as the number of entries accepted/discarded;
// * The library should be configurable by an external configuration (particularly with regards to the DB configuration);

# How to test

0. Setup ENV variables by creating `.env` file from `.env.example`.

## Tools
- [SQLC](https://sqlc.dev) is used to generate Go code and structs based on SQL queries. To add/edit SQL queries, simply modify [the queries file](internal/sqldb/queries.sql) and run `make gen`. This will generate SQL boilerplate code for you.
- [Migrate](https://github.com/golang-migrate/migrate) generates database migration files. Use `bash gen_migration.sh [migration_name]` to generate new migrations. Example: `bash gen_migration.sh secure_link_index`
- [Go Mock] (github.com/golang/mock/mockgen) generates mocks for interfaces.Make sure you've install go package `go install github.com/golang/mock/mockgen@v1.6.0` and exported PATH: `export PATH=$PATH:$(go env GOPATH)/bin`


## Database

-- IPv4 = varchar(15) = xxx.xxx.xxx.xxx, IPv6 = varchar(45) = 0000:0000:0000:0000:0000:0000:0000:0000

### What Is The Longest Country Name In The World?
Rank	Country Name	Character Count
1	The United Kingdom of Great Britain and Northern Ireland	56
2	Independent and Sovereign Republic of Kiribati	46
3	Democratic Republic of Sao Tome and Principe	45


The longest city name:

Krung Thep Mahanakhon Amon Rattanakosin Mahinthara Yuthaya Mahadilok Phop Noppharat Ratchathani Burirom Udomratchaniwet Mahasathan Amon Piman Awatan Sathit Sakkathattiya Witsanukam Prasit

### Geo data 

http://postgis.net/workshops/postgis-intro/geography.html

The most common SRID for geographic coordinates is 4326, which corresponds to “longitude/latitude on the WGS84 spheroid”. You can see the definition here: https://epsg.io/4326


https://www.postgresql.org/docs/current/sql-copy.html