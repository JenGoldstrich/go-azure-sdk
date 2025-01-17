package firewallstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallStatusClient struct {
	Client *resourcemanager.Client
}

func NewFirewallStatusClientWithBaseURI(api environments.Api) (*FirewallStatusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "firewallstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FirewallStatusClient: %+v", err)
	}

	return &FirewallStatusClient{
		Client: client,
	}, nil
}
