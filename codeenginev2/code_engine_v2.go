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
// List all projects in the current resource group.
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

	if listProjectsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listProjectsOptions.Limit))
	}
	if listProjectsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listProjectsOptions.Start))
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

// CreateProject : Create a project
// Create a project in the current resource group.
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
// Display the details of a single project.
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

// DeleteProject : Delete a project
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

// ListBuilds : List builds
// List all builds in a project.
func (codeEngine *CodeEngineV2) ListBuilds(listBuildsOptions *ListBuildsOptions) (result *V2BuildList, response *core.DetailedResponse, err error) {
	return codeEngine.ListBuildsWithContext(context.Background(), listBuildsOptions)
}

// ListBuildsWithContext is an alternate form of the ListBuilds method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListBuildsWithContext(ctx context.Context, listBuildsOptions *ListBuildsOptions) (result *V2BuildList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBuildsOptions, "listBuildsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBuildsOptions, "listBuildsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listBuildsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/builds`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBuildsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListBuilds")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listBuildsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listBuildsOptions.Limit))
	}
	if listBuildsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listBuildsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2BuildList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBuild : Create a build
// Create a build.
func (codeEngine *CodeEngineV2) CreateBuild(createBuildOptions *CreateBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	return codeEngine.CreateBuildWithContext(context.Background(), createBuildOptions)
}

// CreateBuildWithContext is an alternate form of the CreateBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateBuildWithContext(ctx context.Context, createBuildOptions *CreateBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBuildOptions, "createBuildOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBuildOptions, "createBuildOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createBuildOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/builds`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBuildOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateBuild")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createBuildOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = createBuildOptions.CeOwnerReference
	}
	if createBuildOptions.Dockerfile != nil {
		body["dockerfile"] = createBuildOptions.Dockerfile
	}
	if createBuildOptions.Name != nil {
		body["name"] = createBuildOptions.Name
	}
	if createBuildOptions.OutputImage != nil {
		body["output_image"] = createBuildOptions.OutputImage
	}
	if createBuildOptions.OutputSecret != nil {
		body["output_secret"] = createBuildOptions.OutputSecret
	}
	if createBuildOptions.SourceContextDir != nil {
		body["source_context_dir"] = createBuildOptions.SourceContextDir
	}
	if createBuildOptions.SourceRevision != nil {
		body["source_revision"] = createBuildOptions.SourceRevision
	}
	if createBuildOptions.SourceSecret != nil {
		body["source_secret"] = createBuildOptions.SourceSecret
	}
	if createBuildOptions.SourceType != nil {
		body["source_type"] = createBuildOptions.SourceType
	}
	if createBuildOptions.SourceURL != nil {
		body["source_url"] = createBuildOptions.SourceURL
	}
	if createBuildOptions.StrategyName != nil {
		body["strategy_name"] = createBuildOptions.StrategyName
	}
	if createBuildOptions.StrategySize != nil {
		body["strategy_size"] = createBuildOptions.StrategySize
	}
	if createBuildOptions.Timeout != nil {
		body["timeout"] = createBuildOptions.Timeout
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Build)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBuild : Get a build
// Display the details of a build.
func (codeEngine *CodeEngineV2) GetBuild(getBuildOptions *GetBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	return codeEngine.GetBuildWithContext(context.Background(), getBuildOptions)
}

// GetBuildWithContext is an alternate form of the GetBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetBuildWithContext(ctx context.Context, getBuildOptions *GetBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBuildOptions, "getBuildOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBuildOptions, "getBuildOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getBuildOptions.ProjectGuid,
		"build_name": *getBuildOptions.BuildName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/builds/{build_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBuildOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetBuild")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Build)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteBuild : Delete a build
// Delete a build.
func (codeEngine *CodeEngineV2) DeleteBuild(deleteBuildOptions *DeleteBuildOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteBuildWithContext(context.Background(), deleteBuildOptions)
}

// DeleteBuildWithContext is an alternate form of the DeleteBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteBuildWithContext(ctx context.Context, deleteBuildOptions *DeleteBuildOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBuildOptions, "deleteBuildOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBuildOptions, "deleteBuildOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteBuildOptions.ProjectGuid,
		"build_name": *deleteBuildOptions.BuildName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/builds/{build_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBuildOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteBuild")
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

// UpdateBuild : Update a build
// Update a build.
func (codeEngine *CodeEngineV2) UpdateBuild(updateBuildOptions *UpdateBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateBuildWithContext(context.Background(), updateBuildOptions)
}

// UpdateBuildWithContext is an alternate form of the UpdateBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateBuildWithContext(ctx context.Context, updateBuildOptions *UpdateBuildOptions) (result *V2Build, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBuildOptions, "updateBuildOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBuildOptions, "updateBuildOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateBuildOptions.ProjectGuid,
		"build_name": *updateBuildOptions.BuildName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/builds/{build_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBuildOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateBuild")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateBuildOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = updateBuildOptions.CeOwnerReference
	}
	if updateBuildOptions.Dockerfile != nil {
		body["dockerfile"] = updateBuildOptions.Dockerfile
	}
	if updateBuildOptions.Name != nil {
		body["name"] = updateBuildOptions.Name
	}
	if updateBuildOptions.OutputImage != nil {
		body["output_image"] = updateBuildOptions.OutputImage
	}
	if updateBuildOptions.OutputSecret != nil {
		body["output_secret"] = updateBuildOptions.OutputSecret
	}
	if updateBuildOptions.SourceContextDir != nil {
		body["source_context_dir"] = updateBuildOptions.SourceContextDir
	}
	if updateBuildOptions.SourceRevision != nil {
		body["source_revision"] = updateBuildOptions.SourceRevision
	}
	if updateBuildOptions.SourceSecret != nil {
		body["source_secret"] = updateBuildOptions.SourceSecret
	}
	if updateBuildOptions.SourceType != nil {
		body["source_type"] = updateBuildOptions.SourceType
	}
	if updateBuildOptions.SourceURL != nil {
		body["source_url"] = updateBuildOptions.SourceURL
	}
	if updateBuildOptions.StrategyName != nil {
		body["strategy_name"] = updateBuildOptions.StrategyName
	}
	if updateBuildOptions.StrategySize != nil {
		body["strategy_size"] = updateBuildOptions.StrategySize
	}
	if updateBuildOptions.Timeout != nil {
		body["timeout"] = updateBuildOptions.Timeout
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Build)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBuildruns : List build runs
// List all build runs in a project.
func (codeEngine *CodeEngineV2) ListBuildruns(listBuildrunsOptions *ListBuildrunsOptions) (result *V2BuildRunList, response *core.DetailedResponse, err error) {
	return codeEngine.ListBuildrunsWithContext(context.Background(), listBuildrunsOptions)
}

// ListBuildrunsWithContext is an alternate form of the ListBuildruns method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListBuildrunsWithContext(ctx context.Context, listBuildrunsOptions *ListBuildrunsOptions) (result *V2BuildRunList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBuildrunsOptions, "listBuildrunsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBuildrunsOptions, "listBuildrunsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listBuildrunsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/buildruns`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBuildrunsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListBuildruns")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listBuildrunsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listBuildrunsOptions.Limit))
	}
	if listBuildrunsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listBuildrunsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2BuildRunList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBuildrun : Create a build run
// Create a build run.
func (codeEngine *CodeEngineV2) CreateBuildrun(createBuildrunOptions *CreateBuildrunOptions) (result *V2BuildRun, response *core.DetailedResponse, err error) {
	return codeEngine.CreateBuildrunWithContext(context.Background(), createBuildrunOptions)
}

// CreateBuildrunWithContext is an alternate form of the CreateBuildrun method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateBuildrunWithContext(ctx context.Context, createBuildrunOptions *CreateBuildrunOptions) (result *V2BuildRun, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBuildrunOptions, "createBuildrunOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBuildrunOptions, "createBuildrunOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createBuildrunOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/buildruns`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBuildrunOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateBuildrun")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createBuildrunOptions.AppRevision != nil {
		body["app_revision"] = createBuildrunOptions.AppRevision
	}
	if createBuildrunOptions.Build != nil {
		body["build"] = createBuildrunOptions.Build
	}
	if createBuildrunOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = createBuildrunOptions.CeOwnerReference
	}
	if createBuildrunOptions.Dockerfile != nil {
		body["dockerfile"] = createBuildrunOptions.Dockerfile
	}
	if createBuildrunOptions.Name != nil {
		body["name"] = createBuildrunOptions.Name
	}
	if createBuildrunOptions.OutputImage != nil {
		body["output_image"] = createBuildrunOptions.OutputImage
	}
	if createBuildrunOptions.OutputSecret != nil {
		body["output_secret"] = createBuildrunOptions.OutputSecret
	}
	if createBuildrunOptions.ServiceAccount != nil {
		body["service_account"] = createBuildrunOptions.ServiceAccount
	}
	if createBuildrunOptions.SourceContextDir != nil {
		body["source_context_dir"] = createBuildrunOptions.SourceContextDir
	}
	if createBuildrunOptions.SourceRevision != nil {
		body["source_revision"] = createBuildrunOptions.SourceRevision
	}
	if createBuildrunOptions.SourceSecret != nil {
		body["source_secret"] = createBuildrunOptions.SourceSecret
	}
	if createBuildrunOptions.SourceType != nil {
		body["source_type"] = createBuildrunOptions.SourceType
	}
	if createBuildrunOptions.SourceURL != nil {
		body["source_url"] = createBuildrunOptions.SourceURL
	}
	if createBuildrunOptions.StrategyName != nil {
		body["strategy_name"] = createBuildrunOptions.StrategyName
	}
	if createBuildrunOptions.StrategySize != nil {
		body["strategy_size"] = createBuildrunOptions.StrategySize
	}
	if createBuildrunOptions.Timeout != nil {
		body["timeout"] = createBuildrunOptions.Timeout
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2BuildRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBuildrun : Get a build run
// Display the details of a build run.
func (codeEngine *CodeEngineV2) GetBuildrun(getBuildrunOptions *GetBuildrunOptions) (result *V2BuildRun, response *core.DetailedResponse, err error) {
	return codeEngine.GetBuildrunWithContext(context.Background(), getBuildrunOptions)
}

// GetBuildrunWithContext is an alternate form of the GetBuildrun method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetBuildrunWithContext(ctx context.Context, getBuildrunOptions *GetBuildrunOptions) (result *V2BuildRun, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBuildrunOptions, "getBuildrunOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBuildrunOptions, "getBuildrunOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getBuildrunOptions.ProjectGuid,
		"buildrun_name": *getBuildrunOptions.BuildrunName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/buildruns/{buildrun_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBuildrunOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetBuildrun")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2BuildRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteBuildrun : Delete a buildrun
// Delete a buildrun.
func (codeEngine *CodeEngineV2) DeleteBuildrun(deleteBuildrunOptions *DeleteBuildrunOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteBuildrunWithContext(context.Background(), deleteBuildrunOptions)
}

// DeleteBuildrunWithContext is an alternate form of the DeleteBuildrun method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteBuildrunWithContext(ctx context.Context, deleteBuildrunOptions *DeleteBuildrunOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBuildrunOptions, "deleteBuildrunOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBuildrunOptions, "deleteBuildrunOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteBuildrunOptions.ProjectGuid,
		"buildrun_name": *deleteBuildrunOptions.BuildrunName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/buildruns/{buildrun_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBuildrunOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteBuildrun")
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
// List all configmaps in a project.
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

	if listConfigmapsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listConfigmapsOptions.Limit))
	}
	if listConfigmapsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listConfigmapsOptions.Start))
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

