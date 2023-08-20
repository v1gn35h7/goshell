package goshell

type Output struct {
	Id       string
	Agentid  string `json:"agentId,omitempty"`
	Hostname string `json:"hostName,omitempty"`
	Scriptid string `json:"scriptId,omitempty"`
	Output   string `json:"output,omitempty"`
	Score    string `json:"Score,omitempty"`
}
