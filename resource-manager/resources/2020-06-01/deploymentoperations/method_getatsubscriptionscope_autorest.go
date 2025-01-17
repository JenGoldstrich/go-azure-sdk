package deploymentoperations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentOperation
}

// GetAtSubscriptionScope ...
func (c DeploymentOperationsClient) GetAtSubscriptionScope(ctx context.Context, id DeploymentOperationId) (result GetAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForGetAtSubscriptionScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAtSubscriptionScope prepares the GetAtSubscriptionScope request.
func (c DeploymentOperationsClient) preparerForGetAtSubscriptionScope(ctx context.Context, id DeploymentOperationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAtSubscriptionScope handles the response to the GetAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForGetAtSubscriptionScope(resp *http.Response) (result GetAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
