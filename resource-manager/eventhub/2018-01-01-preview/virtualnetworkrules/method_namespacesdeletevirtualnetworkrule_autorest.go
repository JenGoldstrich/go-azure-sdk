package virtualnetworkrules

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesDeleteVirtualNetworkRuleOperationResponse struct {
	HttpResponse *http.Response
}

// NamespacesDeleteVirtualNetworkRule ...
func (c VirtualNetworkRulesClient) NamespacesDeleteVirtualNetworkRule(ctx context.Context, id VirtualnetworkruleId) (result NamespacesDeleteVirtualNetworkRuleOperationResponse, err error) {
	req, err := c.preparerForNamespacesDeleteVirtualNetworkRule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualnetworkrules.VirtualNetworkRulesClient", "NamespacesDeleteVirtualNetworkRule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualnetworkrules.VirtualNetworkRulesClient", "NamespacesDeleteVirtualNetworkRule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNamespacesDeleteVirtualNetworkRule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualnetworkrules.VirtualNetworkRulesClient", "NamespacesDeleteVirtualNetworkRule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNamespacesDeleteVirtualNetworkRule prepares the NamespacesDeleteVirtualNetworkRule request.
func (c VirtualNetworkRulesClient) preparerForNamespacesDeleteVirtualNetworkRule(ctx context.Context, id VirtualnetworkruleId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNamespacesDeleteVirtualNetworkRule handles the response to the NamespacesDeleteVirtualNetworkRule request. The method always
// closes the http.Response Body.
func (c VirtualNetworkRulesClient) responderForNamespacesDeleteVirtualNetworkRule(resp *http.Response) (result NamespacesDeleteVirtualNetworkRuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
