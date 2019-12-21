package udf

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/memfs"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/apache/beam/sdks/go/pkg/beam/util/gcsx"
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

const globNotFound = "glob not found"
const multipleDataLoadedForSameGlob = "multiple data loaded for same glob"

// DownloadGlob shows
func DownloadGlob(ctx context.Context, glob string) ([]byte, error) {
	fmt.Println(glob)
	if strings.HasPrefix(glob, "gs://") {
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		bucket, object, err := gcsx.ParseObject(glob)
		if err != nil {
			return nil, err
		}
		reader, err := client.Bucket(bucket).Object(object).NewReader(ctx)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if strings.HasPrefix(glob, "http://") || strings.HasPrefix(glob, "https://") {
		response, err := http.Get(glob)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	// TODO: check for FTP URL.
	return ioutil.ReadFile(glob)
}

// GetFunction shows
func GetFunction(ctx context.Context, glob string, functionPackage string, functionName string) (interface{}, error) {
	src := ""

	filesystem := memfs.New(ctx)
	defer filesystem.Close()

	filenames, err := filesystem.List(ctx, glob)
	if err != nil {
		return nil, err
	}

	switch len(filenames) {
	// TODO: add redownload mechanism.
	case 0:
		log.Infof(ctx, "downloading glob...")
		downloadedGlob, err := DownloadGlob(ctx, glob)
		if err != nil {
			return nil, err
		}
		if downloadedGlob != nil {
			writer, err := filesystem.OpenWrite(ctx, glob)
			if err != nil {
				return nil, err
			}
			defer writer.Close()
			_, err = writer.Write(downloadedGlob)
			if err != nil {
				return nil, err
			}
			src = string(downloadedGlob)
			break
		}
		return nil, errors.New(globNotFound)
	case 1:
		log.Infof(ctx, "loading glob...")
		reader, err := filesystem.OpenRead(ctx, glob)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(reader)
		src = buffer.String()
		break
	default:
		return nil, errors.New(multipleDataLoadedForSameGlob)
	}

	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err = i.Eval(src)
	if err != nil {
		return nil, err
	}

	v, err := i.Eval(functionPackage + "." + functionName)
	if err != nil {
		return nil, err
	}

	return v.Interface(), nil
}
