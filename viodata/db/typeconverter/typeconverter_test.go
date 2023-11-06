package typeconverter_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stalko/viodata/db/typeconverter"
	"github.com/stretchr/testify/assert"
)

func TestNewNullString(t *testing.T) {
	str := "string"

	res := typeconverter.NewNullString(&str)

	assert.Equal(t, res.Valid, true)
	assert.Equal(t, res.String, str)
}

func TestNewNullStringNull(t *testing.T) {
	res := typeconverter.NewNullString(nil)

	assert.Equal(t, res.Valid, false)
}

func TestNewStringPointer(t *testing.T) {
	v := pgtype.Text{
		String: "string",
		Valid:  true,
	}

	res := typeconverter.NewStringPointer(v)

	assert.NotNil(t, res)
	assert.Equal(t, *res, v.String)
}

func TestNewStringPointerNull(t *testing.T) {
	v := pgtype.Text{
		Valid: false,
	}

	res := typeconverter.NewStringPointer(v)

	assert.Nil(t, res)
}

func TestNewNullInt64(t *testing.T) {
	var v int64 = 123

	res := typeconverter.NewNullInt64(&v)

	assert.Equal(t, res.Valid, true)
	assert.Equal(t, res.Int64, v)
}

func TestNewNullInt64gNull(t *testing.T) {
	res := typeconverter.NewNullInt64(nil)

	assert.Equal(t, res.Valid, false)
}

func TestNewInt64Pointer(t *testing.T) {
	v := pgtype.Int8{
		Int64: 123,
		Valid: true,
	}

	res := typeconverter.NewInt64Pointer(v)

	assert.NotNil(t, res)
	assert.Equal(t, *res, v.Int64)
}

func TestNewInt64PointerNull(t *testing.T) {
	v := pgtype.Int8{
		Valid: false,
	}

	res := typeconverter.NewInt64Pointer(v)

	assert.Nil(t, res)
}

func TestNewNullFloat64(t *testing.T) {
	var v float64 = 123

	res := typeconverter.NewNullFloat64(&v)

	assert.Equal(t, res.Valid, true)
	assert.Equal(t, res.Float64, v)
}

func TestNewNullFloat64Null(t *testing.T) {
	res := typeconverter.NewNullFloat64(nil)

	assert.Equal(t, res.Valid, false)
}
