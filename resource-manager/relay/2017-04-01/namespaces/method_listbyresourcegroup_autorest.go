package namespaces

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

type ListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RelayNamespace

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByResourceGroupOperationResponse, error)
}

type ListByResourceGroupCompleteResult struct {
	Items []RelayNamespace
}

func (r ListByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp ListByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByResourceGroup ...
func (c NamespacesClient) ListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (resp ListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListByResourceGroupComplete retrieves all of the results into a single object
func (c NamespacesClient) ListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ListByResourceGroupCompleteResult, error) {
	return c.ListByResourceGroupCompleteMatchingPredicate(ctx, id, RelayNamespaceOperationPredicate{})
}

// ListByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c NamespacesClient) ListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate RelayNamespaceOperationPredicate) (resp ListByResourceGroupCompleteResult, err error) {
	items := make([]RelayNamespace, 0)

	page, err := c.ListByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListByResourceGroup prepares the ListByResourceGroup request.
func (c NamespacesClient) preparerForListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Relay/namespaces", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByResourceGroupWithNextLink prepares the ListByResourceGroup request with the given nextLink token.
func (c NamespacesClient) preparerForListByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByResourceGroup handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (c NamespacesClient) responderForListByResourceGroup(resp *http.Response) (result ListByResourceGroupOperationResponse, err error) {
	type page struct {
		Values   []RelayNamespace `json:"value"`
		NextLink *string          `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForListByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "namespaces.NamespacesClient", "ListByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
