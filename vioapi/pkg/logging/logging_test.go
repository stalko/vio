package logging_test

import (
	"testing"

	"github.com/stalko/vioapi/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	logger := logging.CreateLogger("staging", "https://ALPHANUMERIC@sentry.io/NUMBER", nil)
	assert.NotNil(t, logger, "Failed to create logger")
}
