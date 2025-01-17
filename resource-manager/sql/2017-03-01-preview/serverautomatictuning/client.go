package serverautomatictuning

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerAutomaticTuningClient struct {
	Client *resourcemanager.Client
}

func NewServerAutomaticTuningClientWithBaseURI(api environments.Api) (*ServerAutomaticTuningClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverautomatictuning", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerAutomaticTuningClient: %+v", err)
	}

	return &ServerAutomaticTuningClient{
		Client: client,
	}, nil
}
