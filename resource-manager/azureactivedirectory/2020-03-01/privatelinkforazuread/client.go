package privatelinkforazuread

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkForAzureAdClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkForAzureAdClientWithBaseURI(api environments.Api) (*PrivateLinkForAzureAdClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privatelinkforazuread", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkForAzureAdClient: %+v", err)
	}

	return &PrivateLinkForAzureAdClient{
		Client: client,
	}, nil
}