// CreateConfigmap : Create a configmap
// Create a configmap.
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
// Display the details of a configmap.
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
// Delete a configmap.
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
// Update a configmap.
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

// ListSecrets : List secret
// List secrets.
func (codeEngine *CodeEngineV2) ListSecrets(listSecretsOptions *ListSecretsOptions) (result *V2SecretList, response *core.DetailedResponse, err error) {
	return codeEngine.ListSecretsWithContext(context.Background(), listSecretsOptions)
}

// ListSecretsWithContext is an alternate form of the ListSecrets method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListSecretsWithContext(ctx context.Context, listSecretsOptions *ListSecretsOptions) (result *V2SecretList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSecretsOptions, "listSecretsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listSecretsOptions, "listSecretsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listSecretsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/secrets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSecretsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListSecrets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSecretsOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*listSecretsOptions.RefreshToken))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2SecretList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSecret : Create a secret
// Create a secret.
func (codeEngine *CodeEngineV2) CreateSecret(createSecretOptions *CreateSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	return codeEngine.CreateSecretWithContext(context.Background(), createSecretOptions)
}

// CreateSecretWithContext is an alternate form of the CreateSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateSecretWithContext(ctx context.Context, createSecretOptions *CreateSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSecretOptions, "createSecretOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSecretOptions, "createSecretOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createSecretOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/secrets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSecretOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateSecret")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSecretOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*createSecretOptions.RefreshToken))
	}

	body := make(map[string]interface{})
	if createSecretOptions.BindingSecretRef != nil {
		body["binding_secret_ref"] = createSecretOptions.BindingSecretRef
	}
	if createSecretOptions.CeComponents != nil {
		body["ce_components"] = createSecretOptions.CeComponents
	}
	if createSecretOptions.Created != nil {
		body["created"] = createSecretOptions.Created
	}
	if createSecretOptions.Data != nil {
		body["data"] = createSecretOptions.Data
	}
	if createSecretOptions.Format != nil {
		body["format"] = createSecretOptions.Format
	}
	if createSecretOptions.ID != nil {
		body["id"] = createSecretOptions.ID
	}
	if createSecretOptions.Immutable != nil {
		body["immutable"] = createSecretOptions.Immutable
	}
	if createSecretOptions.Name != nil {
		body["name"] = createSecretOptions.Name
	}
	if createSecretOptions.ResourceID != nil {
		body["resource_id"] = createSecretOptions.ResourceID
	}
	if createSecretOptions.ResourceType != nil {
		body["resource_type"] = createSecretOptions.ResourceType
	}
	if createSecretOptions.ResourcekeyID != nil {
		body["resourcekey_id"] = createSecretOptions.ResourcekeyID
	}
	if createSecretOptions.Role != nil {
		body["role"] = createSecretOptions.Role
	}
	if createSecretOptions.ServiceidCrn != nil {
		body["serviceid_crn"] = createSecretOptions.ServiceidCrn
	}
	if createSecretOptions.Target != nil {
		body["target"] = createSecretOptions.Target
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Secret)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSecret : Get a secret
// Get a secret.
func (codeEngine *CodeEngineV2) GetSecret(getSecretOptions *GetSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	return codeEngine.GetSecretWithContext(context.Background(), getSecretOptions)
}

