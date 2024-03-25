package nexus

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

func DeleteAllPreconfiguredRepositories(c *nexus.NexusClient) error {
	repos, err := c.Repository.List()
	if err != nil {
		return nil
	}
	for _, rep := range repos {
		common.DeleteRepository(c.Script.Client, rep.Name)
	}
	return nil
}

func CreateRepositoryDockerHosted(c *nexus.NexusClient, name string) {
	writePolicy := repository.StorageWritePolicyAllow
	httpPort := 8082
	rep := repository.DockerHostedRepository{
		Name:   name,
		Online: true,
		Docker: repository.Docker{
			HTTPPort:       &httpPort,
			ForceBasicAuth: false,
			V1Enabled:      false,
		},
		Storage: repository.HostedStorage{
			BlobStoreName:               name,
			WritePolicy:                 &writePolicy,
			StrictContentTypeValidation: true,
		},
	}
	c.Repository.Docker.Hosted.Create(rep)
}

func CreateRepositoryDockerProxy(c *nexus.NexusClient, name string) {
	rep := repository.DockerProxyRepository{
		Name:   name,
		Online: true,
		Docker: repository.Docker{
			ForceBasicAuth: false,
			V1Enabled:      false,
		},
		Proxy: repository.Proxy{
			RemoteURL:      "https://registry-1.docker.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		DockerProxy: repository.DockerProxy{
			IndexType: repository.DockerProxyIndexTypeHub,
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		HTTPClient: repository.HTTPClient{
			Authentication: &repository.HTTPClientAuthentication{
				Type:     repository.HTTPClientAuthenticationTypeUsername,
				Username: hubDockerComUsername,
				Password: hubDockerComPassword,
			},
			AutoBlock: true,
		},
	}
	c.Repository.Docker.Proxy.Create(rep)
}

func CreateRepositoryDockerProxyMcrMicrosoftCom(c *nexus.NexusClient, name string) {
	rep := repository.DockerProxyRepository{
		Name:   name,
		Online: true,
		Docker: repository.Docker{
			ForceBasicAuth: false,
			V1Enabled:      false,
		},
		Proxy: repository.Proxy{
			RemoteURL:      "https://mcr.microsoft.com",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		DockerProxy: repository.DockerProxy{
			IndexType: repository.DockerProxyIndexTypeRegistry,
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		HTTPClient: repository.HTTPClient{
			AutoBlock: true,
		},
	}
	c.Repository.Docker.Proxy.Create(rep)
}

func CreateRepositoryDockerGroup(c *nexus.NexusClient, name string) {
	httpPort := 8083
	rep := repository.DockerGroupRepository{
		Name:   name,
		Online: true,
		Docker: repository.Docker{
			ForceBasicAuth: false,
			HTTPPort:       &httpPort,
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		Group: repository.GroupDeploy{
			MemberNames: []string{"docker-hosted", "docker-proxy-mcr-microsoft-com", "docker-proxy"},
		},
	}
	c.Repository.Docker.Group.Create(rep)
}

func CreateRepositoryNpmProxy(c *nexus.NexusClient, name string) {
	rep := repository.NpmProxyRepository{
		Name:   name,
		Online: true,
		Proxy: repository.Proxy{
			RemoteURL: "https://registry.npmjs.org",
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		HTTPClient: repository.HTTPClient{
			Authentication: &repository.HTTPClientAuthentication{
				Type:     repository.HTTPClientAuthenticationTypeUsername,
				Username: npmjsComUsername,
				Password: npmjsComPassword,
			},
			AutoBlock: true,
		},
	}
	c.Repository.Npm.Proxy.Create(rep)
}

func CreateRepositoryNugetHosted(c *nexus.NexusClient, name string) {
	writePolicy := repository.StorageWritePolicyAllow
	rep := repository.NugetHostedRepository{
		Name:   name,
		Online: true,
		Storage: repository.HostedStorage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
	}
	c.Repository.Nuget.Hosted.Create(rep)
}

func CreateRepositoryNugetProxyDevExpressCom(c *nexus.NexusClient, name string) {
	rep := repository.NugetProxyRepository{
		Name:   name,
		Online: true,
		NugetProxy: repository.NugetProxy{
			NugetVersion:         repository.NugetVersion3,
			QueryCacheItemMaxAge: 3600,
		},
		Proxy: repository.Proxy{
			RemoteURL:      devExpressNugetV3Url,
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		HTTPClient: repository.HTTPClient{
			AutoBlock: true,
		},
	}
	c.Repository.Nuget.Proxy.Create(rep)
}

func CreateRepositoryNugetProxyNugetOrg(c *nexus.NexusClient, name string) {
	rep := repository.NugetProxyRepository{
		Name:   name,
		Online: true,
		NugetProxy: repository.NugetProxy{
			NugetVersion:         repository.NugetVersion3,
			QueryCacheItemMaxAge: 3600,
		},
		Proxy: repository.Proxy{
			RemoteURL:      "https://api.nuget.org/v3/index.json",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		HTTPClient: repository.HTTPClient{
			AutoBlock: true,
		},
	}
	c.Repository.Nuget.Proxy.Create(rep)
}

func CreateRepositoryNugetGroup(c *nexus.NexusClient, name string) {
	rep := repository.NugetGroupRepository{
		Name:   name,
		Online: true,
		Storage: repository.Storage{
			BlobStoreName:               name,
			StrictContentTypeValidation: true,
		},
		Group: repository.Group{
			MemberNames: []string{"nuget-hosted", "nuget.devexpress.com-proxy", "nuget.org-proxy"},
		},
	}
	c.Repository.Nuget.Group.Create(rep)
}
