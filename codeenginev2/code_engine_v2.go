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
 * IBM OpenAPI SDK Code Generator Version: 3.54.0-af6d2126-20220803-151219
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

// ListProjects : List all projects
// List projects.
func (codeEngine *CodeEngineV2) ListProjects(listProjectsOptions *ListProjectsOptions) (result *V2ProjectList, response *core.DetailedResponse, err error) {
	return codeEngine.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *V2ProjectList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProjectsOptions, "listProjectsOptions")
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

	for headerName, headerValue := range listProjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListProjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// CreateProject : Create a Project
// Create a project.
func (codeEngine *CodeEngineV2) CreateProject(createProjectOptions *CreateProjectOptions) (result *V2Project, response *core.DetailedResponse, err error) {
	return codeEngine.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *V2Project, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProjectOptions, "createProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProjectOptions, "createProjectOptions")
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

	for headerName, headerValue := range createProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createProjectOptions.Name != nil {
		body["name"] = createProjectOptions.Name
	}
	if createProjectOptions.Region != nil {
		body["region"] = createProjectOptions.Region
	}
	if createProjectOptions.ResourceGroupID != nil {
		body["resource_group_id"] = createProjectOptions.ResourceGroupID
	}
	if createProjectOptions.Tags != nil {
		body["tags"] = createProjectOptions.Tags
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

// GetProject : Get a project
// Retrieve the project.
func (codeEngine *CodeEngineV2) GetProject(getProjectOptions *GetProjectOptions) (result *V2Project, response *core.DetailedResponse, err error) {
	return codeEngine.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *V2Project, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectOptions, "getProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectOptions, "getProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getProjectOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// DeleteProject : Delete a Project
// Delete a project.
func (codeEngine *CodeEngineV2) DeleteProject(deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteProjectWithContext(context.Background(), deleteProjectOptions)
}

// DeleteProjectWithContext is an alternate form of the DeleteProject method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteProjectWithContext(ctx context.Context, deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProjectOptions, "deleteProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProjectOptions, "deleteProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteProjectOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// ListConfigmaps : List configmaps
// List Configmaps.
func (codeEngine *CodeEngineV2) ListConfigmaps(listConfigmapsOptions *ListConfigmapsOptions) (result *V2ConfigMapList, response *core.DetailedResponse, err error) {
	return codeEngine.ListConfigmapsWithContext(context.Background(), listConfigmapsOptions)
}

// ListConfigmapsWithContext is an alternate form of the ListConfigmaps method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListConfigmapsWithContext(ctx context.Context, listConfigmapsOptions *ListConfigmapsOptions) (result *V2ConfigMapList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigmapsOptions, "listConfigmapsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigmapsOptions, "listConfigmapsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listConfigmapsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigmapsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListConfigmaps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// CreateConfigmap : Create a configmap
// Create a Configmap.
func (codeEngine *CodeEngineV2) CreateConfigmap(createConfigmapOptions *CreateConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.CreateConfigmapWithContext(context.Background(), createConfigmapOptions)
}

// CreateConfigmapWithContext is an alternate form of the CreateConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateConfigmapWithContext(ctx context.Context, createConfigmapOptions *CreateConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createConfigmapOptions, "createConfigmapOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createConfigmapOptions, "createConfigmapOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createConfigmapOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigmapOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateConfigmap")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createConfigmapOptions.Data != nil {
		body["data"] = createConfigmapOptions.Data
	}
	if createConfigmapOptions.Immutable != nil {
		body["immutable"] = createConfigmapOptions.Immutable
	}
	if createConfigmapOptions.Name != nil {
		body["name"] = createConfigmapOptions.Name
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

// GetConfigmap : Get a configmap
// Get a Configmap.
func (codeEngine *CodeEngineV2) GetConfigmap(getConfigmapOptions *GetConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.GetConfigmapWithContext(context.Background(), getConfigmapOptions)
}

// GetConfigmapWithContext is an alternate form of the GetConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetConfigmapWithContext(ctx context.Context, getConfigmapOptions *GetConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigmapOptions, "getConfigmapOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigmapOptions, "getConfigmapOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getConfigmapOptions.ProjectGuid,
		"configmap_name": *getConfigmapOptions.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigmapOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetConfigmap")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// DeleteConfigmap : Delete a configmap
// Delete a Configmap.
func (codeEngine *CodeEngineV2) DeleteConfigmap(deleteConfigmapOptions *DeleteConfigmapOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteConfigmapWithContext(context.Background(), deleteConfigmapOptions)
}

// DeleteConfigmapWithContext is an alternate form of the DeleteConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteConfigmapWithContext(ctx context.Context, deleteConfigmapOptions *DeleteConfigmapOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConfigmapOptions, "deleteConfigmapOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteConfigmapOptions, "deleteConfigmapOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteConfigmapOptions.ProjectGuid,
		"configmap_name": *deleteConfigmapOptions.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigmapOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteConfigmap")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// UpdateConfigmap : Update a configmap
// Update a Configmap.
func (codeEngine *CodeEngineV2) UpdateConfigmap(updateConfigmapOptions *UpdateConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateConfigmapWithContext(context.Background(), updateConfigmapOptions)
}

// UpdateConfigmapWithContext is an alternate form of the UpdateConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateConfigmapWithContext(ctx context.Context, updateConfigmapOptions *UpdateConfigmapOptions) (result *V2ConfigMap, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigmapOptions, "updateConfigmapOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigmapOptions, "updateConfigmapOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateConfigmapOptions.ProjectGuid,
		"configmap_name": *updateConfigmapOptions.ConfigmapName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/configmaps/{configmap_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigmapOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateConfigmap")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateConfigmapOptions.Data != nil {
		body["data"] = updateConfigmapOptions.Data
	}
	if updateConfigmapOptions.Immutable != nil {
		body["immutable"] = updateConfigmapOptions.Immutable
	}
	if updateConfigmapOptions.Name != nil {
		body["name"] = updateConfigmapOptions.Name
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

// ListReclamations : List all reclamations
// List reclamations.
func (codeEngine *CodeEngineV2) ListReclamations(listReclamationsOptions *ListReclamationsOptions) (result *V2ReclamationList, response *core.DetailedResponse, err error) {
	return codeEngine.ListReclamationsWithContext(context.Background(), listReclamationsOptions)
}

// ListReclamationsWithContext is an alternate form of the ListReclamations method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListReclamationsWithContext(ctx context.Context, listReclamationsOptions *ListReclamationsOptions) (result *V2ReclamationList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listReclamationsOptions, "listReclamationsOptions")
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

	for headerName, headerValue := range listReclamationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListReclamations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// GetReclamation : Get a reclamation
// Get a reclamation.
func (codeEngine *CodeEngineV2) GetReclamation(getReclamationOptions *GetReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.GetReclamationWithContext(context.Background(), getReclamationOptions)
}

// GetReclamationWithContext is an alternate form of the GetReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetReclamationWithContext(ctx context.Context, getReclamationOptions *GetReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReclamationOptions, "getReclamationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReclamationOptions, "getReclamationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getReclamationOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReclamationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetReclamation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// ReclaimReclamation : Reclaim a reclamation
// Reclaim a reclaimation to hard delete a project.
func (codeEngine *CodeEngineV2) ReclaimReclamation(reclaimReclamationOptions *ReclaimReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.ReclaimReclamationWithContext(context.Background(), reclaimReclamationOptions)
}

// ReclaimReclamationWithContext is an alternate form of the ReclaimReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) ReclaimReclamationWithContext(ctx context.Context, reclaimReclamationOptions *ReclaimReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reclaimReclamationOptions, "reclaimReclamationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reclaimReclamationOptions, "reclaimReclamationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *reclaimReclamationOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}/reclaim`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range reclaimReclamationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ReclaimReclamation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// RestoreReclamation : Restore a reclamation
// Restore a reclaimation which restores a soft-deleted project.
func (codeEngine *CodeEngineV2) RestoreReclamation(restoreReclamationOptions *RestoreReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.RestoreReclamationWithContext(context.Background(), restoreReclamationOptions)
}

// RestoreReclamationWithContext is an alternate form of the RestoreReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) RestoreReclamationWithContext(ctx context.Context, restoreReclamationOptions *RestoreReclamationOptions) (result *V2Reclamation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restoreReclamationOptions, "restoreReclamationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(restoreReclamationOptions, "restoreReclamationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *restoreReclamationOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/reclamations/{project_guid}/restore`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range restoreReclamationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "RestoreReclamation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// CreateConfigmapOptions : The CreateConfigmap options.
type CreateConfigmapOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Define whether it is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// Specify the Configmap name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigmapOptions : Instantiate CreateConfigmapOptions
func (*CodeEngineV2) NewCreateConfigmapOptions(projectGuid string) *CreateConfigmapOptions {
	return &CreateConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateConfigmapOptions) SetProjectGuid(projectGuid string) *CreateConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetData : Allow user to set Data
func (_options *CreateConfigmapOptions) SetData(data map[string]string) *CreateConfigmapOptions {
	_options.Data = data
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *CreateConfigmapOptions) SetImmutable(immutable bool) *CreateConfigmapOptions {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateConfigmapOptions) SetName(name string) *CreateConfigmapOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigmapOptions) SetHeaders(param map[string]string) *CreateConfigmapOptions {
	options.Headers = param
	return options
}

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The name of the project instance.
	Name *string `json:"name,omitempty"`

	// The deployment location where the project instance should be hosted (us-east, eu-de, ...).
	Region *string `json:"region,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Tags that are attached to the instance after provisioning. These tags can be searched and managed through the
	// Tagging API in IBM Cloud.
	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*CodeEngineV2) NewCreateProjectOptions() *CreateProjectOptions {
	return &CreateProjectOptions{}
}

// SetName : Allow user to set Name
func (_options *CreateProjectOptions) SetName(name string) *CreateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateProjectOptions) SetRegion(region string) *CreateProjectOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (_options *CreateProjectOptions) SetResourceGroupID(resourceGroupID string) *CreateProjectOptions {
	_options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateProjectOptions) SetTags(tags []string) *CreateProjectOptions {
	_options.Tags = tags
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectOptions) SetHeaders(param map[string]string) *CreateProjectOptions {
	options.Headers = param
	return options
}

// DeleteConfigmapOptions : The DeleteConfigmap options.
type DeleteConfigmapOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// ConfigMap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigmapOptions : Instantiate DeleteConfigmapOptions
func (*CodeEngineV2) NewDeleteConfigmapOptions(projectGuid string, configmapName string) *DeleteConfigmapOptions {
	return &DeleteConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteConfigmapOptions) SetProjectGuid(projectGuid string) *DeleteConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *DeleteConfigmapOptions) SetConfigmapName(configmapName string) *DeleteConfigmapOptions {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigmapOptions) SetHeaders(param map[string]string) *DeleteConfigmapOptions {
	options.Headers = param
	return options
}

// DeleteProjectOptions : The DeleteProject options.
type DeleteProjectOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectOptions : Instantiate DeleteProjectOptions
func (*CodeEngineV2) NewDeleteProjectOptions(projectGuid string) *DeleteProjectOptions {
	return &DeleteProjectOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteProjectOptions) SetProjectGuid(projectGuid string) *DeleteProjectOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectOptions) SetHeaders(param map[string]string) *DeleteProjectOptions {
	options.Headers = param
	return options
}

// GetConfigmapOptions : The GetConfigmap options.
type GetConfigmapOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// ConfigMap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigmapOptions : Instantiate GetConfigmapOptions
func (*CodeEngineV2) NewGetConfigmapOptions(projectGuid string, configmapName string) *GetConfigmapOptions {
	return &GetConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetConfigmapOptions) SetProjectGuid(projectGuid string) *GetConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *GetConfigmapOptions) SetConfigmapName(configmapName string) *GetConfigmapOptions {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigmapOptions) SetHeaders(param map[string]string) *GetConfigmapOptions {
	options.Headers = param
	return options
}

// GetProjectOptions : The GetProject options.
type GetProjectOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectOptions : Instantiate GetProjectOptions
func (*CodeEngineV2) NewGetProjectOptions(projectGuid string) *GetProjectOptions {
	return &GetProjectOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetProjectOptions) SetProjectGuid(projectGuid string) *GetProjectOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectOptions) SetHeaders(param map[string]string) *GetProjectOptions {
	options.Headers = param
	return options
}

// GetReclamationOptions : The GetReclamation options.
type GetReclamationOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReclamationOptions : Instantiate GetReclamationOptions
func (*CodeEngineV2) NewGetReclamationOptions(projectGuid string) *GetReclamationOptions {
	return &GetReclamationOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetReclamationOptions) SetProjectGuid(projectGuid string) *GetReclamationOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReclamationOptions) SetHeaders(param map[string]string) *GetReclamationOptions {
	options.Headers = param
	return options
}

// ListConfigmapsOptions : The ListConfigmaps options.
type ListConfigmapsOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigmapsOptions : Instantiate ListConfigmapsOptions
func (*CodeEngineV2) NewListConfigmapsOptions(projectGuid string) *ListConfigmapsOptions {
	return &ListConfigmapsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListConfigmapsOptions) SetProjectGuid(projectGuid string) *ListConfigmapsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigmapsOptions) SetHeaders(param map[string]string) *ListConfigmapsOptions {
	options.Headers = param
	return options
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*CodeEngineV2) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// ListReclamationsOptions : The ListReclamations options.
type ListReclamationsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListReclamationsOptions : Instantiate ListReclamationsOptions
func (*CodeEngineV2) NewListReclamationsOptions() *ListReclamationsOptions {
	return &ListReclamationsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListReclamationsOptions) SetHeaders(param map[string]string) *ListReclamationsOptions {
	options.Headers = param
	return options
}

// ReclaimReclamationOptions : The ReclaimReclamation options.
type ReclaimReclamationOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReclaimReclamationOptions : Instantiate ReclaimReclamationOptions
func (*CodeEngineV2) NewReclaimReclamationOptions(projectGuid string) *ReclaimReclamationOptions {
	return &ReclaimReclamationOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ReclaimReclamationOptions) SetProjectGuid(projectGuid string) *ReclaimReclamationOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReclaimReclamationOptions) SetHeaders(param map[string]string) *ReclaimReclamationOptions {
	options.Headers = param
	return options
}

// RestoreReclamationOptions : The RestoreReclamation options.
type RestoreReclamationOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestoreReclamationOptions : Instantiate RestoreReclamationOptions
func (*CodeEngineV2) NewRestoreReclamationOptions(projectGuid string) *RestoreReclamationOptions {
	return &RestoreReclamationOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *RestoreReclamationOptions) SetProjectGuid(projectGuid string) *RestoreReclamationOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RestoreReclamationOptions) SetHeaders(param map[string]string) *RestoreReclamationOptions {
	options.Headers = param
	return options
}

// UpdateConfigmapOptions : The UpdateConfigmap options.
type UpdateConfigmapOptions struct {
	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Configmap name.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Define whether it is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// Specify the Configmap name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigmapOptions : Instantiate UpdateConfigmapOptions
func (*CodeEngineV2) NewUpdateConfigmapOptions(projectGuid string, configmapName string) *UpdateConfigmapOptions {
	return &UpdateConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigmapName: core.StringPtr(configmapName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateConfigmapOptions) SetProjectGuid(projectGuid string) *UpdateConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigmapName : Allow user to set ConfigmapName
func (_options *UpdateConfigmapOptions) SetConfigmapName(configmapName string) *UpdateConfigmapOptions {
	_options.ConfigmapName = core.StringPtr(configmapName)
	return _options
}

// SetData : Allow user to set Data
func (_options *UpdateConfigmapOptions) SetData(data map[string]string) *UpdateConfigmapOptions {
	_options.Data = data
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *UpdateConfigmapOptions) SetImmutable(immutable bool) *UpdateConfigmapOptions {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateConfigmapOptions) SetName(name string) *UpdateConfigmapOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigmapOptions) SetHeaders(param map[string]string) *UpdateConfigmapOptions {
	options.Headers = param
	return options
}

// PaginationListNextMetadata : PaginationListNextMetadata struct
type PaginationListNextMetadata struct {
	// URL that points to the next page.
	Href *string `json:"href,omitempty"`

	// Token.
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
	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Define whether it is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the Configmap.
	Name *string `json:"name,omitempty"`

	// The type of Code Engine resource.
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
	// List of Configmaps.
	Configmaps []V2ConfigMap `json:"configmaps,omitempty"`

	// Max number of resources per page.
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
	// An alpha-numeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The date when the project was created.
	Created *string `json:"created,omitempty"`

	// The ID associated with the project instance.
	Crn *string `json:"crn,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The GUID of the project instance.
	ID *string `json:"id,omitempty"`

	// Specify the project name.
	Name *string `json:"name,omitempty"`

	// Reason that provides some more context on the given status.
	Reason *string `json:"reason,omitempty"`

	// Specify the id of the regin (us-south, eu-de).
	Region *string `json:"region,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The type of Code Engine resource.
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
	// Max number of resources per page.
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of projects.
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
	// An alpha-numeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The ID associated with the reclamation.
	ID *string `json:"id,omitempty"`

	// The ID of the Code Engine project resource instance.
	ProjectID *string `json:"project_id,omitempty"`

	// Reason that provides some more context on the given status.
	Reason *string `json:"reason,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The target time that the reclamation retention period end.
	TargetTime *string `json:"target_time,omitempty"`

	// The type of Code Engine resource.
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
	// List of reclamations.
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
