package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"whatsdeployed/models"
)

// FetchApplications : Makes API call to fetch applications
func FetchApplications(serverID int) []models.Application {
	res, err := http.Get(fmt.Sprintf(
		"%s%d%s",
		"http://whatsdeployed.herokuapp.com/servers/",
		serverID,
		"/deployed_apps.json",
	))

	if err != nil {
		log.Fatal("Could not fetch applications ", err)
	}

	defer res.Body.Close()

	var apps []models.Application

	err = json.NewDecoder(res.Body).Decode(&apps)
	if err != nil {
		log.Fatal("Could not decode applications json ", err)
	}

	return apps
}

// FetchServerID : Make api call to get servers and find serverID based on server name
func FetchServerID(serverName string) int {
	res, err := http.Get("http://whatsdeployed.herokuapp.com/servers.json")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var servers []models.Server

	err = json.NewDecoder(res.Body).Decode(&servers)

	if err != nil {
		log.Fatal(err)
	}

	var serverID int
	for _, s := range servers {
		if s.Name == strings.ToUpper(serverName) {
			serverID = s.ID
		}
	}

	if serverID == 0 {
		log.Fatal("Server " + serverName + " does not exist")
	}

	return serverID
}
