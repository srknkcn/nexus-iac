package nexus

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
)

func DeleteAllPreConfiguredBlobstores(c *nexus.NexusClient) error {
	blobs, err := c.BlobStore.List()
	if err != nil {
		return err
	}
	for _, bs := range blobs {
		DeleteBlobstore(c, bs.Name)
	}
	return nil
}

func CreateBlobstore(c *nexus.NexusClient, name string) {
	bs := blobstore.File{
		Name: name,
		Path: name,
	}
	c.BlobStore.File.Create(&bs)
}

func DeleteBlobstore(c *nexus.NexusClient, name string) {
	c.BlobStore.File.Delete(name)
}