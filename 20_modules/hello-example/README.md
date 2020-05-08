# Go modules
## Commands
`go test`: Executes a test's module with it's dependencies

`go mod init example.org/hello`: Creates a module named 'example.org/hello' in the current folder, it generates `go.mod` file.

`cat go.mod`: File that contains the name of the module and the required dependencies.

`cat go.sum`: File that contains the checksum of the required modules. You should use it to check if there's any changes in the dependencies.

`go list -m all`: List all available modules.

`go list -m --versions rsc.io/sampler`: List all available specific module's versions.

`go get rsc.io/sampler@v1.3.1`: It gets the specific module's version, ans set it in the current module.
