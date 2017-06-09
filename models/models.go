package models

import (
	"strings"
)

// Server : data model for whatsdeployed server
type Server struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SortPriority int    `json:"sort_priority"`
}

// Deployment : represents an application deployment object
type Deployment struct {
	Server      string
	Application string
	Branch      string
	Version     string
}

// Application : data model for whatsdeployed application
type Application struct {
	ID       int    `json:"id"`
	ServerID int    `json:"server_id"`
	Package  string `json:"package"`
	Active   bool   `json:"active"`
}

// ToDeployment : transform application into deployment model given a serverName
func (a Application) ToDeployment(serverName string) Deployment {
	parts := strings.SplitN(a.Package, "-", 3)

	d := Deployment{
		Application: parts[0],
		Version:     parts[1],
		Branch:      parts[2],
		Server:      serverName,
	}
	return d
}
