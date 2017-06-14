package utils

import (
	"fmt"
	"strings"
	"whatsdeployed/models"
)

/* Used to figure out the length of the longest string */
func determineFieldWidth(longest *int, current string) {
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
		// fmt.Println(d)
	}

	fmt.Println(distinctApps)

	// loop distinct apps, each iteration of distinct apps will be displayed
	// as a row in a table.  The first column will be the application name
	// subsequent columns will display the package deployed on a server

	var applicationLen int

	numberOfServers := len(serverMap)

	fmt.Println(numberOfServers)
	fmt.Println(serverMap)
	for _, app := range distinctApps {

		determineFieldWidth(&applicationLen, app)
		var fieldWidths []int
		i := 0
		for k := range serverMap {
			fieldWidths = make([]int, numberOfServers)
			pkg := getPackageName(apps, app, k)
			determineFieldWidth(&fieldWidths[i], pkg)
			i++
		}

		fmt.Println(fieldWidths)

	}

	// for _, app := range distinctApps {
	// 	determineFieldWidth(&applicationLen, app)
	// }

}

func getPackageName(apps []models.Application, appName string, serverID int) string {
	for _, a := range apps {
		if a.ServerID == serverID &&
			strings.HasPrefix(a.Package, appName) {
			return a.Package
		}
	}

	return ""
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
		determineFieldWidth(&serverLen, d.Server)
		determineFieldWidth(&applicationLen, d.Application)
		determineFieldWidth(&branchLen, d.Branch)
		determineFieldWidth(&versionLen, d.Version)
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
