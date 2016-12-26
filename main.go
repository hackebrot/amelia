package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hackebrot/amelia/amelia"
)

func main() {
	description := flag.String(
		"description",
		"Amelia Gist",
		"A description of the gist.",
	)

	public := flag.Bool(
		"public",
		false,
		"Indicates whether the gist is public. (default false)",
	)

	flag.Parse()

	files := flag.Args()

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", "no files given")
		os.Exit(1)
	}

	g, err := amelia.NewGist(description, public, files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("gist %+v\n", g)
}
