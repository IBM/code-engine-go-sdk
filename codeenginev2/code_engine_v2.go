/**
 * (C) Copyright IBM Corp. 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.53.0-9710cac3-20220713-193508
 */

// Package codeenginev2 : Operations and models for the CodeEngineV2 service
package codeenginev2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/code-engine-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// CodeEngineV2 : REST API for Code Engine
//
// API Version: 2.0.0
type CodeEngineV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.au-syd.codeengine.cloud.ibm.com/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "code_engine"

// CodeEngineV2Options : Service options
type CodeEngineV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCodeEngineV2UsingExternalConfig : constructs an instance of CodeEngineV2 with passed in options and external configuration.
func NewCodeEngineV2UsingExternalConfig(options *CodeEngineV2Options) (codeEngine *CodeEngineV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	codeEngine, err = NewCodeEngineV2(options)
	if err != nil {
		return
	}

	err = codeEngine.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = codeEngine.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCodeEngineV2 : constructs an instance of CodeEngineV2 with passed in options.
func NewCodeEngineV2(options *CodeEngineV2Options) (service *CodeEngineV2, err error) {
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

	service = &CodeEngineV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "codeEngine" suitable for processing requests.
func (codeEngine *CodeEngineV2) Clone() *CodeEngineV2 {
	if core.IsNil(codeEngine) {
		return nil
	}
	clone := *codeEngine
	clone.Service = codeEngine.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (codeEngine *CodeEngineV2) SetServiceURL(url string) error {
	return codeEngine.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (codeEngine *CodeEngineV2) GetServiceURL() string {
	return codeEngine.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (codeEngine *CodeEngineV2) SetDefaultHeaders(headers http.Header) {
	codeEngine.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (codeEngine *CodeEngineV2) SetEnableGzipCompression(enableGzip bool) {
	codeEngine.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (codeEngine *CodeEngineV2) GetEnableGzipCompression() bool {
	return codeEngine.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (codeEngine *CodeEngineV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	codeEngine.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (codeEngine *CodeEngineV2) DisableRetries() {
	codeEngine.Service.DisableRetries()
}

// ListConfigmapsV2 : List configmaps
// List Configmaps.
func (codeEngine *CodeEngineV2) ListConfigmapsV2(listConfigmapsV2Options *ListConfigmapsV2Options) (result *V2ConfigMapList, response *core.DetailedResponse, err error) {
	return codeEngine.ListConfigmapsV2WithContext(context.Background(), listConfigmapsV2Options)
}

// ListConfigmapsV2WithContext is an alternate form of the ListConfigmapsV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListConfigmapsV2WithContext(ctx context.Context, listConfigmapsV2Options *ListConfigmapsV2Options) (result *V2ConfigMapList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigmapsV2Options, "listConfigmapsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigmapsV2Options, "listConfigmapsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listConfigmapsV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigmapsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListConfigmapsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listConfigmapsV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*listConfigmapsV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ConfigMapList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateConfigmapV2 : Create a configmap
// Create a Configmap.
func (codeEngine *CodeEngineV2) CreateConfigmapV2(createConfigmapV2Options *CreateConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.CreateConfigmapV2WithContext(context.Background(), createConfigmapV2Options)
}

// CreateConfigmapV2WithContext is an alternate form of the CreateConfigmapV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateConfigmapV2WithContext(ctx context.Context, createConfigmapV2Options *CreateConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createConfigmapV2Options, "createConfigmapV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createConfigmapV2Options, "createConfigmapV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createConfigmapV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigmapV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateConfigmapV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createConfigmapV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*createConfigmapV2Options.RefreshToken))
	}

	body := make(map[string]interface{})
	if createConfigmapV2Options.Created != nil {
		body["created"] = createConfigmapV2Options.Created
	}
	if createConfigmapV2Options.Data != nil {
		body["data"] = createConfigmapV2Options.Data
	}
	if createConfigmapV2Options.ID != nil {
		body["id"] = createConfigmapV2Options.ID
	}
	if createConfigmapV2Options.Immutable != nil {
		body["immutable"] = createConfigmapV2Options.Immutable
	}
	if createConfigmapV2Options.Name != nil {
		body["name"] = createConfigmapV2Options.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ConfigMap)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfigmapV2 : Get a configmap
// Get a Configmap.
func (codeEngine *CodeEngineV2) GetConfigmapV2(getConfigmapV2Options *GetConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.GetConfigmapV2WithContext(context.Background(), getConfigmapV2Options)
}

// GetConfigmapV2WithContext is an alternate form of the GetConfigmapV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetConfigmapV2WithContext(ctx context.Context, getConfigmapV2Options *GetConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigmapV2Options, "getConfigmapV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigmapV2Options, "getConfigmapV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getConfigmapV2Options.ProjectGuid,
		"configmap_name": *getConfigmapV2Options.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigmapV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetConfigmapV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getConfigmapV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*getConfigmapV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ConfigMap)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteConfigmapV2 : Delete a configmap
// Delete a Configmap.
func (codeEngine *CodeEngineV2) DeleteConfigmapV2(deleteConfigmapV2Options *DeleteConfigmapV2Options) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteConfigmapV2WithContext(context.Background(), deleteConfigmapV2Options)
}

// DeleteConfigmapV2WithContext is an alternate form of the DeleteConfigmapV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteConfigmapV2WithContext(ctx context.Context, deleteConfigmapV2Options *DeleteConfigmapV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConfigmapV2Options, "deleteConfigmapV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteConfigmapV2Options, "deleteConfigmapV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteConfigmapV2Options.ProjectGuid,
		"configmap_name": *deleteConfigmapV2Options.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigmapV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteConfigmapV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteConfigmapV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*deleteConfigmapV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// UpdateConfigmapV2 : Update a configmap
// Update a Configmap.
func (codeEngine *CodeEngineV2) UpdateConfigmapV2(updateConfigmapV2Options *UpdateConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateConfigmapV2WithContext(context.Background(), updateConfigmapV2Options)
}

// UpdateConfigmapV2WithContext is an alternate form of the UpdateConfigmapV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateConfigmapV2WithContext(ctx context.Context, updateConfigmapV2Options *UpdateConfigmapV2Options) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigmapV2Options, "updateConfigmapV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigmapV2Options, "updateConfigmapV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateConfigmapV2Options.ProjectGuid,
		"configmap_name": *updateConfigmapV2Options.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigmapV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateConfigmapV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateConfigmapV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*updateConfigmapV2Options.RefreshToken))
	}

	body := make(map[string]interface{})
	if updateConfigmapV2Options.Created != nil {
		body["created"] = updateConfigmapV2Options.Created
	}
	if updateConfigmapV2Options.Data != nil {
		body["data"] = updateConfigmapV2Options.Data
	}
	if updateConfigmapV2Options.ID != nil {
		body["id"] = updateConfigmapV2Options.ID
	}
	if updateConfigmapV2Options.Immutable != nil {
		body["immutable"] = updateConfigmapV2Options.Immutable
	}
	if updateConfigmapV2Options.Name != nil {
		body["name"] = updateConfigmapV2Options.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ConfigMap)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProjectsV2 : List all projects
// List projects.
func (codeEngine *CodeEngineV2) ListProjectsV2(listProjectsV2Options *ListProjectsV2Options) (result *V2ProjectList, response *core.DetailedResponse, err error) {
	return codeEngine.ListProjectsV2WithContext(context.Background(), listProjectsV2Options)
}

// ListProjectsV2WithContext is an alternate form of the ListProjectsV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListProjectsV2WithContext(ctx context.Context, listProjectsV2Options *ListProjectsV2Options) (result *V2ProjectList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listProjectsV2Options, "listProjectsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listProjectsV2Options, "listProjectsV2Options")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListProjectsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listProjectsV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*listProjectsV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ProjectList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateProjectV2 : Create a Project
// Create a project.
func (codeEngine *CodeEngineV2) CreateProjectV2(createProjectV2Options *CreateProjectV2Options) (result *V2Project, response *core.DetailedResponse, err error) {
	return codeEngine.CreateProjectV2WithContext(context.Background(), createProjectV2Options)
}

// CreateProjectV2WithContext is an alternate form of the CreateProjectV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateProjectV2WithContext(ctx context.Context, createProjectV2Options *CreateProjectV2Options) (result *V2Project, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProjectV2Options, "createProjectV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProjectV2Options, "createProjectV2Options")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProjectV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateProjectV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createProjectV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*createProjectV2Options.RefreshToken))
	}

	body := make(map[string]interface{})
	if createProjectV2Options.Name != nil {
		body["name"] = createProjectV2Options.Name
	}
	if createProjectV2Options.Region != nil {
		body["region"] = createProjectV2Options.Region
	}
	if createProjectV2Options.ResourceGroupID != nil {
		body["resource_group_id"] = createProjectV2Options.ResourceGroupID
	}
	if createProjectV2Options.Tags != nil {
		body["tags"] = createProjectV2Options.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Project)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProjectV2 : Get a project
// Retrieve the project.
func (codeEngine *CodeEngineV2) GetProjectV2(getProjectV2Options *GetProjectV2Options) (result *V2Project, response *core.DetailedResponse, err error) {
	return codeEngine.GetProjectV2WithContext(context.Background(), getProjectV2Options)
}

// GetProjectV2WithContext is an alternate form of the GetProjectV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetProjectV2WithContext(ctx context.Context, getProjectV2Options *GetProjectV2Options) (result *V2Project, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectV2Options, "getProjectV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectV2Options, "getProjectV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getProjectV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetProjectV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getProjectV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*getProjectV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Project)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteProjectV2 : Delete a Project
// Delete a project.
func (codeEngine *CodeEngineV2) DeleteProjectV2(deleteProjectV2Options *DeleteProjectV2Options) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteProjectV2WithContext(context.Background(), deleteProjectV2Options)
}

// DeleteProjectV2WithContext is an alternate form of the DeleteProjectV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteProjectV2WithContext(ctx context.Context, deleteProjectV2Options *DeleteProjectV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProjectV2Options, "deleteProjectV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProjectV2Options, "deleteProjectV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteProjectV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteProjectV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteProjectV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*deleteProjectV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// ListReclamationsV2 : List all reclamations
// List reclamations.
func (codeEngine *CodeEngineV2) ListReclamationsV2(listReclamationsV2Options *ListReclamationsV2Options) (result *V2ReclamationList, response *core.DetailedResponse, err error) {
	return codeEngine.ListReclamationsV2WithContext(context.Background(), listReclamationsV2Options)
}

// ListReclamationsV2WithContext is an alternate form of the ListReclamationsV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListReclamationsV2WithContext(ctx context.Context, listReclamationsV2Options *ListReclamationsV2Options) (result *V2ReclamationList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listReclamationsV2Options, "listReclamationsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listReclamationsV2Options, "listReclamationsV2Options")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listReclamationsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListReclamationsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listReclamationsV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*listReclamationsV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2ReclamationList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReclamationV2 : Get a reclamation
// Get a reclamation.
func (codeEngine *CodeEngineV2) GetReclamationV2(getReclamationV2Options *GetReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.GetReclamationV2WithContext(context.Background(), getReclamationV2Options)
}

// GetReclamationV2WithContext is an alternate form of the GetReclamationV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetReclamationV2WithContext(ctx context.Context, getReclamationV2Options *GetReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReclamationV2Options, "getReclamationV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReclamationV2Options, "getReclamationV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getReclamationV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReclamationV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetReclamationV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReclamationV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*getReclamationV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Reclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReclaimReclamationV2 : Reclaim a reclamation
// Reclaim a reclaimation to hard delete a project.
func (codeEngine *CodeEngineV2) ReclaimReclamationV2(reclaimReclamationV2Options *ReclaimReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.ReclaimReclamationV2WithContext(context.Background(), reclaimReclamationV2Options)
}

// ReclaimReclamationV2WithContext is an alternate form of the ReclaimReclamationV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) ReclaimReclamationV2WithContext(ctx context.Context, reclaimReclamationV2Options *ReclaimReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reclaimReclamationV2Options, "reclaimReclamationV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reclaimReclamationV2Options, "reclaimReclamationV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *reclaimReclamationV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}/reclaim`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range reclaimReclamationV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ReclaimReclamationV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if reclaimReclamationV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*reclaimReclamationV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Reclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RestoreReclamationV2 : Restore a reclamation
// Restore a reclaimation which restores a soft-deleted project.
func (codeEngine *CodeEngineV2) RestoreReclamationV2(restoreReclamationV2Options *RestoreReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.RestoreReclamationV2WithContext(context.Background(), restoreReclamationV2Options)
}

// RestoreReclamationV2WithContext is an alternate form of the RestoreReclamationV2 method which supports a Context parameter
func (codeEngine *CodeEngineV2) RestoreReclamationV2WithContext(ctx context.Context, restoreReclamationV2Options *RestoreReclamationV2Options) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restoreReclamationV2Options, "restoreReclamationV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(restoreReclamationV2Options, "restoreReclamationV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *restoreReclamationV2Options.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}/restore`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range restoreReclamationV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "RestoreReclamationV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if restoreReclamationV2Options.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*restoreReclamationV2Options.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = codeEngine.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Reclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateConfigmapV2Options : The CreateConfigmapV2 options.
type CreateConfigmapV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	Created *string `json:"created,omitempty"`

	Data map[string]string `json:"data,omitempty"`

	ID *string `json:"id,omitempty"`

	Immutable *bool `json:"immutable,omitempty"`

	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigmapV2Options : Instantiate CreateConfigmapV2Options
func (*CodeEngineV2) NewCreateConfigmapV2Options(refreshToken string, projectGuid string) *CreateConfigmapV2Options {
	return &CreateConfigmapV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *CreateConfigmapV2Options) SetRefreshToken(refreshToken string) *CreateConfigmapV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateConfigmapV2Options) SetProjectGuid(projectGuid string) *CreateConfigmapV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetCreated : Allow user to set Created
func (_options *CreateConfigmapV2Options) SetCreated(created string) *CreateConfigmapV2Options {
	_options.Created = core.StringPtr(created)
	return _options
}

// SetData : Allow user to set Data
func (_options *CreateConfigmapV2Options) SetData(data map[string]string) *CreateConfigmapV2Options {
	_options.Data = data
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateConfigmapV2Options) SetID(id string) *CreateConfigmapV2Options {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *CreateConfigmapV2Options) SetImmutable(immutable bool) *CreateConfigmapV2Options {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateConfigmapV2Options) SetName(name string) *CreateConfigmapV2Options {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigmapV2Options) SetHeaders(param map[string]string) *CreateConfigmapV2Options {
	options.Headers = param
	return options
}

// CreateProjectV2Options : The CreateProjectV2 options.
type CreateProjectV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Specify the project name.
	Name *string `json:"name,omitempty"`

	// Specify the id of the regin (us-south, eu-de).
	Region *string `json:"region,omitempty"`

	// Specify the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// resource instance tags.
	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectV2Options : Instantiate CreateProjectV2Options
func (*CodeEngineV2) NewCreateProjectV2Options(refreshToken string) *CreateProjectV2Options {
	return &CreateProjectV2Options{
		RefreshToken: core.StringPtr(refreshToken),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *CreateProjectV2Options) SetRefreshToken(refreshToken string) *CreateProjectV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateProjectV2Options) SetName(name string) *CreateProjectV2Options {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateProjectV2Options) SetRegion(region string) *CreateProjectV2Options {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (_options *CreateProjectV2Options) SetResourceGroupID(resourceGroupID string) *CreateProjectV2Options {
	_options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateProjectV2Options) SetTags(tags []string) *CreateProjectV2Options {
	_options.Tags = tags
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectV2Options) SetHeaders(param map[string]string) *CreateProjectV2Options {
	options.Headers = param
	return options
}

// DeleteConfigmapV2Options : The DeleteConfigmapV2 options.
type DeleteConfigmapV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// ConfigMap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigmapV2Options : Instantiate DeleteConfigmapV2Options
func (*CodeEngineV2) NewDeleteConfigmapV2Options(refreshToken string, projectGuid string, configmapName string) *DeleteConfigmapV2Options {
	return &DeleteConfigmapV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *DeleteConfigmapV2Options) SetRefreshToken(refreshToken string) *DeleteConfigmapV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteConfigmapV2Options) SetProjectGuid(projectGuid string) *DeleteConfigmapV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *DeleteConfigmapV2Options) SetConfigmapName(configmapName string) *DeleteConfigmapV2Options {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigmapV2Options) SetHeaders(param map[string]string) *DeleteConfigmapV2Options {
	options.Headers = param
	return options
}

// DeleteProjectV2Options : The DeleteProjectV2 options.
type DeleteProjectV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectV2Options : Instantiate DeleteProjectV2Options
func (*CodeEngineV2) NewDeleteProjectV2Options(refreshToken string, projectGuid string) *DeleteProjectV2Options {
	return &DeleteProjectV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *DeleteProjectV2Options) SetRefreshToken(refreshToken string) *DeleteProjectV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteProjectV2Options) SetProjectGuid(projectGuid string) *DeleteProjectV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectV2Options) SetHeaders(param map[string]string) *DeleteProjectV2Options {
	options.Headers = param
	return options
}

// GetConfigmapV2Options : The GetConfigmapV2 options.
type GetConfigmapV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// ConfigMap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigmapV2Options : Instantiate GetConfigmapV2Options
func (*CodeEngineV2) NewGetConfigmapV2Options(refreshToken string, projectGuid string, configmapName string) *GetConfigmapV2Options {
	return &GetConfigmapV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *GetConfigmapV2Options) SetRefreshToken(refreshToken string) *GetConfigmapV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetConfigmapV2Options) SetProjectGuid(projectGuid string) *GetConfigmapV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *GetConfigmapV2Options) SetConfigmapName(configmapName string) *GetConfigmapV2Options {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigmapV2Options) SetHeaders(param map[string]string) *GetConfigmapV2Options {
	options.Headers = param
	return options
}

// GetProjectV2Options : The GetProjectV2 options.
type GetProjectV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectV2Options : Instantiate GetProjectV2Options
func (*CodeEngineV2) NewGetProjectV2Options(refreshToken string, projectGuid string) *GetProjectV2Options {
	return &GetProjectV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *GetProjectV2Options) SetRefreshToken(refreshToken string) *GetProjectV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetProjectV2Options) SetProjectGuid(projectGuid string) *GetProjectV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectV2Options) SetHeaders(param map[string]string) *GetProjectV2Options {
	options.Headers = param
	return options
}

// GetReclamationV2Options : The GetReclamationV2 options.
type GetReclamationV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReclamationV2Options : Instantiate GetReclamationV2Options
func (*CodeEngineV2) NewGetReclamationV2Options(refreshToken string, projectGuid string) *GetReclamationV2Options {
	return &GetReclamationV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *GetReclamationV2Options) SetRefreshToken(refreshToken string) *GetReclamationV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetReclamationV2Options) SetProjectGuid(projectGuid string) *GetReclamationV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReclamationV2Options) SetHeaders(param map[string]string) *GetReclamationV2Options {
	options.Headers = param
	return options
}

// ListConfigmapsV2Options : The ListConfigmapsV2 options.
type ListConfigmapsV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigmapsV2Options : Instantiate ListConfigmapsV2Options
func (*CodeEngineV2) NewListConfigmapsV2Options(refreshToken string, projectGuid string) *ListConfigmapsV2Options {
	return &ListConfigmapsV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *ListConfigmapsV2Options) SetRefreshToken(refreshToken string) *ListConfigmapsV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListConfigmapsV2Options) SetProjectGuid(projectGuid string) *ListConfigmapsV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigmapsV2Options) SetHeaders(param map[string]string) *ListConfigmapsV2Options {
	options.Headers = param
	return options
}

// ListProjectsV2Options : The ListProjectsV2 options.
type ListProjectsV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsV2Options : Instantiate ListProjectsV2Options
func (*CodeEngineV2) NewListProjectsV2Options(refreshToken string) *ListProjectsV2Options {
	return &ListProjectsV2Options{
		RefreshToken: core.StringPtr(refreshToken),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *ListProjectsV2Options) SetRefreshToken(refreshToken string) *ListProjectsV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsV2Options) SetHeaders(param map[string]string) *ListProjectsV2Options {
	options.Headers = param
	return options
}

// ListReclamationsV2Options : The ListReclamationsV2 options.
type ListReclamationsV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListReclamationsV2Options : Instantiate ListReclamationsV2Options
func (*CodeEngineV2) NewListReclamationsV2Options(refreshToken string) *ListReclamationsV2Options {
	return &ListReclamationsV2Options{
		RefreshToken: core.StringPtr(refreshToken),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *ListReclamationsV2Options) SetRefreshToken(refreshToken string) *ListReclamationsV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListReclamationsV2Options) SetHeaders(param map[string]string) *ListReclamationsV2Options {
	options.Headers = param
	return options
}

// ReclaimReclamationV2Options : The ReclaimReclamationV2 options.
type ReclaimReclamationV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReclaimReclamationV2Options : Instantiate ReclaimReclamationV2Options
func (*CodeEngineV2) NewReclaimReclamationV2Options(refreshToken string, projectGuid string) *ReclaimReclamationV2Options {
	return &ReclaimReclamationV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *ReclaimReclamationV2Options) SetRefreshToken(refreshToken string) *ReclaimReclamationV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ReclaimReclamationV2Options) SetProjectGuid(projectGuid string) *ReclaimReclamationV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReclaimReclamationV2Options) SetHeaders(param map[string]string) *ReclaimReclamationV2Options {
	options.Headers = param
	return options
}

// RestoreReclamationV2Options : The RestoreReclamationV2 options.
type RestoreReclamationV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestoreReclamationV2Options : Instantiate RestoreReclamationV2Options
func (*CodeEngineV2) NewRestoreReclamationV2Options(refreshToken string, projectGuid string) *RestoreReclamationV2Options {
	return &RestoreReclamationV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *RestoreReclamationV2Options) SetRefreshToken(refreshToken string) *RestoreReclamationV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *RestoreReclamationV2Options) SetProjectGuid(projectGuid string) *RestoreReclamationV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RestoreReclamationV2Options) SetHeaders(param map[string]string) *RestoreReclamationV2Options {
	options.Headers = param
	return options
}

// UpdateConfigmapV2Options : The UpdateConfigmapV2 options.
type UpdateConfigmapV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Configmap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	Created *string `json:"created,omitempty"`

	Data map[string]string `json:"data,omitempty"`

	ID *string `json:"id,omitempty"`

	Immutable *bool `json:"immutable,omitempty"`

	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigmapV2Options : Instantiate UpdateConfigmapV2Options
func (*CodeEngineV2) NewUpdateConfigmapV2Options(refreshToken string, projectGuid string, configmapName string) *UpdateConfigmapV2Options {
	return &UpdateConfigmapV2Options{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *UpdateConfigmapV2Options) SetRefreshToken(refreshToken string) *UpdateConfigmapV2Options {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateConfigmapV2Options) SetProjectGuid(projectGuid string) *UpdateConfigmapV2Options {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *UpdateConfigmapV2Options) SetConfigmapName(configmapName string) *UpdateConfigmapV2Options {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetCreated : Allow user to set Created
func (_options *UpdateConfigmapV2Options) SetCreated(created string) *UpdateConfigmapV2Options {
	_options.Created = core.StringPtr(created)
	return _options
}

// SetData : Allow user to set Data
func (_options *UpdateConfigmapV2Options) SetData(data map[string]string) *UpdateConfigmapV2Options {
	_options.Data = data
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateConfigmapV2Options) SetID(id string) *UpdateConfigmapV2Options {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *UpdateConfigmapV2Options) SetImmutable(immutable bool) *UpdateConfigmapV2Options {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateConfigmapV2Options) SetName(name string) *UpdateConfigmapV2Options {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigmapV2Options) SetHeaders(param map[string]string) *UpdateConfigmapV2Options {
	options.Headers = param
	return options
}

// PaginationListNextMetadata : PaginationListNextMetadata struct
type PaginationListNextMetadata struct {
	Href *string `json:"href,omitempty"`

	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginationListNextMetadata unmarshals an instance of PaginationListNextMetadata from the specified map of raw messages.
func UnmarshalPaginationListNextMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationListNextMetadata)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2ConfigMap : V2ConfigMap struct
type V2ConfigMap struct {
	Created *string `json:"created,omitempty"`

	Data map[string]string `json:"data,omitempty"`

	ID *string `json:"id,omitempty"`

	Immutable *bool `json:"immutable,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`
}

