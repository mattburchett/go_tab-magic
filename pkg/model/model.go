package model

type Results struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	TXT      string `json:"TXT`
}

type UniqResults struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	TXT      string `json:"TXT"`
}
