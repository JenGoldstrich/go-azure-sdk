package customlocations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEnabledResourceTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EnabledResourceType
}

type ListEnabledResourceTypesCompleteResult struct {
	Items []EnabledResourceType
}

// ListEnabledResourceTypes ...
func (c CustomLocationsClient) ListEnabledResourceTypes(ctx context.Context, id CustomLocationId) (result ListEnabledResourceTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/enabledResourceTypes", id.ID()),
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
		Values *[]EnabledResourceType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnabledResourceTypesComplete retrieves all the results into a single object
func (c CustomLocationsClient) ListEnabledResourceTypesComplete(ctx context.Context, id CustomLocationId) (ListEnabledResourceTypesCompleteResult, error) {
	return c.ListEnabledResourceTypesCompleteMatchingPredicate(ctx, id, EnabledResourceTypeOperationPredicate{})
}

// ListEnabledResourceTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomLocationsClient) ListEnabledResourceTypesCompleteMatchingPredicate(ctx context.Context, id CustomLocationId, predicate EnabledResourceTypeOperationPredicate) (result ListEnabledResourceTypesCompleteResult, err error) {
	items := make([]EnabledResourceType, 0)

	resp, err := c.ListEnabledResourceTypes(ctx, id)
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

	result = ListEnabledResourceTypesCompleteResult{
		Items: items,
	}
	return
}
