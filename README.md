# Custom linter

## expectfirst

For checking argument order of assertion in test files. The expect argument should be first replaced and then actual argument.

For example:

```go
require.Equal(t, expectNum, num) // O
require.Equal(t, num, expectNum) // X
```

### Usage

1. Run `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2` to build golangci-lint locally
2. Clone this repo
3. Run `make build` to build custom linter (plugin)
4. Copy `.golingci.yml` to your project root
5. Ready to GO!

### Development

Create a new linter:

1. Add rule to `pkg/analyzer/`
2. Add test to `testdata/` (update vendor if need)
3. Run `make test`
4. Update config in `.golangci.yml`
