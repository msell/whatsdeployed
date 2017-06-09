package models

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

// ToDeployment : transform application into deployment model given a serverName
func (a Application) ToDeployment(serverName string) Deployment {
	//SecurityAPIv2-1.0.1.1-REL-1317
	// TODO: Split a.Package into proper fields
	d := Deployment{
		Application: "todo",
		Branch:      "master",
		Version:     "0.0.0.1",
		Server:      serverName,
	}
	return d
}

// Application : data model for whatsdeployed application
type Application struct {
	ID       int    `json:"id"`
	ServerID int    `json:"server_id"`
	Package  string `json:"package"`
	Active   bool   `json:"active"`
}
