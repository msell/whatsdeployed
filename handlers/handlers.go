package handlers

import (
	"fmt"
	"whatsdeployed/api"
	"whatsdeployed/models"
	"whatsdeployed/utils"
)

// WhatsDeployedOn : Writes a list of deployments to the console for a given server
func WhatsDeployedOn(serverName string) {
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

// Diff : Compares deployments on 2 or more servers and writes results to the console
func Diff(servers []string) {

	var serverIDs []int
	for _, s := range servers {
		serverIDs = append(serverIDs, api.FetchServerID(s))
	}

	fmt.Println(serverIDs)

}
