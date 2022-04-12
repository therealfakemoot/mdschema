package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/therealfakemoot/mdschema"
)

func main() {
	var (
		schemaPath string
		root       string
	)

	flag.StringVar(&schemaPath, "schema", "schema.toml", "path to schema file")
	flag.StringVar(&root, "root", "notes/", "path to root of notes dir")

	flag.Parse()

	f, err := os.Open(schemaPath)
	if err != nil {
		log.Fatalf("error opening schema file: %s\n", err)
	}
	p, err := mdschema.LoadSchema(f)

	fmt.Printf("%#v\n", p)
}
