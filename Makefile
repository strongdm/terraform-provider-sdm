all:
	go build -o example/terraform.d/plugins/darwin_amd64/terraform-provider-sdm
	cd example/ && terraform init && TF_LOG=DEBUG terraform apply -auto-approve