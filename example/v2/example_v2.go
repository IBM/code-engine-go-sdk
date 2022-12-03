package main

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	var (
		codeEngineService *codeenginev2.CodeEngineV2
	)

	// Validate environment
	requiredEnvs := []string{"CE_API_KEY", "CE_API_HOST", "CE_PROJECT_ID", "CE_ACCOUNT_ID"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			fmt.Printf("Environment variable %s must be set\n", env)
			os.Exit(1)
			return
		}
	}

	iamEndpoint := "https://iam.cloud.ibm.com"
	if len(os.Getenv("IAM_ENDPOINT")) > 0 {
		iamEndpoint = os.Getenv("IAM_ENDPOINT")
	}

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey:       os.Getenv("CE_API_KEY"),
		ClientId:     "bx",
		ClientSecret: "bx",
		URL:          iamEndpoint,
	}

	// Setup a Code Engine client
	ceClient, err := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
		Authenticator: authenticator,
		URL:           "https://" + os.Getenv("CE_API_HOST") + "/v2",
	})
	if err != nil {
		fmt.Printf("NewCodeEngineV2 error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// List Code Engine projects using the Code Engine Client
	listResult, _, err := ceClient.ListProjects(&codeenginev2.ListProjectsOptions{})
	if err != nil {
		fmt.Printf("ListProjects error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))

	// Create a new Code Engine project using the Code Engine Client
	projectName := "project-sdk-go-e2e--crud--" + time.Now().Format("060102-150405")
	createdProject, _, err := ceClient.CreateProject(&codeenginev2.CreateProjectOptions{
		Name: &projectName,
	})
	if err != nil {
		fmt.Printf("CreateProject error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created projects '%s' (guid: '%s').\n", *createdProject.Name, *createdProject.ID)

	//
	// Assume that the project creation takes some time
	for x := 0; x < 20; x++ {
		// sleep for 10 seconds and then try to fetch the project
		time.Sleep(10 * time.Second)

		getProjectOptions := codeEngineService.NewGetProjectOptions(
			*createdProject.ID,
		)
		obtainedProject, _, err := ceClient.GetProject(getProjectOptions)
		if err != nil {
			fmt.Printf("GetProject error: %s\n", err.Error())
			os.Exit(1)
			return
		}
		fmt.Printf("Obtained status of project '%s' (guid: '%s'): %s.\n", *obtainedProject.Name, *obtainedProject.ID, *obtainedProject.Status)
		if *obtainedProject.Status == "active" {
			break
		}
	}

	deleteProjectOptions := codeEngineService.NewDeleteProjectOptions(
		*createdProject.ID,
	)

	resp, err := ceClient.DeleteProject(deleteProjectOptions)
	if err != nil {
		fmt.Printf("DeleteProject error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted project: '%d'\n", resp.StatusCode)

	listResult, _, err = ceClient.ListProjects(&codeenginev2.ListProjectsOptions{})
	if err != nil {
		fmt.Printf("ListProjects error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))
}
