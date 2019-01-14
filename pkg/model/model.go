package model

// Results is for DNS results
type Results struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	TXT      string `json:"TXT"`
}

// UniqResults is for Unique DNS Results
type UniqResults struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	TXT      string `json:"TXT"`
}
