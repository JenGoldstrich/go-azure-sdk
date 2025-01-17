package workflowtriggerhistories

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowTriggerHistoriesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkflowTriggerHistoriesClientWithBaseURI(endpoint string) WorkflowTriggerHistoriesClient {
	return WorkflowTriggerHistoriesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
