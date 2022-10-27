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

package codeenginev2_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
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
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}]}`)
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
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}]}`)
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"account_id":"4329073d16d2f3663f74bfa955259139","created":"2021-03-29T12:18:13.992359829Z","crn":"crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::","details":"succeeded","id":"15314cc3-85b4-4338-903f-c28cdee6d005","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"project-name","reason":"create","region":"us-east","resource_group_id":"5c49eabcf5e85881a37e2d100a33b3df","status":"active","type":"project/v2"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"account_id":"4329073d16d2f3663f74bfa955259139","created":"2021-03-29T12:18:13.992359829Z","crn":"crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::","details":"succeeded","id":"15314cc3-85b4-4338-903f-c28cdee6d005","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"project-name","reason":"create","region":"us-east","resource_group_id":"5c49eabcf5e85881a37e2d100a33b3df","status":"active","type":"project/v2"}],"total_count":2,"limit":1}`)
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
				createProjectOptionsModel.Region = core.StringPtr("us-east")
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
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}`)
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
				createProjectOptionsModel.Region = core.StringPtr("us-east")
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
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}`)
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
				createProjectOptionsModel.Region = core.StringPtr("us-east")
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.Tags = []string{"testString"}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(codeenginev2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("my-project")
				createProjectOptionsModel.Region = core.StringPtr("us-east")
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
				createProjectOptionsModel.Region = core.StringPtr("us-east")
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}`)
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "created": "2021-03-29T12:18:13.992359829Z", "crn": "crn:v1:bluemix:public:codeengine:eu-de:a/4329073d16d2f3663f74bfa955259139:15314cc3-85b4-4338-903f-c28cdee6d005::", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "project-name", "reason": "create", "region": "us-east", "resource_group_id": "5c49eabcf5e85881a37e2d100a33b3df", "status": "active", "type": "project/v2"}`)
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				deleteProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				deleteProjectOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listBuildsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"builds": [{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listBuildsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"builds": [{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listBuildsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listBuildsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listBuildsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"builds":[{"ce_owner_reference":"CeOwnerReference","created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","output_image":"stg.icr.io/icr_namespace/image-name","output_secret":"ce-default-icr-us-south","reason":"create","source_context_dir":"SourceContextDir","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"active","strategy_name":"dockerfile","strategy_size":"medium","strategy_spec_file":"Dockerfile","timeout":600,"type":"Type"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"builds":[{"ce_owner_reference":"CeOwnerReference","created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","output_image":"stg.icr.io/icr_namespace/image-name","output_secret":"ce-default-icr-us-south","reason":"create","source_context_dir":"SourceContextDir","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":"active","strategy_name":"dockerfile","strategy_size":"medium","strategy_spec_file":"Dockerfile","timeout":600,"type":"Type"}]}`)
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
				createBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
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
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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
				createBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
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
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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
				createBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
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
				createBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
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
				createBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.Name = core.StringPtr("my-build")
				createBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildOptionsModel.SourceType = core.StringPtr("git")
				createBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
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
				getBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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
				getBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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
				getBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
				getBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
				getBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
				deleteBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.BuildName = core.StringPtr("my-build")
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
				deleteBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.BuildName = core.StringPtr("my-build")
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

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.BuildName = core.StringPtr("my-build")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				updateBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				updateBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				updateBuildOptionsModel.SourceRevision = core.StringPtr("main")
				updateBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				updateBuildOptionsModel.SourceType = core.StringPtr("git")
				updateBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
				updateBuildOptionsModel.StrategySize = core.StringPtr("medium")
				updateBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				updateBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.BuildName = core.StringPtr("my-build")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				updateBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				updateBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				updateBuildOptionsModel.SourceRevision = core.StringPtr("main")
				updateBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				updateBuildOptionsModel.SourceType = core.StringPtr("git")
				updateBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
				updateBuildOptionsModel.StrategySize = core.StringPtr("medium")
				updateBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				updateBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "reason": "create", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": "active", "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
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

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.BuildName = core.StringPtr("my-build")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				updateBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				updateBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				updateBuildOptionsModel.SourceRevision = core.StringPtr("main")
				updateBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				updateBuildOptionsModel.SourceType = core.StringPtr("git")
				updateBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
				updateBuildOptionsModel.StrategySize = core.StringPtr("medium")
				updateBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				updateBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
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

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.BuildName = core.StringPtr("my-build")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				updateBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				updateBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				updateBuildOptionsModel.SourceRevision = core.StringPtr("main")
				updateBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				updateBuildOptionsModel.SourceType = core.StringPtr("git")
				updateBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
				updateBuildOptionsModel.StrategySize = core.StringPtr("medium")
				updateBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				updateBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
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

				// Construct an instance of the UpdateBuildOptions model
				updateBuildOptionsModel := new(codeenginev2.UpdateBuildOptions)
				updateBuildOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.BuildName = core.StringPtr("my-build")
				updateBuildOptionsModel.Name = core.StringPtr("my-build")
				updateBuildOptionsModel.CeOwnerReference = core.StringPtr("testString")
				updateBuildOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				updateBuildOptionsModel.SourceContextDir = core.StringPtr("testString")
				updateBuildOptionsModel.SourceRevision = core.StringPtr("main")
				updateBuildOptionsModel.SourceSecret = core.StringPtr("testString")
				updateBuildOptionsModel.SourceType = core.StringPtr("git")
				updateBuildOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.StrategyName = core.StringPtr("dockerfile")
				updateBuildOptionsModel.StrategySize = core.StringPtr("medium")
				updateBuildOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				updateBuildOptionsModel.Timeout = core.Int64Ptr(int64(600))
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
	Describe(`ListBuildruns(listBuildrunsOptions *ListBuildrunsOptions) - Operation response error`, func() {
		listBuildrunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildrunsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBuildruns with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildrunsOptions model
				listBuildrunsOptionsModel := new(codeenginev2.ListBuildrunsOptions)
				listBuildrunsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildrunsOptionsModel.Start = core.StringPtr("testString")
				listBuildrunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuildruns(listBuildrunsOptions *ListBuildrunsOptions)`, func() {
		listBuildrunsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBuildrunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_runs": [{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuildruns successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListBuildrunsOptions model
				listBuildrunsOptionsModel := new(codeenginev2.ListBuildrunsOptions)
				listBuildrunsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildrunsOptionsModel.Start = core.StringPtr("testString")
				listBuildrunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListBuildrunsWithContext(ctx, listBuildrunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListBuildrunsWithContext(ctx, listBuildrunsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBuildrunsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"build_runs": [{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListBuildruns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListBuildruns(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBuildrunsOptions model
				listBuildrunsOptionsModel := new(codeenginev2.ListBuildrunsOptions)
				listBuildrunsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildrunsOptionsModel.Start = core.StringPtr("testString")
				listBuildrunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBuildruns with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildrunsOptions model
				listBuildrunsOptionsModel := new(codeenginev2.ListBuildrunsOptions)
				listBuildrunsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildrunsOptionsModel.Start = core.StringPtr("testString")
				listBuildrunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBuildrunsOptions model with no property values
				listBuildrunsOptionsModelNew := new(codeenginev2.ListBuildrunsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListBuildruns(listBuildrunsOptionsModelNew)
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
			It(`Invoke ListBuildruns successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListBuildrunsOptions model
				listBuildrunsOptionsModel := new(codeenginev2.ListBuildrunsOptions)
				listBuildrunsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listBuildrunsOptionsModel.Start = core.StringPtr("testString")
				listBuildrunsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListBuildruns(listBuildrunsOptionsModel)
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBuildrunsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"build_runs":[{"app_revision":"AppRevision","build":"Build","ce_owner_reference":"CeOwnerReference","created":"2022-09-13T11:41:35+02:00","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","output_image":"stg.icr.io/icr_namespace/image-name","output_secret":"ce-default-icr-us-south","service_account":"ServiceAccount","source_context_dir":"SourceContextDir","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":{"completion_time":"2022-09-22T17:40:00Z","last_task_run":"LastTaskRun","start_time":"2022-09-22T17:34:00Z"},"strategy_name":"dockerfile","strategy_size":"medium","strategy_spec_file":"Dockerfile","timeout":600,"type":"Type"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"build_runs":[{"app_revision":"AppRevision","build":"Build","ce_owner_reference":"CeOwnerReference","created":"2022-09-13T11:41:35+02:00","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","output_image":"stg.icr.io/icr_namespace/image-name","output_secret":"ce-default-icr-us-south","service_account":"ServiceAccount","source_context_dir":"SourceContextDir","source_revision":"main","source_secret":"SourceSecret","source_type":"git","source_url":"https://github.com/IBM/CodeEngine","status":{"completion_time":"2022-09-22T17:40:00Z","last_task_run":"LastTaskRun","start_time":"2022-09-22T17:34:00Z"},"strategy_name":"dockerfile","strategy_size":"medium","strategy_spec_file":"Dockerfile","timeout":600,"type":"Type"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BuildrunsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildrunsOptionsModel := &codeenginev2.ListBuildrunsOptions{
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildrunsPager(listBuildrunsOptionsModel)
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
			It(`Use BuildrunsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listBuildrunsOptionsModel := &codeenginev2.ListBuildrunsOptions{
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewBuildrunsPager(listBuildrunsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateBuildrun(createBuildrunOptions *CreateBuildrunOptions) - Operation response error`, func() {
		createBuildrunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildrunPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBuildrun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildrunOptions model
				createBuildrunOptionsModel := new(codeenginev2.CreateBuildrunOptions)
				createBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.Name = core.StringPtr("testString")
				createBuildrunOptionsModel.AppRevision = core.StringPtr("testString")
				createBuildrunOptionsModel.Build = core.StringPtr("testString")
				createBuildrunOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildrunOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildrunOptionsModel.ServiceAccount = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildrunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceType = core.StringPtr("git")
				createBuildrunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.StrategyName = core.StringPtr("dockerfile")
				createBuildrunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildrunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildrunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBuildrun(createBuildrunOptions *CreateBuildrunOptions)`, func() {
		createBuildrunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBuildrunPath))
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
					fmt.Fprintf(res, "%s", `{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
				}))
			})
			It(`Invoke CreateBuildrun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateBuildrunOptions model
				createBuildrunOptionsModel := new(codeenginev2.CreateBuildrunOptions)
				createBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.Name = core.StringPtr("testString")
				createBuildrunOptionsModel.AppRevision = core.StringPtr("testString")
				createBuildrunOptionsModel.Build = core.StringPtr("testString")
				createBuildrunOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildrunOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildrunOptionsModel.ServiceAccount = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildrunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceType = core.StringPtr("git")
				createBuildrunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.StrategyName = core.StringPtr("dockerfile")
				createBuildrunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildrunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildrunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateBuildrunWithContext(ctx, createBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateBuildrunWithContext(ctx, createBuildrunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createBuildrunPath))
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
					fmt.Fprintf(res, "%s", `{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
				}))
			})
			It(`Invoke CreateBuildrun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateBuildrun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateBuildrunOptions model
				createBuildrunOptionsModel := new(codeenginev2.CreateBuildrunOptions)
				createBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.Name = core.StringPtr("testString")
				createBuildrunOptionsModel.AppRevision = core.StringPtr("testString")
				createBuildrunOptionsModel.Build = core.StringPtr("testString")
				createBuildrunOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildrunOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildrunOptionsModel.ServiceAccount = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildrunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceType = core.StringPtr("git")
				createBuildrunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.StrategyName = core.StringPtr("dockerfile")
				createBuildrunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildrunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildrunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBuildrun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildrunOptions model
				createBuildrunOptionsModel := new(codeenginev2.CreateBuildrunOptions)
				createBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.Name = core.StringPtr("testString")
				createBuildrunOptionsModel.AppRevision = core.StringPtr("testString")
				createBuildrunOptionsModel.Build = core.StringPtr("testString")
				createBuildrunOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildrunOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildrunOptionsModel.ServiceAccount = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildrunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceType = core.StringPtr("git")
				createBuildrunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.StrategyName = core.StringPtr("dockerfile")
				createBuildrunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildrunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildrunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBuildrunOptions model with no property values
				createBuildrunOptionsModelNew := new(codeenginev2.CreateBuildrunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateBuildrun(createBuildrunOptionsModelNew)
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
			It(`Invoke CreateBuildrun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateBuildrunOptions model
				createBuildrunOptionsModel := new(codeenginev2.CreateBuildrunOptions)
				createBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.Name = core.StringPtr("testString")
				createBuildrunOptionsModel.AppRevision = core.StringPtr("testString")
				createBuildrunOptionsModel.Build = core.StringPtr("testString")
				createBuildrunOptionsModel.CeOwnerReference = core.StringPtr("testString")
				createBuildrunOptionsModel.OutputImage = core.StringPtr("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.OutputSecret = core.StringPtr("ce-default-icr-us-south")
				createBuildrunOptionsModel.ServiceAccount = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceContextDir = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceRevision = core.StringPtr("main")
				createBuildrunOptionsModel.SourceSecret = core.StringPtr("testString")
				createBuildrunOptionsModel.SourceType = core.StringPtr("git")
				createBuildrunOptionsModel.SourceURL = core.StringPtr("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.StrategyName = core.StringPtr("dockerfile")
				createBuildrunOptionsModel.StrategySize = core.StringPtr("medium")
				createBuildrunOptionsModel.StrategySpecFile = core.StringPtr("Dockerfile")
				createBuildrunOptionsModel.Timeout = core.Int64Ptr(int64(600))
				createBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateBuildrun(createBuildrunOptionsModel)
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
	Describe(`GetBuildrun(getBuildrunOptions *GetBuildrunOptions) - Operation response error`, func() {
		getBuildrunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildrunPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBuildrun with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildrunOptions model
				getBuildrunOptionsModel := new(codeenginev2.GetBuildrunOptions)
				getBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				getBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetBuildrun(getBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetBuildrun(getBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuildrun(getBuildrunOptions *GetBuildrunOptions)`, func() {
		getBuildrunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBuildrunPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
				}))
			})
			It(`Invoke GetBuildrun successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetBuildrunOptions model
				getBuildrunOptionsModel := new(codeenginev2.GetBuildrunOptions)
				getBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				getBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetBuildrunWithContext(ctx, getBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetBuildrun(getBuildrunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetBuildrunWithContext(ctx, getBuildrunOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBuildrunPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"app_revision": "AppRevision", "build": "Build", "ce_owner_reference": "CeOwnerReference", "created": "2022-09-13T11:41:35+02:00", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "output_image": "stg.icr.io/icr_namespace/image-name", "output_secret": "ce-default-icr-us-south", "service_account": "ServiceAccount", "source_context_dir": "SourceContextDir", "source_revision": "main", "source_secret": "SourceSecret", "source_type": "git", "source_url": "https://github.com/IBM/CodeEngine", "status": {"completion_time": "2022-09-22T17:40:00Z", "last_task_run": "LastTaskRun", "start_time": "2022-09-22T17:34:00Z"}, "strategy_name": "dockerfile", "strategy_size": "medium", "strategy_spec_file": "Dockerfile", "timeout": 600, "type": "Type"}`)
				}))
			})
			It(`Invoke GetBuildrun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetBuildrun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBuildrunOptions model
				getBuildrunOptionsModel := new(codeenginev2.GetBuildrunOptions)
				getBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				getBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetBuildrun(getBuildrunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBuildrun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildrunOptions model
				getBuildrunOptionsModel := new(codeenginev2.GetBuildrunOptions)
				getBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				getBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetBuildrun(getBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBuildrunOptions model with no property values
				getBuildrunOptionsModelNew := new(codeenginev2.GetBuildrunOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetBuildrun(getBuildrunOptionsModelNew)
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
			It(`Invoke GetBuildrun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetBuildrunOptions model
				getBuildrunOptionsModel := new(codeenginev2.GetBuildrunOptions)
				getBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				getBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetBuildrun(getBuildrunOptionsModel)
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
	Describe(`DeleteBuildrun(deleteBuildrunOptions *DeleteBuildrunOptions)`, func() {
		deleteBuildrunPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/build_runs/my-build-run"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBuildrunPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteBuildrun successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteBuildrun(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBuildrunOptions model
				deleteBuildrunOptionsModel := new(codeenginev2.DeleteBuildrunOptions)
				deleteBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				deleteBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteBuildrun(deleteBuildrunOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBuildrun with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteBuildrunOptions model
				deleteBuildrunOptionsModel := new(codeenginev2.DeleteBuildrunOptions)
				deleteBuildrunOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildrunOptionsModel.BuildRunName = core.StringPtr("my-build-run")
				deleteBuildrunOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteBuildrun(deleteBuildrunOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBuildrunOptions model with no property values
				deleteBuildrunOptionsModelNew := new(codeenginev2.DeleteBuildrunOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteBuildrun(deleteBuildrunOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigmaps(listConfigmapsOptions *ListConfigmapsOptions) - Operation response error`, func() {
		listConfigmapsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigmaps with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsOptions model
				listConfigmapsOptionsModel := new(codeenginev2.ListConfigmapsOptions)
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigmapsOptionsModel.Start = core.StringPtr("testString")
				listConfigmapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigmaps(listConfigmapsOptions *ListConfigmapsOptions)`, func() {
		listConfigmapsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"config_maps": [{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigmaps successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigmapsOptions model
				listConfigmapsOptionsModel := new(codeenginev2.ListConfigmapsOptions)
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigmapsOptionsModel.Start = core.StringPtr("testString")
				listConfigmapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListConfigmapsWithContext(ctx, listConfigmapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListConfigmapsWithContext(ctx, listConfigmapsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"config_maps": [{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigmaps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListConfigmaps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigmapsOptions model
				listConfigmapsOptionsModel := new(codeenginev2.ListConfigmapsOptions)
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigmapsOptionsModel.Start = core.StringPtr("testString")
				listConfigmapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigmaps with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsOptions model
				listConfigmapsOptionsModel := new(codeenginev2.ListConfigmapsOptions)
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigmapsOptionsModel.Start = core.StringPtr("testString")
				listConfigmapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigmapsOptions model with no property values
				listConfigmapsOptionsModelNew := new(codeenginev2.ListConfigmapsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListConfigmaps(listConfigmapsOptionsModelNew)
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
			It(`Invoke ListConfigmaps successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsOptions model
				listConfigmapsOptionsModel := new(codeenginev2.ListConfigmapsOptions)
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listConfigmapsOptionsModel.Start = core.StringPtr("testString")
				listConfigmapsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListConfigmaps(listConfigmapsOptionsModel)
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"config_maps":[{"created":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","immutable":false,"links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","type":"Type"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"config_maps":[{"created":"2022-09-13T11:41:35+02:00","data":{"mapKey":"Inner"},"id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","immutable":false,"links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","type":"Type"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ConfigmapsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listConfigmapsOptionsModel := &codeenginev2.ListConfigmapsOptions{
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewConfigmapsPager(listConfigmapsOptionsModel)
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
			It(`Use ConfigmapsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listConfigmapsOptionsModel := &codeenginev2.ListConfigmapsOptions{
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewConfigmapsPager(listConfigmapsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateConfigmap(createConfigmapOptions *CreateConfigmapOptions) - Operation response error`, func() {
		createConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfigmap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapOptions model
				createConfigmapOptionsModel := new(codeenginev2.CreateConfigmapOptions)
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfigmap(createConfigmapOptions *CreateConfigmapOptions)`, func() {
		createConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapPath))
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateConfigmap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateConfigmapOptions model
				createConfigmapOptionsModel := new(codeenginev2.CreateConfigmapOptions)
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateConfigmapWithContext(ctx, createConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateConfigmapWithContext(ctx, createConfigmapOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapPath))
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateConfigmap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateConfigmapOptions model
				createConfigmapOptionsModel := new(codeenginev2.CreateConfigmapOptions)
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfigmap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapOptions model
				createConfigmapOptionsModel := new(codeenginev2.CreateConfigmapOptions)
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigmapOptions model with no property values
				createConfigmapOptionsModelNew := new(codeenginev2.CreateConfigmapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateConfigmap(createConfigmapOptionsModelNew)
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
			It(`Invoke CreateConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapOptions model
				createConfigmapOptionsModel := new(codeenginev2.CreateConfigmapOptions)
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateConfigmap(createConfigmapOptionsModel)
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
	Describe(`GetConfigmap(getConfigmapOptions *GetConfigmapOptions) - Operation response error`, func() {
		getConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfigmap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapOptions model
				getConfigmapOptionsModel := new(codeenginev2.GetConfigmapOptions)
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				getConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetConfigmap(getConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetConfigmap(getConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigmap(getConfigmapOptions *GetConfigmapOptions)`, func() {
		getConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke GetConfigmap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigmapOptions model
				getConfigmapOptionsModel := new(codeenginev2.GetConfigmapOptions)
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				getConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetConfigmapWithContext(ctx, getConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetConfigmap(getConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetConfigmapWithContext(ctx, getConfigmapOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke GetConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetConfigmap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigmapOptions model
				getConfigmapOptionsModel := new(codeenginev2.GetConfigmapOptions)
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				getConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetConfigmap(getConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigmap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapOptions model
				getConfigmapOptionsModel := new(codeenginev2.GetConfigmapOptions)
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				getConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetConfigmap(getConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigmapOptions model with no property values
				getConfigmapOptionsModelNew := new(codeenginev2.GetConfigmapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetConfigmap(getConfigmapOptionsModelNew)
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
			It(`Invoke GetConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapOptions model
				getConfigmapOptionsModel := new(codeenginev2.GetConfigmapOptions)
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				getConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetConfigmap(getConfigmapOptionsModel)
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
	Describe(`DeleteConfigmap(deleteConfigmapOptions *DeleteConfigmapOptions)`, func() {
		deleteConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigmapPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteConfigmap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteConfigmapOptions model
				deleteConfigmapOptionsModel := new(codeenginev2.DeleteConfigmapOptions)
				deleteConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				deleteConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteConfigmap(deleteConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteConfigmap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigmapOptions model
				deleteConfigmapOptionsModel := new(codeenginev2.DeleteConfigmapOptions)
				deleteConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				deleteConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteConfigmap(deleteConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteConfigmapOptions model with no property values
				deleteConfigmapOptionsModelNew := new(codeenginev2.DeleteConfigmapOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteConfigmap(deleteConfigmapOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfigmap(updateConfigmapOptions *UpdateConfigmapOptions) - Operation response error`, func() {
		updateConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfigmap with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapOptions model
				updateConfigmapOptionsModel := new(codeenginev2.UpdateConfigmapOptions)
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfigmap(updateConfigmapOptions *UpdateConfigmapOptions)`, func() {
		updateConfigmapPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/config_maps/my-config-map"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateConfigmap successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the UpdateConfigmapOptions model
				updateConfigmapOptionsModel := new(codeenginev2.UpdateConfigmapOptions)
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateConfigmapWithContext(ctx, updateConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateConfigmapWithContext(ctx, updateConfigmapOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateConfigmap(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateConfigmapOptions model
				updateConfigmapOptionsModel := new(codeenginev2.UpdateConfigmapOptions)
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfigmap with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapOptions model
				updateConfigmapOptionsModel := new(codeenginev2.UpdateConfigmapOptions)
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigmapOptions model with no property values
				updateConfigmapOptionsModelNew := new(codeenginev2.UpdateConfigmapOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateConfigmap(updateConfigmapOptionsModelNew)
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
			It(`Invoke UpdateConfigmap successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapOptions model
				updateConfigmapOptionsModel := new(codeenginev2.UpdateConfigmapOptions)
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.ConfigMapName = core.StringPtr("my-config-map")
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateConfigmap(updateConfigmapOptionsModel)
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
	Describe(`ListSecrets(listSecretsOptions *ListSecretsOptions) - Operation response error`, func() {
		listSecretsPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSecretsPath))
					Expect(req.Method).To(Equal("GET"))
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
				listSecretsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "secrets": [{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}]}`)
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
				listSecretsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "secrets": [{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}]}`)
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
				listSecretsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listSecretsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listSecretsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				createSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Name = core.StringPtr("testString")
				createSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				createSecretOptionsModel.CeComponents = []string{"testString"}
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Format = core.StringPtr("testString")
				createSecretOptionsModel.Immutable = core.BoolPtr(true)
				createSecretOptionsModel.ResourceID = core.StringPtr("testString")
				createSecretOptionsModel.ResourceType = core.StringPtr("testString")
				createSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				createSecretOptionsModel.Role = core.StringPtr("testString")
				createSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				createSecretOptionsModel.Target = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
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
				createSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Name = core.StringPtr("testString")
				createSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				createSecretOptionsModel.CeComponents = []string{"testString"}
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Format = core.StringPtr("testString")
				createSecretOptionsModel.Immutable = core.BoolPtr(true)
				createSecretOptionsModel.ResourceID = core.StringPtr("testString")
				createSecretOptionsModel.ResourceType = core.StringPtr("testString")
				createSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				createSecretOptionsModel.Role = core.StringPtr("testString")
				createSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				createSecretOptionsModel.Target = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
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
				createSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Name = core.StringPtr("testString")
				createSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				createSecretOptionsModel.CeComponents = []string{"testString"}
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Format = core.StringPtr("testString")
				createSecretOptionsModel.Immutable = core.BoolPtr(true)
				createSecretOptionsModel.ResourceID = core.StringPtr("testString")
				createSecretOptionsModel.ResourceType = core.StringPtr("testString")
				createSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				createSecretOptionsModel.Role = core.StringPtr("testString")
				createSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				createSecretOptionsModel.Target = core.StringPtr("testString")
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
				createSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Name = core.StringPtr("testString")
				createSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				createSecretOptionsModel.CeComponents = []string{"testString"}
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Format = core.StringPtr("testString")
				createSecretOptionsModel.Immutable = core.BoolPtr(true)
				createSecretOptionsModel.ResourceID = core.StringPtr("testString")
				createSecretOptionsModel.ResourceType = core.StringPtr("testString")
				createSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				createSecretOptionsModel.Role = core.StringPtr("testString")
				createSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				createSecretOptionsModel.Target = core.StringPtr("testString")
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
				createSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.Name = core.StringPtr("testString")
				createSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				createSecretOptionsModel.CeComponents = []string{"testString"}
				createSecretOptionsModel.Data = make(map[string]string)
				createSecretOptionsModel.Format = core.StringPtr("testString")
				createSecretOptionsModel.Immutable = core.BoolPtr(true)
				createSecretOptionsModel.ResourceID = core.StringPtr("testString")
				createSecretOptionsModel.ResourceType = core.StringPtr("testString")
				createSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				createSecretOptionsModel.Role = core.StringPtr("testString")
				createSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				createSecretOptionsModel.Target = core.StringPtr("testString")
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
				getSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
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
				getSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
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
				getSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
				getSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
				getSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
				deleteSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
				deleteSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.SecretName = core.StringPtr("my-secret")
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
	Describe(`UpdateSecret(updateSecretOptions *UpdateSecretOptions) - Operation response error`, func() {
		updateSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSecretPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSecret with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateSecretOptions model
				updateSecretOptionsModel := new(codeenginev2.UpdateSecretOptions)
				updateSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SecretName = core.StringPtr("my-secret")
				updateSecretOptionsModel.Name = core.StringPtr("testString")
				updateSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				updateSecretOptionsModel.CeComponents = []string{"testString"}
				updateSecretOptionsModel.Data = make(map[string]string)
				updateSecretOptionsModel.Format = core.StringPtr("testString")
				updateSecretOptionsModel.Immutable = core.BoolPtr(true)
				updateSecretOptionsModel.ResourceID = core.StringPtr("testString")
				updateSecretOptionsModel.ResourceType = core.StringPtr("testString")
				updateSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				updateSecretOptionsModel.Role = core.StringPtr("testString")
				updateSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				updateSecretOptionsModel.Target = core.StringPtr("testString")
				updateSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateSecret(updateSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateSecret(updateSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSecret(updateSecretOptions *UpdateSecretOptions)`, func() {
		updateSecretPath := "/projects/15314cc3-85b4-4338-903f-c28cdee6d005/secrets/my-secret"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSecretPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateSecret successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSecretOptions model
				updateSecretOptionsModel := new(codeenginev2.UpdateSecretOptions)
				updateSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SecretName = core.StringPtr("my-secret")
				updateSecretOptionsModel.Name = core.StringPtr("testString")
				updateSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				updateSecretOptionsModel.CeComponents = []string{"testString"}
				updateSecretOptionsModel.Data = make(map[string]string)
				updateSecretOptionsModel.Format = core.StringPtr("testString")
				updateSecretOptionsModel.Immutable = core.BoolPtr(true)
				updateSecretOptionsModel.ResourceID = core.StringPtr("testString")
				updateSecretOptionsModel.ResourceType = core.StringPtr("testString")
				updateSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				updateSecretOptionsModel.Role = core.StringPtr("testString")
				updateSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				updateSecretOptionsModel.Target = core.StringPtr("testString")
				updateSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateSecretWithContext(ctx, updateSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateSecret(updateSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateSecretWithContext(ctx, updateSecretOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSecretPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"binding_secret_ref": "BindingSecretRef", "ce_components": ["CeComponents"], "created": "2022-09-13T11:41:35+02:00", "data": {"mapKey": "Inner"}, "format": "Format", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "immutable": false, "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "resource_id": "ResourceID", "resource_type": "ResourceType", "resourcekey_id": "ResourcekeyID", "role": "Role", "serviceid_crn": "ServiceidCrn", "target": "Target", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSecretOptions model
				updateSecretOptionsModel := new(codeenginev2.UpdateSecretOptions)
				updateSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SecretName = core.StringPtr("my-secret")
				updateSecretOptionsModel.Name = core.StringPtr("testString")
				updateSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				updateSecretOptionsModel.CeComponents = []string{"testString"}
				updateSecretOptionsModel.Data = make(map[string]string)
				updateSecretOptionsModel.Format = core.StringPtr("testString")
				updateSecretOptionsModel.Immutable = core.BoolPtr(true)
				updateSecretOptionsModel.ResourceID = core.StringPtr("testString")
				updateSecretOptionsModel.ResourceType = core.StringPtr("testString")
				updateSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				updateSecretOptionsModel.Role = core.StringPtr("testString")
				updateSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				updateSecretOptionsModel.Target = core.StringPtr("testString")
				updateSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateSecret(updateSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSecret with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateSecretOptions model
				updateSecretOptionsModel := new(codeenginev2.UpdateSecretOptions)
				updateSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SecretName = core.StringPtr("my-secret")
				updateSecretOptionsModel.Name = core.StringPtr("testString")
				updateSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				updateSecretOptionsModel.CeComponents = []string{"testString"}
				updateSecretOptionsModel.Data = make(map[string]string)
				updateSecretOptionsModel.Format = core.StringPtr("testString")
				updateSecretOptionsModel.Immutable = core.BoolPtr(true)
				updateSecretOptionsModel.ResourceID = core.StringPtr("testString")
				updateSecretOptionsModel.ResourceType = core.StringPtr("testString")
				updateSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				updateSecretOptionsModel.Role = core.StringPtr("testString")
				updateSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				updateSecretOptionsModel.Target = core.StringPtr("testString")
				updateSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateSecret(updateSecretOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSecretOptions model with no property values
				updateSecretOptionsModelNew := new(codeenginev2.UpdateSecretOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateSecret(updateSecretOptionsModelNew)
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
			It(`Invoke UpdateSecret successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateSecretOptions model
				updateSecretOptionsModel := new(codeenginev2.UpdateSecretOptions)
				updateSecretOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SecretName = core.StringPtr("my-secret")
				updateSecretOptionsModel.Name = core.StringPtr("testString")
				updateSecretOptionsModel.BindingSecretRef = core.StringPtr("testString")
				updateSecretOptionsModel.CeComponents = []string{"testString"}
				updateSecretOptionsModel.Data = make(map[string]string)
				updateSecretOptionsModel.Format = core.StringPtr("testString")
				updateSecretOptionsModel.Immutable = core.BoolPtr(true)
				updateSecretOptionsModel.ResourceID = core.StringPtr("testString")
				updateSecretOptionsModel.ResourceType = core.StringPtr("testString")
				updateSecretOptionsModel.ResourcekeyID = core.StringPtr("testString")
				updateSecretOptionsModel.Role = core.StringPtr("testString")
				updateSecretOptionsModel.ServiceidCrn = core.StringPtr("testString")
				updateSecretOptionsModel.Target = core.StringPtr("testString")
				updateSecretOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateSecret(updateSecretOptionsModel)
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
				listAppsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"apps": [{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listAppsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"apps": [{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listAppsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listAppsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listAppsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"apps":[{"ce_managed_domain_mappings":"local+public","created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_protocol":"http1","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"active","type":"Type","version":"1"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"apps":[{"ce_managed_domain_mappings":"local+public","created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_protocol":"http1","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"active","type":"Type","version":"1"}]}`)
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				createAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				createAppOptionsModel.RunArgs = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Version = core.StringPtr("1")
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
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				createAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				createAppOptionsModel.RunArgs = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Version = core.StringPtr("1")
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
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				createAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				createAppOptionsModel.RunArgs = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				createAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				createAppOptionsModel.RunArgs = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateAppOptions model
				createAppOptionsModel := new(codeenginev2.CreateAppOptions)
				createAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.Name = core.StringPtr("my-app")
				createAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				createAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				createAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				createAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				createAppOptionsModel.RunArgs = []string{"testString"}
				createAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createAppOptionsModel.RunCommands = []string{"testString"}
				createAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				createAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				createAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				createAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				createAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				createAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				createAppOptionsModel.Version = core.StringPtr("1")
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
				getAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.AppName = core.StringPtr("my-app")
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
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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
				getAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.AppName = core.StringPtr("my-app")
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
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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
				getAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.AppName = core.StringPtr("my-app")
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
				getAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.AppName = core.StringPtr("my-app")
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
				getAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.AppName = core.StringPtr("my-app")
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
				deleteAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.AppName = core.StringPtr("my-app")
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
				deleteAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.AppName = core.StringPtr("my-app")
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.AppName = core.StringPtr("my-app")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				updateAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				updateAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				updateAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				updateAppOptionsModel.RunArgs = []string{"testString"}
				updateAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateAppOptionsModel.RunCommands = []string{"testString"}
				updateAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				updateAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				updateAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				updateAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				updateAppOptionsModel.Version = core.StringPtr("1")
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.AppName = core.StringPtr("my-app")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				updateAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				updateAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				updateAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				updateAppOptionsModel.RunArgs = []string{"testString"}
				updateAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateAppOptionsModel.RunCommands = []string{"testString"}
				updateAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				updateAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				updateAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				updateAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				updateAppOptionsModel.Version = core.StringPtr("1")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"ce_managed_domain_mappings": "local+public", "created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.AppName = core.StringPtr("my-app")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				updateAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				updateAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				updateAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				updateAppOptionsModel.RunArgs = []string{"testString"}
				updateAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateAppOptionsModel.RunCommands = []string{"testString"}
				updateAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				updateAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				updateAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				updateAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				updateAppOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.AppName = core.StringPtr("my-app")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				updateAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				updateAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				updateAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				updateAppOptionsModel.RunArgs = []string{"testString"}
				updateAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateAppOptionsModel.RunCommands = []string{"testString"}
				updateAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				updateAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				updateAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				updateAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				updateAppOptionsModel.Version = core.StringPtr("1")
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateApp successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateAppOptions model
				updateAppOptionsModel := new(codeenginev2.UpdateAppOptions)
				updateAppOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.AppName = core.StringPtr("my-app")
				updateAppOptionsModel.Name = core.StringPtr("my-app")
				updateAppOptionsModel.CeManagedDomainMappings = core.StringPtr("local+public")
				updateAppOptionsModel.ImagePort = core.Int64Ptr(int64(8080))
				updateAppOptionsModel.ImageProtocol = core.StringPtr("http1")
				updateAppOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateAppOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateAppOptionsModel.RevisionSuffix = core.StringPtr("rev-0001")
				updateAppOptionsModel.RunArgs = []string{"testString"}
				updateAppOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateAppOptionsModel.RunCommands = []string{"testString"}
				updateAppOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateAppOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateAppOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateAppOptionsModel.ScaleConcurrency = core.Int64Ptr(int64(100))
				updateAppOptionsModel.ScaleConcurrencyTarget = core.Int64Ptr(int64(80))
				updateAppOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateAppOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleInitialInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleMaxInstances = core.Int64Ptr(int64(10))
				updateAppOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateAppOptionsModel.ScaleMinInstances = core.Int64Ptr(int64(1))
				updateAppOptionsModel.ScaleRequestTimeout = core.Int64Ptr(int64(300))
				updateAppOptionsModel.Version = core.StringPtr("1")
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
				listAppRevisionsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "revisions": [{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type"}]}`)
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
				listAppRevisionsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "revisions": [{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type"}]}`)
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
				listAppRevisionsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listAppRevisionsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listAppRevisionsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"revisions":[{"created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_protocol":"http1","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"active","type":"Type"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"revisions":[{"created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_port":8080,"image_protocol":"http1","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_concurrency":100,"scale_concurrency_target":80,"scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_initial_instances":1,"scale_max_instances":10,"scale_memory_limit":"4G","scale_min_instances":1,"scale_request_timeout":300,"status":"active","type":"Type"}]}`)
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
				getAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type"}`)
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
				getAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_port": 8080, "image_protocol": "http1", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_concurrency": 100, "scale_concurrency_target": 80, "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_initial_instances": 1, "scale_max_instances": 10, "scale_memory_limit": "4G", "scale_min_instances": 1, "scale_request_timeout": 300, "status": "active", "type": "Type"}`)
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
				getAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
				getAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
				getAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				getAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
				deleteAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				deleteAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
				deleteAppRevisionOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.AppName = core.StringPtr("my-app")
				deleteAppRevisionOptionsModel.RevisionName = core.StringPtr("my-app-001")
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
				listJobsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"jobs": [{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listJobsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
					fmt.Fprintf(res, "%s", `{"jobs": [{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}], "limit": 100, "next": {"href": "Href", "start": "Start"}}`)
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
				listJobsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listJobsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				listJobsOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
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
				nextObject := new(codeenginev2.PaginationListNextMetadata)
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"jobs":[{"created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3,"status":"active","type":"Type","version":"1"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"jobs":[{"created":"2022-09-13T11:41:35+02:00","details":"succeeded","id":"e33b1cv7-7390-4437-a5c2-130d5ccdddc3","image_ref":"icr.io/codeengine/helloworld","image_secret":"my-secret","links":{"mapKey":{"href":"Href","method":"Method"}},"name":"resource-example","reason":"create","run_args":["RunArgs"],"run_as_user":1001,"run_commands":["RunCommands"],"run_env_vars":[{"key":"MY_VARIABLE","name":"SOME","prefix":"PREFIX_","ref":"my-secret","type":"literal","value":"VALUE"}],"run_mode":"daemon","run_service_account":"default","run_volume_mounts":[{"mount_path":"/app","name":"codeengine-mount-b69u90","ref":"my-secret","type":"secret"}],"scale_array_spec":"1-5,7-8,10","scale_cpu_limit":"1","scale_ephemeral_storage_limit":"4G","scale_max_execution_time":7200,"scale_memory_limit":"4G","scale_retry_limit":3,"status":"active","type":"Type","version":"1"}],"limit":1}`)
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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
					ProjectGuid: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.RunArgs = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Version = core.StringPtr("1")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.RunArgs = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Version = core.StringPtr("1")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.RunArgs = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.RunArgs = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(codeenginev2.CreateJobOptions)
				createJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				createJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				createJobOptionsModel.Name = core.StringPtr("my-job")
				createJobOptionsModel.RunArgs = []string{"testString"}
				createJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				createJobOptionsModel.RunCommands = []string{"testString"}
				createJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				createJobOptionsModel.RunMode = core.StringPtr("daemon")
				createJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				createJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				createJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				createJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				createJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				createJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				createJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				createJobOptionsModel.Version = core.StringPtr("1")
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
				getJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.JobName = core.StringPtr("my-job")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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
				getJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.JobName = core.StringPtr("my-job")
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
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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
				getJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.JobName = core.StringPtr("my-job")
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
				getJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.JobName = core.StringPtr("my-job")
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
				getJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.JobName = core.StringPtr("my-job")
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
				deleteJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.JobName = core.StringPtr("my-job")
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
				deleteJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.JobName = core.StringPtr("my-job")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.JobName = core.StringPtr("my-job")
				updateJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.RunArgs = []string{"testString"}
				updateJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateJobOptionsModel.RunCommands = []string{"testString"}
				updateJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateJobOptionsModel.RunMode = core.StringPtr("daemon")
				updateJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				updateJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				updateJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				updateJobOptionsModel.Version = core.StringPtr("1")
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.JobName = core.StringPtr("my-job")
				updateJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.RunArgs = []string{"testString"}
				updateJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateJobOptionsModel.RunCommands = []string{"testString"}
				updateJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateJobOptionsModel.RunMode = core.StringPtr("daemon")
				updateJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				updateJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				updateJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				updateJobOptionsModel.Version = core.StringPtr("1")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "2022-09-13T11:41:35+02:00", "details": "succeeded", "id": "e33b1cv7-7390-4437-a5c2-130d5ccdddc3", "image_ref": "icr.io/codeengine/helloworld", "image_secret": "my-secret", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "name": "resource-example", "reason": "create", "run_args": ["RunArgs"], "run_as_user": 1001, "run_commands": ["RunCommands"], "run_env_vars": [{"key": "MY_VARIABLE", "name": "SOME", "prefix": "PREFIX_", "ref": "my-secret", "type": "literal", "value": "VALUE"}], "run_mode": "daemon", "run_service_account": "default", "run_volume_mounts": [{"mount_path": "/app", "name": "codeengine-mount-b69u90", "ref": "my-secret", "type": "secret"}], "scale_array_spec": "1-5,7-8,10", "scale_cpu_limit": "1", "scale_ephemeral_storage_limit": "4G", "scale_max_execution_time": 7200, "scale_memory_limit": "4G", "scale_retry_limit": 3, "status": "active", "type": "Type", "version": "1"}`)
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.JobName = core.StringPtr("my-job")
				updateJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.RunArgs = []string{"testString"}
				updateJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateJobOptionsModel.RunCommands = []string{"testString"}
				updateJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateJobOptionsModel.RunMode = core.StringPtr("daemon")
				updateJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				updateJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				updateJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				updateJobOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.JobName = core.StringPtr("my-job")
				updateJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.RunArgs = []string{"testString"}
				updateJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateJobOptionsModel.RunCommands = []string{"testString"}
				updateJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateJobOptionsModel.RunMode = core.StringPtr("daemon")
				updateJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				updateJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				updateJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				updateJobOptionsModel.Version = core.StringPtr("1")
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

				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")

				// Construct an instance of the UpdateJobOptions model
				updateJobOptionsModel := new(codeenginev2.UpdateJobOptions)
				updateJobOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.JobName = core.StringPtr("my-job")
				updateJobOptionsModel.ImageRef = core.StringPtr("icr.io/codeengine/helloworld")
				updateJobOptionsModel.ImageSecret = core.StringPtr("my-secret")
				updateJobOptionsModel.Name = core.StringPtr("my-job")
				updateJobOptionsModel.RunArgs = []string{"testString"}
				updateJobOptionsModel.RunAsUser = core.Int64Ptr(int64(1001))
				updateJobOptionsModel.RunCommands = []string{"testString"}
				updateJobOptionsModel.RunEnvVars = []codeenginev2.EnvVar{*envVarModel}
				updateJobOptionsModel.RunMode = core.StringPtr("daemon")
				updateJobOptionsModel.RunServiceAccount = core.StringPtr("default")
				updateJobOptionsModel.RunVolumeMounts = []codeenginev2.VolumeMount{*volumeMountModel}
				updateJobOptionsModel.ScaleArraySpec = core.StringPtr("1-5,7-8,10")
				updateJobOptionsModel.ScaleCpuLimit = core.StringPtr("1")
				updateJobOptionsModel.ScaleEphemeralStorageLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleMaxExecutionTime = core.Int64Ptr(int64(7200))
				updateJobOptionsModel.ScaleMemoryLimit = core.StringPtr("4G")
				updateJobOptionsModel.ScaleRetryLimit = core.Int64Ptr(int64(3))
				updateJobOptionsModel.Version = core.StringPtr("1")
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
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions) - Operation response error`, func() {
		listReclamationsPath := "/reclamations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReclamations with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(codeenginev2.ListReclamationsOptions)
				listReclamationsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listReclamationsOptionsModel.Start = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions)`, func() {
		listReclamationsPath := "/reclamations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "reclamations": [{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}]}`)
				}))
			})
			It(`Invoke ListReclamations successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(codeenginev2.ListReclamationsOptions)
				listReclamationsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listReclamationsOptionsModel.Start = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListReclamationsWithContext(ctx, listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListReclamationsWithContext(ctx, listReclamationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 100, "next": {"href": "Href", "start": "Start"}, "reclamations": [{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}]}`)
				}))
			})
			It(`Invoke ListReclamations successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListReclamations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(codeenginev2.ListReclamationsOptions)
				listReclamationsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listReclamationsOptionsModel.Start = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReclamations with error: Operation request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(codeenginev2.ListReclamationsOptions)
				listReclamationsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listReclamationsOptionsModel.Start = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListReclamations(listReclamationsOptionsModel)
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
			It(`Invoke ListReclamations successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(codeenginev2.ListReclamationsOptions)
				listReclamationsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listReclamationsOptionsModel.Start = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListReclamations(listReclamationsOptionsModel)
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
				responseObject := new(codeenginev2.ReclamationList)
				nextObject := new(codeenginev2.PaginationListNextMetadata)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(codeenginev2.ReclamationList)
	
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
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"reclamations":[{"account_id":"4329073d16d2f3663f74bfa955259139","details":"succeeded","id":"15314cc3-85b4-4338-903f-c28cdee6d005","links":{"mapKey":{"href":"Href","method":"Method"}},"project_id":"15314cc3-85b4-4338-903f-c28cdee6d005","reason":"create","resource_group_id":"b91e849cedb04e7e92bd68c040c672dc","status":"active","target_time":"2022-09-22T17:40:56Z","type":"reclamation/v2"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"reclamations":[{"account_id":"4329073d16d2f3663f74bfa955259139","details":"succeeded","id":"15314cc3-85b4-4338-903f-c28cdee6d005","links":{"mapKey":{"href":"Href","method":"Method"}},"project_id":"15314cc3-85b4-4338-903f-c28cdee6d005","reason":"create","resource_group_id":"b91e849cedb04e7e92bd68c040c672dc","status":"active","target_time":"2022-09-22T17:40:56Z","type":"reclamation/v2"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReclamationsPager.GetNext successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listReclamationsOptionsModel := &codeenginev2.ListReclamationsOptions{
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewReclamationsPager(listReclamationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []codeenginev2.Reclamation
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReclamationsPager.GetAll successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				listReclamationsOptionsModel := &codeenginev2.ListReclamationsOptions{
					Limit: core.Int64Ptr(int64(100)),
				}

				pager, err := codeEngineService.NewReclamationsPager(listReclamationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetReclamation(getReclamationOptions *GetReclamationOptions) - Operation response error`, func() {
		getReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReclamation with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationOptions model
				getReclamationOptionsModel := new(codeenginev2.GetReclamationOptions)
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetReclamation(getReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetReclamation(getReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReclamation(getReclamationOptions *GetReclamationOptions)`, func() {
		getReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke GetReclamation successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetReclamationOptions model
				getReclamationOptionsModel := new(codeenginev2.GetReclamationOptions)
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetReclamationWithContext(ctx, getReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetReclamation(getReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetReclamationWithContext(ctx, getReclamationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke GetReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetReclamation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReclamationOptions model
				getReclamationOptionsModel := new(codeenginev2.GetReclamationOptions)
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetReclamation(getReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReclamation with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationOptions model
				getReclamationOptionsModel := new(codeenginev2.GetReclamationOptions)
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetReclamation(getReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReclamationOptions model with no property values
				getReclamationOptionsModelNew := new(codeenginev2.GetReclamationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetReclamation(getReclamationOptionsModelNew)
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
			It(`Invoke GetReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationOptions model
				getReclamationOptionsModel := new(codeenginev2.GetReclamationOptions)
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetReclamation(getReclamationOptionsModel)
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
	Describe(`ReclaimReclamation(reclaimReclamationOptions *ReclaimReclamationOptions) - Operation response error`, func() {
		reclaimReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005/reclaim"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReclaimReclamation with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationOptions model
				reclaimReclamationOptionsModel := new(codeenginev2.ReclaimReclamationOptions)
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReclaimReclamation(reclaimReclamationOptions *ReclaimReclamationOptions)`, func() {
		reclaimReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005/reclaim"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke ReclaimReclamation successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ReclaimReclamationOptions model
				reclaimReclamationOptionsModel := new(codeenginev2.ReclaimReclamationOptions)
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ReclaimReclamationWithContext(ctx, reclaimReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ReclaimReclamationWithContext(ctx, reclaimReclamationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke ReclaimReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ReclaimReclamation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReclaimReclamationOptions model
				reclaimReclamationOptionsModel := new(codeenginev2.ReclaimReclamationOptions)
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReclaimReclamation with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationOptions model
				reclaimReclamationOptionsModel := new(codeenginev2.ReclaimReclamationOptions)
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReclaimReclamationOptions model with no property values
				reclaimReclamationOptionsModelNew := new(codeenginev2.ReclaimReclamationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModelNew)
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
			It(`Invoke ReclaimReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationOptions model
				reclaimReclamationOptionsModel := new(codeenginev2.ReclaimReclamationOptions)
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ReclaimReclamation(reclaimReclamationOptionsModel)
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
	Describe(`RestoreReclamation(restoreReclamationOptions *RestoreReclamationOptions) - Operation response error`, func() {
		restoreReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005/restore"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RestoreReclamation with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationOptions model
				restoreReclamationOptionsModel := new(codeenginev2.RestoreReclamationOptions)
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreReclamation(restoreReclamationOptions *RestoreReclamationOptions)`, func() {
		restoreReclamationPath := "/reclamations/15314cc3-85b4-4338-903f-c28cdee6d005/restore"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke RestoreReclamation successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the RestoreReclamationOptions model
				restoreReclamationOptionsModel := new(codeenginev2.RestoreReclamationOptions)
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.RestoreReclamationWithContext(ctx, restoreReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.RestoreReclamationWithContext(ctx, restoreReclamationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "4329073d16d2f3663f74bfa955259139", "details": "succeeded", "id": "15314cc3-85b4-4338-903f-c28cdee6d005", "links": {"mapKey": {"href": "Href", "method": "Method"}}, "project_id": "15314cc3-85b4-4338-903f-c28cdee6d005", "reason": "create", "resource_group_id": "b91e849cedb04e7e92bd68c040c672dc", "status": "active", "target_time": "2022-09-22T17:40:56Z", "type": "reclamation/v2"}`)
				}))
			})
			It(`Invoke RestoreReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.RestoreReclamation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestoreReclamationOptions model
				restoreReclamationOptionsModel := new(codeenginev2.RestoreReclamationOptions)
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RestoreReclamation with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationOptions model
				restoreReclamationOptionsModel := new(codeenginev2.RestoreReclamationOptions)
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreReclamationOptions model with no property values
				restoreReclamationOptionsModelNew := new(codeenginev2.RestoreReclamationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.RestoreReclamation(restoreReclamationOptionsModelNew)
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
			It(`Invoke RestoreReclamation successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationOptions model
				restoreReclamationOptionsModel := new(codeenginev2.RestoreReclamationOptions)
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.RestoreReclamation(restoreReclamationOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			codeEngineService, _ := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
				URL:           "http://codeenginev2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateAppOptions successfully`, func() {
				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				Expect(envVarModel).ToNot(BeNil())
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")
				Expect(envVarModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				Expect(volumeMountModel).ToNot(BeNil())
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")
				Expect(volumeMountModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the CreateAppOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createAppOptionsName := "my-app"
				createAppOptionsModel := codeEngineService.NewCreateAppOptions(projectGuid, createAppOptionsName)
				createAppOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createAppOptionsModel.SetName("my-app")
				createAppOptionsModel.SetCeManagedDomainMappings("local+public")
				createAppOptionsModel.SetImagePort(int64(8080))
				createAppOptionsModel.SetImageProtocol("http1")
				createAppOptionsModel.SetImageRef("icr.io/codeengine/helloworld")
				createAppOptionsModel.SetImageSecret("my-secret")
				createAppOptionsModel.SetRevisionSuffix("rev-0001")
				createAppOptionsModel.SetRunArgs([]string{"testString"})
				createAppOptionsModel.SetRunAsUser(int64(1001))
				createAppOptionsModel.SetRunCommands([]string{"testString"})
				createAppOptionsModel.SetRunEnvVars([]codeenginev2.EnvVar{*envVarModel})
				createAppOptionsModel.SetRunServiceAccount("default")
				createAppOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMount{*volumeMountModel})
				createAppOptionsModel.SetScaleConcurrency(int64(100))
				createAppOptionsModel.SetScaleConcurrencyTarget(int64(80))
				createAppOptionsModel.SetScaleCpuLimit("1")
				createAppOptionsModel.SetScaleEphemeralStorageLimit("4G")
				createAppOptionsModel.SetScaleInitialInstances(int64(1))
				createAppOptionsModel.SetScaleMaxInstances(int64(10))
				createAppOptionsModel.SetScaleMemoryLimit("4G")
				createAppOptionsModel.SetScaleMinInstances(int64(1))
				createAppOptionsModel.SetScaleRequestTimeout(int64(300))
				createAppOptionsModel.SetVersion("1")
				createAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAppOptionsModel).ToNot(BeNil())
				Expect(createAppOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(createAppOptionsModel.CeManagedDomainMappings).To(Equal(core.StringPtr("local+public")))
				Expect(createAppOptionsModel.ImagePort).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(createAppOptionsModel.ImageProtocol).To(Equal(core.StringPtr("http1")))
				Expect(createAppOptionsModel.ImageRef).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(createAppOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(createAppOptionsModel.RevisionSuffix).To(Equal(core.StringPtr("rev-0001")))
				Expect(createAppOptionsModel.RunArgs).To(Equal([]string{"testString"}))
				Expect(createAppOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(createAppOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(createAppOptionsModel.RunEnvVars).To(Equal([]codeenginev2.EnvVar{*envVarModel}))
				Expect(createAppOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createAppOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMount{*volumeMountModel}))
				Expect(createAppOptionsModel.ScaleConcurrency).To(Equal(core.Int64Ptr(int64(100))))
				Expect(createAppOptionsModel.ScaleConcurrencyTarget).To(Equal(core.Int64Ptr(int64(80))))
				Expect(createAppOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(createAppOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(createAppOptionsModel.ScaleInitialInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createAppOptionsModel.ScaleMaxInstances).To(Equal(core.Int64Ptr(int64(10))))
				Expect(createAppOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(createAppOptionsModel.ScaleMinInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createAppOptionsModel.ScaleRequestTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(createAppOptionsModel.Version).To(Equal(core.StringPtr("1")))
				Expect(createAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBuildOptions successfully`, func() {
				// Construct an instance of the CreateBuildOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createBuildOptionsName := "my-build"
				createBuildOptionsModel := codeEngineService.NewCreateBuildOptions(projectGuid, createBuildOptionsName)
				createBuildOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildOptionsModel.SetName("my-build")
				createBuildOptionsModel.SetCeOwnerReference("testString")
				createBuildOptionsModel.SetOutputImage("stg.icr.io/icr_namespace/image-name")
				createBuildOptionsModel.SetOutputSecret("ce-default-icr-us-south")
				createBuildOptionsModel.SetSourceContextDir("testString")
				createBuildOptionsModel.SetSourceRevision("main")
				createBuildOptionsModel.SetSourceSecret("testString")
				createBuildOptionsModel.SetSourceType("git")
				createBuildOptionsModel.SetSourceURL("https://github.com/IBM/CodeEngine")
				createBuildOptionsModel.SetStrategyName("dockerfile")
				createBuildOptionsModel.SetStrategySize("medium")
				createBuildOptionsModel.SetStrategySpecFile("Dockerfile")
				createBuildOptionsModel.SetTimeout(int64(600))
				createBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBuildOptionsModel).ToNot(BeNil())
				Expect(createBuildOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(createBuildOptionsModel.CeOwnerReference).To(Equal(core.StringPtr("testString")))
				Expect(createBuildOptionsModel.OutputImage).To(Equal(core.StringPtr("stg.icr.io/icr_namespace/image-name")))
				Expect(createBuildOptionsModel.OutputSecret).To(Equal(core.StringPtr("ce-default-icr-us-south")))
				Expect(createBuildOptionsModel.SourceContextDir).To(Equal(core.StringPtr("testString")))
				Expect(createBuildOptionsModel.SourceRevision).To(Equal(core.StringPtr("main")))
				Expect(createBuildOptionsModel.SourceSecret).To(Equal(core.StringPtr("testString")))
				Expect(createBuildOptionsModel.SourceType).To(Equal(core.StringPtr("git")))
				Expect(createBuildOptionsModel.SourceURL).To(Equal(core.StringPtr("https://github.com/IBM/CodeEngine")))
				Expect(createBuildOptionsModel.StrategyName).To(Equal(core.StringPtr("dockerfile")))
				Expect(createBuildOptionsModel.StrategySize).To(Equal(core.StringPtr("medium")))
				Expect(createBuildOptionsModel.StrategySpecFile).To(Equal(core.StringPtr("Dockerfile")))
				Expect(createBuildOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(600))))
				Expect(createBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBuildrunOptions successfully`, func() {
				// Construct an instance of the CreateBuildrunOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createBuildrunOptionsName := "testString"
				createBuildrunOptionsModel := codeEngineService.NewCreateBuildrunOptions(projectGuid, createBuildrunOptionsName)
				createBuildrunOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createBuildrunOptionsModel.SetName("testString")
				createBuildrunOptionsModel.SetAppRevision("testString")
				createBuildrunOptionsModel.SetBuild("testString")
				createBuildrunOptionsModel.SetCeOwnerReference("testString")
				createBuildrunOptionsModel.SetOutputImage("stg.icr.io/icr_namespace/image-name")
				createBuildrunOptionsModel.SetOutputSecret("ce-default-icr-us-south")
				createBuildrunOptionsModel.SetServiceAccount("testString")
				createBuildrunOptionsModel.SetSourceContextDir("testString")
				createBuildrunOptionsModel.SetSourceRevision("main")
				createBuildrunOptionsModel.SetSourceSecret("testString")
				createBuildrunOptionsModel.SetSourceType("git")
				createBuildrunOptionsModel.SetSourceURL("https://github.com/IBM/CodeEngine")
				createBuildrunOptionsModel.SetStrategyName("dockerfile")
				createBuildrunOptionsModel.SetStrategySize("medium")
				createBuildrunOptionsModel.SetStrategySpecFile("Dockerfile")
				createBuildrunOptionsModel.SetTimeout(int64(600))
				createBuildrunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBuildrunOptionsModel).ToNot(BeNil())
				Expect(createBuildrunOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createBuildrunOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.AppRevision).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.Build).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.CeOwnerReference).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.OutputImage).To(Equal(core.StringPtr("stg.icr.io/icr_namespace/image-name")))
				Expect(createBuildrunOptionsModel.OutputSecret).To(Equal(core.StringPtr("ce-default-icr-us-south")))
				Expect(createBuildrunOptionsModel.ServiceAccount).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.SourceContextDir).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.SourceRevision).To(Equal(core.StringPtr("main")))
				Expect(createBuildrunOptionsModel.SourceSecret).To(Equal(core.StringPtr("testString")))
				Expect(createBuildrunOptionsModel.SourceType).To(Equal(core.StringPtr("git")))
				Expect(createBuildrunOptionsModel.SourceURL).To(Equal(core.StringPtr("https://github.com/IBM/CodeEngine")))
				Expect(createBuildrunOptionsModel.StrategyName).To(Equal(core.StringPtr("dockerfile")))
				Expect(createBuildrunOptionsModel.StrategySize).To(Equal(core.StringPtr("medium")))
				Expect(createBuildrunOptionsModel.StrategySpecFile).To(Equal(core.StringPtr("Dockerfile")))
				Expect(createBuildrunOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(600))))
				Expect(createBuildrunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateConfigmapOptions successfully`, func() {
				// Construct an instance of the CreateConfigmapOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createConfigmapOptionsName := "my-configmap"
				createConfigmapOptionsModel := codeEngineService.NewCreateConfigmapOptions(projectGuid, createConfigmapOptionsName)
				createConfigmapOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createConfigmapOptionsModel.SetName("my-configmap")
				createConfigmapOptionsModel.SetData(make(map[string]string))
				createConfigmapOptionsModel.SetImmutable(false)
				createConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigmapOptionsModel).ToNot(BeNil())
				Expect(createConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createConfigmapOptionsModel.Name).To(Equal(core.StringPtr("my-configmap")))
				Expect(createConfigmapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createConfigmapOptionsModel.Immutable).To(Equal(core.BoolPtr(false)))
				Expect(createConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateJobOptions successfully`, func() {
				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				Expect(envVarModel).ToNot(BeNil())
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")
				Expect(envVarModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				Expect(volumeMountModel).ToNot(BeNil())
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")
				Expect(volumeMountModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the CreateJobOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createJobOptionsModel := codeEngineService.NewCreateJobOptions(projectGuid)
				createJobOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createJobOptionsModel.SetImageRef("icr.io/codeengine/helloworld")
				createJobOptionsModel.SetImageSecret("my-secret")
				createJobOptionsModel.SetName("my-job")
				createJobOptionsModel.SetRunArgs([]string{"testString"})
				createJobOptionsModel.SetRunAsUser(int64(1001))
				createJobOptionsModel.SetRunCommands([]string{"testString"})
				createJobOptionsModel.SetRunEnvVars([]codeenginev2.EnvVar{*envVarModel})
				createJobOptionsModel.SetRunMode("daemon")
				createJobOptionsModel.SetRunServiceAccount("default")
				createJobOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMount{*volumeMountModel})
				createJobOptionsModel.SetScaleArraySpec("1-5,7-8,10")
				createJobOptionsModel.SetScaleCpuLimit("1")
				createJobOptionsModel.SetScaleEphemeralStorageLimit("4G")
				createJobOptionsModel.SetScaleMaxExecutionTime(int64(7200))
				createJobOptionsModel.SetScaleMemoryLimit("4G")
				createJobOptionsModel.SetScaleRetryLimit(int64(3))
				createJobOptionsModel.SetVersion("1")
				createJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createJobOptionsModel).ToNot(BeNil())
				Expect(createJobOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createJobOptionsModel.ImageRef).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(createJobOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(createJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(createJobOptionsModel.RunArgs).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(createJobOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.RunEnvVars).To(Equal([]codeenginev2.EnvVar{*envVarModel}))
				Expect(createJobOptionsModel.RunMode).To(Equal(core.StringPtr("daemon")))
				Expect(createJobOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(createJobOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMount{*volumeMountModel}))
				Expect(createJobOptionsModel.ScaleArraySpec).To(Equal(core.StringPtr("1-5,7-8,10")))
				Expect(createJobOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(createJobOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobOptionsModel.ScaleMaxExecutionTime).To(Equal(core.Int64Ptr(int64(7200))))
				Expect(createJobOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(createJobOptionsModel.ScaleRetryLimit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(createJobOptionsModel.Version).To(Equal(core.StringPtr("1")))
				Expect(createJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := codeEngineService.NewCreateProjectOptions()
				createProjectOptionsModel.SetName("my-project")
				createProjectOptionsModel.SetRegion("us-east")
				createProjectOptionsModel.SetResourceGroupID("b91e849cedb04e7e92bd68c040c672dc")
				createProjectOptionsModel.SetTags([]string{"testString"})
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("my-project")))
				Expect(createProjectOptionsModel.Region).To(Equal(core.StringPtr("us-east")))
				Expect(createProjectOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("b91e849cedb04e7e92bd68c040c672dc")))
				Expect(createProjectOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSecretOptions successfully`, func() {
				// Construct an instance of the CreateSecretOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				createSecretOptionsName := "testString"
				createSecretOptionsModel := codeEngineService.NewCreateSecretOptions(projectGuid, createSecretOptionsName)
				createSecretOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				createSecretOptionsModel.SetName("testString")
				createSecretOptionsModel.SetBindingSecretRef("testString")
				createSecretOptionsModel.SetCeComponents([]string{"testString"})
				createSecretOptionsModel.SetData(make(map[string]string))
				createSecretOptionsModel.SetFormat("testString")
				createSecretOptionsModel.SetImmutable(true)
				createSecretOptionsModel.SetResourceID("testString")
				createSecretOptionsModel.SetResourceType("testString")
				createSecretOptionsModel.SetResourcekeyID("testString")
				createSecretOptionsModel.SetRole("testString")
				createSecretOptionsModel.SetServiceidCrn("testString")
				createSecretOptionsModel.SetTarget("testString")
				createSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSecretOptionsModel).ToNot(BeNil())
				Expect(createSecretOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(createSecretOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.BindingSecretRef).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.CeComponents).To(Equal([]string{"testString"}))
				Expect(createSecretOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createSecretOptionsModel.Format).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(createSecretOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.ResourcekeyID).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.Role).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.ServiceidCrn).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(createSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAppOptions successfully`, func() {
				// Construct an instance of the DeleteAppOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				deleteAppOptionsModel := codeEngineService.NewDeleteAppOptions(projectGuid, appName)
				deleteAppOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppOptionsModel.SetAppName("my-app")
				deleteAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAppOptionsModel).ToNot(BeNil())
				Expect(deleteAppOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteAppOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(deleteAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAppRevisionOptions successfully`, func() {
				// Construct an instance of the DeleteAppRevisionOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				revisionName := "my-app-001"
				deleteAppRevisionOptionsModel := codeEngineService.NewDeleteAppRevisionOptions(projectGuid, appName, revisionName)
				deleteAppRevisionOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteAppRevisionOptionsModel.SetAppName("my-app")
				deleteAppRevisionOptionsModel.SetRevisionName("my-app-001")
				deleteAppRevisionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAppRevisionOptionsModel).ToNot(BeNil())
				Expect(deleteAppRevisionOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteAppRevisionOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(deleteAppRevisionOptionsModel.RevisionName).To(Equal(core.StringPtr("my-app-001")))
				Expect(deleteAppRevisionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBuildOptions successfully`, func() {
				// Construct an instance of the DeleteBuildOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				buildName := "my-build"
				deleteBuildOptionsModel := codeEngineService.NewDeleteBuildOptions(projectGuid, buildName)
				deleteBuildOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildOptionsModel.SetBuildName("my-build")
				deleteBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBuildOptionsModel).ToNot(BeNil())
				Expect(deleteBuildOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteBuildOptionsModel.BuildName).To(Equal(core.StringPtr("my-build")))
				Expect(deleteBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBuildrunOptions successfully`, func() {
				// Construct an instance of the DeleteBuildrunOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				buildRunName := "my-build-run"
				deleteBuildrunOptionsModel := codeEngineService.NewDeleteBuildrunOptions(projectGuid, buildRunName)
				deleteBuildrunOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteBuildrunOptionsModel.SetBuildRunName("my-build-run")
				deleteBuildrunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBuildrunOptionsModel).ToNot(BeNil())
				Expect(deleteBuildrunOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteBuildrunOptionsModel.BuildRunName).To(Equal(core.StringPtr("my-build-run")))
				Expect(deleteBuildrunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigmapOptions successfully`, func() {
				// Construct an instance of the DeleteConfigmapOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				configMapName := "my-config-map"
				deleteConfigmapOptionsModel := codeEngineService.NewDeleteConfigmapOptions(projectGuid, configMapName)
				deleteConfigmapOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteConfigmapOptionsModel.SetConfigMapName("my-config-map")
				deleteConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigmapOptionsModel).ToNot(BeNil())
				Expect(deleteConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteConfigmapOptionsModel.ConfigMapName).To(Equal(core.StringPtr("my-config-map")))
				Expect(deleteConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteJobOptions successfully`, func() {
				// Construct an instance of the DeleteJobOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				jobName := "my-job"
				deleteJobOptionsModel := codeEngineService.NewDeleteJobOptions(projectGuid, jobName)
				deleteJobOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteJobOptionsModel.SetJobName("my-job")
				deleteJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteJobOptionsModel).ToNot(BeNil())
				Expect(deleteJobOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteJobOptionsModel.JobName).To(Equal(core.StringPtr("my-job")))
				Expect(deleteJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				deleteProjectOptionsModel := codeEngineService.NewDeleteProjectOptions(projectGuid)
				deleteProjectOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSecretOptions successfully`, func() {
				// Construct an instance of the DeleteSecretOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				secretName := "my-secret"
				deleteSecretOptionsModel := codeEngineService.NewDeleteSecretOptions(projectGuid, secretName)
				deleteSecretOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				deleteSecretOptionsModel.SetSecretName("my-secret")
				deleteSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSecretOptionsModel).ToNot(BeNil())
				Expect(deleteSecretOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(deleteSecretOptionsModel.SecretName).To(Equal(core.StringPtr("my-secret")))
				Expect(deleteSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAppOptions successfully`, func() {
				// Construct an instance of the GetAppOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				getAppOptionsModel := codeEngineService.NewGetAppOptions(projectGuid, appName)
				getAppOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppOptionsModel.SetAppName("my-app")
				getAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAppOptionsModel).ToNot(BeNil())
				Expect(getAppOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getAppOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(getAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAppRevisionOptions successfully`, func() {
				// Construct an instance of the GetAppRevisionOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				revisionName := "my-app-001"
				getAppRevisionOptionsModel := codeEngineService.NewGetAppRevisionOptions(projectGuid, appName, revisionName)
				getAppRevisionOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getAppRevisionOptionsModel.SetAppName("my-app")
				getAppRevisionOptionsModel.SetRevisionName("my-app-001")
				getAppRevisionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAppRevisionOptionsModel).ToNot(BeNil())
				Expect(getAppRevisionOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getAppRevisionOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(getAppRevisionOptionsModel.RevisionName).To(Equal(core.StringPtr("my-app-001")))
				Expect(getAppRevisionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBuildOptions successfully`, func() {
				// Construct an instance of the GetBuildOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				buildName := "my-build"
				getBuildOptionsModel := codeEngineService.NewGetBuildOptions(projectGuid, buildName)
				getBuildOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildOptionsModel.SetBuildName("my-build")
				getBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBuildOptionsModel).ToNot(BeNil())
				Expect(getBuildOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getBuildOptionsModel.BuildName).To(Equal(core.StringPtr("my-build")))
				Expect(getBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBuildrunOptions successfully`, func() {
				// Construct an instance of the GetBuildrunOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				buildRunName := "my-build-run"
				getBuildrunOptionsModel := codeEngineService.NewGetBuildrunOptions(projectGuid, buildRunName)
				getBuildrunOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getBuildrunOptionsModel.SetBuildRunName("my-build-run")
				getBuildrunOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBuildrunOptionsModel).ToNot(BeNil())
				Expect(getBuildrunOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getBuildrunOptionsModel.BuildRunName).To(Equal(core.StringPtr("my-build-run")))
				Expect(getBuildrunOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigmapOptions successfully`, func() {
				// Construct an instance of the GetConfigmapOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				configMapName := "my-config-map"
				getConfigmapOptionsModel := codeEngineService.NewGetConfigmapOptions(projectGuid, configMapName)
				getConfigmapOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getConfigmapOptionsModel.SetConfigMapName("my-config-map")
				getConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigmapOptionsModel).ToNot(BeNil())
				Expect(getConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getConfigmapOptionsModel.ConfigMapName).To(Equal(core.StringPtr("my-config-map")))
				Expect(getConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetJobOptions successfully`, func() {
				// Construct an instance of the GetJobOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				jobName := "my-job"
				getJobOptionsModel := codeEngineService.NewGetJobOptions(projectGuid, jobName)
				getJobOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getJobOptionsModel.SetJobName("my-job")
				getJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getJobOptionsModel).ToNot(BeNil())
				Expect(getJobOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getJobOptionsModel.JobName).To(Equal(core.StringPtr("my-job")))
				Expect(getJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				getProjectOptionsModel := codeEngineService.NewGetProjectOptions(projectGuid)
				getProjectOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReclamationOptions successfully`, func() {
				// Construct an instance of the GetReclamationOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				getReclamationOptionsModel := codeEngineService.NewGetReclamationOptions(projectGuid)
				getReclamationOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReclamationOptionsModel).ToNot(BeNil())
				Expect(getReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSecretOptions successfully`, func() {
				// Construct an instance of the GetSecretOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				secretName := "my-secret"
				getSecretOptionsModel := codeEngineService.NewGetSecretOptions(projectGuid, secretName)
				getSecretOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				getSecretOptionsModel.SetSecretName("my-secret")
				getSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSecretOptionsModel).ToNot(BeNil())
				Expect(getSecretOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(getSecretOptionsModel.SecretName).To(Equal(core.StringPtr("my-secret")))
				Expect(getSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAppRevisionsOptions successfully`, func() {
				// Construct an instance of the ListAppRevisionsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				listAppRevisionsOptionsModel := codeEngineService.NewListAppRevisionsOptions(projectGuid, appName)
				listAppRevisionsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppRevisionsOptionsModel.SetAppName("my-app")
				listAppRevisionsOptionsModel.SetLimit(int64(100))
				listAppRevisionsOptionsModel.SetStart("testString")
				listAppRevisionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAppRevisionsOptionsModel).ToNot(BeNil())
				Expect(listAppRevisionsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listAppRevisionsOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(listAppRevisionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listAppRevisionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAppRevisionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAppsOptions successfully`, func() {
				// Construct an instance of the ListAppsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listAppsOptionsModel := codeEngineService.NewListAppsOptions(projectGuid)
				listAppsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listAppsOptionsModel.SetLimit(int64(100))
				listAppsOptionsModel.SetStart("testString")
				listAppsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAppsOptionsModel).ToNot(BeNil())
				Expect(listAppsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listAppsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listAppsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAppsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBuildrunsOptions successfully`, func() {
				// Construct an instance of the ListBuildrunsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listBuildrunsOptionsModel := codeEngineService.NewListBuildrunsOptions(projectGuid)
				listBuildrunsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildrunsOptionsModel.SetLimit(int64(100))
				listBuildrunsOptionsModel.SetStart("testString")
				listBuildrunsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBuildrunsOptionsModel).ToNot(BeNil())
				Expect(listBuildrunsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listBuildrunsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listBuildrunsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBuildrunsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBuildsOptions successfully`, func() {
				// Construct an instance of the ListBuildsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listBuildsOptionsModel := codeEngineService.NewListBuildsOptions(projectGuid)
				listBuildsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listBuildsOptionsModel.SetLimit(int64(100))
				listBuildsOptionsModel.SetStart("testString")
				listBuildsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBuildsOptionsModel).ToNot(BeNil())
				Expect(listBuildsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listBuildsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listBuildsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBuildsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigmapsOptions successfully`, func() {
				// Construct an instance of the ListConfigmapsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listConfigmapsOptionsModel := codeEngineService.NewListConfigmapsOptions(projectGuid)
				listConfigmapsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listConfigmapsOptionsModel.SetLimit(int64(100))
				listConfigmapsOptionsModel.SetStart("testString")
				listConfigmapsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigmapsOptionsModel).ToNot(BeNil())
				Expect(listConfigmapsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listConfigmapsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listConfigmapsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listConfigmapsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobsOptions successfully`, func() {
				// Construct an instance of the ListJobsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listJobsOptionsModel := codeEngineService.NewListJobsOptions(projectGuid)
				listJobsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listJobsOptionsModel.SetLimit(int64(100))
				listJobsOptionsModel.SetStart("testString")
				listJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobsOptionsModel).ToNot(BeNil())
				Expect(listJobsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
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
			It(`Invoke NewListReclamationsOptions successfully`, func() {
				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := codeEngineService.NewListReclamationsOptions()
				listReclamationsOptionsModel.SetLimit(int64(100))
				listReclamationsOptionsModel.SetStart("testString")
				listReclamationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReclamationsOptionsModel).ToNot(BeNil())
				Expect(listReclamationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listReclamationsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSecretsOptions successfully`, func() {
				// Construct an instance of the ListSecretsOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				listSecretsOptionsModel := codeEngineService.NewListSecretsOptions(projectGuid)
				listSecretsOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				listSecretsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSecretsOptionsModel).ToNot(BeNil())
				Expect(listSecretsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(listSecretsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReclaimReclamationOptions successfully`, func() {
				// Construct an instance of the ReclaimReclamationOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				reclaimReclamationOptionsModel := codeEngineService.NewReclaimReclamationOptions(projectGuid)
				reclaimReclamationOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				reclaimReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(reclaimReclamationOptionsModel).ToNot(BeNil())
				Expect(reclaimReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(reclaimReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreReclamationOptions successfully`, func() {
				// Construct an instance of the RestoreReclamationOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				restoreReclamationOptionsModel := codeEngineService.NewRestoreReclamationOptions(projectGuid)
				restoreReclamationOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				restoreReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreReclamationOptionsModel).ToNot(BeNil())
				Expect(restoreReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(restoreReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAppOptions successfully`, func() {
				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				Expect(envVarModel).ToNot(BeNil())
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")
				Expect(envVarModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				Expect(volumeMountModel).ToNot(BeNil())
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")
				Expect(volumeMountModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the UpdateAppOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				appName := "my-app"
				updateAppOptionsName := "my-app"
				updateAppOptionsModel := codeEngineService.NewUpdateAppOptions(projectGuid, appName, updateAppOptionsName)
				updateAppOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateAppOptionsModel.SetAppName("my-app")
				updateAppOptionsModel.SetName("my-app")
				updateAppOptionsModel.SetCeManagedDomainMappings("local+public")
				updateAppOptionsModel.SetImagePort(int64(8080))
				updateAppOptionsModel.SetImageProtocol("http1")
				updateAppOptionsModel.SetImageRef("icr.io/codeengine/helloworld")
				updateAppOptionsModel.SetImageSecret("my-secret")
				updateAppOptionsModel.SetRevisionSuffix("rev-0001")
				updateAppOptionsModel.SetRunArgs([]string{"testString"})
				updateAppOptionsModel.SetRunAsUser(int64(1001))
				updateAppOptionsModel.SetRunCommands([]string{"testString"})
				updateAppOptionsModel.SetRunEnvVars([]codeenginev2.EnvVar{*envVarModel})
				updateAppOptionsModel.SetRunServiceAccount("default")
				updateAppOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMount{*volumeMountModel})
				updateAppOptionsModel.SetScaleConcurrency(int64(100))
				updateAppOptionsModel.SetScaleConcurrencyTarget(int64(80))
				updateAppOptionsModel.SetScaleCpuLimit("1")
				updateAppOptionsModel.SetScaleEphemeralStorageLimit("4G")
				updateAppOptionsModel.SetScaleInitialInstances(int64(1))
				updateAppOptionsModel.SetScaleMaxInstances(int64(10))
				updateAppOptionsModel.SetScaleMemoryLimit("4G")
				updateAppOptionsModel.SetScaleMinInstances(int64(1))
				updateAppOptionsModel.SetScaleRequestTimeout(int64(300))
				updateAppOptionsModel.SetVersion("1")
				updateAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAppOptionsModel).ToNot(BeNil())
				Expect(updateAppOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateAppOptionsModel.AppName).To(Equal(core.StringPtr("my-app")))
				Expect(updateAppOptionsModel.Name).To(Equal(core.StringPtr("my-app")))
				Expect(updateAppOptionsModel.CeManagedDomainMappings).To(Equal(core.StringPtr("local+public")))
				Expect(updateAppOptionsModel.ImagePort).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(updateAppOptionsModel.ImageProtocol).To(Equal(core.StringPtr("http1")))
				Expect(updateAppOptionsModel.ImageRef).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(updateAppOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(updateAppOptionsModel.RevisionSuffix).To(Equal(core.StringPtr("rev-0001")))
				Expect(updateAppOptionsModel.RunArgs).To(Equal([]string{"testString"}))
				Expect(updateAppOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(updateAppOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(updateAppOptionsModel.RunEnvVars).To(Equal([]codeenginev2.EnvVar{*envVarModel}))
				Expect(updateAppOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(updateAppOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMount{*volumeMountModel}))
				Expect(updateAppOptionsModel.ScaleConcurrency).To(Equal(core.Int64Ptr(int64(100))))
				Expect(updateAppOptionsModel.ScaleConcurrencyTarget).To(Equal(core.Int64Ptr(int64(80))))
				Expect(updateAppOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(updateAppOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(updateAppOptionsModel.ScaleInitialInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updateAppOptionsModel.ScaleMaxInstances).To(Equal(core.Int64Ptr(int64(10))))
				Expect(updateAppOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(updateAppOptionsModel.ScaleMinInstances).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updateAppOptionsModel.ScaleRequestTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(updateAppOptionsModel.Version).To(Equal(core.StringPtr("1")))
				Expect(updateAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBuildOptions successfully`, func() {
				// Construct an instance of the UpdateBuildOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				buildName := "my-build"
				updateBuildOptionsName := "my-build"
				updateBuildOptionsModel := codeEngineService.NewUpdateBuildOptions(projectGuid, buildName, updateBuildOptionsName)
				updateBuildOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateBuildOptionsModel.SetBuildName("my-build")
				updateBuildOptionsModel.SetName("my-build")
				updateBuildOptionsModel.SetCeOwnerReference("testString")
				updateBuildOptionsModel.SetOutputImage("stg.icr.io/icr_namespace/image-name")
				updateBuildOptionsModel.SetOutputSecret("ce-default-icr-us-south")
				updateBuildOptionsModel.SetSourceContextDir("testString")
				updateBuildOptionsModel.SetSourceRevision("main")
				updateBuildOptionsModel.SetSourceSecret("testString")
				updateBuildOptionsModel.SetSourceType("git")
				updateBuildOptionsModel.SetSourceURL("https://github.com/IBM/CodeEngine")
				updateBuildOptionsModel.SetStrategyName("dockerfile")
				updateBuildOptionsModel.SetStrategySize("medium")
				updateBuildOptionsModel.SetStrategySpecFile("Dockerfile")
				updateBuildOptionsModel.SetTimeout(int64(600))
				updateBuildOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBuildOptionsModel).ToNot(BeNil())
				Expect(updateBuildOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateBuildOptionsModel.BuildName).To(Equal(core.StringPtr("my-build")))
				Expect(updateBuildOptionsModel.Name).To(Equal(core.StringPtr("my-build")))
				Expect(updateBuildOptionsModel.CeOwnerReference).To(Equal(core.StringPtr("testString")))
				Expect(updateBuildOptionsModel.OutputImage).To(Equal(core.StringPtr("stg.icr.io/icr_namespace/image-name")))
				Expect(updateBuildOptionsModel.OutputSecret).To(Equal(core.StringPtr("ce-default-icr-us-south")))
				Expect(updateBuildOptionsModel.SourceContextDir).To(Equal(core.StringPtr("testString")))
				Expect(updateBuildOptionsModel.SourceRevision).To(Equal(core.StringPtr("main")))
				Expect(updateBuildOptionsModel.SourceSecret).To(Equal(core.StringPtr("testString")))
				Expect(updateBuildOptionsModel.SourceType).To(Equal(core.StringPtr("git")))
				Expect(updateBuildOptionsModel.SourceURL).To(Equal(core.StringPtr("https://github.com/IBM/CodeEngine")))
				Expect(updateBuildOptionsModel.StrategyName).To(Equal(core.StringPtr("dockerfile")))
				Expect(updateBuildOptionsModel.StrategySize).To(Equal(core.StringPtr("medium")))
				Expect(updateBuildOptionsModel.StrategySpecFile).To(Equal(core.StringPtr("Dockerfile")))
				Expect(updateBuildOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(600))))
				Expect(updateBuildOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigmapOptions successfully`, func() {
				// Construct an instance of the UpdateConfigmapOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				configMapName := "my-config-map"
				updateConfigmapOptionsName := "my-configmap"
				updateConfigmapOptionsModel := codeEngineService.NewUpdateConfigmapOptions(projectGuid, configMapName, updateConfigmapOptionsName)
				updateConfigmapOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateConfigmapOptionsModel.SetConfigMapName("my-config-map")
				updateConfigmapOptionsModel.SetName("my-configmap")
				updateConfigmapOptionsModel.SetData(make(map[string]string))
				updateConfigmapOptionsModel.SetImmutable(false)
				updateConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigmapOptionsModel).ToNot(BeNil())
				Expect(updateConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateConfigmapOptionsModel.ConfigMapName).To(Equal(core.StringPtr("my-config-map")))
				Expect(updateConfigmapOptionsModel.Name).To(Equal(core.StringPtr("my-configmap")))
				Expect(updateConfigmapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(updateConfigmapOptionsModel.Immutable).To(Equal(core.BoolPtr(false)))
				Expect(updateConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateJobOptions successfully`, func() {
				// Construct an instance of the EnvVar model
				envVarModel := new(codeenginev2.EnvVar)
				Expect(envVarModel).ToNot(BeNil())
				envVarModel.Key = core.StringPtr("MY_VARIABLE")
				envVarModel.Name = core.StringPtr("SOME")
				envVarModel.Prefix = core.StringPtr("PREFIX_")
				envVarModel.Ref = core.StringPtr("my-secret")
				envVarModel.Type = core.StringPtr("literal")
				envVarModel.Value = core.StringPtr("VALUE")
				Expect(envVarModel.Key).To(Equal(core.StringPtr("MY_VARIABLE")))
				Expect(envVarModel.Name).To(Equal(core.StringPtr("SOME")))
				Expect(envVarModel.Prefix).To(Equal(core.StringPtr("PREFIX_")))
				Expect(envVarModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(envVarModel.Type).To(Equal(core.StringPtr("literal")))
				Expect(envVarModel.Value).To(Equal(core.StringPtr("VALUE")))

				// Construct an instance of the VolumeMount model
				volumeMountModel := new(codeenginev2.VolumeMount)
				Expect(volumeMountModel).ToNot(BeNil())
				volumeMountModel.MountPath = core.StringPtr("/app")
				volumeMountModel.Name = core.StringPtr("codeengine-mount-b69u90")
				volumeMountModel.Ref = core.StringPtr("my-secret")
				volumeMountModel.Type = core.StringPtr("secret")
				Expect(volumeMountModel.MountPath).To(Equal(core.StringPtr("/app")))
				Expect(volumeMountModel.Name).To(Equal(core.StringPtr("codeengine-mount-b69u90")))
				Expect(volumeMountModel.Ref).To(Equal(core.StringPtr("my-secret")))
				Expect(volumeMountModel.Type).To(Equal(core.StringPtr("secret")))

				// Construct an instance of the UpdateJobOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				jobName := "my-job"
				updateJobOptionsModel := codeEngineService.NewUpdateJobOptions(projectGuid, jobName)
				updateJobOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateJobOptionsModel.SetJobName("my-job")
				updateJobOptionsModel.SetImageRef("icr.io/codeengine/helloworld")
				updateJobOptionsModel.SetImageSecret("my-secret")
				updateJobOptionsModel.SetName("my-job")
				updateJobOptionsModel.SetRunArgs([]string{"testString"})
				updateJobOptionsModel.SetRunAsUser(int64(1001))
				updateJobOptionsModel.SetRunCommands([]string{"testString"})
				updateJobOptionsModel.SetRunEnvVars([]codeenginev2.EnvVar{*envVarModel})
				updateJobOptionsModel.SetRunMode("daemon")
				updateJobOptionsModel.SetRunServiceAccount("default")
				updateJobOptionsModel.SetRunVolumeMounts([]codeenginev2.VolumeMount{*volumeMountModel})
				updateJobOptionsModel.SetScaleArraySpec("1-5,7-8,10")
				updateJobOptionsModel.SetScaleCpuLimit("1")
				updateJobOptionsModel.SetScaleEphemeralStorageLimit("4G")
				updateJobOptionsModel.SetScaleMaxExecutionTime(int64(7200))
				updateJobOptionsModel.SetScaleMemoryLimit("4G")
				updateJobOptionsModel.SetScaleRetryLimit(int64(3))
				updateJobOptionsModel.SetVersion("1")
				updateJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateJobOptionsModel).ToNot(BeNil())
				Expect(updateJobOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateJobOptionsModel.JobName).To(Equal(core.StringPtr("my-job")))
				Expect(updateJobOptionsModel.ImageRef).To(Equal(core.StringPtr("icr.io/codeengine/helloworld")))
				Expect(updateJobOptionsModel.ImageSecret).To(Equal(core.StringPtr("my-secret")))
				Expect(updateJobOptionsModel.Name).To(Equal(core.StringPtr("my-job")))
				Expect(updateJobOptionsModel.RunArgs).To(Equal([]string{"testString"}))
				Expect(updateJobOptionsModel.RunAsUser).To(Equal(core.Int64Ptr(int64(1001))))
				Expect(updateJobOptionsModel.RunCommands).To(Equal([]string{"testString"}))
				Expect(updateJobOptionsModel.RunEnvVars).To(Equal([]codeenginev2.EnvVar{*envVarModel}))
				Expect(updateJobOptionsModel.RunMode).To(Equal(core.StringPtr("daemon")))
				Expect(updateJobOptionsModel.RunServiceAccount).To(Equal(core.StringPtr("default")))
				Expect(updateJobOptionsModel.RunVolumeMounts).To(Equal([]codeenginev2.VolumeMount{*volumeMountModel}))
				Expect(updateJobOptionsModel.ScaleArraySpec).To(Equal(core.StringPtr("1-5,7-8,10")))
				Expect(updateJobOptionsModel.ScaleCpuLimit).To(Equal(core.StringPtr("1")))
				Expect(updateJobOptionsModel.ScaleEphemeralStorageLimit).To(Equal(core.StringPtr("4G")))
				Expect(updateJobOptionsModel.ScaleMaxExecutionTime).To(Equal(core.Int64Ptr(int64(7200))))
				Expect(updateJobOptionsModel.ScaleMemoryLimit).To(Equal(core.StringPtr("4G")))
				Expect(updateJobOptionsModel.ScaleRetryLimit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(updateJobOptionsModel.Version).To(Equal(core.StringPtr("1")))
				Expect(updateJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSecretOptions successfully`, func() {
				// Construct an instance of the UpdateSecretOptions model
				projectGuid := "15314cc3-85b4-4338-903f-c28cdee6d005"
				secretName := "my-secret"
				updateSecretOptionsName := "testString"
				updateSecretOptionsModel := codeEngineService.NewUpdateSecretOptions(projectGuid, secretName, updateSecretOptionsName)
				updateSecretOptionsModel.SetProjectGuid("15314cc3-85b4-4338-903f-c28cdee6d005")
				updateSecretOptionsModel.SetSecretName("my-secret")
				updateSecretOptionsModel.SetName("testString")
				updateSecretOptionsModel.SetBindingSecretRef("testString")
				updateSecretOptionsModel.SetCeComponents([]string{"testString"})
				updateSecretOptionsModel.SetData(make(map[string]string))
				updateSecretOptionsModel.SetFormat("testString")
				updateSecretOptionsModel.SetImmutable(true)
				updateSecretOptionsModel.SetResourceID("testString")
				updateSecretOptionsModel.SetResourceType("testString")
				updateSecretOptionsModel.SetResourcekeyID("testString")
				updateSecretOptionsModel.SetRole("testString")
				updateSecretOptionsModel.SetServiceidCrn("testString")
				updateSecretOptionsModel.SetTarget("testString")
				updateSecretOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSecretOptionsModel).ToNot(BeNil())
				Expect(updateSecretOptionsModel.ProjectGuid).To(Equal(core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005")))
				Expect(updateSecretOptionsModel.SecretName).To(Equal(core.StringPtr("my-secret")))
				Expect(updateSecretOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.BindingSecretRef).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.CeComponents).To(Equal([]string{"testString"}))
				Expect(updateSecretOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(updateSecretOptionsModel.Format).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(updateSecretOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.ResourcekeyID).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.Role).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.ServiceidCrn).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(updateSecretOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
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
