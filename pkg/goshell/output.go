package goshell

type Output struct {
	AgentId  string `json:"agentId,omitempty"`
	HostName string `json:"hostname,omitempty"`
	ScriptId string `json:"scriptId,omitempty"`
	Output   string `json:"output,omitempty"`
}
