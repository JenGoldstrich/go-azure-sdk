package applicationgatewayprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewApplicationGatewayPrivateLinkResourcesClientWithBaseURI(api environments.Api) (*ApplicationGatewayPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "applicationgatewayprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationGatewayPrivateLinkResourcesClient: %+v", err)
	}

	return &ApplicationGatewayPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
