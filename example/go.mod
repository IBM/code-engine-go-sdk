module github.com/IBM/code-engine-go-sdk/example

go 1.13

require (
	github.com/IBM/code-engine-go-sdk v0.0.0-00010101000000-000000000000
	github.com/IBM/go-sdk-core/v4 v4.6.1
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/go-openapi/errors v0.19.7 // indirect
	github.com/go-openapi/strfmt v0.19.6 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.4.2 // indirect
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sys v0.0.0-20201009025420-dfb3f7c4e634 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/IBM/code-engine-go-sdk => ./../../code-engine-go-sdk
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
