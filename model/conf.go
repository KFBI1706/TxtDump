package model

//Configuration contains all the configurable variables for this app
type Configuration struct {
	Port             int
	DBStringLocation string `json:DBStringLocation,omitempty`
	Path             string
	Production       bool `json:Production,omitempty`
	CSRFString       string
}