// GetSecretWithContext is an alternate form of the GetSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetSecretWithContext(ctx context.Context, getSecretOptions *GetSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSecretOptions, "getSecretOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSecretOptions, "getSecretOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getSecretOptions.ProjectGuid,
		"secret_name": *getSecretOptions.SecretName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/secrets/{secret_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSecretOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetSecret")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSecretOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*getSecretOptions.RefreshToken))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Secret)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSecret : Delete a secret
// Delete a secret.
func (codeEngine *CodeEngineV2) DeleteSecret(deleteSecretOptions *DeleteSecretOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteSecretWithContext(context.Background(), deleteSecretOptions)
}

// DeleteSecretWithContext is an alternate form of the DeleteSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteSecretWithContext(ctx context.Context, deleteSecretOptions *DeleteSecretOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSecretOptions, "deleteSecretOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSecretOptions, "deleteSecretOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteSecretOptions.ProjectGuid,
		"secret_name": *deleteSecretOptions.SecretName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/secrets/{secret_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSecretOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteSecret")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSecretOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*deleteSecretOptions.RefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// UpdateSecret : Update a secret
// Update a secret.
func (codeEngine *CodeEngineV2) UpdateSecret(updateSecretOptions *UpdateSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateSecretWithContext(context.Background(), updateSecretOptions)
}

// UpdateSecretWithContext is an alternate form of the UpdateSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateSecretWithContext(ctx context.Context, updateSecretOptions *UpdateSecretOptions) (result *V2Secret, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSecretOptions, "updateSecretOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSecretOptions, "updateSecretOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateSecretOptions.ProjectGuid,
		"secret_name": *updateSecretOptions.SecretName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/secrets/{secret_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSecretOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateSecret")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateSecretOptions.RefreshToken != nil {
		builder.AddHeader("Refresh-Token", fmt.Sprint(*updateSecretOptions.RefreshToken))
	}

	body := make(map[string]interface{})
	if updateSecretOptions.BindingSecretRef != nil {
		body["binding_secret_ref"] = updateSecretOptions.BindingSecretRef
	}
	if updateSecretOptions.CeComponents != nil {
		body["ce_components"] = updateSecretOptions.CeComponents
	}
	if updateSecretOptions.Created != nil {
		body["created"] = updateSecretOptions.Created
	}
	if updateSecretOptions.Data != nil {
		body["data"] = updateSecretOptions.Data
	}
	if updateSecretOptions.Format != nil {
		body["format"] = updateSecretOptions.Format
	}
	if updateSecretOptions.ID != nil {
		body["id"] = updateSecretOptions.ID
	}
	if updateSecretOptions.Immutable != nil {
		body["immutable"] = updateSecretOptions.Immutable
	}
	if updateSecretOptions.Name != nil {
		body["name"] = updateSecretOptions.Name
	}
	if updateSecretOptions.ResourceID != nil {
		body["resource_id"] = updateSecretOptions.ResourceID
	}
	if updateSecretOptions.ResourceType != nil {
		body["resource_type"] = updateSecretOptions.ResourceType
	}
	if updateSecretOptions.ResourcekeyID != nil {
		body["resourcekey_id"] = updateSecretOptions.ResourcekeyID
	}
	if updateSecretOptions.Role != nil {
		body["role"] = updateSecretOptions.Role
	}
	if updateSecretOptions.ServiceidCrn != nil {
		body["serviceid_crn"] = updateSecretOptions.ServiceidCrn
	}
	if updateSecretOptions.Target != nil {
		body["target"] = updateSecretOptions.Target
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalV2Secret)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListReclamations : List all reclamations
// List all project reclamations.
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

	if listReclamationsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listReclamationsOptions.Limit))
	}
	if listReclamationsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listReclamationsOptions.Start))
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

// ReclaimReclamation : Delete a reclamation
// Delete a project reclamation to permanently delete the project.
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

// RestoreReclamation : Restore a project reclamation
// Restore a project reclamation. Projects that are soft-deleted can be restored within 7 days.
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

