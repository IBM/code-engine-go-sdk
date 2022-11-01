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
// List all projects in the current account.
func (codeEngine *CodeEngineV2) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ProjectList, response *core.DetailedResponse, err error) {
	return codeEngine.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ProjectList, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateProject : Create a project
// Create a project in the current resource group.
func (codeEngine *CodeEngineV2) CreateProject(createProjectOptions *CreateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
	return codeEngine.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProject)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProject : Get a project
// Display the details of a single project.
func (codeEngine *CodeEngineV2) GetProject(getProjectOptions *GetProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
	return codeEngine.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProject)
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
func (codeEngine *CodeEngineV2) ListBuilds(listBuildsOptions *ListBuildsOptions) (result *BuildList, response *core.DetailedResponse, err error) {
	return codeEngine.ListBuildsWithContext(context.Background(), listBuildsOptions)
}

// ListBuildsWithContext is an alternate form of the ListBuilds method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListBuildsWithContext(ctx context.Context, listBuildsOptions *ListBuildsOptions) (result *BuildList, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuildList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBuild : Create a build
// Create a build.
func (codeEngine *CodeEngineV2) CreateBuild(createBuildOptions *CreateBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
	return codeEngine.CreateBuildWithContext(context.Background(), createBuildOptions)
}

// CreateBuildWithContext is an alternate form of the CreateBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateBuildWithContext(ctx context.Context, createBuildOptions *CreateBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
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
	if createBuildOptions.Name != nil {
		body["name"] = createBuildOptions.Name
	}
	if createBuildOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = createBuildOptions.CeOwnerReference
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
	if createBuildOptions.StrategySpecFile != nil {
		body["strategy_spec_file"] = createBuildOptions.StrategySpecFile
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuild)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBuild : Get a build
// Display the details of a build.
func (codeEngine *CodeEngineV2) GetBuild(getBuildOptions *GetBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
	return codeEngine.GetBuildWithContext(context.Background(), getBuildOptions)
}

// GetBuildWithContext is an alternate form of the GetBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetBuildWithContext(ctx context.Context, getBuildOptions *GetBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuild)
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
func (codeEngine *CodeEngineV2) UpdateBuild(updateBuildOptions *UpdateBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateBuildWithContext(context.Background(), updateBuildOptions)
}

// UpdateBuildWithContext is an alternate form of the UpdateBuild method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateBuildWithContext(ctx context.Context, updateBuildOptions *UpdateBuildOptions) (result *Build, response *core.DetailedResponse, err error) {
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
	if updateBuildOptions.Name != nil {
		body["name"] = updateBuildOptions.Name
	}
	if updateBuildOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = updateBuildOptions.CeOwnerReference
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
	if updateBuildOptions.StrategySpecFile != nil {
		body["strategy_spec_file"] = updateBuildOptions.StrategySpecFile
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuild)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBuildruns : List build runs
// List all build runs in a project.
func (codeEngine *CodeEngineV2) ListBuildruns(listBuildrunsOptions *ListBuildrunsOptions) (result *BuildRunList, response *core.DetailedResponse, err error) {
	return codeEngine.ListBuildrunsWithContext(context.Background(), listBuildrunsOptions)
}

// ListBuildrunsWithContext is an alternate form of the ListBuildruns method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListBuildrunsWithContext(ctx context.Context, listBuildrunsOptions *ListBuildrunsOptions) (result *BuildRunList, response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/build_runs`, pathParamsMap)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuildRunList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBuildrun : Create a build run
// Create a build run.
func (codeEngine *CodeEngineV2) CreateBuildrun(createBuildrunOptions *CreateBuildrunOptions) (result *BuildRun, response *core.DetailedResponse, err error) {
	return codeEngine.CreateBuildrunWithContext(context.Background(), createBuildrunOptions)
}

// CreateBuildrunWithContext is an alternate form of the CreateBuildrun method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateBuildrunWithContext(ctx context.Context, createBuildrunOptions *CreateBuildrunOptions) (result *BuildRun, response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/build_runs`, pathParamsMap)
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
	if createBuildrunOptions.Name != nil {
		body["name"] = createBuildrunOptions.Name
	}
	if createBuildrunOptions.AppRevision != nil {
		body["app_revision"] = createBuildrunOptions.AppRevision
	}
	if createBuildrunOptions.Build != nil {
		body["build"] = createBuildrunOptions.Build
	}
	if createBuildrunOptions.CeOwnerReference != nil {
		body["ce_owner_reference"] = createBuildrunOptions.CeOwnerReference
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
	if createBuildrunOptions.StrategySpecFile != nil {
		body["strategy_spec_file"] = createBuildrunOptions.StrategySpecFile
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuildRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBuildrun : Get a build run
// Display the details of a build run.
func (codeEngine *CodeEngineV2) GetBuildrun(getBuildrunOptions *GetBuildrunOptions) (result *BuildRun, response *core.DetailedResponse, err error) {
	return codeEngine.GetBuildrunWithContext(context.Background(), getBuildrunOptions)
}

// GetBuildrunWithContext is an alternate form of the GetBuildrun method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetBuildrunWithContext(ctx context.Context, getBuildrunOptions *GetBuildrunOptions) (result *BuildRun, response *core.DetailedResponse, err error) {
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
		"build_run_name": *getBuildrunOptions.BuildRunName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/build_runs/{build_run_name}`, pathParamsMap)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBuildRun)
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
		"build_run_name": *deleteBuildrunOptions.BuildRunName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/build_runs/{build_run_name}`, pathParamsMap)
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
// List all config maps in a project.
func (codeEngine *CodeEngineV2) ListConfigmaps(listConfigmapsOptions *ListConfigmapsOptions) (result *ConfigMapList, response *core.DetailedResponse, err error) {
	return codeEngine.ListConfigmapsWithContext(context.Background(), listConfigmapsOptions)
}

// ListConfigmapsWithContext is an alternate form of the ListConfigmaps method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListConfigmapsWithContext(ctx context.Context, listConfigmapsOptions *ListConfigmapsOptions) (result *ConfigMapList, response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/config_maps`, pathParamsMap)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigMapList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateConfigmap : Create a configmap
// Create a configmap.
func (codeEngine *CodeEngineV2) CreateConfigmap(createConfigmapOptions *CreateConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.CreateConfigmapWithContext(context.Background(), createConfigmapOptions)
}

// CreateConfigmapWithContext is an alternate form of the CreateConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateConfigmapWithContext(ctx context.Context, createConfigmapOptions *CreateConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/config_maps`, pathParamsMap)
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
	if createConfigmapOptions.Name != nil {
		body["name"] = createConfigmapOptions.Name
	}
	if createConfigmapOptions.Data != nil {
		body["data"] = createConfigmapOptions.Data
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigMap)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfigmap : Get a configmap
// Display the details of a configmap.
func (codeEngine *CodeEngineV2) GetConfigmap(getConfigmapOptions *GetConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.GetConfigmapWithContext(context.Background(), getConfigmapOptions)
}

// GetConfigmapWithContext is an alternate form of the GetConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetConfigmapWithContext(ctx context.Context, getConfigmapOptions *GetConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
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
		"config_map_name": *getConfigmapOptions.ConfigMapName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/config_maps/{config_map_name}`, pathParamsMap)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigMap)
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
		"config_map_name": *deleteConfigmapOptions.ConfigMapName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/config_maps/{config_map_name}`, pathParamsMap)
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
func (codeEngine *CodeEngineV2) UpdateConfigmap(updateConfigmapOptions *UpdateConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateConfigmapWithContext(context.Background(), updateConfigmapOptions)
}

// UpdateConfigmapWithContext is an alternate form of the UpdateConfigmap method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateConfigmapWithContext(ctx context.Context, updateConfigmapOptions *UpdateConfigmapOptions) (result *ConfigMap, response *core.DetailedResponse, err error) {
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
		"config_map_name": *updateConfigmapOptions.ConfigMapName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/config_maps/{config_map_name}`, pathParamsMap)
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
	if updateConfigmapOptions.Name != nil {
		body["name"] = updateConfigmapOptions.Name
	}
	if updateConfigmapOptions.Data != nil {
		body["data"] = updateConfigmapOptions.Data
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigMap)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListSecrets : List secrets
// List all secrets in a project.
func (codeEngine *CodeEngineV2) ListSecrets(listSecretsOptions *ListSecretsOptions) (result *SecretList, response *core.DetailedResponse, err error) {
	return codeEngine.ListSecretsWithContext(context.Background(), listSecretsOptions)
}

// ListSecretsWithContext is an alternate form of the ListSecrets method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListSecretsWithContext(ctx context.Context, listSecretsOptions *ListSecretsOptions) (result *SecretList, response *core.DetailedResponse, err error) {
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

	if listSecretsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listSecretsOptions.Limit))
	}
	if listSecretsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listSecretsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecretList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSecret : Create a secret
// Create a secret.
func (codeEngine *CodeEngineV2) CreateSecret(createSecretOptions *CreateSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
	return codeEngine.CreateSecretWithContext(context.Background(), createSecretOptions)
}

// CreateSecretWithContext is an alternate form of the CreateSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateSecretWithContext(ctx context.Context, createSecretOptions *CreateSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
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

	body := make(map[string]interface{})
	if createSecretOptions.Name != nil {
		body["name"] = createSecretOptions.Name
	}
	if createSecretOptions.CeComponents != nil {
		body["ce_components"] = createSecretOptions.CeComponents
	}
	if createSecretOptions.Data != nil {
		body["data"] = createSecretOptions.Data
	}
	if createSecretOptions.Format != nil {
		body["format"] = createSecretOptions.Format
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
	if createSecretOptions.ResourcekeyName != nil {
		body["resourcekey_name"] = createSecretOptions.ResourcekeyName
	}
	if createSecretOptions.Role != nil {
		body["role"] = createSecretOptions.Role
	}
	if createSecretOptions.ServiceidCrn != nil {
		body["serviceid_crn"] = createSecretOptions.ServiceidCrn
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecret)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSecret : Get a secret
// Get a secret.
func (codeEngine *CodeEngineV2) GetSecret(getSecretOptions *GetSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
	return codeEngine.GetSecretWithContext(context.Background(), getSecretOptions)
}

// GetSecretWithContext is an alternate form of the GetSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetSecretWithContext(ctx context.Context, getSecretOptions *GetSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecret)
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

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = codeEngine.Service.Request(request, nil)

	return
}

// UpdateSecret : Update a secret
// Update a secret.
func (codeEngine *CodeEngineV2) UpdateSecret(updateSecretOptions *UpdateSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateSecretWithContext(context.Background(), updateSecretOptions)
}

// UpdateSecretWithContext is an alternate form of the UpdateSecret method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateSecretWithContext(ctx context.Context, updateSecretOptions *UpdateSecretOptions) (result *Secret, response *core.DetailedResponse, err error) {
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

	body := make(map[string]interface{})
	if updateSecretOptions.Name != nil {
		body["name"] = updateSecretOptions.Name
	}
	if updateSecretOptions.CeComponents != nil {
		body["ce_components"] = updateSecretOptions.CeComponents
	}
	if updateSecretOptions.Data != nil {
		body["data"] = updateSecretOptions.Data
	}
	if updateSecretOptions.Format != nil {
		body["format"] = updateSecretOptions.Format
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
	if updateSecretOptions.ResourcekeyName != nil {
		body["resourcekey_name"] = updateSecretOptions.ResourcekeyName
	}
	if updateSecretOptions.Role != nil {
		body["role"] = updateSecretOptions.Role
	}
	if updateSecretOptions.ServiceidCrn != nil {
		body["serviceid_crn"] = updateSecretOptions.ServiceidCrn
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecret)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListApps : List applications
// List all applications in a project.
func (codeEngine *CodeEngineV2) ListApps(listAppsOptions *ListAppsOptions) (result *AppList, response *core.DetailedResponse, err error) {
	return codeEngine.ListAppsWithContext(context.Background(), listAppsOptions)
}

// ListAppsWithContext is an alternate form of the ListApps method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListAppsWithContext(ctx context.Context, listAppsOptions *ListAppsOptions) (result *AppList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAppsOptions, "listAppsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAppsOptions, "listAppsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listAppsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAppsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListApps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAppsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAppsOptions.Limit))
	}
	if listAppsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listAppsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAppList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateApp : Create an application
// Create an application.
func (codeEngine *CodeEngineV2) CreateApp(createAppOptions *CreateAppOptions) (result *App, response *core.DetailedResponse, err error) {
	return codeEngine.CreateAppWithContext(context.Background(), createAppOptions)
}

// CreateAppWithContext is an alternate form of the CreateApp method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateAppWithContext(ctx context.Context, createAppOptions *CreateAppOptions) (result *App, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAppOptions, "createAppOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAppOptions, "createAppOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createAppOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAppOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateApp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAppOptions.Name != nil {
		body["name"] = createAppOptions.Name
	}
	if createAppOptions.CeManagedDomainMappings != nil {
		body["ce_managed_domain_mappings"] = createAppOptions.CeManagedDomainMappings
	}
	if createAppOptions.ImagePort != nil {
		body["image_port"] = createAppOptions.ImagePort
	}
	if createAppOptions.ImageProtocol != nil {
		body["image_protocol"] = createAppOptions.ImageProtocol
	}
	if createAppOptions.ImageRef != nil {
		body["image_ref"] = createAppOptions.ImageRef
	}
	if createAppOptions.ImageSecret != nil {
		body["image_secret"] = createAppOptions.ImageSecret
	}
	if createAppOptions.RevisionSuffix != nil {
		body["revision_suffix"] = createAppOptions.RevisionSuffix
	}
	if createAppOptions.RunArgs != nil {
		body["run_args"] = createAppOptions.RunArgs
	}
	if createAppOptions.RunAsUser != nil {
		body["run_as_user"] = createAppOptions.RunAsUser
	}
	if createAppOptions.RunCommands != nil {
		body["run_commands"] = createAppOptions.RunCommands
	}
	if createAppOptions.RunEnvVars != nil {
		body["run_env_vars"] = createAppOptions.RunEnvVars
	}
	if createAppOptions.RunServiceAccount != nil {
		body["run_service_account"] = createAppOptions.RunServiceAccount
	}
	if createAppOptions.RunVolumeMounts != nil {
		body["run_volume_mounts"] = createAppOptions.RunVolumeMounts
	}
	if createAppOptions.ScaleConcurrency != nil {
		body["scale_concurrency"] = createAppOptions.ScaleConcurrency
	}
	if createAppOptions.ScaleConcurrencyTarget != nil {
		body["scale_concurrency_target"] = createAppOptions.ScaleConcurrencyTarget
	}
	if createAppOptions.ScaleCpuLimit != nil {
		body["scale_cpu_limit"] = createAppOptions.ScaleCpuLimit
	}
	if createAppOptions.ScaleEphemeralStorageLimit != nil {
		body["scale_ephemeral_storage_limit"] = createAppOptions.ScaleEphemeralStorageLimit
	}
	if createAppOptions.ScaleInitialInstances != nil {
		body["scale_initial_instances"] = createAppOptions.ScaleInitialInstances
	}
	if createAppOptions.ScaleMaxInstances != nil {
		body["scale_max_instances"] = createAppOptions.ScaleMaxInstances
	}
	if createAppOptions.ScaleMemoryLimit != nil {
		body["scale_memory_limit"] = createAppOptions.ScaleMemoryLimit
	}
	if createAppOptions.ScaleMinInstances != nil {
		body["scale_min_instances"] = createAppOptions.ScaleMinInstances
	}
	if createAppOptions.ScaleRequestTimeout != nil {
		body["scale_request_timeout"] = createAppOptions.ScaleRequestTimeout
	}
	if createAppOptions.Version != nil {
		body["version"] = createAppOptions.Version
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetApp : Get an application
// Display the details of an application.
func (codeEngine *CodeEngineV2) GetApp(getAppOptions *GetAppOptions) (result *App, response *core.DetailedResponse, err error) {
	return codeEngine.GetAppWithContext(context.Background(), getAppOptions)
}

// GetAppWithContext is an alternate form of the GetApp method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetAppWithContext(ctx context.Context, getAppOptions *GetAppOptions) (result *App, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAppOptions, "getAppOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAppOptions, "getAppOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getAppOptions.ProjectGuid,
		"app_name": *getAppOptions.AppName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAppOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetApp")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteApp : Delete an application
// Delete an application.
func (codeEngine *CodeEngineV2) DeleteApp(deleteAppOptions *DeleteAppOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteAppWithContext(context.Background(), deleteAppOptions)
}

// DeleteAppWithContext is an alternate form of the DeleteApp method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteAppWithContext(ctx context.Context, deleteAppOptions *DeleteAppOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAppOptions, "deleteAppOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAppOptions, "deleteAppOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteAppOptions.ProjectGuid,
		"app_name": *deleteAppOptions.AppName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAppOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteApp")
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

// UpdateApp : Update an application
// Update the given application.
func (codeEngine *CodeEngineV2) UpdateApp(updateAppOptions *UpdateAppOptions) (result *App, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateAppWithContext(context.Background(), updateAppOptions)
}

// UpdateAppWithContext is an alternate form of the UpdateApp method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateAppWithContext(ctx context.Context, updateAppOptions *UpdateAppOptions) (result *App, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAppOptions, "updateAppOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAppOptions, "updateAppOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateAppOptions.ProjectGuid,
		"app_name": *updateAppOptions.AppName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAppOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateApp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAppOptions.Name != nil {
		body["name"] = updateAppOptions.Name
	}
	if updateAppOptions.CeManagedDomainMappings != nil {
		body["ce_managed_domain_mappings"] = updateAppOptions.CeManagedDomainMappings
	}
	if updateAppOptions.ImagePort != nil {
		body["image_port"] = updateAppOptions.ImagePort
	}
	if updateAppOptions.ImageProtocol != nil {
		body["image_protocol"] = updateAppOptions.ImageProtocol
	}
	if updateAppOptions.ImageRef != nil {
		body["image_ref"] = updateAppOptions.ImageRef
	}
	if updateAppOptions.ImageSecret != nil {
		body["image_secret"] = updateAppOptions.ImageSecret
	}
	if updateAppOptions.RevisionSuffix != nil {
		body["revision_suffix"] = updateAppOptions.RevisionSuffix
	}
	if updateAppOptions.RunArgs != nil {
		body["run_args"] = updateAppOptions.RunArgs
	}
	if updateAppOptions.RunAsUser != nil {
		body["run_as_user"] = updateAppOptions.RunAsUser
	}
	if updateAppOptions.RunCommands != nil {
		body["run_commands"] = updateAppOptions.RunCommands
	}
	if updateAppOptions.RunEnvVars != nil {
		body["run_env_vars"] = updateAppOptions.RunEnvVars
	}
	if updateAppOptions.RunServiceAccount != nil {
		body["run_service_account"] = updateAppOptions.RunServiceAccount
	}
	if updateAppOptions.RunVolumeMounts != nil {
		body["run_volume_mounts"] = updateAppOptions.RunVolumeMounts
	}
	if updateAppOptions.ScaleConcurrency != nil {
		body["scale_concurrency"] = updateAppOptions.ScaleConcurrency
	}
	if updateAppOptions.ScaleConcurrencyTarget != nil {
		body["scale_concurrency_target"] = updateAppOptions.ScaleConcurrencyTarget
	}
	if updateAppOptions.ScaleCpuLimit != nil {
		body["scale_cpu_limit"] = updateAppOptions.ScaleCpuLimit
	}
	if updateAppOptions.ScaleEphemeralStorageLimit != nil {
		body["scale_ephemeral_storage_limit"] = updateAppOptions.ScaleEphemeralStorageLimit
	}
	if updateAppOptions.ScaleInitialInstances != nil {
		body["scale_initial_instances"] = updateAppOptions.ScaleInitialInstances
	}
	if updateAppOptions.ScaleMaxInstances != nil {
		body["scale_max_instances"] = updateAppOptions.ScaleMaxInstances
	}
	if updateAppOptions.ScaleMemoryLimit != nil {
		body["scale_memory_limit"] = updateAppOptions.ScaleMemoryLimit
	}
	if updateAppOptions.ScaleMinInstances != nil {
		body["scale_min_instances"] = updateAppOptions.ScaleMinInstances
	}
	if updateAppOptions.ScaleRequestTimeout != nil {
		body["scale_request_timeout"] = updateAppOptions.ScaleRequestTimeout
	}
	if updateAppOptions.Version != nil {
		body["version"] = updateAppOptions.Version
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAppRevisions : List application revisions
// List all application revisions in a particular application.
func (codeEngine *CodeEngineV2) ListAppRevisions(listAppRevisionsOptions *ListAppRevisionsOptions) (result *AppRevisionList, response *core.DetailedResponse, err error) {
	return codeEngine.ListAppRevisionsWithContext(context.Background(), listAppRevisionsOptions)
}

// ListAppRevisionsWithContext is an alternate form of the ListAppRevisions method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListAppRevisionsWithContext(ctx context.Context, listAppRevisionsOptions *ListAppRevisionsOptions) (result *AppRevisionList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAppRevisionsOptions, "listAppRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAppRevisionsOptions, "listAppRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listAppRevisionsOptions.ProjectGuid,
		"app_name": *listAppRevisionsOptions.AppName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAppRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListAppRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAppRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAppRevisionsOptions.Limit))
	}
	if listAppRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listAppRevisionsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAppRevisionList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAppRevision : Get an application revision
// Display the details of an application revision.
func (codeEngine *CodeEngineV2) GetAppRevision(getAppRevisionOptions *GetAppRevisionOptions) (result *AppRevision, response *core.DetailedResponse, err error) {
	return codeEngine.GetAppRevisionWithContext(context.Background(), getAppRevisionOptions)
}

// GetAppRevisionWithContext is an alternate form of the GetAppRevision method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetAppRevisionWithContext(ctx context.Context, getAppRevisionOptions *GetAppRevisionOptions) (result *AppRevision, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAppRevisionOptions, "getAppRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAppRevisionOptions, "getAppRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getAppRevisionOptions.ProjectGuid,
		"app_name": *getAppRevisionOptions.AppName,
		"revision_name": *getAppRevisionOptions.RevisionName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}/revisions/{revision_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAppRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetAppRevision")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAppRevision)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAppRevision : Delete an application revision
// Delete an application revision.
func (codeEngine *CodeEngineV2) DeleteAppRevision(deleteAppRevisionOptions *DeleteAppRevisionOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteAppRevisionWithContext(context.Background(), deleteAppRevisionOptions)
}

// DeleteAppRevisionWithContext is an alternate form of the DeleteAppRevision method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteAppRevisionWithContext(ctx context.Context, deleteAppRevisionOptions *DeleteAppRevisionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAppRevisionOptions, "deleteAppRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAppRevisionOptions, "deleteAppRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteAppRevisionOptions.ProjectGuid,
		"app_name": *deleteAppRevisionOptions.AppName,
		"revision_name": *deleteAppRevisionOptions.RevisionName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/apps/{app_name}/revisions/{revision_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAppRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteAppRevision")
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

// ListJobs : List jobs
// List all jobs in a project.
func (codeEngine *CodeEngineV2) ListJobs(listJobsOptions *ListJobsOptions) (result *JobList, response *core.DetailedResponse, err error) {
	return codeEngine.ListJobsWithContext(context.Background(), listJobsOptions)
}

// ListJobsWithContext is an alternate form of the ListJobs method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListJobsWithContext(ctx context.Context, listJobsOptions *ListJobsOptions) (result *JobList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listJobsOptions, "listJobsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listJobsOptions, "listJobsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *listJobsOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/jobs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "ListJobs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listJobsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listJobsOptions.Limit))
	}
	if listJobsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listJobsOptions.Start))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateJob : Create an job
// Create an job.
func (codeEngine *CodeEngineV2) CreateJob(createJobOptions *CreateJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	return codeEngine.CreateJobWithContext(context.Background(), createJobOptions)
}

// CreateJobWithContext is an alternate form of the CreateJob method which supports a Context parameter
func (codeEngine *CodeEngineV2) CreateJobWithContext(ctx context.Context, createJobOptions *CreateJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createJobOptions, "createJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createJobOptions, "createJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *createJobOptions.ProjectGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/jobs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "CreateJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createJobOptions.ImageRef != nil {
		body["image_ref"] = createJobOptions.ImageRef
	}
	if createJobOptions.ImageSecret != nil {
		body["image_secret"] = createJobOptions.ImageSecret
	}
	if createJobOptions.Name != nil {
		body["name"] = createJobOptions.Name
	}
	if createJobOptions.RunArgs != nil {
		body["run_args"] = createJobOptions.RunArgs
	}
	if createJobOptions.RunAsUser != nil {
		body["run_as_user"] = createJobOptions.RunAsUser
	}
	if createJobOptions.RunCommands != nil {
		body["run_commands"] = createJobOptions.RunCommands
	}
	if createJobOptions.RunEnvVars != nil {
		body["run_env_vars"] = createJobOptions.RunEnvVars
	}
	if createJobOptions.RunMode != nil {
		body["run_mode"] = createJobOptions.RunMode
	}
	if createJobOptions.RunServiceAccount != nil {
		body["run_service_account"] = createJobOptions.RunServiceAccount
	}
	if createJobOptions.RunVolumeMounts != nil {
		body["run_volume_mounts"] = createJobOptions.RunVolumeMounts
	}
	if createJobOptions.ScaleArraySpec != nil {
		body["scale_array_spec"] = createJobOptions.ScaleArraySpec
	}
	if createJobOptions.ScaleCpuLimit != nil {
		body["scale_cpu_limit"] = createJobOptions.ScaleCpuLimit
	}
	if createJobOptions.ScaleEphemeralStorageLimit != nil {
		body["scale_ephemeral_storage_limit"] = createJobOptions.ScaleEphemeralStorageLimit
	}
	if createJobOptions.ScaleMaxExecutionTime != nil {
		body["scale_max_execution_time"] = createJobOptions.ScaleMaxExecutionTime
	}
	if createJobOptions.ScaleMemoryLimit != nil {
		body["scale_memory_limit"] = createJobOptions.ScaleMemoryLimit
	}
	if createJobOptions.ScaleRetryLimit != nil {
		body["scale_retry_limit"] = createJobOptions.ScaleRetryLimit
	}
	if createJobOptions.Version != nil {
		body["version"] = createJobOptions.Version
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJob)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetJob : Get an job
// Display the details of an job.
func (codeEngine *CodeEngineV2) GetJob(getJobOptions *GetJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	return codeEngine.GetJobWithContext(context.Background(), getJobOptions)
}

// GetJobWithContext is an alternate form of the GetJob method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetJobWithContext(ctx context.Context, getJobOptions *GetJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getJobOptions, "getJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getJobOptions, "getJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *getJobOptions.ProjectGuid,
		"job_name": *getJobOptions.JobName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/jobs/{job_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "GetJob")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJob)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteJob : Delete an job
// Delete an job.
func (codeEngine *CodeEngineV2) DeleteJob(deleteJobOptions *DeleteJobOptions) (response *core.DetailedResponse, err error) {
	return codeEngine.DeleteJobWithContext(context.Background(), deleteJobOptions)
}

// DeleteJobWithContext is an alternate form of the DeleteJob method which supports a Context parameter
func (codeEngine *CodeEngineV2) DeleteJobWithContext(ctx context.Context, deleteJobOptions *DeleteJobOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteJobOptions, "deleteJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteJobOptions, "deleteJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *deleteJobOptions.ProjectGuid,
		"job_name": *deleteJobOptions.JobName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/jobs/{job_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "DeleteJob")
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

// UpdateJob : Update an job
// Update the given job.
func (codeEngine *CodeEngineV2) UpdateJob(updateJobOptions *UpdateJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	return codeEngine.UpdateJobWithContext(context.Background(), updateJobOptions)
}

// UpdateJobWithContext is an alternate form of the UpdateJob method which supports a Context parameter
func (codeEngine *CodeEngineV2) UpdateJobWithContext(ctx context.Context, updateJobOptions *UpdateJobOptions) (result *Job, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateJobOptions, "updateJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateJobOptions, "updateJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_guid": *updateJobOptions.ProjectGuid,
		"job_name": *updateJobOptions.JobName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = codeEngine.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(codeEngine.Service.Options.URL, `/projects/{project_guid}/jobs/{job_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("code_engine", "V2", "UpdateJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateJobOptions.ImageRef != nil {
		body["image_ref"] = updateJobOptions.ImageRef
	}
	if updateJobOptions.ImageSecret != nil {
		body["image_secret"] = updateJobOptions.ImageSecret
	}
	if updateJobOptions.Name != nil {
		body["name"] = updateJobOptions.Name
	}
	if updateJobOptions.RunArgs != nil {
		body["run_args"] = updateJobOptions.RunArgs
	}
	if updateJobOptions.RunAsUser != nil {
		body["run_as_user"] = updateJobOptions.RunAsUser
	}
	if updateJobOptions.RunCommands != nil {
		body["run_commands"] = updateJobOptions.RunCommands
	}
	if updateJobOptions.RunEnvVars != nil {
		body["run_env_vars"] = updateJobOptions.RunEnvVars
	}
	if updateJobOptions.RunMode != nil {
		body["run_mode"] = updateJobOptions.RunMode
	}
	if updateJobOptions.RunServiceAccount != nil {
		body["run_service_account"] = updateJobOptions.RunServiceAccount
	}
	if updateJobOptions.RunVolumeMounts != nil {
		body["run_volume_mounts"] = updateJobOptions.RunVolumeMounts
	}
	if updateJobOptions.ScaleArraySpec != nil {
		body["scale_array_spec"] = updateJobOptions.ScaleArraySpec
	}
	if updateJobOptions.ScaleCpuLimit != nil {
		body["scale_cpu_limit"] = updateJobOptions.ScaleCpuLimit
	}
	if updateJobOptions.ScaleEphemeralStorageLimit != nil {
		body["scale_ephemeral_storage_limit"] = updateJobOptions.ScaleEphemeralStorageLimit
	}
	if updateJobOptions.ScaleMaxExecutionTime != nil {
		body["scale_max_execution_time"] = updateJobOptions.ScaleMaxExecutionTime
	}
	if updateJobOptions.ScaleMemoryLimit != nil {
		body["scale_memory_limit"] = updateJobOptions.ScaleMemoryLimit
	}
	if updateJobOptions.ScaleRetryLimit != nil {
		body["scale_retry_limit"] = updateJobOptions.ScaleRetryLimit
	}
	if updateJobOptions.Version != nil {
		body["version"] = updateJobOptions.Version
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJob)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListReclamations : List all reclamations
// List all project reclamations.
func (codeEngine *CodeEngineV2) ListReclamations(listReclamationsOptions *ListReclamationsOptions) (result *ReclamationList, response *core.DetailedResponse, err error) {
	return codeEngine.ListReclamationsWithContext(context.Background(), listReclamationsOptions)
}

// ListReclamationsWithContext is an alternate form of the ListReclamations method which supports a Context parameter
func (codeEngine *CodeEngineV2) ListReclamationsWithContext(ctx context.Context, listReclamationsOptions *ListReclamationsOptions) (result *ReclamationList, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReclamationList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReclamation : Get a reclamation
// Get a reclamation.
func (codeEngine *CodeEngineV2) GetReclamation(getReclamationOptions *GetReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.GetReclamationWithContext(context.Background(), getReclamationOptions)
}

// GetReclamationWithContext is an alternate form of the GetReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) GetReclamationWithContext(ctx context.Context, getReclamationOptions *GetReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReclaimReclamation : Delete a reclamation
// Delete a project reclamation to permanently delete the project.
func (codeEngine *CodeEngineV2) ReclaimReclamation(reclaimReclamationOptions *ReclaimReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.ReclaimReclamationWithContext(context.Background(), reclaimReclamationOptions)
}

// ReclaimReclamationWithContext is an alternate form of the ReclaimReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) ReclaimReclamationWithContext(ctx context.Context, reclaimReclamationOptions *ReclaimReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RestoreReclamation : Restore a project reclamation
// Restore a project reclamation. Projects that are soft-deleted can be restored within 7 days.
func (codeEngine *CodeEngineV2) RestoreReclamation(restoreReclamationOptions *RestoreReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
	return codeEngine.RestoreReclamationWithContext(context.Background(), restoreReclamationOptions)
}

// RestoreReclamationWithContext is an alternate form of the RestoreReclamation method which supports a Context parameter
func (codeEngine *CodeEngineV2) RestoreReclamationWithContext(ctx context.Context, restoreReclamationOptions *RestoreReclamationOptions) (result *Reclamation, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReclamation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// App : App is the response model for app resources.
type App struct {
	// Controls which of the system managed domain mappings will be setup for the application. Valid values are
	// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
	// application private visibility.
	CeManagedDomainMappings *string `json:"ce_managed_domain_mappings,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The port where the application listens.
	ImagePort *int64 `json:"image_port,omitempty"`

	// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses
	// unencrypted HTTP 2.
	ImageProtocol *string `json:"image_protocol,omitempty"`

	// The name of the image that is used for this application. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where
	// 'REGISTRY' and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not
	// specified, the default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// Set arguments for the application.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the application.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// The maximum number of requests that can be processed concurrently per instance.
	ScaleConcurrency *int64 `json:"scale_concurrency,omitempty"`

	// The threshold of concurrent requests per instance at which one or more additional instances are created. Use this
	// value to scale up instances based on concurrent number of requests. This option defaults to the value of the
	// 'concurrency' option, if not specified.
	ScaleConcurrencyTarget *int64 `json:"scale_concurrency_target,omitempty"`

	// The amount of CPU set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the application.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The initial number of instances that are created upon app creation or app update.
	ScaleInitialInstances *int64 `json:"scale_initial_instances,omitempty"`

	// The maximum number of instances that can be used for this application. If you set this value to '0', the application
	// scales as needed. The application scaling is limited only by the instances per the resource quota for the project of
	// your application. See https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits.
	ScaleMaxInstances *int64 `json:"scale_max_instances,omitempty"`

	// The amount of memory set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The minimum number of instances that can be used for this application.
	ScaleMinInstances *int64 `json:"scale_min_instances,omitempty"`

	// The amount of time in seconds that is allowed for a running application to respond to a request.
	ScaleRequestTimeout *int64 `json:"scale_request_timeout,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`

	// The internal version of the app instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`
}

// Constants associated with the App.CeManagedDomainMappings property.
// Controls which of the system managed domain mappings will be setup for the application. Valid values are
// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
// application private visibility.
const (
	App_CeManagedDomainMappings_Local = "local"
	App_CeManagedDomainMappings_LocalPrivate = "local_private"
	App_CeManagedDomainMappings_LocalPublic = "local_public"
)

// Constants associated with the App.ImageProtocol property.
// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses unencrypted
// HTTP 2.
const (
	App_ImageProtocol_H2c = "h2c"
	App_ImageProtocol_Http1 = "http1"
)

// Constants associated with the App.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	App_RunServiceAccount_Default = "default"
	App_RunServiceAccount_Manager = "manager"
	App_RunServiceAccount_None = "none"
	App_RunServiceAccount_Reader = "reader"
	App_RunServiceAccount_Writer = "writer"
)

// UnmarshalApp unmarshals an instance of App from the specified map of raw messages.
func UnmarshalApp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(App)
	err = core.UnmarshalPrimitive(m, "ce_managed_domain_mappings", &obj.CeManagedDomainMappings)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_port", &obj.ImagePort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_protocol", &obj.ImageProtocol)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_ref", &obj.ImageRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_secret", &obj.ImageSecret)
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
	err = core.UnmarshalPrimitive(m, "run_args", &obj.RunArgs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_as_user", &obj.RunAsUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_commands", &obj.RunCommands)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_env_vars", &obj.RunEnvVars, UnmarshalEnvVar)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_service_account", &obj.RunServiceAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_volume_mounts", &obj.RunVolumeMounts, UnmarshalVolumeMount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_concurrency", &obj.ScaleConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_concurrency_target", &obj.ScaleConcurrencyTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_cpu_limit", &obj.ScaleCpuLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_ephemeral_storage_limit", &obj.ScaleEphemeralStorageLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_initial_instances", &obj.ScaleInitialInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_max_instances", &obj.ScaleMaxInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_memory_limit", &obj.ScaleMemoryLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_min_instances", &obj.ScaleMinInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_request_timeout", &obj.ScaleRequestTimeout)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AppList : Contains a list of apps and pagination information.
