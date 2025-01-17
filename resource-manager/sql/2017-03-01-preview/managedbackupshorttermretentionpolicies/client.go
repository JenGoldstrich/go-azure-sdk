package managedbackupshorttermretentionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedBackupShortTermRetentionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewManagedBackupShortTermRetentionPoliciesClientWithBaseURI(api environments.Api) (*ManagedBackupShortTermRetentionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedbackupshorttermretentionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedBackupShortTermRetentionPoliciesClient: %+v", err)
	}

	return &ManagedBackupShortTermRetentionPoliciesClient{
		Client: client,
	}, nil
}
