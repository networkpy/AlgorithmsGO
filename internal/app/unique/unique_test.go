package unique_test

import (
	"testing"

	"github.com/NetworkPy/algh/internal/app/unique"
	"github.com/stretchr/testify/assert"
)

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestUnique(t *testing.T) {
	res := unique.TestUnique(t, testOk, testOkResult)
	assert.NoError(t, res)
}

var testFail = `1
2
1`

func TestUniqueError(t *testing.T) {
	res := unique.TestUnique(t, testFail, testOkResult)
	assert.Error(t, res)
}
