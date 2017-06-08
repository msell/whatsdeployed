package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func prettyPrint(deployments []deployment) {

	var serverLen, applicationLen, branchLen, versionLen int

	for _, d := range deployments {
		compareAgainstLongest(&serverLen, d.server)
		compareAgainstLongest(&applicationLen, d.application)
		compareAgainstLongest(&branchLen, d.branch)
		compareAgainstLongest(&versionLen, d.version)
	}

	// Print colunn headers
	fmt.Printf("%s %s %s %s\n",
		padRight("Server", serverLen),
		padRight("App", applicationLen),
		padRight("Branch", branchLen),
		padRight("Version", versionLen))

	for _, d := range deployments {
		fmt.Printf("%s %s %s %s\n",
			padRight(d.server, serverLen),
			padRight(d.application, applicationLen),
			padRight(d.branch, branchLen),
			padRight(d.version, versionLen))
	}
}

func fetchDeployments(serverName string) []deployment {
	res, err := http.Get("http://whatsdeployed.herokuapp.com/servers.json")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var servers []Server

	err = json.NewDecoder(res.Body).Decode(&servers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(servers)

	s1 := deployment{
		server:      serverName,
		application: "SPL",
		branch:      "master",
		version:     "3.0.0.2"}

	s2 := deployment{
		server:      serverName,
		application: "Lead API",
		branch:      "master",
		version:     "1.3.0.9"}

	return []deployment{s1, s2}
}

// Server : data represntation of server object
type Server struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SortPriority int    `json:"sort_priority"`
}

type deployment struct {
	server      string
	application string
	branch      string
	version     string
}
