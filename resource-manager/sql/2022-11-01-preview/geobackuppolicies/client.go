package geobackuppolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeoBackupPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewGeoBackupPoliciesClientWithBaseURI(api environments.Api) (*GeoBackupPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "geobackuppolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GeoBackupPoliciesClient: %+v", err)
	}

	return &GeoBackupPoliciesClient{
		Client: client,
	}, nil
}
