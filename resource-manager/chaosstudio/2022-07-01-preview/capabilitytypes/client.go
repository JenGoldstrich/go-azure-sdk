package capabilitytypes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityTypesClient struct {
	Client *resourcemanager.Client
}

func NewCapabilityTypesClientWithBaseURI(api environments.Api) (*CapabilityTypesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "capabilitytypes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CapabilityTypesClient: %+v", err)
	}

	return &CapabilityTypesClient{
		Client: client,
	}, nil
}
