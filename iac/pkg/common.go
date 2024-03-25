package nexus

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/go-nexus-client/nexus3/schema"
)

func EnableRequiredRealms(c *nexus.NexusClient) error {
	yetkiler := []string{
		"DockerToken",
		"NexusAuthenticatingRealm",
		"NpmToken",
		"NuGetApiKey",
	}
	return c.Security.Realm.Activate(yetkiler)
}

func SetupMailServer(c *nexus.NexusClient) error {
	boolptr := true
	userName := mailUsername
	passWord := mailPassword
	subjectPrefix := "[NEXUS]"
	config := schema.MailConfig{
		Enabled:             &boolptr,
		Host:                mailHost,
		Port:                465,
		Username:            &userName,
		Password:            &passWord,
		FromAddress:         mailFromAddress,
		SubjectPrefix:       &subjectPrefix,
		StartTlsRequired:    &boolptr,
		SslOnConnectEnabled: &boolptr,
	}
	return c.MailConfig.Update(&config)
}
