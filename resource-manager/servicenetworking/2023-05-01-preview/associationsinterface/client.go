package associationsinterface

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociationsInterfaceClient struct {
	Client *resourcemanager.Client
}

func NewAssociationsInterfaceClientWithBaseURI(api environments.Api) (*AssociationsInterfaceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "associationsinterface", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssociationsInterfaceClient: %+v", err)
	}

	return &AssociationsInterfaceClient{
		Client: client,
	}, nil
}
