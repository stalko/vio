-- name: GetCountryByID :one
SELECT * FROM "countries"
WHERE id = $1
LIMIT 1;

-- name: InsertCountry :one
WITH new_row AS (
	INSERT INTO "countries" (id, name)
	SELECT sqlc.arg(id)::TEXT, sqlc.arg(name)::TEXT
	WHERE NOT EXISTS (SELECT * FROM "countries" WHERE name = sqlc.arg(name)::TEXT)
	ON CONFLICT DO NOTHING
	RETURNING id, name
)
SELECT id, name FROM new_row
UNION
SELECT id, name FROM "countries" WHERE name = sqlc.arg(name)::TEXT;