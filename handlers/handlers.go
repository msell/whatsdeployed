package handlers

import (
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
