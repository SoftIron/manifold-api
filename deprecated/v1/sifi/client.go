package sifi

import (
	"github.com/softiron/manifold-api/client"
)

// Client is a connection to the SIFI service.
type Client struct {
	ACLService
	ClusterService
	DatastoreService
	DocumentService
	GroupService
	HookService
	HostService
	ImageService
	InstanceService
	MarketService
	SecurityGroupService
	SystemService
	TemplateService
	UserService
	VNetService
	VRouterService
	VirtualDataCenterService
	ZoneService
	*client.Client
}

// NewClient returns a new client.
func NewClient(o *client.Options) *Client {
	c := client.New(o)

	return &Client{
		Client:                   c,
		ACLService:               ACLService{Client: c},
		ClusterService:           ClusterService{Client: c},
		DatastoreService:         DatastoreService{Client: c},
		DocumentService:          DocumentService{Client: c},
		GroupService:             GroupService{Client: c},
		HookService:              HookService{Client: c},
		HostService:              HostService{Client: c},
		ImageService:             ImageService{Client: c},
		InstanceService:          InstanceService{Client: c},
		MarketService:            MarketService{Client: c},
		SecurityGroupService:     SecurityGroupService{Client: c},
		SystemService:            SystemService{Client: c},
		TemplateService:          TemplateService{Client: c},
		UserService:              UserService{Client: c},
		VNetService:              VNetService{Client: c},
		VRouterService:           VRouterService{Client: c},
		VirtualDataCenterService: VirtualDataCenterService{Client: c},
		ZoneService:              ZoneService{Client: c},
	}
}
