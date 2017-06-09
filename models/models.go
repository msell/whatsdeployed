package models

// Server : data represntation of server object
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
