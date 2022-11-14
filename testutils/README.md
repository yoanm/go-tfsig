# testutils

## Functions

### func [EnsureBlockFileEqualsGoldenFile](./terraform_equals.go#L52)

`func EnsureBlockFileEqualsGoldenFile(block *hclwrite.Block, goldenFile string) error`

EnsureBlockFileEqualsGoldenFile checks that provided `hclwrite.Block` content is equals to the content of
the provided golden file.

### func [EnsureFileContentEquals](./terraform_equals.go#L41)

`func EnsureFileContentEquals(file *hclwrite.File, expected string) error`

EnsureFileContentEquals checks that provided `hclwrite.File` content is equals to the expected string.

### func [EnsureFileEqualsGoldenFile](./terraform_equals.go#L64)

`func EnsureFileEqualsGoldenFile(file *hclwrite.File, goldenFile string) error`

EnsureFileEqualsGoldenFile checks that provided `hclwrite.File` content is equals to the content of
the provided golden file.

### func [ExpectPanic](./panic.go#L8)

`func ExpectPanic(t *testing.T, tcname string, callback func(), expectedError string)`

ExpectPanic executes provided 'fn' function and check that:
- `panic(...)` has been called
- related error is the expected one.

### func [LoadGoldenFile](./terraform_equals.go#L24)

`func LoadGoldenFile(filename string) (*string, error)`

LoadGoldenFile loads the golden file filename located under 'testdata' directory

It takes care of suffixing the filename with ".golden.tf".

## Types

### type [ExpectedMismatchError](./terraform_equals.go#L12)

`type ExpectedMismatchError struct { ... }`

#### func (ExpectedMismatchError) [Error](./terraform_equals.go#L17)

`func (e ExpectedMismatchError) Error() string`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
