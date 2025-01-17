package migrationconfigs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndStartMigrationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateAndStartMigration ...
func (c MigrationConfigsClient) CreateAndStartMigration(ctx context.Context, id NamespaceId, input MigrationConfigProperties) (result CreateAndStartMigrationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/migrationConfigurations/$default", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// CreateAndStartMigrationThenPoll performs CreateAndStartMigration then polls until it's completed
func (c MigrationConfigsClient) CreateAndStartMigrationThenPoll(ctx context.Context, id NamespaceId, input MigrationConfigProperties) error {
	result, err := c.CreateAndStartMigration(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateAndStartMigration: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CreateAndStartMigration: %+v", err)
	}

	return nil
}
