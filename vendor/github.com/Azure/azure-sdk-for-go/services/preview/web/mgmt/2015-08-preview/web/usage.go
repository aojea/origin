package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// UsageClient is the webSite Management Client
type UsageClient struct {
	BaseClient
}

// NewUsageClient creates an instance of the UsageClient client.
func NewUsageClient(subscriptionID string) UsageClient {
	return NewUsageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageClientWithBaseURI creates an instance of the UsageClient client.
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return UsageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetUsage sends the get usage request.
// Parameters:
// resourceGroupName - name of resource group
// environmentName - environment name
// lastID - last marker that was returned from the batch
// batchSize - size of the batch to be returned.
func (client UsageClient) GetUsage(ctx context.Context, resourceGroupName string, environmentName string, lastID string, batchSize int32) (result SetObject, err error) {
	req, err := client.GetUsagePreparer(ctx, resourceGroupName, environmentName, lastID, batchSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", resp, "Failure sending request")
		return
	}

	result, err = client.GetUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", resp, "Failure responding to request")
	}

	return
}

// GetUsagePreparer prepares the GetUsage request.
func (client UsageClient) GetUsagePreparer(ctx context.Context, resourceGroupName string, environmentName string, lastID string, batchSize int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"batchSize":   autorest.Encode("query", batchSize),
		"lastId":      autorest.Encode("query", lastID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web.Admin/environments/{environmentName}/usage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUsageSender sends the GetUsage request. The method will close the
// http.Response Body if it receives an error.
func (client UsageClient) GetUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetUsageResponder handles the response to the GetUsage request. The method always
// closes the http.Response Body.
func (client UsageClient) GetUsageResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
