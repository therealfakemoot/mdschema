package mdschema

import (
	"io"
	"io/fs"
	"log"
	"os"

	"github.com/yuin/goldmark/text"
	// "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type MetaAccumulator struct {
	Keys   map[string]bool
	Schema map[string]Config
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

	for k, _ := range metaData {
		ma.Keys[k] = true
	}

	return nil
}
