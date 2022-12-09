// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Code Engine service.
//
// The following configuration properties are assumed to be defined:
// CODE_ENGINE_URL=<service base url>
// CODE_ENGINE_AUTH_TYPE=iam
// CODE_ENGINE_APIKEY=<IAM apikey>
// CODE_ENGINE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`CodeEngineV2 Examples Tests`, func() {

	const externalConfigFile = "../code_engine_v2.env"

	var (
		codeEngineService *codeenginev2.CodeEngineV2
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(codeenginev2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			codeEngineServiceOptions := &codeenginev2.CodeEngineV2Options{}

			codeEngineService, err = codeenginev2.NewCodeEngineV2UsingExternalConfig(codeEngineServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(codeEngineService).ToNot(BeNil())
		})
	})

	Describe(`CodeEngineV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProjects request example`, func() {
			fmt.Println("\nListProjects() result:")
			// begin-list_projects
			listProjectsOptions := &codeenginev2.ListProjectsOptions{
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewProjectsPager(listProjectsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.Project
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_projects
		})
		It(`CreateProject request example`, func() {
			fmt.Println("\nCreateProject() result:")
			// begin-create_project

			createProjectOptions := codeEngineService.NewCreateProjectOptions(
				"my-project",
			)

			project, response, err := codeEngineService.CreateProject(createProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-create_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(project).ToNot(BeNil())
		})
		It(`GetProject request example`, func() {
			fmt.Println("\nGetProject() result:")
			// begin-get_project

			getProjectOptions := codeEngineService.NewGetProjectOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
			)

			project, response, err := codeEngineService.GetProject(getProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-get_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
		It(`ListApps request example`, func() {
			fmt.Println("\nListApps() result:")
			// begin-list_apps
			listAppsOptions := &codeenginev2.ListAppsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewAppsPager(listAppsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.App
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_apps
		})
		It(`CreateApp request example`, func() {
			fmt.Println("\nCreateApp() result:")
			// begin-create_app

			createAppOptions := codeEngineService.NewCreateAppOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"icr.io/codeengine/helloworld",
				"my-app",
			)

			app, response, err := codeEngineService.CreateApp(createAppOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(app, "", "  ")
			fmt.Println(string(b))

			// end-create_app

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(app).ToNot(BeNil())
		})
		It(`GetApp request example`, func() {
			fmt.Println("\nGetApp() result:")
			// begin-get_app

			getAppOptions := codeEngineService.NewGetAppOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-app",
			)

			app, response, err := codeEngineService.GetApp(getAppOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(app, "", "  ")
			fmt.Println(string(b))

			// end-get_app

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(app).ToNot(BeNil())
		})
		It(`UpdateApp request example`, func() {
			fmt.Println("\nUpdateApp() result:")
			// begin-update_app

			appPatchModel := &codeenginev2.AppPatch{
			}
			appPatchModelAsPatch, asPatchErr := appPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateAppOptions := codeEngineService.NewUpdateAppOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-app",
				"testString",
				appPatchModelAsPatch,
			)

			app, response, err := codeEngineService.UpdateApp(updateAppOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(app, "", "  ")
			fmt.Println(string(b))

			// end-update_app

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(app).ToNot(BeNil())
		})
		It(`ListAppRevisions request example`, func() {
			fmt.Println("\nListAppRevisions() result:")
			// begin-list_app_revisions
			listAppRevisionsOptions := &codeenginev2.ListAppRevisionsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				AppName: core.StringPtr("my-app"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewAppRevisionsPager(listAppRevisionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.AppRevision
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_app_revisions
		})
		It(`GetAppRevision request example`, func() {
			fmt.Println("\nGetAppRevision() result:")
			// begin-get_app_revision

			getAppRevisionOptions := codeEngineService.NewGetAppRevisionOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-app",
				"my-app-001",
			)

			appRevision, response, err := codeEngineService.GetAppRevision(getAppRevisionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(appRevision, "", "  ")
			fmt.Println(string(b))

			// end-get_app_revision

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(appRevision).ToNot(BeNil())
		})
		It(`ListJobs request example`, func() {
			fmt.Println("\nListJobs() result:")
			// begin-list_jobs
			listJobsOptions := &codeenginev2.ListJobsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewJobsPager(listJobsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.Job
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_jobs
		})
		It(`CreateJob request example`, func() {
			fmt.Println("\nCreateJob() result:")
			// begin-create_job

			createJobOptions := codeEngineService.NewCreateJobOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"icr.io/codeengine/helloworld",
				"my-job",
			)

			job, response, err := codeEngineService.CreateJob(createJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-create_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(job).ToNot(BeNil())
		})
		It(`GetJob request example`, func() {
			fmt.Println("\nGetJob() result:")
			// begin-get_job

			getJobOptions := codeEngineService.NewGetJobOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-job",
			)

			job, response, err := codeEngineService.GetJob(getJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-get_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
		It(`UpdateJob request example`, func() {
			fmt.Println("\nUpdateJob() result:")
			// begin-update_job

			jobPatchModel := &codeenginev2.JobPatch{
			}
			jobPatchModelAsPatch, asPatchErr := jobPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateJobOptions := codeEngineService.NewUpdateJobOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-job",
				"testString",
				jobPatchModelAsPatch,
			)

			job, response, err := codeEngineService.UpdateJob(updateJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-update_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
		It(`ListJobRuns request example`, func() {
			fmt.Println("\nListJobRuns() result:")
			// begin-list_job_runs
			listJobRunsOptions := &codeenginev2.ListJobRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				JobName: core.StringPtr("my-job"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewJobRunsPager(listJobRunsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.JobRun
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_job_runs
		})
		It(`CreateJobRun request example`, func() {
			fmt.Println("\nCreateJobRun() result:")
			// begin-create_job_run

			createJobRunOptions := codeEngineService.NewCreateJobRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
			)

			jobRun, response, err := codeEngineService.CreateJobRun(createJobRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobRun, "", "  ")
			fmt.Println(string(b))

			// end-create_job_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobRun).ToNot(BeNil())
		})
		It(`GetJobRun request example`, func() {
			fmt.Println("\nGetJobRun() result:")
			// begin-get_job_run

			getJobRunOptions := codeEngineService.NewGetJobRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-job",
			)

			jobRun, response, err := codeEngineService.GetJobRun(getJobRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobRun, "", "  ")
			fmt.Println(string(b))

			// end-get_job_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobRun).ToNot(BeNil())
		})
		It(`ListBuilds request example`, func() {
			fmt.Println("\nListBuilds() result:")
			// begin-list_builds
			listBuildsOptions := &codeenginev2.ListBuildsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewBuildsPager(listBuildsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.Build
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_builds
		})
		It(`CreateBuild request example`, func() {
			fmt.Println("\nCreateBuild() result:")
			// begin-create_build

			createBuildOptions := codeEngineService.NewCreateBuildOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build",
				"private.de.icr.io/icr_namespace/image-name",
				"ce-auto-icr-private-eu-de",
				"https://github.com/IBM/CodeEngine",
				"dockerfile",
			)

			build, response, err := codeEngineService.CreateBuild(createBuildOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(build, "", "  ")
			fmt.Println(string(b))

			// end-create_build

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(build).ToNot(BeNil())
		})
		It(`GetBuild request example`, func() {
			fmt.Println("\nGetBuild() result:")
			// begin-get_build

			getBuildOptions := codeEngineService.NewGetBuildOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build",
			)

			build, response, err := codeEngineService.GetBuild(getBuildOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(build, "", "  ")
			fmt.Println(string(b))

			// end-get_build

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(build).ToNot(BeNil())
		})
		It(`UpdateBuild request example`, func() {
			fmt.Println("\nUpdateBuild() result:")
			// begin-update_build

			buildPatchModel := &codeenginev2.BuildPatch{
			}
			buildPatchModelAsPatch, asPatchErr := buildPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBuildOptions := codeEngineService.NewUpdateBuildOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build",
				"testString",
				buildPatchModelAsPatch,
			)

			build, response, err := codeEngineService.UpdateBuild(updateBuildOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(build, "", "  ")
			fmt.Println(string(b))

			// end-update_build

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(build).ToNot(BeNil())
		})
		It(`ListBuildRuns request example`, func() {
			fmt.Println("\nListBuildRuns() result:")
			// begin-list_build_runs
			listBuildRunsOptions := &codeenginev2.ListBuildRunsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				BuildName: core.StringPtr("my-build"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewBuildRunsPager(listBuildRunsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.BuildRun
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_build_runs
		})
		It(`CreateBuildRun request example`, func() {
			fmt.Println("\nCreateBuildRun() result:")
			// begin-create_build_run

			createBuildRunOptions := codeEngineService.NewCreateBuildRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
			)

			buildRun, response, err := codeEngineService.CreateBuildRun(createBuildRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(buildRun, "", "  ")
			fmt.Println(string(b))

			// end-create_build_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(buildRun).ToNot(BeNil())
		})
		It(`GetBuildRun request example`, func() {
			fmt.Println("\nGetBuildRun() result:")
			// begin-get_build_run

			getBuildRunOptions := codeEngineService.NewGetBuildRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build-run",
			)

			buildRun, response, err := codeEngineService.GetBuildRun(getBuildRunOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(buildRun, "", "  ")
			fmt.Println(string(b))

			// end-get_build_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(buildRun).ToNot(BeNil())
		})
		It(`ListConfigMaps request example`, func() {
			fmt.Println("\nListConfigMaps() result:")
			// begin-list_config_maps
			listConfigMapsOptions := &codeenginev2.ListConfigMapsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewConfigMapsPager(listConfigMapsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.ConfigMap
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_config_maps
		})
		It(`CreateConfigMap request example`, func() {
			fmt.Println("\nCreateConfigMap() result:")
			// begin-create_config_map

			createConfigMapOptions := codeEngineService.NewCreateConfigMapOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-configmap",
			)

			configMap, response, err := codeEngineService.CreateConfigMap(createConfigMapOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configMap, "", "  ")
			fmt.Println(string(b))

			// end-create_config_map

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(configMap).ToNot(BeNil())
		})
		It(`GetConfigMap request example`, func() {
			fmt.Println("\nGetConfigMap() result:")
			// begin-get_config_map

			getConfigMapOptions := codeEngineService.NewGetConfigMapOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-config-map",
			)

			configMap, response, err := codeEngineService.GetConfigMap(getConfigMapOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configMap, "", "  ")
			fmt.Println(string(b))

			// end-get_config_map

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configMap).ToNot(BeNil())
		})
		It(`ReplaceConfigMap request example`, func() {
			fmt.Println("\nReplaceConfigMap() result:")
			// begin-replace_config_map

			replaceConfigMapOptions := codeEngineService.NewReplaceConfigMapOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-config-map",
				"testString",
			)

			configMap, response, err := codeEngineService.ReplaceConfigMap(replaceConfigMapOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configMap, "", "  ")
			fmt.Println(string(b))

			// end-replace_config_map

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configMap).ToNot(BeNil())
		})
		It(`ListSecrets request example`, func() {
			fmt.Println("\nListSecrets() result:")
			// begin-list_secrets
			listSecretsOptions := &codeenginev2.ListSecretsOptions{
				ProjectID: core.StringPtr("15314cc3-85b4-4338-903f-c28cdee6d005"),
				Limit: core.Int64Ptr(int64(100)),
			}

			pager, err := codeEngineService.NewSecretsPager(listSecretsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []codeenginev2.Secret
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_secrets
		})
		It(`CreateSecret request example`, func() {
			fmt.Println("\nCreateSecret() result:")
			// begin-create_secret

			createSecretOptions := codeEngineService.NewCreateSecretOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"generic",
				"my-secret",
			)

			secret, response, err := codeEngineService.CreateSecret(createSecretOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(secret, "", "  ")
			fmt.Println(string(b))

			// end-create_secret

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(secret).ToNot(BeNil())
		})
		It(`GetSecret request example`, func() {
			fmt.Println("\nGetSecret() result:")
			// begin-get_secret

			getSecretOptions := codeEngineService.NewGetSecretOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-secret",
			)

			secret, response, err := codeEngineService.GetSecret(getSecretOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(secret, "", "  ")
			fmt.Println(string(b))

			// end-get_secret

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(secret).ToNot(BeNil())
		})
		It(`ReplaceSecret request example`, func() {
			fmt.Println("\nReplaceSecret() result:")
			// begin-replace_secret

			replaceSecretOptions := codeEngineService.NewReplaceSecretOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-secret",
				"testString",
			)

			secret, response, err := codeEngineService.ReplaceSecret(replaceSecretOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(secret, "", "  ")
			fmt.Println(string(b))

			// end-replace_secret

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(secret).ToNot(BeNil())
		})
		It(`DeleteProject request example`, func() {
			// begin-delete_project

			deleteProjectOptions := codeEngineService.NewDeleteProjectOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
			)

			response, err := codeEngineService.DeleteProject(deleteProjectOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteProject(): %d\n", response.StatusCode)
			}

			// end-delete_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteApp request example`, func() {
			// begin-delete_app

			deleteAppOptions := codeEngineService.NewDeleteAppOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-app",
			)

			response, err := codeEngineService.DeleteApp(deleteAppOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteApp(): %d\n", response.StatusCode)
			}

			// end-delete_app

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteAppRevision request example`, func() {
			// begin-delete_app_revision

			deleteAppRevisionOptions := codeEngineService.NewDeleteAppRevisionOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-app",
				"my-app-001",
			)

			response, err := codeEngineService.DeleteAppRevision(deleteAppRevisionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteAppRevision(): %d\n", response.StatusCode)
			}

			// end-delete_app_revision

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteJob request example`, func() {
			// begin-delete_job

			deleteJobOptions := codeEngineService.NewDeleteJobOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-job",
			)

			response, err := codeEngineService.DeleteJob(deleteJobOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteJob(): %d\n", response.StatusCode)
			}

			// end-delete_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteJobRun request example`, func() {
			// begin-delete_job_run

			deleteJobRunOptions := codeEngineService.NewDeleteJobRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-job",
			)

			response, err := codeEngineService.DeleteJobRun(deleteJobRunOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteJobRun(): %d\n", response.StatusCode)
			}

			// end-delete_job_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteBuild request example`, func() {
			// begin-delete_build

			deleteBuildOptions := codeEngineService.NewDeleteBuildOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build",
			)

			response, err := codeEngineService.DeleteBuild(deleteBuildOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteBuild(): %d\n", response.StatusCode)
			}

			// end-delete_build

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteBuildRun request example`, func() {
			// begin-delete_build_run

			deleteBuildRunOptions := codeEngineService.NewDeleteBuildRunOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-build-run",
			)

			response, err := codeEngineService.DeleteBuildRun(deleteBuildRunOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteBuildRun(): %d\n", response.StatusCode)
			}

			// end-delete_build_run

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteConfigMap request example`, func() {
			// begin-delete_config_map

			deleteConfigMapOptions := codeEngineService.NewDeleteConfigMapOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-config-map",
			)

			response, err := codeEngineService.DeleteConfigMap(deleteConfigMapOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteConfigMap(): %d\n", response.StatusCode)
			}

			// end-delete_config_map

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteSecret request example`, func() {
			// begin-delete_secret

			deleteSecretOptions := codeEngineService.NewDeleteSecretOptions(
				"15314cc3-85b4-4338-903f-c28cdee6d005",
				"my-secret",
			)

			response, err := codeEngineService.DeleteSecret(deleteSecretOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteSecret(): %d\n", response.StatusCode)
			}

			// end-delete_secret

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})
})
