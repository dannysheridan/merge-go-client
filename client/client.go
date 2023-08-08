// This file was auto-generated by Fern from our API Definition.

package client

import (
	accountingclient "github.com/merge-api/merge-go-client/accounting/client"
	atsclient "github.com/merge-api/merge-go-client/ats/client"
	core "github.com/merge-api/merge-go-client/core"
	crmclient "github.com/merge-api/merge-go-client/crm/client"
	filestorageclient "github.com/merge-api/merge-go-client/filestorage/client"
	hrisclient "github.com/merge-api/merge-go-client/hris/client"
	ticketingclient "github.com/merge-api/merge-go-client/ticketing/client"
	http "net/http"
)

type Client interface {
	Ats() atsclient.Client
	Crm() crmclient.Client
	Filestorage() filestorageclient.Client
	Hris() hrisclient.Client
	Ticketing() ticketingclient.Client
	Accounting() accountingclient.Client
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:           options.BaseURL,
		httpClient:        options.HTTPClient,
		header:            options.ToHeader(),
		atsClient:         atsclient.NewClient(opts...),
		crmClient:         crmclient.NewClient(opts...),
		filestorageClient: filestorageclient.NewClient(opts...),
		hrisClient:        hrisclient.NewClient(opts...),
		ticketingClient:   ticketingclient.NewClient(opts...),
		accountingClient:  accountingclient.NewClient(opts...),
	}
}

type client struct {
	baseURL           string
	httpClient        core.HTTPClient
	header            http.Header
	atsClient         atsclient.Client
	crmClient         crmclient.Client
	filestorageClient filestorageclient.Client
	hrisClient        hrisclient.Client
	ticketingClient   ticketingclient.Client
	accountingClient  accountingclient.Client
}

func (c *client) Ats() atsclient.Client {
	return c.atsClient
}

func (c *client) Crm() crmclient.Client {
	return c.crmClient
}

func (c *client) Filestorage() filestorageclient.Client {
	return c.filestorageClient
}

func (c *client) Hris() hrisclient.Client {
	return c.hrisClient
}

func (c *client) Ticketing() ticketingclient.Client {
	return c.ticketingClient
}

func (c *client) Accounting() accountingclient.Client {
	return c.accountingClient
}
