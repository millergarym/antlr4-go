# Antlr4 Go runtime

The intention of this repository is to make it easier to get up a running with Antlr4 in Go.
This is Go gettable and contains a build Antlr4 jar file, an example grammar, Go test file and of course the runtime source.
The test file contain a go:generate comment, so the go tool can be used to call the antlr4 tool.

## Getting started

```
mkdir antlr4test
cd antlr4test
mkdir src
export GOPATH=`pwd`
go get github.com/millergarym/antlr4-go
cd src/github.com/millergarym/antlr4-go/examples/expr
go generate
go test
```

The `go generate` command read the source and executes commands specified by in `//go:generate xxx` lines.
The example has such a line, which calls the Antlr4 tool and then `sed` to replace the import statement in the generated code.

## Issues
This only works on OSX as the `sed` for OSX is different.
To get this to work in Linux or Windows, change the sed command or edit the generate files manually.

## Notes

Go runtime package for Antlr4. Also contains Antlr4 jar file and examples.

This is simply a split of the Go runtime from the main github repository.
See https://help.github.com/articles/splitting-a-subfolder-out-into-a-new-repository/ regarding split a repo
