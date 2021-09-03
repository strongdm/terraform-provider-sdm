# strongDM SDK for Go

This is the official [strongDM](https://www.strongdm.com/) SDK for the Go programming language. 

Learn more with our [📚strongDM API docs](https://www.strongdm.com/docs/api/) or [📓browse the SDK reference](https://pkg.go.dev/github.com/strongdm/strongdm-sdk-go?tab=doc).

## Installation

```bash
$ go get github.com/strongdm/strongdm-sdk-go
```

## Authentication

If you don't already have them you will need to generate a set of API keys, instructions are here: [API Credentials](https://www.strongdm.com/docs/admin-guide/api-credentials/)

Add the keys as environment variables; the SDK will need to access these keys for every request.
```bash
$ export SDM_API_ACCESS_KEY=<YOUR ACCESS KEY>
$ export SDM_API_SECRET_KEY=<YOUR SECRET KEY>
```

## List Users
The following code lists all registered users:

```go
func main() {
	log.SetFlags(0)
	accessKey := os.Getenv("SDM_API_ACCESS_KEY")
	secretKey := os.Getenv("SDM_API_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		log.Fatal("SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY must be provided")
	}

	client, err := sdm.New(accessKey, secretKey)
	if err != nil {
		log.Fatal("failed to create strongDM client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	users, err := client.Accounts().List(ctx, "")
	if err != nil {
		log.Fatal("failed to query accounts:", err)
	}
	for users.Next() {
		user := users.Value()
		fmt.Println(user)
	}
	if err := users.Err(); err != nil {
		log.Fatal("error while iterating users:", err)
	}
}
```

## Useful Links

* Documentation:  [sdm package · pkg.go.dev](https://pkg.go.dev/github.com/strongdm/strongdm-sdk-go?tab=doc)
* Examples: [GitHub - strongdm/strongdm-sdk-go-examples](https://github.com/strongdm/strongdm-sdk-go-examples)
	1. [Managing Resources](https://github.com/strongdm/strongdm-sdk-go-examples/tree/master/1_managing_resources)
	1. [Managing Accounts](https://github.com/strongdm/strongdm-sdk-go-examples/tree/master/2_managing_accounts)
	1. [Managing Roles](https://github.com/strongdm/strongdm-sdk-go-examples/tree/master/3_managing_roles)
	1. [Managing Gateways](https://github.com/strongdm/strongdm-sdk-go-examples/tree/master/4_managing_gateways)
   
## License

[Apache 2](https://github.com/strongdm/strongdm-sdk-go/blob/master/LICENSE)

## Contributing 

Currently, we are not accepting pull requests directly to this repository, but our users are some of the most resourceful and ambitious folks out there. So, if you have something to contribute, find a bug, or just want to give us some feedback, please email <support@strongdm.com>.