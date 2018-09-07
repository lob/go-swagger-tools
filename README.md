# Go Swagger Tools

This repository contains some tools to make integrating Swagger with Go services more straight-forward. Specifically, it has a utility for generating swagger specifications and HTTP middleware for serving the spec + human-readable documentation.

## Generating Swagger Specifications

This repository contains a helper, `mkswagger`, which generates `swagger.json` and `bindata.go` file based on swagger [comment docs](https://goswagger.io/use/spec.html).

This can be installed via `go get`:

```sh
$ go get github.com/lob/go-swagger-tools/cmd/mkswagger
```

Once installed it can be used with a `//go:generate` comment:

```go
//go:generate mkswagger -i github.com/lob/tracking-api/cmd/serve
package swagger
```

This will generate `swagger.json` and `bindata.go` in the same directory and package as its invocation (in this example the `swagger` package). The specification will be generating by loading the input packge (`-i`) and walking the source tree.

Invoke `mkswagger -help` for a complete description of options.

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
