// This file was auto-generated by Fern from our API Definition.

package contacts

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	crm "github.com/merge-api/merge-go-client/crm"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	List(ctx context.Context, request *crm.ContactsListRequest) (*crm.PaginatedContactList, error)
	Create(ctx context.Context, request *crm.CrmContactEndpointRequest) (*crm.CrmContactResponse, error)
	Retrieve(ctx context.Context, id string, request *crm.ContactsRetrieveRequest) (*crm.Contact, error)
	PartialUpdate(ctx context.Context, id string, request *crm.PatchedCrmContactEndpointRequest) (*crm.CrmContactResponse, error)
	IgnoreCreate(ctx context.Context, modelId string, request *crm.IgnoreCommonModelRequest) error
	MetaPatchRetrieve(ctx context.Context, id string) (*crm.MetaResponse, error)
	MetaPostRetrieve(ctx context.Context) (*crm.MetaResponse, error)
	RemoteFieldClassesList(ctx context.Context, request *crm.ContactsRemoteFieldClassesListRequest) (*crm.PaginatedRemoteFieldClassList, error)
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

// Returns a list of `Contact` objects.
func (c *client) List(ctx context.Context, request *crm.ContactsListRequest) (*crm.PaginatedContactList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/contacts"

	queryParams := make(url.Values)
	if request.AccountId != nil {
		queryParams.Add("account_id", fmt.Sprintf("%v", *request.AccountId))
	}
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.PaginatedContactList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Creates a `Contact` object with the given values.
func (c *client) Create(ctx context.Context, request *crm.CrmContactEndpointRequest) (*crm.CrmContactResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/contacts"

	queryParams := make(url.Values)
	if request.IsDebugMode != nil {
		queryParams.Add("is_debug_mode", fmt.Sprintf("%v", *request.IsDebugMode))
	}
	if request.RunAsync != nil {
		queryParams.Add("run_async", fmt.Sprintf("%v", *request.RunAsync))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.CrmContactResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns a `Contact` object with the given `id`.
func (c *client) Retrieve(ctx context.Context, id string, request *crm.ContactsRetrieveRequest) (*crm.Contact, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/contacts/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.Contact
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Updates a `Contact` object with the given `id`.
func (c *client) PartialUpdate(ctx context.Context, id string, request *crm.PatchedCrmContactEndpointRequest) (*crm.CrmContactResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/contacts/%v", id)

	queryParams := make(url.Values)
	if request.IsDebugMode != nil {
		queryParams.Add("is_debug_mode", fmt.Sprintf("%v", *request.IsDebugMode))
	}
	if request.RunAsync != nil {
		queryParams.Add("run_async", fmt.Sprintf("%v", *request.RunAsync))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.CrmContactResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
func (c *client) IgnoreCreate(ctx context.Context, modelId string, request *crm.IgnoreCommonModelRequest) error {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/contacts/ignore/%v", modelId)

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		nil,
		false,
		c.header,
		nil,
	); err != nil {
		return err
	}
	return nil
}

// Returns metadata for `CRMContact` PATCHs.
func (c *client) MetaPatchRetrieve(ctx context.Context, id string) (*crm.MetaResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/contacts/meta/patch/%v", id)

	var response *crm.MetaResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns metadata for `CRMContact` POSTs.
func (c *client) MetaPostRetrieve(ctx context.Context) (*crm.MetaResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/contacts/meta/post"

	var response *crm.MetaResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns a list of `RemoteFieldClass` objects.
func (c *client) RemoteFieldClassesList(ctx context.Context, request *crm.ContactsRemoteFieldClassesListRequest) (*crm.PaginatedRemoteFieldClassList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/contacts/remote-field-classes"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.PaginatedRemoteFieldClassList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
