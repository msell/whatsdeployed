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

// Application : data model for whatsdeployed application
type Application struct {
	ID       int    `json:"id"`
	ServerID int    `json:"server_id"`
	Package  string `json:"package"`
	Active   bool   `json:"active"`
}
