BEGIN;

DROP INDEX IF EXISTS ip_locations_ip_address_idx;
DROP TABLE IF EXISTS "ip_locations";
DROP TABLE IF EXISTS "countries";

DROP EXTENSION IF EXISTS postgis;

COMMIT;