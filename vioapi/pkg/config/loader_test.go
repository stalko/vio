package config_test

import (
	"os"
	"testing"

	"github.com/stalko/vioapi/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	type TestStruct struct {
		Foo  string `env:"foo"`
		Port int    `env:"test_port"`
	}
	os.Setenv("foo", "bar")
	os.Setenv("test_port", "22001")

	expectData := TestStruct{
		Foo:  "bar",
		Port: 22001,
	}

	actualData, err := config.Load[TestStruct]()
	assert.Equal(t, expectData, actualData)
	assert.NoError(t, err)
}
