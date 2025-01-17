package firewalls

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSupportInfoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SupportInfo
}

type GetSupportInfoOperationOptions struct {
	Email *string
}

func DefaultGetSupportInfoOperationOptions() GetSupportInfoOperationOptions {
	return GetSupportInfoOperationOptions{}
}

func (o GetSupportInfoOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSupportInfoOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetSupportInfoOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Email != nil {
		out.Append("email", fmt.Sprintf("%v", *o.Email))
	}
	return &out
}

// GetSupportInfo ...
func (c FirewallsClient) GetSupportInfo(ctx context.Context, id FirewallId, options GetSupportInfoOperationOptions) (result GetSupportInfoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/getSupportInfo", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
