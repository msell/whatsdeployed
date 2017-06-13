package utils

import (
	"fmt"
	"whatsdeployed/models"
)

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

// PrettyPrintDiff : given a list of deployments and server maps write the results to console
func PrettyPrintDiff(apps []models.Application, serverMap map[int]string) {
	var distinctApps []string
	for _, app := range apps {
		d := app.ToDeployment(serverMap[app.ID])
		if isAppDistinct(distinctApps, d.Application) {
			distinctApps = append(distinctApps, d.Application)
		}
		fmt.Println(d)
	}

	fmt.Println(distinctApps)

	// var applicationLen, server1Len, server2Len int
	// for _, app := range distinctApps {
	// 	// will add a row for each distinct app
	// }
}

func isAppDistinct(apps []string, app string) bool {
	for _, a := range apps {
		if a == app {
			return false
		}
	}
	return true
}

// PrettyPrint : given a list of deployments format and write results to console
func PrettyPrint(deployments []models.Deployment) {

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
