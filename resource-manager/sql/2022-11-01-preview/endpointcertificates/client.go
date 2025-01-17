package endpointcertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewEndpointCertificatesClientWithBaseURI(api environments.Api) (*EndpointCertificatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "endpointcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EndpointCertificatesClient: %+v", err)
	}

	return &EndpointCertificatesClient{
		Client: client,
	}, nil
}
