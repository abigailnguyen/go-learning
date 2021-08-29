## To run the tests

```bash
go test -v
```

```bash
go test isvalidBST.go isvalidBST_test.go -v
```

```bash
cd memoization
go run main.go
```

Troubleshooting:

1. When getting the error message compile version of golang does not match go tool version, you can set the global golang version to be same as the go tools used by VS with `asdf global golang <version>`
2. Inside VS Code, search for command `Go: Install/Update tools` and update all the tools
3. In the folder you can also set `.tool-versions` file with `golang 1.16.6` to match with the global golang version
