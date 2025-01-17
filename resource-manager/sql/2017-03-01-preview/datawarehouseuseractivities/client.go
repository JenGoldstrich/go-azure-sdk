package datawarehouseuseractivities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataWarehouseUserActivitiesClient struct {
	Client *resourcemanager.Client
}

func NewDataWarehouseUserActivitiesClientWithBaseURI(api environments.Api) (*DataWarehouseUserActivitiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "datawarehouseuseractivities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataWarehouseUserActivitiesClient: %+v", err)
	}

	return &DataWarehouseUserActivitiesClient{
		Client: client,
	}, nil
}
