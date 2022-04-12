package main

import (
	"flag"
	"io"
	"io/fs"
	"log"
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
	if info.IsDir() {
		return nil
	}

	if err != nil {
		log.Printf("error handed to walkfunc: %s", err)
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		log.Printf("error opening note: %s\n", path)
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Printf("error ingesting note: %s\n", err)
	}

	document := ma.Parser.Parse(text.NewReader(b))
	metaData := document.OwnerDocument().Meta()
	log.Printf("%#v\n", metaData)

	return nil
}
