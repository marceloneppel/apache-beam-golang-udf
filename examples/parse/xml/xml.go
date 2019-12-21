package main

import (
	"context"
	"flag"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/apache/beam/sdks/go/pkg/beam"
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/gcs"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/go/pkg/beam/x/debug"
	udf "github.com/marceloneppel/apache-beam-golang-udf"
	"github.com/marceloneppel/apache-beam-golang-udf/examples/common"
)

const exampleDirectory = "examples/parse/xml"

type xmlParse struct {
	Location string
}

func (c *xmlParse) ProcessElement(ctx context.Context, line string, emit func(string)) {
	function, err := udf.GetFunction(ctx, common.JoinPath(c.Location, filepath.Join(exampleDirectory, "xmlparse.go")), "xml", "Parse")
	if err != nil {
		log.Infof(ctx, err.Error())
		return
	}
	emit(function.(func(string) string)(line))
}

func init() {
	beam.RegisterType(reflect.TypeOf((*xmlParse)(nil)).Elem())
}

func main() {
	flag.Parse()

	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	/*var lines beam.PCollection
	if strings.HasPrefix(*common.Location, "http://") || strings.HasPrefix(*common.Location, "https://") {
		lines, _ = beam.ParDo2(s, common.DownloadFile, beam.Impulse(s))
	} else {
		lines = textio.Read(s, common.JoinPath(*common.Location, filepath.Join(exampleDirectory, "file.xml")))
	}*/

	location := *common.Location
	if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
		location = ""
	}

	lines := textio.Read(s, common.JoinPath(location, filepath.Join(exampleDirectory, "file.xml")))

	// Combine XML in one line.
	concatenatedLines := beam.Combine(s, func(a, b string) string {
		return strings.Trim(a, " ") + strings.Trim(b, " ")
	}, lines)

	processedLines := beam.ParDo(s, &xmlParse{
		Location: *common.Location,
	}, concatenatedLines)

	debug.Print(s, processedLines)

	beamx.Run(context.Background(), p)
}
