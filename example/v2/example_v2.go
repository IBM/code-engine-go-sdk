package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

type ResourceGroup struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Default bool   `json:"default,omitempty"`
}
type ResourceGroups struct {
	Resources []ResourceGroup `json:"resources,omitempty"`
}

func getDefaultResourceGroupId(accessToken string, resourceControllerEndpoint string, accountId string) (*string, error) {

	// build the request payload
	data := url.Values{}
	data.Set("account_id", accountId)
	data.Set("default", "true")

	// initialize the HTTP client
	client := &http.Client{}
	req, _ := http.NewRequest("GET", resourceControllerEndpoint+"/v2/resource_groups?"+data.Encode(), nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", `application/json`)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// perform the RC retrieval operation
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read the response body and convert it to a byte array
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// map the bytes to a defined struct
	resourceGroups := ResourceGroups{}
	err = json.Unmarshal(responseData, &resourceGroups)
	if err != nil {
		return nil, err
	}

	// for debugging purposes
	// fmt.Println("resourceGroups: " + resourceGroups)

	for _, resourceGroup := range resourceGroups.Resources {
		if resourceGroup.Default {
			fmt.Println("Identified resource group '" + resourceGroup.Name + "' as default")

			// assigning the resource group id
			resourceGroupId := resourceGroup.Id

			return &resourceGroupId, nil
		}
	}

	return nil, nil
}

func main() {

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

	accessToken, err := authenticator.GetToken()
	if err != nil {
		fmt.Printf("IAM GetToken error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	resourceControllerEndpoint := "https://resource-controller.cloud.ibm.com"
	if len(os.Getenv("RESOURCECONTROLLER_ENDPOINT")) > 0 {
		resourceControllerEndpoint = os.Getenv("RESOURCECONTROLLER_ENDPOINT")
	}
	resourceGroupId, err := getDefaultResourceGroupId(accessToken, resourceControllerEndpoint, os.Getenv("CE_ACCOUNT_ID"))
	if err != nil {
		fmt.Printf("ResourceController GetResourceGroups error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Resolved %s as default resource group id.\n", *resourceGroupId)

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
	listResult, _, err := ceClient.ListProjectsV2(&codeenginev2.ListProjectsV2Options{
		RefreshToken: &authenticator.RefreshToken,
	})
	if err != nil {
		fmt.Printf("ListProjectsV2 error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))

	// Create a new Code Engine project using the Code Engine Client
	projectName := "project-sdk-go-e2e--crud--" + time.Now().Format("060102-150405")
	region := "eu-de"
	createdProject, _, err := ceClient.CreateProjectV2(&codeenginev2.CreateProjectV2Options{
		RefreshToken:    &authenticator.RefreshToken,
		Name:            &projectName,
		ResourceGroupID: resourceGroupId,
		Region:          &region,
	})
	if err != nil {
		fmt.Printf("CreateProjectV2 error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created projects '%s' (guid: '%s').\n", *createdProject.Name, *createdProject.ID)

	//
	// Assume that the project creation takes some time
	for x := 0; x < 20; x++ {
		// sleep for 10 seconds and then try to fetch the project
		time.Sleep(10 * time.Second)

		obtainedProject, _, err := ceClient.GetProjectV2(&codeenginev2.GetProjectV2Options{
			RefreshToken: &authenticator.RefreshToken,
			ProjectGuid:  createdProject.ID,
		})
		if err != nil {
			fmt.Printf("GetProjectV2 error: %s\n", err.Error())
			os.Exit(1)
			return
		}
		fmt.Printf("Obtained status of project '%s' (guid: '%s'): %s.\n", *obtainedProject.Name, *obtainedProject.ID, *obtainedProject.Status)
		if *obtainedProject.Status == "active" {
			break
		}
	}

	resp, err := ceClient.DeleteProjectV2(&codeenginev2.DeleteProjectV2Options{
		RefreshToken: &authenticator.RefreshToken,
		ProjectGuid:  createdProject.ID,
	})
	if err != nil {
		fmt.Printf("DeleteProjectV2 error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted project: '%d'\n", resp.StatusCode)

	listResult, _, err = ceClient.ListProjectsV2(&codeenginev2.ListProjectsV2Options{
		RefreshToken: &authenticator.RefreshToken,
	})
	if err != nil {
		fmt.Printf("ListProjectsV2 error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Found %d projects.\n", len(listResult.Projects))
}
