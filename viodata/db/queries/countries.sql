-- name: GetCountryByID :one
SELECT * FROM "countries"
WHERE id = $1
LIMIT 1;

-- name: InsertCountry :one
WITH new_row AS (
	INSERT INTO "countries" (id, name)
	SELECT $1::text, $2::text
	WHERE NOT EXISTS (SELECT * FROM "countries" WHERE name = $2::text)
	RETURNING id, name
)
SELECT id, name FROM new_row
UNION
SELECT id, name FROM "countries" WHERE name = $2::text;