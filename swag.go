// Package main imports the SIFI Daemon swagger documentation and writes it to a
// file. This requires access to Sofiron's private git instance, and is not
// intended to be used outside of Softiron.
//
// Prior to any software release use this code to re-generate the swagger.json
// file.
package main

import (
	"os"
	
	"git.softiron.com/sw/hc/sifi.git/cmd/sifid/docs"
)

//go:generate go run .

func main() {
	fout, err := os.OpenFile("swagger.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer fout.Close()
	
	doc := docs.SwaggerInfo.ReadDoc()
	if _, err := fout.WriteString(doc); err != nil {
		panic(err)
	}
}
