package serverendpointresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointResourceClient struct {
	Client *resourcemanager.Client
}

func NewServerEndpointResourceClientWithBaseURI(api environments.Api) (*ServerEndpointResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverendpointresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerEndpointResourceClient: %+v", err)
	}

	return &ServerEndpointResourceClient{
		Client: client,
	}, nil
}
