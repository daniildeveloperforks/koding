package openstack

import (
	"crypto/tls"
	"net/http"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type Config struct {
	Username         string
	UserID           string
	Password         string
	APIKey           string
	IdentityEndpoint string
	TenantID         string
	TenantName       string
	DomainID         string
	DomainName       string
	Insecure         bool

	osClient *gophercloud.ProviderClient
}

func (c *Config) loadAndValidate() error {
	ao := gophercloud.AuthOptions{
		Username:         c.Username,
		UserID:           c.UserID,
		Password:         c.Password,
		APIKey:           c.APIKey,
		IdentityEndpoint: c.IdentityEndpoint,
		TenantID:         c.TenantID,
		TenantName:       c.TenantName,
		DomainID:         c.DomainID,
		DomainName:       c.DomainName,
	}

	client, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		return err
	}

	if c.Insecure {
		// Configure custom TLS settings.
		config := &tls.Config{InsecureSkipVerify: true}
		transport := &http.Transport{TLSClientConfig: config}
		client.HTTPClient.Transport = transport
	}

	err = openstack.Authenticate(client, ao)
	if err != nil {
		return err
	}

	c.osClient = client

	return nil
}

func (c *Config) blockStorageV1Client(region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV1(c.osClient, gophercloud.EndpointOpts{
		Region: region,
	})
}

func (c *Config) computeV2Client(region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(c.osClient, gophercloud.EndpointOpts{
		Region: region,
	})
}

func (c *Config) networkingV2Client(region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(c.osClient, gophercloud.EndpointOpts{
		Region: region,
	})
}

func (c *Config) objectStorageV1Client(region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewObjectStorageV1(c.osClient, gophercloud.EndpointOpts{
		Region: region,
	})
}
