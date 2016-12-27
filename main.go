package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hackebrot/amelia/amelia"
)

func loadFromEnv(keys ...string) (map[string]string, error) {
	env := make(map[string]string)

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("environment variable %q is required", key)
		}
		env[key] = v
	}

	return env, nil
}

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

	env, err := loadFromEnv("AMELIA_USERNAME", "AMELIA_TOKEN")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	n, err := amelia.NewGist(description, public, files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	client := &amelia.GitHubClient{
		Username: env["AMELIA_USERNAME"],
		Token:    env["AMELIA_TOKEN"],
	}

	g, err := client.CreateGist(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("gist created at %s\n", *g.HTMLURL)
}
