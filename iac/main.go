package main

import (
	"fmt"

	nexusClient "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	nexus "github.com/srknkcn/nexus-iac/pkg"
)

func main() {
	c := nexusClient.NewClient(getDefaultConfig())

	// Delete ALL!!! Preconfigured Repositories & Blobstores...
	if err := nexus.DeleteAllPreconfiguredRepositories(c); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := nexus.DeleteAllPreConfiguredBlobstores(c); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Enabled Required Realms...
	if err := nexus.EnableRequiredRealms(c); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Setup Mail Server...
	if err := nexus.SetupMailServer(c); err != nil {
		fmt.Println(err.Error())
		return
	}

	// DOCKER repositories...
	nexus.CreateBlobstore(c, "docker-hosted")
	nexus.CreateRepositoryDockerHosted(c, "docker-hosted")
	nexus.CreateBlobstore(c, "docker-proxy")
	nexus.CreateRepositoryDockerProxy(c, "docker-proxy")
	nexus.CreateBlobstore(c, "docker-proxy-mcr-microsoft-com")
	nexus.CreateRepositoryDockerProxyMcrMicrosoftCom(c, "docker-proxy-mcr-microsoft-com")
	nexus.CreateBlobstore(c, "docker-group")
	nexus.CreateRepositoryDockerGroup(c, "docker-group")

	// NPM repositories
	nexus.CreateBlobstore(c, "npm-proxy")
	nexus.CreateRepositoryNpmProxy(c, "npm-proxy")

	// NUGET repositories...
	nexus.CreateBlobstore(c, "nuget-hosted")
	nexus.CreateRepositoryNugetHosted(c, "nuget-hosted")
	nexus.CreateBlobstore(c, "nuget.devexpress.com-proxy")
	nexus.CreateRepositoryNugetProxyDevExpressCom(c, "nuget.devexpress.com-proxy")
	nexus.CreateBlobstore(c, "nuget.org-proxy")
	nexus.CreateRepositoryNugetProxyNugetOrg(c, "nuget.org-proxy")
	nexus.CreateBlobstore(c, "nuget-group")
	nexus.CreateRepositoryNugetGroup(c, "nuget-group")
}

func getDefaultConfig() client.Config {
	return client.Config{
		URL:      nexus.NexusURL,
		Username: nexus.NexusUsername,
		Password: nexus.NexusPassword,
		Insecure: true,
	}
}
