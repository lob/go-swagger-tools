package main

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
	flags "github.com/jessevdk/go-flags"
	bindata "github.com/jteeuwen/go-bindata"
)

var gofiles *regexp.Regexp

var opts struct {
	OutputPackage flags.Filename `long:"package" short:"p" description:"destination package" env:"GOPACKAGE"`
	Input         flags.Filename `long:"input" short:"i" description:"go package to use as input" required:"t"`
}

func init() {
	gofiles = regexp.MustCompile(`.+\.go$`)
}

func main() {

	pwd, _ := os.Getwd()
	args, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	bdConfig := bindata.Config{
		Package:    string(opts.OutputPackage),
		Output:     path.Join(pwd, "bindata.go"),
		Prefix:     pwd,
		NoMemCopy:  true,
		NoMetadata: true,
		Ignore:     []*regexp.Regexp{gofiles},
		Input: []bindata.InputConfig{
			bindata.InputConfig{
				Path:      pwd,
				Recursive: true,
			},
		},
	}
	// generate bindata to bootstrap, if necessary
	err = bindata.Translate(&bdConfig)
	if err != nil {
		fmt.Println("Error bootstrapping bindata.go: ", err)
	}

	// generate the swagger specification
	outputFile := flags.Filename(path.Join(pwd, "swagger.json"))
	fmt.Println("Generate swagger.json...")
	spec := generate.SpecFile{
		ScanModels: true,
		BasePath:   string(opts.Input),
		Compact:    false,
		Output:     outputFile,
		Input:      opts.Input,
	}
	err = spec.Execute(args)
	if err != nil {
		// Note that go-swagger has the annoying habit of Fatal-ing,
		// so we may have already exit-ed before reaching this point.
		fmt.Println("Error generating swagger.json: ", err)
	}

	// regenerate bindata.go
	fmt.Println("Generating bindata.go...")
	err = bindata.Translate(&bdConfig)
	if err != nil {
		fmt.Println("Error generating bindata.go: ", err)
	}
}
