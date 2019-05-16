terraform-provider-jsonmultiplex: $(shell find cmd -type f)
	go build -o terraform-provider-jsonmultiplex cmd/terraform-provider-jsonmultiplex/*.go
