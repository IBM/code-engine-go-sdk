package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/code-engine-go-sdk/v2/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func main() {

	var (
		codeEngineService     *codeenginev2.CodeEngineV2
		codeEngineApiEndpoint string
	)

	// Validate environment
	requiredEnvs := []string{"CE_API_KEY", "CE_DOMAIN_MAPPING_NAME", "CE_TLS_KEY_FILE_PATH", "CE_TLS_CERT_FILE_PATH"}
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

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey:       os.Getenv("CE_API_KEY"),
		ClientId:     "bx",
		ClientSecret: "bx",
		URL:          iamEndpoint,
	}

	if len(os.Getenv("CE_API_HOST")) > 0 {
		codeEngineApiEndpoint = "https://" + os.Getenv("CE_API_HOST") + "/v2"
	} else {
		codeEngineApiEndpoint = "https://api." + os.Getenv("CE_PROJECT_REGION") + ".codeengine.cloud.ibm.com/v2"
	}
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

	if len(os.Getenv("CE_ACCOUNT_ID")) > 0 {
		accountID := os.Getenv("CE_ACCOUNT_ID")
		fmt.Printf("Using account: '%s'\n", accountID)

		// Cleanup projects that have been created by prior runs
		cleanupProjectReclamations(authenticator, rcEndpoint, accountID)
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
		if err != nil && !strings.Contains(err.Error(), "Project is not yet active") {
			fmt.Printf("GetProject error: %s\n", err.Error())
			os.Exit(1)
			return
		}
		fmt.Printf("Obtained status of project '%s' (guid: '%s'): %s.\n", *obtainedProject.Name, *obtainedProject.ID, *obtainedProject.Status)
		if *obtainedProject.Status == "active" {
			break
		}
	}

	// Create ssh secret
	createSecretOpts := codeEngineService.NewCreateSecretOptions(
		*createdProject.ID,
		"ssh_auth",
		"ssh-secret",
	)

	sshKey := "-----BEGIN RSA PRIVATE KEY----------END RSA PRIVATE KEY-----"
	createSecretOpts.Data = &codeenginev2.SecretDataSSHSecretData{
		SshKey: &sshKey,
	}

	createdSecret, _, err := codeEngineService.CreateSecret(createSecretOpts)
	if err != nil {
		fmt.Printf("CreateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created ssh secret '%s'\n", *createdSecret.Name)

	// Get ssh secret
	getSecretOpts := codeEngineService.NewGetSecretOptions(
		*createdProject.ID,
		"ssh-secret",
	)
	obtainedSecret, _, err := codeEngineService.GetSecret(getSecretOpts)
	if err != nil {
		fmt.Printf("GetSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained secret '%s', format: %s", *obtainedSecret.Name, *obtainedSecret.Format)

	// Update ssh secret
	replaceSecretopts := codeEngineService.NewReplaceSecretOptions(
		*createdProject.ID,
		"ssh-secret",
		"*",
		"ssh_auth",
	)
	sshKeyUpdated := "-----BEGIN RSA PRIVATE KEY-----udpated-----END RSA PRIVATE KEY-----"
	replaceSecretopts.Data = &codeenginev2.SecretDataSSHSecretData{
		SshKey: &sshKeyUpdated,
	}
	format := "ssh_auth"
	replaceSecretopts.Format = &format
	updatedSecret, _, err := codeEngineService.ReplaceSecret(replaceSecretopts)
	if err != nil {
		fmt.Printf("UpdateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated secret '%s', format: %s", *updatedSecret.Name, *updatedSecret.Format)

	listSecretOpts := codeEngineService.NewListSecretsOptions(
		*createdProject.ID,
	)
	secretList, _, err := codeEngineService.ListSecrets(listSecretOpts)
	if err != nil {
		fmt.Printf("GetSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained secret list '%d'", len(secretList.Secrets))

	// Delete ssh secret
	deleteSecretOpts := codeEngineService.NewDeleteSecretOptions(
		*createdProject.ID,
		"ssh-secret",
	)
	resp, err := codeEngineService.DeleteSecret(deleteSecretOpts)
	if err != nil {
		fmt.Printf("DeleteSecret error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted secret: '%d'\n", resp.StatusCode)

	// Create basic auth secret
	createBasicAuthSecretOpts := codeEngineService.NewCreateSecretOptions(
		*createdProject.ID,
		"basic_auth",
		"basic-auth-secret",
	)

	username := "username"
	password := "password"
	createBasicAuthSecretOpts.Data = &codeenginev2.SecretDataBasicAuthSecretData{
		Username: &username,
		Password: &password,
	}

	createdBASecret, _, err := codeEngineService.CreateSecret(createBasicAuthSecretOpts)
	if err != nil {
		fmt.Printf("CreateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created basic auth secret '%s'\n", *createdBASecret.Name)

	// Get basic auth secret
	getBASecretOpts := codeEngineService.NewGetSecretOptions(
		*createdProject.ID,
		"basic-auth-secret",
	)
	obtainedBASecret, _, err := codeEngineService.GetSecret(getBASecretOpts)
	if err != nil {
		fmt.Printf("GetSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained secret '%s', format: %s", *obtainedBASecret.Name, *obtainedBASecret.Format)

	// Update basic auth secret
	replaceBASecretopts := codeEngineService.NewReplaceSecretOptions(
		*createdProject.ID,
		"basic-auth-secret",
		"*",
		"basic_auth",
	)
	username = "user2"
	password = "password2"
	replaceBASecretopts.Data = &codeenginev2.SecretDataBasicAuthSecretData{
		Username: &username,
		Password: &password,
	}
	format = "basic_auth"
	replaceBASecretopts.Format = &format
	updatedBASecret, _, err := codeEngineService.ReplaceSecret(replaceBASecretopts)
	if err != nil {
		fmt.Printf("UpdateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated secret '%s', format: %s", *updatedBASecret.Name, *updatedBASecret.Format)

	// Delete basic auth secret
	deleteBASecretOpts := codeEngineService.NewDeleteSecretOptions(
		*createdProject.ID,
		"basic-auth-secret",
	)
	resp, err = codeEngineService.DeleteSecret(deleteBASecretOpts)
	if err != nil {
		fmt.Printf("DeleteSecret error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted secret: '%d'\n", resp.StatusCode)

	// Create registry secret
	createRegistrySecretOpts := codeEngineService.NewCreateSecretOptions(
		*createdProject.ID,
		"registry",
		"registry-secret",
	)
	username = "user"
	password = "password"
	server := "github.com"
	email := "email@email.com"
	createRegistrySecretOpts.Data = &codeenginev2.SecretDataRegistrySecretData{
		Username: &username,
		Password: &password,
		Email:    &email,
		Server:   &server,
	}

	createdRegistrySecret, _, err := codeEngineService.CreateSecret(createRegistrySecretOpts)
	if err != nil {
		fmt.Printf("CreateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created tls secret '%s'\n", *createdRegistrySecret.Name)

	// Get registry secret
	getRegistrySecretOpts := codeEngineService.NewGetSecretOptions(
		*createdProject.ID,
		"registry-secret",
	)
	obtainedRegistrySecret, _, err := codeEngineService.GetSecret(getRegistrySecretOpts)
	if err != nil {
		fmt.Printf("GetSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained secret '%s', format: %s", *obtainedRegistrySecret.Name, *obtainedRegistrySecret.Format)

	// Update registry secret
	replaceRegistrySecretopts := codeEngineService.NewReplaceSecretOptions(
		*createdProject.ID,
		"registry-secret",
		"*",
		"registry",
	)
	username = "user2"
	password = "password2"
	replaceRegistrySecretopts.Data = &codeenginev2.SecretDataRegistrySecretData{
		Username: &username,
		Password: &password,
		Email:    &email,
		Server:   &server,
	}
	format = "registry"
	replaceRegistrySecretopts.Format = &format

	updatedRegistrySecret, _, err := codeEngineService.ReplaceSecret(replaceRegistrySecretopts)
	if err != nil {
		fmt.Printf("UpdateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated secret '%s', format: %s", *updatedRegistrySecret.Name, *updatedRegistrySecret.Format)

	// Delete registry secret
	deleteRegistrySecretOpts := codeEngineService.NewDeleteSecretOptions(
		*createdProject.ID,
		"registry-secret",
	)
	resp, err = codeEngineService.DeleteSecret(deleteRegistrySecretOpts)
	if err != nil {
		fmt.Printf("DeleteSecret error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted secret: '%d'\n", resp.StatusCode)

	// List Allowed Outbound Destinations
	listAllowedOutboundDestinationOptions := codeEngineService.NewListAllowedOutboundDestinationOptions(
		*createdProject.ID,
	)

	allowedOutboundDestinationList, _, err := codeEngineService.ListAllowedOutboundDestination(listAllowedOutboundDestinationOptions)
	if err != nil {
		fmt.Printf("ListAllowedOutboundDestination error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained AllowedOutboundDestination list '%d'", len(allowedOutboundDestinationList.AllowedOutboundDestinations))

	var cidrTypeDefault = "cidr_block"
	var cidrBlock = "192.68.4.0/24"
	var cidrBlockName = "my-allowed-outbound-destination"

	// Create allowed outbound destination
	createAllowedOutboundDestinationOpts := codeEngineService.NewCreateAllowedOutboundDestinationOptions(
		*createdProject.ID,
		&codeenginev2.AllowedOutboundDestinationPrototype{
			Type:      &cidrTypeDefault,
			CidrBlock: &cidrBlock,
			Name:      &cidrBlockName,
		},
	)

	createdAllowedOutboundDestination, _, err := codeEngineService.CreateAllowedOutboundDestination(createAllowedOutboundDestinationOpts)
	if err != nil {
		fmt.Printf("CreateAllowedOutboundDestination error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created allowed outbound destination '%s'\n", createdAllowedOutboundDestination)

	// Get allowed outbound destination
	getAllowedOutboundDestinationOpts := codeEngineService.NewGetAllowedOutboundDestinationOptions(
		*createdProject.ID,
		cidrBlockName,
	)

	obtainedAllowedOutboundDestination, _, err := codeEngineService.GetAllowedOutboundDestination(getAllowedOutboundDestinationOpts)
	if err != nil {
		fmt.Printf("GetAllowedOutboundDestination error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained allowed outbound destination '%s'\n", obtainedAllowedOutboundDestination)

	var updatedCidrBlock = "192.68.3.0/24"

	// Update allowed outbound destination
	allowedOutboundDestinationUpdateModel := &codeenginev2.AllowedOutboundDestinationPatch{
		Type:      &cidrTypeDefault,
		CidrBlock: &updatedCidrBlock,
	}
	updateAllowedOutboundDestinationAsPatch, err := allowedOutboundDestinationUpdateModel.AsPatch()
	if err != nil {
		fmt.Printf("allowedOutboundDestinationUpdateModel.AsPatch error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	updateAllowedOutboundDestinationOptions := codeEngineService.NewUpdateAllowedOutboundDestinationOptions(
		*createdProject.ID,
		cidrBlockName,
		"*",
		updateAllowedOutboundDestinationAsPatch,
	)
	updatedAllowedOutboundDestination, _, err := codeEngineService.UpdateAllowedOutboundDestination(updateAllowedOutboundDestinationOptions)
	if err != nil {
		fmt.Printf("UpdateAllowedOutboundDestination error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated allowed outbound destination'%s'\n", updatedAllowedOutboundDestination)

	// Delete allowed outbound destination
	deleteAllowedOutboundDestinationOpts := codeEngineService.NewDeleteAllowedOutboundDestinationOptions(
		*createdProject.ID,
		cidrBlockName,
	)

	resp, err = codeEngineService.DeleteAllowedOutboundDestination(deleteAllowedOutboundDestinationOpts)
	if err != nil {
		fmt.Printf("DeleteAllowedOutboundDestination error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted allowed outbound destination: '%d'\n", resp.StatusCode)

	// Create app
	createAppOpts := codeEngineService.NewCreateAppOptions(
		*createdProject.ID,
		"icr.io/codeengine/helloworld",
		"app-1",
	)

	createdApp, _, err := codeEngineService.CreateApp(createAppOpts)
	if err != nil {
		fmt.Printf("CreateApp error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created app '%s'\n", *createdApp.Name)

	// Create tls secret
	createTLSSecretOpts := codeEngineService.NewCreateSecretOptions(
		*createdProject.ID,
		"tls",
		"tls-secret",
	)

	tlsCert, err := os.ReadFile(os.Getenv("CE_TLS_CERT_FILE_PATH"))
	if err != nil {
		fmt.Printf("ReadFile error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	tlsKey, err := os.ReadFile(os.Getenv("CE_TLS_KEY_FILE_PATH"))
	if err != nil {
		fmt.Printf("ReadFile error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	createTLSSecretOpts.Data = &codeenginev2.SecretDataTLSSecretData{
		TlsCert: core.StringPtr(string(tlsCert)),
		TlsKey:  core.StringPtr(string(tlsKey)),
	}

	createdTLSSecret, _, err := codeEngineService.CreateSecret(createTLSSecretOpts)
	if err != nil {
		fmt.Printf("CreateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created tls secret '%s'\n", *createdTLSSecret.Name)

	// Create domain mapping
	domainMappingName := os.Getenv("CE_DOMAIN_MAPPING_NAME")
	appComponentRef := &codeenginev2.ComponentRef{
		Name:         createdApp.Name,
		ResourceType: core.StringPtr("app_v2"),
	}

	createDomainMappingOpts := codeEngineService.NewCreateDomainMappingOptions(
		*createdProject.ID,
		appComponentRef,
		domainMappingName,
		*createdTLSSecret.Name,
	)

	createdDomainMapping, _, err := codeEngineService.CreateDomainMapping(createDomainMappingOpts)
	if err != nil {
		fmt.Printf("CreateDomainMapping error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created domain mapping '%s'\n", *createdDomainMapping.Name)

	// Get domain mapping
	getDomainMappingOpts := codeEngineService.NewGetDomainMappingOptions(
		*createdProject.ID,
		domainMappingName,
	)

	obtainedDomainMapping, _, err := codeEngineService.GetDomainMapping(getDomainMappingOpts)
	if err != nil {
		fmt.Printf("GetDomainMapping error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained domain mapping '%s'\n", *obtainedDomainMapping.Name)

	// Create another app
	createAppOpts2 := codeEngineService.NewCreateAppOptions(
		*createdProject.ID,
		"icr.io/codeengine/helloworld",
		"app-2",
	)

	createdApp2, _, err := codeEngineService.CreateApp(createAppOpts2)
	if err != nil {
		fmt.Printf("CreateApp error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created app '%s'\n", *createdApp2.Name)

	app2ComponentRef := &codeenginev2.ComponentRef{
		Name:         createdApp2.Name,
		ResourceType: core.StringPtr("app_v2"),
	}

	// Update domain mapping
	updateDomainMappingComponentRef := &codeenginev2.DomainMappingPatch{
		Component: app2ComponentRef,
	}
	updateDomainMappingAsPatch, err := updateDomainMappingComponentRef.AsPatch()
	if err != nil {
		fmt.Printf("updateDomainMappingComponentRef.AsPatch error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	updateDomainMappingOptions := codeEngineService.NewUpdateDomainMappingOptions(
		*createdProject.ID,
		domainMappingName,
		"*",
		updateDomainMappingAsPatch,
	)

	updatedDomainMapping, _, err := codeEngineService.UpdateDomainMapping(updateDomainMappingOptions)
	if err != nil {
		fmt.Printf("UpdateDomainMapping error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated domain mapping '%s'\n", *updatedDomainMapping.Name)

	// List domain mappings
	listDomainMappingsOpts := codeEngineService.NewListDomainMappingsOptions(
		*createdProject.ID,
	)
	domainMappingList, _, err := codeEngineService.ListDomainMappings(listDomainMappingsOpts)
	if err != nil {
		fmt.Printf("ListDomainMappings error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained domain mapping list '%d'\n", len(domainMappingList.DomainMappings))

	// Delete domain mapping
	deleteDomainMappingOpts := codeEngineService.NewDeleteDomainMappingOptions(
		*createdProject.ID,
		domainMappingName,
	)

	resp, err = codeEngineService.DeleteDomainMapping(deleteDomainMappingOpts)
	if err != nil {
		fmt.Printf("DeleteDomainMapping error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted domain mapping: '%d'\n", resp.StatusCode)

	// Get tls secret
	getTLSSecretOpts := codeEngineService.NewGetSecretOptions(
		*createdProject.ID,
		"tls-secret",
	)
	obtainedTLSSecret, _, err := codeEngineService.GetSecret(getTLSSecretOpts)
	if err != nil {
		fmt.Printf("GetSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained secret '%s', format: %s", *obtainedTLSSecret.Name, *obtainedTLSSecret.Format)

	// Update ssh secret
	replaceTLSSecretopts := codeEngineService.NewReplaceSecretOptions(
		*createdProject.ID,
		"tls-secret",
		"*",
		"tls",
	)
	tlsKeyUpdate := "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCr+Qx5OrAHLWwm\nmstn7aEo317g/Lxv/Dmb/N/lanbGZfaVlnE1JrASNnEjps5CrVBLkjctbRYuAWOb\nvX4OKIGbSmT2JDu3gieg1v2gg0iuMmfqh9pgP8szlfB8lG7/rZ5m4ApEEB8iszIe\n+BrPsmlBBqd+tuJ3+t/BY9a7PjphkaMCbGlvoaDZEjT6KqubAMmZqkkYFT8mYx+A\nkwImgqVR5zMs4R2XSEl0QGLsFjnDtWLDvrHGdeGE0hnqTS5OusJ8bmNLJDOvSJSd\nZSWPtyahNQT4wAnp3RKxd3D2pdChqmxGdIs+eeNwzoXD42M2VEE/MgPLu7hPuPmC\nnN6AsET9AgMBAAECggEAc9d1cYv42zzbpz2KWt2VO6ULkl5syLqMS+kRIMaQb6Br\nc+Q9KeJ/pCUMHUnVktCQT/eUN4NN93t0D4qbiQn8FBEO5UcO+tQvwYZQnnkQ0lad\n7TvJ/B+8z2jm7+REyPG4y++KusJpVsSCtJ3H4bR6dhT3asHi15Mkem64TLTkOqf2\n5lWg5BUi3ZR5qFjriZdb7N3A+/Cb1fwOObCwNjRUJX6FAPpCdwEr+L9/o6bod+1N\nUArBYlSP8yMNyct3WzkPSpFnZxaYapjl0Nm9ipOfR5b9CHThoHg007WxdDF+6a/e\nSEJOZ0jRHwSctLhjSuL8/EOIuQGSHsyOK4SOmeHRgQKBgQDYlrafbArou+pStqIU\nZCmV51UqSfqZAAJ+YzV9rqhsM97yQKQYEESeIbgAnWCGlAbY7XrysIA/aOdglOuF\no60oRqlnkYZJT8SXjvnwmyxor67f3G0jbVuoefYL1G1EPdcL9l2K0xehOa2huYm0\n8lvlI8PPKKJkmu22r/TNyp6VEQKBgQDLRAHsDjNdwyMKVGe2G6ZmnyDWhGzVOOZf\n+Ixfmt0BK5AnmJBeABM6WRC/6EM0eX31lcev7sJMpWF4Iw0Op+tW2gmtfphi3j/l\nG7B3lU4V/M6jw0CrASy1RGY257ou3o+/yS4N6/lafZw/V8KDjgJngCeyRhgFf+Rj\nVNC3FIsBLQKBgERN43ILZLVY7eD/78V2gRbhSZ54jitKMX8iUnA8cKkPArRrZlSg\nbMNh5uFqwFIwxKgM3MVEnG1i6/Utgck3gRg+kJY08qCUI2+Yi4IxraOmJAQ9Q730\ncv+C1vGMIJlw1yzSmVV6lO0nf3aNSLxj4k81JD9klTIdGfKPMyjjSXfBAoGBALhl\nWI0JkOWlSZtsWK1mxfzgrMyOU6DWvn8fnlB4z7bpCxwwlf8AeHD9LWm6zYTEFlV8\n7CsZIOChQxvWSFkcUi13HUJrztgaIMK57Mt/AdiGf/sl/Ptk1GcYxtVWQJuWQbfN\nTN9KS+oge2cnOQlZAatdIiXi2pXaoJjP74u2sid9AoGAFuustiKF2vffjhyEg+HL\nU57p6LG7y6x02COLDhKTX4c/bEa6MX4f91ZKXy2S47tCgLSf4SYd49k1H0wQEDkl\nYs+pznN30O/Jxu063JfvFbLZxJkeayLpQL12w+NQUDwsF6MGvIYTnUefhkfb3LWC\njBKCTCcw9u4SVX1jK4f2/OU=\n-----END PRIVATE KEY-----"
	tlsCertUpdate := "-----BEGIN CERTIFICATE-----\nMIICqDCCAZACCQDB2CY2jE7CCjANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtm\nb28uYmFyLmNvbTAeFw0yMzA2MjkyMDM5MzhaFw0yNDA2MjgyMDM5MzhaMBYxFDAS\nBgNVBAMMC2Zvby5iYXIuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\nAQEAq/kMeTqwBy1sJprLZ+2hKN9e4Py8b/w5m/zf5Wp2xmX2lZZxNSawEjZxI6bO\nQq1QS5I3LW0WLgFjm71+DiiBm0pk9iQ7t4InoNb9oINIrjJn6ofaYD/LM5XwfJRu\n/62eZuAKRBAfIrMyHvgaz7JpQQanfrbid/rfwWPWuz46YZGjAmxpb6Gg2RI0+iqr\nmwDJmapJGBU/JmMfgJMCJoKlUeczLOEdl0hJdEBi7BY5w7Viw76xxnXhhNIZ6k0u\nTrrCfG5jSyQzr0iUnWUlj7cmoTUE+MAJ6d0SsXdw9qXQoapsRnSLPnnjcM6Fw+Nj\nNlRBPzIDy7u4T7j5gpzegLBE/QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCXRwhk\nwjvOzKh5R+QKHGjtcjutSkwZbMj5+5enN/8IwX2BbX0i/aALxEPcZExMK5aIS5rm\n+kUkDyZkYVaMQQoTGNHSnnET8WJf8zGqd/GdiVxZRVXjOnQ5tEezdwFm0a3TEEKw\n/2HG9chz24ywhbIZZMEFmse7LLrcy5XSUQzOTMWBKZ8fTEXBYaEVhD/9b4SPuLpw\ni4vDZPt+e+p96NcGNf0b932aod+X34dARUd55UM9PY4i4Z7UzzV7zK+U6tHjzzmg\nrv+JA2kDt3mwQXn7bfgRxLcpBZFpUHjLRe+MGlQJM2xFYAXop9ZzF1go58ErHbsT\nCyXJ56cw0ffDrXSn\n-----END CERTIFICATE-----"
	replaceTLSSecretopts.Data = &codeenginev2.SecretDataTLSSecretData{
		TlsCert: &tlsCertUpdate,
		TlsKey:  &tlsKeyUpdate,
	}
	format = "tls"
	replaceTLSSecretopts.Format = &format
	updatedTLSSecret, _, err := codeEngineService.ReplaceSecret(replaceTLSSecretopts)
	if err != nil {
		fmt.Printf("UpdateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated secret '%s', format: %s", *updatedTLSSecret.Name, *updatedTLSSecret.Format)

	// Delete tls secret
	deleteTLSSecretOpts := codeEngineService.NewDeleteSecretOptions(
		*createdProject.ID,
		"tls-secret",
	)
	resp, err = codeEngineService.DeleteSecret(deleteTLSSecretOpts)
	if err != nil {
		fmt.Printf("DeleteSecret error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted secret: '%d'\n", resp.StatusCode)

	// List Function Runtimes
	listFunctionRuntimesOptions := codeEngineService.NewListFunctionRuntimesOptions()

	functionRuntimeList, _, err := codeEngineService.ListFunctionRuntimes(listFunctionRuntimesOptions)
	if err != nil {
		fmt.Printf("ListFunctionRuntimes error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained Function runtime list '%d'", len(functionRuntimeList.FunctionRuntimes))

	// List Functions
	listFunctionsOptions := codeEngineService.NewListFunctionsOptions(
		*createdProject.ID,
	)

	functionsList, _, err := codeEngineService.ListFunctions(listFunctionsOptions)
	if err != nil {
		fmt.Printf("ListFunctions error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained Functions list '%d'", len(functionsList.Functions))

	// Create Function
	createFunctionOptions := codeEngineService.NewCreateFunctionOptions(
		*createdProject.ID,
		"data:text/plain;base64,YXN5bmMgZnVuY3Rpb24gbWFpbihwYXJhbXMpIHsKICByZXR1cm4gewogICAgICBzdGF0dXNDb2RlOiAyMDAsCiAgICAgIGhlYWRlcnM6IHsgJ0NvbnRlbnQtVHlwZSc6ICdhcHBsaWNhdGlvbi9qc29uJyB9LAogICAgICBib2R5OiBwYXJhbXMgfTsKfQptb2R1bGUuZXhwb3J0cy5tYWluID0gbWFpbjs=",
		"my-function",
		"nodejs-20",
	)

	createdFunction, _, err := codeEngineService.CreateFunction(createFunctionOptions)
	if err != nil {
		fmt.Printf("CreateFunction error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created Function '%s'\n", *createdFunction.Name)

	//Get Function
	getFunctionOptions := codeEngineService.NewGetFunctionOptions(
		*createdProject.ID,
		"my-function",
	)

	obtainedFunction, _, err := codeEngineService.GetFunction(getFunctionOptions)
	if err != nil {
		fmt.Printf("GetFunction error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Obtained Function '%s'\n", *obtainedFunction.Name)

	// Update Function
	functionUpdateModel := &codeenginev2.FunctionPatch{
		ScaleMaxExecutionTime: core.Int64Ptr(30),
	}
	updateFunctionAsPatch, err := functionUpdateModel.AsPatch()
	if err != nil {
		fmt.Printf("functionUpdateModel.AsPatch error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	updateFunctionOptions := codeEngineService.NewUpdateFunctionOptions(
		*createdProject.ID,
		"my-function",
		"*",
		updateFunctionAsPatch,
	)
	updatedFunction, _, err := codeEngineService.UpdateFunction(updateFunctionOptions)
	if err != nil {
		fmt.Printf("UpdateFunction error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Updated Function '%s'\n", *updatedFunction.Name)

	// Delete Function
	deleteFunctionOptions := codeEngineService.NewDeleteFunctionOptions(
		*createdProject.ID,
		"my-function",
	)

	resp, err = codeEngineService.DeleteFunction(deleteFunctionOptions)
	if err != nil {
		fmt.Printf("DeleteFunction error: %s (transaction-id: '%s')\n", err.Error(), resp.Headers.Get("X-Transaction-Id"))
		os.Exit(1)
		return
	}
	fmt.Printf("Deleted Function: '%d'\n", resp.StatusCode)

	// Delete Project
	deleteProjectOptions := codeEngineService.NewDeleteProjectOptions(
		*createdProject.ID,
	)

	resp, err = codeEngineService.DeleteProject(deleteProjectOptions)
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
