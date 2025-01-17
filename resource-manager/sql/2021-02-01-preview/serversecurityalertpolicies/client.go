package serversecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewServerSecurityAlertPoliciesClientWithBaseURI(api environments.Api) (*ServerSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serversecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerSecurityAlertPoliciesClient: %+v", err)
	}

	return &ServerSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