// UnmarshalV2ConfigMap unmarshals an instance of V2ConfigMap from the specified map of raw messages.
func UnmarshalV2ConfigMap(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2ConfigMap)
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "immutable", &obj.Immutable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2ConfigMapList : V2ConfigMapList struct
type V2ConfigMapList struct {
	Configmaps []V2ConfigMap `json:"configmaps,omitempty"`

	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalV2ConfigMapList unmarshals an instance of V2ConfigMapList from the specified map of raw messages.
func UnmarshalV2ConfigMapList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2ConfigMapList)
	err = core.UnmarshalModel(m, "configmaps", &obj.Configmaps, UnmarshalV2ConfigMap)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationListNextMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2Project : V2Project struct
type V2Project struct {
	AccountID *string `json:"account_id,omitempty"`

	Created *string `json:"created,omitempty"`

	Crn *string `json:"crn,omitempty"`

	Details *string `json:"details,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Region *string `json:"region,omitempty"`

	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`
}

// UnmarshalV2Project unmarshals an instance of V2Project from the specified map of raw messages.
func UnmarshalV2Project(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2Project)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2ProjectList : V2ProjectList struct
type V2ProjectList struct {
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`

	Projects []V2Project `json:"projects,omitempty"`
}

// UnmarshalV2ProjectList unmarshals an instance of V2ProjectList from the specified map of raw messages.
func UnmarshalV2ProjectList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2ProjectList)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationListNextMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalV2Project)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2Reclamation : V2Reclamation struct
type V2Reclamation struct {
	AccountID *string `json:"account_id,omitempty"`

	Details *string `json:"details,omitempty"`

	ID *string `json:"id,omitempty"`

	ProjectID *string `json:"project_id,omitempty"`

	Reason *string `json:"reason,omitempty"`

	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	Status *string `json:"status,omitempty"`

	TargetTime *string `json:"target_time,omitempty"`

	Type *string `json:"type,omitempty"`
}

// UnmarshalV2Reclamation unmarshals an instance of V2Reclamation from the specified map of raw messages.
func UnmarshalV2Reclamation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2Reclamation)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_time", &obj.TargetTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2ReclamationList : V2ReclamationList struct
type V2ReclamationList struct {
	Reclamations []V2Reclamation `json:"reclamations,omitempty"`
}

// UnmarshalV2ReclamationList unmarshals an instance of V2ReclamationList from the specified map of raw messages.
func UnmarshalV2ReclamationList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2ReclamationList)
	err = core.UnmarshalModel(m, "reclamations", &obj.Reclamations, UnmarshalV2Reclamation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
