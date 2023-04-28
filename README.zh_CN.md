# Errorcode 是一个附带错误码的error生成工具

源代码基于  https://github.com/golang/tools/tree/master/cmd/stringer ，并在此基础上，增加了根据 error 获取错误码的功能。

## 安装

```
go install github.com/fengxuway/errorcode@latest
```

## 用法


命令行帮助文档

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

在 `example/example.go` 中，定义了两个枚举类型 `Err` 和 `Err2`，使用命令：

```
errcode -type=Err,Err2 --linecomment -unknowncode=-1 -codefunc=Code
```
即可生成源代码文件 `err_errcode.go`

PS: 文件总是以 `*_errcode.go` 作为后缀
