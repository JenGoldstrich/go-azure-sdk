package devicecapacityinfo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityInfoClient struct {
	Client *resourcemanager.Client
}

func NewDeviceCapacityInfoClientWithBaseURI(api environments.Api) (*DeviceCapacityInfoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "devicecapacityinfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCapacityInfoClient: %+v", err)
	}

	return &DeviceCapacityInfoClient{
		Client: client,
	}, nil
}
