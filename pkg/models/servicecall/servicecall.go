package servicecall

// ServiceCall is a struct that defines how http/json communication will be structured
type ServiceCall struct {
	Cmd  string `json:"cmd"`  // Commands are: notify, verify & destroy
	IP   string `json:"ip"`   // IP of the microservice
	Tipo string `json:"type"` // Type of servie can be anything e.g. "database", "webserver", "gapi"
	Port string `json:"port"`
}
