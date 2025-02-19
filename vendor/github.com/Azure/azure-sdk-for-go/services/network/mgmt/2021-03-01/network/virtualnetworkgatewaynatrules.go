package network

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

// VirtualNetworkGatewayNatRulesClient is the network Client
type VirtualNetworkGatewayNatRulesClient struct {
	BaseClient
}

// NewVirtualNetworkGatewayNatRulesClient creates an instance of the VirtualNetworkGatewayNatRulesClient client.
func NewVirtualNetworkGatewayNatRulesClient(subscriptionID string) VirtualNetworkGatewayNatRulesClient {
	return NewVirtualNetworkGatewayNatRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewayNatRulesClientWithBaseURI creates an instance of the VirtualNetworkGatewayNatRulesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewVirtualNetworkGatewayNatRulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewayNatRulesClient {
	return VirtualNetworkGatewayNatRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a nat rule to a scalable virtual network gateway if it doesn't exist else updates the
// existing nat rules.
// Parameters:
// resourceGroupName - the resource group name of the Virtual Network Gateway.
// virtualNetworkGatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
// natRuleParameters - parameters supplied to create or Update a Nat Rule.
func (client VirtualNetworkGatewayNatRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule) (result VirtualNetworkGatewayNatRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewayNatRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, natRuleParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewayNatRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natRuleName":               autorest.Encode("path", natRuleName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	natRuleParameters.Etag = nil
	natRuleParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithJSON(natRuleParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayNatRulesClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworkGatewayNatRulesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayNatRulesClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkGatewayNatRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a nat rule.
// Parameters:
// resourceGroupName - the resource group name of the Virtual Network Gateway.
// virtualNetworkGatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
func (client VirtualNetworkGatewayNatRulesClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string) (result VirtualNetworkGatewayNatRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewayNatRulesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkGatewayNatRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natRuleName":               autorest.Encode("path", natRuleName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayNatRulesClient) DeleteSender(req *http.Request) (future VirtualNetworkGatewayNatRulesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayNatRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a nat rule.
// Parameters:
// resourceGroupName - the resource group name of the Virtual Network Gateway.
// virtualNetworkGatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
func (client VirtualNetworkGatewayNatRulesClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string) (result VirtualNetworkGatewayNatRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewayNatRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkGatewayNatRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natRuleName":               autorest.Encode("path", natRuleName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayNatRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayNatRulesClient) GetResponder(resp *http.Response) (result VirtualNetworkGatewayNatRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVirtualNetworkGateway retrieves all nat rules for a particular virtual network gateway.
// Parameters:
// resourceGroupName - the resource group name of the virtual network gateway.
// virtualNetworkGatewayName - the name of the gateway.
func (client VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGateway(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result ListVirtualNetworkGatewayNatRulesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewayNatRulesClient.ListByVirtualNetworkGateway")
		defer func() {
			sc := -1
			if result.lvngnrr.Response.Response != nil {
				sc = result.lvngnrr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVirtualNetworkGatewayNextResults
	req, err := client.ListByVirtualNetworkGatewayPreparer(ctx, resourceGroupName, virtualNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "ListByVirtualNetworkGateway", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVirtualNetworkGatewaySender(req)
	if err != nil {
		result.lvngnrr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "ListByVirtualNetworkGateway", resp, "Failure sending request")
		return
	}

	result.lvngnrr, err = client.ListByVirtualNetworkGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "ListByVirtualNetworkGateway", resp, "Failure responding to request")
		return
	}
	if result.lvngnrr.hasNextLink() && result.lvngnrr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVirtualNetworkGatewayPreparer prepares the ListByVirtualNetworkGateway request.
func (client VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGatewayPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVirtualNetworkGatewaySender sends the ListByVirtualNetworkGateway request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGatewaySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVirtualNetworkGatewayResponder handles the response to the ListByVirtualNetworkGateway request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGatewayResponder(resp *http.Response) (result ListVirtualNetworkGatewayNatRulesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVirtualNetworkGatewayNextResults retrieves the next set of results, if any.
func (client VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayNextResults(ctx context.Context, lastResults ListVirtualNetworkGatewayNatRulesResult) (result ListVirtualNetworkGatewayNatRulesResult, err error) {
	req, err := lastResults.listVirtualNetworkGatewayNatRulesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "listByVirtualNetworkGatewayNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVirtualNetworkGatewaySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "listByVirtualNetworkGatewayNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVirtualNetworkGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayNatRulesClient", "listByVirtualNetworkGatewayNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVirtualNetworkGatewayComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkGatewayNatRulesClient) ListByVirtualNetworkGatewayComplete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result ListVirtualNetworkGatewayNatRulesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewayNatRulesClient.ListByVirtualNetworkGateway")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVirtualNetworkGateway(ctx, resourceGroupName, virtualNetworkGatewayName)
	return
}
