package main

import (
	"fmt"
	"os"

	"github.com/IBM/code-engine-go-sdk/ibmcloudcodeenginev1"
	"github.com/IBM/go-sdk-core/v4/core"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// Validate environment
	requiredEnvs := []string{"CE_API_KEY", "CE_PROJECT_REGION", "CE_PROJECT_ID"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			fmt.Printf("Environment variable %s must be set\n", env)
			os.Exit(1)
			return
		}
	}

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey:       os.Getenv("CE_API_KEY"),
		ClientId:     "bx",
		ClientSecret: "bx",
	}

	// Setup a Code Engine client
	ceClient, err := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
		Authenticator: authenticator,
		URL:           "https://api." + os.Getenv("CE_PROJECT_REGION") + ".codeengine.cloud.ibm.com/api/v1",
	})
	if err != nil {
		fmt.Printf("NewIbmCloudCodeEngineV1 error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// Get an IAM token
	iamToken, err := authenticator.RequestToken()
	if err != nil {
		fmt.Printf("RequestToken error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// Get Code Engine project config using the Code Engine Client
	projectID := os.Getenv("CE_PROJECT_ID")
	refreshToken := iamToken.RefreshToken
	result, _, err := ceClient.ListKubeconfig(&ibmcloudcodeenginev1.ListKubeconfigOptions{
		RefreshToken: &refreshToken,
		ID:           &projectID,
	})
	if err != nil {
		fmt.Printf("ListKubeconfig error: %s\n", err.Error())
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
	configMapList, err := kubeClient.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Pods list error: %s\n", err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("Project %s has %d configmaps.\n", os.Getenv("CE_PROJECT_ID"), len(configMapList.Items))

}
