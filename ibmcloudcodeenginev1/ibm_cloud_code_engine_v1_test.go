/**
 * (C) Copyright IBM Corp. 2021.
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

package ibmcloudcodeenginev1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/code-engine-go-sdk/ibmcloudcodeenginev1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IbmCloudCodeEngineV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudCodeEngineService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudCodeEngineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
				URL: "https://ibmcloudcodeenginev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudCodeEngineService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_CODE_ENGINE_URL":       "https://ibmcloudcodeenginev1/api",
				"IBM_CLOUD_CODE_ENGINE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1UsingExternalConfig(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{})
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1UsingExternalConfig(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1UsingExternalConfig(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{})
				err := ibmCloudCodeEngineService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_CODE_ENGINE_URL":       "https://ibmcloudcodeenginev1/api",
				"IBM_CLOUD_CODE_ENGINE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1UsingExternalConfig(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudCodeEngineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_CODE_ENGINE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1UsingExternalConfig(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudCodeEngineService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`ListKubeconfig(listKubeconfigOptions *ListKubeconfigOptions)`, func() {
		listKubeconfigPath := "/namespaces/testString/config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKubeconfigPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "text/plain")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke ListKubeconfig successfully`, func() {
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudCodeEngineService.ListKubeconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListKubeconfigOptions model
				listKubeconfigOptionsModel := new(ibmcloudcodeenginev1.ListKubeconfigOptions)
				listKubeconfigOptionsModel.RefreshToken = core.StringPtr("testString")
				listKubeconfigOptionsModel.ID = core.StringPtr("testString")
				listKubeconfigOptionsModel.Accept = core.StringPtr("text/plain")
				listKubeconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudCodeEngineService.ListKubeconfig(listKubeconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr = ibmCloudCodeEngineService.ListKubeconfigWithContext(ctx, listKubeconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			It(`Invoke ListKubeconfig with error: Operation validation and request error`, func() {
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())

				// Construct an instance of the ListKubeconfigOptions model
				listKubeconfigOptionsModel := new(ibmcloudcodeenginev1.ListKubeconfigOptions)
				listKubeconfigOptionsModel.RefreshToken = core.StringPtr("testString")
				listKubeconfigOptionsModel.ID = core.StringPtr("testString")
				listKubeconfigOptionsModel.Accept = core.StringPtr("text/plain")
				listKubeconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudCodeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudCodeEngineService.ListKubeconfig(listKubeconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListKubeconfigOptions model with no property values
				listKubeconfigOptionsModelNew := new(ibmcloudcodeenginev1.ListKubeconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudCodeEngineService.ListKubeconfig(listKubeconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetKubeconfig(getKubeconfigOptions *GetKubeconfigOptions)`, func() {
		getKubeconfigPath := "/project/testString/config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKubeconfigPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Delegated-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Delegated-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "text/plain")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke GetKubeconfig successfully`, func() {
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudCodeEngineService.GetKubeconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKubeconfigOptions model
				getKubeconfigOptionsModel := new(ibmcloudcodeenginev1.GetKubeconfigOptions)
				getKubeconfigOptionsModel.XDelegatedRefreshToken = core.StringPtr("testString")
				getKubeconfigOptionsModel.ID = core.StringPtr("testString")
				getKubeconfigOptionsModel.Accept = core.StringPtr("text/plain")
				getKubeconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudCodeEngineService.GetKubeconfig(getKubeconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr = ibmCloudCodeEngineService.GetKubeconfigWithContext(ctx, getKubeconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			It(`Invoke GetKubeconfig with error: Operation validation and request error`, func() {
				ibmCloudCodeEngineService, serviceErr := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudCodeEngineService).ToNot(BeNil())

				// Construct an instance of the GetKubeconfigOptions model
				getKubeconfigOptionsModel := new(ibmcloudcodeenginev1.GetKubeconfigOptions)
				getKubeconfigOptionsModel.XDelegatedRefreshToken = core.StringPtr("testString")
				getKubeconfigOptionsModel.ID = core.StringPtr("testString")
				getKubeconfigOptionsModel.Accept = core.StringPtr("text/plain")
				getKubeconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudCodeEngineService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudCodeEngineService.GetKubeconfig(getKubeconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKubeconfigOptions model with no property values
				getKubeconfigOptionsModelNew := new(ibmcloudcodeenginev1.GetKubeconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudCodeEngineService.GetKubeconfig(getKubeconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmCloudCodeEngineService, _ := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
				URL:           "http://ibmcloudcodeenginev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetKubeconfigOptions successfully`, func() {
				// Construct an instance of the GetKubeconfigOptions model
				xDelegatedRefreshToken := "testString"
				id := "testString"
				getKubeconfigOptionsModel := ibmCloudCodeEngineService.NewGetKubeconfigOptions(xDelegatedRefreshToken, id)
				getKubeconfigOptionsModel.SetXDelegatedRefreshToken("testString")
				getKubeconfigOptionsModel.SetID("testString")
				getKubeconfigOptionsModel.SetAccept("text/plain")
				getKubeconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKubeconfigOptionsModel).ToNot(BeNil())
				Expect(getKubeconfigOptionsModel.XDelegatedRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getKubeconfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKubeconfigOptionsModel.Accept).To(Equal(core.StringPtr("text/plain")))
				Expect(getKubeconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListKubeconfigOptions successfully`, func() {
				// Construct an instance of the ListKubeconfigOptions model
				refreshToken := "testString"
				id := "testString"
				listKubeconfigOptionsModel := ibmCloudCodeEngineService.NewListKubeconfigOptions(refreshToken, id)
				listKubeconfigOptionsModel.SetRefreshToken("testString")
				listKubeconfigOptionsModel.SetID("testString")
				listKubeconfigOptionsModel.SetAccept("text/plain")
				listKubeconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listKubeconfigOptionsModel).ToNot(BeNil())
				Expect(listKubeconfigOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(listKubeconfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listKubeconfigOptionsModel.Accept).To(Equal(core.StringPtr("text/plain")))
				Expect(listKubeconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
