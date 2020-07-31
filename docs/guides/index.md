# SDM Provider

The SDM provider is used to manage your strongDM resources. 

## Configuration

This provider needs to be configured with API credentials from your strongDM account. These credentials can be supplied to the provider directly or set using environment variables.

### Example Usage

main.tf
```hcl
provider "sdm" {
  api_access_key = file("~/.secrets/sdm_access_key.txt")
  api_secret_key = file("~/.secrets/sdm_secret_key.txt")
}
```

#### Using environment variables

```shell
$ export SDM_API_ACCESS_KEY=<YOUR KEY>
$ export SDM_API_SECRET_KEY=<YOUR SECRET KEY>
```

main.tf
```hcl
provider "sdm" {}
```

## Guides

* [Adding Resources to strongDM](./adding_resources)
* [How to Create a Gateway](./how_to_create_a_gateway)
* [Managing Roles](./roles)
* [Managing Users](./user_management)