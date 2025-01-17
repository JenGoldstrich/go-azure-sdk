package vcenter

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterClient struct {
	Client *resourcemanager.Client
}

func NewVCenterClientWithBaseURI(api environments.Api) (*VCenterClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vcenter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VCenterClient: %+v", err)
	}

	return &VCenterClient{
		Client: client,
	}, nil
}
