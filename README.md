# Errorcode is an error generation tool with error codes

The source code is based on https://github.com/golang/tools/tree/master/cmd/stringer, and on this basis, the function of obtaining the error code according to error is added.

## Installation

```
go install github.com/fengxuway/errorcode@latest
```

## Usage

Here is help documention

```
$ errcode  -h                                                         
Usage of errcode:
        errcode [flags] -type T [directory]
        errcode [flags] -type T files... # Must be a single package
For more information, see:
        https://pkg.go.dev/golang.org/x/tools/cmd/errcode
Flags:
  -codefunc string
        error code function name (default "Code")
  -linecomment
        use line comment text as printed text when present
  -output string
        output file name; default srcdir/<type>_errcode.go
  -tags string
        comma-separated list of build tags to apply
  -trimprefix prefix
        trim the prefix from the generated constant names
  -type string
        comma-separated list of type names; must be set
  -unknowncode int
        set unknown error code when not match (default -1)
```

In `example/example.go` source file, define `Err` and `Err2` enum for error code, we can command:

```
errcode -type=Err,Err2 --linecomment -unknowncode=-1 -codefunc=Code
```

Then `err_errcode.go` source file is generated!

PS: generated file name always with `_errcode.go` as suffix.
