# go-test-expect-first

This is a VSCode extension for checking argument order of assertion in test files. The expect argument should be first replaced and then actual argument.

For example:

```go
require.Equal(t, expectNum, num) // O
```

```go
require.Equal(t, num, expectNum) // X
```
