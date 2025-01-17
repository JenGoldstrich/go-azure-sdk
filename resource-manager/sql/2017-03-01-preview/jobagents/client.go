package jobagents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobAgentsClient struct {
	Client *resourcemanager.Client
}

func NewJobAgentsClientWithBaseURI(api environments.Api) (*JobAgentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobagents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobAgentsClient: %+v", err)
	}

	return &JobAgentsClient{
		Client: client,
	}, nil
}
