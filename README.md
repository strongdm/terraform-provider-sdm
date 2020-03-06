Terraform Provider for strongDM
==================

Requirements
------------
- [Terraform](https://www.terraform.io/downloads.html) 0.12+


Using the provider
----------------------
TODO



Testing the provider
----------------------
In order to run the tests these environment variables must be set:
- `SDM_API_ACCESS_KEY`
- `SDM_API_SECRET_KEY`

_Note_: The acceptance tests create and destroy real resources.

Then from the cloned repo, run: 
```
TF_ACC=yes go test ./sdm -v -count=1 -mod=vendor
```

The `TF_ACC=yes` sets up terraform's testing utilties to allow the tests to run. Without this being set, all tests will pass without any action.