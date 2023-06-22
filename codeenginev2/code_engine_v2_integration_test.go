// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the codeenginev2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CodeEngineV2 Integration Tests`, func() {
	const externalConfigFile = "../code_engine_v2.env"

	var (
		err          error
		codeEngineService *codeenginev2.CodeEngineV2
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(codeenginev2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			codeEngineServiceOptions := &codeenginev2.CodeEngineV2Options{}

			codeEngineService, err = codeenginev2.NewCodeEngineV2UsingExternalConfig(codeEngineServiceOptions)
			Expect(err).To(BeNil())
			Expect(codeEngineService).ToNot(BeNil())
			Expect(codeEngineService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			codeEngineService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListProjects - List all projects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProjects(listProjectsOptions *ListProjectsOptions) with pagination`, func(){
			listProjectsOptions := &codeenginev2.ListProjectsOptions{
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listProjectsOptions.Start = nil
			listProjectsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.Project
			for {
				projectList, response, err := codeEngineService.ListProjects(listProjectsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(projectList).ToNot(BeNil())
				allResults = append(allResults, projectList.Projects...)

				listProjectsOptions.Start, err = projectList.GetNextStart()
				Expect(err).To(BeNil())

				if listProjectsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListProjects(listProjectsOptions *ListProjectsOptions) using ProjectsPager`, func(){
			listProjectsOptions := &codeenginev2.ListProjectsOptions{
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewProjectsPager(listProjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.Project
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewProjectsPager(listProjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProjects() returned a total of %d item(s) using ProjectsPager.\n", len(allResults))
		})
	})

	Describe(`CreateProject - Create a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
			createProjectOptions := &codeenginev2.CreateProjectOptions{
				Name: core.StringPtr("my-project"),
				ResourceGroupID: core.StringPtr("b91e849cedb04e7e92bd68c040c672dc"),
				Tags: []string{"testString"},
			}

			project, response, err := codeEngineService.CreateProject(createProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`GetProject - Get a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
			getProjectOptions := &codeenginev2.GetProjectOptions{
				ID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
			}

			project, response, err := codeEngineService.GetProject(getProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`GetProjectEgressIps - List egress IP addresses`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProjectEgressIps(getProjectEgressIpsOptions *GetProjectEgressIpsOptions)`, func() {
			getProjectEgressIpsOptions := &codeenginev2.GetProjectEgressIpsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
			}

			projectEgressIpAddresses, response, err := codeEngineService.GetProjectEgressIps(getProjectEgressIpsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectEgressIpAddresses).ToNot(BeNil())
		})
	})

	Describe(`GetProjectStatusDetails - Get the status details for a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProjectStatusDetails(getProjectStatusDetailsOptions *GetProjectStatusDetailsOptions)`, func() {
			getProjectStatusDetailsOptions := &codeenginev2.GetProjectStatusDetailsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
			}

			projectStatusDetails, response, err := codeEngineService.GetProjectStatusDetails(getProjectStatusDetailsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectStatusDetails).ToNot(BeNil())
		})
	})

	Describe(`ListApps - List applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApps(listAppsOptions *ListAppsOptions) with pagination`, func(){
			listAppsOptions := &codeenginev2.ListAppsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listAppsOptions.Start = nil
			listAppsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.App
			for {
				appList, response, err := codeEngineService.ListApps(listAppsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(appList).ToNot(BeNil())
				allResults = append(allResults, appList.Apps...)

				listAppsOptions.Start, err = appList.GetNextStart()
				Expect(err).To(BeNil())

				if listAppsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListApps(listAppsOptions *ListAppsOptions) using AppsPager`, func(){
			listAppsOptions := &codeenginev2.ListAppsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewAppsPager(listAppsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.App
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewAppsPager(listAppsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListApps() returned a total of %d item(s) using AppsPager.\n", len(allResults))
		})
	})

	Describe(`CreateApp - Create an application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApp(createAppOptions *CreateAppOptions)`, func() {
			envVarPrototypeModel := &codeenginev2.EnvVarPrototype{
				Key: core.StringPtr("MY_VARIABLE"),
				Name: core.StringPtr("SOME"),
				Prefix: core.StringPtr("PREFIX_"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("literal"),
				Value: core.StringPtr("VALUE"),
			}

			volumeMountPrototypeModel := &codeenginev2.VolumeMountPrototype{
				MountPath: core.StringPtr("/app"),
				Name: core.StringPtr("codeengine-mount-b69u90"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("secret"),
			}

			createAppOptions := &codeenginev2.CreateAppOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				ImageReference: core.StringPtr("icr.io/codeengine/helloworld"),
				Name: core.StringPtr("my-app"),
				ImagePort: core.Int64Ptr(int64(8080)),
				ImageSecret: core.StringPtr("my-secret"),
				ManagedDomainMappings: core.StringPtr("local_public"),
				RunArguments: []string{"testString"},
				RunAsUser: core.Int64Ptr(int64(1001)),
				RunCommands: []string{"testString"},
				RunEnvVariables: []codeenginev2.EnvVarPrototype{*envVarPrototypeModel},
				RunServiceAccount: core.StringPtr("default"),
				RunVolumeMounts: []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel},
				ScaleConcurrency: core.Int64Ptr(int64(100)),
				ScaleConcurrencyTarget: core.Int64Ptr(int64(80)),
				ScaleCpuLimit: core.StringPtr("1"),
				ScaleDownDelay: core.Int64Ptr(int64(300)),
				ScaleEphemeralStorageLimit: core.StringPtr("4G"),
				ScaleInitialInstances: core.Int64Ptr(int64(1)),
				ScaleMaxInstances: core.Int64Ptr(int64(10)),
				ScaleMemoryLimit: core.StringPtr("4G"),
				ScaleMinInstances: core.Int64Ptr(int64(1)),
				ScaleRequestTimeout: core.Int64Ptr(int64(300)),
			}

			app, response, err := codeEngineService.CreateApp(createAppOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(app).ToNot(BeNil())
		})
	})

	Describe(`GetApp - Get an application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApp(getAppOptions *GetAppOptions)`, func() {
			getAppOptions := &codeenginev2.GetAppOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-app"),
			}

			app, response, err := codeEngineService.GetApp(getAppOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(app).ToNot(BeNil())
		})
	})

	Describe(`UpdateApp - Update an application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateApp(updateAppOptions *UpdateAppOptions)`, func() {
			envVarPrototypeModel := &codeenginev2.EnvVarPrototype{
				Key: core.StringPtr("MY_VARIABLE"),
				Name: core.StringPtr("SOME"),
				Prefix: core.StringPtr("PREFIX_"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("literal"),
				Value: core.StringPtr("VALUE"),
			}

			volumeMountPrototypeModel := &codeenginev2.VolumeMountPrototype{
				MountPath: core.StringPtr("/app"),
				Name: core.StringPtr("codeengine-mount-b69u90"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("secret"),
			}

			appPatchModel := &codeenginev2.AppPatch{
				ImagePort: core.Int64Ptr(int64(8080)),
				ImageReference: core.StringPtr("icr.io/codeengine/helloworld"),
				ImageSecret: core.StringPtr("my-secret"),
				ManagedDomainMappings: core.StringPtr("local_public"),
				RunArguments: []string{"testString"},
				RunAsUser: core.Int64Ptr(int64(1001)),
				RunCommands: []string{"testString"},
				RunEnvVariables: []codeenginev2.EnvVarPrototype{*envVarPrototypeModel},
				RunServiceAccount: core.StringPtr("default"),
				RunVolumeMounts: []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel},
				ScaleConcurrency: core.Int64Ptr(int64(100)),
				ScaleConcurrencyTarget: core.Int64Ptr(int64(80)),
				ScaleCpuLimit: core.StringPtr("1"),
				ScaleDownDelay: core.Int64Ptr(int64(300)),
				ScaleEphemeralStorageLimit: core.StringPtr("4G"),
				ScaleInitialInstances: core.Int64Ptr(int64(1)),
				ScaleMaxInstances: core.Int64Ptr(int64(10)),
				ScaleMemoryLimit: core.StringPtr("4G"),
				ScaleMinInstances: core.Int64Ptr(int64(1)),
				ScaleRequestTimeout: core.Int64Ptr(int64(300)),
			}
			appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateAppOptions := &codeenginev2.UpdateAppOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-app"),
				IfMatch: core.StringPtr("testString"),
				App: appPatchModelAsPatch,
			}

			app, response, err := codeEngineService.UpdateApp(updateAppOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(app).ToNot(BeNil())
		})
	})

	Describe(`ListAppRevisions - List application revisions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAppRevisions(listAppRevisionsOptions *ListAppRevisionsOptions) with pagination`, func(){
			listAppRevisionsOptions := &codeenginev2.ListAppRevisionsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				AppName: core.StringPtr("my-app"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listAppRevisionsOptions.Start = nil
			listAppRevisionsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.AppRevision
			for {
				appRevisionList, response, err := codeEngineService.ListAppRevisions(listAppRevisionsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(appRevisionList).ToNot(BeNil())
				allResults = append(allResults, appRevisionList.Revisions...)

				listAppRevisionsOptions.Start, err = appRevisionList.GetNextStart()
				Expect(err).To(BeNil())

				if listAppRevisionsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListAppRevisions(listAppRevisionsOptions *ListAppRevisionsOptions) using AppRevisionsPager`, func(){
			listAppRevisionsOptions := &codeenginev2.ListAppRevisionsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				AppName: core.StringPtr("my-app"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewAppRevisionsPager(listAppRevisionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.AppRevision
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewAppRevisionsPager(listAppRevisionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAppRevisions() returned a total of %d item(s) using AppRevisionsPager.\n", len(allResults))
		})
	})

	Describe(`GetAppRevision - Get an application revision`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAppRevision(getAppRevisionOptions *GetAppRevisionOptions)`, func() {
			getAppRevisionOptions := &codeenginev2.GetAppRevisionOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				AppName: core.StringPtr("my-app"),
				Name: core.StringPtr("my-app-00001"),
			}

			appRevision, response, err := codeEngineService.GetAppRevision(getAppRevisionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(appRevision).ToNot(BeNil())
		})
	})

	Describe(`ListJobs - List jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListJobs(listJobsOptions *ListJobsOptions) with pagination`, func(){
			listJobsOptions := &codeenginev2.ListJobsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listJobsOptions.Start = nil
			listJobsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.Job
			for {
				jobList, response, err := codeEngineService.ListJobs(listJobsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(jobList).ToNot(BeNil())
				allResults = append(allResults, jobList.Jobs...)

				listJobsOptions.Start, err = jobList.GetNextStart()
				Expect(err).To(BeNil())

				if listJobsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListJobs(listJobsOptions *ListJobsOptions) using JobsPager`, func(){
			listJobsOptions := &codeenginev2.ListJobsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewJobsPager(listJobsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.Job
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewJobsPager(listJobsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListJobs() returned a total of %d item(s) using JobsPager.\n", len(allResults))
		})
	})

	Describe(`CreateJob - Create a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
			envVarPrototypeModel := &codeenginev2.EnvVarPrototype{
				Key: core.StringPtr("MY_VARIABLE"),
				Name: core.StringPtr("SOME"),
				Prefix: core.StringPtr("PREFIX_"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("literal"),
				Value: core.StringPtr("VALUE"),
			}

			volumeMountPrototypeModel := &codeenginev2.VolumeMountPrototype{
				MountPath: core.StringPtr("/app"),
				Name: core.StringPtr("codeengine-mount-b69u90"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("secret"),
			}

			createJobOptions := &codeenginev2.CreateJobOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				ImageReference: core.StringPtr("icr.io/codeengine/helloworld"),
				Name: core.StringPtr("my-job"),
				ImageSecret: core.StringPtr("my-secret"),
				RunArguments: []string{"testString"},
				RunAsUser: core.Int64Ptr(int64(1001)),
				RunCommands: []string{"testString"},
				RunEnvVariables: []codeenginev2.EnvVarPrototype{*envVarPrototypeModel},
				RunMode: core.StringPtr("task"),
				RunServiceAccount: core.StringPtr("default"),
				RunVolumeMounts: []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel},
				ScaleArraySpec: core.StringPtr("1-5,7-8,10"),
				ScaleCpuLimit: core.StringPtr("1"),
				ScaleEphemeralStorageLimit: core.StringPtr("4G"),
				ScaleMaxExecutionTime: core.Int64Ptr(int64(7200)),
				ScaleMemoryLimit: core.StringPtr("4G"),
				ScaleRetryLimit: core.Int64Ptr(int64(3)),
			}

			job, response, err := codeEngineService.CreateJob(createJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`GetJob - Get a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetJob(getJobOptions *GetJobOptions)`, func() {
			getJobOptions := &codeenginev2.GetJobOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-job"),
			}

			job, response, err := codeEngineService.GetJob(getJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`UpdateJob - Update a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateJob(updateJobOptions *UpdateJobOptions)`, func() {
			envVarPrototypeModel := &codeenginev2.EnvVarPrototype{
				Key: core.StringPtr("MY_VARIABLE"),
				Name: core.StringPtr("SOME"),
				Prefix: core.StringPtr("PREFIX_"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("literal"),
				Value: core.StringPtr("VALUE"),
			}

			volumeMountPrototypeModel := &codeenginev2.VolumeMountPrototype{
				MountPath: core.StringPtr("/app"),
				Name: core.StringPtr("codeengine-mount-b69u90"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("secret"),
			}

			jobPatchModel := &codeenginev2.JobPatch{
				ImageReference: core.StringPtr("icr.io/codeengine/helloworld"),
				ImageSecret: core.StringPtr("my-secret"),
				RunArguments: []string{"testString"},
				RunAsUser: core.Int64Ptr(int64(1001)),
				RunCommands: []string{"testString"},
				RunEnvVariables: []codeenginev2.EnvVarPrototype{*envVarPrototypeModel},
				RunMode: core.StringPtr("task"),
				RunServiceAccount: core.StringPtr("default"),
				RunVolumeMounts: []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel},
				ScaleArraySpec: core.StringPtr("1-5,7-8,10"),
				ScaleCpuLimit: core.StringPtr("1"),
				ScaleEphemeralStorageLimit: core.StringPtr("4G"),
				ScaleMaxExecutionTime: core.Int64Ptr(int64(7200)),
				ScaleMemoryLimit: core.StringPtr("4G"),
				ScaleRetryLimit: core.Int64Ptr(int64(3)),
			}
			jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateJobOptions := &codeenginev2.UpdateJobOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-job"),
				IfMatch: core.StringPtr("testString"),
				Job: jobPatchModelAsPatch,
			}

			job, response, err := codeEngineService.UpdateJob(updateJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`ListJobRuns - List job runs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListJobRuns(listJobRunsOptions *ListJobRunsOptions) with pagination`, func(){
			listJobRunsOptions := &codeenginev2.ListJobRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				JobName: core.StringPtr("my-job"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listJobRunsOptions.Start = nil
			listJobRunsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.JobRun
			for {
				jobRunList, response, err := codeEngineService.ListJobRuns(listJobRunsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(jobRunList).ToNot(BeNil())
				allResults = append(allResults, jobRunList.JobRuns...)

				listJobRunsOptions.Start, err = jobRunList.GetNextStart()
				Expect(err).To(BeNil())

				if listJobRunsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListJobRuns(listJobRunsOptions *ListJobRunsOptions) using JobRunsPager`, func(){
			listJobRunsOptions := &codeenginev2.ListJobRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				JobName: core.StringPtr("my-job"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewJobRunsPager(listJobRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.JobRun
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewJobRunsPager(listJobRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListJobRuns() returned a total of %d item(s) using JobRunsPager.\n", len(allResults))
		})
	})

	Describe(`CreateJobRun - Create a job run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateJobRun(createJobRunOptions *CreateJobRunOptions)`, func() {
			envVarPrototypeModel := &codeenginev2.EnvVarPrototype{
				Key: core.StringPtr("MY_VARIABLE"),
				Name: core.StringPtr("SOME"),
				Prefix: core.StringPtr("PREFIX_"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("literal"),
				Value: core.StringPtr("VALUE"),
			}

			volumeMountPrototypeModel := &codeenginev2.VolumeMountPrototype{
				MountPath: core.StringPtr("/app"),
				Name: core.StringPtr("codeengine-mount-b69u90"),
				Reference: core.StringPtr("my-secret"),
				Type: core.StringPtr("secret"),
			}

			createJobRunOptions := &codeenginev2.CreateJobRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				ImageReference: core.StringPtr("icr.io/codeengine/helloworld"),
				ImageSecret: core.StringPtr("my-secret"),
				JobName: core.StringPtr("my-job"),
				Name: core.StringPtr("my-job-run"),
				RunArguments: []string{"testString"},
				RunAsUser: core.Int64Ptr(int64(1001)),
				RunCommands: []string{"testString"},
				RunEnvVariables: []codeenginev2.EnvVarPrototype{*envVarPrototypeModel},
				RunMode: core.StringPtr("task"),
				RunServiceAccount: core.StringPtr("default"),
				RunVolumeMounts: []codeenginev2.VolumeMountPrototype{*volumeMountPrototypeModel},
				ScaleArraySpec: core.StringPtr("1-5,7-8,10"),
				ScaleCpuLimit: core.StringPtr("1"),
				ScaleEphemeralStorageLimit: core.StringPtr("4G"),
				ScaleMaxExecutionTime: core.Int64Ptr(int64(7200)),
				ScaleMemoryLimit: core.StringPtr("4G"),
				ScaleRetryLimit: core.Int64Ptr(int64(3)),
			}

			jobRun, response, err := codeEngineService.CreateJobRun(createJobRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobRun).ToNot(BeNil())
		})
	})

	Describe(`GetJobRun - Get a job run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetJobRun(getJobRunOptions *GetJobRunOptions)`, func() {
			getJobRunOptions := &codeenginev2.GetJobRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-job-run"),
			}

			jobRun, response, err := codeEngineService.GetJobRun(getJobRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobRun).ToNot(BeNil())
		})
	})

	Describe(`ListBindings - List bindings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBindings(listBindingsOptions *ListBindingsOptions) with pagination`, func(){
			listBindingsOptions := &codeenginev2.ListBindingsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listBindingsOptions.Start = nil
			listBindingsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.Binding
			for {
				bindingList, response, err := codeEngineService.ListBindings(listBindingsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(bindingList).ToNot(BeNil())
				allResults = append(allResults, bindingList.Bindings...)

				listBindingsOptions.Start, err = bindingList.GetNextStart()
				Expect(err).To(BeNil())

				if listBindingsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListBindings(listBindingsOptions *ListBindingsOptions) using BindingsPager`, func(){
			listBindingsOptions := &codeenginev2.ListBindingsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewBindingsPager(listBindingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.Binding
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewBindingsPager(listBindingsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListBindings() returned a total of %d item(s) using BindingsPager.\n", len(allResults))
		})
	})

	Describe(`CreateBinding - Create a binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBinding(createBindingOptions *CreateBindingOptions)`, func() {
			componentRefModel := &codeenginev2.ComponentRef{
				Name: core.StringPtr("my-app-1"),
				ResourceType: core.StringPtr("app_v2"),
			}

			createBindingOptions := &codeenginev2.CreateBindingOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Component: componentRefModel,
				Prefix: core.StringPtr("MY_COS"),
				SecretName: core.StringPtr("my-service-access"),
			}

			binding, response, err := codeEngineService.CreateBinding(createBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(binding).ToNot(BeNil())
		})
	})

	Describe(`GetBinding - Get a binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBinding(getBindingOptions *GetBindingOptions)`, func() {
			getBindingOptions := &codeenginev2.GetBindingOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				ID: core.StringPtr("a172ced-b5f21bc-71ba50c-1638604"),
			}

			binding, response, err := codeEngineService.GetBinding(getBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(binding).ToNot(BeNil())
		})
	})

	Describe(`ListBuilds - List builds`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBuilds(listBuildsOptions *ListBuildsOptions) with pagination`, func(){
			listBuildsOptions := &codeenginev2.ListBuildsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listBuildsOptions.Start = nil
			listBuildsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.Build
			for {
				buildList, response, err := codeEngineService.ListBuilds(listBuildsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(buildList).ToNot(BeNil())
				allResults = append(allResults, buildList.Builds...)

				listBuildsOptions.Start, err = buildList.GetNextStart()
				Expect(err).To(BeNil())

				if listBuildsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListBuilds(listBuildsOptions *ListBuildsOptions) using BuildsPager`, func(){
			listBuildsOptions := &codeenginev2.ListBuildsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewBuildsPager(listBuildsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.Build
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewBuildsPager(listBuildsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListBuilds() returned a total of %d item(s) using BuildsPager.\n", len(allResults))
		})
	})

	Describe(`CreateBuild - Create a build`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBuild(createBuildOptions *CreateBuildOptions)`, func() {
			createBuildOptions := &codeenginev2.CreateBuildOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build"),
				OutputImage: core.StringPtr("private.de.icr.io/icr_namespace/image-name"),
				OutputSecret: core.StringPtr("ce-auto-icr-private-eu-de"),
				StrategyType: core.StringPtr("dockerfile"),
				SourceContextDir: core.StringPtr("some/subfolder"),
				SourceRevision: core.StringPtr("main"),
				SourceSecret: core.StringPtr("testString"),
				SourceType: core.StringPtr("git"),
				SourceURL: core.StringPtr("https://github.com/IBM/CodeEngine"),
				StrategySize: core.StringPtr("medium"),
				StrategySpecFile: core.StringPtr("Dockerfile"),
				Timeout: core.Int64Ptr(int64(600)),
			}

			build, response, err := codeEngineService.CreateBuild(createBuildOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(build).ToNot(BeNil())
		})
	})

	Describe(`GetBuild - Get a build`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBuild(getBuildOptions *GetBuildOptions)`, func() {
			getBuildOptions := &codeenginev2.GetBuildOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build"),
			}

			build, response, err := codeEngineService.GetBuild(getBuildOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(build).ToNot(BeNil())
		})
	})

	Describe(`UpdateBuild - Update a build`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBuild(updateBuildOptions *UpdateBuildOptions)`, func() {
			buildPatchModel := &codeenginev2.BuildPatch{
				OutputImage: core.StringPtr("private.de.icr.io/icr_namespace/image-name"),
				OutputSecret: core.StringPtr("ce-auto-icr-private-eu-de"),
				SourceContextDir: core.StringPtr("some/subfolder"),
				SourceRevision: core.StringPtr("main"),
				SourceSecret: core.StringPtr("testString"),
				SourceType: core.StringPtr("git"),
				SourceURL: core.StringPtr("https://github.com/IBM/CodeEngine"),
				StrategySize: core.StringPtr("medium"),
				StrategySpecFile: core.StringPtr("Dockerfile"),
				StrategyType: core.StringPtr("dockerfile"),
				Timeout: core.Int64Ptr(int64(600)),
			}
			buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBuildOptions := &codeenginev2.UpdateBuildOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build"),
				IfMatch: core.StringPtr("testString"),
				Build: buildPatchModelAsPatch,
			}

			build, response, err := codeEngineService.UpdateBuild(updateBuildOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(build).ToNot(BeNil())
		})
	})

	Describe(`ListBuildRuns - List build runs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBuildRuns(listBuildRunsOptions *ListBuildRunsOptions) with pagination`, func(){
			listBuildRunsOptions := &codeenginev2.ListBuildRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				BuildName: core.StringPtr("my-build"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listBuildRunsOptions.Start = nil
			listBuildRunsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.BuildRun
			for {
				buildRunList, response, err := codeEngineService.ListBuildRuns(listBuildRunsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(buildRunList).ToNot(BeNil())
				allResults = append(allResults, buildRunList.BuildRuns...)

				listBuildRunsOptions.Start, err = buildRunList.GetNextStart()
				Expect(err).To(BeNil())

				if listBuildRunsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListBuildRuns(listBuildRunsOptions *ListBuildRunsOptions) using BuildRunsPager`, func(){
			listBuildRunsOptions := &codeenginev2.ListBuildRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				BuildName: core.StringPtr("my-build"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewBuildRunsPager(listBuildRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.BuildRun
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewBuildRunsPager(listBuildRunsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListBuildRuns() returned a total of %d item(s) using BuildRunsPager.\n", len(allResults))
		})
	})

	Describe(`CreateBuildRun - Create a build run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBuildRun(createBuildRunOptions *CreateBuildRunOptions)`, func() {
			createBuildRunOptions := &codeenginev2.CreateBuildRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				BuildName: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				OutputImage: core.StringPtr("private.de.icr.io/icr_namespace/image-name"),
				OutputSecret: core.StringPtr("ce-auto-icr-private-eu-de"),
				ServiceAccount: core.StringPtr("default"),
				SourceContextDir: core.StringPtr("some/subfolder"),
				SourceRevision: core.StringPtr("main"),
				SourceSecret: core.StringPtr("testString"),
				SourceType: core.StringPtr("git"),
				SourceURL: core.StringPtr("https://github.com/IBM/CodeEngine"),
				StrategySize: core.StringPtr("medium"),
				StrategySpecFile: core.StringPtr("Dockerfile"),
				StrategyType: core.StringPtr("dockerfile"),
				Timeout: core.Int64Ptr(int64(600)),
			}

			buildRun, response, err := codeEngineService.CreateBuildRun(createBuildRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(buildRun).ToNot(BeNil())
		})
	})

	Describe(`GetBuildRun - Get a build run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBuildRun(getBuildRunOptions *GetBuildRunOptions)`, func() {
			getBuildRunOptions := &codeenginev2.GetBuildRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build-run"),
			}

			buildRun, response, err := codeEngineService.GetBuildRun(getBuildRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(buildRun).ToNot(BeNil())
		})
	})

	Describe(`ListConfigMaps - List config maps`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigMaps(listConfigMapsOptions *ListConfigMapsOptions) with pagination`, func(){
			listConfigMapsOptions := &codeenginev2.ListConfigMapsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listConfigMapsOptions.Start = nil
			listConfigMapsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.ConfigMap
			for {
				configMapList, response, err := codeEngineService.ListConfigMaps(listConfigMapsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(configMapList).ToNot(BeNil())
				allResults = append(allResults, configMapList.ConfigMaps...)

				listConfigMapsOptions.Start, err = configMapList.GetNextStart()
				Expect(err).To(BeNil())

				if listConfigMapsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListConfigMaps(listConfigMapsOptions *ListConfigMapsOptions) using ConfigMapsPager`, func(){
			listConfigMapsOptions := &codeenginev2.ListConfigMapsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewConfigMapsPager(listConfigMapsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.ConfigMap
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewConfigMapsPager(listConfigMapsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListConfigMaps() returned a total of %d item(s) using ConfigMapsPager.\n", len(allResults))
		})
	})

	Describe(`CreateConfigMap - Create a config map`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateConfigMap(createConfigMapOptions *CreateConfigMapOptions)`, func() {
			createConfigMapOptions := &codeenginev2.CreateConfigMapOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-config-map"),
				Data: map[string]string{"key1": "testString"},
			}

			configMap, response, err := codeEngineService.CreateConfigMap(createConfigMapOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(configMap).ToNot(BeNil())
		})
	})

	Describe(`GetConfigMap - Get a config map`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfigMap(getConfigMapOptions *GetConfigMapOptions)`, func() {
			getConfigMapOptions := &codeenginev2.GetConfigMapOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-config-map"),
			}

			configMap, response, err := codeEngineService.GetConfigMap(getConfigMapOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configMap).ToNot(BeNil())
		})
	})

	Describe(`ReplaceConfigMap - Update a config map`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceConfigMap(replaceConfigMapOptions *ReplaceConfigMapOptions)`, func() {
			replaceConfigMapOptions := &codeenginev2.ReplaceConfigMapOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-config-map"),
				IfMatch: core.StringPtr("testString"),
				Data: map[string]string{"key1": "testString"},
			}

			configMap, response, err := codeEngineService.ReplaceConfigMap(replaceConfigMapOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configMap).ToNot(BeNil())
		})
	})

	Describe(`ListSecrets - List secrets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSecrets(listSecretsOptions *ListSecretsOptions) with pagination`, func(){
			listSecretsOptions := &codeenginev2.ListSecretsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
				Start: core.StringPtr("testString"),
			}

			listSecretsOptions.Start = nil
			listSecretsOptions.Limit = core.Int64Ptr(1)

			var allResults []codeenginev2.Secret
			for {
				secretList, response, err := codeEngineService.ListSecrets(listSecretsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(secretList).ToNot(BeNil())
				allResults = append(allResults, secretList.Secrets...)

				listSecretsOptions.Start, err = secretList.GetNextStart()
				Expect(err).To(BeNil())

				if listSecretsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSecrets(listSecretsOptions *ListSecretsOptions) using SecretsPager`, func(){
			listSecretsOptions := &codeenginev2.ListSecretsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := codeEngineService.NewSecretsPager(listSecretsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []codeenginev2.Secret
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = codeEngineService.NewSecretsPager(listSecretsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSecrets() returned a total of %d item(s) using SecretsPager.\n", len(allResults))
		})
	})

	Describe(`CreateSecret - Create a secret`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSecret(createSecretOptions *CreateSecretOptions)`, func() {
			secretDataModel := &codeenginev2.SecretDataSSHSecretData{
				SshKey: core.StringPtr("testString"),
				KnownHosts: core.StringPtr("testString"),
			}
			secretDataModel.SetProperty("foo", core.StringPtr("testString"))

			resourceKeyRefPrototypeModel := &codeenginev2.ResourceKeyRefPrototype{
				ID: core.StringPtr("4e49b3e0-27a8-48d2-a784-c7ee48bb863b"),
			}

			roleRefPrototypeModel := &codeenginev2.RoleRefPrototype{
				Crn: core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Writer"),
			}

			serviceInstanceRefPrototypeModel := &codeenginev2.ServiceInstanceRefPrototype{
				ID: core.StringPtr("4e49b3e0-27a8-48d2-a784-c7ee48bb863b"),
			}

			serviceIdRefModel := &codeenginev2.ServiceIDRef{
				Crn: core.StringPtr("testString"),
			}

			serviceAccessSecretPrototypePropsModel := &codeenginev2.ServiceAccessSecretPrototypeProps{
				ResourceKey: resourceKeyRefPrototypeModel,
				Role: roleRefPrototypeModel,
				ServiceInstance: serviceInstanceRefPrototypeModel,
				Serviceid: serviceIdRefModel,
			}

			createSecretOptions := &codeenginev2.CreateSecretOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Format: core.StringPtr("generic"),
				Name: core.StringPtr("my-secret"),
				Data: secretDataModel,
				ServiceAccess: serviceAccessSecretPrototypePropsModel,
			}

			secret, response, err := codeEngineService.CreateSecret(createSecretOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(secret).ToNot(BeNil())
		})
	})

	Describe(`GetSecret - Get a secret`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecret(getSecretOptions *GetSecretOptions)`, func() {
			getSecretOptions := &codeenginev2.GetSecretOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-secret"),
			}

			secret, response, err := codeEngineService.GetSecret(getSecretOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(secret).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSecret - Update a secret`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSecret(replaceSecretOptions *ReplaceSecretOptions)`, func() {
			secretDataModel := &codeenginev2.SecretDataSSHSecretData{
				SshKey: core.StringPtr("testString"),
				KnownHosts: core.StringPtr("testString"),
			}
			secretDataModel.SetProperty("foo", core.StringPtr("testString"))

			replaceSecretOptions := &codeenginev2.ReplaceSecretOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-secret"),
				IfMatch: core.StringPtr("testString"),
				Format: core.StringPtr("generic"),
				Data: secretDataModel,
			}

			secret, response, err := codeEngineService.ReplaceSecret(replaceSecretOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(secret).ToNot(BeNil())
		})
	})

	Describe(`DeleteProject - Delete a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
			deleteProjectOptions := &codeenginev2.DeleteProjectOptions{
				ID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
			}

			response, err := codeEngineService.DeleteProject(deleteProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteApp - Delete an application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApp(deleteAppOptions *DeleteAppOptions)`, func() {
			deleteAppOptions := &codeenginev2.DeleteAppOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-app"),
			}

			response, err := codeEngineService.DeleteApp(deleteAppOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteAppRevision - Delete an application revision`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAppRevision(deleteAppRevisionOptions *DeleteAppRevisionOptions)`, func() {
			deleteAppRevisionOptions := &codeenginev2.DeleteAppRevisionOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				AppName: core.StringPtr("my-app"),
				Name: core.StringPtr("my-app-00001"),
			}

			response, err := codeEngineService.DeleteAppRevision(deleteAppRevisionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteJob - Delete a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
			deleteJobOptions := &codeenginev2.DeleteJobOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-job"),
			}

			response, err := codeEngineService.DeleteJob(deleteJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteJobRun - Delete a job run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteJobRun(deleteJobRunOptions *DeleteJobRunOptions)`, func() {
			deleteJobRunOptions := &codeenginev2.DeleteJobRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-job-run"),
			}

			response, err := codeEngineService.DeleteJobRun(deleteJobRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteBinding - Delete a binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBinding(deleteBindingOptions *DeleteBindingOptions)`, func() {
			deleteBindingOptions := &codeenginev2.DeleteBindingOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				ID: core.StringPtr("a172ced-b5f21bc-71ba50c-1638604"),
			}

			response, err := codeEngineService.DeleteBinding(deleteBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteBuild - Delete a build`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBuild(deleteBuildOptions *DeleteBuildOptions)`, func() {
			deleteBuildOptions := &codeenginev2.DeleteBuildOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build"),
			}

			response, err := codeEngineService.DeleteBuild(deleteBuildOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteBuildRun - Delete a build run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBuildRun(deleteBuildRunOptions *DeleteBuildRunOptions)`, func() {
			deleteBuildRunOptions := &codeenginev2.DeleteBuildRunOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-build-run"),
			}

			response, err := codeEngineService.DeleteBuildRun(deleteBuildRunOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteConfigMap - Delete a config map`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfigMap(deleteConfigMapOptions *DeleteConfigMapOptions)`, func() {
			deleteConfigMapOptions := &codeenginev2.DeleteConfigMapOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-config-map"),
			}

			response, err := codeEngineService.DeleteConfigMap(deleteConfigMapOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`DeleteSecret - Delete a secret`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSecret(deleteSecretOptions *DeleteSecretOptions)`, func() {
			deleteSecretOptions := &codeenginev2.DeleteSecretOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Name: core.StringPtr("my-secret"),
			}

			response, err := codeEngineService.DeleteSecret(deleteSecretOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