// CreateBuildOptions : The CreateBuild options.
type CreateBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The path to the Dockerfile that is used for build strategies for building an image.
	Dockerfile *string `json:"dockerfile,omitempty"`

	// The name of the build. Use a name that is unique within the project.
	Name *string `json:"name,omitempty"`

	// The name of the image.
	OutputImage *string `json:"output_image,omitempty"`

	// The secret that is required to access the image registry.
	OutputSecret *string `json:"output_secret,omitempty"`

	// The directory in the repository that contains the buildpacks file or the Dockerfile.
	SourceContextDir *string `json:"source_context_dir,omitempty"`

	// The commit, tag, or branch in the source repository to pull.
	SourceRevision *string `json:"source_revision,omitempty"`

	// The name of the secret that is required to access the repository source.
	SourceSecret *string `json:"source_secret,omitempty"`

	// Specifies the type of source to determine if your build source is in a repository or based on local source code.
	SourceType *string `json:"source_type,omitempty"`

	// The URL of the repository.
	SourceURL *string `json:"source_url,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateBuildOptions : Instantiate CreateBuildOptions
func (*CodeEngineV2) NewCreateBuildOptions(projectGuid string) *CreateBuildOptions {
	return &CreateBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateBuildOptions) SetProjectGuid(projectGuid string) *CreateBuildOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetCeOwnerReference : Allow user to set CeOwnerReference
func (_options *CreateBuildOptions) SetCeOwnerReference(ceOwnerReference string) *CreateBuildOptions {
	_options.CeOwnerReference = core.StringPtr(ceOwnerReference)
	return _options
}

// SetDockerfile : Allow user to set Dockerfile
func (_options *CreateBuildOptions) SetDockerfile(dockerfile string) *CreateBuildOptions {
	_options.Dockerfile = core.StringPtr(dockerfile)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateBuildOptions) SetName(name string) *CreateBuildOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetOutputImage : Allow user to set OutputImage
func (_options *CreateBuildOptions) SetOutputImage(outputImage string) *CreateBuildOptions {
	_options.OutputImage = core.StringPtr(outputImage)
	return _options
}

// SetOutputSecret : Allow user to set OutputSecret
func (_options *CreateBuildOptions) SetOutputSecret(outputSecret string) *CreateBuildOptions {
	_options.OutputSecret = core.StringPtr(outputSecret)
	return _options
}

// SetSourceContextDir : Allow user to set SourceContextDir
func (_options *CreateBuildOptions) SetSourceContextDir(sourceContextDir string) *CreateBuildOptions {
	_options.SourceContextDir = core.StringPtr(sourceContextDir)
	return _options
}

// SetSourceRevision : Allow user to set SourceRevision
func (_options *CreateBuildOptions) SetSourceRevision(sourceRevision string) *CreateBuildOptions {
	_options.SourceRevision = core.StringPtr(sourceRevision)
	return _options
}

// SetSourceSecret : Allow user to set SourceSecret
func (_options *CreateBuildOptions) SetSourceSecret(sourceSecret string) *CreateBuildOptions {
	_options.SourceSecret = core.StringPtr(sourceSecret)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *CreateBuildOptions) SetSourceType(sourceType string) *CreateBuildOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetSourceURL : Allow user to set SourceURL
func (_options *CreateBuildOptions) SetSourceURL(sourceURL string) *CreateBuildOptions {
	_options.SourceURL = core.StringPtr(sourceURL)
	return _options
}

// SetStrategyName : Allow user to set StrategyName
func (_options *CreateBuildOptions) SetStrategyName(strategyName string) *CreateBuildOptions {
	_options.StrategyName = core.StringPtr(strategyName)
	return _options
}

// SetStrategySize : Allow user to set StrategySize
func (_options *CreateBuildOptions) SetStrategySize(strategySize string) *CreateBuildOptions {
	_options.StrategySize = core.StringPtr(strategySize)
	return _options
}

// SetTimeout : Allow user to set Timeout
func (_options *CreateBuildOptions) SetTimeout(timeout int64) *CreateBuildOptions {
	_options.Timeout = core.Int64Ptr(timeout)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBuildOptions) SetHeaders(param map[string]string) *CreateBuildOptions {
	options.Headers = param
	return options
}

// CreateBuildrunOptions : The CreateBuildrun options.
type CreateBuildrunOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of the app revision with which this build run is associated.
	AppRevision *string `json:"app_revision,omitempty"`

	// The name of the build on which this build run is associated.
	Build *string `json:"build,omitempty"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The path to the Dockerfile that is used for build strategies for building an image.
	Dockerfile *string `json:"dockerfile,omitempty"`

	// Name of the build run.
	Name *string `json:"name,omitempty"`

	// The name of the image.
	OutputImage *string `json:"output_image,omitempty"`

	// The secret that is required to access the image registry.
	OutputSecret *string `json:"output_secret,omitempty"`

	// ServiceAccount refers to the serviceaccount which is used for resource control.
	ServiceAccount *string `json:"service_account,omitempty"`

	// The directory in the repository that contains the buildpacks file or the Dockerfile.
	SourceContextDir *string `json:"source_context_dir,omitempty"`

	// The commit, tag, or branch in the source repository to pull.
	SourceRevision *string `json:"source_revision,omitempty"`

	// The name of the secret that is required to access the repository source.
	SourceSecret *string `json:"source_secret,omitempty"`

	// Specifies the type of source to determine if your build source is in a repository or based on local source code.
	SourceType *string `json:"source_type,omitempty"`

	// The URL of the repository.
	SourceURL *string `json:"source_url,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateBuildrunOptions : Instantiate CreateBuildrunOptions
func (*CodeEngineV2) NewCreateBuildrunOptions(projectGuid string) *CreateBuildrunOptions {
	return &CreateBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateBuildrunOptions) SetProjectGuid(projectGuid string) *CreateBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppRevision : Allow user to set AppRevision
func (_options *CreateBuildrunOptions) SetAppRevision(appRevision string) *CreateBuildrunOptions {
	_options.AppRevision = core.StringPtr(appRevision)
	return _options
}

// SetBuild : Allow user to set Build
func (_options *CreateBuildrunOptions) SetBuild(build string) *CreateBuildrunOptions {
	_options.Build = core.StringPtr(build)
	return _options
}

// SetCeOwnerReference : Allow user to set CeOwnerReference
func (_options *CreateBuildrunOptions) SetCeOwnerReference(ceOwnerReference string) *CreateBuildrunOptions {
	_options.CeOwnerReference = core.StringPtr(ceOwnerReference)
	return _options
}

// SetDockerfile : Allow user to set Dockerfile
func (_options *CreateBuildrunOptions) SetDockerfile(dockerfile string) *CreateBuildrunOptions {
	_options.Dockerfile = core.StringPtr(dockerfile)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateBuildrunOptions) SetName(name string) *CreateBuildrunOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetOutputImage : Allow user to set OutputImage
func (_options *CreateBuildrunOptions) SetOutputImage(outputImage string) *CreateBuildrunOptions {
	_options.OutputImage = core.StringPtr(outputImage)
	return _options
}

// SetOutputSecret : Allow user to set OutputSecret
func (_options *CreateBuildrunOptions) SetOutputSecret(outputSecret string) *CreateBuildrunOptions {
	_options.OutputSecret = core.StringPtr(outputSecret)
	return _options
}

// SetServiceAccount : Allow user to set ServiceAccount
func (_options *CreateBuildrunOptions) SetServiceAccount(serviceAccount string) *CreateBuildrunOptions {
	_options.ServiceAccount = core.StringPtr(serviceAccount)
	return _options
}

// SetSourceContextDir : Allow user to set SourceContextDir
func (_options *CreateBuildrunOptions) SetSourceContextDir(sourceContextDir string) *CreateBuildrunOptions {
	_options.SourceContextDir = core.StringPtr(sourceContextDir)
	return _options
}

// SetSourceRevision : Allow user to set SourceRevision
func (_options *CreateBuildrunOptions) SetSourceRevision(sourceRevision string) *CreateBuildrunOptions {
	_options.SourceRevision = core.StringPtr(sourceRevision)
	return _options
}

// SetSourceSecret : Allow user to set SourceSecret
func (_options *CreateBuildrunOptions) SetSourceSecret(sourceSecret string) *CreateBuildrunOptions {
	_options.SourceSecret = core.StringPtr(sourceSecret)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *CreateBuildrunOptions) SetSourceType(sourceType string) *CreateBuildrunOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetSourceURL : Allow user to set SourceURL
func (_options *CreateBuildrunOptions) SetSourceURL(sourceURL string) *CreateBuildrunOptions {
	_options.SourceURL = core.StringPtr(sourceURL)
	return _options
}

// SetStrategyName : Allow user to set StrategyName
func (_options *CreateBuildrunOptions) SetStrategyName(strategyName string) *CreateBuildrunOptions {
	_options.StrategyName = core.StringPtr(strategyName)
	return _options
}

// SetStrategySize : Allow user to set StrategySize
func (_options *CreateBuildrunOptions) SetStrategySize(strategySize string) *CreateBuildrunOptions {
	_options.StrategySize = core.StringPtr(strategySize)
	return _options
}

// SetTimeout : Allow user to set Timeout
func (_options *CreateBuildrunOptions) SetTimeout(timeout int64) *CreateBuildrunOptions {
	_options.Timeout = core.Int64Ptr(timeout)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBuildrunOptions) SetHeaders(param map[string]string) *CreateBuildrunOptions {
	options.Headers = param
	return options
}

// CreateConfigmapOptions : The CreateConfigmap options.
type CreateConfigmapOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// Indicates that the key-value pair cannot be edited.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the configmap. Use a name that is unique within the project.
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
	// The name of the project. Use a name that is unique to your region.
	Name *string `json:"name,omitempty"`

	// The region for your project deployment.
	Region *string `json:"region,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// A list of labels to assign to your project. You can manage tags through the Tagging API in IBM Cloud.
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

// CreateSecretOptions : The CreateSecret options.
type CreateSecretOptions struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Name of the secret.
	BindingSecretRef *string `json:"binding_secret_ref,omitempty"`

	// List of bound CE Components.
	CeComponents []string `json:"ce_components,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Define whether the secret is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the Secret.
	Name *string `json:"name,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associtaed with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential (resource key) associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Role of the service credential (resource key).
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential (resource key).
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Specify the target of the secret (aka how the secret will be used) label format: "target_<target>: <target>".
	Target *string `json:"target,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSecretOptions : Instantiate CreateSecretOptions
func (*CodeEngineV2) NewCreateSecretOptions(refreshToken string, projectGuid string) *CreateSecretOptions {
	return &CreateSecretOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *CreateSecretOptions) SetRefreshToken(refreshToken string) *CreateSecretOptions {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateSecretOptions) SetProjectGuid(projectGuid string) *CreateSecretOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBindingSecretRef : Allow user to set BindingSecretRef
func (_options *CreateSecretOptions) SetBindingSecretRef(bindingSecretRef string) *CreateSecretOptions {
	_options.BindingSecretRef = core.StringPtr(bindingSecretRef)
	return _options
}

// SetCeComponents : Allow user to set CeComponents
func (_options *CreateSecretOptions) SetCeComponents(ceComponents []string) *CreateSecretOptions {
	_options.CeComponents = ceComponents
	return _options
}

// SetCreated : Allow user to set Created
func (_options *CreateSecretOptions) SetCreated(created string) *CreateSecretOptions {
	_options.Created = core.StringPtr(created)
	return _options
}

// SetData : Allow user to set Data
func (_options *CreateSecretOptions) SetData(data map[string]string) *CreateSecretOptions {
	_options.Data = data
	return _options
}

// SetFormat : Allow user to set Format
func (_options *CreateSecretOptions) SetFormat(format string) *CreateSecretOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateSecretOptions) SetID(id string) *CreateSecretOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *CreateSecretOptions) SetImmutable(immutable bool) *CreateSecretOptions {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateSecretOptions) SetName(name string) *CreateSecretOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceID : Allow user to set ResourceID
func (_options *CreateSecretOptions) SetResourceID(resourceID string) *CreateSecretOptions {
	_options.ResourceID = core.StringPtr(resourceID)
	return _options
}

// SetResourceType : Allow user to set ResourceType
func (_options *CreateSecretOptions) SetResourceType(resourceType string) *CreateSecretOptions {
	_options.ResourceType = core.StringPtr(resourceType)
	return _options
}

// SetResourcekeyID : Allow user to set ResourcekeyID
func (_options *CreateSecretOptions) SetResourcekeyID(resourcekeyID string) *CreateSecretOptions {
	_options.ResourcekeyID = core.StringPtr(resourcekeyID)
	return _options
}

// SetRole : Allow user to set Role
func (_options *CreateSecretOptions) SetRole(role string) *CreateSecretOptions {
	_options.Role = core.StringPtr(role)
	return _options
}

// SetServiceidCrn : Allow user to set ServiceidCrn
func (_options *CreateSecretOptions) SetServiceidCrn(serviceidCrn string) *CreateSecretOptions {
	_options.ServiceidCrn = core.StringPtr(serviceidCrn)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *CreateSecretOptions) SetTarget(target string) *CreateSecretOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSecretOptions) SetHeaders(param map[string]string) *CreateSecretOptions {
	options.Headers = param
	return options
}

// DeleteBuildOptions : The DeleteBuild options.
type DeleteBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build.
	BuildName *string `json:"build_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBuildOptions : Instantiate DeleteBuildOptions
