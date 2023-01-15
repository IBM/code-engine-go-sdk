/**
 * (C) Copyright IBM Corp. 2023.
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

package codeenginev2_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`CodeEngineV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(codeEngineService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(codeEngineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
				URL: "https://codeenginev2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(codeEngineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CODE_ENGINE_URL": "https://codeenginev2/api",
				"CODE_ENGINE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2UsingExternalConfig(&codeenginev2.CodeEngineV2Options{
				})
				Expect(codeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := codeEngineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != codeEngineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(codeEngineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(codeEngineService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2UsingExternalConfig(&codeenginev2.CodeEngineV2Options{
					URL: "https://testService/api",
				})
				Expect(codeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := codeEngineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != codeEngineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(codeEngineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(codeEngineService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2UsingExternalConfig(&codeenginev2.CodeEngineV2Options{
				})
				err := codeEngineService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := codeEngineService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != codeEngineService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(codeEngineService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(codeEngineService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CODE_ENGINE_URL": "https://codeenginev2/api",
				"CODE_ENGINE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2UsingExternalConfig(&codeenginev2.CodeEngineV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(codeEngineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CODE_ENGINE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2UsingExternalConfig(&codeenginev2.CodeEngineV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(codeEngineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = codeenginev2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions) - Operation response error`, func() {
		listProjectsPath := "/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjects with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(codeenginev2.ListProjectsOptions)
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions)`, func() {
		listProjectsPath := "/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}]}`)
				}))
			})
			It(`Invoke ListProjects successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(codeenginev2.ListProjectsOptions)
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}]}`)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(codeenginev2.ListProjectsOptions)
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjects with error: Operation request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(codeenginev2.ListProjectsOptions)
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(codeenginev2.ListProjectsOptions)
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.ProjectList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.ProjectList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"account_id":"4329073d16d2f3663f74bfa955259139","created_at":"2021-03-29T12:18:13.992359829Z","crn":"crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b","id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","name":"project-name","region":"us-east","resource_group_id":"5c49eabcf5e85881a37e2d100a33b3df","resource_type":"project_v2","status":"active"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"account_id":"4329073d16d2f3663f74bfa955259139","created_at":"2021-03-29T12:18:13.992359829Z","crn":"crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b","id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","name":"project-name","region":"us-east","resource_group_id":"5c49eabcf5e85881a37e2d100a33b3df","resource_type":"project_v2","status":"active"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ProjectsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listProjectsOptionsModel := &codeenginev2.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.Project
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ProjectsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listProjectsOptionsModel := &codeenginev2.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions) - Operation response error`, func() {
		createProjectPath := "/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProject with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
		createProjectPath := "/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}`)
				}))
			})
			It(`Invoke CreateProject successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}`)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectOptions model with no property values
				createProjectOptionsModelNew := new(codeenginev2.CreateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateProject(createProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions) - Operation response error`, func() {
		getProjectPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProject with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(codeenginev2.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
		getProjectPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}`)
				}))
			})
			It(`Invoke GetProject successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(codeenginev2.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created_at": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:4e49b3e0-27a8-48d2-a784-c7ee48bb863b::", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "name": "project-name", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "resource_type": "project_v2", "status": "active"}`)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(codeenginev2.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProject with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(codeenginev2.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectOptions model with no property values
				getProjectOptionsModelNew := new(codeenginev2.GetProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetProject(getProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(codeenginev2.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		deleteProjectPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteProject successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(codeenginev2.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProject with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(codeenginev2.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectOptions model with no property values
				deleteProjectOptionsModelNew := new(codeenginev2.DeleteProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteProject(deleteProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListApps(listAppsOptions *ListAppsOptions) - Operation response error`, func() {
		listAppsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListApps with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppsOptions model
				listAppsOptionsModel := new(codeenginev2.ListAppsOptions)
				listAppsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppsOptionsModel.Start = core.StringPtr("testString")
				listAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListApps(listAppsOptions *ListAppsOptions)`, func() {
		listAppsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apps": [{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListApps successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListAppsOptions model
				listAppsOptionsModel := new(codeenginev2.ListAppsOptions)
				listAppsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppsOptionsModel.Start = core.StringPtr("testString")
				listAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListAppsWithContext(ctx, listAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListAppsWithContext(ctx, listAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apps": [{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListApps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListApps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAppsOptions model
				listAppsOptionsModel := new(codeenginev2.ListAppsOptions)
				listAppsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppsOptionsModel.Start = core.StringPtr("testString")
				listAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListApps with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppsOptions model
				listAppsOptionsModel := new(codeenginev2.ListAppsOptions)
				listAppsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppsOptionsModel.Start = core.StringPtr("testString")
				listAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAppsOptions model with no property values
				listAppsOptionsModelNew := new(codeenginev2.ListAppsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListApps(listAppsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListApps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppsOptions model
				listAppsOptionsModel := new(codeenginev2.ListAppsOptions)
				listAppsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppsOptionsModel.Start = core.StringPtr("testString")
				listAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListApps(listAppsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.AppList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.AppList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"apps":[{"created_at":"2022-09-13T11:41:35+02:00","endpoint":"https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud","endpoint_internal":"http://my-app.vg67hzldruk.svc.cluster.local","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","managed_domain_mappings":"local_public","name":"my-app","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"app_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"ready","status_details":{"latest_created_revision":"my-app-00001","latest_ready_revision":"my-app-00001","reason":"ready"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"apps":[{"created_at":"2022-09-13T11:41:35+02:00","endpoint":"https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud","endpoint_internal":"http://my-app.vg67hzldruk.svc.cluster.local","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","managed_domain_mappings":"local_public","name":"my-app","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"app_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"ready","status_details":{"latest_created_revision":"my-app-00001","latest_ready_revision":"my-app-00001","reason":"ready"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AppsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listAppsOptionsModel := &codeenginev2.ListAppsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewAppsPager(listAppsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.App
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AppsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listAppsOptionsModel := &codeenginev2.ListAppsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewAppsPager(listAppsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateApp(createAppOptions *CreateAppOptions) - Operation response error`, func() {
		createAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAppPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApp with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.ManagedDomainMappings = core.StringPtr("local_public")
				createAppOptionsModel.RunArguments = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApp(createAppOptions *CreateAppOptions)`, func() {
		createAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAppPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke CreateApp successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.ManagedDomainMappings = core.StringPtr("local_public")
				createAppOptionsModel.RunArguments = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateAppWithContext(ctx, createAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateAppWithContext(ctx, createAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAppPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke CreateApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.ManagedDomainMappings = core.StringPtr("local_public")
				createAppOptionsModel.RunArguments = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateApp with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.ManagedDomainMappings = core.StringPtr("local_public")
				createAppOptionsModel.RunArguments = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAppOptions model with no property values
				createAppOptionsModelNew := new(codeenginev2.CreateAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateApp(createAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.ManagedDomainMappings = core.StringPtr("local_public")
				createAppOptionsModel.RunArguments = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateApp(createAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApp(getAppOptions *GetAppOptions) - Operation response error`, func() {
		getAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApp with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppOptions model
				getAppOptionsModel := new(codeenginev2.GetAppOptions)
				getAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.Name = core.StringPtr("my-app")
				getAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApp(getAppOptions *GetAppOptions)`, func() {
		getAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke GetApp successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetAppOptions model
				getAppOptionsModel := new(codeenginev2.GetAppOptions)
				getAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.Name = core.StringPtr("my-app")
				getAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetAppWithContext(ctx, getAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetAppWithContext(ctx, getAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke GetApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAppOptions model
				getAppOptionsModel := new(codeenginev2.GetAppOptions)
				getAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.Name = core.StringPtr("my-app")
				getAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApp with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppOptions model
				getAppOptionsModel := new(codeenginev2.GetAppOptions)
				getAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.Name = core.StringPtr("my-app")
				getAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAppOptions model with no property values
				getAppOptionsModelNew := new(codeenginev2.GetAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetApp(getAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppOptions model
				getAppOptionsModel := new(codeenginev2.GetAppOptions)
				getAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.Name = core.StringPtr("my-app")
				getAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetApp(getAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteApp(deleteAppOptions *DeleteAppOptions)`, func() {
		deleteAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAppPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAppOptions model
				deleteAppOptionsModel := new(codeenginev2.DeleteAppOptions)
				deleteAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.Name = core.StringPtr("my-app")
				deleteAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteApp(deleteAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApp with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteAppOptions model
				deleteAppOptionsModel := new(codeenginev2.DeleteAppOptions)
				deleteAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.Name = core.StringPtr("my-app")
				deleteAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteApp(deleteAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAppOptions model with no property values
				deleteAppOptionsModelNew := new(codeenginev2.DeleteAppOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteApp(deleteAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApp(updateAppOptions *UpdateAppOptions) - Operation response error`, func() {
		updateAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAppPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateApp with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the AppPatch model
				appPatchModel := new(codeenginev2.AppPatch)
				appPatchModel.ImagePort = core.Int64Ptr(int64(8080))
				appPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				appPatchModel.ImageSecret = core.StringPtr("my-secret")
				appPatchModel.ManagedDomainMappings = core.StringPtr("local_public")
				appPatchModel.RunArguments = []string{"testString"}
				appPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				appPatchModel.RunCommands = []string{"testString"}
				appPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				appPatchModel.RunServiceAccount = core.StringPtr("default")
				appPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				appPatchModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				appPatchModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				appPatchModel.ScaleCpuLimit = core.StringPtr("1")
				appPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				appPatchModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				appPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				appPatchModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.IfMatch = core.StringPtr("testString")
				updateAppOptionsModel.App = appPatchModelAsPatch
				updateAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApp(updateAppOptions *UpdateAppOptions)`, func() {
		updateAppPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAppPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke UpdateApp successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the AppPatch model
				appPatchModel := new(codeenginev2.AppPatch)
				appPatchModel.ImagePort = core.Int64Ptr(int64(8080))
				appPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				appPatchModel.ImageSecret = core.StringPtr("my-secret")
				appPatchModel.ManagedDomainMappings = core.StringPtr("local_public")
				appPatchModel.RunArguments = []string{"testString"}
				appPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				appPatchModel.RunCommands = []string{"testString"}
				appPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				appPatchModel.RunServiceAccount = core.StringPtr("default")
				appPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				appPatchModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				appPatchModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				appPatchModel.ScaleCpuLimit = core.StringPtr("1")
				appPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				appPatchModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				appPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				appPatchModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.IfMatch = core.StringPtr("testString")
				updateAppOptionsModel.App = appPatchModelAsPatch
				updateAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateAppWithContext(ctx, updateAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateAppWithContext(ctx, updateAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAppPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "endpoint": "https://my-app.vg67hzldruk.eu-de.codeengine.appdomain.cloud", "endpoint_internal": "http://my-app.vg67hzldruk.svc.cluster.local", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "managed_domain_mappings": "local_public", "name": "my-app", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"latest_created_revision": "my-app-00001", "latest_ready_revision": "my-app-00001", "reason": "ready"}}`)
				}))
			})
			It(`Invoke UpdateApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the AppPatch model
				appPatchModel := new(codeenginev2.AppPatch)
				appPatchModel.ImagePort = core.Int64Ptr(int64(8080))
				appPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				appPatchModel.ImageSecret = core.StringPtr("my-secret")
				appPatchModel.ManagedDomainMappings = core.StringPtr("local_public")
				appPatchModel.RunArguments = []string{"testString"}
				appPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				appPatchModel.RunCommands = []string{"testString"}
				appPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				appPatchModel.RunServiceAccount = core.StringPtr("default")
				appPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				appPatchModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				appPatchModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				appPatchModel.ScaleCpuLimit = core.StringPtr("1")
				appPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				appPatchModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				appPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				appPatchModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.IfMatch = core.StringPtr("testString")
				updateAppOptionsModel.App = appPatchModelAsPatch
				updateAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateApp with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the AppPatch model
				appPatchModel := new(codeenginev2.AppPatch)
				appPatchModel.ImagePort = core.Int64Ptr(int64(8080))
				appPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				appPatchModel.ImageSecret = core.StringPtr("my-secret")
				appPatchModel.ManagedDomainMappings = core.StringPtr("local_public")
				appPatchModel.RunArguments = []string{"testString"}
				appPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				appPatchModel.RunCommands = []string{"testString"}
				appPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				appPatchModel.RunServiceAccount = core.StringPtr("default")
				appPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				appPatchModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				appPatchModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				appPatchModel.ScaleCpuLimit = core.StringPtr("1")
				appPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				appPatchModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				appPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				appPatchModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.IfMatch = core.StringPtr("testString")
				updateAppOptionsModel.App = appPatchModelAsPatch
				updateAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAppOptions model with no property values
				updateAppOptionsModelNew := new(codeenginev2.UpdateAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateApp(updateAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the AppPatch model
				appPatchModel := new(codeenginev2.AppPatch)
				appPatchModel.ImagePort = core.Int64Ptr(int64(8080))
				appPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				appPatchModel.ImageSecret = core.StringPtr("my-secret")
				appPatchModel.ManagedDomainMappings = core.StringPtr("local_public")
				appPatchModel.RunArguments = []string{"testString"}
				appPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				appPatchModel.RunCommands = []string{"testString"}
				appPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				appPatchModel.RunServiceAccount = core.StringPtr("default")
				appPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				appPatchModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				appPatchModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				appPatchModel.ScaleCpuLimit = core.StringPtr("1")
				appPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				appPatchModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				appPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				appPatchModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				appPatchModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.IfMatch = core.StringPtr("testString")
				updateAppOptionsModel.App = appPatchModelAsPatch
				updateAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateApp(updateAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAppRevisions(listAppRevisionsOptions *ListAppRevisionsOptions) - Operation response error`, func() {
		listAppRevisionsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app/revisions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppRevisionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAppRevisions with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppRevisionsOptions model
				listAppRevisionsOptionsModel := new(codeenginev2.ListAppRevisionsOptions)
				listAppRevisionsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.AppName = core.StringPtr("my-app")
				listAppRevisionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppRevisionsOptionsModel.Start = core.StringPtr("testString")
				listAppRevisionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAppRevisions(listAppRevisionsOptions *ListAppRevisionsOptions)`, func() {
		listAppRevisionsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app/revisions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppRevisionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "revisions": [{"app_name": "my-app", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-app-00001", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_revision_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"actual_instances": 1, "reason": "ready"}}]}`)
				}))
			})
			It(`Invoke ListAppRevisions successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListAppRevisionsOptions model
				listAppRevisionsOptionsModel := new(codeenginev2.ListAppRevisionsOptions)
				listAppRevisionsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.AppName = core.StringPtr("my-app")
				listAppRevisionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppRevisionsOptionsModel.Start = core.StringPtr("testString")
				listAppRevisionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListAppRevisionsWithContext(ctx, listAppRevisionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListAppRevisionsWithContext(ctx, listAppRevisionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppRevisionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "revisions": [{"app_name": "my-app", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-app-00001", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_revision_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"actual_instances": 1, "reason": "ready"}}]}`)
				}))
			})
			It(`Invoke ListAppRevisions successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListAppRevisions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAppRevisionsOptions model
				listAppRevisionsOptionsModel := new(codeenginev2.ListAppRevisionsOptions)
				listAppRevisionsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.AppName = core.StringPtr("my-app")
				listAppRevisionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppRevisionsOptionsModel.Start = core.StringPtr("testString")
				listAppRevisionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAppRevisions with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppRevisionsOptions model
				listAppRevisionsOptionsModel := new(codeenginev2.ListAppRevisionsOptions)
				listAppRevisionsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.AppName = core.StringPtr("my-app")
				listAppRevisionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppRevisionsOptionsModel.Start = core.StringPtr("testString")
				listAppRevisionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAppRevisionsOptions model with no property values
				listAppRevisionsOptionsModelNew := new(codeenginev2.ListAppRevisionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListAppRevisions(listAppRevisionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAppRevisions successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListAppRevisionsOptions model
				listAppRevisionsOptionsModel := new(codeenginev2.ListAppRevisionsOptions)
				listAppRevisionsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.AppName = core.StringPtr("my-app")
				listAppRevisionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listAppRevisionsOptionsModel.Start = core.StringPtr("testString")
				listAppRevisionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListAppRevisions(listAppRevisionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.AppRevisionList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.AppRevisionList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAppRevisionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"revisions":[{"app_name":"my-app","created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","name":"my-app-00001","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"app_revision_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"ready","status_details":{"actual_instances":1,"reason":"ready"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"revisions":[{"app_name":"my-app","created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","name":"my-app-00001","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"app_revision_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"ready","status_details":{"actual_instances":1,"reason":"ready"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AppRevisionsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listAppRevisionsOptionsModel := &codeenginev2.ListAppRevisionsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					AppName: core.StringPtr("my-app"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewAppRevisionsPager(listAppRevisionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.AppRevision
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AppRevisionsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listAppRevisionsOptionsModel := &codeenginev2.ListAppRevisionsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					AppName: core.StringPtr("my-app"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewAppRevisionsPager(listAppRevisionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAppRevision(getAppRevisionOptions *GetAppRevisionOptions) - Operation response error`, func() {
		getAppRevisionPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app/revisions/my-app-001"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppRevisionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAppRevision with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppRevisionOptions model
				getAppRevisionOptionsModel := new(codeenginev2.GetAppRevisionOptions)
				getAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				getAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAppRevision(getAppRevisionOptions *GetAppRevisionOptions)`, func() {
		getAppRevisionPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app/revisions/my-app-001"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppRevisionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"app_name": "my-app", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-app-00001", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_revision_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"actual_instances": 1, "reason": "ready"}}`)
				}))
			})
			It(`Invoke GetAppRevision successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetAppRevisionOptions model
				getAppRevisionOptionsModel := new(codeenginev2.GetAppRevisionOptions)
				getAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				getAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetAppRevisionWithContext(ctx, getAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetAppRevisionWithContext(ctx, getAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAppRevisionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"app_name": "my-app", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/apps/my-app/revisions/my-app-00001", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-app-00001", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "app_revision_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "ready", "status_details": {"actual_instances": 1, "reason": "ready"}}`)
				}))
			})
			It(`Invoke GetAppRevision successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetAppRevision(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAppRevisionOptions model
				getAppRevisionOptionsModel := new(codeenginev2.GetAppRevisionOptions)
				getAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				getAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAppRevision with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppRevisionOptions model
				getAppRevisionOptionsModel := new(codeenginev2.GetAppRevisionOptions)
				getAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				getAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAppRevisionOptions model with no property values
				getAppRevisionOptionsModelNew := new(codeenginev2.GetAppRevisionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetAppRevision(getAppRevisionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAppRevision successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetAppRevisionOptions model
				getAppRevisionOptionsModel := new(codeenginev2.GetAppRevisionOptions)
				getAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				getAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetAppRevision(getAppRevisionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAppRevision(deleteAppRevisionOptions *DeleteAppRevisionOptions)`, func() {
		deleteAppRevisionPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/apps/my-app/revisions/my-app-001"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAppRevisionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteAppRevision successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteAppRevision(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAppRevisionOptions model
				deleteAppRevisionOptionsModel := new(codeenginev2.DeleteAppRevisionOptions)
				deleteAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				deleteAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				deleteAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteAppRevision(deleteAppRevisionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAppRevision with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteAppRevisionOptions model
				deleteAppRevisionOptionsModel := new(codeenginev2.DeleteAppRevisionOptions)
				deleteAppRevisionOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				deleteAppRevisionOptionsModel.Name = core.StringPtr("my-app-001")
				deleteAppRevisionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteAppRevision(deleteAppRevisionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAppRevisionOptions model with no property values
				deleteAppRevisionOptionsModelNew := new(codeenginev2.DeleteAppRevisionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteAppRevision(deleteAppRevisionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobs(listJobsOptions *ListJobsOptions) - Operation response error`, func() {
		listJobsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListJobs with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(codeenginev2.ListJobsOptions)
				listJobsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobsOptionsModel.Start = core.StringPtr("testString")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobs(listJobsOptions *ListJobsOptions)`, func() {
		listJobsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "jobs": [{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListJobs successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(codeenginev2.ListJobsOptions)
				listJobsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobsOptionsModel.Start = core.StringPtr("testString")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListJobsWithContext(ctx, listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListJobsWithContext(ctx, listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "jobs": [{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListJobs successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(codeenginev2.ListJobsOptions)
				listJobsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobsOptionsModel.Start = core.StringPtr("testString")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListJobs with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(codeenginev2.ListJobsOptions)
				listJobsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobsOptionsModel.Start = core.StringPtr("testString")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListJobsOptions model with no property values
				listJobsOptionsModelNew := new(codeenginev2.ListJobsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListJobs(listJobsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListJobs successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(codeenginev2.ListJobsOptions)
				listJobsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobsOptionsModel.Start = core.StringPtr("testString")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.JobList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.JobList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"jobs":[{"created_at":"2022-09-13T11:41:35+02:00","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","name":"my-job","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"job_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"jobs":[{"created_at":"2022-09-13T11:41:35+02:00","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","name":"my-job","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"job_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use JobsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listJobsOptionsModel := &codeenginev2.ListJobsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewJobsPager(listJobsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.Job
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use JobsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listJobsOptionsModel := &codeenginev2.ListJobsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewJobsPager(listJobsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateJob(createJobOptions *CreateJobOptions) - Operation response error`, func() {
		createJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateJob with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.RunArguments = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
		createJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke CreateJob successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.RunArguments = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke CreateJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.RunArguments = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateJob with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.RunArguments = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateJobOptions model with no property values
				createJobOptionsModelNew := new(codeenginev2.CreateJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateJob(createJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.RunArguments = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetJob(getJobOptions *GetJobOptions) - Operation response error`, func() {
		getJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs/my-job"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetJob with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(codeenginev2.GetJobOptions)
				getJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.Name = core.StringPtr("my-job")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetJob(getJobOptions *GetJobOptions)`, func() {
		getJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs/my-job"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke GetJob successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(codeenginev2.GetJobOptions)
				getJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.Name = core.StringPtr("my-job")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetJobWithContext(ctx, getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetJobWithContext(ctx, getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke GetJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(codeenginev2.GetJobOptions)
				getJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.Name = core.StringPtr("my-job")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetJob with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(codeenginev2.GetJobOptions)
				getJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.Name = core.StringPtr("my-job")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetJobOptions model with no property values
				getJobOptionsModelNew := new(codeenginev2.GetJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetJob(getJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(codeenginev2.GetJobOptions)
				getJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.Name = core.StringPtr("my-job")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetJob(getJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
		deleteJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs/my-job"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteJobPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(codeenginev2.DeleteJobOptions)
				deleteJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.Name = core.StringPtr("my-job")
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteJob with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(codeenginev2.DeleteJobOptions)
				deleteJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.Name = core.StringPtr("my-job")
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteJobOptions model with no property values
				deleteJobOptionsModelNew := new(codeenginev2.DeleteJobOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteJob(deleteJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateJob(updateJobOptions *UpdateJobOptions) - Operation response error`, func() {
		updateJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs/my-job"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateJobPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateJob with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the JobPatch model
				jobPatchModel := new(codeenginev2.JobPatch)
				jobPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				jobPatchModel.ImageSecret = core.StringPtr("my-secret")
				jobPatchModel.RunArguments = []string{"testString"}
				jobPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				jobPatchModel.RunCommands = []string{"testString"}
				jobPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				jobPatchModel.RunMode = core.StringPtr("daemon")
				jobPatchModel.RunServiceAccount = core.StringPtr("default")
				jobPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				jobPatchModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				jobPatchModel.ScaleCpuLimit = core.StringPtr("1")
				jobPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				jobPatchModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				jobPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				jobPatchModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.IfMatch = core.StringPtr("testString")
				updateJobOptionsModel.Job = jobPatchModelAsPatch
				updateJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateJob(updateJobOptions *UpdateJobOptions)`, func() {
		updateJobPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/jobs/my-job"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateJobPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke UpdateJob successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the JobPatch model
				jobPatchModel := new(codeenginev2.JobPatch)
				jobPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				jobPatchModel.ImageSecret = core.StringPtr("my-secret")
				jobPatchModel.RunArguments = []string{"testString"}
				jobPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				jobPatchModel.RunCommands = []string{"testString"}
				jobPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				jobPatchModel.RunMode = core.StringPtr("daemon")
				jobPatchModel.RunServiceAccount = core.StringPtr("default")
				jobPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				jobPatchModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				jobPatchModel.ScaleCpuLimit = core.StringPtr("1")
				jobPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				jobPatchModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				jobPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				jobPatchModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.IfMatch = core.StringPtr("testString")
				updateJobOptionsModel.Job = jobPatchModelAsPatch
				updateJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateJobWithContext(ctx, updateJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateJobWithContext(ctx, updateJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateJobPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/jobs/my-job", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "name": "my-job", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3}`)
				}))
			})
			It(`Invoke UpdateJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the JobPatch model
				jobPatchModel := new(codeenginev2.JobPatch)
				jobPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				jobPatchModel.ImageSecret = core.StringPtr("my-secret")
				jobPatchModel.RunArguments = []string{"testString"}
				jobPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				jobPatchModel.RunCommands = []string{"testString"}
				jobPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				jobPatchModel.RunMode = core.StringPtr("daemon")
				jobPatchModel.RunServiceAccount = core.StringPtr("default")
				jobPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				jobPatchModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				jobPatchModel.ScaleCpuLimit = core.StringPtr("1")
				jobPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				jobPatchModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				jobPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				jobPatchModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.IfMatch = core.StringPtr("testString")
				updateJobOptionsModel.Job = jobPatchModelAsPatch
				updateJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateJob with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the JobPatch model
				jobPatchModel := new(codeenginev2.JobPatch)
				jobPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				jobPatchModel.ImageSecret = core.StringPtr("my-secret")
				jobPatchModel.RunArguments = []string{"testString"}
				jobPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				jobPatchModel.RunCommands = []string{"testString"}
				jobPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				jobPatchModel.RunMode = core.StringPtr("daemon")
				jobPatchModel.RunServiceAccount = core.StringPtr("default")
				jobPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				jobPatchModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				jobPatchModel.ScaleCpuLimit = core.StringPtr("1")
				jobPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				jobPatchModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				jobPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				jobPatchModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.IfMatch = core.StringPtr("testString")
				updateJobOptionsModel.Job = jobPatchModelAsPatch
				updateJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateJobOptions model with no property values
				updateJobOptionsModelNew := new(codeenginev2.UpdateJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateJob(updateJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateJob successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the JobPatch model
				jobPatchModel := new(codeenginev2.JobPatch)
				jobPatchModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				jobPatchModel.ImageSecret = core.StringPtr("my-secret")
				jobPatchModel.RunArguments = []string{"testString"}
				jobPatchModel.RunAsUser = core.Int64Ptr(int64(1001))
				jobPatchModel.RunCommands = []string{"testString"}
				jobPatchModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				jobPatchModel.RunMode = core.StringPtr("daemon")
				jobPatchModel.RunServiceAccount = core.StringPtr("default")
				jobPatchModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				jobPatchModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				jobPatchModel.ScaleCpuLimit = core.StringPtr("1")
				jobPatchModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				jobPatchModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				jobPatchModel.ScaleMemoryLimit = core.StringPtr("4G")
				jobPatchModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.IfMatch = core.StringPtr("testString")
				updateJobOptionsModel.Job = jobPatchModelAsPatch
				updateJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateJob(updateJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobRuns(listJobRunsOptions *ListJobRunsOptions) - Operation response error`, func() {
		listJobRunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobRunsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["job_name"]).To(Equal([]string{"my-job"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListJobRuns with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobRunsOptions model
				listJobRunsOptionsModel := new(codeenginev2.ListJobRunsOptions)
				listJobRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.JobName = core.StringPtr("my-job")
				listJobRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobRunsOptionsModel.Start = core.StringPtr("testString")
				listJobRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobRuns(listJobRunsOptions *ListJobRunsOptions)`, func() {
		listJobRunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["job_name"]).To(Equal([]string{"my-job"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "job_runs": [{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListJobRuns successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListJobRunsOptions model
				listJobRunsOptionsModel := new(codeenginev2.ListJobRunsOptions)
				listJobRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.JobName = core.StringPtr("my-job")
				listJobRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobRunsOptionsModel.Start = core.StringPtr("testString")
				listJobRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListJobRunsWithContext(ctx, listJobRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListJobRunsWithContext(ctx, listJobRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["job_name"]).To(Equal([]string{"my-job"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "job_runs": [{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListJobRuns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListJobRuns(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListJobRunsOptions model
				listJobRunsOptionsModel := new(codeenginev2.ListJobRunsOptions)
				listJobRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.JobName = core.StringPtr("my-job")
				listJobRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobRunsOptionsModel.Start = core.StringPtr("testString")
				listJobRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListJobRuns with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobRunsOptions model
				listJobRunsOptionsModel := new(codeenginev2.ListJobRunsOptions)
				listJobRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.JobName = core.StringPtr("my-job")
				listJobRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobRunsOptionsModel.Start = core.StringPtr("testString")
				listJobRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListJobRunsOptions model with no property values
				listJobRunsOptionsModelNew := new(codeenginev2.ListJobRunsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListJobRuns(listJobRunsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListJobRuns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListJobRunsOptions model
				listJobRunsOptionsModel := new(codeenginev2.ListJobRunsOptions)
				listJobRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.JobName = core.StringPtr("my-job")
				listJobRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listJobRunsOptionsModel.Start = core.StringPtr("testString")
				listJobRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListJobRuns(listJobRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.JobRunList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.JobRunList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobRunsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"job_runs":[{"created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","job_name":"my-job","name":"my-job-run","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"job_run_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3,"status":"completed","status_details":{"completion_time":"2022-09-22T17:40:00Z","failed":0,"pending":0,"requested":0,"running":0,"start_time":"2022-09-22T17:34:00Z","succeeded":1,"unknown":0}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"job_runs":[{"created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_reference":"icr.io/codeengine/helloworld","image_secret":"my-secret","job_name":"my-job","name":"my-job-run","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"job_run_v2","run_arguments":["RunArguments"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_variables":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","reference":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","reference":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3,"status":"completed","status_details":{"completion_time":"2022-09-22T17:40:00Z","failed":0,"pending":0,"requested":0,"running":0,"start_time":"2022-09-22T17:34:00Z","succeeded":1,"unknown":0}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use JobRunsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listJobRunsOptionsModel := &codeenginev2.ListJobRunsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					JobName: core.StringPtr("my-job"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewJobRunsPager(listJobRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.JobRun
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use JobRunsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listJobRunsOptionsModel := &codeenginev2.ListJobRunsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					JobName: core.StringPtr("my-job"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewJobRunsPager(listJobRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateJobRun(createJobRunOptions *CreateJobRunOptions) - Operation response error`, func() {
		createJobRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobRunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateJobRun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobRunOptions model
				createJobRunOptionsModel := new(codeenginev2.CreateJobRunOptions)
				createJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobRunOptionsModel.JobName = core.StringPtr("my-job")
				createJobRunOptionsModel.Name = core.StringPtr("my-job-run")
				createJobRunOptionsModel.RunArguments = []string{"testString"}
				createJobRunOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobRunOptionsModel.RunCommands = []string{"testString"}
				createJobRunOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobRunOptionsModel.RunMode = core.StringPtr("daemon")
				createJobRunOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobRunOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobRunOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobRunOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobRunOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobRunOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateJobRun(createJobRunOptions *CreateJobRunOptions)`, func() {
		createJobRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobRunPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}`)
				}))
			})
			It(`Invoke CreateJobRun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobRunOptions model
				createJobRunOptionsModel := new(codeenginev2.CreateJobRunOptions)
				createJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobRunOptionsModel.JobName = core.StringPtr("my-job")
				createJobRunOptionsModel.Name = core.StringPtr("my-job-run")
				createJobRunOptionsModel.RunArguments = []string{"testString"}
				createJobRunOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobRunOptionsModel.RunCommands = []string{"testString"}
				createJobRunOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobRunOptionsModel.RunMode = core.StringPtr("daemon")
				createJobRunOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobRunOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobRunOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobRunOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobRunOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobRunOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateJobRunWithContext(ctx, createJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateJobRunWithContext(ctx, createJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobRunPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}`)
				}))
			})
			It(`Invoke CreateJobRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateJobRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobRunOptions model
				createJobRunOptionsModel := new(codeenginev2.CreateJobRunOptions)
				createJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobRunOptionsModel.JobName = core.StringPtr("my-job")
				createJobRunOptionsModel.Name = core.StringPtr("my-job-run")
				createJobRunOptionsModel.RunArguments = []string{"testString"}
				createJobRunOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobRunOptionsModel.RunCommands = []string{"testString"}
				createJobRunOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobRunOptionsModel.RunMode = core.StringPtr("daemon")
				createJobRunOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobRunOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobRunOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobRunOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobRunOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobRunOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateJobRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobRunOptions model
				createJobRunOptionsModel := new(codeenginev2.CreateJobRunOptions)
				createJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobRunOptionsModel.JobName = core.StringPtr("my-job")
				createJobRunOptionsModel.Name = core.StringPtr("my-job-run")
				createJobRunOptionsModel.RunArguments = []string{"testString"}
				createJobRunOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobRunOptionsModel.RunCommands = []string{"testString"}
				createJobRunOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobRunOptionsModel.RunMode = core.StringPtr("daemon")
				createJobRunOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobRunOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobRunOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobRunOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobRunOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobRunOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateJobRunOptions model with no property values
				createJobRunOptionsModelNew := new(codeenginev2.CreateJobRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateJobRun(createJobRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateJobRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobRunOptions model
				createJobRunOptionsModel := new(codeenginev2.CreateJobRunOptions)
				createJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.ImageReference = core.StringPtr("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobRunOptionsModel.JobName = core.StringPtr("my-job")
				createJobRunOptionsModel.Name = core.StringPtr("my-job-run")
				createJobRunOptionsModel.RunArguments = []string{"testString"}
				createJobRunOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobRunOptionsModel.RunCommands = []string{"testString"}
				createJobRunOptionsModel.RunEnvVariables = []codeenginev2.EnvVarPrototype{*envVarPrototypeModel}
				createJobRunOptionsModel.RunMode = core.StringPtr("daemon")
				createJobRunOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobRunOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}
				createJobRunOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobRunOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobRunOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobRunOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobRunOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateJobRun(createJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetJobRun(getJobRunOptions *GetJobRunOptions) - Operation response error`, func() {
		getJobRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs/my-job"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobRunPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetJobRun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobRunOptions model
				getJobRunOptionsModel := new(codeenginev2.GetJobRunOptions)
				getJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.Name = core.StringPtr("my-job")
				getJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetJobRun(getJobRunOptions *GetJobRunOptions)`, func() {
		getJobRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs/my-job"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobRunPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}`)
				}))
			})
			It(`Invoke GetJobRun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetJobRunOptions model
				getJobRunOptionsModel := new(codeenginev2.GetJobRunOptions)
				getJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.Name = core.StringPtr("my-job")
				getJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetJobRunWithContext(ctx, getJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetJobRunWithContext(ctx, getJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobRunPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/job_runs/my-job-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_reference": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "job_name": "my-job", "name": "my-job-run", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "job_run_v2", "run_arguments": ["RunArguments"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_variables": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "reference": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "reference": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "completed", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "failed": 0, "pending": 0, "requested": 0, "running": 0, "start_time": "2022-09-22T17:34:00Z", "succeeded": 1, "unknown": 0}}`)
				}))
			})
			It(`Invoke GetJobRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetJobRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetJobRunOptions model
				getJobRunOptionsModel := new(codeenginev2.GetJobRunOptions)
				getJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.Name = core.StringPtr("my-job")
				getJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetJobRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobRunOptions model
				getJobRunOptionsModel := new(codeenginev2.GetJobRunOptions)
				getJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.Name = core.StringPtr("my-job")
				getJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetJobRunOptions model with no property values
				getJobRunOptionsModelNew := new(codeenginev2.GetJobRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetJobRun(getJobRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetJobRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetJobRunOptions model
				getJobRunOptionsModel := new(codeenginev2.GetJobRunOptions)
				getJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.Name = core.StringPtr("my-job")
				getJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetJobRun(getJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteJobRun(deleteJobRunOptions *DeleteJobRunOptions)`, func() {
		deleteJobRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/job_runs/my-job"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteJobRunPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteJobRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteJobRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteJobRunOptions model
				deleteJobRunOptionsModel := new(codeenginev2.DeleteJobRunOptions)
				deleteJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobRunOptionsModel.Name = core.StringPtr("my-job")
				deleteJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteJobRun(deleteJobRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteJobRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteJobRunOptions model
				deleteJobRunOptionsModel := new(codeenginev2.DeleteJobRunOptions)
				deleteJobRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobRunOptionsModel.Name = core.StringPtr("my-job")
				deleteJobRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteJobRun(deleteJobRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteJobRunOptions model with no property values
				deleteJobRunOptionsModelNew := new(codeenginev2.DeleteJobRunOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteJobRun(deleteJobRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuilds(listBuildsOptions *ListBuildsOptions) - Operation response error`, func() {
		listBuildsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBuilds with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildsOptions model
				listBuildsOptionsModel := new(codeenginev2.ListBuildsOptions)
				listBuildsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildsOptionsModel.Start = core.StringPtr("testString")
				listBuildsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuilds(listBuildsOptions *ListBuildsOptions)`, func() {
		listBuildsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"builds": [{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuilds successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListBuildsOptions model
				listBuildsOptionsModel := new(codeenginev2.ListBuildsOptions)
				listBuildsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildsOptionsModel.Start = core.StringPtr("testString")
				listBuildsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListBuildsWithContext(ctx, listBuildsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListBuildsWithContext(ctx, listBuildsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"builds": [{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuilds successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListBuilds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBuildsOptions model
				listBuildsOptionsModel := new(codeenginev2.ListBuildsOptions)
				listBuildsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildsOptionsModel.Start = core.StringPtr("testString")
				listBuildsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBuilds with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildsOptions model
				listBuildsOptionsModel := new(codeenginev2.ListBuildsOptions)
				listBuildsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildsOptionsModel.Start = core.StringPtr("testString")
				listBuildsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBuildsOptions model with no property values
				listBuildsOptionsModelNew := new(codeenginev2.ListBuildsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListBuilds(listBuildsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBuilds successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildsOptions model
				listBuildsOptionsModel := new(codeenginev2.ListBuildsOptions)
				listBuildsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildsOptionsModel.Start = core.StringPtr("testString")
				listBuildsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListBuilds(listBuildsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.BuildList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.BuildList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"builds":[{"created_at":"2022-09-13T11:41:35+02:00","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-build","output_image":"private.de.icr.io/icr_namespace/image-name","output_secret":"ce-auto-icr-private-eu-de","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"build_v2","source_context_dir":"some/subfolder","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"ready","status_details":{"reason":"registered"},"strategy_size":"medium","strategy_spec_file":"Dockerfile","strategy_type":"dockerfile","timeout":600}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"builds":[{"created_at":"2022-09-13T11:41:35+02:00","entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-build","output_image":"private.de.icr.io/icr_namespace/image-name","output_secret":"ce-auto-icr-private-eu-de","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"build_v2","source_context_dir":"some/subfolder","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"ready","status_details":{"reason":"registered"},"strategy_size":"medium","strategy_spec_file":"Dockerfile","strategy_type":"dockerfile","timeout":600}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BuildsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildsOptionsModel := &codeenginev2.ListBuildsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildsPager(listBuildsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.Build
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use BuildsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildsOptionsModel := &codeenginev2.ListBuildsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildsPager(listBuildsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateBuild(createBuildOptions *CreateBuildOptions) - Operation response error`, func() {
		createBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBuild with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildOptions model
				createBuildOptionsModel := new(codeenginev2.CreateBuildOptions)
				createBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBuild(createBuildOptions *CreateBuildOptions)`, func() {
		createBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke CreateBuild successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateBuildOptions model
				createBuildOptionsModel := new(codeenginev2.CreateBuildOptions)
				createBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateBuildWithContext(ctx, createBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateBuildWithContext(ctx, createBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke CreateBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateBuild(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateBuildOptions model
				createBuildOptionsModel := new(codeenginev2.CreateBuildOptions)
				createBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBuild with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildOptions model
				createBuildOptionsModel := new(codeenginev2.CreateBuildOptions)
				createBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBuildOptions model with no property values
				createBuildOptionsModelNew := new(codeenginev2.CreateBuildOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateBuild(createBuildOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildOptions model
				createBuildOptionsModel := new(codeenginev2.CreateBuildOptions)
				createBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateBuild(createBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuild(getBuildOptions *GetBuildOptions) - Operation response error`, func() {
		getBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds/my-build"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBuild with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildOptions model
				getBuildOptionsModel := new(codeenginev2.GetBuildOptions)
				getBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.Name = core.StringPtr("my-build")
				getBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuild(getBuildOptions *GetBuildOptions)`, func() {
		getBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds/my-build"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke GetBuild successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetBuildOptions model
				getBuildOptionsModel := new(codeenginev2.GetBuildOptions)
				getBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.Name = core.StringPtr("my-build")
				getBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetBuildWithContext(ctx, getBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetBuildWithContext(ctx, getBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke GetBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetBuild(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBuildOptions model
				getBuildOptionsModel := new(codeenginev2.GetBuildOptions)
				getBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.Name = core.StringPtr("my-build")
				getBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBuild with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildOptions model
				getBuildOptionsModel := new(codeenginev2.GetBuildOptions)
				getBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.Name = core.StringPtr("my-build")
				getBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBuildOptions model with no property values
				getBuildOptionsModelNew := new(codeenginev2.GetBuildOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetBuild(getBuildOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildOptions model
				getBuildOptionsModel := new(codeenginev2.GetBuildOptions)
				getBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.Name = core.StringPtr("my-build")
				getBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetBuild(getBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteBuild(deleteBuildOptions *DeleteBuildOptions)`, func() {
		deleteBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds/my-build"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBuildPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteBuild(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBuildOptions model
				deleteBuildOptionsModel := new(codeenginev2.DeleteBuildOptions)
				deleteBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.Name = core.StringPtr("my-build")
				deleteBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteBuild(deleteBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBuild with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteBuildOptions model
				deleteBuildOptionsModel := new(codeenginev2.DeleteBuildOptions)
				deleteBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.Name = core.StringPtr("my-build")
				deleteBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteBuild(deleteBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBuildOptions model with no property values
				deleteBuildOptionsModelNew := new(codeenginev2.DeleteBuildOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteBuild(deleteBuildOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBuild(updateBuildOptions *UpdateBuildOptions) - Operation response error`, func() {
		updateBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds/my-build"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBuildPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBuild with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the BuildPatch model
				buildPatchModel := new(codeenginev2.BuildPatch)
				buildPatchModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				buildPatchModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				buildPatchModel.SourceContextDir = core.StringPtr("some/subfolder")
				buildPatchModel.SourceRevision = core.StringPtr("main")
				buildPatchModel.SourceSecret = core.StringPtr("testString")
				buildPatchModel.SourceType = core.StringPtr("git")
				buildPatchModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				buildPatchModel.StrategySize = core.StringPtr("medium")
				buildPatchModel.StrategySpecFile = core.StringPtr("Dockerfile")
				buildPatchModel.StrategyType = core.StringPtr("dockerfile")
				buildPatchModel.Timeout = core.Int64Ptr(int64(600))
				buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.IfMatch = core.StringPtr("testString")
				updateBuildOptionsModel.Build = buildPatchModelAsPatch
				updateBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBuild(updateBuildOptions *UpdateBuildOptions)`, func() {
		updateBuildPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/builds/my-build"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBuildPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke UpdateBuild successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the BuildPatch model
				buildPatchModel := new(codeenginev2.BuildPatch)
				buildPatchModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				buildPatchModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				buildPatchModel.SourceContextDir = core.StringPtr("some/subfolder")
				buildPatchModel.SourceRevision = core.StringPtr("main")
				buildPatchModel.SourceSecret = core.StringPtr("testString")
				buildPatchModel.SourceType = core.StringPtr("git")
				buildPatchModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				buildPatchModel.StrategySize = core.StringPtr("medium")
				buildPatchModel.StrategySpecFile = core.StringPtr("Dockerfile")
				buildPatchModel.StrategyType = core.StringPtr("dockerfile")
				buildPatchModel.Timeout = core.Int64Ptr(int64(600))
				buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.IfMatch = core.StringPtr("testString")
				updateBuildOptionsModel.Build = buildPatchModelAsPatch
				updateBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateBuildWithContext(ctx, updateBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateBuildWithContext(ctx, updateBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBuildPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/builds/my-build", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_v2", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "ready", "status_details": {"reason": "registered"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke UpdateBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateBuild(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BuildPatch model
				buildPatchModel := new(codeenginev2.BuildPatch)
				buildPatchModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				buildPatchModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				buildPatchModel.SourceContextDir = core.StringPtr("some/subfolder")
				buildPatchModel.SourceRevision = core.StringPtr("main")
				buildPatchModel.SourceSecret = core.StringPtr("testString")
				buildPatchModel.SourceType = core.StringPtr("git")
				buildPatchModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				buildPatchModel.StrategySize = core.StringPtr("medium")
				buildPatchModel.StrategySpecFile = core.StringPtr("Dockerfile")
				buildPatchModel.StrategyType = core.StringPtr("dockerfile")
				buildPatchModel.Timeout = core.Int64Ptr(int64(600))
				buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.IfMatch = core.StringPtr("testString")
				updateBuildOptionsModel.Build = buildPatchModelAsPatch
				updateBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBuild with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the BuildPatch model
				buildPatchModel := new(codeenginev2.BuildPatch)
				buildPatchModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				buildPatchModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				buildPatchModel.SourceContextDir = core.StringPtr("some/subfolder")
				buildPatchModel.SourceRevision = core.StringPtr("main")
				buildPatchModel.SourceSecret = core.StringPtr("testString")
				buildPatchModel.SourceType = core.StringPtr("git")
				buildPatchModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				buildPatchModel.StrategySize = core.StringPtr("medium")
				buildPatchModel.StrategySpecFile = core.StringPtr("Dockerfile")
				buildPatchModel.StrategyType = core.StringPtr("dockerfile")
				buildPatchModel.Timeout = core.Int64Ptr(int64(600))
				buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.IfMatch = core.StringPtr("testString")
				updateBuildOptionsModel.Build = buildPatchModelAsPatch
				updateBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBuildOptions model with no property values
				updateBuildOptionsModelNew := new(codeenginev2.UpdateBuildOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateBuild(updateBuildOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBuild successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the BuildPatch model
				buildPatchModel := new(codeenginev2.BuildPatch)
				buildPatchModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				buildPatchModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				buildPatchModel.SourceContextDir = core.StringPtr("some/subfolder")
				buildPatchModel.SourceRevision = core.StringPtr("main")
				buildPatchModel.SourceSecret = core.StringPtr("testString")
				buildPatchModel.SourceType = core.StringPtr("git")
				buildPatchModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				buildPatchModel.StrategySize = core.StringPtr("medium")
				buildPatchModel.StrategySpecFile = core.StringPtr("Dockerfile")
				buildPatchModel.StrategyType = core.StringPtr("dockerfile")
				buildPatchModel.Timeout = core.Int64Ptr(int64(600))
				buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.IfMatch = core.StringPtr("testString")
				updateBuildOptionsModel.Build = buildPatchModelAsPatch
				updateBuildOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateBuild(updateBuildOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuildRuns(listBuildRunsOptions *ListBuildRunsOptions) - Operation response error`, func() {
		listBuildRunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildRunsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["build_name"]).To(Equal([]string{"my-build"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBuildRuns with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildRunsOptions model
				listBuildRunsOptionsModel := new(codeenginev2.ListBuildRunsOptions)
				listBuildRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.BuildName = core.StringPtr("my-build")
				listBuildRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildRunsOptionsModel.Start = core.StringPtr("testString")
				listBuildRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuildRuns(listBuildRunsOptions *ListBuildRunsOptions)`, func() {
		listBuildRunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["build_name"]).To(Equal([]string{"my-build"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_runs": [{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuildRuns successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListBuildRunsOptions model
				listBuildRunsOptionsModel := new(codeenginev2.ListBuildRunsOptions)
				listBuildRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.BuildName = core.StringPtr("my-build")
				listBuildRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildRunsOptionsModel.Start = core.StringPtr("testString")
				listBuildRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListBuildRunsWithContext(ctx, listBuildRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListBuildRunsWithContext(ctx, listBuildRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildRunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["build_name"]).To(Equal([]string{"my-build"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_runs": [{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuildRuns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListBuildRuns(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBuildRunsOptions model
				listBuildRunsOptionsModel := new(codeenginev2.ListBuildRunsOptions)
				listBuildRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.BuildName = core.StringPtr("my-build")
				listBuildRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildRunsOptionsModel.Start = core.StringPtr("testString")
				listBuildRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBuildRuns with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildRunsOptions model
				listBuildRunsOptionsModel := new(codeenginev2.ListBuildRunsOptions)
				listBuildRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.BuildName = core.StringPtr("my-build")
				listBuildRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildRunsOptionsModel.Start = core.StringPtr("testString")
				listBuildRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBuildRunsOptions model with no property values
				listBuildRunsOptionsModelNew := new(codeenginev2.ListBuildRunsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListBuildRuns(listBuildRunsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBuildRuns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildRunsOptions model
				listBuildRunsOptionsModel := new(codeenginev2.ListBuildRunsOptions)
				listBuildRunsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.BuildName = core.StringPtr("my-build")
				listBuildRunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildRunsOptionsModel.Start = core.StringPtr("testString")
				listBuildRunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListBuildRuns(listBuildRunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.BuildRunList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.BuildRunList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildRunsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"build_runs":[{"build_name":"BuildName","created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-build-run","output_image":"private.de.icr.io/icr_namespace/image-name","output_secret":"ce-auto-icr-private-eu-de","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"build_run_v2","service_account":"default","source_context_dir":"some/subfolder","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"succeeded","status_details":{"completion_time":"2022-09-22T17:40:00Z","output_digest":"sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384","reason":"succeeded","start_time":"2022-09-22T17:34:00Z"},"strategy_size":"medium","strategy_spec_file":"Dockerfile","strategy_type":"dockerfile","timeout":600}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"build_runs":[{"build_name":"BuildName","created_at":"2022-09-13T11:41:35+02:00","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-build-run","output_image":"private.de.icr.io/icr_namespace/image-name","output_secret":"ce-auto-icr-private-eu-de","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"build_run_v2","service_account":"default","source_context_dir":"some/subfolder","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"succeeded","status_details":{"completion_time":"2022-09-22T17:40:00Z","output_digest":"sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384","reason":"succeeded","start_time":"2022-09-22T17:34:00Z"},"strategy_size":"medium","strategy_spec_file":"Dockerfile","strategy_type":"dockerfile","timeout":600}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BuildRunsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildRunsOptionsModel := &codeenginev2.ListBuildRunsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					BuildName: core.StringPtr("my-build"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildRunsPager(listBuildRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.BuildRun
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use BuildRunsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildRunsOptionsModel := &codeenginev2.ListBuildRunsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					BuildName: core.StringPtr("my-build"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildRunsPager(listBuildRunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateBuildRun(createBuildRunOptions *CreateBuildRunOptions) - Operation response error`, func() {
		createBuildRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildRunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBuildRun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildRunOptions model
				createBuildRunOptionsModel := new(codeenginev2.CreateBuildRunOptions)
				createBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.BuildName = core.StringPtr("testString")
				createBuildRunOptionsModel.Name = core.StringPtr("testString")
				createBuildRunOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.ServiceAccount = core.StringPtr("default")
				createBuildRunOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildRunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildRunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildRunOptionsModel.SourceType = core.StringPtr("git")
				createBuildRunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildRunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildRunOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildRunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBuildRun(createBuildRunOptions *CreateBuildRunOptions)`, func() {
		createBuildRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildRunPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke CreateBuildRun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateBuildRunOptions model
				createBuildRunOptionsModel := new(codeenginev2.CreateBuildRunOptions)
				createBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.BuildName = core.StringPtr("testString")
				createBuildRunOptionsModel.Name = core.StringPtr("testString")
				createBuildRunOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.ServiceAccount = core.StringPtr("default")
				createBuildRunOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildRunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildRunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildRunOptionsModel.SourceType = core.StringPtr("git")
				createBuildRunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildRunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildRunOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildRunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateBuildRunWithContext(ctx, createBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateBuildRunWithContext(ctx, createBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildRunPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke CreateBuildRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateBuildRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateBuildRunOptions model
				createBuildRunOptionsModel := new(codeenginev2.CreateBuildRunOptions)
				createBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.BuildName = core.StringPtr("testString")
				createBuildRunOptionsModel.Name = core.StringPtr("testString")
				createBuildRunOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.ServiceAccount = core.StringPtr("default")
				createBuildRunOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildRunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildRunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildRunOptionsModel.SourceType = core.StringPtr("git")
				createBuildRunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildRunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildRunOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildRunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBuildRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildRunOptions model
				createBuildRunOptionsModel := new(codeenginev2.CreateBuildRunOptions)
				createBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.BuildName = core.StringPtr("testString")
				createBuildRunOptionsModel.Name = core.StringPtr("testString")
				createBuildRunOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.ServiceAccount = core.StringPtr("default")
				createBuildRunOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildRunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildRunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildRunOptionsModel.SourceType = core.StringPtr("git")
				createBuildRunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildRunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildRunOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildRunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBuildRunOptions model with no property values
				createBuildRunOptionsModelNew := new(codeenginev2.CreateBuildRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateBuildRun(createBuildRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateBuildRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildRunOptions model
				createBuildRunOptionsModel := new(codeenginev2.CreateBuildRunOptions)
				createBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.BuildName = core.StringPtr("testString")
				createBuildRunOptionsModel.Name = core.StringPtr("testString")
				createBuildRunOptionsModel.OutputImage = core.StringPtr("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.OutputSecret = core.StringPtr("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.ServiceAccount = core.StringPtr("default")
				createBuildRunOptionsModel.SourceContextDir = core.StringPtr("some/subfolder")
				createBuildRunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildRunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildRunOptionsModel.SourceType = core.StringPtr("git")
				createBuildRunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildRunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildRunOptionsModel.StrategyType = core.StringPtr("dockerfile")
				createBuildRunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateBuildRun(createBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuildRun(getBuildRunOptions *GetBuildRunOptions) - Operation response error`, func() {
		getBuildRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildRunPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBuildRun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildRunOptions model
				getBuildRunOptionsModel := new(codeenginev2.GetBuildRunOptions)
				getBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				getBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuildRun(getBuildRunOptions *GetBuildRunOptions)`, func() {
		getBuildRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildRunPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke GetBuildRun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetBuildRunOptions model
				getBuildRunOptionsModel := new(codeenginev2.GetBuildRunOptions)
				getBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				getBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetBuildRunWithContext(ctx, getBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetBuildRunWithContext(ctx, getBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildRunPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_name": "BuildName", "created_at": "2022-09-13T11:41:35+02:00", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/build_runs/my-build-run", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-build-run", "output_image": "private.de.icr.io/icr_namespace/image-name", "output_secret": "ce-auto-icr-private-eu-de", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "build_run_v2", "service_account": "default", "source_context_dir": "some/subfolder", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "succeeded", "status_details": {"completion_time": "2022-09-22T17:40:00Z", "output_digest": "sha256:9a3d845c629d2b4a6b271b1d526dfafc1e7d9511f8863b43b5bb0483ef626384", "reason": "succeeded", "start_time": "2022-09-22T17:34:00Z"}, "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "strategy_type": "dockerfile", "timeout": 600}`)
				}))
			})
			It(`Invoke GetBuildRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetBuildRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBuildRunOptions model
				getBuildRunOptionsModel := new(codeenginev2.GetBuildRunOptions)
				getBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				getBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBuildRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildRunOptions model
				getBuildRunOptionsModel := new(codeenginev2.GetBuildRunOptions)
				getBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				getBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBuildRunOptions model with no property values
				getBuildRunOptionsModelNew := new(codeenginev2.GetBuildRunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetBuildRun(getBuildRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBuildRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildRunOptions model
				getBuildRunOptionsModel := new(codeenginev2.GetBuildRunOptions)
				getBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				getBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetBuildRun(getBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteBuildRun(deleteBuildRunOptions *DeleteBuildRunOptions)`, func() {
		deleteBuildRunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBuildRunPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteBuildRun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteBuildRun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBuildRunOptions model
				deleteBuildRunOptionsModel := new(codeenginev2.DeleteBuildRunOptions)
				deleteBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				deleteBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteBuildRun(deleteBuildRunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBuildRun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteBuildRunOptions model
				deleteBuildRunOptionsModel := new(codeenginev2.DeleteBuildRunOptions)
				deleteBuildRunOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildRunOptionsModel.Name = core.StringPtr("my-build-run")
				deleteBuildRunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteBuildRun(deleteBuildRunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBuildRunOptions model with no property values
				deleteBuildRunOptionsModelNew := new(codeenginev2.DeleteBuildRunOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteBuildRun(deleteBuildRunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigMaps(listConfigMapsOptions *ListConfigMapsOptions) - Operation response error`, func() {
		listConfigMapsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigMapsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigMaps with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigMapsOptions model
				listConfigMapsOptionsModel := new(codeenginev2.ListConfigMapsOptions)
				listConfigMapsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigMapsOptionsModel.Start = core.StringPtr("testString")
				listConfigMapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigMaps(listConfigMapsOptions *ListConfigMapsOptions)`, func() {
		listConfigMapsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigMapsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"config_maps": [{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigMaps successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigMapsOptions model
				listConfigMapsOptionsModel := new(codeenginev2.ListConfigMapsOptions)
				listConfigMapsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigMapsOptionsModel.Start = core.StringPtr("testString")
				listConfigMapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListConfigMapsWithContext(ctx, listConfigMapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListConfigMapsWithContext(ctx, listConfigMapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigMapsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"config_maps": [{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}], "first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigMaps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListConfigMaps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigMapsOptions model
				listConfigMapsOptionsModel := new(codeenginev2.ListConfigMapsOptions)
				listConfigMapsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigMapsOptionsModel.Start = core.StringPtr("testString")
				listConfigMapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigMaps with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigMapsOptions model
				listConfigMapsOptionsModel := new(codeenginev2.ListConfigMapsOptions)
				listConfigMapsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigMapsOptionsModel.Start = core.StringPtr("testString")
				listConfigMapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigMapsOptions model with no property values
				listConfigMapsOptionsModelNew := new(codeenginev2.ListConfigMapsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListConfigMaps(listConfigMapsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListConfigMaps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigMapsOptions model
				listConfigMapsOptionsModel := new(codeenginev2.ListConfigMapsOptions)
				listConfigMapsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigMapsOptionsModel.Start = core.StringPtr("testString")
				listConfigMapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListConfigMaps(listConfigMapsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.ConfigMapList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.ConfigMapList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigMapsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"config_maps":[{"created_at":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-config-map","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"config_map_v2"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"config_maps":[{"created_at":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"entity_tag":"2385407409","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-config-map","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"config_map_v2"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ConfigMapsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listConfigMapsOptionsModel := &codeenginev2.ListConfigMapsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewConfigMapsPager(listConfigMapsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.ConfigMap
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ConfigMapsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listConfigMapsOptionsModel := &codeenginev2.ListConfigMapsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewConfigMapsPager(listConfigMapsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateConfigMap(createConfigMapOptions *CreateConfigMapOptions) - Operation response error`, func() {
		createConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigMapPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfigMap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigMapOptions model
				createConfigMapOptionsModel := new(codeenginev2.CreateConfigMapOptions)
				createConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigMapOptionsModel.Data = make(map[string]string)
				createConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfigMap(createConfigMapOptions *CreateConfigMapOptions)`, func() {
		createConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigMapPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke CreateConfigMap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateConfigMapOptions model
				createConfigMapOptionsModel := new(codeenginev2.CreateConfigMapOptions)
				createConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigMapOptionsModel.Data = make(map[string]string)
				createConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateConfigMapWithContext(ctx, createConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateConfigMapWithContext(ctx, createConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigMapPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke CreateConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateConfigMap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateConfigMapOptions model
				createConfigMapOptionsModel := new(codeenginev2.CreateConfigMapOptions)
				createConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigMapOptionsModel.Data = make(map[string]string)
				createConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfigMap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigMapOptions model
				createConfigMapOptionsModel := new(codeenginev2.CreateConfigMapOptions)
				createConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigMapOptionsModel.Data = make(map[string]string)
				createConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigMapOptions model with no property values
				createConfigMapOptionsModelNew := new(codeenginev2.CreateConfigMapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateConfigMap(createConfigMapOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigMapOptions model
				createConfigMapOptionsModel := new(codeenginev2.CreateConfigMapOptions)
				createConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigMapOptionsModel.Data = make(map[string]string)
				createConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateConfigMap(createConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigMap(getConfigMapOptions *GetConfigMapOptions) - Operation response error`, func() {
		getConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigMapPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfigMap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigMapOptions model
				getConfigMapOptionsModel := new(codeenginev2.GetConfigMapOptions)
				getConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				getConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigMap(getConfigMapOptions *GetConfigMapOptions)`, func() {
		getConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigMapPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke GetConfigMap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigMapOptions model
				getConfigMapOptionsModel := new(codeenginev2.GetConfigMapOptions)
				getConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				getConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetConfigMapWithContext(ctx, getConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetConfigMapWithContext(ctx, getConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigMapPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke GetConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetConfigMap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigMapOptions model
				getConfigMapOptionsModel := new(codeenginev2.GetConfigMapOptions)
				getConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				getConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigMap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigMapOptions model
				getConfigMapOptionsModel := new(codeenginev2.GetConfigMapOptions)
				getConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				getConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigMapOptions model with no property values
				getConfigMapOptionsModelNew := new(codeenginev2.GetConfigMapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetConfigMap(getConfigMapOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigMapOptions model
				getConfigMapOptionsModel := new(codeenginev2.GetConfigMapOptions)
				getConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				getConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetConfigMap(getConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceConfigMap(replaceConfigMapOptions *ReplaceConfigMapOptions) - Operation response error`, func() {
		replaceConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceConfigMapPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceConfigMap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceConfigMapOptions model
				replaceConfigMapOptionsModel := new(codeenginev2.ReplaceConfigMapOptions)
				replaceConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				replaceConfigMapOptionsModel.IfMatch = core.StringPtr("testString")
				replaceConfigMapOptionsModel.Data = make(map[string]string)
				replaceConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceConfigMap(replaceConfigMapOptions *ReplaceConfigMapOptions)`, func() {
		replaceConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceConfigMapPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke ReplaceConfigMap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceConfigMapOptions model
				replaceConfigMapOptionsModel := new(codeenginev2.ReplaceConfigMapOptions)
				replaceConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				replaceConfigMapOptionsModel.IfMatch = core.StringPtr("testString")
				replaceConfigMapOptionsModel.Data = make(map[string]string)
				replaceConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ReplaceConfigMapWithContext(ctx, replaceConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ReplaceConfigMapWithContext(ctx, replaceConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceConfigMapPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/config_maps/my-config-map", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-config-map", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "config_map_v2"}`)
				}))
			})
			It(`Invoke ReplaceConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ReplaceConfigMap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceConfigMapOptions model
				replaceConfigMapOptionsModel := new(codeenginev2.ReplaceConfigMapOptions)
				replaceConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				replaceConfigMapOptionsModel.IfMatch = core.StringPtr("testString")
				replaceConfigMapOptionsModel.Data = make(map[string]string)
				replaceConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceConfigMap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceConfigMapOptions model
				replaceConfigMapOptionsModel := new(codeenginev2.ReplaceConfigMapOptions)
				replaceConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				replaceConfigMapOptionsModel.IfMatch = core.StringPtr("testString")
				replaceConfigMapOptionsModel.Data = make(map[string]string)
				replaceConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceConfigMapOptions model with no property values
				replaceConfigMapOptionsModelNew := new(codeenginev2.ReplaceConfigMapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceConfigMapOptions model
				replaceConfigMapOptionsModel := new(codeenginev2.ReplaceConfigMapOptions)
				replaceConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				replaceConfigMapOptionsModel.IfMatch = core.StringPtr("testString")
				replaceConfigMapOptionsModel.Data = make(map[string]string)
				replaceConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ReplaceConfigMap(replaceConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfigMap(deleteConfigMapOptions *DeleteConfigMapOptions)`, func() {
		deleteConfigMapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigMapPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteConfigMap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteConfigMap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteConfigMapOptions model
				deleteConfigMapOptionsModel := new(codeenginev2.DeleteConfigMapOptions)
				deleteConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				deleteConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteConfigMap(deleteConfigMapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteConfigMap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigMapOptions model
				deleteConfigMapOptionsModel := new(codeenginev2.DeleteConfigMapOptions)
				deleteConfigMapOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigMapOptionsModel.Name = core.StringPtr("my-config-map")
				deleteConfigMapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteConfigMap(deleteConfigMapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteConfigMapOptions model with no property values
				deleteConfigMapOptionsModelNew := new(codeenginev2.DeleteConfigMapOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteConfigMap(deleteConfigMapOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSecrets(listSecretsOptions *ListSecretsOptions) - Operation response error`, func() {
		listSecretsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSecretsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSecrets with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListSecretsOptions model
				listSecretsOptionsModel := new(codeenginev2.ListSecretsOptions)
				listSecretsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listSecretsOptionsModel.Start = core.StringPtr("testString")
				listSecretsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSecrets(listSecretsOptions *ListSecretsOptions)`, func() {
		listSecretsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSecretsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "secrets": [{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}]}`)
				}))
			})
			It(`Invoke ListSecrets successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListSecretsOptions model
				listSecretsOptionsModel := new(codeenginev2.ListSecretsOptions)
				listSecretsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listSecretsOptionsModel.Start = core.StringPtr("testString")
				listSecretsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListSecretsWithContext(ctx, listSecretsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListSecretsWithContext(ctx, listSecretsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSecretsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "Href"}, "limit": 100, "next": {"href": "Href", "start": "Start"}, "secrets": [{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}]}`)
				}))
			})
			It(`Invoke ListSecrets successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListSecrets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSecretsOptions model
				listSecretsOptionsModel := new(codeenginev2.ListSecretsOptions)
				listSecretsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listSecretsOptionsModel.Start = core.StringPtr("testString")
				listSecretsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSecrets with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListSecretsOptions model
				listSecretsOptionsModel := new(codeenginev2.ListSecretsOptions)
				listSecretsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listSecretsOptionsModel.Start = core.StringPtr("testString")
				listSecretsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSecretsOptions model with no property values
				listSecretsOptionsModelNew := new(codeenginev2.ListSecretsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListSecrets(listSecretsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSecrets successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListSecretsOptions model
				listSecretsOptionsModel := new(codeenginev2.ListSecretsOptions)
				listSecretsOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listSecretsOptionsModel.Start = core.StringPtr("testString")
				listSecretsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListSecrets(listSecretsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(codeenginev2.SecretList)
				nextObject := new(codeenginev2.ListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.SecretList)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSecretsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"secrets":[{"created_at":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"entity_tag":"2385407409","format":"generic","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-secret","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"ResourceType"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"secrets":[{"created_at":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"entity_tag":"2385407409","format":"generic","href":"https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","name":"my-secret","project_id":"4e49b3e0-27a8-48d2-a784-c7ee48bb863b","resource_type":"ResourceType"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SecretsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listSecretsOptionsModel := &codeenginev2.ListSecretsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewSecretsPager(listSecretsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.Secret
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SecretsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listSecretsOptionsModel := &codeenginev2.ListSecretsOptions{
					ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewSecretsPager(listSecretsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateSecret(createSecretOptions *CreateSecretOptions) - Operation response error`, func() {
		createSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSecretPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSecret with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateSecretOptions model
				createSecretOptionsModel := new(codeenginev2.CreateSecretOptions)
				createSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Format = core.StringPtr("generic")
				createSecretOptionsModel.Name = core.StringPtr("my-secret")
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSecret(createSecretOptions *CreateSecretOptions)`, func() {
		createSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSecretPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke CreateSecret successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateSecretOptions model
				createSecretOptionsModel := new(codeenginev2.CreateSecretOptions)
				createSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Format = core.StringPtr("generic")
				createSecretOptionsModel.Name = core.StringPtr("my-secret")
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateSecretWithContext(ctx, createSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateSecretWithContext(ctx, createSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSecretPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke CreateSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSecretOptions model
				createSecretOptionsModel := new(codeenginev2.CreateSecretOptions)
				createSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Format = core.StringPtr("generic")
				createSecretOptionsModel.Name = core.StringPtr("my-secret")
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSecret with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateSecretOptions model
				createSecretOptionsModel := new(codeenginev2.CreateSecretOptions)
				createSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Format = core.StringPtr("generic")
				createSecretOptionsModel.Name = core.StringPtr("my-secret")
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSecretOptions model with no property values
				createSecretOptionsModelNew := new(codeenginev2.CreateSecretOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateSecret(createSecretOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateSecretOptions model
				createSecretOptionsModel := new(codeenginev2.CreateSecretOptions)
				createSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Format = core.StringPtr("generic")
				createSecretOptionsModel.Name = core.StringPtr("my-secret")
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateSecret(createSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSecret(getSecretOptions *GetSecretOptions) - Operation response error`, func() {
		getSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecretPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSecret with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetSecretOptions model
				getSecretOptionsModel := new(codeenginev2.GetSecretOptions)
				getSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.Name = core.StringPtr("my-secret")
				getSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSecret(getSecretOptions *GetSecretOptions)`, func() {
		getSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecretPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke GetSecret successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetSecretOptions model
				getSecretOptionsModel := new(codeenginev2.GetSecretOptions)
				getSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.Name = core.StringPtr("my-secret")
				getSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetSecretWithContext(ctx, getSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetSecretWithContext(ctx, getSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecretPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke GetSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecretOptions model
				getSecretOptionsModel := new(codeenginev2.GetSecretOptions)
				getSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.Name = core.StringPtr("my-secret")
				getSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSecret with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetSecretOptions model
				getSecretOptionsModel := new(codeenginev2.GetSecretOptions)
				getSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.Name = core.StringPtr("my-secret")
				getSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSecretOptions model with no property values
				getSecretOptionsModelNew := new(codeenginev2.GetSecretOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetSecret(getSecretOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetSecretOptions model
				getSecretOptionsModel := new(codeenginev2.GetSecretOptions)
				getSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.Name = core.StringPtr("my-secret")
				getSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetSecret(getSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSecret(replaceSecretOptions *ReplaceSecretOptions) - Operation response error`, func() {
		replaceSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSecretPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceSecret with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceSecretOptions model
				replaceSecretOptionsModel := new(codeenginev2.ReplaceSecretOptions)
				replaceSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.Name = core.StringPtr("my-secret")
				replaceSecretOptionsModel.IfMatch = core.StringPtr("testString")
				replaceSecretOptionsModel.Data = make(map[string]string)
				replaceSecretOptionsModel.Format = core.StringPtr("generic")
				replaceSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSecret(replaceSecretOptions *ReplaceSecretOptions)`, func() {
		replaceSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSecretPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke ReplaceSecret successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceSecretOptions model
				replaceSecretOptionsModel := new(codeenginev2.ReplaceSecretOptions)
				replaceSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.Name = core.StringPtr("my-secret")
				replaceSecretOptionsModel.IfMatch = core.StringPtr("testString")
				replaceSecretOptionsModel.Data = make(map[string]string)
				replaceSecretOptionsModel.Format = core.StringPtr("generic")
				replaceSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ReplaceSecretWithContext(ctx, replaceSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ReplaceSecretWithContext(ctx, replaceSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSecretPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "entity_tag": "2385407409", "format": "generic", "href": "https://api.eu-de.codeengine.cloud.ibm.com/v2/projects/4e49b3e0-27a8-48d2-a784-c7ee48bb863b/secrets/my-secret", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "name": "my-secret", "project_id": "4e49b3e0-27a8-48d2-a784-c7ee48bb863b", "resource_type": "ResourceType"}`)
				}))
			})
			It(`Invoke ReplaceSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ReplaceSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceSecretOptions model
				replaceSecretOptionsModel := new(codeenginev2.ReplaceSecretOptions)
				replaceSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.Name = core.StringPtr("my-secret")
				replaceSecretOptionsModel.IfMatch = core.StringPtr("testString")
				replaceSecretOptionsModel.Data = make(map[string]string)
				replaceSecretOptionsModel.Format = core.StringPtr("generic")
				replaceSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceSecret with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceSecretOptions model
				replaceSecretOptionsModel := new(codeenginev2.ReplaceSecretOptions)
				replaceSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.Name = core.StringPtr("my-secret")
				replaceSecretOptionsModel.IfMatch = core.StringPtr("testString")
				replaceSecretOptionsModel.Data = make(map[string]string)
				replaceSecretOptionsModel.Format = core.StringPtr("generic")
				replaceSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceSecretOptions model with no property values
				replaceSecretOptionsModelNew := new(codeenginev2.ReplaceSecretOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ReplaceSecret(replaceSecretOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReplaceSecretOptions model
				replaceSecretOptionsModel := new(codeenginev2.ReplaceSecretOptions)
				replaceSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.Name = core.StringPtr("my-secret")
				replaceSecretOptionsModel.IfMatch = core.StringPtr("testString")
				replaceSecretOptionsModel.Data = make(map[string]string)
				replaceSecretOptionsModel.Format = core.StringPtr("generic")
				replaceSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ReplaceSecret(replaceSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSecret(deleteSecretOptions *DeleteSecretOptions)`, func() {
		deleteSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSecretPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSecretOptions model
				deleteSecretOptionsModel := new(codeenginev2.DeleteSecretOptions)
				deleteSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.Name = core.StringPtr("my-secret")
				deleteSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteSecret(deleteSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSecret with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteSecretOptions model
				deleteSecretOptionsModel := new(codeenginev2.DeleteSecretOptions)
				deleteSecretOptionsModel.ProjectID = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.Name = core.StringPtr("my-secret")
				deleteSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteSecret(deleteSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSecretOptions model with no property values
				deleteSecretOptionsModelNew := new(codeenginev2.DeleteSecretOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteSecret(deleteSecretOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			codeEngineService, _ := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
				URL:           "http://codeenginev2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateAppOptions successfully`, func() {
				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				Expect(envVarPrototypeModel).ToNot(BeNil())
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")
				Expect(envVarPrototypeModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarPrototypeModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarPrototypeModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarPrototypeModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarPrototypeModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				Expect(volumeMountPrototypeModel).ToNot(BeNil())
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")
				Expect(volumeMountPrototypeModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountPrototypeModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountPrototypeModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the CreateAppOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createAppOptionsImageReference := "icr.io/codeengine/helloworld"
				createAppOptionsName := "my-app"
				createAppOptionsModel := codeEngineService.NewCreateAppOptions(projectID, createAppOptionsImageReference, createAppOptionsName)
				createAppOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.SetImageReference("icr.io/codeengine/helloworld")
				createAppOptionsModel.SetName("my-app")
				createAppOptionsModel.SetImagePort(int64(8080))
				createAppOptionsModel.SetImageSecret("my-secret")
				createAppOptionsModel.SetManagedDomainMappings("local_public")
				createAppOptionsModel.SetRunArguments([]string{"testString"})
				createAppOptionsModel.SetRunAsUser(int64(1001))
				createAppOptionsModel.SetRunCommands([]string{"testString"})
				createAppOptionsModel.SetRunEnvVariables([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel})
				createAppOptionsModel.SetRunServiceAccount("default")
				createAppOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel})
				createAppOptionsModel.SetScaleConcurrency(int64(100))
				createAppOptionsModel.SetScaleConcurrencyTarget(int64(80))
				createAppOptionsModel.SetScaleCpuLimit("1")
				createAppOptionsModel.SetScaleEphemeralStorageLimit("4G")
				createAppOptionsModel.SetScaleInitialInstances(int64(1))
				createAppOptionsModel.SetScaleMaxInstances(int64(10))
				createAppOptionsModel.SetScaleMemoryLimit("4G")
				createAppOptionsModel.SetScaleMinInstances(int64(1))
				createAppOptionsModel.SetScaleRequestTimeout(int64(300))
				createAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAppOptionsModel).ToNot(BeNil())
				Expect(createAppOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createAppOptionsModel.ImageReference).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(createAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(createAppOptionsModel.ImagePort).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(createAppOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(createAppOptionsModel.ManagedDomainMappings).To(Equal(core.StringPtr("local_public")))
				Expect(createAppOptionsModel.RunArguments).To(Equal([]string{"testString"}))
				Expect(createAppOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(createAppOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(createAppOptionsModel.RunEnvVariables).To(Equal([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel}))
				Expect(createAppOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createAppOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}))
				Expect(createAppOptionsModel.ScaleConcurrency).To(Equal(core.Int64Ptr(int64(100))))
				Expect(createAppOptionsModel.ScaleConcurrencyTarget).To(Equal(core.Int64Ptr(int64(80))))
				Expect(createAppOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(createAppOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(createAppOptionsModel.ScaleInitialInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createAppOptionsModel.ScaleMaxInstances).To(Equal(core.Int64Ptr(int64(10))))
				Expect(createAppOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(createAppOptionsModel.ScaleMinInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createAppOptionsModel.ScaleRequestTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(createAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBuildOptions successfully`, func() {
				// Construct an instance of the CreateBuildOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createBuildOptionsName := "my-build"
				createBuildOptionsOutputImage := "private.de.icr.io/icr_namespace/image-name"
				createBuildOptionsOutputSecret := "ce-auto-icr-private-eu-de"
				createBuildOptionsSourceURL := "https://github.com/IBM/CodeEngine"
				createBuildOptionsStrategyType := "dockerfile"
				createBuildOptionsModel := codeEngineService.NewCreateBuildOptions(projectID, createBuildOptionsName, createBuildOptionsOutputImage, createBuildOptionsOutputSecret, createBuildOptionsSourceURL, createBuildOptionsStrategyType)
				createBuildOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.SetName("my-build")
				createBuildOptionsModel.SetOutputImage("private.de.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.SetOutputSecret("ce-auto-icr-private-eu-de")
				createBuildOptionsModel.SetSourceURL("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.SetStrategyType("dockerfile")
				createBuildOptionsModel.SetSourceContextDir("some/subfolder")
				createBuildOptionsModel.SetSourceRevision("main")
				createBuildOptionsModel.SetSourceSecret("testString")
				createBuildOptionsModel.SetSourceType("git")
				createBuildOptionsModel.SetStrategySize("medium")
				createBuildOptionsModel.SetStrategySpecFile("Dockerfile")
				createBuildOptionsModel.SetTimeout(int64(600))
				createBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBuildOptionsModel).ToNot(BeNil())
				Expect(createBuildOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(createBuildOptionsModel.OutputImage).To(Equal(core.StringPtr("private.de.icr.io/icr_namespace/image-name")))
				Expect(createBuildOptionsModel.OutputSecret).To(Equal(core.StringPtr("ce-auto-icr-private-eu-de")))
				Expect(createBuildOptionsModel.SourceURL).To(Equal(core.StringPtr("https://github.com/IBM/CodeEngine")))
				Expect(createBuildOptionsModel.StrategyType).To(Equal(core.StringPtr("dockerfile")))
				Expect(createBuildOptionsModel.SourceContextDir).To(Equal(core.StringPtr("some/subfolder")))
				Expect(createBuildOptionsModel.SourceRevision).To(Equal(core.StringPtr("main")))
				Expect(createBuildOptionsModel.SourceSecret).To(Equal(core.StringPtr("testString")))
				Expect(createBuildOptionsModel.SourceType).To(Equal(core.StringPtr("git")))
				Expect(createBuildOptionsModel.StrategySize).To(Equal(core.StringPtr("medium")))
				Expect(createBuildOptionsModel.StrategySpecFile).To(Equal(core.StringPtr("Dockerfile")))
				Expect(createBuildOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(600))))
				Expect(createBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBuildRunOptions successfully`, func() {
				// Construct an instance of the CreateBuildRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createBuildRunOptionsModel := codeEngineService.NewCreateBuildRunOptions(projectID)
				createBuildRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildRunOptionsModel.SetBuildName("testString")
				createBuildRunOptionsModel.SetName("testString")
				createBuildRunOptionsModel.SetOutputImage("private.de.icr.io/icr_namespace/image-name")
				createBuildRunOptionsModel.SetOutputSecret("ce-auto-icr-private-eu-de")
				createBuildRunOptionsModel.SetServiceAccount("default")
				createBuildRunOptionsModel.SetSourceContextDir("some/subfolder")
				createBuildRunOptionsModel.SetSourceRevision("main")
				createBuildRunOptionsModel.SetSourceSecret("testString")
				createBuildRunOptionsModel.SetSourceType("git")
				createBuildRunOptionsModel.SetSourceURL("https://github.com/IBM/CodeEngine")
				createBuildRunOptionsModel.SetStrategySize("medium")
				createBuildRunOptionsModel.SetStrategySpecFile("Dockerfile")
				createBuildRunOptionsModel.SetStrategyType("dockerfile")
				createBuildRunOptionsModel.SetTimeout(int64(600))
				createBuildRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBuildRunOptionsModel).ToNot(BeNil())
				Expect(createBuildRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createBuildRunOptionsModel.BuildName).To(Equal(core.StringPtr("testString")))
				Expect(createBuildRunOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createBuildRunOptionsModel.OutputImage).To(Equal(core.StringPtr("private.de.icr.io/icr_namespace/image-name")))
				Expect(createBuildRunOptionsModel.OutputSecret).To(Equal(core.StringPtr("ce-auto-icr-private-eu-de")))
				Expect(createBuildRunOptionsModel.ServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createBuildRunOptionsModel.SourceContextDir).To(Equal(core.StringPtr("some/subfolder")))
				Expect(createBuildRunOptionsModel.SourceRevision).To(Equal(core.StringPtr("main")))
				Expect(createBuildRunOptionsModel.SourceSecret).To(Equal(core.StringPtr("testString")))
				Expect(createBuildRunOptionsModel.SourceType).To(Equal(core.StringPtr("git")))
				Expect(createBuildRunOptionsModel.SourceURL).To(Equal(core.StringPtr("https://github.com/IBM/CodeEngine")))
				Expect(createBuildRunOptionsModel.StrategySize).To(Equal(core.StringPtr("medium")))
				Expect(createBuildRunOptionsModel.StrategySpecFile).To(Equal(core.StringPtr("Dockerfile")))
				Expect(createBuildRunOptionsModel.StrategyType).To(Equal(core.StringPtr("dockerfile")))
				Expect(createBuildRunOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(600))))
				Expect(createBuildRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateConfigMapOptions successfully`, func() {
				// Construct an instance of the CreateConfigMapOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createConfigMapOptionsName := "my-configmap"
				createConfigMapOptionsModel := codeEngineService.NewCreateConfigMapOptions(projectID, createConfigMapOptionsName)
				createConfigMapOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigMapOptionsModel.SetName("my-configmap")
				createConfigMapOptionsModel.SetData(make(map[string]string))
				createConfigMapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigMapOptionsModel).ToNot(BeNil())
				Expect(createConfigMapOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createConfigMapOptionsModel.Name).To(Equal(core.StringPtr("my-configmap")))
				Expect(createConfigMapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createConfigMapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateJobOptions successfully`, func() {
				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				Expect(envVarPrototypeModel).ToNot(BeNil())
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")
				Expect(envVarPrototypeModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarPrototypeModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarPrototypeModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarPrototypeModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarPrototypeModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				Expect(volumeMountPrototypeModel).ToNot(BeNil())
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")
				Expect(volumeMountPrototypeModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountPrototypeModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountPrototypeModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the CreateJobOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createJobOptionsImageReference := "icr.io/codeengine/helloworld"
				createJobOptionsName := "my-job"
				createJobOptionsModel := codeEngineService.NewCreateJobOptions(projectID, createJobOptionsImageReference, createJobOptionsName)
				createJobOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.SetImageReference("icr.io/codeengine/helloworld")
				createJobOptionsModel.SetName("my-job")
				createJobOptionsModel.SetImageSecret("my-secret")
				createJobOptionsModel.SetRunArguments([]string{"testString"})
				createJobOptionsModel.SetRunAsUser(int64(1001))
				createJobOptionsModel.SetRunCommands([]string{"testString"})
				createJobOptionsModel.SetRunEnvVariables([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel})
				createJobOptionsModel.SetRunMode("daemon")
				createJobOptionsModel.SetRunServiceAccount("default")
				createJobOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel})
				createJobOptionsModel.SetScaleArraySpec("1-5,7-8,10")
				createJobOptionsModel.SetScaleCpuLimit("1")
				createJobOptionsModel.SetScaleEphemeralStorageLimit("4G")
				createJobOptionsModel.SetScaleMaxExecutionTime(int64(7200))
				createJobOptionsModel.SetScaleMemoryLimit("4G")
				createJobOptionsModel.SetScaleRetryLimit(int64(3))
				createJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createJobOptionsModel).ToNot(BeNil())
				Expect(createJobOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createJobOptionsModel.ImageReference).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(createJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(createJobOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(createJobOptionsModel.RunArguments).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(createJobOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.RunEnvVariables).To(Equal([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel}))
				Expect(createJobOptionsModel.RunMode).To(Equal(core.StringPtr("daemon")))
				Expect(createJobOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createJobOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}))
				Expect(createJobOptionsModel.ScaleArraySpec).To(Equal(core.StringPtr("1-5,7-8,10")))
				Expect(createJobOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(createJobOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobOptionsModel.ScaleMaxExecutionTime).To(Equal(core.Int64Ptr(int64(7200))))
				Expect(createJobOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobOptionsModel.ScaleRetryLimit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(createJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateJobRunOptions successfully`, func() {
				// Construct an instance of the EnvVarPrototype model
				envVarPrototypeModel := new(codeenginev2.EnvVarPrototype)
				Expect(envVarPrototypeModel).ToNot(BeNil())
				envVarPrototypeModel.Key = core.StringPtr("MY_VARIABLE")
				envVarPrototypeModel.Name = core.StringPtr("SOME")
				envVarPrototypeModel.Prefix = core.StringPtr("PREFIX_")
				envVarPrototypeModel.Reference = core.StringPtr("my-secret")
				envVarPrototypeModel.Type = core.StringPtr("literal")
				envVarPrototypeModel.Value = core.StringPtr("VALUE")
				Expect(envVarPrototypeModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarPrototypeModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarPrototypeModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarPrototypeModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarPrototypeModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMountPrototype model
				volumeMountPrototypeModel := new(codeenginev2.VolumeMountPrototype)
				Expect(volumeMountPrototypeModel).ToNot(BeNil())
				volumeMountPrototypeModel.MountPath = core.StringPtr("/app")
				volumeMountPrototypeModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountPrototypeModel.Reference = core.StringPtr("my-secret")
				volumeMountPrototypeModel.Type = core.StringPtr("secret")
				Expect(volumeMountPrototypeModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountPrototypeModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountPrototypeModel.Reference).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountPrototypeModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the CreateJobRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createJobRunOptionsModel := codeEngineService.NewCreateJobRunOptions(projectID)
				createJobRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobRunOptionsModel.SetImageReference("icr.io/codeengine/helloworld")
				createJobRunOptionsModel.SetImageSecret("my-secret")
				createJobRunOptionsModel.SetJobName("my-job")
				createJobRunOptionsModel.SetName("my-job-run")
				createJobRunOptionsModel.SetRunArguments([]string{"testString"})
				createJobRunOptionsModel.SetRunAsUser(int64(1001))
				createJobRunOptionsModel.SetRunCommands([]string{"testString"})
				createJobRunOptionsModel.SetRunEnvVariables([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel})
				createJobRunOptionsModel.SetRunMode("daemon")
				createJobRunOptionsModel.SetRunServiceAccount("default")
				createJobRunOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel})
				createJobRunOptionsModel.SetScaleArraySpec("1-5,7-8,10")
				createJobRunOptionsModel.SetScaleCpuLimit("1")
				createJobRunOptionsModel.SetScaleEphemeralStorageLimit("4G")
				createJobRunOptionsModel.SetScaleMaxExecutionTime(int64(7200))
				createJobRunOptionsModel.SetScaleMemoryLimit("4G")
				createJobRunOptionsModel.SetScaleRetryLimit(int64(3))
				createJobRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createJobRunOptionsModel).ToNot(BeNil())
				Expect(createJobRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createJobRunOptionsModel.ImageReference).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(createJobRunOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(createJobRunOptionsModel.JobName).To(Equal(core.StringPtr("my-job")))
				Expect(createJobRunOptionsModel.Name).To(Equal(core.StringPtr("my-job-run")))
				Expect(createJobRunOptionsModel.RunArguments).To(Equal([]string{"testString"}))
				Expect(createJobRunOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(createJobRunOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(createJobRunOptionsModel.RunEnvVariables).To(Equal([]codeenginev2.EnvVarPrototype{*envVarPrototypeModel}))
				Expect(createJobRunOptionsModel.RunMode).To(Equal(core.StringPtr("daemon")))
				Expect(createJobRunOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createJobRunOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel}))
				Expect(createJobRunOptionsModel.ScaleArraySpec).To(Equal(core.StringPtr("1-5,7-8,10")))
				Expect(createJobRunOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(createJobRunOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobRunOptionsModel.ScaleMaxExecutionTime).To(Equal(core.Int64Ptr(int64(7200))))
				Expect(createJobRunOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobRunOptionsModel.ScaleRetryLimit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(createJobRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "my-project"
				createProjectOptionsModel := codeEngineService.NewCreateProjectOptions(createProjectOptionsName)
				createProjectOptionsModel.SetName("my-project")
				createProjectOptionsModel.SetResourceGroupID("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.SetTags([]string{"testString"})
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("my-project")))
				Expect(createProjectOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")))
				Expect(createProjectOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSecretOptions successfully`, func() {
				// Construct an instance of the CreateSecretOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createSecretOptionsFormat := "generic"
				createSecretOptionsName := "my-secret"
				createSecretOptionsModel := codeEngineService.NewCreateSecretOptions(projectID, createSecretOptionsFormat, createSecretOptionsName)
				createSecretOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.SetFormat("generic")
				createSecretOptionsModel.SetName("my-secret")
				createSecretOptionsModel.SetData(make(map[string]string))
				createSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSecretOptionsModel).ToNot(BeNil())
				Expect(createSecretOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createSecretOptionsModel.Format).To(Equal(core.StringPtr("generic")))
				Expect(createSecretOptionsModel.Name).To(Equal(core.StringPtr("my-secret")))
				Expect(createSecretOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAppOptions successfully`, func() {
				// Construct an instance of the DeleteAppOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-app"
				deleteAppOptionsModel := codeEngineService.NewDeleteAppOptions(projectID, name)
				deleteAppOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.SetName("my-app")
				deleteAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAppOptionsModel).ToNot(BeNil())
				Expect(deleteAppOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(deleteAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAppRevisionOptions successfully`, func() {
				// Construct an instance of the DeleteAppRevisionOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				name := "my-app-001"
				deleteAppRevisionOptionsModel := codeEngineService.NewDeleteAppRevisionOptions(projectID, appName, name)
				deleteAppRevisionOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.SetAppName("my-app")
				deleteAppRevisionOptionsModel.SetName("my-app-001")
				deleteAppRevisionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAppRevisionOptionsModel).ToNot(BeNil())
				Expect(deleteAppRevisionOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteAppRevisionOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(deleteAppRevisionOptionsModel.Name).To(Equal(core.StringPtr("my-app-001")))
				Expect(deleteAppRevisionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBuildOptions successfully`, func() {
				// Construct an instance of the DeleteBuildOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-build"
				deleteBuildOptionsModel := codeEngineService.NewDeleteBuildOptions(projectID, name)
				deleteBuildOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.SetName("my-build")
				deleteBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBuildOptionsModel).ToNot(BeNil())
				Expect(deleteBuildOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(deleteBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBuildRunOptions successfully`, func() {
				// Construct an instance of the DeleteBuildRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-build-run"
				deleteBuildRunOptionsModel := codeEngineService.NewDeleteBuildRunOptions(projectID, name)
				deleteBuildRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildRunOptionsModel.SetName("my-build-run")
				deleteBuildRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBuildRunOptionsModel).ToNot(BeNil())
				Expect(deleteBuildRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteBuildRunOptionsModel.Name).To(Equal(core.StringPtr("my-build-run")))
				Expect(deleteBuildRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigMapOptions successfully`, func() {
				// Construct an instance of the DeleteConfigMapOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-config-map"
				deleteConfigMapOptionsModel := codeEngineService.NewDeleteConfigMapOptions(projectID, name)
				deleteConfigMapOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigMapOptionsModel.SetName("my-config-map")
				deleteConfigMapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigMapOptionsModel).ToNot(BeNil())
				Expect(deleteConfigMapOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteConfigMapOptionsModel.Name).To(Equal(core.StringPtr("my-config-map")))
				Expect(deleteConfigMapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteJobOptions successfully`, func() {
				// Construct an instance of the DeleteJobOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-job"
				deleteJobOptionsModel := codeEngineService.NewDeleteJobOptions(projectID, name)
				deleteJobOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.SetName("my-job")
				deleteJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteJobOptionsModel).ToNot(BeNil())
				Expect(deleteJobOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(deleteJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteJobRunOptions successfully`, func() {
				// Construct an instance of the DeleteJobRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-job"
				deleteJobRunOptionsModel := codeEngineService.NewDeleteJobRunOptions(projectID, name)
				deleteJobRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobRunOptionsModel.SetName("my-job")
				deleteJobRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteJobRunOptionsModel).ToNot(BeNil())
				Expect(deleteJobRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteJobRunOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(deleteJobRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "15314cc3-85b4-4338-903f-c28cdee6d005"
				deleteProjectOptionsModel := codeEngineService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSecretOptions successfully`, func() {
				// Construct an instance of the DeleteSecretOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-secret"
				deleteSecretOptionsModel := codeEngineService.NewDeleteSecretOptions(projectID, name)
				deleteSecretOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.SetName("my-secret")
				deleteSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSecretOptionsModel).ToNot(BeNil())
				Expect(deleteSecretOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteSecretOptionsModel.Name).To(Equal(core.StringPtr("my-secret")))
				Expect(deleteSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAppOptions successfully`, func() {
				// Construct an instance of the GetAppOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-app"
				getAppOptionsModel := codeEngineService.NewGetAppOptions(projectID, name)
				getAppOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.SetName("my-app")
				getAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAppOptionsModel).ToNot(BeNil())
				Expect(getAppOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(getAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAppRevisionOptions successfully`, func() {
				// Construct an instance of the GetAppRevisionOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				name := "my-app-001"
				getAppRevisionOptionsModel := codeEngineService.NewGetAppRevisionOptions(projectID, appName, name)
				getAppRevisionOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.SetAppName("my-app")
				getAppRevisionOptionsModel.SetName("my-app-001")
				getAppRevisionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAppRevisionOptionsModel).ToNot(BeNil())
				Expect(getAppRevisionOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getAppRevisionOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(getAppRevisionOptionsModel.Name).To(Equal(core.StringPtr("my-app-001")))
				Expect(getAppRevisionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBuildOptions successfully`, func() {
				// Construct an instance of the GetBuildOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-build"
				getBuildOptionsModel := codeEngineService.NewGetBuildOptions(projectID, name)
				getBuildOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.SetName("my-build")
				getBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBuildOptionsModel).ToNot(BeNil())
				Expect(getBuildOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(getBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBuildRunOptions successfully`, func() {
				// Construct an instance of the GetBuildRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-build-run"
				getBuildRunOptionsModel := codeEngineService.NewGetBuildRunOptions(projectID, name)
				getBuildRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildRunOptionsModel.SetName("my-build-run")
				getBuildRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBuildRunOptionsModel).ToNot(BeNil())
				Expect(getBuildRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getBuildRunOptionsModel.Name).To(Equal(core.StringPtr("my-build-run")))
				Expect(getBuildRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigMapOptions successfully`, func() {
				// Construct an instance of the GetConfigMapOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-config-map"
				getConfigMapOptionsModel := codeEngineService.NewGetConfigMapOptions(projectID, name)
				getConfigMapOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigMapOptionsModel.SetName("my-config-map")
				getConfigMapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigMapOptionsModel).ToNot(BeNil())
				Expect(getConfigMapOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getConfigMapOptionsModel.Name).To(Equal(core.StringPtr("my-config-map")))
				Expect(getConfigMapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetJobOptions successfully`, func() {
				// Construct an instance of the GetJobOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-job"
				getJobOptionsModel := codeEngineService.NewGetJobOptions(projectID, name)
				getJobOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.SetName("my-job")
				getJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getJobOptionsModel).ToNot(BeNil())
				Expect(getJobOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(getJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetJobRunOptions successfully`, func() {
				// Construct an instance of the GetJobRunOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-job"
				getJobRunOptionsModel := codeEngineService.NewGetJobRunOptions(projectID, name)
				getJobRunOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobRunOptionsModel.SetName("my-job")
				getJobRunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getJobRunOptionsModel).ToNot(BeNil())
				Expect(getJobRunOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getJobRunOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(getJobRunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "15314cc3-85b4-4338-903f-c28cdee6d005"
				getProjectOptionsModel := codeEngineService.NewGetProjectOptions(id)
				getProjectOptionsModel.SetID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSecretOptions successfully`, func() {
				// Construct an instance of the GetSecretOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-secret"
				getSecretOptionsModel := codeEngineService.NewGetSecretOptions(projectID, name)
				getSecretOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SetName("my-secret")
				getSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSecretOptionsModel).ToNot(BeNil())
				Expect(getSecretOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getSecretOptionsModel.Name).To(Equal(core.StringPtr("my-secret")))
				Expect(getSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAppRevisionsOptions successfully`, func() {
				// Construct an instance of the ListAppRevisionsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				listAppRevisionsOptionsModel := codeEngineService.NewListAppRevisionsOptions(projectID, appName)
				listAppRevisionsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.SetAppName("my-app")
				listAppRevisionsOptionsModel.SetLimit(int64(100))
				listAppRevisionsOptionsModel.SetStart("testString")
				listAppRevisionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAppRevisionsOptionsModel).ToNot(BeNil())
				Expect(listAppRevisionsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listAppRevisionsOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(listAppRevisionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listAppRevisionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAppRevisionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAppsOptions successfully`, func() {
				// Construct an instance of the ListAppsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listAppsOptionsModel := codeEngineService.NewListAppsOptions(projectID)
				listAppsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.SetLimit(int64(100))
				listAppsOptionsModel.SetStart("testString")
				listAppsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAppsOptionsModel).ToNot(BeNil())
				Expect(listAppsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listAppsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listAppsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAppsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBuildRunsOptions successfully`, func() {
				// Construct an instance of the ListBuildRunsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listBuildRunsOptionsModel := codeEngineService.NewListBuildRunsOptions(projectID)
				listBuildRunsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildRunsOptionsModel.SetBuildName("my-build")
				listBuildRunsOptionsModel.SetLimit(int64(100))
				listBuildRunsOptionsModel.SetStart("testString")
				listBuildRunsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBuildRunsOptionsModel).ToNot(BeNil())
				Expect(listBuildRunsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listBuildRunsOptionsModel.BuildName).To(Equal(core.StringPtr("my-build")))
				Expect(listBuildRunsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listBuildRunsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBuildRunsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBuildsOptions successfully`, func() {
				// Construct an instance of the ListBuildsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listBuildsOptionsModel := codeEngineService.NewListBuildsOptions(projectID)
				listBuildsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.SetLimit(int64(100))
				listBuildsOptionsModel.SetStart("testString")
				listBuildsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBuildsOptionsModel).ToNot(BeNil())
				Expect(listBuildsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listBuildsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listBuildsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBuildsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigMapsOptions successfully`, func() {
				// Construct an instance of the ListConfigMapsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listConfigMapsOptionsModel := codeEngineService.NewListConfigMapsOptions(projectID)
				listConfigMapsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigMapsOptionsModel.SetLimit(int64(100))
				listConfigMapsOptionsModel.SetStart("testString")
				listConfigMapsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigMapsOptionsModel).ToNot(BeNil())
				Expect(listConfigMapsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listConfigMapsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listConfigMapsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listConfigMapsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobRunsOptions successfully`, func() {
				// Construct an instance of the ListJobRunsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listJobRunsOptionsModel := codeEngineService.NewListJobRunsOptions(projectID)
				listJobRunsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobRunsOptionsModel.SetJobName("my-job")
				listJobRunsOptionsModel.SetLimit(int64(100))
				listJobRunsOptionsModel.SetStart("testString")
				listJobRunsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobRunsOptionsModel).ToNot(BeNil())
				Expect(listJobRunsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listJobRunsOptionsModel.JobName).To(Equal(core.StringPtr("my-job")))
				Expect(listJobRunsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listJobRunsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listJobRunsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobsOptions successfully`, func() {
				// Construct an instance of the ListJobsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listJobsOptionsModel := codeEngineService.NewListJobsOptions(projectID)
				listJobsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.SetLimit(int64(100))
				listJobsOptionsModel.SetStart("testString")
				listJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobsOptionsModel).ToNot(BeNil())
				Expect(listJobsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listJobsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listJobsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := codeEngineService.NewListProjectsOptions()
				listProjectsOptionsModel.SetLimit(int64(100))
				listProjectsOptionsModel.SetStart("testString")
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listProjectsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSecretsOptions successfully`, func() {
				// Construct an instance of the ListSecretsOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listSecretsOptionsModel := codeEngineService.NewListSecretsOptions(projectID)
				listSecretsOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.SetLimit(int64(100))
				listSecretsOptionsModel.SetStart("testString")
				listSecretsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSecretsOptionsModel).ToNot(BeNil())
				Expect(listSecretsOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listSecretsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listSecretsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listSecretsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceConfigMapOptions successfully`, func() {
				// Construct an instance of the ReplaceConfigMapOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-config-map"
				ifMatch := "testString"
				replaceConfigMapOptionsModel := codeEngineService.NewReplaceConfigMapOptions(projectID, name, ifMatch)
				replaceConfigMapOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceConfigMapOptionsModel.SetName("my-config-map")
				replaceConfigMapOptionsModel.SetIfMatch("testString")
				replaceConfigMapOptionsModel.SetData(make(map[string]string))
				replaceConfigMapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceConfigMapOptionsModel).ToNot(BeNil())
				Expect(replaceConfigMapOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(replaceConfigMapOptionsModel.Name).To(Equal(core.StringPtr("my-config-map")))
				Expect(replaceConfigMapOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceConfigMapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(replaceConfigMapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceSecretOptions successfully`, func() {
				// Construct an instance of the ReplaceSecretOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-secret"
				ifMatch := "testString"
				replaceSecretOptionsModel := codeEngineService.NewReplaceSecretOptions(projectID, name, ifMatch)
				replaceSecretOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				replaceSecretOptionsModel.SetName("my-secret")
				replaceSecretOptionsModel.SetIfMatch("testString")
				replaceSecretOptionsModel.SetData(make(map[string]string))
				replaceSecretOptionsModel.SetFormat("generic")
				replaceSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceSecretOptionsModel).ToNot(BeNil())
				Expect(replaceSecretOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(replaceSecretOptionsModel.Name).To(Equal(core.StringPtr("my-secret")))
				Expect(replaceSecretOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceSecretOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(replaceSecretOptionsModel.Format).To(Equal(core.StringPtr("generic")))
				Expect(replaceSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAppOptions successfully`, func() {
				// Construct an instance of the UpdateAppOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-app"
				ifMatch := "testString"
				app := map[string]interface{}{"anyKey": "anyValue"}
				updateAppOptionsModel := codeEngineService.NewUpdateAppOptions(projectID, name, ifMatch, app)
				updateAppOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.SetName("my-app")
				updateAppOptionsModel.SetIfMatch("testString")
				updateAppOptionsModel.SetApp(map[string]interface{}{"anyKey": "anyValue"})
				updateAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAppOptionsModel).ToNot(BeNil())
				Expect(updateAppOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(updateAppOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAppOptionsModel.App).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBuildOptions successfully`, func() {
				// Construct an instance of the UpdateBuildOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-build"
				ifMatch := "testString"
				build := map[string]interface{}{"anyKey": "anyValue"}
				updateBuildOptionsModel := codeEngineService.NewUpdateBuildOptions(projectID, name, ifMatch, build)
				updateBuildOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.SetName("my-build")
				updateBuildOptionsModel.SetIfMatch("testString")
				updateBuildOptionsModel.SetBuild(map[string]interface{}{"anyKey": "anyValue"})
				updateBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBuildOptionsModel).ToNot(BeNil())
				Expect(updateBuildOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(updateBuildOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateBuildOptionsModel.Build).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateJobOptions successfully`, func() {
				// Construct an instance of the UpdateJobOptions model
				projectID := "15314cc3-85b4-4338-903f-c28cdee6d005"
				name := "my-job"
				ifMatch := "testString"
				job := map[string]interface{}{"anyKey": "anyValue"}
				updateJobOptionsModel := codeEngineService.NewUpdateJobOptions(projectID, name, ifMatch, job)
				updateJobOptionsModel.SetProjectID("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.SetName("my-job")
				updateJobOptionsModel.SetIfMatch("testString")
				updateJobOptionsModel.SetJob(map[string]interface{}{"anyKey": "anyValue"})
				updateJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateJobOptionsModel).ToNot(BeNil())
				Expect(updateJobOptionsModel.ProjectID).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(updateJobOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateJobOptionsModel.Job).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVolumeMountPrototype successfully`, func() {
				mountPath := "/app"
				reference := "my-secret"
				typeVar := "secret"
				_model, err := codeEngineService.NewVolumeMountPrototype(mountPath, reference, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
