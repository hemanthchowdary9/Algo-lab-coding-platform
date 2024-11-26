package models

type JDoodleRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script,omitempty"`
	Stdin        string `json:"stdin,omitempty"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

type CreditSpent struct {
	Used int `json:"used"`
}
