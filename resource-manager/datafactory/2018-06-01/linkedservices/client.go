package linkedservices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedServicesClient struct {
	Client *resourcemanager.Client
}

func NewLinkedServicesClientWithBaseURI(api environments.Api) (*LinkedServicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "linkedservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LinkedServicesClient: %+v", err)
	}

	return &LinkedServicesClient{
		Client: client,
	}, nil
}
