// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/merge-api/merge-go-client/core"
	accountdetails "github.com/merge-api/merge-go-client/ticketing/accountdetails"
	accounts "github.com/merge-api/merge-go-client/ticketing/accounts"
	accounttoken "github.com/merge-api/merge-go-client/ticketing/accounttoken"
	asyncpassthrough "github.com/merge-api/merge-go-client/ticketing/asyncpassthrough"
	attachments "github.com/merge-api/merge-go-client/ticketing/attachments"
	audittrail "github.com/merge-api/merge-go-client/ticketing/audittrail"
	availableactions "github.com/merge-api/merge-go-client/ticketing/availableactions"
	collections "github.com/merge-api/merge-go-client/ticketing/collections"
	comments "github.com/merge-api/merge-go-client/ticketing/comments"
	contacts "github.com/merge-api/merge-go-client/ticketing/contacts"
	deleteaccount "github.com/merge-api/merge-go-client/ticketing/deleteaccount"
	forceresync "github.com/merge-api/merge-go-client/ticketing/forceresync"
	generatekey "github.com/merge-api/merge-go-client/ticketing/generatekey"
	issues "github.com/merge-api/merge-go-client/ticketing/issues"
	linkedaccounts "github.com/merge-api/merge-go-client/ticketing/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/ticketing/linktoken"
	passthrough "github.com/merge-api/merge-go-client/ticketing/passthrough"
	projects "github.com/merge-api/merge-go-client/ticketing/projects"
	regeneratekey "github.com/merge-api/merge-go-client/ticketing/regeneratekey"
	selectivesync "github.com/merge-api/merge-go-client/ticketing/selectivesync"
	syncstatus "github.com/merge-api/merge-go-client/ticketing/syncstatus"
	tags "github.com/merge-api/merge-go-client/ticketing/tags"
	teams "github.com/merge-api/merge-go-client/ticketing/teams"
	tickets "github.com/merge-api/merge-go-client/ticketing/tickets"
	users "github.com/merge-api/merge-go-client/ticketing/users"
	webhookreceivers "github.com/merge-api/merge-go-client/ticketing/webhookreceivers"
	http "net/http"
)

type Client interface {
	AccountDetails() accountdetails.Client
	AccountToken() accounttoken.Client
	Accounts() accounts.Client
	AsyncPassthrough() asyncpassthrough.Client
	Attachments() attachments.Client
	AuditTrail() audittrail.Client
	AvailableActions() availableactions.Client
	Collections() collections.Client
	Comments() comments.Client
	Contacts() contacts.Client
	DeleteAccount() deleteaccount.Client
	GenerateKey() generatekey.Client
	Issues() issues.Client
	LinkToken() linktoken.Client
	LinkedAccounts() linkedaccounts.Client
	Passthrough() passthrough.Client
	Projects() projects.Client
	RegenerateKey() regeneratekey.Client
	SelectiveSync() selectivesync.Client
	SyncStatus() syncstatus.Client
	ForceResync() forceresync.Client
	Tags() tags.Client
	Teams() teams.Client
	Tickets() tickets.Client
	Users() users.Client
	WebhookReceivers() webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:                options.BaseURL,
		httpClient:             options.HTTPClient,
		header:                 options.ToHeader(),
		accountDetailsClient:   accountdetails.NewClient(opts...),
		accountTokenClient:     accounttoken.NewClient(opts...),
		accountsClient:         accounts.NewClient(opts...),
		asyncPassthroughClient: asyncpassthrough.NewClient(opts...),
		attachmentsClient:      attachments.NewClient(opts...),
		auditTrailClient:       audittrail.NewClient(opts...),
		availableActionsClient: availableactions.NewClient(opts...),
		collectionsClient:      collections.NewClient(opts...),
		commentsClient:         comments.NewClient(opts...),
		contactsClient:         contacts.NewClient(opts...),
		deleteAccountClient:    deleteaccount.NewClient(opts...),
		generateKeyClient:      generatekey.NewClient(opts...),
		issuesClient:           issues.NewClient(opts...),
		linkTokenClient:        linktoken.NewClient(opts...),
		linkedAccountsClient:   linkedaccounts.NewClient(opts...),
		passthroughClient:      passthrough.NewClient(opts...),
		projectsClient:         projects.NewClient(opts...),
		regenerateKeyClient:    regeneratekey.NewClient(opts...),
		selectiveSyncClient:    selectivesync.NewClient(opts...),
		syncStatusClient:       syncstatus.NewClient(opts...),
		forceResyncClient:      forceresync.NewClient(opts...),
		tagsClient:             tags.NewClient(opts...),
		teamsClient:            teams.NewClient(opts...),
		ticketsClient:          tickets.NewClient(opts...),
		usersClient:            users.NewClient(opts...),
		webhookReceiversClient: webhookreceivers.NewClient(opts...),
	}
}

