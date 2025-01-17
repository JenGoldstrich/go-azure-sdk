package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ResourcesListByWorkspaceCompleteResult struct {
	Items []PrivateLinkResource
}

// ResourcesListByWorkspace ...
func (c PrivateLinkClient) ResourcesListByWorkspace(ctx context.Context, id WorkspaceId) (result ResourcesListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourcesListByWorkspaceComplete retrieves all the results into a single object
func (c PrivateLinkClient) ResourcesListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (ResourcesListByWorkspaceCompleteResult, error) {
	return c.ResourcesListByWorkspaceCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// ResourcesListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) ResourcesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate PrivateLinkResourceOperationPredicate) (result ResourcesListByWorkspaceCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ResourcesListByWorkspace(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ResourcesListByWorkspaceCompleteResult{
		Items: items,
	}
	return
}
