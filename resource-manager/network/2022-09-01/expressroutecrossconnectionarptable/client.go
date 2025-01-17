package expressroutecrossconnectionarptable

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCrossConnectionArpTableClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCrossConnectionArpTableClientWithBaseURI(api environments.Api) (*ExpressRouteCrossConnectionArpTableClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecrossconnectionarptable", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCrossConnectionArpTableClient: %+v", err)
	}

	return &ExpressRouteCrossConnectionArpTableClient{
		Client: client,
	}, nil
}
