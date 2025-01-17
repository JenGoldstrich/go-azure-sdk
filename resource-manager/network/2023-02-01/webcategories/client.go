package webcategories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebCategoriesClient struct {
	Client *resourcemanager.Client
}

func NewWebCategoriesClientWithBaseURI(api environments.Api) (*WebCategoriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "webcategories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebCategoriesClient: %+v", err)
	}

	return &WebCategoriesClient{
		Client: client,
	}, nil
}
