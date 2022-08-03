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

### Use an HTTP library to get a Delegated Refresh Token from IAM
```go
iamRequestData := url.Values{}
iamRequestData.Set("grant_type", "urn:ibm:params:oauth:grant-type:apikey")
iamRequestData.Set("apikey", os.Getenv("CE_API_KEY"))
iamRequestData.Set("response_type", "delegated_refresh_token")
iamRequestData.Set("receiver_client_ids", "ce")
iamRequestData.Set("delegated_refresh_token_expiry", "3600")

client := &http.Client{}
req, _ := http.NewRequest("POST", "https://iam.cloud.ibm.com/identity/token", strings.NewReader(iamRequestData.Encode()))
req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
resp, _ := client.Do(req)

var iamResponseData map[string]string
json.NewDecoder(resp.Body).Decode(&iamResponseData)
delegatedRefreshToken := iamResponseData["delegated_refresh_token"]
```

### Use the Code Engine client to get a Kubernetes config
```go
projectID := os.Getenv("CE_PROJECT_ID")
result, _, err := ceClient.GetKubeconfig(&ibmcloudcodeenginev1.GetKubeconfigOptions{
    XDelegatedRefreshToken: &delegatedRefreshToken,
    ID:                     &projectID,
})
```

## Deprecated endpoint

The `/namespaces/{id}/config` endpoint function, `ListKubeconfig()`, is deprecated, and will be removed before Code Engine is out of Beta. Please use the `GetKubeconfig` function, demonstrated in the example above.