package provideroperationsmetadata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderOperationsMetadataClient struct {
	Client *resourcemanager.Client
}

func NewProviderOperationsMetadataClientWithBaseURI(api environments.Api) (*ProviderOperationsMetadataClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "provideroperationsmetadata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProviderOperationsMetadataClient: %+v", err)
	}

	return &ProviderOperationsMetadataClient{
		Client: client,
	}, nil
}
