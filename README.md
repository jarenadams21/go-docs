# Go(lang) Quick Docs


## Commands
* go help - displays all available go commands
* go mod tidy - add missing and remove unused modules

<br>

## Official Quick Links & Notes

### hello
* Search packages or symbols - https://pkg.go.dev/
* Managing dependencies - https://go.dev/doc/modules/managing-dependencies#naming_module
* Authenticating Modules - https://go.dev/ref/mod#authenticating
<br>

### Create a Go module
* The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.
* Exported names - https://go.dev/tour/basics/3
