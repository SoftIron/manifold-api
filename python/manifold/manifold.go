// Package main writes sifi constants to a python file.
package main

import (
	"fmt"
	"os"

	"github.com/softiron/manifold-api/manifold"
)

//go:generate go run .

func main() {
	fout, err := os.Create("__init__.py")
	if err != nil {
		fmt.Printf("failed to create __init__.py: %v\n", err)
		os.Exit(1)
	}

	defer fout.Close()

	fmt.Fprintf(fout, "__default_api_prefix__ = %q\n", manifold.APIPrefix)
	fmt.Fprintf(fout, "__default_api_version__ = %q\n", manifold.APIVersion)
	fmt.Fprintf(fout, "__default_api_port__ = %d\n", manifold.PortNumber+1) // ha-proxy public port
}
