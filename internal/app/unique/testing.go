package unique

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestUnique(t *testing.T, testOk, testOkResult string) error {
	input := bufio.NewReader(strings.NewReader(testOk))
	output := new(bytes.Buffer)
	err := Unique(input, output)
	if err != nil {
		return err
	}
	result := output.String()
	if result != testOkResult {
		return err
	}
	return nil
}
