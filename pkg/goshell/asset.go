package goshell

type Asset struct {
	Agentid         string `json:"agentId"`
	Platform        string `json:"platform"`
	Operatingsystem string `json:"operatingSystem"`
	Architecture    string `json:"architecture"`
	Hostname        string `json:"hostName"`
	Synctime        string `json:"syncTime"`
}
