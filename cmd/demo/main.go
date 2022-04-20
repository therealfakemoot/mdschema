package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"

	"github.com/therealfakemoot/mdschema"
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

	ma := mdschema.MetaAccumulator{
		Keys:   make(map[string]bool),
		Schema: make(map[string]mdschema.SchemaKey),
		Parser: markdown.Parser(),
	}
	err := filepath.Walk(vaultRoot, ma.WalkFunc)
	if err != nil {
		log.Printf("error walking root: %s", err)
	}

	log.Printf("%#v\n", ma)
}
