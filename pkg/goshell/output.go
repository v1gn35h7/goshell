package goshell

type Output struct {
	Id       string
	Agentid  string `json:"agentid,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Scriptid string `json:"scriptid,omitempty"`
	Output   string `json:"output,omitempty"`
	Score    string `json:"Score,omitempty"`
}
