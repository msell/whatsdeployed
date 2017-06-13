package handlers

import (
	"fmt"
	"strings"
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

	var apps []models.Application
	var serverMap = make(map[int]string)

	for _, s := range servers {
		serverID := api.FetchServerID(s)
		apps = append(apps, api.FetchApplications(serverID)...)
		serverMap[serverID] = strings.ToUpper(s)
	}

	fmt.Println(apps)
	fmt.Println(serverMap)

}
