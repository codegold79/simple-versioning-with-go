# Example of how Package Versions Can Work in Go

Notice the following:

- The greeting package is called `greeting` in both the `v1` and `v2` directories.
- The greeting package is NOT called `v1` or `v2` like how it is in this [Kubernetes example](https://github.com/kubernetes/api/blob/master/rbac/v1/types.go).
- Note that I have a single `go.mod` file in the root directory, where `main.go` is.
- The first line in the `go.mod` file is `module github.com/codegold79/simple-versioning-with-go`. That sets the namespace for all the packages in the repo.
- I ran `go mod tidy` to bring in dependencies (probably not needed, but good practice).
- Run `go run main.go` and see that I can use whatever package I imported using `greeting`.

Try it out yourself. Change the import path to V2 and see the greeting change. You didn't need to reset all the imported package aliases and update them all in the code.
