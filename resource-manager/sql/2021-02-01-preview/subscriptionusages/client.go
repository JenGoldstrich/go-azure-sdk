package subscriptionusages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionUsagesClient struct {
	Client *resourcemanager.Client
}

func NewSubscriptionUsagesClientWithBaseURI(api environments.Api) (*SubscriptionUsagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "subscriptionusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubscriptionUsagesClient: %+v", err)
	}

	return &SubscriptionUsagesClient{
		Client: client,
	}, nil
}
