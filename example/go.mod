module github.com/IBM/code-engine-go-sdk/example

go 1.13

require (
	github.com/IBM/code-engine-go-sdk v0.0.0-00010101000000-000000000000
	github.com/IBM/go-sdk-core/v5 v5.9.5
	k8s.io/apimachinery v0.21.7
	k8s.io/client-go v0.21.7
)

require (
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	golang.org/x/net v0.0.0-20211020060615-d418f374d309 // indirect
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

replace github.com/IBM/code-engine-go-sdk => ./..
