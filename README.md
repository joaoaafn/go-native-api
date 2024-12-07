## This is my studies about native methods to implement API.
This example uses the method [net/http](https://pkg.go.dev/net/http@go1.23.4) from default lib in go without any external packages.

I created a fake database to use.

To run the project:
``` go
go run .
```

In Postman run `localhost:8080/user/profile?clientId={USER}` and methods POST or PATCH