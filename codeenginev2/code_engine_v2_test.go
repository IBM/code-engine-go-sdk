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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}]}`)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}]}`)
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
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
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
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
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
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")
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
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")
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
				createProjectOptionsModel.ResourceGroupID = core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")
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
		getProjectPath := "/projects/testString"
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		getProjectPath := "/projects/testString"
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				getProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		deleteProjectPath := "/projects/testString"
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
				deleteProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				deleteProjectOptionsModel.ProjectGuid = core.StringPtr("testString")
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
	Describe(`ListConfigmaps(listConfigmapsOptions *ListConfigmapsOptions) - Operation response error`, func() {
		listConfigmapsPath := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))
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
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		listConfigmapsPath := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configmaps": [{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}], "limit": 5, "next": {"href": "Href", "start": "Start"}}`)
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
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("testString")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configmaps": [{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}], "limit": 5, "next": {"href": "Href", "start": "Start"}}`)
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
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				listConfigmapsOptionsModel.ProjectGuid = core.StringPtr("testString")
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
	})
	Describe(`CreateConfigmap(createConfigmapOptions *CreateConfigmapOptions) - Operation response error`, func() {
		createConfigmapPath := "/projects/testString/configmaps"
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
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
		createConfigmapPath := "/projects/testString/configmaps"
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
				createConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapOptionsModel.Data = make(map[string]string)
				createConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				createConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
		getConfigmapPath := "/projects/testString/configmaps/testString"
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
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
		getConfigmapPath := "/projects/testString/configmaps/testString"
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
				getConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
		deleteConfigmapPath := "/projects/testString/configmaps/testString"
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
				deleteConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
				deleteConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
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
		updateConfigmapPath := "/projects/testString/configmaps/testString"
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
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
		updateConfigmapPath := "/projects/testString/configmaps/testString"
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
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
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
				updateConfigmapOptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapOptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapOptionsModel.Data = make(map[string]string)
				updateConfigmapOptionsModel.Immutable = core.BoolPtr(false)
				updateConfigmapOptionsModel.Name = core.StringPtr("my-configmap")
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
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions) - Operation response error`, func() {
		listReclamationsPath := "/reclamations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"reclamations": [{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}]}`)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"reclamations": [{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}]}`)
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
	})
	Describe(`GetReclamation(getReclamationOptions *GetReclamationOptions) - Operation response error`, func() {
		getReclamationPath := "/reclamations/testString"
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
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		getReclamationPath := "/reclamations/testString"
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				getReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		reclaimReclamationPath := "/reclamations/testString/reclaim"
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
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		reclaimReclamationPath := "/reclamations/testString/reclaim"
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				reclaimReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		restoreReclamationPath := "/reclamations/testString/restore"
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
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
		restoreReclamationPath := "/reclamations/testString/restore"
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
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
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
				restoreReclamationOptionsModel.ProjectGuid = core.StringPtr("testString")
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
			It(`Invoke NewCreateConfigmapOptions successfully`, func() {
				// Construct an instance of the CreateConfigmapOptions model
				projectGuid := "testString"
				createConfigmapOptionsModel := codeEngineService.NewCreateConfigmapOptions(projectGuid)
				createConfigmapOptionsModel.SetProjectGuid("testString")
				createConfigmapOptionsModel.SetData(make(map[string]string))
				createConfigmapOptionsModel.SetImmutable(false)
				createConfigmapOptionsModel.SetName("my-configmap")
				createConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigmapOptionsModel).ToNot(BeNil())
				Expect(createConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createConfigmapOptionsModel.Immutable).To(Equal(core.BoolPtr(false)))
				Expect(createConfigmapOptionsModel.Name).To(Equal(core.StringPtr("my-configmap")))
				Expect(createConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := codeEngineService.NewCreateProjectOptions()
				createProjectOptionsModel.SetName("my-project")
				createProjectOptionsModel.SetRegion("us-east")
				createProjectOptionsModel.SetResourceGroupID("5c49eabcf5e85881a37e2d100a33b3df")
				createProjectOptionsModel.SetTags([]string{"testString"})
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("my-project")))
				Expect(createProjectOptionsModel.Region).To(Equal(core.StringPtr("us-east")))
				Expect(createProjectOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("5c49eabcf5e85881a37e2d100a33b3df")))
				Expect(createProjectOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigmapOptions successfully`, func() {
				// Construct an instance of the DeleteConfigmapOptions model
				projectGuid := "testString"
				configmapName := "testString"
				deleteConfigmapOptionsModel := codeEngineService.NewDeleteConfigmapOptions(projectGuid, configmapName)
				deleteConfigmapOptionsModel.SetProjectGuid("testString")
				deleteConfigmapOptionsModel.SetConfigmapName("testString")
				deleteConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigmapOptionsModel).ToNot(BeNil())
				Expect(deleteConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigmapOptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				projectGuid := "testString"
				deleteProjectOptionsModel := codeEngineService.NewDeleteProjectOptions(projectGuid)
				deleteProjectOptionsModel.SetProjectGuid("testString")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigmapOptions successfully`, func() {
				// Construct an instance of the GetConfigmapOptions model
				projectGuid := "testString"
				configmapName := "testString"
				getConfigmapOptionsModel := codeEngineService.NewGetConfigmapOptions(projectGuid, configmapName)
				getConfigmapOptionsModel.SetProjectGuid("testString")
				getConfigmapOptionsModel.SetConfigmapName("testString")
				getConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigmapOptionsModel).ToNot(BeNil())
				Expect(getConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getConfigmapOptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(getConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				projectGuid := "testString"
				getProjectOptionsModel := codeEngineService.NewGetProjectOptions(projectGuid)
				getProjectOptionsModel.SetProjectGuid("testString")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReclamationOptions successfully`, func() {
				// Construct an instance of the GetReclamationOptions model
				projectGuid := "testString"
				getReclamationOptionsModel := codeEngineService.NewGetReclamationOptions(projectGuid)
				getReclamationOptionsModel.SetProjectGuid("testString")
				getReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReclamationOptionsModel).ToNot(BeNil())
				Expect(getReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigmapsOptions successfully`, func() {
				// Construct an instance of the ListConfigmapsOptions model
				projectGuid := "testString"
				listConfigmapsOptionsModel := codeEngineService.NewListConfigmapsOptions(projectGuid)
				listConfigmapsOptionsModel.SetProjectGuid("testString")
				listConfigmapsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigmapsOptionsModel).ToNot(BeNil())
				Expect(listConfigmapsOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(listConfigmapsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := codeEngineService.NewListProjectsOptions()
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReclamationsOptions successfully`, func() {
				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := codeEngineService.NewListReclamationsOptions()
				listReclamationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReclamationsOptionsModel).ToNot(BeNil())
				Expect(listReclamationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReclaimReclamationOptions successfully`, func() {
				// Construct an instance of the ReclaimReclamationOptions model
				projectGuid := "testString"
				reclaimReclamationOptionsModel := codeEngineService.NewReclaimReclamationOptions(projectGuid)
				reclaimReclamationOptionsModel.SetProjectGuid("testString")
				reclaimReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(reclaimReclamationOptionsModel).ToNot(BeNil())
				Expect(reclaimReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(reclaimReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreReclamationOptions successfully`, func() {
				// Construct an instance of the RestoreReclamationOptions model
				projectGuid := "testString"
				restoreReclamationOptionsModel := codeEngineService.NewRestoreReclamationOptions(projectGuid)
				restoreReclamationOptionsModel.SetProjectGuid("testString")
				restoreReclamationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreReclamationOptionsModel).ToNot(BeNil())
				Expect(restoreReclamationOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(restoreReclamationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigmapOptions successfully`, func() {
				// Construct an instance of the UpdateConfigmapOptions model
				projectGuid := "testString"
				configmapName := "testString"
				updateConfigmapOptionsModel := codeEngineService.NewUpdateConfigmapOptions(projectGuid, configmapName)
				updateConfigmapOptionsModel.SetProjectGuid("testString")
				updateConfigmapOptionsModel.SetConfigmapName("testString")
				updateConfigmapOptionsModel.SetData(make(map[string]string))
				updateConfigmapOptionsModel.SetImmutable(false)
				updateConfigmapOptionsModel.SetName("my-configmap")
				updateConfigmapOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigmapOptionsModel).ToNot(BeNil())
				Expect(updateConfigmapOptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapOptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapOptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(updateConfigmapOptionsModel.Immutable).To(Equal(core.BoolPtr(false)))
				Expect(updateConfigmapOptionsModel.Name).To(Equal(core.StringPtr("my-configmap")))
				Expect(updateConfigmapOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
