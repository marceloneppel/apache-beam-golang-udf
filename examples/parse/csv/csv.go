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

const exampleDirectory = "examples/parse/csv"

type csvParse struct {
	Location string
}

func (c *csvParse) ProcessElement(ctx context.Context, line string, emit func(string)) {
	function, err := udf.GetFunction(ctx, common.JoinPath(c.Location, filepath.Join(exampleDirectory, "csvparse.go")), "csv", "Parse")
	if err != nil {
		log.Infof(ctx, err.Error())
		return
	}
	emit(function.(func(string) string)(line))
}

func init() {
	beam.RegisterType(reflect.TypeOf((*csvParse)(nil)).Elem())
}

func main() {
	flag.Parse()

	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	location := *common.Location
	if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
		location = ""
	}

	lines := textio.Read(s, common.JoinPath(location, filepath.Join(exampleDirectory, "file.csv")))

	processedLines := beam.ParDo(s, &csvParse{
		Location: *common.Location,
	}, lines)

	debug.Print(s, processedLines)

	beamx.Run(context.Background(), p)
}
