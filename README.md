# Go Swagger Tools

This repository contains some tools to make integrating Swagger with Go services more straight-forward. Specifically, it has a utility for generating swagger specifications and HTTP middleware for serving the spec + human-readable documentation.

## Generating Swagger Specifications

Running `go generate` on the `generate` sub-package will create a `swagger.json` and `bindata.go` file based on swagger [comment docs](https://goswagger.io/use/spec.html. This relies on [go-swagger](https://goswagger.io) and [go-bindata](https://github.com/jteeuwen/go-bindata); the consumer is responsible for ensuring both utilities are available.

There are three environment variables used to pass "parameters" to the generator:

* `PKG_DIR`
    The absolute path where the resulting files will be placed.
* `PKG_NAME`
    The name for the resulting Go package where `bindata.go` will be placed.
* `SOURCE_PKG`
    The package name where `go-swagger` will begin searching for swagger definitions. This should be in the form of a Go import path.

Invoke `go generate` on the `github.com/lob/swagger/generate` package with these values set in the environment. An example invocation might look like:

```sh
$ PKG_NAME=data PKG_DIR=$(PWD)/data SOURCE_PKG=github.com/lob/tracking-api/cmd/serve go generate github.com/lob/go-swagger-tools/generate
```

## HTTP Middleware

This package includes `WithSwagger` middleware which can be used to serve the generated Swagger specification and HTML documentation with your service. This is stock `net/http` middleware which takes a `SwaggerOpts` struct and a handler to wrap. By default it serves the specification at `/swagger.json` and the docs as `/docs`.

For example,

```go
srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.Config.ServerPort),
		Handler: swagger.WithSwagger(swagger.SwaggerOpts{
			SwaggerSpec: swaggerJSON,
			Title:       "Tracking API",
		}, e),
    }
```

Run `godoc github.com/lob/go-swagger-tools` for full documentation.
