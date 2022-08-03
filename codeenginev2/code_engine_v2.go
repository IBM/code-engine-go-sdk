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
	if createProjectV2Options.AccountID != nil {
		body["account_id"] = createProjectV2Options.AccountID
	}
	if createProjectV2Options.Created != nil {
		body["created"] = createProjectV2Options.Created
	}
	if createProjectV2Options.Crn != nil {
		body["crn"] = createProjectV2Options.Crn
	}
	if createProjectV2Options.Details != nil {
		body["details"] = createProjectV2Options.Details
	}
	if createProjectV2Options.ID != nil {
		body["id"] = createProjectV2Options.ID
	}
	if createProjectV2Options.Name != nil {
		body["name"] = createProjectV2Options.Name
	}
	if createProjectV2Options.Reason != nil {
		body["reason"] = createProjectV2Options.Reason
	}
	if createProjectV2Options.Region != nil {
		body["region"] = createProjectV2Options.Region
	}
	if createProjectV2Options.ResourceGroupID != nil {
		body["resource_group_id"] = createProjectV2Options.ResourceGroupID
	}
	if createProjectV2Options.Status != nil {
		body["status"] = createProjectV2Options.Status
	}
	if createProjectV2Options.Type != nil {
		body["type"] = createProjectV2Options.Type
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

// CreateProjectV2Options : The CreateProjectV2 options.
type CreateProjectV2Options struct {
	// Refresh Token.
	RefreshToken *string `json:"Refresh-Token" validate:"required"`

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

// SetAccountID : Allow user to set AccountID
func (_options *CreateProjectV2Options) SetAccountID(accountID string) *CreateProjectV2Options {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetCreated : Allow user to set Created
func (_options *CreateProjectV2Options) SetCreated(created string) *CreateProjectV2Options {
	_options.Created = core.StringPtr(created)
	return _options
}

// SetCrn : Allow user to set Crn
func (_options *CreateProjectV2Options) SetCrn(crn string) *CreateProjectV2Options {
	_options.Crn = core.StringPtr(crn)
	return _options
}

// SetDetails : Allow user to set Details
func (_options *CreateProjectV2Options) SetDetails(details string) *CreateProjectV2Options {
	_options.Details = core.StringPtr(details)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateProjectV2Options) SetID(id string) *CreateProjectV2Options {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateProjectV2Options) SetName(name string) *CreateProjectV2Options {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetReason : Allow user to set Reason
func (_options *CreateProjectV2Options) SetReason(reason string) *CreateProjectV2Options {
	_options.Reason = core.StringPtr(reason)
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

// SetStatus : Allow user to set Status
func (_options *CreateProjectV2Options) SetStatus(status string) *CreateProjectV2Options {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateProjectV2Options) SetType(typeVar string) *CreateProjectV2Options {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectV2Options) SetHeaders(param map[string]string) *CreateProjectV2Options {
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
