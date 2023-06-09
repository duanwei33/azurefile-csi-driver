package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CommunityGalleryImageVersionsClient is the compute Client
type CommunityGalleryImageVersionsClient struct {
	BaseClient
}

// NewCommunityGalleryImageVersionsClient creates an instance of the CommunityGalleryImageVersionsClient client.
func NewCommunityGalleryImageVersionsClient(subscriptionID string) CommunityGalleryImageVersionsClient {
	return NewCommunityGalleryImageVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommunityGalleryImageVersionsClientWithBaseURI creates an instance of the CommunityGalleryImageVersionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewCommunityGalleryImageVersionsClientWithBaseURI(baseURI string, subscriptionID string) CommunityGalleryImageVersionsClient {
	return CommunityGalleryImageVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a community gallery image version.
// Parameters:
// location - resource location.
// publicGalleryName - the public name of the community gallery.
// galleryImageName - the name of the community gallery image definition.
// galleryImageVersionName - the name of the community gallery image version. Needs to follow semantic version
// name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit
// integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
func (client CommunityGalleryImageVersionsClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string) (result CommunityGalleryImageVersion, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunityGalleryImageVersionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, publicGalleryName, galleryImageName, galleryImageVersionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommunityGalleryImageVersionsClient) GetPreparer(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":        autorest.Encode("path", galleryImageName),
		"galleryImageVersionName": autorest.Encode("path", galleryImageVersionName),
		"location":                autorest.Encode("path", location),
		"publicGalleryName":       autorest.Encode("path", publicGalleryName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommunityGalleryImageVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommunityGalleryImageVersionsClient) GetResponder(resp *http.Response) (result CommunityGalleryImageVersion, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list community gallery image versions inside an image.
// Parameters:
// location - resource location.
// publicGalleryName - the public name of the community gallery.
// galleryImageName - the name of the community gallery image definition.
func (client CommunityGalleryImageVersionsClient) List(ctx context.Context, location string, publicGalleryName string, galleryImageName string) (result CommunityGalleryImageVersionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunityGalleryImageVersionsClient.List")
		defer func() {
			sc := -1
			if result.cgivl.Response.Response != nil {
				sc = result.cgivl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, publicGalleryName, galleryImageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cgivl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "List", resp, "Failure sending request")
		return
	}

	result.cgivl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.cgivl.hasNextLink() && result.cgivl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CommunityGalleryImageVersionsClient) ListPreparer(ctx context.Context, location string, publicGalleryName string, galleryImageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"location":          autorest.Encode("path", location),
		"publicGalleryName": autorest.Encode("path", publicGalleryName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CommunityGalleryImageVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CommunityGalleryImageVersionsClient) ListResponder(resp *http.Response) (result CommunityGalleryImageVersionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CommunityGalleryImageVersionsClient) listNextResults(ctx context.Context, lastResults CommunityGalleryImageVersionList) (result CommunityGalleryImageVersionList, err error) {
	req, err := lastResults.communityGalleryImageVersionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleryImageVersionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CommunityGalleryImageVersionsClient) ListComplete(ctx context.Context, location string, publicGalleryName string, galleryImageName string) (result CommunityGalleryImageVersionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunityGalleryImageVersionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location, publicGalleryName, galleryImageName)
	return
}