type client struct {
	baseURL                string
	httpClient             core.HTTPClient
	header                 http.Header
	accountDetailsClient   accountdetails.Client
	accountTokenClient     accounttoken.Client
	accountsClient         accounts.Client
	asyncPassthroughClient asyncpassthrough.Client
	attachmentsClient      attachments.Client
	auditTrailClient       audittrail.Client
	availableActionsClient availableactions.Client
	collectionsClient      collections.Client
	commentsClient         comments.Client
	contactsClient         contacts.Client
	deleteAccountClient    deleteaccount.Client
	generateKeyClient      generatekey.Client
	issuesClient           issues.Client
	linkTokenClient        linktoken.Client
	linkedAccountsClient   linkedaccounts.Client
	passthroughClient      passthrough.Client
	projectsClient         projects.Client
	regenerateKeyClient    regeneratekey.Client
	selectiveSyncClient    selectivesync.Client
	syncStatusClient       syncstatus.Client
	forceResyncClient      forceresync.Client
	tagsClient             tags.Client
	teamsClient            teams.Client
	ticketsClient          tickets.Client
	usersClient            users.Client
	webhookReceiversClient webhookreceivers.Client
}

func (c *client) AccountDetails() accountdetails.Client {
	return c.accountDetailsClient
}

func (c *client) AccountToken() accounttoken.Client {
	return c.accountTokenClient
}

func (c *client) Accounts() accounts.Client {
	return c.accountsClient
}

func (c *client) AsyncPassthrough() asyncpassthrough.Client {
	return c.asyncPassthroughClient
}

func (c *client) Attachments() attachments.Client {
	return c.attachmentsClient
}

func (c *client) AuditTrail() audittrail.Client {
	return c.auditTrailClient
}

func (c *client) AvailableActions() availableactions.Client {
	return c.availableActionsClient
}

func (c *client) Collections() collections.Client {
	return c.collectionsClient
}

func (c *client) Comments() comments.Client {
	return c.commentsClient
}

func (c *client) Contacts() contacts.Client {
	return c.contactsClient
}

func (c *client) DeleteAccount() deleteaccount.Client {
	return c.deleteAccountClient
}

func (c *client) GenerateKey() generatekey.Client {
	return c.generateKeyClient
}

func (c *client) Issues() issues.Client {
	return c.issuesClient
}

func (c *client) LinkToken() linktoken.Client {
	return c.linkTokenClient
}

func (c *client) LinkedAccounts() linkedaccounts.Client {
	return c.linkedAccountsClient
}

func (c *client) Passthrough() passthrough.Client {
	return c.passthroughClient
}

func (c *client) Projects() projects.Client {
	return c.projectsClient
}

func (c *client) RegenerateKey() regeneratekey.Client {
	return c.regenerateKeyClient
}

func (c *client) SelectiveSync() selectivesync.Client {
	return c.selectiveSyncClient
}

func (c *client) SyncStatus() syncstatus.Client {
	return c.syncStatusClient
}

func (c *client) ForceResync() forceresync.Client {
	return c.forceResyncClient
}

func (c *client) Tags() tags.Client {
	return c.tagsClient
}

func (c *client) Teams() teams.Client {
	return c.teamsClient
}

func (c *client) Tickets() tickets.Client {
	return c.ticketsClient
}

func (c *client) Users() users.Client {
	return c.usersClient
}

func (c *client) WebhookReceivers() webhookreceivers.Client {
	return c.webhookReceiversClient
}
