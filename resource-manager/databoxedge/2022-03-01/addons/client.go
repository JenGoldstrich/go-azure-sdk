package addons

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAddonsClientWithBaseURI(endpoint string) AddonsClient {
	return AddonsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
