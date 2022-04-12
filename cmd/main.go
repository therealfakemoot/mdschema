package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	// "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func main() {
	var (
		vaultRoot string
	)

	flag.StringVar(&vaultRoot, "root", "root directory to walk", ".")

	flag.Parse()

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)

	ma := MetaAccumulator{
		Keys:   make([]string, 0),
		Schema: make(map[string]Policy),
		Parser: markdown.Parser(),
	}
	filepath.Walk(vaultRoot, ma.WalkFunc)

}

type PolicyRange interface {
	Valid(interface{}) bool
}

type Policy struct {
	Key   string
	Type  string
	Range PolicyRange
}

type MetaAccumulator struct {
	Keys   []string
	Schema map[string]Policy
	Parser parser.Parser
}

func (ma *MetaAccumulator) WalkFunc(path string, info fs.FileInfo, err error) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	io.Copy(f, &buf)
	document := ma.Parser.Parse(text.NewReader(buf.Bytes()))
	metaData := document.OwnerDocument().Meta()
	fmt.Printf("%#v\n", metaData)

	return nil
}
