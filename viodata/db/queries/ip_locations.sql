-- name: GetIPLocationsByIPAddress :one
SELECT id,
    ip_address,
    country_id,
    country_code,
    city,
    -- COALESCE(ST_X(location), 0)::FLOAT as latitude,
    -- COALESCE(ST_Y(location), 0)::FLOAT as longitude,
    latitude,
    longitude,
    mystery_value
FROM "ip_locations"
WHERE ip_address = $1
LIMIT 1;

-- name: GetCountIPLocationsByIPAddress :one
SELECT COUNT(*) FROM "ip_locations"
WHERE ip_address = $1;

-- name: InsertIPLocation :copyfrom
INSERT INTO "ip_locations" (
    id, 
    ip_address, 
    country_id, 
    country_code, 
    city, 
    latitude, 
    longitude,
    mystery_value
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8
);

-- name: InsertIPLocationWIP :exec
INSERT INTO "ip_locations" (
    id, 
    ip_address, 
    country_id, 
    country_code, 
    city, 
    location, 
    mystery_value
) VALUES (
    sqlc.arg(id)::TEXT, 
    sqlc.arg(ip_address)::TEXT,
    sqlc.narg(country_id)::INT,
    sqlc.narg(country_code)::TEXT,
    sqlc.narg(city)::TEXT,
    ST_MakePoint(sqlc.narg('latitude')::FLOAT, 
    sqlc.narg('longitude')::FLOAT), 
    sqlc.narg(mystery_value)::BIGINT
);