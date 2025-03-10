module github.com/IBM/v2

go 1.23.1

toolchain go1.23.5

require (
	github.com/IBM/code-engine-go-sdk/v2 v2.0.3
	github.com/IBM/go-sdk-core/v5 v5.18.5
	github.com/IBM/platform-services-go-sdk v0.77.1
)

require (
	github.com/IBM/code-engine-go-sdk v0.0.0-20221209153711-82472bae75eb // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/strfmt v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.24.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	go.mongodb.org/mongo-driver v1.17.2 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/IBM/code-engine-go-sdk/v2 => ./../..

replace github.com/IBM/code-engine-go-sdk => ../..
