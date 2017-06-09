package main

import (
	"fmt"
	"os"
	"whatsdeployed/api"
	"whatsdeployed/models"
)

func main() {

	var serverNameProvided bool
	var serverName string

	if len(os.Args) >= 2 {
		serverNameProvided = true
		serverName = os.Args[1]
		fmt.Printf("Searching for deployments on %s...\n", serverName)
	}

	if serverNameProvided {
		deployments := fetchDeployments(serverName)
		prettyPrint(deployments)
	}
}

/* Used to figure out the length of the longest string */
func compareAgainstLongest(longest *int, current string) {
	currentLength := len(current)
	if currentLength > *longest {
		*longest = currentLength
	}
}

func padRight(str string, length int) string {
	for {
		str += " "
		if len(str) > length {
			return str[0:length]
		}
	}
}

func prettyPrint(deployments []models.Deployment) {

	var serverLen, applicationLen, branchLen, versionLen int

	for _, d := range deployments {
		compareAgainstLongest(&serverLen, d.Server)
		compareAgainstLongest(&applicationLen, d.Application)
		compareAgainstLongest(&branchLen, d.Branch)
		compareAgainstLongest(&versionLen, d.Version)
	}

	// Print colunn headers
	fmt.Printf("%s %s %s %s\n",
		padRight("Server", serverLen),
		padRight("App", applicationLen),
		padRight("Branch", branchLen),
		padRight("Version", versionLen))

	for _, d := range deployments {
		fmt.Printf("%s %s %s %s\n",
			padRight(d.Server, serverLen),
			padRight(d.Application, applicationLen),
			padRight(d.Branch, branchLen),
			padRight(d.Version, versionLen))
	}
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
