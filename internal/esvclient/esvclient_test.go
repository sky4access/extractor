package esvclient

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestESVClient_GetVerses(t *testing.T) {

	assert.NotEmpty(t, os.Getenv("ESV_API_KEY"))
	esv, err := NewClient(2 * time.Second)
	assert.Nil(t, err)

	passage, err := esv.GetVerses("john 1:1")
	assert.Nil(t, err)
	assert.NotEmpty(t, passage)
	t.Log(passage)

	passage, err = esv.GetVerses("rev 1:1")
	assert.Nil(t, err)
	assert.NotEmpty(t, passage)
	t.Log(passage)
}
