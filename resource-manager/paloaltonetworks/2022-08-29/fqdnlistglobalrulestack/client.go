package fqdnlistglobalrulestack

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FqdnListGlobalRulestackClient struct {
	Client *resourcemanager.Client
}

func NewFqdnListGlobalRulestackClientWithBaseURI(api environments.Api) (*FqdnListGlobalRulestackClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "fqdnlistglobalrulestack", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FqdnListGlobalRulestackClient: %+v", err)
	}

	return &FqdnListGlobalRulestackClient{
		Client: client,
	}, nil
}
