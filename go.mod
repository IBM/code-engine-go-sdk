module github.com/IBM/code-engine-go-sdk

go 1.18

require (
	github.com/IBM/go-sdk-core/v5 v5.15.0
	github.com/go-openapi/strfmt v0.22.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.29.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-openapi/errors v0.21.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.13.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	// Replace golang.org/x/net as the latest level, thereby logically excluding vulnerable versions prior to v0.17.0
	golang.org/x/net => golang.org/x/net v0.20.0
	// Replace golang.org/x/text as the latest level, thereby logically excluding vulnerable versions prior to v0.3.7
	golang.org/x/text => golang.org/x/text v0.14.0
)
