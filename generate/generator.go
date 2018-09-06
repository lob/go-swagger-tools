//go:generate go-bindata -o "$PKG_DIR/bindata.go" -prefix $PKG_DIR -ignore .+go$ -nomemcopy -nometadata -pkg $PKG_NAME $PKG_DIR
//go:generate swagger generate spec -o $PKG_DIR/swagger.json -b $SOURCE_PKG --scan-models
//go:generate go-bindata -o "$PKG_DIR/bindata.go" -prefix $PKG_DIR -ignore .+go$ -nomemcopy -nometadata -pkg $PKG_NAME $PKG_DIR

// Package generate provides go:generate shorthand for building swagger.json.
//
// Invoking go generate on this package will create a swagger specification and
// corresponding bindata.go file for your project.
//
// There are three environment variables used to pass "parameters" to the generator:
//
// * PKG_DIR defines the absolute path where the resulting files will be placed.
//
// * PKG_NAME specifies the name for the resulting Go package where `bindata.go` will be placed.
//
// * SOURCE_PKG is the package name where `go-swagger` will begin searching for swagger definitions. This should be in the form of a Go import path.
//
// Example:
//
//   $ PKG_NAME=data SOURCE_PKG=github.com/lob/tracking-api/cmd/serve PKG_DIR=$(PWD)/data go generate github.com/lob/go-swagger-tools/generate
//
// This generate relies on go-swagger and go-bindata; they must be available on the path.
//
package generate
