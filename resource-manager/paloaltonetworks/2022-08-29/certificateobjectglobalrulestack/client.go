package certificateobjectglobalrulestack

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateObjectGlobalRulestackClient struct {
	Client *resourcemanager.Client
}

func NewCertificateObjectGlobalRulestackClientWithBaseURI(api environments.Api) (*CertificateObjectGlobalRulestackClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "certificateobjectglobalrulestack", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateObjectGlobalRulestackClient: %+v", err)
	}

	return &CertificateObjectGlobalRulestackClient{
		Client: client,
	}, nil
}
