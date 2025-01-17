package instancepools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstancePoolsClient struct {
	Client *resourcemanager.Client
}

func NewInstancePoolsClientWithBaseURI(api environments.Api) (*InstancePoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "instancepools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InstancePoolsClient: %+v", err)
	}

	return &InstancePoolsClient{
		Client: client,
	}, nil
}
