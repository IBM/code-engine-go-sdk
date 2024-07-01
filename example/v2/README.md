# Code Engine Go SDK Example

## Running example_v2.go

To run the example, create a Code Engine project from the Console or Code Engine CLI, and run the following commands from this directory:
1. `export CE_API_KEY=<Your IBM Cloud API key>`
2. `export CE_PROJECT_REGION=<The region (e.g. 'us-south') of your Code Engine project>`
3. `export CE_DOMAIN_MAPPING_NAME=<The name of your domain>`
4. `export CE_TLS_CERT_FILE_PATH=<The path to your TLS certificate file>`
5. `export CE_TLS_KEY_FILE_PATH=<The path to your TLS key file>`
6. `go run example_v2.go`

## How-to

### Set up an authenticator
```go
authenticator := &core.IamAuthenticator{
    ApiKey:       os.Getenv("CE_API_KEY"),
    ClientId:     "bx", 
    ClientSecret: "bx",
    URL:          "https://iam.cloud.ibm.com",
}
```

### Set up a Code Engine client
```go
codeEngineServiceOptions := &codeenginev2.CodeEngineV2Options{
    Authenticator: authenticator,
    URL:           "https://api." + os.Getenv("CE_PROJECT_REGION") + ".codeengine.cloud.ibm.com/v2",
}
codeEngineService, err := codeenginev2.NewCodeEngineV2UsingExternalConfig(codeEngineServiceOptions)
```

### Create a Code Engine project
```go
projectName := "my-project"
createdProject, _, err := codeEngineService.CreateProject(&codeenginev2.CreateProjectOptions{
    Name: &projectName,
})
```

### Create a Code Engine application
```go
createAppOpts := codeEngineService.NewCreateAppOptions(
    *createdProject.ID,
    "icr.io/codeengine/helloworld",
    "my-app",
)
createdApp, _, err := codeEngineService.CreateApp(createAppOpts)
```

### Create a Code Engine TLS secret
```go
createTLSSecretOpts := codeEngineService.NewCreateSecretOptions(
    *createdProject.ID,
    "tls",
    "my-tls-secret",
)

tlsCert, _ := os.ReadFile(os.Getenv("CE_TLS_CERT_FILE_PATH"))
tlsKey, _ := os.ReadFile(os.Getenv("CE_TLS_KEY_FILE_PATH"))

createTLSSecretOpts.Data = &codeenginev2.SecretDataTLSSecretData{
    TlsCert: core.StringPtr(string(tlsCert)),
    TlsKey:  core.StringPtr(string(tlsKey)),
}
createdTLSSecret, _, err := codeEngineService.CreateSecret(createTLSSecretOpts)
```

### Create a Code Engine domain mapping
```go
domainMappingName := os.Getenv("CE_DOMAIN_MAPPING_NAME")
appComponentRef := &codeenginev2.ComponentRef{
    Name:          createdApp.Name,
    ResourceType: core.StringPtr("app_v2"),
}

createDomainMappingOpts := codeEngineService.NewCreateDomainMappingOptions(
    *createdProject.ID,
    appComponentRef,
    domainMappingName,
    *createdTLSSecret.Name,
)
createdDomainMapping, _, err := codeEngineService.CreateDomainMapping(createDomainMappingOpts)
```

### Create a Code Engine function
```go
createFunctionOpts := codeEngineService.NewCreateFunctionOptions(
    *createdProject.ID,
    "data:text/plain;base64,YXN5bmMgZnVuY3Rpb24gbWFpbihwYXJhbXMpIHsKICByZXR1cm4gewogICAgICBzdGF0dXNDb2RlOiAyMDAsCiAgICAgIGhlYWRlcnM6IHsgJ0NvbnRlbnQtVHlwZSc6ICdhcHBsaWNhdGlvbi9qc29uJyB9LAogICAgICBib2R5OiBwYXJhbXMgfTsKfQptb2R1bGUuZXhwb3J0cy5tYWluID0gbWFpbjs=",
    "my-function",
    "nodejs-18",
)

createdFunction, _, err := codeEngineService.CreateFunction(createFunctionOpts)
```