// This file was auto-generated by Fern from our API Definition.

package deleteaccount

import (
	context "context"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client interface {
	Create(ctx context.Context) error
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Delete a linked account.
func (c *client) Create(ctx context.Context) error {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/ticketing/v1/delete-account"

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		nil,
		false,
		c.header,
		nil,
	); err != nil {
		return err
	}
	return nil
}
