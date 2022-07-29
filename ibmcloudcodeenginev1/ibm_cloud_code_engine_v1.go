/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.15.0-45841b53-20201019-214802
 */
 

// Package ibmcloudcodeenginev1 : Operations and models for the IbmCloudCodeEngineV1 service
package ibmcloudcodeenginev1

import (
	"context"
	"fmt"
	common "github.com/IBM/code-engine-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// IbmCloudCodeEngineV1 : The purpose is to provide an API to get Kubeconfig file for IBM Cloud Code Engine Project
//
// Version: 0.0
type IbmCloudCodeEngineV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://ibm-cloud-code-engine.cloud.ibm.com/api/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_code_engine"

// IbmCloudCodeEngineV1Options : Service options
type IbmCloudCodeEngineV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudCodeEngineV1UsingExternalConfig : constructs an instance of IbmCloudCodeEngineV1 with passed in options and external configuration.
func NewIbmCloudCodeEngineV1UsingExternalConfig(options *IbmCloudCodeEngineV1Options) (ibmCloudCodeEngine *IbmCloudCodeEngineV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudCodeEngine, err = NewIbmCloudCodeEngineV1(options)
	if err != nil {
		return
	}

	err = ibmCloudCodeEngine.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudCodeEngine.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudCodeEngineV1 : constructs an instance of IbmCloudCodeEngineV1 with passed in options.
func NewIbmCloudCodeEngineV1(options *IbmCloudCodeEngineV1Options) (service *IbmCloudCodeEngineV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &IbmCloudCodeEngineV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) SetServiceURL(url string) error {
	return ibmCloudCodeEngine.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) GetServiceURL() string {
	return ibmCloudCodeEngine.Service.GetServiceURL()
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) SetEnableGzipCompression(enableGzip bool) {
	ibmCloudCodeEngine.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) GetEnableGzipCompression() bool {
	return ibmCloudCodeEngine.Service.GetEnableGzipCompression()
}

// ListKubeconfig : Deprecated soon: Retrieve KUBECONFIG for a specified project
// **Deprecated soon**: This API will be deprecated soon. Use the [GET /project/{id}/config](#get-kubeconfig) API
// instead. Returns the KUBECONFIG file, similar to the output of `kubectl config view --minify=true`.
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) ListKubeconfig(listKubeconfigOptions *ListKubeconfigOptions) (result *string, response *core.DetailedResponse, err error) {
	return ibmCloudCodeEngine.ListKubeconfigWithContext(context.Background(), listKubeconfigOptions)
}

// ListKubeconfigWithContext is an alternate form of the ListKubeconfig method which supports a Context parameter
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) ListKubeconfigWithContext(ctx context.Context, listKubeconfigOptions *ListKubeconfigOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listKubeconfigOptions, "listKubeconfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listKubeconfigOptions, "listKubeconfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listKubeconfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudCodeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudCodeEngine.Service.Options.URL, `/namespaces/{id}/config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listKubeconfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_code_engine", "V1", "ListKubeconfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain")
	if listKubeconfigOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*listKubeconfigOptions.RefreshToken))
	}
	if listKubeconfigOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listKubeconfigOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudCodeEngine.Service.Request(request, &result)

	return
}

// GetKubeconfig : Retrieve KUBECONFIG for a specified project
// Returns the KUBECONFIG, similar to the output of `kubectl config view --minify=true`. There are 2 tokens in the
// Request Header and a query parameter that you must provide.
//  These values can be generated as follows: 1. Auth Header Pass the generated IAM Token as the Authorization header
// from the CLI as `token=cat $HOME/.bluemix/config.json | jq .IAMToken -r`. Generate the token with the [Create an IAM
// access token for a user or service ID using an API
// key](https://cloud.ibm.com/apidocs/iam-identity-token-api#gettoken-apikey) API.
//
// 2. X-Delegated-Refresh-Token Header Generate an IAM Delegated Refresh Token for Code Engine with the [Create an IAM
// access token and delegated refresh token for a user or service
// ID](https://cloud.ibm.com/apidocs/iam-identity-token-api#gettoken-apikey-delegatedrefreshtoken) API. Specify the
// `receiver_client_ids` value to be `ce` and the `delegated_refresh_token_expiry` value to be `3600`.
//
// 3. Project ID In order to retrieve the Kubeconfig file for a specific Code Engine project, use the CLI to extract the
// ID
// `id=ibmcloud ce project get -n ${CE_PROJECT_NAME} -o jsonpath={.guid}` You must be logged into the account where the
// project was created to retrieve the ID.
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) GetKubeconfig(getKubeconfigOptions *GetKubeconfigOptions) (result *string, response *core.DetailedResponse, err error) {
	return ibmCloudCodeEngine.GetKubeconfigWithContext(context.Background(), getKubeconfigOptions)
}

// GetKubeconfigWithContext is an alternate form of the GetKubeconfig method which supports a Context parameter
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) GetKubeconfigWithContext(ctx context.Context, getKubeconfigOptions *GetKubeconfigOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getKubeconfigOptions, "getKubeconfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getKubeconfigOptions, "getKubeconfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getKubeconfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudCodeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudCodeEngine.Service.Options.URL, `/project/{id}/config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getKubeconfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_code_engine", "V1", "GetKubeconfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain")
	if getKubeconfigOptions.XDelegatedRefreshToken != nil {
		builder.AddHeader("X-Delegated-Refresh-Token", fmt.Sprint(*getKubeconfigOptions.XDelegatedRefreshToken))
	}
	if getKubeconfigOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getKubeconfigOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudCodeEngine.Service.Request(request, &result)

	return
}

