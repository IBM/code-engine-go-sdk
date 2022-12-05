package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
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
	fmt.Printf("Using IAM endpoint: '%s'\n", iamEndpoint)

	rcEndpoint := "https://resource-controller.cloud.ibm.com"
	if len(os.Getenv("RESOURCECONTROLLER_ENDPOINT")) > 0 {
		rcEndpoint = os.Getenv("RESOURCECONTROLLER_ENDPOINT")
	}
	fmt.Printf("Using Resource Controller endpoint: '%s'\n", rcEndpoint)

	accountID := os.Getenv("CE_ACCOUNT_ID")
	fmt.Printf("Using account: '%s'\n", accountID)

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey:       os.Getenv("CE_API_KEY"),
		ClientId:     "bx",
		ClientSecret: "bx",
		URL:          iamEndpoint,
	}

	// Cleanup projects that have been created by prior runs
	cleanupProjectReclamations(authenticator, rcEndpoint, accountID)

	codeEngineApiEndpoint := "https://" + os.Getenv("CE_API_HOST") + "/v2"
	fmt.Printf("Using Code Engine API endpoint: '%s'\n", codeEngineApiEndpoint)

	// Setup a Code Engine client
	codeEngineServiceOptions := &codeenginev2.CodeEngineV2Options{
		Authenticator: authenticator,
		URL:           codeEngineApiEndpoint,
	}
	codeEngineService, err := codeenginev2.NewCodeEngineV2UsingExternalConfig(codeEngineServiceOptions)
	if err != nil {
		fmt.Printf("NewCodeEngineV2UsingExternalConfig error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// List Code Engine projects using the Code Engine Client
	listResult, _, err := codeEngineService.ListProjects(&codeenginev2.ListProjectsOptions{})
	if err != nil {
		fmt.Printf("ListProjects error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))

	// Create a new Code Engine project using the Code Engine Client
	projectName := "project-sdk-go-e2e--crud--" + time.Now().Format("060102-150405")
	createdProject, _, err := codeEngineService.CreateProject(&codeenginev2.CreateProjectOptions{
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
		obtainedProject, _, err := codeEngineService.GetProject(getProjectOptions)
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

	resp, err := codeEngineService.DeleteProject(deleteProjectOptions)
	if err != nil {
		fmt.Printf("DeleteProject error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted project: '%d'\n", resp.StatusCode)

	listResult, _, err = codeEngineService.ListProjects(&codeenginev2.ListProjectsOptions{})
	if err != nil {
		fmt.Printf("ListProjects error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))

}

func cleanupProjectReclamations(authenticator *core.IamAuthenticator, rcEndpoint string, accountID string) {

	// Init the resource controller client to cleanup leftovers
	resourceControllerServiceOptions := &resourcecontrollerv2.ResourceControllerV2Options{
		Authenticator: authenticator,
		URL:           rcEndpoint,
	}
	resourceControllerService, rcInitErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(resourceControllerServiceOptions)
	if rcInitErr != nil {
		fmt.Printf("NewResourceControllerV2UsingExternalConfig error: %s\n", rcInitErr.Error())
		os.Exit(1)
		return
	}

	// 1 list all reclamations
	listReclamationsOptions := resourceControllerService.NewListReclamationsOptions()
	listReclamationsOptions = listReclamationsOptions.SetAccountID(accountID)
	reclamationsList, _, rcErr := resourceControllerService.ListReclamations(listReclamationsOptions)
	if rcErr != nil {
		fmt.Printf("ListReclamations error: %s\n", rcErr.Error())
		os.Exit(1)
	}
	fmt.Printf("Found %d reclamations\n", len(reclamationsList.Resources))

	// 2 iterate over all reclamations
	for _, reclamation := range reclamationsList.Resources {

		// examine whether we are dealing with a code engine instance
		if !strings.Contains(*reclamation.EntityCRN, ":public:codeengine:") {
			continue
		}

		// examine the data of the reclamation and check whether it is older than 5 min
		if *reclamation.State != "SCHEDULED" {
			continue
		}

		created, parserErr := time.Parse(time.RFC3339, reclamation.CreatedAt.String())
		if parserErr != nil {
			fmt.Printf("faild to parse '%s' error: %s\n", reclamation.CreatedAt, parserErr.Error())
			os.Exit(1)
		}

		// 3 examine the data of the reclamation and check whether it is older than 5 min
		if created.After(time.Now().Add(-5 * time.Minute)) {
			continue
		}

		// 4 delete it, if it is too old
		fmt.Printf("Deleting reclamation: '%s'\n", *reclamation.ID)
		runReclamationActionOptions := resourceControllerService.NewRunReclamationActionOptions(
			*reclamation.ID,
			"reclaim",
		)
		reclamation, _, reclaimErr := resourceControllerService.RunReclamationAction(runReclamationActionOptions)
		if reclaimErr != nil {
			fmt.Printf("RunReclamationAction error: %s\n", reclaimErr.Error())
			os.Exit(1)
		}
		b, marshalErr := json.MarshalIndent(reclamation, "", "  ")
		if marshalErr != nil {
			fmt.Printf("faild to print reclamation: %s\n", parserErr.Error())
			os.Exit(1)
		}
		fmt.Println(string(b))

	}

	fmt.Printf("Done cleaning up!\n")
}
