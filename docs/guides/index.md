# SDM Provider

The SDM provider is used to manage your strongDM resources. This provider needs to be configured with API credentials from the strongDM account you wish to manage.

## Example Usage

    # main.tf

    provider "sdm" {
      api_access_key = file("~/.secrets/sdm_access_key.txt")
      api_secret_key = file("~/.secrets/sdm_secret_key.txt")
    }

## Authentication

Generate your authentication credentials from your strongDM account settings.

place holder for API key creation images

**These keys are secrets and should be treated as such.**