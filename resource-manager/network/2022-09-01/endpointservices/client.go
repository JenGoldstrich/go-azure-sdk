package endpointservices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointServicesClient struct {
	Client *resourcemanager.Client
}

func NewEndpointServicesClientWithBaseURI(api environments.Api) (*EndpointServicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "endpointservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EndpointServicesClient: %+v", err)
	}

	return &EndpointServicesClient{
		Client: client,
	}, nil
}
