package iotconnectors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type DeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Delete ...
func (c IotConnectorsClient) Delete(ctx context.Context, id IotConnectorId) (result DeleteOperationResponse, err error) {
	req, err := c.preparerForDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotconnectors.IotConnectorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotconnectors.IotConnectorsClient", "Delete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteThenPoll performs Delete then polls until it's completed
func (c IotConnectorsClient) DeleteThenPoll(ctx context.Context, id IotConnectorId) error {
	result, err := c.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Delete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Delete: %+v", err)
	}

	return nil
}

// preparerForDelete prepares the Delete request.
func (c IotConnectorsClient) preparerForDelete(ctx context.Context, id IotConnectorId) (*http.Request, error) {
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

// senderForDelete sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (c IotConnectorsClient) senderForDelete(ctx context.Context, req *http.Request) (future DeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}