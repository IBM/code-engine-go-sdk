/**
 * (C) Copyright IBM Corp. 2020.
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
 * IBM OpenAPI SDK Code Generator Version: 3.12.0-64fe8d3f-20200820-144050
 */
 

// Package ibmcloudcodeenginev1 : Operations and models for the IbmCloudCodeEngineV1 service
package ibmcloudcodeenginev1

import (
	"fmt"
	common "github.com/IBM/code-engine-go-sdk/common"
	"github.com/IBM/go-sdk-core/v4/core"
)

// IbmCloudCodeEngineV1 : The purpose is to provide an API to get Kubeconfig for IBM Cloud Code Engine Project
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

// ListKubeconfig : Retrieve KUBECONFIG for a specified project
// Returns the KUBECONFIG, similar to the output of `kubectl config view --minify=true`.
func (ibmCloudCodeEngine *IbmCloudCodeEngineV1) ListKubeconfig(listKubeconfigOptions *ListKubeconfigOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listKubeconfigOptions, "listKubeconfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listKubeconfigOptions, "listKubeconfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces", "config"}
	pathParameters := []string{*listKubeconfigOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmCloudCodeEngine.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddHeader("Accept", "text/html")
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

// ListKubeconfigOptions : The ListKubeconfig options.
type ListKubeconfigOptions struct {
	// The IAM Refresh token associated with the IBM Cloud account.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// The id of the IBM Cloud Code Engine project.
	ID *string `json:"id" validate:"required"`

	// The type of the response: text/html or application/json. A character encoding can be specified by including a
	// `charset` parameter. For example, 'text/html;charset=utf-8'.
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
