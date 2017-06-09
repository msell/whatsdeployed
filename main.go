package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"whatsdeployed/handlers"
)

func main() {

	var serverNameProvided bool
	var serverName string
	var isDiff bool

	flag.BoolVar(&isDiff, "diff", false, "Performs a diff against one or more servers")
	flag.Parse()

	if isDiff {
		fmt.Println("Diff requested...")
		if len(os.Args) < 4 {
			log.Fatal("You must provide at least 2 server names after the diff flag")
		}

		// TODO: make Diff a variadic function that takes N servers for comparison
		servers := []string{os.Args[2], os.Args[3]}
		handlers.Diff(servers)

		fmt.Println("Diff functionality is not implemented yet. 🐳")
		os.Exit(0)
	} else {
		// User is trying to get a list of deployments on a single server
		if len(os.Args) >= 2 {
			serverNameProvided = true
			serverName = os.Args[1]
			fmt.Printf("Searching for deployments on %s...\n", serverName)
		}

		if !serverNameProvided {
			log.Fatal("You must pass a server name as the first argument")
		}

		handlers.WhatsDeployedOn(serverName)
	}
}
