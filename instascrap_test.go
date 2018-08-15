package instascrap

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstascrapDefaultCreation(t *testing.T) {
	NewInstascrap(nil)
}

func getEmptyHTTPClient() *http.Client {
	return nil
}
func TestCustomHTTPClientProvider(t *testing.T) {
	instascrap := NewInstascrap(getEmptyHTTPClient)
	_, error := instascrap.getDataFromURL("http://google.com", nil)
	assert.Error(t, error)
}

func getTestDataPath() string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, "test-data")
}

func ReadTestDataFile(filename string) []byte {
	data, _ := ioutil.ReadFile(filepath.Join(getTestDataPath(), filename))
	return data
}
