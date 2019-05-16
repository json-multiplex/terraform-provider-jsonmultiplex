module github.com/json-multiplex/terraform-provider-jsonmultiplex

go 1.12

require github.com/hashicorp/terraform v0.11.13

require (
	github.com/json-multiplex/iam v0.0.0
	google.golang.org/grpc v1.20.1
)

replace github.com/json-multiplex/iam v0.0.0 => ../iam
