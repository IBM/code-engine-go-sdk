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

	// Create tls secret
	createTLSSecretOpts := codeEngineService.NewCreateSecretOptions(
		*createdProject.ID,
		"tls",
		"tls-secret",
	)

	tlsKey := "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDGfJO9qkAlq8Vy\nKNyJEAwJ+VGurknonWKL+/B/8uS45qDYHP9McyokfHR6GEeL3p/vk4zf+QI/+5Dn\n0IK6jyiLzl4x0FlEVbEesLubi/6B8r6I+pYfWlOX+ShJkryuZcMjuEtvP6sli+Wj\nr5yILu8YHgAVVdvLs7XJmlDPv/kmq9R66Nsl02PgLazJfztijcdBGkQAPxClwwkJ\nzVCWE/G7fS0iYUb76ScHrxLwN7Bh+wTOIMHk6qqK2UA45a8MmyGOkD4SoB4K3z3y\nGNTQrxQbj+wCyK9kY2/sTs++kcsiwTfTx+17UYO05S0+ExqIWrD6bJpnYmWART/2\niBvcAfLvAgMBAAECggEBAKzVj6SJGmBzKXQVxquHEKSiuBC+bVcjrMsuL6aKb8Xd\n9VMaNOhyI9EvmhEzESHnUidAuVvSLbZfLTfeZedjfy/2HCmOPhz17UxHIqX4ij7H\njEgkxBI7Ci18ZStjne7SZ9CzyuPtce842VbmNQyUqde7T+FEKSdArlwFhrbQeHjF\ngJQrsroY8d0h9Xt6UlfVzX/CeNWP98YJLJ7my9WYRhZlcBE5qwyaMRIY2UKcjgpx\nXaViny79P5GwiaVGgOUYZ32bA+GHf7u5WP7lCiqT32SgTZzQ9dov9KN+QSkui7qO\nj0tC0c7OI59zatPAbp+t1LDjsgTjkuoReHes4nhupvECgYEA5b+S9uLtTn3XxLMf\nR00anvek8EUbTA/TRrJWUhCgvyVCafpyx0BqJC9eR5LnmD3f3yFXDvD5DF201zTn\n1Py+sk6oUfuPLXz8P76L5Wpz8ryRjR4LfLu0CTGMuUMDfE3NRHQJHNKnkrPwu0mX\njwbZrI08Xs8yjyx4gapdwE1cEvkCgYEA3SqPHW1AjCdnhSpnzf9QwxX0hzUKvUBK\neuuhKvmwh/AnE2y4b/6VH7TRj+fUvbaSFl63tTKXvUA2J7gHvz8o3j24EYAibwe/\nTvcloLjNxHOEq42vwB9zoZZ1UjvNhRo7lB6626/ffQRHXeSfoyMr2GTFYdpCAZds\nf8/fHA14RycCgYASzd9FfcVWi05Btzd0KodnQ3WohL97NkBgpPATv3CotG//JJSI\nYmlNlOLukMOL3mSYaq4pduerb3ABvT7MW/NvvKhiLWjGnFg5D2t7136t+2keV7sw\n9lwB9KBD+YwrfGK0m5qzVTqJ81hcu+U/u5vNV7H9QJAuz8D9O+h4eNx0YQKBgQC+\naa3dv/oasLJHzEKi8HYv/+8PmXMtjPSS79tKjL6XywNZjfkdMypgqeTi6M4Yp98O\ns22m63AI2AfIGoFQ/qfI74pSRudegGUNL2uN/I3r3SkUKmBuIKYFMOzBaAuB1RwG\nYo6uJbVchRqMlBF8+wL8w4XMwYSiqiQXxnhoRpCPcQKBgE2UoeHxvydgQjcfx7/M\n8BmmLqohWUF6tU62TVBMhYeO77H5Qkn/y0K6UPvar7x0lNAz2ljiUtYvvHc3S9Mc\nwSQ7GGIZu4ro/tLfi1xVfeQH5Ibm/pdk+1BZfcGeYAHC9Gr+LAT0iJRu8nFyWroB\nq/Tq26sIqxRotUdtRDJ6D6jf\n-----END PRIVATE KEY-----"
	tlsCert := "-----BEGIN CERTIFICATE-----\nMIICqDCCAZACCQDE3bI657EkGTANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtm\nb28uYmFyLmNvbTAeFw0yMzA3MDUxOTA5MzRaFw0yMzA4MDQxOTA5MzRaMBYxFDAS\nBgNVBAMMC2Zvby5iYXIuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\nAQEAxnyTvapAJavFcijciRAMCflRrq5J6J1ii/vwf/LkuOag2Bz/THMqJHx0ehhH\ni96f75OM3/kCP/uQ59CCuo8oi85eMdBZRFWxHrC7m4v+gfK+iPqWH1pTl/koSZK8\nrmXDI7hLbz+rJYvlo6+ciC7vGB4AFVXby7O1yZpQz7/5JqvUeujbJdNj4C2syX87\nYo3HQRpEAD8QpcMJCc1QlhPxu30tImFG++knB68S8DewYfsEziDB5OqqitlAOOWv\nDJshjpA+EqAeCt898hjU0K8UG4/sAsivZGNv7E7PvpHLIsE308fte1GDtOUtPhMa\niFqw+myaZ2JlgEU/9ogb3AHy7wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBgbjph\n/srZdtk5Fg1KVQCDg6h73ss0sfdvryz9mGQbhCsFxzcZgg5eF+5UKQPCKwunvREI\nSCemXjojfKa6UpfWe51MQs1ehigxiUrdEVt0DOgW7C78JXPZ+vdTMm5XTyaoim9H\nqB2Fhf4An+kxHIQdRwlf20YWnM8/nihKctbPF4hwWGvw88Ob5ugvDQ805JyCkiX1\nWBkSuIcNmv8L0D/BY2dvwhOB2fhg9cmm5m5OrY+MrEAA1Bk982sJV1x0rFzJK7oA\n3TuE9L09bd9lZD2BCF1ZXq8aAhExtCEMgrI3btnRcpKvcuKl+q1UHRBanDHIWDiY\nN/AOez7b5poi09yV\n-----END CERTIFICATE-----"
	createTLSSecretOpts.Data = &codeenginev2.SecretDataTLSSecretData{
		TlsCert: &tlsCert,
		TlsKey:  &tlsKey,
	}

	createdTLSSecret, _, err := codeEngineService.CreateSecret(createTLSSecretOpts)
	if err != nil {
		fmt.Printf("CreateSecret error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Created tls secret '%s'\n", *createdTLSSecret.Name)

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
	tlsKey = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCr+Qx5OrAHLWwm\nmstn7aEo317g/Lxv/Dmb/N/lanbGZfaVlnE1JrASNnEjps5CrVBLkjctbRYuAWOb\nvX4OKIGbSmT2JDu3gieg1v2gg0iuMmfqh9pgP8szlfB8lG7/rZ5m4ApEEB8iszIe\n+BrPsmlBBqd+tuJ3+t/BY9a7PjphkaMCbGlvoaDZEjT6KqubAMmZqkkYFT8mYx+A\nkwImgqVR5zMs4R2XSEl0QGLsFjnDtWLDvrHGdeGE0hnqTS5OusJ8bmNLJDOvSJSd\nZSWPtyahNQT4wAnp3RKxd3D2pdChqmxGdIs+eeNwzoXD42M2VEE/MgPLu7hPuPmC\nnN6AsET9AgMBAAECggEAc9d1cYv42zzbpz2KWt2VO6ULkl5syLqMS+kRIMaQb6Br\nc+Q9KeJ/pCUMHUnVktCQT/eUN4NN93t0D4qbiQn8FBEO5UcO+tQvwYZQnnkQ0lad\n7TvJ/B+8z2jm7+REyPG4y++KusJpVsSCtJ3H4bR6dhT3asHi15Mkem64TLTkOqf2\n5lWg5BUi3ZR5qFjriZdb7N3A+/Cb1fwOObCwNjRUJX6FAPpCdwEr+L9/o6bod+1N\nUArBYlSP8yMNyct3WzkPSpFnZxaYapjl0Nm9ipOfR5b9CHThoHg007WxdDF+6a/e\nSEJOZ0jRHwSctLhjSuL8/EOIuQGSHsyOK4SOmeHRgQKBgQDYlrafbArou+pStqIU\nZCmV51UqSfqZAAJ+YzV9rqhsM97yQKQYEESeIbgAnWCGlAbY7XrysIA/aOdglOuF\no60oRqlnkYZJT8SXjvnwmyxor67f3G0jbVuoefYL1G1EPdcL9l2K0xehOa2huYm0\n8lvlI8PPKKJkmu22r/TNyp6VEQKBgQDLRAHsDjNdwyMKVGe2G6ZmnyDWhGzVOOZf\n+Ixfmt0BK5AnmJBeABM6WRC/6EM0eX31lcev7sJMpWF4Iw0Op+tW2gmtfphi3j/l\nG7B3lU4V/M6jw0CrASy1RGY257ou3o+/yS4N6/lafZw/V8KDjgJngCeyRhgFf+Rj\nVNC3FIsBLQKBgERN43ILZLVY7eD/78V2gRbhSZ54jitKMX8iUnA8cKkPArRrZlSg\nbMNh5uFqwFIwxKgM3MVEnG1i6/Utgck3gRg+kJY08qCUI2+Yi4IxraOmJAQ9Q730\ncv+C1vGMIJlw1yzSmVV6lO0nf3aNSLxj4k81JD9klTIdGfKPMyjjSXfBAoGBALhl\nWI0JkOWlSZtsWK1mxfzgrMyOU6DWvn8fnlB4z7bpCxwwlf8AeHD9LWm6zYTEFlV8\n7CsZIOChQxvWSFkcUi13HUJrztgaIMK57Mt/AdiGf/sl/Ptk1GcYxtVWQJuWQbfN\nTN9KS+oge2cnOQlZAatdIiXi2pXaoJjP74u2sid9AoGAFuustiKF2vffjhyEg+HL\nU57p6LG7y6x02COLDhKTX4c/bEa6MX4f91ZKXy2S47tCgLSf4SYd49k1H0wQEDkl\nYs+pznN30O/Jxu063JfvFbLZxJkeayLpQL12w+NQUDwsF6MGvIYTnUefhkfb3LWC\njBKCTCcw9u4SVX1jK4f2/OU=\n-----END PRIVATE KEY-----"
	tlsCert = "-----BEGIN CERTIFICATE-----\nMIICqDCCAZACCQDB2CY2jE7CCjANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtm\nb28uYmFyLmNvbTAeFw0yMzA2MjkyMDM5MzhaFw0yNDA2MjgyMDM5MzhaMBYxFDAS\nBgNVBAMMC2Zvby5iYXIuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\nAQEAq/kMeTqwBy1sJprLZ+2hKN9e4Py8b/w5m/zf5Wp2xmX2lZZxNSawEjZxI6bO\nQq1QS5I3LW0WLgFjm71+DiiBm0pk9iQ7t4InoNb9oINIrjJn6ofaYD/LM5XwfJRu\n/62eZuAKRBAfIrMyHvgaz7JpQQanfrbid/rfwWPWuz46YZGjAmxpb6Gg2RI0+iqr\nmwDJmapJGBU/JmMfgJMCJoKlUeczLOEdl0hJdEBi7BY5w7Viw76xxnXhhNIZ6k0u\nTrrCfG5jSyQzr0iUnWUlj7cmoTUE+MAJ6d0SsXdw9qXQoapsRnSLPnnjcM6Fw+Nj\nNlRBPzIDy7u4T7j5gpzegLBE/QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCXRwhk\nwjvOzKh5R+QKHGjtcjutSkwZbMj5+5enN/8IwX2BbX0i/aALxEPcZExMK5aIS5rm\n+kUkDyZkYVaMQQoTGNHSnnET8WJf8zGqd/GdiVxZRVXjOnQ5tEezdwFm0a3TEEKw\n/2HG9chz24ywhbIZZMEFmse7LLrcy5XSUQzOTMWBKZ8fTEXBYaEVhD/9b4SPuLpw\ni4vDZPt+e+p96NcGNf0b932aod+X34dARUd55UM9PY4i4Z7UzzV7zK+U6tHjzzmg\nrv+JA2kDt3mwQXn7bfgRxLcpBZFpUHjLRe+MGlQJM2xFYAXop9ZzF1go58ErHbsT\nCyXJ56cw0ffDrXSn\n-----END CERTIFICATE-----"
	replaceTLSSecretopts.Data = &codeenginev2.SecretDataTLSSecretData{
		TlsCert: &tlsCert,
		TlsKey:  &tlsKey,
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
	fmt.Printf("Created tls secret '%s'\n", *createdBASecret.Name)

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
