package syncsets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncSetsClient struct {
	Client *resourcemanager.Client
}

func NewSyncSetsClientWithBaseURI(api environments.Api) (*SyncSetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "syncsets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SyncSetsClient: %+v", err)
	}

	return &SyncSetsClient{
		Client: client,
	}, nil
}
