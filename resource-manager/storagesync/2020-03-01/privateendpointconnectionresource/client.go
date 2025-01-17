package privateendpointconnectionresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionResourceClient struct {
	Client *resourcemanager.Client
}

func NewPrivateEndpointConnectionResourceClientWithBaseURI(api environments.Api) (*PrivateEndpointConnectionResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privateendpointconnectionresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateEndpointConnectionResourceClient: %+v", err)
	}

	return &PrivateEndpointConnectionResourceClient{
		Client: client,
	}, nil
}
