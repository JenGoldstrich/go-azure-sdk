package appplatform

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppPlatformClient struct {
	Client *resourcemanager.Client
}

func NewAppPlatformClientWithBaseURI(api environments.Api) (*AppPlatformClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "appplatform", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppPlatformClient: %+v", err)
	}

	return &AppPlatformClient{
		Client: client,
	}, nil
}
