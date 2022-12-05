package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/IBM/code-engine-go-sdk/ibmcloudcodeenginev1"
	"github.com/IBM/go-sdk-core/v5/core"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// Validate environment
	requiredEnvs := []string{"CE_API_KEY", "CE_API_HOST", "CE_PROJECT_ID"}
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
	ceClient, err := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
		Authenticator: authenticator,
		URL:           "https://" + os.Getenv("CE_API_HOST") + "/api/v1",
	})
	if err != nil {
		fmt.Printf("NewIbmCloudCodeEngineV1 error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// Use the http library to get an IAM Delegated Refresh Token
	iamRequestData := url.Values{}
	iamRequestData.Set("grant_type", "urn:ibm:params:oauth:grant-type:apikey")
	iamRequestData.Set("apikey", os.Getenv("CE_API_KEY"))
	iamRequestData.Set("response_type", "delegated_refresh_token")
	iamRequestData.Set("receiver_client_ids", "ce")
	iamRequestData.Set("delegated_refresh_token_expiry", "3600")

	client := &http.Client{}
	req, err := http.NewRequest("POST", iamEndpoint+"/identity/token", strings.NewReader(iamRequestData.Encode()))
	if err != nil {
		fmt.Printf("NewRequest err: %s\n", err)
		os.Exit(1)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("POST /identity/token err: %s\n", err)
		os.Exit(1)
		return
	}

	var iamResponseData map[string]string
	err := json.NewDecoder(resp.Body).Decode(&iamResponseData)
	if err != nil {
		fmt.Printf("Failed to decode IAM response data: %s\n", err.Error())
		os.Exit(1)
		return
	}
	err := resp.Body.Close()
	if err != nil {
		fmt.Printf("Failed to close the response body: %s\n", err.Error())
		os.Exit(1)
		return
	}
	delegatedRefreshToken := iamResponseData["delegated_refresh_token"]

	// Get Code Engine project config using the Code Engine Client
	projectID := os.Getenv("CE_PROJECT_ID")
	fmt.Printf("Obtaining a kube config of project '%s'\n", projectID)
	result, _, err := ceClient.GetKubeconfig(&ibmcloudcodeenginev1.GetKubeconfigOptions{
		XDelegatedRefreshToken: &delegatedRefreshToken,
		ID:                     &projectID,
	})
	if err != nil {
		fmt.Printf("GetKubeconfig error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// Get Kubernetes client using Code Engine project config
	kubeConfig, err := clientcmd.NewClientConfigFromBytes([]byte(*result))
	if err != nil {
		fmt.Printf("NewClientConfigFromBytes error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	kubeClientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		fmt.Printf("ClientConfig error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	kubeClient, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		fmt.Printf("NewForConfig error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// Get something from project
	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		fmt.Printf("Namespace error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	configMapList, err := kubeClient.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Pods list error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Project %s has %d configmaps.\n", os.Getenv("CE_PROJECT_ID"), len(configMapList.Items))
}
