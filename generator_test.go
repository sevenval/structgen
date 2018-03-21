package structgen_test

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/sevenval/structgen"
)

func TestNewGenerator(t *testing.T) {
	t.SkipNow() // currently just a debug test and example usage, TODO test :)
	cwd, err := os.Getwd()
	must(err)
	schemaFile, err := ioutil.ReadFile(path.Join(cwd, "testdata", "config_schema_v3.6.json"))
	must(err)
	schema, err := structgen.NewSchema(schemaFile)
	must(err)
	gen := structgen.NewGenerator("ComposeFile", "structgen_test", schema)
	file, err := os.OpenFile(path.Join(cwd, "out.go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	must(err)
	// io.Copy(ioutil.Discard, gen)
	_, err = io.Copy(file, gen)
	must(err)
	file.Close()
}

func must(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
