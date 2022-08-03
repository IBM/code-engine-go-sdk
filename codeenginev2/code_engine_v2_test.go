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
	Describe(`ListProjectsV2(listProjectsV2Options *ListProjectsV2Options) - Operation response error`, func() {
		listProjectsV2Path := "/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjectsV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsV2Options model
				listProjectsV2OptionsModel := new(codeenginev2.ListProjectsV2Options)
				listProjectsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listProjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjectsV2(listProjectsV2Options *ListProjectsV2Options)`, func() {
		listProjectsV2Path := "/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListProjectsV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsV2Options model
				listProjectsV2OptionsModel := new(codeenginev2.ListProjectsV2Options)
				listProjectsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listProjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListProjectsV2WithContext(ctx, listProjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListProjectsV2WithContext(ctx, listProjectsV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "next": {"href": "Href", "start": "Start"}, "projects": [{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListProjectsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListProjectsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsV2Options model
				listProjectsV2OptionsModel := new(codeenginev2.ListProjectsV2Options)
				listProjectsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listProjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjectsV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsV2Options model
				listProjectsV2OptionsModel := new(codeenginev2.ListProjectsV2Options)
				listProjectsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listProjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProjectsV2Options model with no property values
				listProjectsV2OptionsModelNew := new(codeenginev2.ListProjectsV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListProjectsV2(listProjectsV2OptionsModelNew)
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
			It(`Invoke ListProjectsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListProjectsV2Options model
				listProjectsV2OptionsModel := new(codeenginev2.ListProjectsV2Options)
				listProjectsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listProjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListProjectsV2(listProjectsV2OptionsModel)
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
	Describe(`CreateProjectV2(createProjectV2Options *CreateProjectV2Options) - Operation response error`, func() {
		createProjectV2Path := "/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectV2Path))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProjectV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectV2Options model
				createProjectV2OptionsModel := new(codeenginev2.CreateProjectV2Options)
				createProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createProjectV2OptionsModel.AccountID = core.StringPtr("testString")
				createProjectV2OptionsModel.Created = core.StringPtr("testString")
				createProjectV2OptionsModel.Crn = core.StringPtr("testString")
				createProjectV2OptionsModel.Details = core.StringPtr("testString")
				createProjectV2OptionsModel.ID = core.StringPtr("testString")
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Reason = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Status = core.StringPtr("testString")
				createProjectV2OptionsModel.Type = core.StringPtr("testString")
				createProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProjectV2(createProjectV2Options *CreateProjectV2Options)`, func() {
		createProjectV2Path := "/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectV2Path))
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

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateProjectV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateProjectV2Options model
				createProjectV2OptionsModel := new(codeenginev2.CreateProjectV2Options)
				createProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createProjectV2OptionsModel.AccountID = core.StringPtr("testString")
				createProjectV2OptionsModel.Created = core.StringPtr("testString")
				createProjectV2OptionsModel.Crn = core.StringPtr("testString")
				createProjectV2OptionsModel.Details = core.StringPtr("testString")
				createProjectV2OptionsModel.ID = core.StringPtr("testString")
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Reason = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Status = core.StringPtr("testString")
				createProjectV2OptionsModel.Type = core.StringPtr("testString")
				createProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateProjectV2WithContext(ctx, createProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateProjectV2WithContext(ctx, createProjectV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProjectV2Path))
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

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateProjectV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateProjectV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProjectV2Options model
				createProjectV2OptionsModel := new(codeenginev2.CreateProjectV2Options)
				createProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createProjectV2OptionsModel.AccountID = core.StringPtr("testString")
				createProjectV2OptionsModel.Created = core.StringPtr("testString")
				createProjectV2OptionsModel.Crn = core.StringPtr("testString")
				createProjectV2OptionsModel.Details = core.StringPtr("testString")
				createProjectV2OptionsModel.ID = core.StringPtr("testString")
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Reason = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Status = core.StringPtr("testString")
				createProjectV2OptionsModel.Type = core.StringPtr("testString")
				createProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProjectV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectV2Options model
				createProjectV2OptionsModel := new(codeenginev2.CreateProjectV2Options)
				createProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createProjectV2OptionsModel.AccountID = core.StringPtr("testString")
				createProjectV2OptionsModel.Created = core.StringPtr("testString")
				createProjectV2OptionsModel.Crn = core.StringPtr("testString")
				createProjectV2OptionsModel.Details = core.StringPtr("testString")
				createProjectV2OptionsModel.ID = core.StringPtr("testString")
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Reason = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Status = core.StringPtr("testString")
				createProjectV2OptionsModel.Type = core.StringPtr("testString")
				createProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectV2Options model with no property values
				createProjectV2OptionsModelNew := new(codeenginev2.CreateProjectV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateProjectV2(createProjectV2OptionsModelNew)
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
			It(`Invoke CreateProjectV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateProjectV2Options model
				createProjectV2OptionsModel := new(codeenginev2.CreateProjectV2Options)
				createProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createProjectV2OptionsModel.AccountID = core.StringPtr("testString")
				createProjectV2OptionsModel.Created = core.StringPtr("testString")
				createProjectV2OptionsModel.Crn = core.StringPtr("testString")
				createProjectV2OptionsModel.Details = core.StringPtr("testString")
				createProjectV2OptionsModel.ID = core.StringPtr("testString")
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Reason = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Status = core.StringPtr("testString")
				createProjectV2OptionsModel.Type = core.StringPtr("testString")
				createProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateProjectV2(createProjectV2OptionsModel)
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
	Describe(`GetProjectV2(getProjectV2Options *GetProjectV2Options) - Operation response error`, func() {
		getProjectV2Path := "/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectV2Options model
				getProjectV2OptionsModel := new(codeenginev2.GetProjectV2Options)
				getProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetProjectV2(getProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetProjectV2(getProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectV2(getProjectV2Options *GetProjectV2Options)`, func() {
		getProjectV2Path := "/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
				}))
			})
			It(`Invoke GetProjectV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectV2Options model
				getProjectV2OptionsModel := new(codeenginev2.GetProjectV2Options)
				getProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetProjectV2WithContext(ctx, getProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetProjectV2(getProjectV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetProjectV2WithContext(ctx, getProjectV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProjectV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "created": "Created", "crn": "Crn", "details": "Details", "id": "ID", "name": "Name", "reason": "Reason", "region": "Region", "resource_group_id": "ResourceGroupID", "status": "Status", "type": "Type"}`)
				}))
			})
			It(`Invoke GetProjectV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetProjectV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectV2Options model
				getProjectV2OptionsModel := new(codeenginev2.GetProjectV2Options)
				getProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetProjectV2(getProjectV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectV2Options model
				getProjectV2OptionsModel := new(codeenginev2.GetProjectV2Options)
				getProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetProjectV2(getProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectV2Options model with no property values
				getProjectV2OptionsModelNew := new(codeenginev2.GetProjectV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetProjectV2(getProjectV2OptionsModelNew)
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
			It(`Invoke GetProjectV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetProjectV2Options model
				getProjectV2OptionsModel := new(codeenginev2.GetProjectV2Options)
				getProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetProjectV2(getProjectV2OptionsModel)
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
	Describe(`DeleteProjectV2(deleteProjectV2Options *DeleteProjectV2Options)`, func() {
		deleteProjectV2Path := "/projects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteProjectV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteProjectV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectV2Options model
				deleteProjectV2OptionsModel := new(codeenginev2.DeleteProjectV2Options)
				deleteProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				deleteProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteProjectV2(deleteProjectV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProjectV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectV2Options model
				deleteProjectV2OptionsModel := new(codeenginev2.DeleteProjectV2Options)
				deleteProjectV2OptionsModel.RefreshToken = core.StringPtr("testString")
				deleteProjectV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteProjectV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteProjectV2(deleteProjectV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectV2Options model with no property values
				deleteProjectV2OptionsModelNew := new(codeenginev2.DeleteProjectV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteProjectV2(deleteProjectV2OptionsModelNew)
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
			It(`Invoke NewCreateProjectV2Options successfully`, func() {
				// Construct an instance of the CreateProjectV2Options model
				refreshToken := "testString"
				createProjectV2OptionsModel := codeEngineService.NewCreateProjectV2Options(refreshToken)
				createProjectV2OptionsModel.SetRefreshToken("testString")
				createProjectV2OptionsModel.SetAccountID("testString")
				createProjectV2OptionsModel.SetCreated("testString")
				createProjectV2OptionsModel.SetCrn("testString")
				createProjectV2OptionsModel.SetDetails("testString")
				createProjectV2OptionsModel.SetID("testString")
				createProjectV2OptionsModel.SetName("testString")
				createProjectV2OptionsModel.SetReason("testString")
				createProjectV2OptionsModel.SetRegion("testString")
				createProjectV2OptionsModel.SetResourceGroupID("testString")
				createProjectV2OptionsModel.SetStatus("testString")
				createProjectV2OptionsModel.SetType("testString")
				createProjectV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectV2OptionsModel).ToNot(BeNil())
				Expect(createProjectV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Created).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Details).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Reason).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Status).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectV2Options successfully`, func() {
				// Construct an instance of the DeleteProjectV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				deleteProjectV2OptionsModel := codeEngineService.NewDeleteProjectV2Options(refreshToken, projectGuid)
				deleteProjectV2OptionsModel.SetRefreshToken("testString")
				deleteProjectV2OptionsModel.SetProjectGuid("testString")
				deleteProjectV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectV2OptionsModel).ToNot(BeNil())
				Expect(deleteProjectV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectV2Options successfully`, func() {
				// Construct an instance of the GetProjectV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				getProjectV2OptionsModel := codeEngineService.NewGetProjectV2Options(refreshToken, projectGuid)
				getProjectV2OptionsModel.SetRefreshToken("testString")
				getProjectV2OptionsModel.SetProjectGuid("testString")
				getProjectV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectV2OptionsModel).ToNot(BeNil())
				Expect(getProjectV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getProjectV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getProjectV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsV2Options successfully`, func() {
				// Construct an instance of the ListProjectsV2Options model
				refreshToken := "testString"
				listProjectsV2OptionsModel := codeEngineService.NewListProjectsV2Options(refreshToken)
				listProjectsV2OptionsModel.SetRefreshToken("testString")
				listProjectsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsV2OptionsModel).ToNot(BeNil())
				Expect(listProjectsV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
