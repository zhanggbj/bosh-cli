package fakes

import (
	bmagentclient "github.com/cloudfoundry/bosh-micro-cli/deployment/agentclient"
)

type FakeAgentClientFactory struct {
	CreateAgentClient bmagentclient.AgentClient
	CreateMbusURL     string
}

func NewFakeAgentClientFactory() *FakeAgentClientFactory {
	return &FakeAgentClientFactory{}
}

func (f *FakeAgentClientFactory) Create(mbusURL string) bmagentclient.AgentClient {
	f.CreateMbusURL = mbusURL
	return f.CreateAgentClient
}