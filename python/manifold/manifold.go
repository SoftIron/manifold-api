// Package main writes sifi constants to a python file.
package main

import (
	"fmt"
	"os"

	"github.com/softiron/manifold-api/manifold"
)

//go:generate go run .

func main() {
	fout, err := os.Create("defaults.py")
	if err != nil {
		fmt.Printf("failed to create defaults.py: %v\n", err)
		os.Exit(1)
	}

	defer fout.Close()

	fmt.Fprintf(fout, "__api_prefix__ = %q\n", manifold.APIPrefix)
	fmt.Fprintf(fout, "__api_version__ = %q\n", manifold.APIVersion)
}