// GetKubeconfigOptions : The GetKubeconfig options.
type GetKubeconfigOptions struct {
	// This IAM Delegated Refresh Token is specifically valid for Code Engine. Generate this token with the [Create an IAM
	// access token and delegated refresh token for a user or service
	// ID](https://cloud.ibm.com/apidocs/iam-identity-token-api#gettoken-apikey-delegatedrefreshtoken) API. Specify the
	// `receiver_client_ids` value to be `ce` and the `delegated_refresh_token_expiry` value to be `3600`.
	XDelegatedRefreshToken *string `json:"X-Delegated-Refresh-Token" validate:"required"`

	// The id of the IBM Cloud Code Engine project.
	ID *string `json:"id" validate:"required,ne="`

	// The type of the response: text/plain or application/json. A character encoding can be specified by including a
	// `charset` parameter. For example, 'text/plain;charset=utf-8'.
	Accept *string `json:"Accept,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetKubeconfigOptions : Instantiate GetKubeconfigOptions
func (*IbmCloudCodeEngineV1) NewGetKubeconfigOptions(xDelegatedRefreshToken string, id string) *GetKubeconfigOptions {
	return &GetKubeconfigOptions{
		XDelegatedRefreshToken: core.StringPtr(xDelegatedRefreshToken),
		ID: core.StringPtr(id),
	}
}

// SetXDelegatedRefreshToken : Allow user to set XDelegatedRefreshToken
func (options *GetKubeconfigOptions) SetXDelegatedRefreshToken(xDelegatedRefreshToken string) *GetKubeconfigOptions {
	options.XDelegatedRefreshToken = core.StringPtr(xDelegatedRefreshToken)
	return options
}

// SetID : Allow user to set ID
func (options *GetKubeconfigOptions) SetID(id string) *GetKubeconfigOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccept : Allow user to set Accept
func (options *GetKubeconfigOptions) SetAccept(accept string) *GetKubeconfigOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetKubeconfigOptions) SetHeaders(param map[string]string) *GetKubeconfigOptions {
	options.Headers = param
	return options
}

// ListKubeconfigOptions : The ListKubeconfig options.
type ListKubeconfigOptions struct {
	// The IAM Refresh token associated with the IBM Cloud account. To retrieve your IAM token, run `ibmcloud iam
	// oauth-tokens`.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// The id of the IBM Cloud Code Engine project. To retrieve your project ID, run `ibmcloud ce project get -n
	// <PROJECT_NAME>`.
	ID *string `json:"id" validate:"required,ne="`

	// The type of the response: text/plain or application/json. A character encoding can be specified by including a
	// `charset` parameter. For example, 'text/plain;charset=utf-8'.
	Accept *string `json:"Accept,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListKubeconfigOptions : Instantiate ListKubeconfigOptions
func (*IbmCloudCodeEngineV1) NewListKubeconfigOptions(refreshToken string, id string) *ListKubeconfigOptions {
	return &ListKubeconfigOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ID: core.StringPtr(id),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (options *ListKubeconfigOptions) SetRefreshToken(refreshToken string) *ListKubeconfigOptions {
	options.RefreshToken = core.StringPtr(refreshToken)
	return options
}

// SetID : Allow user to set ID
func (options *ListKubeconfigOptions) SetID(id string) *ListKubeconfigOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccept : Allow user to set Accept
func (options *ListKubeconfigOptions) SetAccept(accept string) *ListKubeconfigOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListKubeconfigOptions) SetHeaders(param map[string]string) *ListKubeconfigOptions {
	options.Headers = param
	return options
}
