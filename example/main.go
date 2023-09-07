// Package main is an example sifid client that prints out the name of
// the instance with an ID of 0.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/sifi"
)

func main() {
	fs := flag.NewFlagSet("example", flag.ExitOnError)

	// Both flag.FlagSet and cobra's pflag.FlagSet implement
	// client.flagSet. NewOptions adds the flags that all REST clients
	// will need.

	opts := client.NewOptions(fs, "", sifi.PortNumber)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// If running on the same host as the HyperCloud dashboard and
	// running as the admin (or root) user, it is possible to grab HC
	// admin username and password and use that to Login.
	//
	// But this is not an option for opensource code as the auth
	// package is part of the sifi module.
	//
	// opts.Username, opts.Password, err = auth.HyperCloudAdmin()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	c := sifi.NewClient(opts)

	// Set c.Logger to attach a logger to the client. The slog package
	// implements this interface (structured logger new to go 1.21).
	// All REST calls are logged at the debug level.

	// The Login call is optional as any unauthorized call will
	// trigger a Login.
	//
	// c.Login(context.Background())

	// Request information about the first instance.
	resp, err := c.Instance(context.Background(), 0)
	if err != nil {
		log.Fatal(err)
	}

	// This could also be written as below. The sifi.Client is
	// composed of more than a dozen services to organize calls into
	// logical groups. It is possible to pass a single service to a
	// function to limit what groups of calls are available.
	//
	// resp, err := c.InstanceService.Instance(context.Background(), 0)

	fmt.Printf("Instance %v is called: %v", 0, resp.Instance.Name)
}
