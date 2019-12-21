package common

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

var Location = flag.String("location", filepath.Join(os.Getenv("GOPATH"), "src/github.com/marceloneppel/apache-beam-golang-udf"), "example files location")

/*type DownloadFile struct {
	Glob string
}

func (d *DownloadFile) ProcessElement(glob string) (string, error) {
	response, err := http.Get(glob)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}*/

func JoinPath(location, path string) string {
	if strings.HasPrefix(location, "gs://") {
		if strings.HasSuffix(location, "/") {
			return location + path
		}
		return location + "/" + path
	}
	if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
		if strings.HasSuffix(location, "/") {
			return location + path
		}
		return location + "/" + path
	}
	return filepath.Join(location, path)
}
