package importer_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stalko/viodata/importer"
	"github.com/stalko/viodata/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestImport(t *testing.T) {
	const (
		csv = `ip_address,country_code,country,city,latitude,longitude,mystery_value
200.106.141.15,SI,Nepal,DuBuquemouth,-84.87503094689836,7.206435933364332,7823011346
160.103.7.140,CZ,Nicaragua,New Neva,-68.31023296602508,-37.62435199624531,7301823115
70.95.73.73,TL,Saudi Arabia,Gradymouth,-49.16675918861615,-86.05920084416894,2559997162
,PY,Falkland Islands (Malvinas),,75.41685191518815,-144.6943217219469,0
125.159.20.54,LI,Guyana,Port Karson,-78.2274228596799,-163.26218895343357,1337885276`
		fileName                 = "test.csv"
		countGoRoutine           = 1
		expectedAcceptedEntries  = 4
		expectedDiscardedEntries = 1
		expectedDurationLessThan = 2 * time.Second
		countBulkInsert          = 1
	)
	t.Parallel()
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)

	for i := 0; i < expectedAcceptedEntries; i++ {
		mockStorage.EXPECT().BulkInsertIPLocation(ctx, gomock.Any()).Return(nil)
	}

	//------create tmp directory and file-----------
	dir := os.TempDir()
	defer os.RemoveAll(dir)

	fileAddress := fmt.Sprintf("%s%s", dir, fileName)

	file, err := os.Create(fileAddress)
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = file.Close()
	assert.NoError(t, err)
	//-----------------------------------------------

	csvImporter := importer.NewCSVImporter(mockStorage, countBulkInsert, zap.NewExample(), ctx)

	out, err := csvImporter.Import(fileAddress, countGoRoutine)
	assert.NoError(t, err)

	if assert.NotNil(t, out) {
		assert.Equal(t, out.AcceptedEntries, expectedAcceptedEntries)
		assert.Equal(t, out.DiscardedEntries, expectedDiscardedEntries)
		if out.Duration > expectedDurationLessThan {
			assert.Error(t, fmt.Errorf("duration took more than expected: %s", out.Duration))
		}
	}
}
