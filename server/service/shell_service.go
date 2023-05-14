package service

import (
	"github.com/v1gn35h7/goshell/server/goshell"
)

type shellService interface {
	ExecuteCmd(cmd string) (string, error)
	ConnectToRemoteHost(hostId string) (bool, error)
	GetScripts(AgentId string) ([]*goshell.ShellScript, error)
	//EndpointHeartBeat(hostId string) ([]execu)
}

func (srvc service) ExecuteCmd(cmd string) (string, error) {
	return "ellow!!", nil
}

func (srvc service) ConnectToRemoteHost(hostId string) (bool, error) {
	return true, nil
}

func (srvc service) GetScripts(agentId string) ([]*goshell.ShellScript, error) {
	scripts := make([]*goshell.ShellScript, 0)
	scripts = append(scripts, &goshell.ShellScript{
		Script: "ls",
		Args:   "",
		Type:   "SINGLE",
	})
	scripts = append(scripts, &goshell.ShellScript{
		Script: "date",
		Args:   "",
		Type:   "SINGLE",
	})
	scripts = append(scripts, &goshell.ShellScript{
		Script: "Get-NetAdapter -Name * -IncludeHidden",
		Args:   "",
		Type:   "SINGLE",
	})
	return scripts, nil
}
