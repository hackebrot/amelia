package main

import (
	"flag"
	"fmt"
	"os"
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

	fmt.Printf("description %q\n", *description)
	fmt.Printf("public %v\n", *public)
	fmt.Printf("files %v\n", files)
}
