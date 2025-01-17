package actions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionsClient struct {
	Client *resourcemanager.Client
}

func NewActionsClientWithBaseURI(api environments.Api) (*ActionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "actions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActionsClient: %+v", err)
	}

	return &ActionsClient{
		Client: client,
	}, nil
}
