package blobauditing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedServerBlobAuditingPoliciesListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExtendedServerBlobAuditingPolicy
}

type ExtendedServerBlobAuditingPoliciesListByServerCompleteResult struct {
	Items []ExtendedServerBlobAuditingPolicy
}

// ExtendedServerBlobAuditingPoliciesListByServer ...
func (c BlobAuditingClient) ExtendedServerBlobAuditingPoliciesListByServer(ctx context.Context, id ServerId) (result ExtendedServerBlobAuditingPoliciesListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extendedAuditingSettings", id.ID()),
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
		Values *[]ExtendedServerBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExtendedServerBlobAuditingPoliciesListByServerComplete retrieves all the results into a single object
func (c BlobAuditingClient) ExtendedServerBlobAuditingPoliciesListByServerComplete(ctx context.Context, id ServerId) (ExtendedServerBlobAuditingPoliciesListByServerCompleteResult, error) {
	return c.ExtendedServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate(ctx, id, ExtendedServerBlobAuditingPolicyOperationPredicate{})
}

// ExtendedServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BlobAuditingClient) ExtendedServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate(ctx context.Context, id ServerId, predicate ExtendedServerBlobAuditingPolicyOperationPredicate) (result ExtendedServerBlobAuditingPoliciesListByServerCompleteResult, err error) {
	items := make([]ExtendedServerBlobAuditingPolicy, 0)

	resp, err := c.ExtendedServerBlobAuditingPoliciesListByServer(ctx, id)
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

	result = ExtendedServerBlobAuditingPoliciesListByServerCompleteResult{
		Items: items,
	}
	return
}
