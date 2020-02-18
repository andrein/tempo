package friggdb

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/grafana/frigg/friggdb/backend/local"
	"github.com/stretchr/testify/assert"
)

func TestCompactor(t *testing.T) {
	tempDir, err := ioutil.TempDir("/tmp", "")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err, "unexpected error creating temp dir")

	r, w, err := New(&Config{
		Backend: "local",
		Local: &local.Config{
			Path: path.Join(tempDir, "traces"),
		},
		WALFilepath:              path.Join(tempDir, "wal"),
		IndexDownsample:          17,
		BloomFilterFalsePositive: .01,
		BlocklistRefreshRate:     30 * time.Minute,
	}, log.NewNopLogger())
	assert.NoError(t, err)

}

/*func createAndWriteBlock(w Writer) ([][]ID, [][]byte, uuid.UUID) {
	wal, err := w.WAL()
}*/