func (*CodeEngineV2) NewDeleteBuildOptions(projectGuid string, buildName string) *DeleteBuildOptions {
	return &DeleteBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildName: core.StringPtr(buildName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteBuildOptions) SetProjectGuid(projectGuid string) *DeleteBuildOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildName : Allow user to set BuildName
func (_options *DeleteBuildOptions) SetBuildName(buildName string) *DeleteBuildOptions {
	_options.BuildName = core.StringPtr(buildName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBuildOptions) SetHeaders(param map[string]string) *DeleteBuildOptions {
	options.Headers = param
	return options
}

// DeleteBuildrunOptions : The DeleteBuildrun options.
type DeleteBuildrunOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build run.
	BuildrunName *string `json:"buildrun_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBuildrunOptions : Instantiate DeleteBuildrunOptions
func (*CodeEngineV2) NewDeleteBuildrunOptions(projectGuid string, buildrunName string) *DeleteBuildrunOptions {
	return &DeleteBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildrunName: core.StringPtr(buildrunName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteBuildrunOptions) SetProjectGuid(projectGuid string) *DeleteBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildrunName : Allow user to set BuildrunName
func (_options *DeleteBuildrunOptions) SetBuildrunName(buildrunName string) *DeleteBuildrunOptions {
	_options.BuildrunName = core.StringPtr(buildrunName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBuildrunOptions) SetHeaders(param map[string]string) *DeleteBuildrunOptions {
	options.Headers = param
	return options
}

// DeleteConfigmapOptions : The DeleteConfigmap options.
type DeleteConfigmapOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your configmap.
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
	// The ID of the project.
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

// DeleteSecretOptions : The DeleteSecret options.
type DeleteSecretOptions struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Secret name.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSecretOptions : Instantiate DeleteSecretOptions
func (*CodeEngineV2) NewDeleteSecretOptions(refreshToken string, projectGuid string, secretName string) *DeleteSecretOptions {
	return &DeleteSecretOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *DeleteSecretOptions) SetRefreshToken(refreshToken string) *DeleteSecretOptions {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteSecretOptions) SetProjectGuid(projectGuid string) *DeleteSecretOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetSecretName : Allow user to set SecretName
func (_options *DeleteSecretOptions) SetSecretName(secretName string) *DeleteSecretOptions {
	_options.SecretName = core.StringPtr(secretName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSecretOptions) SetHeaders(param map[string]string) *DeleteSecretOptions {
	options.Headers = param
	return options
}

// GetBuildOptions : The GetBuild options.
type GetBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build.
	BuildName *string `json:"build_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBuildOptions : Instantiate GetBuildOptions
func (*CodeEngineV2) NewGetBuildOptions(projectGuid string, buildName string) *GetBuildOptions {
	return &GetBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildName: core.StringPtr(buildName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetBuildOptions) SetProjectGuid(projectGuid string) *GetBuildOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildName : Allow user to set BuildName
func (_options *GetBuildOptions) SetBuildName(buildName string) *GetBuildOptions {
	_options.BuildName = core.StringPtr(buildName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBuildOptions) SetHeaders(param map[string]string) *GetBuildOptions {
	options.Headers = param
	return options
}

// GetBuildrunOptions : The GetBuildrun options.
type GetBuildrunOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build run.
	BuildrunName *string `json:"buildrun_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBuildrunOptions : Instantiate GetBuildrunOptions
func (*CodeEngineV2) NewGetBuildrunOptions(projectGuid string, buildrunName string) *GetBuildrunOptions {
	return &GetBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildrunName: core.StringPtr(buildrunName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetBuildrunOptions) SetProjectGuid(projectGuid string) *GetBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildrunName : Allow user to set BuildrunName
func (_options *GetBuildrunOptions) SetBuildrunName(buildrunName string) *GetBuildrunOptions {
	_options.BuildrunName = core.StringPtr(buildrunName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBuildrunOptions) SetHeaders(param map[string]string) *GetBuildrunOptions {
	options.Headers = param
	return options
}

// GetConfigmapOptions : The GetConfigmap options.
type GetConfigmapOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your configmap.
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
	// The ID of the project.
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
	// The ID of the project.
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

// GetSecretOptions : The GetSecret options.
type GetSecretOptions struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Secret name.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSecretOptions : Instantiate GetSecretOptions
func (*CodeEngineV2) NewGetSecretOptions(refreshToken string, projectGuid string, secretName string) *GetSecretOptions {
	return &GetSecretOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *GetSecretOptions) SetRefreshToken(refreshToken string) *GetSecretOptions {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetSecretOptions) SetProjectGuid(projectGuid string) *GetSecretOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetSecretName : Allow user to set SecretName
func (_options *GetSecretOptions) SetSecretName(secretName string) *GetSecretOptions {
	_options.SecretName = core.StringPtr(secretName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSecretOptions) SetHeaders(param map[string]string) *GetSecretOptions {
	options.Headers = param
	return options
}

// ListBuildrunsOptions : The ListBuildruns options.
type ListBuildrunsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of build runs per page.
	Limit *int64 `json:"limit,omitempty"`

	// Token to continue traversing paginated list of build runs.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBuildrunsOptions : Instantiate ListBuildrunsOptions
func (*CodeEngineV2) NewListBuildrunsOptions(projectGuid string) *ListBuildrunsOptions {
	return &ListBuildrunsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListBuildrunsOptions) SetProjectGuid(projectGuid string) *ListBuildrunsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListBuildrunsOptions) SetLimit(limit int64) *ListBuildrunsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListBuildrunsOptions) SetStart(start string) *ListBuildrunsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBuildrunsOptions) SetHeaders(param map[string]string) *ListBuildrunsOptions {
	options.Headers = param
	return options
}

// ListBuildsOptions : The ListBuilds options.
type ListBuildsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of builds per page.
	Limit *int64 `json:"limit,omitempty"`

	// The token to continue traversing paginated list of builds.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBuildsOptions : Instantiate ListBuildsOptions
func (*CodeEngineV2) NewListBuildsOptions(projectGuid string) *ListBuildsOptions {
	return &ListBuildsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListBuildsOptions) SetProjectGuid(projectGuid string) *ListBuildsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListBuildsOptions) SetLimit(limit int64) *ListBuildsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListBuildsOptions) SetStart(start string) *ListBuildsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBuildsOptions) SetHeaders(param map[string]string) *ListBuildsOptions {
	options.Headers = param
	return options
}

// ListConfigmapsOptions : The ListConfigmaps options.
type ListConfigmapsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of configmaps per page.
	Limit *int64 `json:"limit,omitempty"`

	// Token to continue traversing paginated list of configmaps.
	Start *string `json:"start,omitempty"`

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

// SetLimit : Allow user to set Limit
func (_options *ListConfigmapsOptions) SetLimit(limit int64) *ListConfigmapsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListConfigmapsOptions) SetStart(start string) *ListConfigmapsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigmapsOptions) SetHeaders(param map[string]string) *ListConfigmapsOptions {
	options.Headers = param
	return options
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {
	// The maximum number of projects per page.
	Limit *int64 `json:"limit,omitempty"`

	// Token to continue traversing paginated list of projects.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*CodeEngineV2) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *ListProjectsOptions) SetLimit(limit int64) *ListProjectsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListProjectsOptions) SetStart(start string) *ListProjectsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// ListReclamationsOptions : The ListReclamations options.
type ListReclamationsOptions struct {
	// The maximum number of reclamations per page.
	Limit *int64 `json:"limit,omitempty"`

	// Token to continue traversing paginated list of reclamations.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListReclamationsOptions : Instantiate ListReclamationsOptions
func (*CodeEngineV2) NewListReclamationsOptions() *ListReclamationsOptions {
	return &ListReclamationsOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *ListReclamationsOptions) SetLimit(limit int64) *ListReclamationsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListReclamationsOptions) SetStart(start string) *ListReclamationsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListReclamationsOptions) SetHeaders(param map[string]string) *ListReclamationsOptions {
	options.Headers = param
	return options
}

// ListSecretsOptions : The ListSecrets options.
type ListSecretsOptions struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSecretsOptions : Instantiate ListSecretsOptions
func (*CodeEngineV2) NewListSecretsOptions(refreshToken string, projectGuid string) *ListSecretsOptions {
	return &ListSecretsOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *ListSecretsOptions) SetRefreshToken(refreshToken string) *ListSecretsOptions {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListSecretsOptions) SetProjectGuid(projectGuid string) *ListSecretsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSecretsOptions) SetHeaders(param map[string]string) *ListSecretsOptions {
	options.Headers = param
	return options
}

// ReclaimReclamationOptions : The ReclaimReclamation options.
type ReclaimReclamationOptions struct {
	// The ID of the project.
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
	// The ID of the project.
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

// UpdateBuildOptions : The UpdateBuild options.
type UpdateBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build.
	BuildName *string `json:"build_name" validate:"required,ne="`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The path to the Dockerfile that is used for build strategies for building an image.
	Dockerfile *string `json:"dockerfile,omitempty"`

	// The name of the build. Use a name that is unique within the project.
	Name *string `json:"name,omitempty"`

	// The name of the image.
	OutputImage *string `json:"output_image,omitempty"`

	// The secret that is required to access the image registry.
	OutputSecret *string `json:"output_secret,omitempty"`

	// The directory in the repository that contains the buildpacks file or the Dockerfile.
	SourceContextDir *string `json:"source_context_dir,omitempty"`

	// The commit, tag, or branch in the source repository to pull.
	SourceRevision *string `json:"source_revision,omitempty"`

	// The name of the secret that is required to access the repository source.
	SourceSecret *string `json:"source_secret,omitempty"`

	// Specifies the type of source to determine if your build source is in a repository or based on local source code.
	SourceType *string `json:"source_type,omitempty"`

	// The URL of the repository.
	SourceURL *string `json:"source_url,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBuildOptions : Instantiate UpdateBuildOptions
func (*CodeEngineV2) NewUpdateBuildOptions(projectGuid string, buildName string) *UpdateBuildOptions {
	return &UpdateBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildName: core.StringPtr(buildName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateBuildOptions) SetProjectGuid(projectGuid string) *UpdateBuildOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildName : Allow user to set BuildName
func (_options *UpdateBuildOptions) SetBuildName(buildName string) *UpdateBuildOptions {
	_options.BuildName = core.StringPtr(buildName)
	return _options
}

// SetCeOwnerReference : Allow user to set CeOwnerReference
func (_options *UpdateBuildOptions) SetCeOwnerReference(ceOwnerReference string) *UpdateBuildOptions {
	_options.CeOwnerReference = core.StringPtr(ceOwnerReference)
	return _options
}

// SetDockerfile : Allow user to set Dockerfile
func (_options *UpdateBuildOptions) SetDockerfile(dockerfile string) *UpdateBuildOptions {
	_options.Dockerfile = core.StringPtr(dockerfile)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateBuildOptions) SetName(name string) *UpdateBuildOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetOutputImage : Allow user to set OutputImage
func (_options *UpdateBuildOptions) SetOutputImage(outputImage string) *UpdateBuildOptions {
	_options.OutputImage = core.StringPtr(outputImage)
	return _options
}

// SetOutputSecret : Allow user to set OutputSecret
func (_options *UpdateBuildOptions) SetOutputSecret(outputSecret string) *UpdateBuildOptions {
	_options.OutputSecret = core.StringPtr(outputSecret)
	return _options
}

// SetSourceContextDir : Allow user to set SourceContextDir
func (_options *UpdateBuildOptions) SetSourceContextDir(sourceContextDir string) *UpdateBuildOptions {
	_options.SourceContextDir = core.StringPtr(sourceContextDir)
	return _options
}

// SetSourceRevision : Allow user to set SourceRevision
func (_options *UpdateBuildOptions) SetSourceRevision(sourceRevision string) *UpdateBuildOptions {
	_options.SourceRevision = core.StringPtr(sourceRevision)
	return _options
}

// SetSourceSecret : Allow user to set SourceSecret
func (_options *UpdateBuildOptions) SetSourceSecret(sourceSecret string) *UpdateBuildOptions {
	_options.SourceSecret = core.StringPtr(sourceSecret)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *UpdateBuildOptions) SetSourceType(sourceType string) *UpdateBuildOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetSourceURL : Allow user to set SourceURL
func (_options *UpdateBuildOptions) SetSourceURL(sourceURL string) *UpdateBuildOptions {
	_options.SourceURL = core.StringPtr(sourceURL)
	return _options
}

// SetStrategyName : Allow user to set StrategyName
func (_options *UpdateBuildOptions) SetStrategyName(strategyName string) *UpdateBuildOptions {
	_options.StrategyName = core.StringPtr(strategyName)
	return _options
}

// SetStrategySize : Allow user to set StrategySize
func (_options *UpdateBuildOptions) SetStrategySize(strategySize string) *UpdateBuildOptions {
	_options.StrategySize = core.StringPtr(strategySize)
	return _options
}

// SetTimeout : Allow user to set Timeout
func (_options *UpdateBuildOptions) SetTimeout(timeout int64) *UpdateBuildOptions {
	_options.Timeout = core.Int64Ptr(timeout)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBuildOptions) SetHeaders(param map[string]string) *UpdateBuildOptions {
	options.Headers = param
	return options
}

// UpdateConfigmapOptions : The UpdateConfigmap options.
type UpdateConfigmapOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your configmap.
	ConfigmapName *string `json:"configmap_name" validate:"required,ne="`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// Indicates that the key-value pair cannot be edited.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the configmap. Use a name that is unique within the project.
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

// UpdateSecretOptions : The UpdateSecret options.
type UpdateSecretOptions struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

	// Project GUID.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// Secret name.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// Name of the secret.
	BindingSecretRef *string `json:"binding_secret_ref,omitempty"`

	// List of bound CE Components.
	CeComponents []string `json:"ce_components,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Define whether the secret is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the Secret.
	Name *string `json:"name,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associtaed with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential (resource key) associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Role of the service credential (resource key).
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential (resource key).
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Specify the target of the secret (aka how the secret will be used) label format: "target_<target>: <target>".
	Target *string `json:"target,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSecretOptions : Instantiate UpdateSecretOptions
func (*CodeEngineV2) NewUpdateSecretOptions(refreshToken string, projectGuid string, secretName string) *UpdateSecretOptions {
	return &UpdateSecretOptions{
		RefreshToken: core.StringPtr(refreshToken),
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
	}
}

// SetRefreshToken : Allow user to set RefreshToken
func (_options *UpdateSecretOptions) SetRefreshToken(refreshToken string) *UpdateSecretOptions {
	_options.RefreshToken = core.StringPtr(refreshToken)
	return _options
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateSecretOptions) SetProjectGuid(projectGuid string) *UpdateSecretOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetSecretName : Allow user to set SecretName
func (_options *UpdateSecretOptions) SetSecretName(secretName string) *UpdateSecretOptions {
	_options.SecretName = core.StringPtr(secretName)
	return _options
}

// SetBindingSecretRef : Allow user to set BindingSecretRef
func (_options *UpdateSecretOptions) SetBindingSecretRef(bindingSecretRef string) *UpdateSecretOptions {
	_options.BindingSecretRef = core.StringPtr(bindingSecretRef)
	return _options
}

// SetCeComponents : Allow user to set CeComponents
func (_options *UpdateSecretOptions) SetCeComponents(ceComponents []string) *UpdateSecretOptions {
	_options.CeComponents = ceComponents
	return _options
}

// SetCreated : Allow user to set Created
func (_options *UpdateSecretOptions) SetCreated(created string) *UpdateSecretOptions {
	_options.Created = core.StringPtr(created)
	return _options
}

// SetData : Allow user to set Data
func (_options *UpdateSecretOptions) SetData(data map[string]string) *UpdateSecretOptions {
	_options.Data = data
	return _options
}

// SetFormat : Allow user to set Format
func (_options *UpdateSecretOptions) SetFormat(format string) *UpdateSecretOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateSecretOptions) SetID(id string) *UpdateSecretOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetImmutable : Allow user to set Immutable
func (_options *UpdateSecretOptions) SetImmutable(immutable bool) *UpdateSecretOptions {
	_options.Immutable = core.BoolPtr(immutable)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateSecretOptions) SetName(name string) *UpdateSecretOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceID : Allow user to set ResourceID
func (_options *UpdateSecretOptions) SetResourceID(resourceID string) *UpdateSecretOptions {
	_options.ResourceID = core.StringPtr(resourceID)
	return _options
}

// SetResourceType : Allow user to set ResourceType
func (_options *UpdateSecretOptions) SetResourceType(resourceType string) *UpdateSecretOptions {
	_options.ResourceType = core.StringPtr(resourceType)
	return _options
}

// SetResourcekeyID : Allow user to set ResourcekeyID
func (_options *UpdateSecretOptions) SetResourcekeyID(resourcekeyID string) *UpdateSecretOptions {
	_options.ResourcekeyID = core.StringPtr(resourcekeyID)
	return _options
}

// SetRole : Allow user to set Role
func (_options *UpdateSecretOptions) SetRole(role string) *UpdateSecretOptions {
	_options.Role = core.StringPtr(role)
	return _options
}

// SetServiceidCrn : Allow user to set ServiceidCrn
func (_options *UpdateSecretOptions) SetServiceidCrn(serviceidCrn string) *UpdateSecretOptions {
	_options.ServiceidCrn = core.StringPtr(serviceidCrn)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *UpdateSecretOptions) SetTarget(target string) *UpdateSecretOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSecretOptions) SetHeaders(param map[string]string) *UpdateSecretOptions {
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

// V2Build : Build is the response model for build resources.
type V2Build struct {
	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The path to the Dockerfile that is used for build strategies for building an image.
	Dockerfile *string `json:"dockerfile,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The name of the image.
	OutputImage *string `json:"output_image,omitempty"`

	// The secret that is required to access the image registry.
	OutputSecret *string `json:"output_secret,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// The directory in the repository that contains the buildpacks file or the Dockerfile.
	SourceContextDir *string `json:"source_context_dir,omitempty"`

	// The commit, tag, or branch in the source repository to pull.
	SourceRevision *string `json:"source_revision,omitempty"`

	// The name of the secret that is required to access the repository source.
	SourceSecret *string `json:"source_secret,omitempty"`

	// Specifies the type of source to determine if your build source is in a repository or based on local source code.
	SourceType *string `json:"source_type,omitempty"`

	// The URL of the repository.
	SourceURL *string `json:"source_url,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`
}

// UnmarshalV2Build unmarshals an instance of V2Build from the specified map of raw messages.
func UnmarshalV2Build(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2Build)
	err = core.UnmarshalPrimitive(m, "ce_owner_reference", &obj.CeOwnerReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dockerfile", &obj.Dockerfile)
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
	err = core.UnmarshalPrimitive(m, "output_image", &obj.OutputImage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output_secret", &obj.OutputSecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_context_dir", &obj.SourceContextDir)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_revision", &obj.SourceRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_secret", &obj.SourceSecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_type", &obj.SourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_url", &obj.SourceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "strategy_name", &obj.StrategyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "strategy_size", &obj.StrategySize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
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

// V2BuildList : Contains a list of builds and pagination information.
type V2BuildList struct {
	// List of all builds.
	Builds []V2Build `json:"builds,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalV2BuildList unmarshals an instance of V2BuildList from the specified map of raw messages.
func UnmarshalV2BuildList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2BuildList)
	err = core.UnmarshalModel(m, "builds", &obj.Builds, UnmarshalV2Build)
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *V2BuildList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// V2BuildRun : V2BuildRun struct
type V2BuildRun struct {
	// The name of the app revision with which this build run is associated.
	AppRevision *string `json:"app_revision,omitempty"`

	// The name of the build on which this build run is associated.
	Build *string `json:"build,omitempty"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// The path to the Dockerfile that is used for build strategies for building an image.
	Dockerfile *string `json:"dockerfile,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The name of the image.
	OutputImage *string `json:"output_image,omitempty"`

	// The secret that is required to access the image registry.
	OutputSecret *string `json:"output_secret,omitempty"`

	// ServiceAccount refers to the serviceaccount which is used for resource control.
	ServiceAccount *string `json:"service_account,omitempty"`

	// The directory in the repository that contains the buildpacks file or the Dockerfile.
	SourceContextDir *string `json:"source_context_dir,omitempty"`

	// The commit, tag, or branch in the source repository to pull.
	SourceRevision *string `json:"source_revision,omitempty"`

	// The name of the secret that is required to access the repository source.
	SourceSecret *string `json:"source_secret,omitempty"`

	// Specifies the type of source to determine if your build source is in a repository or based on local source code.
	SourceType *string `json:"source_type,omitempty"`

	// The URL of the repository.
	SourceURL *string `json:"source_url,omitempty"`

	// Describes the current status condition of a build run.
	Status *V2BuildRunStatus `json:"status,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`
}

// UnmarshalV2BuildRun unmarshals an instance of V2BuildRun from the specified map of raw messages.
func UnmarshalV2BuildRun(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2BuildRun)
	err = core.UnmarshalPrimitive(m, "app_revision", &obj.AppRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build", &obj.Build)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ce_owner_reference", &obj.CeOwnerReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dockerfile", &obj.Dockerfile)
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
	err = core.UnmarshalPrimitive(m, "output_image", &obj.OutputImage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output_secret", &obj.OutputSecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_account", &obj.ServiceAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_context_dir", &obj.SourceContextDir)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_revision", &obj.SourceRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_secret", &obj.SourceSecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_type", &obj.SourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_url", &obj.SourceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalV2BuildRunStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "strategy_name", &obj.StrategyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "strategy_size", &obj.StrategySize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
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

// V2BuildRunList : Contains a list of build runs and pagination information.
type V2BuildRunList struct {
	// List of all build runs.
	Buildruns []V2BuildRun `json:"buildruns,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalV2BuildRunList unmarshals an instance of V2BuildRunList from the specified map of raw messages.
func UnmarshalV2BuildRunList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2BuildRunList)
	err = core.UnmarshalModel(m, "buildruns", &obj.Buildruns, UnmarshalV2BuildRun)
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *V2BuildRunList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// V2BuildRunStatus : Describes the current status condition of a build run.
type V2BuildRunStatus struct {
	// Describes the time the build run completed.
	CompletionTime *string `json:"completion_time,omitempty"`

	// Describes the name of the task run responsible for executing this build run.
	LastTaskRun *string `json:"last_task_run,omitempty"`

	// Describes the time the build run started.
	StartTime *string `json:"start_time,omitempty"`
}

// UnmarshalV2BuildRunStatus unmarshals an instance of V2BuildRunStatus from the specified map of raw messages.
func UnmarshalV2BuildRunStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2BuildRunStatus)
	err = core.UnmarshalPrimitive(m, "completion_time", &obj.CompletionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_task_run", &obj.LastTaskRun)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V2ConfigMap : Describes the model of a configmap.
type V2ConfigMap struct {
	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Specifies that the key-value pair cannot be edited.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The type of the resource.
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

// V2ConfigMapList : Contains a list of configmaps and pagination information.
type V2ConfigMapList struct {
	// List of all configmaps.
	Configmaps []V2ConfigMap `json:"configmaps,omitempty"`

	// Maximum number of resources per page.
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *V2ConfigMapList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// V2Project : Describes the model of a project.
type V2Project struct {
	// An alphanumeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The date when the project was created.
	Created *string `json:"created,omitempty"`

	// The ID associated with the project.
	Crn *string `json:"crn,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The ID of the project.
	ID *string `json:"id,omitempty"`

	// The name of the project.
	Name *string `json:"name,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// The region for your project deployment.
	Region *string `json:"region,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The type of the project.
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

// V2ProjectList : Contains a list of projects and pagination information.
type V2ProjectList struct {
	// Maximum number of resources per page.
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

// Retrieve the value to be passed to a request to access the next page of results
func (resp *V2ProjectList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// V2Reclamation : Describes the model of a reclamation.
type V2Reclamation struct {
	// An alphanumeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The ID of the reclamation.
	ID *string `json:"id,omitempty"`

	// The ID of the Code Engine project resource instance.
	ProjectID *string `json:"project_id,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The target time that the reclamation retention period end.
	TargetTime *string `json:"target_time,omitempty"`

	// The type of the reclamation.
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

// V2ReclamationList : Contains a list of reclamations and pagination information.
type V2ReclamationList struct {
	// Maximum number of resources per page.
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of all project reclamations.
	Reclamations []V2Reclamation `json:"reclamations,omitempty"`
}

// UnmarshalV2ReclamationList unmarshals an instance of V2ReclamationList from the specified map of raw messages.
func UnmarshalV2ReclamationList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2ReclamationList)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationListNextMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reclamations", &obj.Reclamations, UnmarshalV2Reclamation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *V2ReclamationList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// V2Secret : A secret resource.
type V2Secret struct {
	// Name of the secret.
	BindingSecretRef *string `json:"binding_secret_ref,omitempty"`

	// List of bound CE Components.
	CeComponents []string `json:"ce_components,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// Define whether the secret is immutable.
	Immutable *bool `json:"immutable,omitempty"`

	// The name of the resource. Use a name that is unique within the project.
	Name *string `json:"name,omitempty"`

	ProjectID *string `json:"project_id,omitempty"`

	Region *string `json:"region,omitempty"`

	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associtaed with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential (resource key) associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Role of the service credential (resource key).
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential (resource key).
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Specify the target of the secret (aka how the secret will be used).
	Target *string `json:"target,omitempty"`

	// Defines the resource type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalV2Secret unmarshals an instance of V2Secret from the specified map of raw messages.
func UnmarshalV2Secret(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2Secret)
	err = core.UnmarshalPrimitive(m, "binding_secret_ref", &obj.BindingSecretRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ce_components", &obj.CeComponents)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "format", &obj.Format)
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
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
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
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resourcekey_id", &obj.ResourcekeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "serviceid_crn", &obj.ServiceidCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
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

// V2SecretList : List of secret resources.
type V2SecretList struct {
	// Maximum number of resources per page.
	Limit *int64 `json:"limit,omitempty"`

	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of Secrets.
	Secrets []V2Secret `json:"secrets,omitempty"`
}

// UnmarshalV2SecretList unmarshals an instance of V2SecretList from the specified map of raw messages.
func UnmarshalV2SecretList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(V2SecretList)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationListNextMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "secrets", &obj.Secrets, UnmarshalV2Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// ProjectsPager can be used to simplify the use of the "ListProjects" method.
//
type ProjectsPager struct {
	hasNext bool
	options *ListProjectsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewProjectsPager returns a new ProjectsPager instance.
func (codeEngine *CodeEngineV2) NewProjectsPager(options *ListProjectsOptions) (pager *ProjectsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListProjectsOptions = *options
	pager = &ProjectsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ProjectsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ProjectsPager) GetNextWithContext(ctx context.Context) (page []V2Project, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListProjectsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Projects

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ProjectsPager) GetAllWithContext(ctx context.Context) (allItems []V2Project, err error) {
	for pager.HasNext() {
		var nextPage []V2Project
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetNext() (page []V2Project, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetAll() (allItems []V2Project, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// BuildsPager can be used to simplify the use of the "ListBuilds" method.
//
type BuildsPager struct {
	hasNext bool
	options *ListBuildsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewBuildsPager returns a new BuildsPager instance.
func (codeEngine *CodeEngineV2) NewBuildsPager(options *ListBuildsOptions) (pager *BuildsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListBuildsOptions = *options
	pager = &BuildsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *BuildsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *BuildsPager) GetNextWithContext(ctx context.Context) (page []V2Build, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListBuildsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Builds

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *BuildsPager) GetAllWithContext(ctx context.Context) (allItems []V2Build, err error) {
	for pager.HasNext() {
		var nextPage []V2Build
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BuildsPager) GetNext() (page []V2Build, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BuildsPager) GetAll() (allItems []V2Build, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// BuildrunsPager can be used to simplify the use of the "ListBuildruns" method.
//
type BuildrunsPager struct {
	hasNext bool
	options *ListBuildrunsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewBuildrunsPager returns a new BuildrunsPager instance.
func (codeEngine *CodeEngineV2) NewBuildrunsPager(options *ListBuildrunsOptions) (pager *BuildrunsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListBuildrunsOptions = *options
	pager = &BuildrunsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *BuildrunsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *BuildrunsPager) GetNextWithContext(ctx context.Context) (page []V2BuildRun, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListBuildrunsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Buildruns

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *BuildrunsPager) GetAllWithContext(ctx context.Context) (allItems []V2BuildRun, err error) {
	for pager.HasNext() {
		var nextPage []V2BuildRun
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BuildrunsPager) GetNext() (page []V2BuildRun, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BuildrunsPager) GetAll() (allItems []V2BuildRun, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// ConfigmapsPager can be used to simplify the use of the "ListConfigmaps" method.
//
type ConfigmapsPager struct {
	hasNext bool
	options *ListConfigmapsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewConfigmapsPager returns a new ConfigmapsPager instance.
func (codeEngine *CodeEngineV2) NewConfigmapsPager(options *ListConfigmapsOptions) (pager *ConfigmapsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListConfigmapsOptions = *options
	pager = &ConfigmapsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ConfigmapsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ConfigmapsPager) GetNextWithContext(ctx context.Context) (page []V2ConfigMap, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListConfigmapsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Configmaps

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ConfigmapsPager) GetAllWithContext(ctx context.Context) (allItems []V2ConfigMap, err error) {
	for pager.HasNext() {
		var nextPage []V2ConfigMap
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ConfigmapsPager) GetNext() (page []V2ConfigMap, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ConfigmapsPager) GetAll() (allItems []V2ConfigMap, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// ReclamationsPager can be used to simplify the use of the "ListReclamations" method.
//
type ReclamationsPager struct {
	hasNext bool
	options *ListReclamationsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewReclamationsPager returns a new ReclamationsPager instance.
func (codeEngine *CodeEngineV2) NewReclamationsPager(options *ListReclamationsOptions) (pager *ReclamationsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListReclamationsOptions = *options
	pager = &ReclamationsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ReclamationsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ReclamationsPager) GetNextWithContext(ctx context.Context) (page []V2Reclamation, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListReclamationsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Reclamations

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ReclamationsPager) GetAllWithContext(ctx context.Context) (allItems []V2Reclamation, err error) {
	for pager.HasNext() {
		var nextPage []V2Reclamation
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ReclamationsPager) GetNext() (page []V2Reclamation, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ReclamationsPager) GetAll() (allItems []V2Reclamation, err error) {
	return pager.GetAllWithContext(context.Background())
}
