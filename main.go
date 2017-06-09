package main

import (
	"fmt"
	"log"
	"os"
	"whatsdeployed/api"
	"whatsdeployed/models"
	"whatsdeployed/utils"
)

func main() {

	var serverNameProvided bool
	var serverName string

	if len(os.Args) >= 2 {
		serverNameProvided = true
		serverName = os.Args[1]
		fmt.Printf("Searching for deployments on %s...\n", serverName)
	}

	if !serverNameProvided {
		log.Fatal("You must pass a server name as the first argument")
	}

	deployments := fetchDeployments(serverName)
	utils.PrettyPrint(deployments)

}

func fetchDeployments(serverName string) []models.Deployment {
	serverID := api.FetchServerID(serverName)
	apps := api.FetchApplications(serverID)

	var deployments []models.Deployment
	for _, app := range apps {
		deployments = append(deployments, app.ToDeployment(serverName))
	}

	return deployments
}
