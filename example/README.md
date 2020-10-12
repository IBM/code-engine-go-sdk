# Code Engine Go SDK Example

## Running example.go

To run the example, create a Code Engine project from the Console or Code Engine CLI, and run the following commands from this directory:
1. `export CE_API_KEY=<Your IBM Cloud API key>`
2. `export CE_PROJECT_ID=<Your Code Engine project ID>`
3. `export CE_PROJECT_REGION=<The region (e.g. 'us-south') of your Code Engine project>`
4. `go run example.go`

## How-to

### Set up an authenticator
```go
authenticator := &core.IamAuthenticator{
    ApiKey:       os.Getenv("CE_API_KEY"),
    ClientId:     "bx",
    ClientSecret: "bx",
}
```

### Set up a Code Engine client
```go
ceClient, err := ibmcloudcodeenginev1.NewIbmCloudCodeEngineV1(&ibmcloudcodeenginev1.IbmCloudCodeEngineV1Options{
    Authenticator: authenticator,
    URL:           "https://api." + os.Getenv("CE_PROJECT_REGION") + ".codeengine.cloud.ibm.com/api/v1",
})
```

### Use the Code Engine client to get a Kubernetes config
```go
projectID := os.Getenv("CE_PROJECT_ID")
iamToken, _ := authenticator.RequestToken()
refreshToken := iamToken.RefreshToken
result, _, err := ceClient.ListKubeconfig(&ibmcloudcodeenginev1.ListKubeconfigOptions{
    RefreshToken: &refreshToken,
    ID:           &projectID,
})
```
