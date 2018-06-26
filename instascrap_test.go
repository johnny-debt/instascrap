package instascrap

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func getTestDataPath() string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, "test-data")
}

func ReadTestDataFile(filename string) []byte {
	data, _ := ioutil.ReadFile(filepath.Join(getTestDataPath(), filename))
	return data
}