type AppList struct {
	// List of all apps.
	Apps []App `json:"apps,omitempty"`

	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalAppList unmarshals an instance of AppList from the specified map of raw messages.
func UnmarshalAppList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AppList)
	err = core.UnmarshalModel(m, "apps", &obj.Apps, UnmarshalApp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
func (resp *AppList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// AppRevision : AppRevision is the response model for app revision resources.
type AppRevision struct {
	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The port where the application listens.
	ImagePort *int64 `json:"image_port,omitempty"`

	// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses
	// unencrypted HTTP 2.
	ImageProtocol *string `json:"image_protocol,omitempty"`

	// The name of the image that is used for this application. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where
	// 'REGISTRY' and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not
	// specified, the default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// Set arguments for the application.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the application.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// The maximum number of requests that can be processed concurrently per instance.
	ScaleConcurrency *int64 `json:"scale_concurrency,omitempty"`

	// The threshold of concurrent requests per instance at which one or more additional instances are created. Use this
	// value to scale up instances based on concurrent number of requests. This option defaults to the value of the
	// 'concurrency' option, if not specified.
	ScaleConcurrencyTarget *int64 `json:"scale_concurrency_target,omitempty"`

	// The amount of CPU set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the application.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The initial number of instances that are created upon app creation or app update.
	ScaleInitialInstances *int64 `json:"scale_initial_instances,omitempty"`

	// The maximum number of instances that can be used for this application. If you set this value to '0', the application
	// scales as needed. The application scaling is limited only by the instances per the resource quota for the project of
	// your application. See https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits.
	ScaleMaxInstances *int64 `json:"scale_max_instances,omitempty"`

	// The amount of memory set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The minimum number of instances that can be used for this application.
	ScaleMinInstances *int64 `json:"scale_min_instances,omitempty"`

	// The amount of time in seconds that is allowed for a running application to respond to a request.
	ScaleRequestTimeout *int64 `json:"scale_request_timeout,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// Constants associated with the AppRevision.ImageProtocol property.
// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses unencrypted
// HTTP 2.
const (
	AppRevision_ImageProtocol_H2c = "h2c"
	AppRevision_ImageProtocol_Http1 = "http1"
)

// Constants associated with the AppRevision.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	AppRevision_RunServiceAccount_Default = "default"
	AppRevision_RunServiceAccount_Manager = "manager"
	AppRevision_RunServiceAccount_None = "none"
	AppRevision_RunServiceAccount_Reader = "reader"
	AppRevision_RunServiceAccount_Writer = "writer"
)

// UnmarshalAppRevision unmarshals an instance of AppRevision from the specified map of raw messages.
func UnmarshalAppRevision(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AppRevision)
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
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
	err = core.UnmarshalPrimitive(m, "image_port", &obj.ImagePort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_protocol", &obj.ImageProtocol)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_ref", &obj.ImageRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_secret", &obj.ImageSecret)
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
	err = core.UnmarshalPrimitive(m, "run_args", &obj.RunArgs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_as_user", &obj.RunAsUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_commands", &obj.RunCommands)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_env_vars", &obj.RunEnvVars, UnmarshalEnvVar)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_service_account", &obj.RunServiceAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_volume_mounts", &obj.RunVolumeMounts, UnmarshalVolumeMount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_concurrency", &obj.ScaleConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_concurrency_target", &obj.ScaleConcurrencyTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_cpu_limit", &obj.ScaleCpuLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_ephemeral_storage_limit", &obj.ScaleEphemeralStorageLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_initial_instances", &obj.ScaleInitialInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_max_instances", &obj.ScaleMaxInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_memory_limit", &obj.ScaleMemoryLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_min_instances", &obj.ScaleMinInstances)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_request_timeout", &obj.ScaleRequestTimeout)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AppRevisionList : Contains a list of app revisions and pagination information.
type AppRevisionList struct {
	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of all app revisions.
	Revisions []AppRevision `json:"revisions,omitempty"`
}

// UnmarshalAppRevisionList unmarshals an instance of AppRevisionList from the specified map of raw messages.
func UnmarshalAppRevisionList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AppRevisionList)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
	err = core.UnmarshalModel(m, "revisions", &obj.Revisions, UnmarshalAppRevision)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *AppRevisionList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// Build : Response model for build definitions.
type Build struct {
	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

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

	// The path to the specification file that is used for build strategies for building an image.
	StrategySpecFile *string `json:"strategy_spec_file,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// Constants associated with the Build.SourceType property.
// Specifies the type of source to determine if your build source is in a repository or based on local source code.
const (
	Build_SourceType_Git = "git"
	Build_SourceType_Local = "local"
)

// UnmarshalBuild unmarshals an instance of Build from the specified map of raw messages.
func UnmarshalBuild(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Build)
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
	err = core.UnmarshalPrimitive(m, "strategy_spec_file", &obj.StrategySpecFile)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BuildList : Contains a list of builds and pagination information.
type BuildList struct {
	// List of all builds.
	Builds []Build `json:"builds,omitempty"`

	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalBuildList unmarshals an instance of BuildList from the specified map of raw messages.
func UnmarshalBuildList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BuildList)
	err = core.UnmarshalModel(m, "builds", &obj.Builds, UnmarshalBuild)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
func (resp *BuildList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// BuildRun : Response model for build run objects.
type BuildRun struct {
	// The name of the app revision with which this build run is associated.
	AppRevision *string `json:"app_revision,omitempty"`

	// The name of the build on which this build run is associated.
	Build *string `json:"build,omitempty"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

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

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The strategy to use for building the image.
	StrategyName *string `json:"strategy_name,omitempty"`

	// The size for the build, which determines the amount of resources used.  Build sizes are `small`, `medium`,
	// `large`,`xlarge`.
	StrategySize *string `json:"strategy_size,omitempty"`

	// The path to the specification file that is used for build strategies for building an image.
	StrategySpecFile *string `json:"strategy_spec_file,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// Constants associated with the BuildRun.ServiceAccount property.
// ServiceAccount refers to the serviceaccount which is used for resource control.
const (
	BuildRun_ServiceAccount_Default = "default"
	BuildRun_ServiceAccount_Manager = "manager"
	BuildRun_ServiceAccount_None = "none"
	BuildRun_ServiceAccount_Reader = "reader"
	BuildRun_ServiceAccount_Writer = "writer"
)

// Constants associated with the BuildRun.SourceType property.
// Specifies the type of source to determine if your build source is in a repository or based on local source code.
const (
	BuildRun_SourceType_Git = "git"
	BuildRun_SourceType_Local = "local"
)

// UnmarshalBuildRun unmarshals an instance of BuildRun from the specified map of raw messages.
func UnmarshalBuildRun(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BuildRun)
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
	err = core.UnmarshalPrimitive(m, "strategy_spec_file", &obj.StrategySpecFile)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BuildRunList : Contains a list of build runs and pagination information.
type BuildRunList struct {
	// List of all build runs.
	BuildRuns []BuildRun `json:"build_runs,omitempty"`

	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalBuildRunList unmarshals an instance of BuildRunList from the specified map of raw messages.
func UnmarshalBuildRunList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BuildRunList)
	err = core.UnmarshalModel(m, "build_runs", &obj.BuildRuns, UnmarshalBuildRun)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
func (resp *BuildRunList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// ConfigMap : Describes the model of a configmap.
type ConfigMap struct {
	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// UnmarshalConfigMap unmarshals an instance of ConfigMap from the specified map of raw messages.
func UnmarshalConfigMap(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigMap)
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
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigMapList : Contains a list of configmaps and pagination information.
type ConfigMapList struct {
	// List of all configmaps.
	ConfigMaps []ConfigMap `json:"config_maps,omitempty"`

	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalConfigMapList unmarshals an instance of ConfigMapList from the specified map of raw messages.
func UnmarshalConfigMapList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigMapList)
	err = core.UnmarshalModel(m, "config_maps", &obj.ConfigMaps, UnmarshalConfigMap)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
func (resp *ConfigMapList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// CreateAppOptions : The CreateApp options.
type CreateAppOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of the app. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// Controls which of the system managed domain mappings will be setup for the application. Valid values are
	// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
	// application private visibility.
	CeManagedDomainMappings *string `json:"ce_managed_domain_mappings,omitempty"`

	// The port where the application listens.
	ImagePort *int64 `json:"image_port,omitempty"`

	// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses
	// unencrypted HTTP 2.
	ImageProtocol *string `json:"image_protocol,omitempty"`

	// The name of the image that is used for this application. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where
	// 'REGISTRY' and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not
	// specified, the default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The suffix of the new revision. Use a suffix that is unique for this application.
	RevisionSuffix *string `json:"revision_suffix,omitempty"`

	// Set arguments for the application.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the application.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// The maximum number of requests that can be processed concurrently per instance.
	ScaleConcurrency *int64 `json:"scale_concurrency,omitempty"`

	// The threshold of concurrent requests per instance at which one or more additional instances are created. Use this
	// value to scale up instances based on concurrent number of requests. This option defaults to the value of the
	// 'concurrency' option, if not specified.
	ScaleConcurrencyTarget *int64 `json:"scale_concurrency_target,omitempty"`

	// The amount of CPU set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the application.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The initial number of instances that are created upon app creation or app update.
	ScaleInitialInstances *int64 `json:"scale_initial_instances,omitempty"`

	// The maximum number of instances that can be used for this application. If you set this value to '0', the application
	// scales as needed. The application scaling is limited only by the instances per the resource quota for the project of
	// your application. See https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits.
	ScaleMaxInstances *int64 `json:"scale_max_instances,omitempty"`

	// The amount of memory set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The minimum number of instances that can be used for this application.
	ScaleMinInstances *int64 `json:"scale_min_instances,omitempty"`

	// The amount of time in seconds that is allowed for a running application to respond to a request.
	ScaleRequestTimeout *int64 `json:"scale_request_timeout,omitempty"`

	// The internal version of the app instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateAppOptions.CeManagedDomainMappings property.
// Controls which of the system managed domain mappings will be setup for the application. Valid values are
// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
// application private visibility.
const (
	CreateAppOptions_CeManagedDomainMappings_Local = "local"
	CreateAppOptions_CeManagedDomainMappings_LocalPrivate = "local_private"
	CreateAppOptions_CeManagedDomainMappings_LocalPublic = "local_public"
)

// Constants associated with the CreateAppOptions.ImageProtocol property.
// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses unencrypted
// HTTP 2.
const (
	CreateAppOptions_ImageProtocol_H2c = "h2c"
	CreateAppOptions_ImageProtocol_Http1 = "http1"
)

// Constants associated with the CreateAppOptions.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	CreateAppOptions_RunServiceAccount_Default = "default"
	CreateAppOptions_RunServiceAccount_Manager = "manager"
	CreateAppOptions_RunServiceAccount_None = "none"
	CreateAppOptions_RunServiceAccount_Reader = "reader"
	CreateAppOptions_RunServiceAccount_Writer = "writer"
)

// NewCreateAppOptions : Instantiate CreateAppOptions
func (*CodeEngineV2) NewCreateAppOptions(projectGuid string, name string) *CreateAppOptions {
	return &CreateAppOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateAppOptions) SetProjectGuid(projectGuid string) *CreateAppOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateAppOptions) SetName(name string) *CreateAppOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeManagedDomainMappings : Allow user to set CeManagedDomainMappings
func (_options *CreateAppOptions) SetCeManagedDomainMappings(ceManagedDomainMappings string) *CreateAppOptions {
	_options.CeManagedDomainMappings = core.StringPtr(ceManagedDomainMappings)
	return _options
}

// SetImagePort : Allow user to set ImagePort
func (_options *CreateAppOptions) SetImagePort(imagePort int64) *CreateAppOptions {
	_options.ImagePort = core.Int64Ptr(imagePort)
	return _options
}

// SetImageProtocol : Allow user to set ImageProtocol
func (_options *CreateAppOptions) SetImageProtocol(imageProtocol string) *CreateAppOptions {
	_options.ImageProtocol = core.StringPtr(imageProtocol)
	return _options
}

// SetImageRef : Allow user to set ImageRef
func (_options *CreateAppOptions) SetImageRef(imageRef string) *CreateAppOptions {
	_options.ImageRef = core.StringPtr(imageRef)
	return _options
}

// SetImageSecret : Allow user to set ImageSecret
func (_options *CreateAppOptions) SetImageSecret(imageSecret string) *CreateAppOptions {
	_options.ImageSecret = core.StringPtr(imageSecret)
	return _options
}

// SetRevisionSuffix : Allow user to set RevisionSuffix
func (_options *CreateAppOptions) SetRevisionSuffix(revisionSuffix string) *CreateAppOptions {
	_options.RevisionSuffix = core.StringPtr(revisionSuffix)
	return _options
}

// SetRunArgs : Allow user to set RunArgs
func (_options *CreateAppOptions) SetRunArgs(runArgs []string) *CreateAppOptions {
	_options.RunArgs = runArgs
	return _options
}

// SetRunAsUser : Allow user to set RunAsUser
func (_options *CreateAppOptions) SetRunAsUser(runAsUser int64) *CreateAppOptions {
	_options.RunAsUser = core.Int64Ptr(runAsUser)
	return _options
}

// SetRunCommands : Allow user to set RunCommands
func (_options *CreateAppOptions) SetRunCommands(runCommands []string) *CreateAppOptions {
	_options.RunCommands = runCommands
	return _options
}

// SetRunEnvVars : Allow user to set RunEnvVars
func (_options *CreateAppOptions) SetRunEnvVars(runEnvVars []EnvVar) *CreateAppOptions {
	_options.RunEnvVars = runEnvVars
	return _options
}

// SetRunServiceAccount : Allow user to set RunServiceAccount
func (_options *CreateAppOptions) SetRunServiceAccount(runServiceAccount string) *CreateAppOptions {
	_options.RunServiceAccount = core.StringPtr(runServiceAccount)
	return _options
}

// SetRunVolumeMounts : Allow user to set RunVolumeMounts
func (_options *CreateAppOptions) SetRunVolumeMounts(runVolumeMounts []VolumeMount) *CreateAppOptions {
	_options.RunVolumeMounts = runVolumeMounts
	return _options
}

// SetScaleConcurrency : Allow user to set ScaleConcurrency
func (_options *CreateAppOptions) SetScaleConcurrency(scaleConcurrency int64) *CreateAppOptions {
	_options.ScaleConcurrency = core.Int64Ptr(scaleConcurrency)
	return _options
}

// SetScaleConcurrencyTarget : Allow user to set ScaleConcurrencyTarget
func (_options *CreateAppOptions) SetScaleConcurrencyTarget(scaleConcurrencyTarget int64) *CreateAppOptions {
	_options.ScaleConcurrencyTarget = core.Int64Ptr(scaleConcurrencyTarget)
	return _options
}

// SetScaleCpuLimit : Allow user to set ScaleCpuLimit
func (_options *CreateAppOptions) SetScaleCpuLimit(scaleCpuLimit string) *CreateAppOptions {
	_options.ScaleCpuLimit = core.StringPtr(scaleCpuLimit)
	return _options
}

// SetScaleEphemeralStorageLimit : Allow user to set ScaleEphemeralStorageLimit
func (_options *CreateAppOptions) SetScaleEphemeralStorageLimit(scaleEphemeralStorageLimit string) *CreateAppOptions {
	_options.ScaleEphemeralStorageLimit = core.StringPtr(scaleEphemeralStorageLimit)
	return _options
}

// SetScaleInitialInstances : Allow user to set ScaleInitialInstances
func (_options *CreateAppOptions) SetScaleInitialInstances(scaleInitialInstances int64) *CreateAppOptions {
	_options.ScaleInitialInstances = core.Int64Ptr(scaleInitialInstances)
	return _options
}

// SetScaleMaxInstances : Allow user to set ScaleMaxInstances
func (_options *CreateAppOptions) SetScaleMaxInstances(scaleMaxInstances int64) *CreateAppOptions {
	_options.ScaleMaxInstances = core.Int64Ptr(scaleMaxInstances)
	return _options
}

// SetScaleMemoryLimit : Allow user to set ScaleMemoryLimit
func (_options *CreateAppOptions) SetScaleMemoryLimit(scaleMemoryLimit string) *CreateAppOptions {
	_options.ScaleMemoryLimit = core.StringPtr(scaleMemoryLimit)
	return _options
}

// SetScaleMinInstances : Allow user to set ScaleMinInstances
func (_options *CreateAppOptions) SetScaleMinInstances(scaleMinInstances int64) *CreateAppOptions {
	_options.ScaleMinInstances = core.Int64Ptr(scaleMinInstances)
	return _options
}

// SetScaleRequestTimeout : Allow user to set ScaleRequestTimeout
func (_options *CreateAppOptions) SetScaleRequestTimeout(scaleRequestTimeout int64) *CreateAppOptions {
	_options.ScaleRequestTimeout = core.Int64Ptr(scaleRequestTimeout)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreateAppOptions) SetVersion(version string) *CreateAppOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAppOptions) SetHeaders(param map[string]string) *CreateAppOptions {
	options.Headers = param
	return options
}

// CreateBuildOptions : The CreateBuild options.
type CreateBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of the build. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

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

	// The path to the specification file that is used for build strategies for building an image.
	StrategySpecFile *string `json:"strategy_spec_file,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateBuildOptions.SourceType property.
// Specifies the type of source to determine if your build source is in a repository or based on local source code.
const (
	CreateBuildOptions_SourceType_Git = "git"
	CreateBuildOptions_SourceType_Local = "local"
)

// NewCreateBuildOptions : Instantiate CreateBuildOptions
func (*CodeEngineV2) NewCreateBuildOptions(projectGuid string, name string) *CreateBuildOptions {
	return &CreateBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateBuildOptions) SetProjectGuid(projectGuid string) *CreateBuildOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateBuildOptions) SetName(name string) *CreateBuildOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeOwnerReference : Allow user to set CeOwnerReference
func (_options *CreateBuildOptions) SetCeOwnerReference(ceOwnerReference string) *CreateBuildOptions {
	_options.CeOwnerReference = core.StringPtr(ceOwnerReference)
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

// SetStrategySpecFile : Allow user to set StrategySpecFile
func (_options *CreateBuildOptions) SetStrategySpecFile(strategySpecFile string) *CreateBuildOptions {
	_options.StrategySpecFile = core.StringPtr(strategySpecFile)
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

	// Name of the build run.
	Name *string `json:"name" validate:"required"`

	// The name of the app revision with which this build run is associated.
	AppRevision *string `json:"app_revision,omitempty"`

	// The name of the build on which this build run is associated.
	Build *string `json:"build,omitempty"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

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

	// The path to the specification file that is used for build strategies for building an image.
	StrategySpecFile *string `json:"strategy_spec_file,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateBuildrunOptions.ServiceAccount property.
// ServiceAccount refers to the serviceaccount which is used for resource control.
const (
	CreateBuildrunOptions_ServiceAccount_Default = "default"
	CreateBuildrunOptions_ServiceAccount_Manager = "manager"
	CreateBuildrunOptions_ServiceAccount_None = "none"
	CreateBuildrunOptions_ServiceAccount_Reader = "reader"
	CreateBuildrunOptions_ServiceAccount_Writer = "writer"
)

// Constants associated with the CreateBuildrunOptions.SourceType property.
// Specifies the type of source to determine if your build source is in a repository or based on local source code.
const (
	CreateBuildrunOptions_SourceType_Git = "git"
	CreateBuildrunOptions_SourceType_Local = "local"
)

// NewCreateBuildrunOptions : Instantiate CreateBuildrunOptions
func (*CodeEngineV2) NewCreateBuildrunOptions(projectGuid string, name string) *CreateBuildrunOptions {
	return &CreateBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateBuildrunOptions) SetProjectGuid(projectGuid string) *CreateBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateBuildrunOptions) SetName(name string) *CreateBuildrunOptions {
	_options.Name = core.StringPtr(name)
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

// SetStrategySpecFile : Allow user to set StrategySpecFile
func (_options *CreateBuildrunOptions) SetStrategySpecFile(strategySpecFile string) *CreateBuildrunOptions {
	_options.StrategySpecFile = core.StringPtr(strategySpecFile)
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

	// The name of the configmap. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigmapOptions : Instantiate CreateConfigmapOptions
func (*CodeEngineV2) NewCreateConfigmapOptions(projectGuid string, name string) *CreateConfigmapOptions {
	return &CreateConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateConfigmapOptions) SetProjectGuid(projectGuid string) *CreateConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateConfigmapOptions) SetName(name string) *CreateConfigmapOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetData : Allow user to set Data
func (_options *CreateConfigmapOptions) SetData(data map[string]string) *CreateConfigmapOptions {
	_options.Data = data
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigmapOptions) SetHeaders(param map[string]string) *CreateConfigmapOptions {
	options.Headers = param
	return options
}

// CreateJobOptions : The CreateJob options.
type CreateJobOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of the image that is used for this job. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where 'REGISTRY'
	// and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not specified, the
	// default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The name of the job. Use a name that is unique within the project.
	Name *string `json:"name,omitempty"`

	// Set arguments for the job.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the job.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
	// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
	// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
	RunMode *string `json:"run_mode,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges
	// like "5,12-14,23,27". Each instance can pick up its array index via environment variable JOB_INDEX. The number of
	// unique array indices specified here determines the number of job instances to run.
	ScaleArraySpec *string `json:"scale_array_spec,omitempty"`

	// The amount of CPU set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the job.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The maximum execution time in seconds for runs of the job. This option can only be specified if 'mode' is 'task'.
	ScaleMaxExecutionTime *int64 `json:"scale_max_execution_time,omitempty"`

	// The amount of memory set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The number of times to rerun an instance of the job before the job is marked as failed. This option can only be
	// specified if 'mode' is 'task'.
	ScaleRetryLimit *int64 `json:"scale_retry_limit,omitempty"`

	// The internal version of the job instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateJobOptions.RunMode property.
// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
const (
	CreateJobOptions_RunMode_Daemon = "daemon"
	CreateJobOptions_RunMode_Task = "task"
)

// Constants associated with the CreateJobOptions.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	CreateJobOptions_RunServiceAccount_Default = "default"
	CreateJobOptions_RunServiceAccount_Manager = "manager"
	CreateJobOptions_RunServiceAccount_None = "none"
	CreateJobOptions_RunServiceAccount_Reader = "reader"
	CreateJobOptions_RunServiceAccount_Writer = "writer"
)

// NewCreateJobOptions : Instantiate CreateJobOptions
func (*CodeEngineV2) NewCreateJobOptions(projectGuid string) *CreateJobOptions {
	return &CreateJobOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateJobOptions) SetProjectGuid(projectGuid string) *CreateJobOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetImageRef : Allow user to set ImageRef
func (_options *CreateJobOptions) SetImageRef(imageRef string) *CreateJobOptions {
	_options.ImageRef = core.StringPtr(imageRef)
	return _options
}

// SetImageSecret : Allow user to set ImageSecret
func (_options *CreateJobOptions) SetImageSecret(imageSecret string) *CreateJobOptions {
	_options.ImageSecret = core.StringPtr(imageSecret)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateJobOptions) SetName(name string) *CreateJobOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRunArgs : Allow user to set RunArgs
func (_options *CreateJobOptions) SetRunArgs(runArgs []string) *CreateJobOptions {
	_options.RunArgs = runArgs
	return _options
}

// SetRunAsUser : Allow user to set RunAsUser
func (_options *CreateJobOptions) SetRunAsUser(runAsUser int64) *CreateJobOptions {
	_options.RunAsUser = core.Int64Ptr(runAsUser)
	return _options
}

// SetRunCommands : Allow user to set RunCommands
func (_options *CreateJobOptions) SetRunCommands(runCommands []string) *CreateJobOptions {
	_options.RunCommands = runCommands
	return _options
}

// SetRunEnvVars : Allow user to set RunEnvVars
func (_options *CreateJobOptions) SetRunEnvVars(runEnvVars []EnvVar) *CreateJobOptions {
	_options.RunEnvVars = runEnvVars
	return _options
}

// SetRunMode : Allow user to set RunMode
func (_options *CreateJobOptions) SetRunMode(runMode string) *CreateJobOptions {
	_options.RunMode = core.StringPtr(runMode)
	return _options
}

// SetRunServiceAccount : Allow user to set RunServiceAccount
func (_options *CreateJobOptions) SetRunServiceAccount(runServiceAccount string) *CreateJobOptions {
	_options.RunServiceAccount = core.StringPtr(runServiceAccount)
	return _options
}

// SetRunVolumeMounts : Allow user to set RunVolumeMounts
func (_options *CreateJobOptions) SetRunVolumeMounts(runVolumeMounts []VolumeMount) *CreateJobOptions {
	_options.RunVolumeMounts = runVolumeMounts
	return _options
}

// SetScaleArraySpec : Allow user to set ScaleArraySpec
func (_options *CreateJobOptions) SetScaleArraySpec(scaleArraySpec string) *CreateJobOptions {
	_options.ScaleArraySpec = core.StringPtr(scaleArraySpec)
	return _options
}

// SetScaleCpuLimit : Allow user to set ScaleCpuLimit
func (_options *CreateJobOptions) SetScaleCpuLimit(scaleCpuLimit string) *CreateJobOptions {
	_options.ScaleCpuLimit = core.StringPtr(scaleCpuLimit)
	return _options
}

// SetScaleEphemeralStorageLimit : Allow user to set ScaleEphemeralStorageLimit
func (_options *CreateJobOptions) SetScaleEphemeralStorageLimit(scaleEphemeralStorageLimit string) *CreateJobOptions {
	_options.ScaleEphemeralStorageLimit = core.StringPtr(scaleEphemeralStorageLimit)
	return _options
}

// SetScaleMaxExecutionTime : Allow user to set ScaleMaxExecutionTime
func (_options *CreateJobOptions) SetScaleMaxExecutionTime(scaleMaxExecutionTime int64) *CreateJobOptions {
	_options.ScaleMaxExecutionTime = core.Int64Ptr(scaleMaxExecutionTime)
	return _options
}

// SetScaleMemoryLimit : Allow user to set ScaleMemoryLimit
func (_options *CreateJobOptions) SetScaleMemoryLimit(scaleMemoryLimit string) *CreateJobOptions {
	_options.ScaleMemoryLimit = core.StringPtr(scaleMemoryLimit)
	return _options
}

// SetScaleRetryLimit : Allow user to set ScaleRetryLimit
func (_options *CreateJobOptions) SetScaleRetryLimit(scaleRetryLimit int64) *CreateJobOptions {
	_options.ScaleRetryLimit = core.Int64Ptr(scaleRetryLimit)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreateJobOptions) SetVersion(version string) *CreateJobOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateJobOptions) SetHeaders(param map[string]string) *CreateJobOptions {
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
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of the secret.
	Name *string `json:"name" validate:"required"`

	// List of bound Code Engine components.
	CeComponents []string `json:"ce_components,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associated with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Name of the service credential associated with the secret.
	ResourcekeyName *string `json:"resourcekey_name,omitempty"`

	// Role of the service credential.
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential.
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateSecretOptions.Format property.
// Specify the format of the secret.
const (
	CreateSecretOptions_Format_BasicAuth = "basic_auth"
	CreateSecretOptions_Format_Generic = "generic"
	CreateSecretOptions_Format_Other = "other"
	CreateSecretOptions_Format_Registry = "registry"
	CreateSecretOptions_Format_ServiceAccess = "service_access"
	CreateSecretOptions_Format_SshAuth = "ssh_auth"
	CreateSecretOptions_Format_Tls = "tls"
)

// NewCreateSecretOptions : Instantiate CreateSecretOptions
func (*CodeEngineV2) NewCreateSecretOptions(projectGuid string, name string) *CreateSecretOptions {
	return &CreateSecretOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *CreateSecretOptions) SetProjectGuid(projectGuid string) *CreateSecretOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateSecretOptions) SetName(name string) *CreateSecretOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeComponents : Allow user to set CeComponents
func (_options *CreateSecretOptions) SetCeComponents(ceComponents []string) *CreateSecretOptions {
	_options.CeComponents = ceComponents
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

// SetResourcekeyName : Allow user to set ResourcekeyName
func (_options *CreateSecretOptions) SetResourcekeyName(resourcekeyName string) *CreateSecretOptions {
	_options.ResourcekeyName = core.StringPtr(resourcekeyName)
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

// SetHeaders : Allow user to set Headers
func (options *CreateSecretOptions) SetHeaders(param map[string]string) *CreateSecretOptions {
	options.Headers = param
	return options
}

// DeleteAppOptions : The DeleteApp options.
type DeleteAppOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAppOptions : Instantiate DeleteAppOptions
func (*CodeEngineV2) NewDeleteAppOptions(projectGuid string, appName string) *DeleteAppOptions {
	return &DeleteAppOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteAppOptions) SetProjectGuid(projectGuid string) *DeleteAppOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *DeleteAppOptions) SetAppName(appName string) *DeleteAppOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAppOptions) SetHeaders(param map[string]string) *DeleteAppOptions {
	options.Headers = param
	return options
}

// DeleteAppRevisionOptions : The DeleteAppRevision options.
type DeleteAppRevisionOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// The name of your application revision.
	RevisionName *string `json:"revision_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAppRevisionOptions : Instantiate DeleteAppRevisionOptions
func (*CodeEngineV2) NewDeleteAppRevisionOptions(projectGuid string, appName string, revisionName string) *DeleteAppRevisionOptions {
	return &DeleteAppRevisionOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
		RevisionName: core.StringPtr(revisionName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteAppRevisionOptions) SetProjectGuid(projectGuid string) *DeleteAppRevisionOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *DeleteAppRevisionOptions) SetAppName(appName string) *DeleteAppRevisionOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetRevisionName : Allow user to set RevisionName
func (_options *DeleteAppRevisionOptions) SetRevisionName(revisionName string) *DeleteAppRevisionOptions {
	_options.RevisionName = core.StringPtr(revisionName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAppRevisionOptions) SetHeaders(param map[string]string) *DeleteAppRevisionOptions {
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
	BuildRunName *string `json:"build_run_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBuildrunOptions : Instantiate DeleteBuildrunOptions
func (*CodeEngineV2) NewDeleteBuildrunOptions(projectGuid string, buildRunName string) *DeleteBuildrunOptions {
	return &DeleteBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildRunName: core.StringPtr(buildRunName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteBuildrunOptions) SetProjectGuid(projectGuid string) *DeleteBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildRunName : Allow user to set BuildRunName
func (_options *DeleteBuildrunOptions) SetBuildRunName(buildRunName string) *DeleteBuildrunOptions {
	_options.BuildRunName = core.StringPtr(buildRunName)
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

	// The name of your config map.
	ConfigMapName *string `json:"config_map_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigmapOptions : Instantiate DeleteConfigmapOptions
func (*CodeEngineV2) NewDeleteConfigmapOptions(projectGuid string, configMapName string) *DeleteConfigmapOptions {
	return &DeleteConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigMapName: core.StringPtr(configMapName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteConfigmapOptions) SetProjectGuid(projectGuid string) *DeleteConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigMapName : Allow user to set ConfigMapName
func (_options *DeleteConfigmapOptions) SetConfigMapName(configMapName string) *DeleteConfigmapOptions {
	_options.ConfigMapName = core.StringPtr(configMapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigmapOptions) SetHeaders(param map[string]string) *DeleteConfigmapOptions {
	options.Headers = param
	return options
}

// DeleteJobOptions : The DeleteJob options.
type DeleteJobOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your job.
	JobName *string `json:"job_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteJobOptions : Instantiate DeleteJobOptions
func (*CodeEngineV2) NewDeleteJobOptions(projectGuid string, jobName string) *DeleteJobOptions {
	return &DeleteJobOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		JobName: core.StringPtr(jobName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *DeleteJobOptions) SetProjectGuid(projectGuid string) *DeleteJobOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetJobName : Allow user to set JobName
func (_options *DeleteJobOptions) SetJobName(jobName string) *DeleteJobOptions {
	_options.JobName = core.StringPtr(jobName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteJobOptions) SetHeaders(param map[string]string) *DeleteJobOptions {
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
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your secret.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSecretOptions : Instantiate DeleteSecretOptions
func (*CodeEngineV2) NewDeleteSecretOptions(projectGuid string, secretName string) *DeleteSecretOptions {
	return &DeleteSecretOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
	}
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

// EnvVar : EnvVar described an environent variable.
type EnvVar struct {
	// The key to reference as environment variable.
	Key *string `json:"key,omitempty"`

	// The name of the environment variable.
	Name *string `json:"name,omitempty"`

	// A prefix that can be added to all keys of a full secret or config map reference.
	Prefix *string `json:"prefix,omitempty"`

	// The name of the secret or config map.
	Ref *string `json:"ref,omitempty"`

	// Specify the type of the environment variable. Allowed types are: 'literal', 'config_map_key_ref',
	// 'config_map_full_ref', 'secret_key_ref', 'secret_full_ref'.
	Type *string `json:"type,omitempty"`

	// The literal value of the environment variable.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the EnvVar.Type property.
// Specify the type of the environment variable. Allowed types are: 'literal', 'config_map_key_ref',
// 'config_map_full_ref', 'secret_key_ref', 'secret_full_ref'.
const (
	EnvVar_Type_ConfigMapFullRef = "config_map_full_ref"
	EnvVar_Type_ConfigMapKeyRef = "config_map_key_ref"
	EnvVar_Type_Literal = "literal"
	EnvVar_Type_SecretFullRef = "secret_full_ref"
	EnvVar_Type_SecretKeyRef = "secret_key_ref"
)

// UnmarshalEnvVar unmarshals an instance of EnvVar from the specified map of raw messages.
func UnmarshalEnvVar(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnvVar)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ref", &obj.Ref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAppOptions : The GetApp options.
type GetAppOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAppOptions : Instantiate GetAppOptions
func (*CodeEngineV2) NewGetAppOptions(projectGuid string, appName string) *GetAppOptions {
	return &GetAppOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetAppOptions) SetProjectGuid(projectGuid string) *GetAppOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *GetAppOptions) SetAppName(appName string) *GetAppOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAppOptions) SetHeaders(param map[string]string) *GetAppOptions {
	options.Headers = param
	return options
}

// GetAppRevisionOptions : The GetAppRevision options.
type GetAppRevisionOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// The name of your application revision.
	RevisionName *string `json:"revision_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAppRevisionOptions : Instantiate GetAppRevisionOptions
func (*CodeEngineV2) NewGetAppRevisionOptions(projectGuid string, appName string, revisionName string) *GetAppRevisionOptions {
	return &GetAppRevisionOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
		RevisionName: core.StringPtr(revisionName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetAppRevisionOptions) SetProjectGuid(projectGuid string) *GetAppRevisionOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *GetAppRevisionOptions) SetAppName(appName string) *GetAppRevisionOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetRevisionName : Allow user to set RevisionName
func (_options *GetAppRevisionOptions) SetRevisionName(revisionName string) *GetAppRevisionOptions {
	_options.RevisionName = core.StringPtr(revisionName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAppRevisionOptions) SetHeaders(param map[string]string) *GetAppRevisionOptions {
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
	BuildRunName *string `json:"build_run_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBuildrunOptions : Instantiate GetBuildrunOptions
func (*CodeEngineV2) NewGetBuildrunOptions(projectGuid string, buildRunName string) *GetBuildrunOptions {
	return &GetBuildrunOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildRunName: core.StringPtr(buildRunName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetBuildrunOptions) SetProjectGuid(projectGuid string) *GetBuildrunOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetBuildRunName : Allow user to set BuildRunName
func (_options *GetBuildrunOptions) SetBuildRunName(buildRunName string) *GetBuildrunOptions {
	_options.BuildRunName = core.StringPtr(buildRunName)
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

	// The name of your config map.
	ConfigMapName *string `json:"config_map_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigmapOptions : Instantiate GetConfigmapOptions
func (*CodeEngineV2) NewGetConfigmapOptions(projectGuid string, configMapName string) *GetConfigmapOptions {
	return &GetConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigMapName: core.StringPtr(configMapName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetConfigmapOptions) SetProjectGuid(projectGuid string) *GetConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigMapName : Allow user to set ConfigMapName
func (_options *GetConfigmapOptions) SetConfigMapName(configMapName string) *GetConfigmapOptions {
	_options.ConfigMapName = core.StringPtr(configMapName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigmapOptions) SetHeaders(param map[string]string) *GetConfigmapOptions {
	options.Headers = param
	return options
}

// GetJobOptions : The GetJob options.
type GetJobOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your job.
	JobName *string `json:"job_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetJobOptions : Instantiate GetJobOptions
func (*CodeEngineV2) NewGetJobOptions(projectGuid string, jobName string) *GetJobOptions {
	return &GetJobOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		JobName: core.StringPtr(jobName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *GetJobOptions) SetProjectGuid(projectGuid string) *GetJobOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetJobName : Allow user to set JobName
func (_options *GetJobOptions) SetJobName(jobName string) *GetJobOptions {
	_options.JobName = core.StringPtr(jobName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetJobOptions) SetHeaders(param map[string]string) *GetJobOptions {
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
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your secret.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSecretOptions : Instantiate GetSecretOptions
func (*CodeEngineV2) NewGetSecretOptions(projectGuid string, secretName string) *GetSecretOptions {
	return &GetSecretOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
	}
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

// Job : Job is the response model for job resources.
type Job struct {
	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Detailed information on the status.
	Details *string `json:"details,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the image that is used for this job. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where 'REGISTRY'
	// and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not specified, the
	// default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The reason to provide more context for the status.
	Reason *string `json:"reason,omitempty"`

	// Set arguments for the job.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the job.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
	// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
	// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
	RunMode *string `json:"run_mode,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges
	// like "5,12-14,23,27". Each instance can pick up its array index via environment variable JOB_INDEX. The number of
	// unique array indices specified here determines the number of job instances to run.
	ScaleArraySpec *string `json:"scale_array_spec,omitempty"`

	// The amount of CPU set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the job.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The maximum execution time in seconds for runs of the job. This option can only be specified if 'mode' is 'task'.
	ScaleMaxExecutionTime *int64 `json:"scale_max_execution_time,omitempty"`

	// The amount of memory set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The number of times to rerun an instance of the job before the job is marked as failed. This option can only be
	// specified if 'mode' is 'task'.
	ScaleRetryLimit *int64 `json:"scale_retry_limit,omitempty"`

	// The current state of the Code Engine resource.
	Status *string `json:"status,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`

	// The internal version of the job instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`
}

// Constants associated with the Job.RunMode property.
// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
const (
	Job_RunMode_Daemon = "daemon"
	Job_RunMode_Task = "task"
)

// Constants associated with the Job.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	Job_RunServiceAccount_Default = "default"
	Job_RunServiceAccount_Manager = "manager"
	Job_RunServiceAccount_None = "none"
	Job_RunServiceAccount_Reader = "reader"
	Job_RunServiceAccount_Writer = "writer"
)

// UnmarshalJob unmarshals an instance of Job from the specified map of raw messages.
func UnmarshalJob(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Job)
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
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
	err = core.UnmarshalPrimitive(m, "image_ref", &obj.ImageRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_secret", &obj.ImageSecret)
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
	err = core.UnmarshalPrimitive(m, "run_args", &obj.RunArgs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_as_user", &obj.RunAsUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_commands", &obj.RunCommands)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_env_vars", &obj.RunEnvVars, UnmarshalEnvVar)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_mode", &obj.RunMode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_service_account", &obj.RunServiceAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_volume_mounts", &obj.RunVolumeMounts, UnmarshalVolumeMount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_array_spec", &obj.ScaleArraySpec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_cpu_limit", &obj.ScaleCpuLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_ephemeral_storage_limit", &obj.ScaleEphemeralStorageLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_max_execution_time", &obj.ScaleMaxExecutionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_memory_limit", &obj.ScaleMemoryLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale_retry_limit", &obj.ScaleRetryLimit)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobList : Contains a list of jobs and pagination information.
type JobList struct {
	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// List of all jobs.
	Jobs []Job `json:"jobs,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`
}

// UnmarshalJobList unmarshals an instance of JobList from the specified map of raw messages.
func UnmarshalJobList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobList)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "jobs", &obj.Jobs, UnmarshalJob)
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
func (resp *JobList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// ListAppRevisionsOptions : The ListAppRevisions options.
type ListAppRevisionsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// The maximum number of apps per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAppRevisionsOptions : Instantiate ListAppRevisionsOptions
func (*CodeEngineV2) NewListAppRevisionsOptions(projectGuid string, appName string) *ListAppRevisionsOptions {
	return &ListAppRevisionsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListAppRevisionsOptions) SetProjectGuid(projectGuid string) *ListAppRevisionsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *ListAppRevisionsOptions) SetAppName(appName string) *ListAppRevisionsOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAppRevisionsOptions) SetLimit(limit int64) *ListAppRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListAppRevisionsOptions) SetStart(start string) *ListAppRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAppRevisionsOptions) SetHeaders(param map[string]string) *ListAppRevisionsOptions {
	options.Headers = param
	return options
}

// ListAppsOptions : The ListApps options.
type ListAppsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of apps per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAppsOptions : Instantiate ListAppsOptions
func (*CodeEngineV2) NewListAppsOptions(projectGuid string) *ListAppsOptions {
	return &ListAppsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListAppsOptions) SetProjectGuid(projectGuid string) *ListAppsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAppsOptions) SetLimit(limit int64) *ListAppsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListAppsOptions) SetStart(start string) *ListAppsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAppsOptions) SetHeaders(param map[string]string) *ListAppsOptions {
	options.Headers = param
	return options
}

// ListBuildrunsOptions : The ListBuildruns options.
type ListBuildrunsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of build runs per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
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

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
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

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
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

// ListJobsOptions : The ListJobs options.
type ListJobsOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of jobs per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListJobsOptions : Instantiate ListJobsOptions
func (*CodeEngineV2) NewListJobsOptions(projectGuid string) *ListJobsOptions {
	return &ListJobsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListJobsOptions) SetProjectGuid(projectGuid string) *ListJobsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListJobsOptions) SetLimit(limit int64) *ListJobsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListJobsOptions) SetStart(start string) *ListJobsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListJobsOptions) SetHeaders(param map[string]string) *ListJobsOptions {
	options.Headers = param
	return options
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {
	// The maximum number of projects per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
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

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
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
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The maximum number of configmaps per page.
	Limit *int64 `json:"limit,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'start' query parameter in the 'next_url' field of the
	// operation response.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSecretsOptions : Instantiate ListSecretsOptions
func (*CodeEngineV2) NewListSecretsOptions(projectGuid string) *ListSecretsOptions {
	return &ListSecretsOptions{
		ProjectGuid: core.StringPtr(projectGuid),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *ListSecretsOptions) SetProjectGuid(projectGuid string) *ListSecretsOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListSecretsOptions) SetLimit(limit int64) *ListSecretsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListSecretsOptions) SetStart(start string) *ListSecretsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSecretsOptions) SetHeaders(param map[string]string) *ListSecretsOptions {
	options.Headers = param
	return options
}

// Project : Describes the model of a project.
type Project struct {
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

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// UnmarshalProject unmarshals an instance of Project from the specified map of raw messages.
func UnmarshalProject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Project)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectList : Contains a list of projects and pagination information.
type ProjectList struct {
	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of projects.
	Projects []Project `json:"projects,omitempty"`
}

// UnmarshalProjectList unmarshals an instance of ProjectList from the specified map of raw messages.
func UnmarshalProjectList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectList)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ProjectList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
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

// Reclamation : Describes the model of a reclamation.
type Reclamation struct {
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

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// UnmarshalReclamation unmarshals an instance of Reclamation from the specified map of raw messages.
func UnmarshalReclamation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Reclamation)
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
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReclamationList : Contains a list of reclamations and pagination information.
type ReclamationList struct {
	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of all project reclamations.
	Reclamations []Reclamation `json:"reclamations,omitempty"`
}

// UnmarshalReclamationList unmarshals an instance of ReclamationList from the specified map of raw messages.
func UnmarshalReclamationList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReclamationList)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
	err = core.UnmarshalModel(m, "reclamations", &obj.Reclamations, UnmarshalReclamation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ReclamationList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
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

// Secret : Describes the model of a secret.
type Secret struct {
	// List of bound Code Engine components.
	CeComponents []string `json:"ce_components,omitempty"`

	// The date when the resource was created.
	Created *string `json:"created,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// The identifier of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associated with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Name of the service credential associated with the secret.
	ResourcekeyName *string `json:"resourcekey_name,omitempty"`

	// Role of the service credential.
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential.
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`
}

// Constants associated with the Secret.Format property.
// Specify the format of the secret.
const (
	Secret_Format_BasicAuth = "basic_auth"
	Secret_Format_Generic = "generic"
	Secret_Format_Other = "other"
	Secret_Format_Registry = "registry"
	Secret_Format_ServiceAccess = "service_access"
	Secret_Format_SshAuth = "ssh_auth"
	Secret_Format_Tls = "tls"
)

// UnmarshalSecret unmarshals an instance of Secret from the specified map of raw messages.
func UnmarshalSecret(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Secret)
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
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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
	err = core.UnmarshalPrimitive(m, "resourcekey_name", &obj.ResourcekeyName)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecretList : List of secret resources.
type SecretList struct {
	// Describes properties needed to retrieve the first page of a result list.
	First *PaginationListFirstMetadata `json:"first,omitempty"`

	// Maximum number of resources per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Describes properties needed to retrieve the next page of a result list.
	Next *PaginationListNextMetadata `json:"next,omitempty"`

	// List of Secrets.
	Secrets []Secret `json:"secrets,omitempty"`
}

// UnmarshalSecretList unmarshals an instance of SecretList from the specified map of raw messages.
func UnmarshalSecretList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecretList)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationListFirstMetadata)
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
	err = core.UnmarshalModel(m, "secrets", &obj.Secrets, UnmarshalSecret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *SecretList) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// UpdateAppOptions : The UpdateApp options.
type UpdateAppOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your application.
	AppName *string `json:"app_name" validate:"required,ne="`

	// The name of the app. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// Controls which of the system managed domain mappings will be setup for the application. Valid values are
	// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
	// application private visibility.
	CeManagedDomainMappings *string `json:"ce_managed_domain_mappings,omitempty"`

	// The port where the application listens.
	ImagePort *int64 `json:"image_port,omitempty"`

	// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses
	// unencrypted HTTP 2.
	ImageProtocol *string `json:"image_protocol,omitempty"`

	// The name of the image that is used for this application. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where
	// 'REGISTRY' and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not
	// specified, the default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The suffix of the new revision. Use a suffix that is unique for this application.
	RevisionSuffix *string `json:"revision_suffix,omitempty"`

	// Set arguments for the application.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the application.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// The maximum number of requests that can be processed concurrently per instance.
	ScaleConcurrency *int64 `json:"scale_concurrency,omitempty"`

	// The threshold of concurrent requests per instance at which one or more additional instances are created. Use this
	// value to scale up instances based on concurrent number of requests. This option defaults to the value of the
	// 'concurrency' option, if not specified.
	ScaleConcurrencyTarget *int64 `json:"scale_concurrency_target,omitempty"`

	// The amount of CPU set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the application.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The initial number of instances that are created upon app creation or app update.
	ScaleInitialInstances *int64 `json:"scale_initial_instances,omitempty"`

	// The maximum number of instances that can be used for this application. If you set this value to '0', the application
	// scales as needed. The application scaling is limited only by the instances per the resource quota for the project of
	// your application. See https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits.
	ScaleMaxInstances *int64 `json:"scale_max_instances,omitempty"`

	// The amount of memory set for the instance of the application. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The minimum number of instances that can be used for this application.
	ScaleMinInstances *int64 `json:"scale_min_instances,omitempty"`

	// The amount of time in seconds that is allowed for a running application to respond to a request.
	ScaleRequestTimeout *int64 `json:"scale_request_timeout,omitempty"`

	// The internal version of the app instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateAppOptions.CeManagedDomainMappings property.
// Controls which of the system managed domain mappings will be setup for the application. Valid values are
// 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports
// application private visibility.
const (
	UpdateAppOptions_CeManagedDomainMappings_Local = "local"
	UpdateAppOptions_CeManagedDomainMappings_LocalPrivate = "local_private"
	UpdateAppOptions_CeManagedDomainMappings_LocalPublic = "local_public"
)

// Constants associated with the UpdateAppOptions.ImageProtocol property.
// Specifies the protocol that the image uses. For 'http1' the image uses HTTP 1.1. For 'h2c' the image uses unencrypted
// HTTP 2.
const (
	UpdateAppOptions_ImageProtocol_H2c = "h2c"
	UpdateAppOptions_ImageProtocol_Http1 = "http1"
)

// Constants associated with the UpdateAppOptions.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	UpdateAppOptions_RunServiceAccount_Default = "default"
	UpdateAppOptions_RunServiceAccount_Manager = "manager"
	UpdateAppOptions_RunServiceAccount_None = "none"
	UpdateAppOptions_RunServiceAccount_Reader = "reader"
	UpdateAppOptions_RunServiceAccount_Writer = "writer"
)

// NewUpdateAppOptions : Instantiate UpdateAppOptions
func (*CodeEngineV2) NewUpdateAppOptions(projectGuid string, appName string, name string) *UpdateAppOptions {
	return &UpdateAppOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		AppName: core.StringPtr(appName),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateAppOptions) SetProjectGuid(projectGuid string) *UpdateAppOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetAppName : Allow user to set AppName
func (_options *UpdateAppOptions) SetAppName(appName string) *UpdateAppOptions {
	_options.AppName = core.StringPtr(appName)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateAppOptions) SetName(name string) *UpdateAppOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeManagedDomainMappings : Allow user to set CeManagedDomainMappings
func (_options *UpdateAppOptions) SetCeManagedDomainMappings(ceManagedDomainMappings string) *UpdateAppOptions {
	_options.CeManagedDomainMappings = core.StringPtr(ceManagedDomainMappings)
	return _options
}

// SetImagePort : Allow user to set ImagePort
func (_options *UpdateAppOptions) SetImagePort(imagePort int64) *UpdateAppOptions {
	_options.ImagePort = core.Int64Ptr(imagePort)
	return _options
}

// SetImageProtocol : Allow user to set ImageProtocol
func (_options *UpdateAppOptions) SetImageProtocol(imageProtocol string) *UpdateAppOptions {
	_options.ImageProtocol = core.StringPtr(imageProtocol)
	return _options
}

// SetImageRef : Allow user to set ImageRef
func (_options *UpdateAppOptions) SetImageRef(imageRef string) *UpdateAppOptions {
	_options.ImageRef = core.StringPtr(imageRef)
	return _options
}

// SetImageSecret : Allow user to set ImageSecret
func (_options *UpdateAppOptions) SetImageSecret(imageSecret string) *UpdateAppOptions {
	_options.ImageSecret = core.StringPtr(imageSecret)
	return _options
}

// SetRevisionSuffix : Allow user to set RevisionSuffix
func (_options *UpdateAppOptions) SetRevisionSuffix(revisionSuffix string) *UpdateAppOptions {
	_options.RevisionSuffix = core.StringPtr(revisionSuffix)
	return _options
}

// SetRunArgs : Allow user to set RunArgs
func (_options *UpdateAppOptions) SetRunArgs(runArgs []string) *UpdateAppOptions {
	_options.RunArgs = runArgs
	return _options
}

// SetRunAsUser : Allow user to set RunAsUser
func (_options *UpdateAppOptions) SetRunAsUser(runAsUser int64) *UpdateAppOptions {
	_options.RunAsUser = core.Int64Ptr(runAsUser)
	return _options
}

// SetRunCommands : Allow user to set RunCommands
func (_options *UpdateAppOptions) SetRunCommands(runCommands []string) *UpdateAppOptions {
	_options.RunCommands = runCommands
	return _options
}

// SetRunEnvVars : Allow user to set RunEnvVars
func (_options *UpdateAppOptions) SetRunEnvVars(runEnvVars []EnvVar) *UpdateAppOptions {
	_options.RunEnvVars = runEnvVars
	return _options
}

// SetRunServiceAccount : Allow user to set RunServiceAccount
func (_options *UpdateAppOptions) SetRunServiceAccount(runServiceAccount string) *UpdateAppOptions {
	_options.RunServiceAccount = core.StringPtr(runServiceAccount)
	return _options
}

// SetRunVolumeMounts : Allow user to set RunVolumeMounts
func (_options *UpdateAppOptions) SetRunVolumeMounts(runVolumeMounts []VolumeMount) *UpdateAppOptions {
	_options.RunVolumeMounts = runVolumeMounts
	return _options
}

// SetScaleConcurrency : Allow user to set ScaleConcurrency
func (_options *UpdateAppOptions) SetScaleConcurrency(scaleConcurrency int64) *UpdateAppOptions {
	_options.ScaleConcurrency = core.Int64Ptr(scaleConcurrency)
	return _options
}

// SetScaleConcurrencyTarget : Allow user to set ScaleConcurrencyTarget
func (_options *UpdateAppOptions) SetScaleConcurrencyTarget(scaleConcurrencyTarget int64) *UpdateAppOptions {
	_options.ScaleConcurrencyTarget = core.Int64Ptr(scaleConcurrencyTarget)
	return _options
}

// SetScaleCpuLimit : Allow user to set ScaleCpuLimit
func (_options *UpdateAppOptions) SetScaleCpuLimit(scaleCpuLimit string) *UpdateAppOptions {
	_options.ScaleCpuLimit = core.StringPtr(scaleCpuLimit)
	return _options
}

// SetScaleEphemeralStorageLimit : Allow user to set ScaleEphemeralStorageLimit
func (_options *UpdateAppOptions) SetScaleEphemeralStorageLimit(scaleEphemeralStorageLimit string) *UpdateAppOptions {
	_options.ScaleEphemeralStorageLimit = core.StringPtr(scaleEphemeralStorageLimit)
	return _options
}

// SetScaleInitialInstances : Allow user to set ScaleInitialInstances
func (_options *UpdateAppOptions) SetScaleInitialInstances(scaleInitialInstances int64) *UpdateAppOptions {
	_options.ScaleInitialInstances = core.Int64Ptr(scaleInitialInstances)
	return _options
}

// SetScaleMaxInstances : Allow user to set ScaleMaxInstances
func (_options *UpdateAppOptions) SetScaleMaxInstances(scaleMaxInstances int64) *UpdateAppOptions {
	_options.ScaleMaxInstances = core.Int64Ptr(scaleMaxInstances)
	return _options
}

// SetScaleMemoryLimit : Allow user to set ScaleMemoryLimit
func (_options *UpdateAppOptions) SetScaleMemoryLimit(scaleMemoryLimit string) *UpdateAppOptions {
	_options.ScaleMemoryLimit = core.StringPtr(scaleMemoryLimit)
	return _options
}

// SetScaleMinInstances : Allow user to set ScaleMinInstances
func (_options *UpdateAppOptions) SetScaleMinInstances(scaleMinInstances int64) *UpdateAppOptions {
	_options.ScaleMinInstances = core.Int64Ptr(scaleMinInstances)
	return _options
}

// SetScaleRequestTimeout : Allow user to set ScaleRequestTimeout
func (_options *UpdateAppOptions) SetScaleRequestTimeout(scaleRequestTimeout int64) *UpdateAppOptions {
	_options.ScaleRequestTimeout = core.Int64Ptr(scaleRequestTimeout)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *UpdateAppOptions) SetVersion(version string) *UpdateAppOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAppOptions) SetHeaders(param map[string]string) *UpdateAppOptions {
	options.Headers = param
	return options
}

// UpdateBuildOptions : The UpdateBuild options.
type UpdateBuildOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your build.
	BuildName *string `json:"build_name" validate:"required,ne="`

	// The name of the build. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// The resource that owns this build, such as a Code Engine application or job.
	CeOwnerReference *string `json:"ce_owner_reference,omitempty"`

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

	// The path to the specification file that is used for build strategies for building an image.
	StrategySpecFile *string `json:"strategy_spec_file,omitempty"`

	// The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
	Timeout *int64 `json:"timeout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateBuildOptions.SourceType property.
// Specifies the type of source to determine if your build source is in a repository or based on local source code.
const (
	UpdateBuildOptions_SourceType_Git = "git"
	UpdateBuildOptions_SourceType_Local = "local"
)

// NewUpdateBuildOptions : Instantiate UpdateBuildOptions
func (*CodeEngineV2) NewUpdateBuildOptions(projectGuid string, buildName string, name string) *UpdateBuildOptions {
	return &UpdateBuildOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		BuildName: core.StringPtr(buildName),
		Name: core.StringPtr(name),
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

// SetName : Allow user to set Name
func (_options *UpdateBuildOptions) SetName(name string) *UpdateBuildOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeOwnerReference : Allow user to set CeOwnerReference
func (_options *UpdateBuildOptions) SetCeOwnerReference(ceOwnerReference string) *UpdateBuildOptions {
	_options.CeOwnerReference = core.StringPtr(ceOwnerReference)
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

// SetStrategySpecFile : Allow user to set StrategySpecFile
func (_options *UpdateBuildOptions) SetStrategySpecFile(strategySpecFile string) *UpdateBuildOptions {
	_options.StrategySpecFile = core.StringPtr(strategySpecFile)
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

	// The name of your config map.
	ConfigMapName *string `json:"config_map_name" validate:"required,ne="`

	// The name of the configmap. Use a name that is unique within the project.
	Name *string `json:"name" validate:"required"`

	// The key-value pair for the configmap. Values must be specified in `KEY=VALUE` format.
	Data map[string]string `json:"data,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigmapOptions : Instantiate UpdateConfigmapOptions
func (*CodeEngineV2) NewUpdateConfigmapOptions(projectGuid string, configMapName string, name string) *UpdateConfigmapOptions {
	return &UpdateConfigmapOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		ConfigMapName: core.StringPtr(configMapName),
		Name: core.StringPtr(name),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateConfigmapOptions) SetProjectGuid(projectGuid string) *UpdateConfigmapOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetConfigMapName : Allow user to set ConfigMapName
func (_options *UpdateConfigmapOptions) SetConfigMapName(configMapName string) *UpdateConfigmapOptions {
	_options.ConfigMapName = core.StringPtr(configMapName)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateConfigmapOptions) SetName(name string) *UpdateConfigmapOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetData : Allow user to set Data
func (_options *UpdateConfigmapOptions) SetData(data map[string]string) *UpdateConfigmapOptions {
	_options.Data = data
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigmapOptions) SetHeaders(param map[string]string) *UpdateConfigmapOptions {
	options.Headers = param
	return options
}

// UpdateJobOptions : The UpdateJob options.
type UpdateJobOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your job.
	JobName *string `json:"job_name" validate:"required,ne="`

	// The name of the image that is used for this job. The format is 'REGISTRY/NAMESPACE/REPOSITORY:TAG' where 'REGISTRY'
	// and 'TAG' are optional. If 'REGISTRY' is not specified, the default is 'docker.io'. If 'TAG' is not specified, the
	// default is 'latest'.
	ImageRef *string `json:"image_ref,omitempty"`

	// The name of the image registry access secret. The image registry access secret is used to authenticate with a
	// private registry when you download the container image.
	ImageSecret *string `json:"image_secret,omitempty"`

	// The name of the job. Use a name that is unique within the project.
	Name *string `json:"name,omitempty"`

	// Set arguments for the job.
	RunArgs []string `json:"run_args,omitempty"`

	// The user ID (UID) to run the application (e.g., 1001).
	RunAsUser *int64 `json:"run_as_user,omitempty"`

	// Set commands for the job.
	RunCommands []string `json:"run_commands,omitempty"`

	// Mount a configmap or a secret.
	RunEnvVars []EnvVar `json:"run_env_vars,omitempty"`

	// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
	// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
	// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
	RunMode *string `json:"run_mode,omitempty"`

	// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
	// 'reader', and 'writer'.
	RunServiceAccount *string `json:"run_service_account,omitempty"`

	// Mount a configmap or a secret.
	RunVolumeMounts []VolumeMount `json:"run_volume_mounts,omitempty"`

	// Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges
	// like "5,12-14,23,27". Each instance can pick up its array index via environment variable JOB_INDEX. The number of
	// unique array indices specified here determines the number of job instances to run.
	ScaleArraySpec *string `json:"scale_array_spec,omitempty"`

	// The amount of CPU set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleCpuLimit *string `json:"scale_cpu_limit,omitempty"`

	// The amount of ephemeral storage to set for the instance of the job.
	ScaleEphemeralStorageLimit *string `json:"scale_ephemeral_storage_limit,omitempty"`

	// The maximum execution time in seconds for runs of the job. This option can only be specified if 'mode' is 'task'.
	ScaleMaxExecutionTime *int64 `json:"scale_max_execution_time,omitempty"`

	// The amount of memory set for the instance of the job. For valid values see
	// https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo.
	ScaleMemoryLimit *string `json:"scale_memory_limit,omitempty"`

	// The number of times to rerun an instance of the job before the job is marked as failed. This option can only be
	// specified if 'mode' is 'task'.
	ScaleRetryLimit *int64 `json:"scale_retry_limit,omitempty"`

	// The internal version of the job instance, which is used to achieve optimistic concurrency.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateJobOptions.RunMode property.
// The mode for runs of the job. Valid values are 'task' and 'daemon'. In 'task' mode, the 'max_execution_time' and
// 'retry_limit' options apply. In 'daemon' mode, since there is no timeout and failed instances are restarted
// indefinitely, the 'max_execution_time' and 'retry_limit' options are not allowed.
const (
	UpdateJobOptions_RunMode_Daemon = "daemon"
	UpdateJobOptions_RunMode_Task = "task"
)

// Constants associated with the UpdateJobOptions.RunServiceAccount property.
// The name of the service account. For built-in service accounts, you can use the shortened names 'manager', 'none',
// 'reader', and 'writer'.
const (
	UpdateJobOptions_RunServiceAccount_Default = "default"
	UpdateJobOptions_RunServiceAccount_Manager = "manager"
	UpdateJobOptions_RunServiceAccount_None = "none"
	UpdateJobOptions_RunServiceAccount_Reader = "reader"
	UpdateJobOptions_RunServiceAccount_Writer = "writer"
)

// NewUpdateJobOptions : Instantiate UpdateJobOptions
func (*CodeEngineV2) NewUpdateJobOptions(projectGuid string, jobName string) *UpdateJobOptions {
	return &UpdateJobOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		JobName: core.StringPtr(jobName),
	}
}

// SetProjectGuid : Allow user to set ProjectGuid
func (_options *UpdateJobOptions) SetProjectGuid(projectGuid string) *UpdateJobOptions {
	_options.ProjectGuid = core.StringPtr(projectGuid)
	return _options
}

// SetJobName : Allow user to set JobName
func (_options *UpdateJobOptions) SetJobName(jobName string) *UpdateJobOptions {
	_options.JobName = core.StringPtr(jobName)
	return _options
}

// SetImageRef : Allow user to set ImageRef
func (_options *UpdateJobOptions) SetImageRef(imageRef string) *UpdateJobOptions {
	_options.ImageRef = core.StringPtr(imageRef)
	return _options
}

// SetImageSecret : Allow user to set ImageSecret
func (_options *UpdateJobOptions) SetImageSecret(imageSecret string) *UpdateJobOptions {
	_options.ImageSecret = core.StringPtr(imageSecret)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateJobOptions) SetName(name string) *UpdateJobOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRunArgs : Allow user to set RunArgs
func (_options *UpdateJobOptions) SetRunArgs(runArgs []string) *UpdateJobOptions {
	_options.RunArgs = runArgs
	return _options
}

// SetRunAsUser : Allow user to set RunAsUser
func (_options *UpdateJobOptions) SetRunAsUser(runAsUser int64) *UpdateJobOptions {
	_options.RunAsUser = core.Int64Ptr(runAsUser)
	return _options
}

// SetRunCommands : Allow user to set RunCommands
func (_options *UpdateJobOptions) SetRunCommands(runCommands []string) *UpdateJobOptions {
	_options.RunCommands = runCommands
	return _options
}

// SetRunEnvVars : Allow user to set RunEnvVars
func (_options *UpdateJobOptions) SetRunEnvVars(runEnvVars []EnvVar) *UpdateJobOptions {
	_options.RunEnvVars = runEnvVars
	return _options
}

// SetRunMode : Allow user to set RunMode
func (_options *UpdateJobOptions) SetRunMode(runMode string) *UpdateJobOptions {
	_options.RunMode = core.StringPtr(runMode)
	return _options
}

// SetRunServiceAccount : Allow user to set RunServiceAccount
func (_options *UpdateJobOptions) SetRunServiceAccount(runServiceAccount string) *UpdateJobOptions {
	_options.RunServiceAccount = core.StringPtr(runServiceAccount)
	return _options
}

// SetRunVolumeMounts : Allow user to set RunVolumeMounts
func (_options *UpdateJobOptions) SetRunVolumeMounts(runVolumeMounts []VolumeMount) *UpdateJobOptions {
	_options.RunVolumeMounts = runVolumeMounts
	return _options
}

// SetScaleArraySpec : Allow user to set ScaleArraySpec
func (_options *UpdateJobOptions) SetScaleArraySpec(scaleArraySpec string) *UpdateJobOptions {
	_options.ScaleArraySpec = core.StringPtr(scaleArraySpec)
	return _options
}

// SetScaleCpuLimit : Allow user to set ScaleCpuLimit
func (_options *UpdateJobOptions) SetScaleCpuLimit(scaleCpuLimit string) *UpdateJobOptions {
	_options.ScaleCpuLimit = core.StringPtr(scaleCpuLimit)
	return _options
}

// SetScaleEphemeralStorageLimit : Allow user to set ScaleEphemeralStorageLimit
func (_options *UpdateJobOptions) SetScaleEphemeralStorageLimit(scaleEphemeralStorageLimit string) *UpdateJobOptions {
	_options.ScaleEphemeralStorageLimit = core.StringPtr(scaleEphemeralStorageLimit)
	return _options
}

// SetScaleMaxExecutionTime : Allow user to set ScaleMaxExecutionTime
func (_options *UpdateJobOptions) SetScaleMaxExecutionTime(scaleMaxExecutionTime int64) *UpdateJobOptions {
	_options.ScaleMaxExecutionTime = core.Int64Ptr(scaleMaxExecutionTime)
	return _options
}

// SetScaleMemoryLimit : Allow user to set ScaleMemoryLimit
func (_options *UpdateJobOptions) SetScaleMemoryLimit(scaleMemoryLimit string) *UpdateJobOptions {
	_options.ScaleMemoryLimit = core.StringPtr(scaleMemoryLimit)
	return _options
}

// SetScaleRetryLimit : Allow user to set ScaleRetryLimit
func (_options *UpdateJobOptions) SetScaleRetryLimit(scaleRetryLimit int64) *UpdateJobOptions {
	_options.ScaleRetryLimit = core.Int64Ptr(scaleRetryLimit)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *UpdateJobOptions) SetVersion(version string) *UpdateJobOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateJobOptions) SetHeaders(param map[string]string) *UpdateJobOptions {
	options.Headers = param
	return options
}

// UpdateSecretOptions : The UpdateSecret options.
type UpdateSecretOptions struct {
	// The ID of the project.
	ProjectGuid *string `json:"project_guid" validate:"required,ne="`

	// The name of your secret.
	SecretName *string `json:"secret_name" validate:"required,ne="`

	// The name of the secret.
	Name *string `json:"name" validate:"required"`

	// List of bound Code Engine components.
	CeComponents []string `json:"ce_components,omitempty"`

	// Data container that allows to specify config parameters and their values as a key-value map.
	Data map[string]string `json:"data,omitempty"`

	// Specify the format of the secret.
	Format *string `json:"format,omitempty"`

	// ID of the IBM Cloud service instance associated with the secret.
	ResourceID *string `json:"resource_id,omitempty"`

	// Type of IBM Cloud service associated with the secret.
	ResourceType *string `json:"resource_type,omitempty"`

	// ID of the service credential associated with the secret.
	ResourcekeyID *string `json:"resourcekey_id,omitempty"`

	// Name of the service credential associated with the secret.
	ResourcekeyName *string `json:"resourcekey_name,omitempty"`

	// Role of the service credential.
	Role *string `json:"role,omitempty"`

	// CRN of a Service ID used to create the service credential.
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateSecretOptions.Format property.
// Specify the format of the secret.
const (
	UpdateSecretOptions_Format_BasicAuth = "basic_auth"
	UpdateSecretOptions_Format_Generic = "generic"
	UpdateSecretOptions_Format_Other = "other"
	UpdateSecretOptions_Format_Registry = "registry"
	UpdateSecretOptions_Format_ServiceAccess = "service_access"
	UpdateSecretOptions_Format_SshAuth = "ssh_auth"
	UpdateSecretOptions_Format_Tls = "tls"
)

// NewUpdateSecretOptions : Instantiate UpdateSecretOptions
func (*CodeEngineV2) NewUpdateSecretOptions(projectGuid string, secretName string, name string) *UpdateSecretOptions {
	return &UpdateSecretOptions{
		ProjectGuid: core.StringPtr(projectGuid),
		SecretName: core.StringPtr(secretName),
		Name: core.StringPtr(name),
	}
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

// SetName : Allow user to set Name
func (_options *UpdateSecretOptions) SetName(name string) *UpdateSecretOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetCeComponents : Allow user to set CeComponents
func (_options *UpdateSecretOptions) SetCeComponents(ceComponents []string) *UpdateSecretOptions {
	_options.CeComponents = ceComponents
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

// SetResourcekeyName : Allow user to set ResourcekeyName
func (_options *UpdateSecretOptions) SetResourcekeyName(resourcekeyName string) *UpdateSecretOptions {
	_options.ResourcekeyName = core.StringPtr(resourcekeyName)
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

// SetHeaders : Allow user to set Headers
func (options *UpdateSecretOptions) SetHeaders(param map[string]string) *UpdateSecretOptions {
	options.Headers = param
	return options
}

// VolumeMount : VolumeMount described a volume mount.
type VolumeMount struct {
	// The path that should be mounted.
	MountPath *string `json:"mount_path,omitempty"`

	// The name of the mount.
	Name *string `json:"name,omitempty"`

	// The name of the referenced secret or config map.
	Ref *string `json:"ref,omitempty"`

	// Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the VolumeMount.Type property.
// Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.
const (
	VolumeMount_Type_ConfigMap = "config_map"
	VolumeMount_Type_Secret = "secret"
)

// UnmarshalVolumeMount unmarshals an instance of VolumeMount from the specified map of raw messages.
func UnmarshalVolumeMount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMount)
	err = core.UnmarshalPrimitive(m, "mount_path", &obj.MountPath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ref", &obj.Ref)
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

// PaginationListFirstMetadata : Describes properties needed to retrieve the first page of a result list.
type PaginationListFirstMetadata struct {
	// Href that points to the first page.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPaginationListFirstMetadata unmarshals an instance of PaginationListFirstMetadata from the specified map of raw messages.
func UnmarshalPaginationListFirstMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationListFirstMetadata)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationListNextMetadata : Describes properties needed to retrieve the next page of a result list.
type PaginationListNextMetadata struct {
	// Href that points to the next page.
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
func (pager *ProjectsPager) GetNextWithContext(ctx context.Context) (page []Project, err error) {
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
func (pager *ProjectsPager) GetAllWithContext(ctx context.Context) (allItems []Project, err error) {
	for pager.HasNext() {
		var nextPage []Project
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetNext() (page []Project, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetAll() (allItems []Project, err error) {
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
func (pager *BuildsPager) GetNextWithContext(ctx context.Context) (page []Build, err error) {
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
func (pager *BuildsPager) GetAllWithContext(ctx context.Context) (allItems []Build, err error) {
	for pager.HasNext() {
		var nextPage []Build
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BuildsPager) GetNext() (page []Build, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BuildsPager) GetAll() (allItems []Build, err error) {
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
func (pager *BuildrunsPager) GetNextWithContext(ctx context.Context) (page []BuildRun, err error) {
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
	page = result.BuildRuns

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *BuildrunsPager) GetAllWithContext(ctx context.Context) (allItems []BuildRun, err error) {
	for pager.HasNext() {
		var nextPage []BuildRun
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BuildrunsPager) GetNext() (page []BuildRun, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BuildrunsPager) GetAll() (allItems []BuildRun, err error) {
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
func (pager *ConfigmapsPager) GetNextWithContext(ctx context.Context) (page []ConfigMap, err error) {
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
	page = result.ConfigMaps

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ConfigmapsPager) GetAllWithContext(ctx context.Context) (allItems []ConfigMap, err error) {
	for pager.HasNext() {
		var nextPage []ConfigMap
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ConfigmapsPager) GetNext() (page []ConfigMap, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ConfigmapsPager) GetAll() (allItems []ConfigMap, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// SecretsPager can be used to simplify the use of the "ListSecrets" method.
//
type SecretsPager struct {
	hasNext bool
	options *ListSecretsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewSecretsPager returns a new SecretsPager instance.
func (codeEngine *CodeEngineV2) NewSecretsPager(options *ListSecretsOptions) (pager *SecretsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListSecretsOptions = *options
	pager = &SecretsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *SecretsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *SecretsPager) GetNextWithContext(ctx context.Context) (page []Secret, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListSecretsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Secrets

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *SecretsPager) GetAllWithContext(ctx context.Context) (allItems []Secret, err error) {
	for pager.HasNext() {
		var nextPage []Secret
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *SecretsPager) GetNext() (page []Secret, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *SecretsPager) GetAll() (allItems []Secret, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// AppsPager can be used to simplify the use of the "ListApps" method.
//
type AppsPager struct {
	hasNext bool
	options *ListAppsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewAppsPager returns a new AppsPager instance.
func (codeEngine *CodeEngineV2) NewAppsPager(options *ListAppsOptions) (pager *AppsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListAppsOptions = *options
	pager = &AppsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AppsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AppsPager) GetNextWithContext(ctx context.Context) (page []App, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListAppsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Apps

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AppsPager) GetAllWithContext(ctx context.Context) (allItems []App, err error) {
	for pager.HasNext() {
		var nextPage []App
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AppsPager) GetNext() (page []App, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AppsPager) GetAll() (allItems []App, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// AppRevisionsPager can be used to simplify the use of the "ListAppRevisions" method.
//
type AppRevisionsPager struct {
	hasNext bool
	options *ListAppRevisionsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewAppRevisionsPager returns a new AppRevisionsPager instance.
func (codeEngine *CodeEngineV2) NewAppRevisionsPager(options *ListAppRevisionsOptions) (pager *AppRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListAppRevisionsOptions = *options
	pager = &AppRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AppRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AppRevisionsPager) GetNextWithContext(ctx context.Context) (page []AppRevision, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListAppRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Revisions

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AppRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []AppRevision, err error) {
	for pager.HasNext() {
		var nextPage []AppRevision
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AppRevisionsPager) GetNext() (page []AppRevision, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AppRevisionsPager) GetAll() (allItems []AppRevision, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// JobsPager can be used to simplify the use of the "ListJobs" method.
//
type JobsPager struct {
	hasNext bool
	options *ListJobsOptions
	client  *CodeEngineV2
	pageContext struct {
		next *string
	}
}

// NewJobsPager returns a new JobsPager instance.
func (codeEngine *CodeEngineV2) NewJobsPager(options *ListJobsOptions) (pager *JobsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListJobsOptions = *options
	pager = &JobsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  codeEngine,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *JobsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *JobsPager) GetNextWithContext(ctx context.Context) (page []Job, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListJobsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Jobs

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *JobsPager) GetAllWithContext(ctx context.Context) (allItems []Job, err error) {
	for pager.HasNext() {
		var nextPage []Job
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *JobsPager) GetNext() (page []Job, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *JobsPager) GetAll() (allItems []Job, err error) {
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
func (pager *ReclamationsPager) GetNextWithContext(ctx context.Context) (page []Reclamation, err error) {
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
func (pager *ReclamationsPager) GetAllWithContext(ctx context.Context) (allItems []Reclamation, err error) {
	for pager.HasNext() {
		var nextPage []Reclamation
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ReclamationsPager) GetNext() (page []Reclamation, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ReclamationsPager) GetAll() (allItems []Reclamation, err error) {
	return pager.GetAllWithContext(context.Background())
}
