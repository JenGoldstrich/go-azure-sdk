package managedinstanceprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstancePrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstancePrivateLinkResourcesClientWithBaseURI(api environments.Api) (*ManagedInstancePrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedinstanceprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstancePrivateLinkResourcesClient: %+v", err)
	}

	return &ManagedInstancePrivateLinkResourcesClient{
		Client: client,
	}, nil
}
