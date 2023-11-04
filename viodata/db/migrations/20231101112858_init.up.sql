BEGIN;

CREATE EXTENSION postgis;

CREATE TABLE IF NOT EXISTS "countries" (
  "id" VARCHAR(32) PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL -- The longest name has 56 characters, but just in case we will have max as 100
);

ALTER TABLE IF EXISTS public.countries
    ADD CONSTRAINT countries_name_uq UNIQUE (name);

CREATE TABLE IF NOT EXISTS "ip_locations" (
  "id" VARCHAR(32) PRIMARY KEY,
  "ip_address" VARCHAR(45) NOT NULL, -- IPv4 = varchar(15) = xxx.xxx.xxx.xxx, IPv6 = varchar(45) = 0000:0000:0000:0000:0000:0000:0000:0000
  "country_id" VARCHAR(32),
  "country_code" VARCHAR(2), -- example: NL
  "city" VARCHAR(200),
  "location" geometry(Point, 4326), -- SRID = 4326, more here - https://epsg.io/4326
  "latitude" FLOAT,
  "longitude" FLOAT,
  "mystery_value" BIGINT,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "is_deleted" BOOLEAN NOT NULL DEFAULT false
);

CREATE INDEX ip_locations_ip_address_idx ON public.ip_locations (ip_address);

ALTER TABLE IF EXISTS public.ip_locations
    ADD CONSTRAINT ip_locations_countries_fk FOREIGN KEY (country_id)
    REFERENCES public.countries (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

COMMIT;