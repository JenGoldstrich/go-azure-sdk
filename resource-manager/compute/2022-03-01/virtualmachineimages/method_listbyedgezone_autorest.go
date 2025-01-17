package virtualmachineimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByEdgeZoneOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMImagesInEdgeZoneListResult
}

// ListByEdgeZone ...
func (c VirtualMachineImagesClient) ListByEdgeZone(ctx context.Context, id EdgeZoneId) (result ListByEdgeZoneOperationResponse, err error) {
	req, err := c.preparerForListByEdgeZone(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListByEdgeZone", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListByEdgeZone", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByEdgeZone(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListByEdgeZone", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByEdgeZone prepares the ListByEdgeZone request.
func (c VirtualMachineImagesClient) preparerForListByEdgeZone(ctx context.Context, id EdgeZoneId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/vmimages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByEdgeZone handles the response to the ListByEdgeZone request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForListByEdgeZone(resp *http.Response) (result ListByEdgeZoneOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
