package guestconfigurationassignmentreports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentReportsClient struct {
	Client *resourcemanager.Client
}

func NewGuestConfigurationAssignmentReportsClientWithBaseURI(api environments.Api) (*GuestConfigurationAssignmentReportsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "guestconfigurationassignmentreports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestConfigurationAssignmentReportsClient: %+v", err)
	}

	return &GuestConfigurationAssignmentReportsClient{
		Client: client,
	}, nil
}
