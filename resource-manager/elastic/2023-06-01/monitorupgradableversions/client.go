package monitorupgradableversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorUpgradableVersionsClient struct {
	Client *resourcemanager.Client
}

func NewMonitorUpgradableVersionsClientWithBaseURI(api environments.Api) (*MonitorUpgradableVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "monitorupgradableversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitorUpgradableVersionsClient: %+v", err)
	}

	return &MonitorUpgradableVersionsClient{
		Client: client,
	}, nil
}
