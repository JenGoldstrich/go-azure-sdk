package collectorpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CollectorPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewCollectorPoliciesClientWithBaseURI(api environments.Api) (*CollectorPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "collectorpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CollectorPoliciesClient: %+v", err)
	}

	return &CollectorPoliciesClient{
		Client: client,
	}, nil
}
