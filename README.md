# nexus-iac

Configure your Nexus repository with Infrastructure As Code (IAC)

All blobstores and repositories are created and configured by iac.
First, change your values in pkg/common.go

**This repository deletes all repositories and blobstores.**

Create blobstores for each repository.

Creating repositories are as follows;

* Docker Group
  * Docker Hosted
  * Docker Proxy - mcr.microsoft.com
  * Docker Proxy - docker.io
* NPM Proxy
* Nuget Group
  * Nuget Hosted
  * Nuget Proxy - nuget.devexpress.com
  * Nuget Proxy - nuget.org

This project uses [Golang Nexus Client](https://pkg.go.dev/github.com/datadrivers/go-nexus-client)
