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
	Describe(`ListConfigmapsV2(listConfigmapsV2Options *ListConfigmapsV2Options) - Operation response error`, func() {
		listConfigmapsV2Path := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigmapsV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsV2Options model
				listConfigmapsV2OptionsModel := new(codeenginev2.ListConfigmapsV2Options)
				listConfigmapsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigmapsV2(listConfigmapsV2Options *ListConfigmapsV2Options)`, func() {
		listConfigmapsV2Path := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configmaps": [{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}], "limit": 5, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigmapsV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigmapsV2Options model
				listConfigmapsV2OptionsModel := new(codeenginev2.ListConfigmapsV2Options)
				listConfigmapsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListConfigmapsV2WithContext(ctx, listConfigmapsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListConfigmapsV2WithContext(ctx, listConfigmapsV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConfigmapsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configmaps": [{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}], "limit": 5, "next": {"href": "Href", "start": "Start"}}`)
				}))
			})
			It(`Invoke ListConfigmapsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListConfigmapsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigmapsV2Options model
				listConfigmapsV2OptionsModel := new(codeenginev2.ListConfigmapsV2Options)
				listConfigmapsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigmapsV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsV2Options model
				listConfigmapsV2OptionsModel := new(codeenginev2.ListConfigmapsV2Options)
				listConfigmapsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigmapsV2Options model with no property values
				listConfigmapsV2OptionsModelNew := new(codeenginev2.ListConfigmapsV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModelNew)
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
			It(`Invoke ListConfigmapsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListConfigmapsV2Options model
				listConfigmapsV2OptionsModel := new(codeenginev2.ListConfigmapsV2Options)
				listConfigmapsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				listConfigmapsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListConfigmapsV2(listConfigmapsV2OptionsModel)
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
	Describe(`CreateConfigmapV2(createConfigmapV2Options *CreateConfigmapV2Options) - Operation response error`, func() {
		createConfigmapV2Path := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapV2Path))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfigmapV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapV2Options model
				createConfigmapV2OptionsModel := new(codeenginev2.CreateConfigmapV2Options)
				createConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Data = make(map[string]string)
				createConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				createConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfigmapV2(createConfigmapV2Options *CreateConfigmapV2Options)`, func() {
		createConfigmapV2Path := "/projects/testString/configmaps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapV2Path))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateConfigmapV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the CreateConfigmapV2Options model
				createConfigmapV2OptionsModel := new(codeenginev2.CreateConfigmapV2Options)
				createConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Data = make(map[string]string)
				createConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				createConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.CreateConfigmapV2WithContext(ctx, createConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.CreateConfigmapV2WithContext(ctx, createConfigmapV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createConfigmapV2Path))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.CreateConfigmapV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateConfigmapV2Options model
				createConfigmapV2OptionsModel := new(codeenginev2.CreateConfigmapV2Options)
				createConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Data = make(map[string]string)
				createConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				createConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfigmapV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapV2Options model
				createConfigmapV2OptionsModel := new(codeenginev2.CreateConfigmapV2Options)
				createConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Data = make(map[string]string)
				createConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				createConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigmapV2Options model with no property values
				createConfigmapV2OptionsModelNew := new(codeenginev2.CreateConfigmapV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModelNew)
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
			It(`Invoke CreateConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the CreateConfigmapV2Options model
				createConfigmapV2OptionsModel := new(codeenginev2.CreateConfigmapV2Options)
				createConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				createConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Data = make(map[string]string)
				createConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				createConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				createConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.CreateConfigmapV2(createConfigmapV2OptionsModel)
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
	Describe(`GetConfigmapV2(getConfigmapV2Options *GetConfigmapV2Options) - Operation response error`, func() {
		getConfigmapV2Path := "/projects/testString/configmaps/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfigmapV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapV2Options model
				getConfigmapV2OptionsModel := new(codeenginev2.GetConfigmapV2Options)
				getConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				getConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigmapV2(getConfigmapV2Options *GetConfigmapV2Options)`, func() {
		getConfigmapV2Path := "/projects/testString/configmaps/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke GetConfigmapV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigmapV2Options model
				getConfigmapV2OptionsModel := new(codeenginev2.GetConfigmapV2Options)
				getConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				getConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetConfigmapV2WithContext(ctx, getConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetConfigmapV2WithContext(ctx, getConfigmapV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConfigmapV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke GetConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetConfigmapV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigmapV2Options model
				getConfigmapV2OptionsModel := new(codeenginev2.GetConfigmapV2Options)
				getConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				getConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigmapV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapV2Options model
				getConfigmapV2OptionsModel := new(codeenginev2.GetConfigmapV2Options)
				getConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				getConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigmapV2Options model with no property values
				getConfigmapV2OptionsModelNew := new(codeenginev2.GetConfigmapV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModelNew)
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
			It(`Invoke GetConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetConfigmapV2Options model
				getConfigmapV2OptionsModel := new(codeenginev2.GetConfigmapV2Options)
				getConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				getConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetConfigmapV2(getConfigmapV2OptionsModel)
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
	Describe(`DeleteConfigmapV2(deleteConfigmapV2Options *DeleteConfigmapV2Options)`, func() {
		deleteConfigmapV2Path := "/projects/testString/configmaps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigmapV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := codeEngineService.DeleteConfigmapV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteConfigmapV2Options model
				deleteConfigmapV2OptionsModel := new(codeenginev2.DeleteConfigmapV2Options)
				deleteConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = codeEngineService.DeleteConfigmapV2(deleteConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteConfigmapV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigmapV2Options model
				deleteConfigmapV2OptionsModel := new(codeenginev2.DeleteConfigmapV2Options)
				deleteConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				deleteConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := codeEngineService.DeleteConfigmapV2(deleteConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteConfigmapV2Options model with no property values
				deleteConfigmapV2OptionsModelNew := new(codeenginev2.DeleteConfigmapV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = codeEngineService.DeleteConfigmapV2(deleteConfigmapV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfigmapV2(updateConfigmapV2Options *UpdateConfigmapV2Options) - Operation response error`, func() {
		updateConfigmapV2Path := "/projects/testString/configmaps/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapV2Path))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfigmapV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapV2Options model
				updateConfigmapV2OptionsModel := new(codeenginev2.UpdateConfigmapV2Options)
				updateConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Data = make(map[string]string)
				updateConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				updateConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfigmapV2(updateConfigmapV2Options *UpdateConfigmapV2Options)`, func() {
		updateConfigmapV2Path := "/projects/testString/configmaps/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapV2Path))
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

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateConfigmapV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the UpdateConfigmapV2Options model
				updateConfigmapV2OptionsModel := new(codeenginev2.UpdateConfigmapV2Options)
				updateConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Data = make(map[string]string)
				updateConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				updateConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.UpdateConfigmapV2WithContext(ctx, updateConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.UpdateConfigmapV2WithContext(ctx, updateConfigmapV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigmapV2Path))
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

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": "Created", "data": {"mapKey": "Inner"}, "id": "ID", "immutable": false, "name": "Name", "type": "Type"}`)
				}))
			})
			It(`Invoke UpdateConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.UpdateConfigmapV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateConfigmapV2Options model
				updateConfigmapV2OptionsModel := new(codeenginev2.UpdateConfigmapV2Options)
				updateConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Data = make(map[string]string)
				updateConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				updateConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfigmapV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapV2Options model
				updateConfigmapV2OptionsModel := new(codeenginev2.UpdateConfigmapV2Options)
				updateConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Data = make(map[string]string)
				updateConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				updateConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigmapV2Options model with no property values
				updateConfigmapV2OptionsModelNew := new(codeenginev2.UpdateConfigmapV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModelNew)
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
			It(`Invoke UpdateConfigmapV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigmapV2Options model
				updateConfigmapV2OptionsModel := new(codeenginev2.UpdateConfigmapV2Options)
				updateConfigmapV2OptionsModel.RefreshToken = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.ConfigmapName = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Created = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Data = make(map[string]string)
				updateConfigmapV2OptionsModel.ID = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Immutable = core.BoolPtr(true)
				updateConfigmapV2OptionsModel.Name = core.StringPtr("testString")
				updateConfigmapV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.UpdateConfigmapV2(updateConfigmapV2OptionsModel)
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
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Tags = []string{"testString"}
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
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Tags = []string{"testString"}
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
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Tags = []string{"testString"}
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
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Tags = []string{"testString"}
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
				createProjectV2OptionsModel.Name = core.StringPtr("testString")
				createProjectV2OptionsModel.Region = core.StringPtr("testString")
				createProjectV2OptionsModel.ResourceGroupID = core.StringPtr("testString")
				createProjectV2OptionsModel.Tags = []string{"testString"}
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
	Describe(`ListReclamationsV2(listReclamationsV2Options *ListReclamationsV2Options) - Operation response error`, func() {
		listReclamationsV2Path := "/reclamations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReclamationsV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsV2Options model
				listReclamationsV2OptionsModel := new(codeenginev2.ListReclamationsV2Options)
				listReclamationsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listReclamationsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReclamationsV2(listReclamationsV2Options *ListReclamationsV2Options)`, func() {
		listReclamationsV2Path := "/reclamations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"reclamations": [{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListReclamationsV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ListReclamationsV2Options model
				listReclamationsV2OptionsModel := new(codeenginev2.ListReclamationsV2Options)
				listReclamationsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listReclamationsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ListReclamationsV2WithContext(ctx, listReclamationsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ListReclamationsV2WithContext(ctx, listReclamationsV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"reclamations": [{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}]}`)
				}))
			})
			It(`Invoke ListReclamationsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ListReclamationsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReclamationsV2Options model
				listReclamationsV2OptionsModel := new(codeenginev2.ListReclamationsV2Options)
				listReclamationsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listReclamationsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReclamationsV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsV2Options model
				listReclamationsV2OptionsModel := new(codeenginev2.ListReclamationsV2Options)
				listReclamationsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listReclamationsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListReclamationsV2Options model with no property values
				listReclamationsV2OptionsModelNew := new(codeenginev2.ListReclamationsV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModelNew)
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
			It(`Invoke ListReclamationsV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsV2Options model
				listReclamationsV2OptionsModel := new(codeenginev2.ListReclamationsV2Options)
				listReclamationsV2OptionsModel.RefreshToken = core.StringPtr("testString")
				listReclamationsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ListReclamationsV2(listReclamationsV2OptionsModel)
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
	Describe(`GetReclamationV2(getReclamationV2Options *GetReclamationV2Options) - Operation response error`, func() {
		getReclamationV2Path := "/reclamations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReclamationV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationV2Options model
				getReclamationV2OptionsModel := new(codeenginev2.GetReclamationV2Options)
				getReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReclamationV2(getReclamationV2Options *GetReclamationV2Options)`, func() {
		getReclamationV2Path := "/reclamations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke GetReclamationV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the GetReclamationV2Options model
				getReclamationV2OptionsModel := new(codeenginev2.GetReclamationV2Options)
				getReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.GetReclamationV2WithContext(ctx, getReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.GetReclamationV2WithContext(ctx, getReclamationV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReclamationV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke GetReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.GetReclamationV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReclamationV2Options model
				getReclamationV2OptionsModel := new(codeenginev2.GetReclamationV2Options)
				getReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReclamationV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationV2Options model
				getReclamationV2OptionsModel := new(codeenginev2.GetReclamationV2Options)
				getReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReclamationV2Options model with no property values
				getReclamationV2OptionsModelNew := new(codeenginev2.GetReclamationV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.GetReclamationV2(getReclamationV2OptionsModelNew)
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
			It(`Invoke GetReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the GetReclamationV2Options model
				getReclamationV2OptionsModel := new(codeenginev2.GetReclamationV2Options)
				getReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				getReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				getReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.GetReclamationV2(getReclamationV2OptionsModel)
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
	Describe(`ReclaimReclamationV2(reclaimReclamationV2Options *ReclaimReclamationV2Options) - Operation response error`, func() {
		reclaimReclamationV2Path := "/reclamations/testString/reclaim"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReclaimReclamationV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationV2Options model
				reclaimReclamationV2OptionsModel := new(codeenginev2.ReclaimReclamationV2Options)
				reclaimReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReclaimReclamationV2(reclaimReclamationV2Options *ReclaimReclamationV2Options)`, func() {
		reclaimReclamationV2Path := "/reclamations/testString/reclaim"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke ReclaimReclamationV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the ReclaimReclamationV2Options model
				reclaimReclamationV2OptionsModel := new(codeenginev2.ReclaimReclamationV2Options)
				reclaimReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.ReclaimReclamationV2WithContext(ctx, reclaimReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.ReclaimReclamationV2WithContext(ctx, reclaimReclamationV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(reclaimReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke ReclaimReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.ReclaimReclamationV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReclaimReclamationV2Options model
				reclaimReclamationV2OptionsModel := new(codeenginev2.ReclaimReclamationV2Options)
				reclaimReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReclaimReclamationV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationV2Options model
				reclaimReclamationV2OptionsModel := new(codeenginev2.ReclaimReclamationV2Options)
				reclaimReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReclaimReclamationV2Options model with no property values
				reclaimReclamationV2OptionsModelNew := new(codeenginev2.ReclaimReclamationV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModelNew)
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
			It(`Invoke ReclaimReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the ReclaimReclamationV2Options model
				reclaimReclamationV2OptionsModel := new(codeenginev2.ReclaimReclamationV2Options)
				reclaimReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				reclaimReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.ReclaimReclamationV2(reclaimReclamationV2OptionsModel)
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
	Describe(`RestoreReclamationV2(restoreReclamationV2Options *RestoreReclamationV2Options) - Operation response error`, func() {
		restoreReclamationV2Path := "/reclamations/testString/restore"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RestoreReclamationV2 with error: Operation response processing error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationV2Options model
				restoreReclamationV2OptionsModel := new(codeenginev2.RestoreReclamationV2Options)
				restoreReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				codeEngineService.EnableRetries(0, 0)
				result, response, operationErr = codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreReclamationV2(restoreReclamationV2Options *RestoreReclamationV2Options)`, func() {
		restoreReclamationV2Path := "/reclamations/testString/restore"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke RestoreReclamationV2 successfully with retries`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())
				codeEngineService.EnableRetries(0, 0)

				// Construct an instance of the RestoreReclamationV2Options model
				restoreReclamationV2OptionsModel := new(codeenginev2.RestoreReclamationV2Options)
				restoreReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := codeEngineService.RestoreReclamationV2WithContext(ctx, restoreReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				codeEngineService.DisableRetries()
				result, response, operationErr := codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = codeEngineService.RestoreReclamationV2WithContext(ctx, restoreReclamationV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(restoreReclamationV2Path))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "details": "Details", "id": "ID", "project_id": "ProjectID", "reason": "Reason", "resource_group_id": "ResourceGroupID", "status": "Status", "target_time": "TargetTime", "type": "Type"}`)
				}))
			})
			It(`Invoke RestoreReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := codeEngineService.RestoreReclamationV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestoreReclamationV2Options model
				restoreReclamationV2OptionsModel := new(codeenginev2.RestoreReclamationV2Options)
				restoreReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RestoreReclamationV2 with error: Operation validation and request error`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationV2Options model
				restoreReclamationV2OptionsModel := new(codeenginev2.RestoreReclamationV2Options)
				restoreReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := codeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreReclamationV2Options model with no property values
				restoreReclamationV2OptionsModelNew := new(codeenginev2.RestoreReclamationV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModelNew)
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
			It(`Invoke RestoreReclamationV2 successfully`, func() {
				codeEngineService, serviceErr := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(codeEngineService).ToNot(BeNil())

				// Construct an instance of the RestoreReclamationV2Options model
				restoreReclamationV2OptionsModel := new(codeenginev2.RestoreReclamationV2Options)
				restoreReclamationV2OptionsModel.RefreshToken = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.ProjectGuid = core.StringPtr("testString")
				restoreReclamationV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := codeEngineService.RestoreReclamationV2(restoreReclamationV2OptionsModel)
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
			It(`Invoke NewCreateConfigmapV2Options successfully`, func() {
				// Construct an instance of the CreateConfigmapV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				createConfigmapV2OptionsModel := codeEngineService.NewCreateConfigmapV2Options(refreshToken, projectGuid)
				createConfigmapV2OptionsModel.SetRefreshToken("testString")
				createConfigmapV2OptionsModel.SetProjectGuid("testString")
				createConfigmapV2OptionsModel.SetCreated("testString")
				createConfigmapV2OptionsModel.SetData(make(map[string]string))
				createConfigmapV2OptionsModel.SetID("testString")
				createConfigmapV2OptionsModel.SetImmutable(true)
				createConfigmapV2OptionsModel.SetName("testString")
				createConfigmapV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigmapV2OptionsModel).ToNot(BeNil())
				Expect(createConfigmapV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapV2OptionsModel.Created).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapV2OptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(createConfigmapV2OptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapV2OptionsModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(createConfigmapV2OptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createConfigmapV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectV2Options successfully`, func() {
				// Construct an instance of the CreateProjectV2Options model
				refreshToken := "testString"
				createProjectV2OptionsModel := codeEngineService.NewCreateProjectV2Options(refreshToken)
				createProjectV2OptionsModel.SetRefreshToken("testString")
				createProjectV2OptionsModel.SetName("testString")
				createProjectV2OptionsModel.SetRegion("testString")
				createProjectV2OptionsModel.SetResourceGroupID("testString")
				createProjectV2OptionsModel.SetTags([]string{"testString"})
				createProjectV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectV2OptionsModel).ToNot(BeNil())
				Expect(createProjectV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(createProjectV2OptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createProjectV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigmapV2Options successfully`, func() {
				// Construct an instance of the DeleteConfigmapV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				configmapName := "testString"
				deleteConfigmapV2OptionsModel := codeEngineService.NewDeleteConfigmapV2Options(refreshToken, projectGuid, configmapName)
				deleteConfigmapV2OptionsModel.SetRefreshToken("testString")
				deleteConfigmapV2OptionsModel.SetProjectGuid("testString")
				deleteConfigmapV2OptionsModel.SetConfigmapName("testString")
				deleteConfigmapV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigmapV2OptionsModel).ToNot(BeNil())
				Expect(deleteConfigmapV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigmapV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigmapV2OptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigmapV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetConfigmapV2Options successfully`, func() {
				// Construct an instance of the GetConfigmapV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				configmapName := "testString"
				getConfigmapV2OptionsModel := codeEngineService.NewGetConfigmapV2Options(refreshToken, projectGuid, configmapName)
				getConfigmapV2OptionsModel.SetRefreshToken("testString")
				getConfigmapV2OptionsModel.SetProjectGuid("testString")
				getConfigmapV2OptionsModel.SetConfigmapName("testString")
				getConfigmapV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigmapV2OptionsModel).ToNot(BeNil())
				Expect(getConfigmapV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getConfigmapV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getConfigmapV2OptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(getConfigmapV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetReclamationV2Options successfully`, func() {
				// Construct an instance of the GetReclamationV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				getReclamationV2OptionsModel := codeEngineService.NewGetReclamationV2Options(refreshToken, projectGuid)
				getReclamationV2OptionsModel.SetRefreshToken("testString")
				getReclamationV2OptionsModel.SetProjectGuid("testString")
				getReclamationV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReclamationV2OptionsModel).ToNot(BeNil())
				Expect(getReclamationV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getReclamationV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(getReclamationV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigmapsV2Options successfully`, func() {
				// Construct an instance of the ListConfigmapsV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				listConfigmapsV2OptionsModel := codeEngineService.NewListConfigmapsV2Options(refreshToken, projectGuid)
				listConfigmapsV2OptionsModel.SetRefreshToken("testString")
				listConfigmapsV2OptionsModel.SetProjectGuid("testString")
				listConfigmapsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigmapsV2OptionsModel).ToNot(BeNil())
				Expect(listConfigmapsV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(listConfigmapsV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(listConfigmapsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewListReclamationsV2Options successfully`, func() {
				// Construct an instance of the ListReclamationsV2Options model
				refreshToken := "testString"
				listReclamationsV2OptionsModel := codeEngineService.NewListReclamationsV2Options(refreshToken)
				listReclamationsV2OptionsModel.SetRefreshToken("testString")
				listReclamationsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReclamationsV2OptionsModel).ToNot(BeNil())
				Expect(listReclamationsV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReclaimReclamationV2Options successfully`, func() {
				// Construct an instance of the ReclaimReclamationV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				reclaimReclamationV2OptionsModel := codeEngineService.NewReclaimReclamationV2Options(refreshToken, projectGuid)
				reclaimReclamationV2OptionsModel.SetRefreshToken("testString")
				reclaimReclamationV2OptionsModel.SetProjectGuid("testString")
				reclaimReclamationV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(reclaimReclamationV2OptionsModel).ToNot(BeNil())
				Expect(reclaimReclamationV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(reclaimReclamationV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(reclaimReclamationV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreReclamationV2Options successfully`, func() {
				// Construct an instance of the RestoreReclamationV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				restoreReclamationV2OptionsModel := codeEngineService.NewRestoreReclamationV2Options(refreshToken, projectGuid)
				restoreReclamationV2OptionsModel.SetRefreshToken("testString")
				restoreReclamationV2OptionsModel.SetProjectGuid("testString")
				restoreReclamationV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreReclamationV2OptionsModel).ToNot(BeNil())
				Expect(restoreReclamationV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(restoreReclamationV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(restoreReclamationV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigmapV2Options successfully`, func() {
				// Construct an instance of the UpdateConfigmapV2Options model
				refreshToken := "testString"
				projectGuid := "testString"
				configmapName := "testString"
				updateConfigmapV2OptionsModel := codeEngineService.NewUpdateConfigmapV2Options(refreshToken, projectGuid, configmapName)
				updateConfigmapV2OptionsModel.SetRefreshToken("testString")
				updateConfigmapV2OptionsModel.SetProjectGuid("testString")
				updateConfigmapV2OptionsModel.SetConfigmapName("testString")
				updateConfigmapV2OptionsModel.SetCreated("testString")
				updateConfigmapV2OptionsModel.SetData(make(map[string]string))
				updateConfigmapV2OptionsModel.SetID("testString")
				updateConfigmapV2OptionsModel.SetImmutable(true)
				updateConfigmapV2OptionsModel.SetName("testString")
				updateConfigmapV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigmapV2OptionsModel).ToNot(BeNil())
				Expect(updateConfigmapV2OptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.ProjectGuid).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.ConfigmapName).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.Created).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.Data).To(Equal(make(map[string]string)))
				Expect(updateConfigmapV2OptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(updateConfigmapV2OptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigmapV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
